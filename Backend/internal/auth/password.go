package auth

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"golang.org/x/crypto/argon2"
)

const (
	memoryCost       uint32 = 19 * 1024 // Note: argon2.IDKey() uses KiB
	timeCost         uint32 = 2
	parallelism      uint8  = 1
	saltLength              = 16
	keyLength               = 32
	maxPasswordBytes        = 1024

	/* Limits applied when parsing stored hashes. */
	maxMemoryCost  uint32 = 64 * 1024
	maxTimeCost    uint32 = 10
	maxParallelism uint32 = 4
)

var (
	ErrPasswordTooLong = errors.New("password exceeds maximum length")
	ErrInvalidHash     = errors.New("invalid password hash")
)

/*
HashPassword creates an Argon2id hash with a unique random salt.
The returned string contains the algorithm, parameters, salt, and hash.
*/
func HashPassword(password string) (string, error) {
	if len(password) > maxPasswordBytes {
		return "", ErrPasswordTooLong
	}

	salt := make([]byte, saltLength)
	rand.Read(salt)

	hash := argon2.IDKey(
		[]byte(password),
		salt,
		timeCost,
		memoryCost,
		parallelism,
		keyLength,
	)
	defer clear(hash)

	encodedSalt := base64.RawStdEncoding.EncodeToString(salt)
	encodedHash := base64.RawStdEncoding.EncodeToString(hash)

	/* PHC-style encoded representation */
	encoded := fmt.Sprintf(
		"$argon2id$v=%d$m=%d,t=%d,p=%d$%s$%s",
		argon2.Version,
		memoryCost,
		timeCost,
		parallelism,
		encodedSalt,
		encodedHash,
	)

	return encoded, nil
}

func VerifyPassword(encodedHash, password string) (bool, error) {
	if len(password) > maxPasswordBytes {
		return false, ErrPasswordTooLong
	}

	params, err := decodePasswordHash(encodedHash)
	if err != nil {
		return false, err
	}

	computedHash := argon2.IDKey(
		[]byte(password),
		params.salt,
		params.time,
		params.memory,
		params.parallelism,
		uint32(len(params.hash)),
	)
	defer clear(computedHash)

	return subtle.ConstantTimeCompare(computedHash, params.hash) == 1, nil
}

type hashParams struct {
	memory      uint32
	time        uint32
	parallelism uint8
	salt        []byte
	hash        []byte
}

/*
decodePasswordHash parses and validates the stored hash format.
It also limits parameters to prevent malformed stored data from
requesting excessive memory or computation during verification.
*/
func decodePasswordHash(encoded string) (hashParams, error) {
	parts := strings.Split(encoded, "$")
	if len(parts) != 6 ||
		parts[0] != "" ||
		parts[1] != "argon2id" ||
		!strings.HasPrefix(parts[2], "v=") {
		return hashParams{}, ErrInvalidHash
	}

	version, err := strconv.ParseUint(strings.TrimPrefix(parts[2], "v="), 10, 32)
	if err != nil || version != uint64(argon2.Version) {
		return hashParams{}, ErrInvalidHash
	}

	var memory, time, threads uint64
	seen := make(map[string]bool)

	for _, item := range strings.Split(parts[3], ",") {
		keyValue := strings.SplitN(item, "=", 2)
		if len(keyValue) != 2 || seen[keyValue[0]] {
			return hashParams{}, ErrInvalidHash
		}

		seen[keyValue[0]] = true

		value, err := strconv.ParseUint(keyValue[1], 10, 32)
		if err != nil {
			return hashParams{}, ErrInvalidHash
		}

		switch keyValue[0] {
		case "m":
			memory = value
		case "t":
			time = value
		case "p":
			threads = value
		default:
			return hashParams{}, ErrInvalidHash
		}
	}

	if len(seen) != 3 {
		return hashParams{}, ErrInvalidHash
	}

	/*
		Accept the current OWASP minimum profile and stronger
		configurations, while imposing upper limits.
	*/
	validCost := (memory >= 47*1024 && time >= 1) ||
		(memory >= 19*1024 && time >= 2)

	if !validCost ||
		memory > uint64(maxMemoryCost) ||
		time > uint64(maxTimeCost) ||
		threads < 1 ||
		threads > uint64(maxParallelism) {
		return hashParams{}, ErrInvalidHash
	}

	salt, err := base64.RawStdEncoding.DecodeString(parts[4])
	if err != nil || len(salt) != saltLength {
		return hashParams{}, ErrInvalidHash
	}

	hash, err := base64.RawStdEncoding.DecodeString(parts[5])
	if err != nil || len(hash) != keyLength {
		return hashParams{}, ErrInvalidHash
	}

	return hashParams{
		memory:      uint32(memory),
		time:        uint32(time),
		parallelism: uint8(threads),
		salt:        salt,
		hash:        hash,
	}, nil
}
