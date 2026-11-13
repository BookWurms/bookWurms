// Im using GORM as per discussion with robin, should make it a bit easier to query 
// without constantly and explicitly writing SQL queries every time.
// the setup was a pain but running AutoMigrate to create the table was nice to do.
// see GORMDOC.md if you want to understand this better.
package database

import (
	"time"
	"gorm.io/gorm"
)

type Book struct {
	ID				int64 `gorm:"primaryKey"`
	Title			string `gorm:"type:text;not null"`
	Author			string `gorm:"type:text;not null"`
	Summary 		*string `gorm:"type:text"`
	DescriptionHTML *string `gorm:"column:description_html;type:text"`
	Cover 			*string `gorm:"type:text"`
	Ebook 			string `gorm:"type:text;not null;uniqueIndex"`
	ISBN 			*string `gorm:"column:isbn;type:text"`
	WordCount 		*int32 `gorm:"type:integer"`
	SourceURL 		*string `gorm:"type:text;uniqueIndex"`
	CreatedAt 		time.Time `gorm:"type:timestamptz;not null; default:now()"`
}

func CreateSeededBookTable(db *gorm.DB) error {
	return db.AutoMigrate(&Book{})
}