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
	"time"

	"../app"
	"github.com/goadesign/goa"
	"github.com/jinzhu/gorm"
	"golang.org/x/net/context"
)

// 都道府県
type Pref struct {
	ID              int     `gorm:"primary_key"` // primary key
	Events          []Event // has many Events
	Name            string
	UserFollowPrefs []UserFollowPref // has many UserFollowPrefs
	CreatedAt       time.Time        // timestamp
	DeletedAt       *time.Time       // nullable timestamp (soft delete)
	UpdatedAt       time.Time        // timestamp
}

// TableName overrides the table name settings in Gorm to force a specific table name
// in the database.
func (m Pref) TableName() string {
	return "prefs"

}

// PrefDB is the implementation of the storage interface for
// Pref.
type PrefDB struct {
	Db *gorm.DB
}

// NewPrefDB creates a new storage type.
func NewPrefDB(db *gorm.DB) *PrefDB {
	return &PrefDB{Db: db}
}

// DB returns the underlying database.
func (m *PrefDB) DB() interface{} {
	return m.Db
}

// PrefStorage represents the storage interface.
type PrefStorage interface {
	DB() interface{}
	List(ctx context.Context) ([]*Pref, error)
	Get(ctx context.Context, id int) (*Pref, error)
	Add(ctx context.Context, pref *Pref) error
	Update(ctx context.Context, pref *Pref) error
	Delete(ctx context.Context, id int) error

	ListPref(ctx context.Context) []*app.Pref
	OnePref(ctx context.Context, id int) (*app.Pref, error)
}

// TableName overrides the table name settings in Gorm to force a specific table name
// in the database.
func (m *PrefDB) TableName() string {
	return "prefs"

}

// CRUD Functions

// Get returns a single Pref as a Database Model
// This is more for use internally, and probably not what you want in  your controllers
func (m *PrefDB) Get(ctx context.Context, id int) (*Pref, error) {
	defer goa.MeasureSince([]string{"goa", "db", "pref", "get"}, time.Now())

	var native Pref
	err := m.Db.Table(m.TableName()).Where("id = ?", id).Find(&native).Error
	if err == gorm.ErrRecordNotFound {
		return nil, err
	}

	return &native, err
}

// List returns an array of Pref
func (m *PrefDB) List(ctx context.Context) ([]*Pref, error) {
	defer goa.MeasureSince([]string{"goa", "db", "pref", "list"}, time.Now())

	var objs []*Pref
	err := m.Db.Table(m.TableName()).Find(&objs).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return objs, nil
}

// Add creates a new record.
func (m *PrefDB) Add(ctx context.Context, model *Pref) error {
	defer goa.MeasureSince([]string{"goa", "db", "pref", "add"}, time.Now())

	err := m.Db.Create(model).Error
	if err != nil {
		goa.LogError(ctx, "error adding Pref", "error", err.Error())
		return err
	}

	return nil
}

// Update modifies a single record.
func (m *PrefDB) Update(ctx context.Context, model *Pref) error {
	defer goa.MeasureSince([]string{"goa", "db", "pref", "update"}, time.Now())

	obj, err := m.Get(ctx, model.ID)
	if err != nil {
		goa.LogError(ctx, "error updating Pref", "error", err.Error())
		return err
	}
	err = m.Db.Model(obj).Updates(model).Error

	return err
}

// Delete removes a single record.
func (m *PrefDB) Delete(ctx context.Context, id int) error {
	defer goa.MeasureSince([]string{"goa", "db", "pref", "delete"}, time.Now())

	var obj Pref

	err := m.Db.Delete(&obj, id).Error

	if err != nil {
		goa.LogError(ctx, "error deleting Pref", "error", err.Error())
		return err
	}

	return nil
}
