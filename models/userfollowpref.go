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

// ユーザーのフォロー都道府県
type UserFollowPref struct {
	ID        int        `gorm:"primary_key"` // primary key
	PrefID    int        // Belongs To Pref
	UserID    int        // Belongs To User
	CreatedAt time.Time  // timestamp
	DeletedAt *time.Time // nullable timestamp (soft delete)
	UpdatedAt time.Time  // timestamp
	Pref      Pref
	User      User
}

// TableName overrides the table name settings in Gorm to force a specific table name
// in the database.
func (m UserFollowPref) TableName() string {
	return "user_follow_prefs"

}

// UserFollowPrefDB is the implementation of the storage interface for
// UserFollowPref.
type UserFollowPrefDB struct {
	Db *gorm.DB
}

// NewUserFollowPrefDB creates a new storage type.
func NewUserFollowPrefDB(db *gorm.DB) *UserFollowPrefDB {
	return &UserFollowPrefDB{Db: db}
}

// DB returns the underlying database.
func (m *UserFollowPrefDB) DB() interface{} {
	return m.Db
}

// UserFollowPrefStorage represents the storage interface.
type UserFollowPrefStorage interface {
	DB() interface{}
	List(ctx context.Context) ([]*UserFollowPref, error)
	Get(ctx context.Context, id int) (*UserFollowPref, error)
	Add(ctx context.Context, userfollowpref *UserFollowPref) error
	Update(ctx context.Context, userfollowpref *UserFollowPref) error
	Delete(ctx context.Context, id int) error
}

// TableName overrides the table name settings in Gorm to force a specific table name
// in the database.
func (m *UserFollowPrefDB) TableName() string {
	return "user_follow_prefs"

}

// Belongs To Relationships

// UserFollowPrefFilterByPref is a gorm filter for a Belongs To relationship.
func UserFollowPrefFilterByPref(prefID int, originaldb *gorm.DB) func(db *gorm.DB) *gorm.DB {

	if prefID > 0 {

		return func(db *gorm.DB) *gorm.DB {
			return db.Where("pref_id = ?", prefID)

		}
	}
	return func(db *gorm.DB) *gorm.DB { return db }
}

// Belongs To Relationships

// UserFollowPrefFilterByUser is a gorm filter for a Belongs To relationship.
func UserFollowPrefFilterByUser(userID int, originaldb *gorm.DB) func(db *gorm.DB) *gorm.DB {

	if userID > 0 {

		return func(db *gorm.DB) *gorm.DB {
			return db.Where("user_id = ?", userID)

		}
	}
	return func(db *gorm.DB) *gorm.DB { return db }
}

// CRUD Functions

// Get returns a single UserFollowPref as a Database Model
// This is more for use internally, and probably not what you want in  your controllers
func (m *UserFollowPrefDB) Get(ctx context.Context, id int) (*UserFollowPref, error) {
	defer goa.MeasureSince([]string{"goa", "db", "userFollowPref", "get"}, time.Now())

	var native UserFollowPref
	err := m.Db.Table(m.TableName()).Where("id = ?", id).Find(&native).Error
	if err == gorm.ErrRecordNotFound {
		return nil, err
	}

	return &native, err
}

// List returns an array of UserFollowPref
func (m *UserFollowPrefDB) List(ctx context.Context) ([]*UserFollowPref, error) {
	defer goa.MeasureSince([]string{"goa", "db", "userFollowPref", "list"}, time.Now())

	var objs []*UserFollowPref
	err := m.Db.Table(m.TableName()).Find(&objs).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return objs, nil
}

// Add creates a new record.
func (m *UserFollowPrefDB) Add(ctx context.Context, model *UserFollowPref) error {
	defer goa.MeasureSince([]string{"goa", "db", "userFollowPref", "add"}, time.Now())

	err := m.Db.Create(model).Error
	if err != nil {
		goa.LogError(ctx, "error adding UserFollowPref", "error", err.Error())
		return err
	}

	return nil
}

// Update modifies a single record.
func (m *UserFollowPrefDB) Update(ctx context.Context, model *UserFollowPref) error {
	defer goa.MeasureSince([]string{"goa", "db", "userFollowPref", "update"}, time.Now())

	obj, err := m.Get(ctx, model.ID)
	if err != nil {
		goa.LogError(ctx, "error updating UserFollowPref", "error", err.Error())
		return err
	}
	err = m.Db.Model(obj).Updates(model).Error

	return err
}

// Delete removes a single record.
func (m *UserFollowPrefDB) Delete(ctx context.Context, id int) error {
	defer goa.MeasureSince([]string{"goa", "db", "userFollowPref", "delete"}, time.Now())

	var obj UserFollowPref

	err := m.Db.Delete(&obj, id).Error

	if err != nil {
		goa.LogError(ctx, "error deleting UserFollowPref", "error", err.Error())
		return err
	}

	return nil
}
