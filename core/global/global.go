package global

import (
	"github.com/sirupsen/logrus"
	"gopkg.in/ini.v1"
	"gorm.io/gorm"
)

var DB *gorm.DB
var LOG *logrus.Logger = logrus.New()
var CONFIG *ini.File
