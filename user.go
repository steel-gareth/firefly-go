// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package firefly

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/steel-gareth/firefly-go/internal/apijson"
	"github.com/steel-gareth/firefly-go/internal/apiquery"
	shimjson "github.com/steel-gareth/firefly-go/internal/encoding/json"
	"github.com/steel-gareth/firefly-go/internal/requestconfig"
	"github.com/steel-gareth/firefly-go/option"
	"github.com/steel-gareth/firefly-go/packages/param"
	"github.com/steel-gareth/firefly-go/packages/respjson"
)

// Use these endpoints to manage the users registered within Firefly III. You need
// to have the &quot;owner&quot; role to access these endpoints.
//
// UserService contains methods and other services that help with interacting with
// the firefly API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewUserService] method instead.
type UserService struct {
	options []option.RequestOption
}

// NewUserService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewUserService(opts ...option.RequestOption) (r UserService) {
	r = UserService{}
	r.options = opts
	return
}

// Creates a new user. The data required can be submitted as a JSON body or as a
// list of parameters. The user will be given a random password, which they can
// reset using the "forgot password" function.
func (r *UserService) New(ctx context.Context, params UserNewParams, opts ...option.RequestOption) (res *UserSingle, err error) {
	if !param.IsOmitted(params.XTraceID) {
		opts = append(opts, option.WithHeader("X-Trace-Id", fmt.Sprintf("%v", params.XTraceID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/vnd.api+json")}, opts...)
	path := "v1/users"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Gets all info of a single user.
func (r *UserService) Get(ctx context.Context, id string, query UserGetParams, opts ...option.RequestOption) (res *UserSingle, err error) {
	if !param.IsOmitted(query.XTraceID) {
		opts = append(opts, option.WithHeader("X-Trace-Id", fmt.Sprintf("%v", query.XTraceID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/vnd.api+json")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/users/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Update existing user.
func (r *UserService) Update(ctx context.Context, id string, params UserUpdateParams, opts ...option.RequestOption) (res *UserSingle, err error) {
	if !param.IsOmitted(params.XTraceID) {
		opts = append(opts, option.WithHeader("X-Trace-Id", fmt.Sprintf("%v", params.XTraceID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/vnd.api+json")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/users/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return res, err
}

// List all the users in this instance of Firefly III.
func (r *UserService) List(ctx context.Context, params UserListParams, opts ...option.RequestOption) (res *UserListResponse, err error) {
	if !param.IsOmitted(params.XTraceID) {
		opts = append(opts, option.WithHeader("X-Trace-Id", fmt.Sprintf("%v", params.XTraceID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/vnd.api+json")}, opts...)
	path := "v1/users"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

// Delete a user. You cannot delete the user you're authenticated with. This cannot
// be undone. Be careful.
func (r *UserService) Delete(ctx context.Context, id string, body UserDeleteParams, opts ...option.RequestOption) (err error) {
	if !param.IsOmitted(body.XTraceID) {
		opts = append(opts, option.WithHeader("X-Trace-Id", fmt.Sprintf("%v", body.XTraceID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return err
	}
	path := fmt.Sprintf("v1/users/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

type User struct {
	// The new users email address.
	Email string `json:"email" api:"required" format:"email"`
	// Boolean to indicate if the user is blocked.
	Blocked bool `json:"blocked"`
	// If you say the user must be blocked, this will be the reason code.
	//
	// Any of "email_changed".
	BlockedCode UserBlockedCode `json:"blocked_code" api:"nullable"`
	CreatedAt   time.Time       `json:"created_at" format:"date-time"`
	// Role for the user. Can be empty or omitted.
	//
	// Any of "owner", "demo".
	Role      UserRole  `json:"role" api:"nullable"`
	UpdatedAt time.Time `json:"updated_at" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Email       respjson.Field
		Blocked     respjson.Field
		BlockedCode respjson.Field
		CreatedAt   respjson.Field
		Role        respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r User) RawJSON() string { return r.JSON.raw }
func (r *User) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this User to a UserParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// UserParam.Overrides()
func (r User) ToParam() UserParam {
	return param.Override[UserParam](json.RawMessage(r.RawJSON()))
}

// If you say the user must be blocked, this will be the reason code.
type UserBlockedCode string

const (
	UserBlockedCodeEmailChanged UserBlockedCode = "email_changed"
)

// Role for the user. Can be empty or omitted.
type UserRole string

const (
	UserRoleOwner UserRole = "owner"
	UserRoleDemo  UserRole = "demo"
)

// The property Email is required.
type UserParam struct {
	// The new users email address.
	Email string `json:"email" api:"required" format:"email"`
	// Boolean to indicate if the user is blocked.
	Blocked param.Opt[bool] `json:"blocked,omitzero"`
	// If you say the user must be blocked, this will be the reason code.
	//
	// Any of "email_changed".
	BlockedCode UserBlockedCode `json:"blocked_code,omitzero"`
	// Role for the user. Can be empty or omitted.
	//
	// Any of "owner", "demo".
	Role UserRole `json:"role,omitzero"`
	paramObj
}

func (r UserParam) MarshalJSON() (data []byte, err error) {
	type shadow UserParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UserParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserListResponse struct {
	Data  []UserRead `json:"data" api:"required"`
	Links PageLink   `json:"links" api:"required"`
	Meta  Meta       `json:"meta" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Links       respjson.Field
		Meta        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UserListResponse) RawJSON() string { return r.JSON.raw }
func (r *UserListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserNewParams struct {
	User     UserParam
	XTraceID param.Opt[string] `header:"X-Trace-Id,omitzero" format:"uuid" json:"-"`
	paramObj
}

func (r UserNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.User)
}
func (r *UserNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserGetParams struct {
	XTraceID param.Opt[string] `header:"X-Trace-Id,omitzero" format:"uuid" json:"-"`
	paramObj
}

type UserUpdateParams struct {
	User     UserParam
	XTraceID param.Opt[string] `header:"X-Trace-Id,omitzero" format:"uuid" json:"-"`
	paramObj
}

func (r UserUpdateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.User)
}
func (r *UserUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserListParams struct {
	// Number of items per page. The default pagination is per 50 items.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Page number. The default pagination is per 50 items.
	Page     param.Opt[int64]  `query:"page,omitzero" json:"-"`
	XTraceID param.Opt[string] `header:"X-Trace-Id,omitzero" format:"uuid" json:"-"`
	paramObj
}

// URLQuery serializes [UserListParams]'s query parameters as `url.Values`.
func (r UserListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type UserDeleteParams struct {
	XTraceID param.Opt[string] `header:"X-Trace-Id,omitzero" format:"uuid" json:"-"`
	paramObj
}
