package models

import (
	"time"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

// Blog
type Blog struct {
	ID        uuid.UUID `form:"id" json:"id" xml:"id" gorm:"primary_key" sql:"type:uuid" name:"ID"`
	Title     string    `json:"created_by" name:"作成者"`
	Content   string
	CreatedAt time.Time  `name:"作成日"`
	UpdatedAt time.Time  `name:"更新日"`
	DeletedAt *time.Time `name:"削除日"`
}

// BlogDB is the implementation of the storage interface for
// Blog.
type BlogDB struct {
	Db *gorm.DB
}

// NewBlogDB creates a new storage type.
func NewBlogDB(db *gorm.DB) *BlogDB {
	return &BlogDB{Db: db}
}

// DB returns the underlying database.
func (m *BlogDB) DB() interface{} {
	return m.Db
}

// BlogStorage represents the storage interface.
type BlogStorage interface {
	DB() interface{}
}
