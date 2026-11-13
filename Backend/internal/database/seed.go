// go's fs package works with things like folders or zip files, so we can use this to read
// the epub and extract the metadata content.
package database

import (
	"archive/zip"
	"encoding/xml" //epub is formatted as xml
	"io/fs"
	"path/filepath"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// OPF = Open Packaging Format for epub books
type BookMetadata struct {
	Metadata struct {
		Titles      []string `xml:"title"`
		Author      string   `xml:"creator"`
		Language    string   `xml:"language"`
		SourceURL   string   `xml:"identifier"`
		Description string   `xml:"description"`
	} `xml:"metadata"`
}

func ParseOPF(data []byte) (BookMetadata, error) {
	var meta BookMetadata
	// xml.Unmarshal will take the raw text from xml and fill
	// a go struct with the fields and elements, by taking and tagging
	// matching copies.
	err := xml.Unmarshal(data, &meta)
	return meta, err
}

func ReadOPF(path string) ([]byte, error) {
	read, err := zip.OpenReader(path)
	if err != nil {
		return nil, err
	}
	defer read.Close()

	return fs.ReadFile(read, "epub/content.opf")
}

func SeedBook(db *gorm.DB, path string) error {
	data, err := ReadOPF(path)
	if err != nil {
		return err
	}

	meta, err := ParseOPF(data)
	if err != nil {
		return err
	}

	book := Book{
		Title:           meta.Metadata.Titles[0],
		Author:          meta.Metadata.Author,
		DescriptionHTML: &meta.Metadata.Description,
		Ebook:           filepath.Base(path),
		SourceURL:       &meta.Metadata.SourceURL,
	}

	return db.Clauses(clause.OnConflict{DoNothing: true}).Create(&book).Error
}

func SeedAllBooks(db *gorm.DB, dir string) error {
	paths, err := filepath.Glob(filepath.Join(dir, "*.epub")) // get all files in a directory that contain .epub
	if err != nil {
		return err
	}

	for _, path := range paths {
		if err := SeedBook(db, path); err != nil {
			return err
		}
	}
	return nil
}
