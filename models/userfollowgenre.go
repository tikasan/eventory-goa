// Code generated by goagen v1.1.0, command line:
// $ goagen
// --design=github.com/tikasan/eventory-goa/design
// --out=$(GOPATH)
// --version=v1.1.0-dirty
//
// API "eventory": Models
//
// The content of this file is auto-generated, DO NOT MODIFY

package models

import (
	"github.com/goadesign/goa"
	"github.com/jinzhu/gorm"
	"golang.org/x/net/context"
	"time"
)

// ユーザーのフォロージャンル
type UserFollowGenre struct {
	ID        int        `gorm:"primary_key"` // primary key
	GenreID   int        // Belongs To Genre
	UserID    int        // Belongs To User
	CreatedAt time.Time  // timestamp
	DeletedAt *time.Time // nullable timestamp (soft delete)
	UpdatedAt time.Time  // timestamp
	Genre     Genre
	User      User
}

// TableName overrides the table name settings in Gorm to force a specific table name
// in the database.
func (m UserFollowGenre) TableName() string {
	return "user_follow_genres"

}

// UserFollowGenreDB is the implementation of the storage interface for
// UserFollowGenre.
type UserFollowGenreDB struct {
	Db *gorm.DB
}

// NewUserFollowGenreDB creates a new storage type.
func NewUserFollowGenreDB(db *gorm.DB) *UserFollowGenreDB {
	return &UserFollowGenreDB{Db: db}
}

// DB returns the underlying database.
func (m *UserFollowGenreDB) DB() interface{} {
	return m.Db
}

// UserFollowGenreStorage represents the storage interface.
type UserFollowGenreStorage interface {
	DB() interface{}
	List(ctx context.Context) ([]*UserFollowGenre, error)
	Get(ctx context.Context, id int) (*UserFollowGenre, error)
	Add(ctx context.Context, userfollowgenre *UserFollowGenre) error
	Update(ctx context.Context, userfollowgenre *UserFollowGenre) error
	Delete(ctx context.Context, id int) error
}

// TableName overrides the table name settings in Gorm to force a specific table name
// in the database.
func (m *UserFollowGenreDB) TableName() string {
	return "user_follow_genres"

}

// Belongs To Relationships

// UserFollowGenreFilterByGenre is a gorm filter for a Belongs To relationship.
func UserFollowGenreFilterByGenre(genreID int, originaldb *gorm.DB) func(db *gorm.DB) *gorm.DB {

	if genreID > 0 {

		return func(db *gorm.DB) *gorm.DB {
			return db.Where("genre_id = ?", genreID)

		}
	}
	return func(db *gorm.DB) *gorm.DB { return db }
}

// Belongs To Relationships

// UserFollowGenreFilterByUser is a gorm filter for a Belongs To relationship.
func UserFollowGenreFilterByUser(userID int, originaldb *gorm.DB) func(db *gorm.DB) *gorm.DB {

	if userID > 0 {

		return func(db *gorm.DB) *gorm.DB {
			return db.Where("user_id = ?", userID)

		}
	}
	return func(db *gorm.DB) *gorm.DB { return db }
}

// CRUD Functions

// Get returns a single UserFollowGenre as a Database Model
// This is more for use internally, and probably not what you want in  your controllers
func (m *UserFollowGenreDB) Get(ctx context.Context, id int) (*UserFollowGenre, error) {
	defer goa.MeasureSince([]string{"goa", "db", "userFollowGenre", "get"}, time.Now())

	var native UserFollowGenre
	err := m.Db.Table(m.TableName()).Where("id = ?", id).Find(&native).Error
	if err == gorm.ErrRecordNotFound {
		return nil, err
	}

	return &native, err
}

// List returns an array of UserFollowGenre
func (m *UserFollowGenreDB) List(ctx context.Context) ([]*UserFollowGenre, error) {
	defer goa.MeasureSince([]string{"goa", "db", "userFollowGenre", "list"}, time.Now())

	var objs []*UserFollowGenre
	err := m.Db.Table(m.TableName()).Find(&objs).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return objs, nil
}

// Add creates a new record.
func (m *UserFollowGenreDB) Add(ctx context.Context, model *UserFollowGenre) error {
	defer goa.MeasureSince([]string{"goa", "db", "userFollowGenre", "add"}, time.Now())

	err := m.Db.Create(model).Error
	if err != nil {
		goa.LogError(ctx, "error adding UserFollowGenre", "error", err.Error())
		return err
	}

	return nil
}

// Update modifies a single record.
func (m *UserFollowGenreDB) Update(ctx context.Context, model *UserFollowGenre) error {
	defer goa.MeasureSince([]string{"goa", "db", "userFollowGenre", "update"}, time.Now())

	obj, err := m.Get(ctx, model.ID)
	if err != nil {
		goa.LogError(ctx, "error updating UserFollowGenre", "error", err.Error())
		return err
	}
	err = m.Db.Model(obj).Updates(model).Error

	return err
}

// Delete removes a single record.
func (m *UserFollowGenreDB) Delete(ctx context.Context, id int) error {
	defer goa.MeasureSince([]string{"goa", "db", "userFollowGenre", "delete"}, time.Now())

	var obj UserFollowGenre

	err := m.Db.Delete(&obj, id).Error

	if err != nil {
		goa.LogError(ctx, "error deleting UserFollowGenre", "error", err.Error())
		return err
	}

	return nil
}