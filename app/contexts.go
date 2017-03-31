// Code generated by goagen v1.1.0, command line:
// $ goagen
// --design=github.com/tikasan/eventory-goa/design
// --out=$(GOPATH)
// --version=v1.1.0-dirty
//
// API "eventory": Application Contexts
//
// The content of this file is auto-generated, DO NOT MODIFY

package app

import (
	"github.com/goadesign/goa"
	"golang.org/x/net/context"
	"strconv"
	"unicode/utf8"
)

// KeepEventEventsContext provides the events keep event action context.
type KeepEventEventsContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	EventID int
	IsKeep  bool
}

// NewKeepEventEventsContext parses the incoming request URL and body, performs validations and creates the
// context used by the events controller keep event action.
func NewKeepEventEventsContext(ctx context.Context, service *goa.Service) (*KeepEventEventsContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	rctx := KeepEventEventsContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramEventID := req.Params["eventID"]
	if len(paramEventID) > 0 {
		rawEventID := paramEventID[0]
		if eventID, err2 := strconv.Atoi(rawEventID); err2 == nil {
			rctx.EventID = eventID
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("eventID", rawEventID, "integer"))
		}
	}
	paramIsKeep := req.Params["isKeep"]
	if len(paramIsKeep) == 0 {
		err = goa.MergeErrors(err, goa.MissingParamError("isKeep"))
	} else {
		rawIsKeep := paramIsKeep[0]
		if isKeep, err2 := strconv.ParseBool(rawIsKeep); err2 == nil {
			rctx.IsKeep = isKeep
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("isKeep", rawIsKeep, "boolean"))
		}
	}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *KeepEventEventsContext) OK(resp []byte) error {
	ctx.ResponseData.Header().Set("Content-Type", "text/plain")
	ctx.ResponseData.WriteHeader(200)
	_, err := ctx.ResponseData.Write(resp)
	return err
}

// BadRequest sends a HTTP response with status code 400.
func (ctx *KeepEventEventsContext) BadRequest(r error) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goa.error")
	return ctx.ResponseData.Service.Send(ctx.Context, 400, r)
}

// Unauthorized sends a HTTP response with status code 401.
func (ctx *KeepEventEventsContext) Unauthorized() error {
	ctx.ResponseData.WriteHeader(401)
	return nil
}

// NotFound sends a HTTP response with status code 404.
func (ctx *KeepEventEventsContext) NotFound() error {
	ctx.ResponseData.WriteHeader(404)
	return nil
}

// ListEventsContext provides the events list action context.
type ListEventsContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	ID   string
	Page *int
	Q    string
	Sort string
}

// NewListEventsContext parses the incoming request URL and body, performs validations and creates the
// context used by the events controller list action.
func NewListEventsContext(ctx context.Context, service *goa.Service) (*ListEventsContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	rctx := ListEventsContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramID := req.Params["id"]
	if len(paramID) > 0 {
		rawID := paramID[0]
		rctx.ID = rawID
	}
	paramPage := req.Params["page"]
	if len(paramPage) > 0 {
		rawPage := paramPage[0]
		if page, err2 := strconv.Atoi(rawPage); err2 == nil {
			tmp4 := page
			tmp3 := &tmp4
			rctx.Page = tmp3
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("page", rawPage, "integer"))
		}
		if rctx.Page != nil {
			if *rctx.Page < 0 {
				err = goa.MergeErrors(err, goa.InvalidRangeError(`page`, *rctx.Page, 0, true))
			}
		}
	}
	paramQ := req.Params["q"]
	if len(paramQ) > 0 {
		rawQ := paramQ[0]
		rctx.Q = rawQ
	}
	paramSort := req.Params["sort"]
	if len(paramSort) > 0 {
		rawSort := paramSort[0]
		rctx.Sort = rawSort
	}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *ListEventsContext) OK(r EventCollection) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.event+json; type=collection")
	if r == nil {
		r = EventCollection{}
	}
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// BadRequest sends a HTTP response with status code 400.
func (ctx *ListEventsContext) BadRequest(r error) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goa.error")
	return ctx.ResponseData.Service.Send(ctx.Context, 400, r)
}

// Unauthorized sends a HTTP response with status code 401.
func (ctx *ListEventsContext) Unauthorized() error {
	ctx.ResponseData.WriteHeader(401)
	return nil
}

// CreateGenresContext provides the genres create action context.
type CreateGenresContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	Name string
}

// NewCreateGenresContext parses the incoming request URL and body, performs validations and creates the
// context used by the genres controller create action.
func NewCreateGenresContext(ctx context.Context, service *goa.Service) (*CreateGenresContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	rctx := CreateGenresContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramName := req.Params["name"]
	if len(paramName) == 0 {
		err = goa.MergeErrors(err, goa.MissingParamError("name"))
	} else {
		rawName := paramName[0]
		rctx.Name = rawName
		if utf8.RuneCountInString(rctx.Name) < 1 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`name`, rctx.Name, utf8.RuneCountInString(rctx.Name), 1, true))
		}
		if utf8.RuneCountInString(rctx.Name) > 30 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`name`, rctx.Name, utf8.RuneCountInString(rctx.Name), 30, false))
		}
	}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *CreateGenresContext) OK(r *Genre) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.genre+json")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// BadRequest sends a HTTP response with status code 400.
func (ctx *CreateGenresContext) BadRequest(r error) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goa.error")
	return ctx.ResponseData.Service.Send(ctx.Context, 400, r)
}

// Unauthorized sends a HTTP response with status code 401.
func (ctx *CreateGenresContext) Unauthorized() error {
	ctx.ResponseData.WriteHeader(401)
	return nil
}

// FollowGenreGenresContext provides the genres follow genre action context.
type FollowGenreGenresContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	GenreID int
}

// NewFollowGenreGenresContext parses the incoming request URL and body, performs validations and creates the
// context used by the genres controller follow genre action.
func NewFollowGenreGenresContext(ctx context.Context, service *goa.Service) (*FollowGenreGenresContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	rctx := FollowGenreGenresContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramGenreID := req.Params["genreID"]
	if len(paramGenreID) > 0 {
		rawGenreID := paramGenreID[0]
		if genreID, err2 := strconv.Atoi(rawGenreID); err2 == nil {
			rctx.GenreID = genreID
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("genreID", rawGenreID, "integer"))
		}
	}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *FollowGenreGenresContext) OK(resp []byte) error {
	ctx.ResponseData.Header().Set("Content-Type", "text/plain")
	ctx.ResponseData.WriteHeader(200)
	_, err := ctx.ResponseData.Write(resp)
	return err
}

// BadRequest sends a HTTP response with status code 400.
func (ctx *FollowGenreGenresContext) BadRequest(r error) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goa.error")
	return ctx.ResponseData.Service.Send(ctx.Context, 400, r)
}

// Unauthorized sends a HTTP response with status code 401.
func (ctx *FollowGenreGenresContext) Unauthorized() error {
	ctx.ResponseData.WriteHeader(401)
	return nil
}

// NotFound sends a HTTP response with status code 404.
func (ctx *FollowGenreGenresContext) NotFound() error {
	ctx.ResponseData.WriteHeader(404)
	return nil
}

// ListGenresContext provides the genres list action context.
type ListGenresContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	Q string
}

// NewListGenresContext parses the incoming request URL and body, performs validations and creates the
// context used by the genres controller list action.
func NewListGenresContext(ctx context.Context, service *goa.Service) (*ListGenresContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	rctx := ListGenresContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramQ := req.Params["q"]
	if len(paramQ) > 0 {
		rawQ := paramQ[0]
		rctx.Q = rawQ
		if utf8.RuneCountInString(rctx.Q) < 0 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`q`, rctx.Q, utf8.RuneCountInString(rctx.Q), 0, true))
		}
		if utf8.RuneCountInString(rctx.Q) > 30 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`q`, rctx.Q, utf8.RuneCountInString(rctx.Q), 30, false))
		}
	}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *ListGenresContext) OK(r GenreCollection) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.genre+json; type=collection")
	if r == nil {
		r = GenreCollection{}
	}
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// BadRequest sends a HTTP response with status code 400.
func (ctx *ListGenresContext) BadRequest(r error) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goa.error")
	return ctx.ResponseData.Service.Send(ctx.Context, 400, r)
}

// Unauthorized sends a HTTP response with status code 401.
func (ctx *ListGenresContext) Unauthorized() error {
	ctx.ResponseData.WriteHeader(401)
	return nil
}

// PrefFollowPrefsContext provides the prefs pref follow action context.
type PrefFollowPrefsContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	PrefID int
}

// NewPrefFollowPrefsContext parses the incoming request URL and body, performs validations and creates the
// context used by the prefs controller pref follow action.
func NewPrefFollowPrefsContext(ctx context.Context, service *goa.Service) (*PrefFollowPrefsContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	rctx := PrefFollowPrefsContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramPrefID := req.Params["prefID"]
	if len(paramPrefID) > 0 {
		rawPrefID := paramPrefID[0]
		if prefID, err2 := strconv.Atoi(rawPrefID); err2 == nil {
			rctx.PrefID = prefID
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("prefID", rawPrefID, "integer"))
		}
	}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *PrefFollowPrefsContext) OK(resp []byte) error {
	ctx.ResponseData.Header().Set("Content-Type", "text/plain")
	ctx.ResponseData.WriteHeader(200)
	_, err := ctx.ResponseData.Write(resp)
	return err
}

// BadRequest sends a HTTP response with status code 400.
func (ctx *PrefFollowPrefsContext) BadRequest(r error) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goa.error")
	return ctx.ResponseData.Service.Send(ctx.Context, 400, r)
}

// Unauthorized sends a HTTP response with status code 401.
func (ctx *PrefFollowPrefsContext) Unauthorized() error {
	ctx.ResponseData.WriteHeader(401)
	return nil
}

// NotFound sends a HTTP response with status code 404.
func (ctx *PrefFollowPrefsContext) NotFound() error {
	ctx.ResponseData.WriteHeader(404)
	return nil
}

// AccountCreateUsersContext provides the users account create action context.
type AccountCreateUsersContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	ClientVersion *string
	Email         *string
	Identifier    *string
	Platform      *string
}

// NewAccountCreateUsersContext parses the incoming request URL and body, performs validations and creates the
// context used by the users controller account create action.
func NewAccountCreateUsersContext(ctx context.Context, service *goa.Service) (*AccountCreateUsersContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	rctx := AccountCreateUsersContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramClientVersion := req.Params["client_version"]
	if len(paramClientVersion) > 0 {
		rawClientVersion := paramClientVersion[0]
		rctx.ClientVersion = &rawClientVersion
		if rctx.ClientVersion != nil {
			if utf8.RuneCountInString(*rctx.ClientVersion) < 1 {
				err = goa.MergeErrors(err, goa.InvalidLengthError(`client_version`, *rctx.ClientVersion, utf8.RuneCountInString(*rctx.ClientVersion), 1, true))
			}
		}
		if rctx.ClientVersion != nil {
			if utf8.RuneCountInString(*rctx.ClientVersion) > 10 {
				err = goa.MergeErrors(err, goa.InvalidLengthError(`client_version`, *rctx.ClientVersion, utf8.RuneCountInString(*rctx.ClientVersion), 10, false))
			}
		}
	}
	paramEmail := req.Params["email"]
	if len(paramEmail) > 0 {
		rawEmail := paramEmail[0]
		rctx.Email = &rawEmail
		if rctx.Email != nil {
			if err2 := goa.ValidateFormat(goa.FormatEmail, *rctx.Email); err2 != nil {
				err = goa.MergeErrors(err, goa.InvalidFormatError(`email`, *rctx.Email, goa.FormatEmail, err2))
			}
		}
	}
	paramIdentifier := req.Params["identifier"]
	if len(paramIdentifier) > 0 {
		rawIdentifier := paramIdentifier[0]
		rctx.Identifier = &rawIdentifier
		if rctx.Identifier != nil {
			if ok := goa.ValidatePattern(`(^[a-z0-9]{16}$|^[a-z0-9\-]{32}$)`, *rctx.Identifier); !ok {
				err = goa.MergeErrors(err, goa.InvalidPatternError(`identifier`, *rctx.Identifier, `(^[a-z0-9]{16}$|^[a-z0-9\-]{32}$)`))
			}
		}
	}
	paramPlatform := req.Params["platform"]
	if len(paramPlatform) > 0 {
		rawPlatform := paramPlatform[0]
		rctx.Platform = &rawPlatform
		if rctx.Platform != nil {
			if !(*rctx.Platform == "ios" || *rctx.Platform == "android") {
				err = goa.MergeErrors(err, goa.InvalidEnumValueError(`platform`, *rctx.Platform, []interface{}{"ios", "android"}))
			}
		}
	}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *AccountCreateUsersContext) OK(r *Token) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.token+json")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// BadRequest sends a HTTP response with status code 400.
func (ctx *AccountCreateUsersContext) BadRequest(r error) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goa.error")
	return ctx.ResponseData.Service.Send(ctx.Context, 400, r)
}

// Unauthorized sends a HTTP response with status code 401.
func (ctx *AccountCreateUsersContext) Unauthorized() error {
	ctx.ResponseData.WriteHeader(401)
	return nil
}

// TmpAccountCreateUsersContext provides the users tmp account create action context.
type TmpAccountCreateUsersContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	ClientVersion *string
	Identifier    *string
	Platform      *string
}

// NewTmpAccountCreateUsersContext parses the incoming request URL and body, performs validations and creates the
// context used by the users controller tmp account create action.
func NewTmpAccountCreateUsersContext(ctx context.Context, service *goa.Service) (*TmpAccountCreateUsersContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	rctx := TmpAccountCreateUsersContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramClientVersion := req.Params["client_version"]
	if len(paramClientVersion) > 0 {
		rawClientVersion := paramClientVersion[0]
		rctx.ClientVersion = &rawClientVersion
		if rctx.ClientVersion != nil {
			if utf8.RuneCountInString(*rctx.ClientVersion) < 1 {
				err = goa.MergeErrors(err, goa.InvalidLengthError(`client_version`, *rctx.ClientVersion, utf8.RuneCountInString(*rctx.ClientVersion), 1, true))
			}
		}
		if rctx.ClientVersion != nil {
			if utf8.RuneCountInString(*rctx.ClientVersion) > 10 {
				err = goa.MergeErrors(err, goa.InvalidLengthError(`client_version`, *rctx.ClientVersion, utf8.RuneCountInString(*rctx.ClientVersion), 10, false))
			}
		}
	}
	paramIdentifier := req.Params["identifier"]
	if len(paramIdentifier) > 0 {
		rawIdentifier := paramIdentifier[0]
		rctx.Identifier = &rawIdentifier
		if rctx.Identifier != nil {
			if ok := goa.ValidatePattern(`(^[a-z0-9]{16}$|^[a-z0-9\-]{32}$)`, *rctx.Identifier); !ok {
				err = goa.MergeErrors(err, goa.InvalidPatternError(`identifier`, *rctx.Identifier, `(^[a-z0-9]{16}$|^[a-z0-9\-]{32}$)`))
			}
		}
	}
	paramPlatform := req.Params["platform"]
	if len(paramPlatform) > 0 {
		rawPlatform := paramPlatform[0]
		rctx.Platform = &rawPlatform
		if rctx.Platform != nil {
			if !(*rctx.Platform == "ios" || *rctx.Platform == "android") {
				err = goa.MergeErrors(err, goa.InvalidEnumValueError(`platform`, *rctx.Platform, []interface{}{"ios", "android"}))
			}
		}
	}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *TmpAccountCreateUsersContext) OK(r *Token) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.token+json")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// BadRequest sends a HTTP response with status code 400.
func (ctx *TmpAccountCreateUsersContext) BadRequest(r error) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goa.error")
	return ctx.ResponseData.Service.Send(ctx.Context, 400, r)
}
