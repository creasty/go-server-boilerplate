package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

const (
	sampleDbContextName = "SampleDB"
)

// SetSampleDBWrapper sets a connection of Sample database to the context
func SetSampleDBWrapper(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set(notificationsDbContextName, db)
		c.Next()
	}
}

// GetSampleDB retrives a connection from the context
func GetSampleDB(c *gin.Context) *gorm.DB {
	v := c.MustGet(sampleDbContextName)

	db, ok := v.(*gorm.DB)
	if !ok {
		panic("Cannot retrive value from context")
	}

	return db
}
