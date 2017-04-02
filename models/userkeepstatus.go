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

// ユーザーのキープ状態
type UserKeepStatus struct {
	ID             int `gorm:"primary_key"` // primary key
	BatchProcessed bool
	EventID        int // Belongs To Event
	Status         string
	UserID         int        // Belongs To User
	CreatedAt      time.Time  // timestamp
	DeletedAt      *time.Time // nullable timestamp (soft delete)
	UpdatedAt      time.Time  // timestamp
	Event          Event
	User           User
}

// TableName overrides the table name settings in Gorm to force a specific table name
// in the database.
func (m UserKeepStatus) TableName() string {
	return "user_keep_status"

}

// UserKeepStatusDB is the implementation of the storage interface for
// UserKeepStatus.
type UserKeepStatusDB struct {
	Db *gorm.DB
}

// NewUserKeepStatusDB creates a new storage type.
func NewUserKeepStatusDB(db *gorm.DB) *UserKeepStatusDB {
	return &UserKeepStatusDB{Db: db}
}

// DB returns the underlying database.
func (m *UserKeepStatusDB) DB() interface{} {
	return m.Db
}

// UserKeepStatusStorage represents the storage interface.
type UserKeepStatusStorage interface {
	DB() interface{}
	List(ctx context.Context) ([]*UserKeepStatus, error)
	Get(ctx context.Context, id int) (*UserKeepStatus, error)
	Add(ctx context.Context, userkeepstatus *UserKeepStatus) error
	Update(ctx context.Context, userkeepstatus *UserKeepStatus) error
	Delete(ctx context.Context, id int) error
}

// TableName overrides the table name settings in Gorm to force a specific table name
// in the database.
func (m *UserKeepStatusDB) TableName() string {
	return "user_keep_status"

}

// Belongs To Relationships

// UserKeepStatusFilterByEvent is a gorm filter for a Belongs To relationship.
func UserKeepStatusFilterByEvent(eventID int, originaldb *gorm.DB) func(db *gorm.DB) *gorm.DB {

	if eventID > 0 {

		return func(db *gorm.DB) *gorm.DB {
			return db.Where("event_id = ?", eventID)

		}
	}
	return func(db *gorm.DB) *gorm.DB { return db }
}

// Belongs To Relationships

// UserKeepStatusFilterByUser is a gorm filter for a Belongs To relationship.
func UserKeepStatusFilterByUser(userID int, originaldb *gorm.DB) func(db *gorm.DB) *gorm.DB {

	if userID > 0 {

		return func(db *gorm.DB) *gorm.DB {
			return db.Where("user_id = ?", userID)

		}
	}
	return func(db *gorm.DB) *gorm.DB { return db }
}

// CRUD Functions

// Get returns a single UserKeepStatus as a Database Model
// This is more for use internally, and probably not what you want in  your controllers
func (m *UserKeepStatusDB) Get(ctx context.Context, id int) (*UserKeepStatus, error) {
	defer goa.MeasureSince([]string{"goa", "db", "userKeepStatus", "get"}, time.Now())

	var native UserKeepStatus
	err := m.Db.Table(m.TableName()).Where("id = ?", id).Find(&native).Error
	if err == gorm.ErrRecordNotFound {
		return nil, err
	}

	return &native, err
}

// List returns an array of UserKeepStatus
func (m *UserKeepStatusDB) List(ctx context.Context) ([]*UserKeepStatus, error) {
	defer goa.MeasureSince([]string{"goa", "db", "userKeepStatus", "list"}, time.Now())

	var objs []*UserKeepStatus
	err := m.Db.Table(m.TableName()).Find(&objs).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return objs, nil
}

// Add creates a new record.
func (m *UserKeepStatusDB) Add(ctx context.Context, model *UserKeepStatus) error {
	defer goa.MeasureSince([]string{"goa", "db", "userKeepStatus", "add"}, time.Now())

	err := m.Db.Create(model).Error
	if err != nil {
		goa.LogError(ctx, "error adding UserKeepStatus", "error", err.Error())
		return err
	}

	return nil
}

// Update modifies a single record.
func (m *UserKeepStatusDB) Update(ctx context.Context, model *UserKeepStatus) error {
	defer goa.MeasureSince([]string{"goa", "db", "userKeepStatus", "update"}, time.Now())

	obj, err := m.Get(ctx, model.ID)
	if err != nil {
		goa.LogError(ctx, "error updating UserKeepStatus", "error", err.Error())
		return err
	}
	err = m.Db.Model(obj).Updates(model).Error

	return err
}

// Delete removes a single record.
func (m *UserKeepStatusDB) Delete(ctx context.Context, id int) error {
	defer goa.MeasureSince([]string{"goa", "db", "userKeepStatus", "delete"}, time.Now())

	var obj UserKeepStatus

	err := m.Db.Delete(&obj, id).Error

	if err != nil {
		goa.LogError(ctx, "error deleting UserKeepStatus", "error", err.Error())
		return err
	}

	return nil
}
