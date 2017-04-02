// Code generated by goagen v1.1.0, command line:
// $ goagen
// --design=github.com/tikasan/eventory-goa/design
// --out=$(GOPATH)
// --version=v1.1.0-dirty
//
// API "eventory": Model Helpers
//
// The content of this file is auto-generated, DO NOT MODIFY

package models

import (
	"../app"
	"github.com/goadesign/goa"
	"github.com/jinzhu/gorm"
	"golang.org/x/net/context"
	"time"
)

// MediaType Retrieval Functions

// ListEvent returns an array of view: default.
func (m *EventDB) ListEvent(ctx context.Context, prefID int) []*app.Event {
	defer goa.MeasureSince([]string{"goa", "db", "event", "listevent"}, time.Now())

	var native []*Event
	var objs []*app.Event
	err := m.Db.Scopes(EventFilterByPref(prefID, m.Db)).Table(m.TableName()).Find(&native).Error

	if err != nil {
		goa.LogError(ctx, "error listing Event", "error", err.Error())
		return objs
	}

	for _, t := range native {
		objs = append(objs, t.EventToEvent())
	}

	return objs
}

// EventToEvent loads a Event and builds the default view of media type Event.
func (m *Event) EventToEvent() *app.Event {
	event := &app.Event{}
	event.Accepte = m.Accepte
	event.Address = m.Address
	tmp1 := &m.EndAt
	event.EndAt = tmp1.EndAtToEndAt() // %!s(MISSING)
	event.Identifier = m.Identifier
	event.Limits = m.Limits
	tmp2 := &m.StartAt
	event.StartAt = tmp2.StartAtToStartAt() // %!s(MISSING)
	event.URL = m.URL
	event.Wait = m.Wait

	return event
}

// OneEvent loads a Event and builds the default view of media type Event.
func (m *EventDB) OneEvent(ctx context.Context, id int, prefID int) (*app.Event, error) {
	defer goa.MeasureSince([]string{"goa", "db", "event", "oneevent"}, time.Now())

	var native Event
	err := m.Db.Scopes(EventFilterByPref(prefID, m.Db)).Table(m.TableName()).Preload("EventGenres").Preload("UserKeepStatus").Preload("Pref").Where("id = ?", id).Find(&native).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		goa.LogError(ctx, "error getting Event", "error", err.Error())
		return nil, err
	}

	view := *native.EventToEvent()
	return &view, err
}
