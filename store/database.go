package store

import (
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq" // Use with gorm
)

// NewDatabase opens a new database connection for the url
func NewDatabase(url string, log bool) *gorm.DB {
	db, err := gorm.Open("postgres", url)
	if err != nil {
		panic(err)
	}
	db.LogMode(log)

	dbconn := db.DB()

	// health check
	if err := dbconn.Ping(); err != nil {
		panic(err)
	}
	// dbconn.SetMaxIdleConns(10)
	// dbconn.SetMaxOpenConns(100)

	return db
}

// Transaction executes the function in a transactional block,
// and rollbacks if an error is returned
func Transaction(db *gorm.DB, fn func(*gorm.DB) error) (err error) {
	tx := db.Begin()
	err = fn(tx)

	if err == nil {
		tx.Commit()
	} else {
		tx.Rollback()
	}

	return
}
