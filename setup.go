package main

import (
	"git.si-media.biz/takutoarao/start/db"
	"github.com/goadesign/goa"
	"github.com/jinzhu/gorm"
)

var database *gorm.DB

var service *goa.Service

func setup() {
	database = db.ConnectDB()
}
