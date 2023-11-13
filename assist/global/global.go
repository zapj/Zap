package global

import (
	"github.com/rs/zerolog"
	"gorm.io/gorm"
)

var DB *gorm.DB
var LOG zerolog.Logger
