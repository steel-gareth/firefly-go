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

	"github.com/stainless-sdks/emcees-prod-testing-5-go/internal/apijson"
	"github.com/stainless-sdks/emcees-prod-testing-5-go/internal/apiquery"
	shimjson "github.com/stainless-sdks/emcees-prod-testing-5-go/internal/encoding/json"
	"github.com/stainless-sdks/emcees-prod-testing-5-go/internal/requestconfig"
	"github.com/stainless-sdks/emcees-prod-testing-5-go/option"
	"github.com/stainless-sdks/emcees-prod-testing-5-go/packages/param"
	"github.com/stainless-sdks/emcees-prod-testing-5-go/packages/respjson"
)

// These endpoints can be used to manage the user&#039;s preferences, including
// some hidden ones.
//
// PreferenceService contains methods and other services that help with interacting
// with the firefly API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPreferenceService] method instead.
type PreferenceService struct {
	options []option.RequestOption
}

// NewPreferenceService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewPreferenceService(opts ...option.RequestOption) (r PreferenceService) {
	r = PreferenceService{}
	r.options = opts
	return
}

// This endpoint creates a new preference. The name and data are free-format, and
// entirely up to you. If the preference is not used in Firefly III itself it may
// not be configurable through the user interface, but you can use this endpoint to
// persist custom data for your own app.
func (r *PreferenceService) New(ctx context.Context, params PreferenceNewParams, opts ...option.RequestOption) (res *PreferenceSingle, err error) {
	if !param.IsOmitted(params.XTraceID) {
		opts = append(opts, option.WithHeader("X-Trace-Id", fmt.Sprintf("%v", params.XTraceID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/vnd.api+json")}, opts...)
	path := "v1/preferences"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Return a single preference and the value.
func (r *PreferenceService) Get(ctx context.Context, name string, query PreferenceGetParams, opts ...option.RequestOption) (res *PreferenceSingle, err error) {
	if !param.IsOmitted(query.XTraceID) {
		opts = append(opts, option.WithHeader("X-Trace-Id", fmt.Sprintf("%v", query.XTraceID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/vnd.api+json")}, opts...)
	if name == "" {
		err = errors.New("missing required name parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/preferences/%s", url.PathEscape(name))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Update a user's preference.
func (r *PreferenceService) Update(ctx context.Context, name string, params PreferenceUpdateParams, opts ...option.RequestOption) (res *PreferenceSingle, err error) {
	if !param.IsOmitted(params.XTraceID) {
		opts = append(opts, option.WithHeader("X-Trace-Id", fmt.Sprintf("%v", params.XTraceID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/vnd.api+json")}, opts...)
	if name == "" {
		err = errors.New("missing required name parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/preferences/%s", url.PathEscape(name))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return res, err
}

// List all of the preferences of the user.
func (r *PreferenceService) List(ctx context.Context, params PreferenceListParams, opts ...option.RequestOption) (res *PreferenceListResponse, err error) {
	if !param.IsOmitted(params.XTraceID) {
		opts = append(opts, option.WithHeader("X-Trace-Id", fmt.Sprintf("%v", params.XTraceID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/vnd.api+json")}, opts...)
	path := "v1/preferences"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

type Preference struct {
	Data      PolymorphicPropertyUnion `json:"data" api:"required"`
	Name      string                   `json:"name" api:"required"`
	CreatedAt time.Time                `json:"created_at" format:"date-time"`
	UpdatedAt time.Time                `json:"updated_at" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Name        respjson.Field
		CreatedAt   respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Preference) RawJSON() string { return r.JSON.raw }
func (r *Preference) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this Preference to a PreferenceParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// PreferenceParam.Overrides()
func (r Preference) ToParam() PreferenceParam {
	return param.Override[PreferenceParam](json.RawMessage(r.RawJSON()))
}

// The properties Data, Name are required.
type PreferenceParam struct {
	Data PolymorphicPropertyUnionParam `json:"data,omitzero" api:"required"`
	Name string                        `json:"name" api:"required"`
	paramObj
}

func (r PreferenceParam) MarshalJSON() (data []byte, err error) {
	type shadow PreferenceParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PreferenceParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PreferenceRead struct {
	ID         string     `json:"id" api:"required"`
	Attributes Preference `json:"attributes" api:"required"`
	// Immutable value
	Type string `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Attributes  respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PreferenceRead) RawJSON() string { return r.JSON.raw }
func (r *PreferenceRead) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PreferenceSingle struct {
	Data PreferenceRead `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PreferenceSingle) RawJSON() string { return r.JSON.raw }
func (r *PreferenceSingle) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PreferenceListResponse struct {
	Data  []PreferenceRead `json:"data" api:"required"`
	Links PageLink         `json:"links" api:"required"`
	Meta  Meta             `json:"meta" api:"required"`
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
func (r PreferenceListResponse) RawJSON() string { return r.JSON.raw }
func (r *PreferenceListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PreferenceNewParams struct {
	Preference PreferenceParam
	XTraceID   param.Opt[string] `header:"X-Trace-Id,omitzero" format:"uuid" json:"-"`
	paramObj
}

func (r PreferenceNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.Preference)
}
func (r *PreferenceNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PreferenceGetParams struct {
	XTraceID param.Opt[string] `header:"X-Trace-Id,omitzero" format:"uuid" json:"-"`
	paramObj
}

type PreferenceUpdateParams struct {
	Data     PolymorphicPropertyUnionParam `json:"data,omitzero" api:"required"`
	XTraceID param.Opt[string]             `header:"X-Trace-Id,omitzero" format:"uuid" json:"-"`
	paramObj
}

func (r PreferenceUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow PreferenceUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PreferenceUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PreferenceListParams struct {
	// Number of items per page. The default pagination is per 50 items.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Page number. The default pagination is per 50 items.
	Page     param.Opt[int64]  `query:"page,omitzero" json:"-"`
	XTraceID param.Opt[string] `header:"X-Trace-Id,omitzero" format:"uuid" json:"-"`
	paramObj
}

// URLQuery serializes [PreferenceListParams]'s query parameters as `url.Values`.
func (r PreferenceListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
