
# GORM Resources

Gorm is a "Object-Relational Mapping (ORM) library, so its a library where we define our SQL tables as structs, and let the library  functions handle the abstraction into SQL.

### How to install GORM package
Use the following commands in your folder to add the package in the go mod file:
- ```go get gorm.io/gorm```
- ```go get gorm.io/driver/postgres```

### Importing the driver
- use the following blockout to import the driver wherever youre actually using gorm.
``` 
import (
	"gorm.io/gorm" <--------------
	etc...
)
```

## Creating a Table
So with gorm, instead of emplacing direct sql queries into the code like so:
```
const createSeededBooksTable = `
CREATE TABLE IF NOT EXISTS books (
	id					BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
	title				TEXT NOT NULL,
	author				TEXT NOT NULL,
	etc...
```
We can define it as a struct like this:
```
type Book struct {
	ID				int64 `gorm:"primaryKey;type:bigint GENERATED ALWAYS AS IDENTITY`
	Title			string `gorm:"type:text;not null"`
	author			string `gorm:"type:text;not null"`
	etc...
}
```

But this alone in either case will not create a table for us magically. we can use gorms automigrate function on our struct, which transliterates it into SQL for postgres.

```
func CreateTable(db *gorm.DB) error {
	return db.AutoMigrate(&Book{})
}
```