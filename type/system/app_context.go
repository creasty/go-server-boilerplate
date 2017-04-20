package system

import (
	"github.com/jinzhu/gorm"
)

// AppContext represents a context of application
type AppContext struct {
	Config   *Config
	SampleDB *gorm.DB
}
