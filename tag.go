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

// This endpoint manages all of the user&#039;s tags.
//
// TagService contains methods and other services that help with interacting with
// the emcees-prod-testing-5 API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTagService] method instead.
type TagService struct {
	options []option.RequestOption
}

// NewTagService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewTagService(opts ...option.RequestOption) (r TagService) {
	r = TagService{}
	r.options = opts
	return
}

// Creates a new tag. The data required can be submitted as a JSON body or as a
// list of parameters.
func (r *TagService) New(ctx context.Context, params TagNewParams, opts ...option.RequestOption) (res *TagSingle, err error) {
	if !param.IsOmitted(params.XTraceID) {
		opts = append(opts, option.WithHeader("X-Trace-Id", fmt.Sprintf("%v", params.XTraceID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/vnd.api+json")}, opts...)
	path := "v1/tags"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Get a single tag.
func (r *TagService) Get(ctx context.Context, tag string, params TagGetParams, opts ...option.RequestOption) (res *TagSingle, err error) {
	if !param.IsOmitted(params.XTraceID) {
		opts = append(opts, option.WithHeader("X-Trace-Id", fmt.Sprintf("%v", params.XTraceID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/vnd.api+json")}, opts...)
	if tag == "" {
		err = errors.New("missing required tag parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/tags/%s", url.PathEscape(tag))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

// Update existing tag.
func (r *TagService) Update(ctx context.Context, tag string, params TagUpdateParams, opts ...option.RequestOption) (res *TagSingle, err error) {
	if !param.IsOmitted(params.XTraceID) {
		opts = append(opts, option.WithHeader("X-Trace-Id", fmt.Sprintf("%v", params.XTraceID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/vnd.api+json")}, opts...)
	if tag == "" {
		err = errors.New("missing required tag parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/tags/%s", url.PathEscape(tag))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return res, err
}

// List all of the user's tags.
func (r *TagService) List(ctx context.Context, params TagListParams, opts ...option.RequestOption) (res *TagListResponse, err error) {
	if !param.IsOmitted(params.XTraceID) {
		opts = append(opts, option.WithHeader("X-Trace-Id", fmt.Sprintf("%v", params.XTraceID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/vnd.api+json")}, opts...)
	path := "v1/tags"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

// Delete an tag.
func (r *TagService) Delete(ctx context.Context, tag string, body TagDeleteParams, opts ...option.RequestOption) (err error) {
	if !param.IsOmitted(body.XTraceID) {
		opts = append(opts, option.WithHeader("X-Trace-Id", fmt.Sprintf("%v", body.XTraceID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if tag == "" {
		err = errors.New("missing required tag parameter")
		return err
	}
	path := fmt.Sprintf("v1/tags/%s", url.PathEscape(tag))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// Lists all attachments.
func (r *TagService) ListAttachments(ctx context.Context, tag string, params TagListAttachmentsParams, opts ...option.RequestOption) (res *AttachmentArray, err error) {
	if !param.IsOmitted(params.XTraceID) {
		opts = append(opts, option.WithHeader("X-Trace-Id", fmt.Sprintf("%v", params.XTraceID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/vnd.api+json")}, opts...)
	if tag == "" {
		err = errors.New("missing required tag parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/tags/%s/attachments", url.PathEscape(tag))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

// List all transactions with this tag.
func (r *TagService) ListTransactions(ctx context.Context, tag string, params TagListTransactionsParams, opts ...option.RequestOption) (res *TransactionArray, err error) {
	if !param.IsOmitted(params.XTraceID) {
		opts = append(opts, option.WithHeader("X-Trace-Id", fmt.Sprintf("%v", params.XTraceID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/vnd.api+json")}, opts...)
	if tag == "" {
		err = errors.New("missing required tag parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/tags/%s/transactions", url.PathEscape(tag))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

type TagRead struct {
	ID         string            `json:"id" api:"required"`
	Attributes TagReadAttributes `json:"attributes" api:"required"`
	Links      ObjectLink        `json:"links" api:"required"`
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
func (r TagRead) RawJSON() string { return r.JSON.raw }
func (r *TagRead) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TagReadAttributes struct {
	// The tag
	Tag       string    `json:"tag" api:"required"`
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// The date to which the tag is applicable.
	Date        time.Time `json:"date" api:"nullable" format:"date"`
	Description string    `json:"description" api:"nullable"`
	// Latitude of the tag's location, if applicable. Can be used to draw a map.
	Latitude float64 `json:"latitude" api:"nullable"`
	// Latitude of the tag's location, if applicable. Can be used to draw a map.
	Longitude float64   `json:"longitude" api:"nullable"`
	UpdatedAt time.Time `json:"updated_at" format:"date-time"`
	// Zoom level for the map, if drawn. This to set the box right. Unfortunately this
	// is a proprietary value because each map provider has different zoom levels.
	ZoomLevel int64 `json:"zoom_level" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Tag         respjson.Field
		CreatedAt   respjson.Field
		Date        respjson.Field
		Description respjson.Field
		Latitude    respjson.Field
		Longitude   respjson.Field
		UpdatedAt   respjson.Field
		ZoomLevel   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TagReadAttributes) RawJSON() string { return r.JSON.raw }
func (r *TagReadAttributes) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TagSingle struct {
	Data TagRead `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TagSingle) RawJSON() string { return r.JSON.raw }
func (r *TagSingle) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TagListResponse struct {
	Data  []TagRead `json:"data" api:"required"`
	Links PageLink  `json:"links" api:"required"`
	Meta  Meta      `json:"meta" api:"required"`
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
func (r TagListResponse) RawJSON() string { return r.JSON.raw }
func (r *TagListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TagNewParams struct {
	// The tag
	Tag string `json:"tag" api:"required"`
	// The date to which the tag is applicable.
	Date        param.Opt[time.Time] `json:"date,omitzero" format:"date"`
	Description param.Opt[string]    `json:"description,omitzero"`
	// Latitude of the tag's location, if applicable. Can be used to draw a map.
	Latitude param.Opt[float64] `json:"latitude,omitzero"`
	// Latitude of the tag's location, if applicable. Can be used to draw a map.
	Longitude param.Opt[float64] `json:"longitude,omitzero"`
	// Zoom level for the map, if drawn. This to set the box right. Unfortunately this
	// is a proprietary value because each map provider has different zoom levels.
	ZoomLevel param.Opt[int64]  `json:"zoom_level,omitzero"`
	XTraceID  param.Opt[string] `header:"X-Trace-Id,omitzero" format:"uuid" json:"-"`
	paramObj
}

func (r TagNewParams) MarshalJSON() (data []byte, err error) {
	type shadow TagNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TagNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TagGetParams struct {
	// Number of items per page. The default pagination is per 50 items.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Page number. The default pagination is per 50 items.
	Page     param.Opt[int64]  `query:"page,omitzero" json:"-"`
	XTraceID param.Opt[string] `header:"X-Trace-Id,omitzero" format:"uuid" json:"-"`
	paramObj
}

// URLQuery serializes [TagGetParams]'s query parameters as `url.Values`.
func (r TagGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type TagUpdateParams struct {
	// The date to which the tag is applicable.
	Date        param.Opt[time.Time] `json:"date,omitzero" format:"date"`
	Description param.Opt[string]    `json:"description,omitzero"`
	// Latitude of the tag's location, if applicable. Can be used to draw a map.
	Latitude param.Opt[float64] `json:"latitude,omitzero"`
	// Latitude of the tag's location, if applicable. Can be used to draw a map.
	Longitude param.Opt[float64] `json:"longitude,omitzero"`
	// Zoom level for the map, if drawn. This to set the box right. Unfortunately this
	// is a proprietary value because each map provider has different zoom levels.
	ZoomLevel param.Opt[int64] `json:"zoom_level,omitzero"`
	// The tag
	Tag      param.Opt[string] `json:"tag,omitzero"`
	XTraceID param.Opt[string] `header:"X-Trace-Id,omitzero" format:"uuid" json:"-"`
	paramObj
}

func (r TagUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow TagUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TagUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TagListParams struct {
	// Number of items per page. The default pagination is per 50 items.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Page number. The default pagination is per 50 items.
	Page     param.Opt[int64]  `query:"page,omitzero" json:"-"`
	XTraceID param.Opt[string] `header:"X-Trace-Id,omitzero" format:"uuid" json:"-"`
	paramObj
}

// URLQuery serializes [TagListParams]'s query parameters as `url.Values`.
func (r TagListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type TagDeleteParams struct {
	XTraceID param.Opt[string] `header:"X-Trace-Id,omitzero" format:"uuid" json:"-"`
	paramObj
}

type TagListAttachmentsParams struct {
	// Number of items per page. The default pagination is per 50 items.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Page number. The default pagination is per 50 items.
	Page     param.Opt[int64]  `query:"page,omitzero" json:"-"`
	XTraceID param.Opt[string] `header:"X-Trace-Id,omitzero" format:"uuid" json:"-"`
	paramObj
}

// URLQuery serializes [TagListAttachmentsParams]'s query parameters as
// `url.Values`.
func (r TagListAttachmentsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type TagListTransactionsParams struct {
	// A date formatted YYYY-MM-DD. This is the end date of the selected range
	// (inclusive).
	End param.Opt[time.Time] `query:"end,omitzero" format:"date" json:"-"`
	// Number of items per page. The default pagination is per 50 items.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Page number. The default pagination is per 50 items.
	Page param.Opt[int64] `query:"page,omitzero" json:"-"`
	// A date formatted YYYY-MM-DD. This is the start date of the selected range
	// (inclusive).
	Start    param.Opt[time.Time] `query:"start,omitzero" format:"date" json:"-"`
	XTraceID param.Opt[string]    `header:"X-Trace-Id,omitzero" format:"uuid" json:"-"`
	// Any of "all", "withdrawal", "withdrawals", "expense", "deposit", "deposits",
	// "income", "transfer", "transfers", "opening_balance", "reconciliation",
	// "special", "specials", "default".
	Type TransactionTypeFilter `query:"type,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [TagListTransactionsParams]'s query parameters as
// `url.Values`.
func (r TagListTransactionsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
