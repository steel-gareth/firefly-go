// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package emceesprodtesting5

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/stainless-sdks/emcees-prod-testing-5-go/internal/apijson"
	"github.com/stainless-sdks/emcees-prod-testing-5-go/internal/apiquery"
	"github.com/stainless-sdks/emcees-prod-testing-5-go/internal/requestconfig"
	"github.com/stainless-sdks/emcees-prod-testing-5-go/option"
	"github.com/stainless-sdks/emcees-prod-testing-5-go/packages/param"
	"github.com/stainless-sdks/emcees-prod-testing-5-go/packages/respjson"
)

// User groups are the objects around which &quot;financial administrations&quot;
// are built.
//
// UserGroupService contains methods and other services that help with interacting
// with the emcees-prod-testing-5 API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewUserGroupService] method instead.
type UserGroupService struct {
	options []option.RequestOption
}

// NewUserGroupService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewUserGroupService(opts ...option.RequestOption) (r UserGroupService) {
	r = UserGroupService{}
	r.options = opts
	return
}

// Returns a single user group by its ID.
func (r *UserGroupService) Get(ctx context.Context, id string, query UserGroupGetParams, opts ...option.RequestOption) (res *UserGroupSingle, err error) {
	if !param.IsOmitted(query.XTraceID) {
		opts = append(opts, option.WithHeader("X-Trace-Id", fmt.Sprintf("%v", query.XTraceID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/vnd.api+json")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/user-groups/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Used to update a single user group. The available fields are still limited.
func (r *UserGroupService) Update(ctx context.Context, id string, params UserGroupUpdateParams, opts ...option.RequestOption) (res *UserGroupSingle, err error) {
	if !param.IsOmitted(params.XTraceID) {
		opts = append(opts, option.WithHeader("X-Trace-Id", fmt.Sprintf("%v", params.XTraceID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/vnd.api+json")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/user-groups/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return res, err
}

// List all the user groups available to this user. These are essentially the
// 'financial administrations' that Firefly III supports.
func (r *UserGroupService) List(ctx context.Context, params UserGroupListParams, opts ...option.RequestOption) (res *UserGroupListResponse, err error) {
	if !param.IsOmitted(params.XTraceID) {
		opts = append(opts, option.WithHeader("X-Trace-Id", fmt.Sprintf("%v", params.XTraceID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/vnd.api+json")}, opts...)
	path := "v1/user-groups"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

type UserGroupRead struct {
	ID         string                  `json:"id" api:"required"`
	Attributes UserGroupReadAttributes `json:"attributes" api:"required"`
	Links      ObjectLink              `json:"links" api:"required"`
	// Immutable value
	Type string `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Attributes  respjson.Field
		Links       respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UserGroupRead) RawJSON() string { return r.JSON.raw }
func (r *UserGroupRead) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserGroupReadAttributes struct {
	// Can the current user see the members of this user group?
	CanSeeMembers bool      `json:"can_see_members"`
	CreatedAt     time.Time `json:"created_at" format:"date-time"`
	// Is this user group ('financial administration') currently the active
	// administration?
	InUse   bool                            `json:"in_use"`
	Members []UserGroupReadAttributesMember `json:"members"`
	// Returns the primary currency code of the user group.
	PrimaryCurrencyCode string `json:"primary_currency_code"`
	// Returns the primary currency decimal places of the user group.
	PrimaryCurrencyDecimalPlaces int64 `json:"primary_currency_decimal_places"`
	// Returns the primary currency ID of the user group.
	PrimaryCurrencyID string `json:"primary_currency_id"`
	// Returns the primary currency symbol of the user group.
	PrimaryCurrencySymbol string `json:"primary_currency_symbol"`
	// Title of the user group. By default, it is the same as the user's email address.
	Title     string    `json:"title"`
	UpdatedAt time.Time `json:"updated_at" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CanSeeMembers                respjson.Field
		CreatedAt                    respjson.Field
		InUse                        respjson.Field
		Members                      respjson.Field
		PrimaryCurrencyCode          respjson.Field
		PrimaryCurrencyDecimalPlaces respjson.Field
		PrimaryCurrencyID            respjson.Field
		PrimaryCurrencySymbol        respjson.Field
		Title                        respjson.Field
		UpdatedAt                    respjson.Field
		ExtraFields                  map[string]respjson.Field
		raw                          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UserGroupReadAttributes) RawJSON() string { return r.JSON.raw }
func (r *UserGroupReadAttributes) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserGroupReadAttributesMember struct {
	// Any of "ro", "mng_trx", "mng_meta", "read_budgets", "read_piggies",
	// "read_subscriptions", "read_rules", "read_recurring", "read_webhooks",
	// "read_currencies", "mng_budgets", "mng_piggies", "mng_subscriptions",
	// "mng_rules", "mng_recurring", "mng_webhooks", "mng_currencies", "view_reports",
	// "view_memberships", "full", "owner".
	Roles []string `json:"roles"`
	// The email address of the member
	UserEmail string `json:"user_email" format:"email"`
	// The ID of the member.
	UserID string `json:"user_id"`
	// Is this you? (the current user)
	You bool `json:"you"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Roles       respjson.Field
		UserEmail   respjson.Field
		UserID      respjson.Field
		You         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UserGroupReadAttributesMember) RawJSON() string { return r.JSON.raw }
func (r *UserGroupReadAttributesMember) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserGroupSingle struct {
	Data UserGroupRead `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UserGroupSingle) RawJSON() string { return r.JSON.raw }
func (r *UserGroupSingle) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserGroupListResponse struct {
	Data  []UserGroupRead `json:"data" api:"required"`
	Links PageLink        `json:"links" api:"required"`
	Meta  Meta            `json:"meta" api:"required"`
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
func (r UserGroupListResponse) RawJSON() string { return r.JSON.raw }
func (r *UserGroupListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserGroupGetParams struct {
	XTraceID param.Opt[string] `header:"X-Trace-Id,omitzero" format:"uuid" json:"-"`
	paramObj
}

type UserGroupUpdateParams struct {
	// A descriptive title for the user group.
	Title string `json:"title" api:"required"`
	// Use either primary_currency_id or primary_currency_code. This will set the
	// primary currency for the user group ('financial administration').
	PrimaryCurrencyCode param.Opt[string] `json:"primary_currency_code,omitzero"`
	// Use either primary_currency_id or primary_currency_code. This will set the
	// primary currency for the user group ('financial administration').
	PrimaryCurrencyID param.Opt[string] `json:"primary_currency_id,omitzero"`
	XTraceID          param.Opt[string] `header:"X-Trace-Id,omitzero" format:"uuid" json:"-"`
	paramObj
}

func (r UserGroupUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow UserGroupUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UserGroupUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserGroupListParams struct {
	// Number of items per page. The default pagination is per 50 items.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Page number. The default pagination is per 50 items.
	Page     param.Opt[int64]  `query:"page,omitzero" json:"-"`
	XTraceID param.Opt[string] `header:"X-Trace-Id,omitzero" format:"uuid" json:"-"`
	paramObj
}

// URLQuery serializes [UserGroupListParams]'s query parameters as `url.Values`.
func (r UserGroupListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
