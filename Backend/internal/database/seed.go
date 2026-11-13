// go's fs package works with things like folders or zip files, so we can use this to read
// the epub and extract the metadata content.
package database

import (
	"archive/zip"
	"io/fs"
)

// OPF = Open Packaging Format for epub books
func ReadOPF(path string) ([]byte, error) {
	read, err := zip.OpenReader(path)
	if err != nil {
		return nil, err
	}
	defer read.Close()

	return fs.ReadFile(read, "epub/content.opf")

}
