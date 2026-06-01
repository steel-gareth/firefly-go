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

// Endpoints to manage links between transactions, and manage the type of links
// available.
//
// TransactionLinkService contains methods and other services that help with
// interacting with the emcees-prod-testing-5 API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTransactionLinkService] method instead.
type TransactionLinkService struct {
	options []option.RequestOption
}

// NewTransactionLinkService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewTransactionLinkService(opts ...option.RequestOption) (r TransactionLinkService) {
	r = TransactionLinkService{}
	r.options = opts
	return
}

// Store a new link between two transactions. For this end point you need the
// journal_id from a transaction.
func (r *TransactionLinkService) New(ctx context.Context, params TransactionLinkNewParams, opts ...option.RequestOption) (res *TransactionLinkSingle, err error) {
	if !param.IsOmitted(params.XTraceID) {
		opts = append(opts, option.WithHeader("X-Trace-Id", fmt.Sprintf("%v", params.XTraceID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/vnd.api+json")}, opts...)
	path := "v1/transaction-links"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Returns a single link by its ID.
func (r *TransactionLinkService) Get(ctx context.Context, id string, query TransactionLinkGetParams, opts ...option.RequestOption) (res *TransactionLinkSingle, err error) {
	if !param.IsOmitted(query.XTraceID) {
		opts = append(opts, option.WithHeader("X-Trace-Id", fmt.Sprintf("%v", query.XTraceID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/vnd.api+json")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/transaction-links/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Used to update a single existing link.
func (r *TransactionLinkService) Update(ctx context.Context, id string, params TransactionLinkUpdateParams, opts ...option.RequestOption) (res *TransactionLinkSingle, err error) {
	if !param.IsOmitted(params.XTraceID) {
		opts = append(opts, option.WithHeader("X-Trace-Id", fmt.Sprintf("%v", params.XTraceID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/vnd.api+json")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/transaction-links/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return res, err
}

// List all the transaction links.
func (r *TransactionLinkService) List(ctx context.Context, params TransactionLinkListParams, opts ...option.RequestOption) (res *TransactionLinkArray, err error) {
	if !param.IsOmitted(params.XTraceID) {
		opts = append(opts, option.WithHeader("X-Trace-Id", fmt.Sprintf("%v", params.XTraceID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/vnd.api+json")}, opts...)
	path := "v1/transaction-links"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

// Will permanently delete link. Transactions remain.
func (r *TransactionLinkService) Delete(ctx context.Context, id string, body TransactionLinkDeleteParams, opts ...option.RequestOption) (err error) {
	if !param.IsOmitted(body.XTraceID) {
		opts = append(opts, option.WithHeader("X-Trace-Id", fmt.Sprintf("%v", body.XTraceID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return err
	}
	path := fmt.Sprintf("v1/transaction-links/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

type TransactionLinkArray struct {
	Data  []TransactionLinkRead `json:"data" api:"required"`
	Links PageLink              `json:"links" api:"required"`
	Meta  Meta                  `json:"meta" api:"required"`
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
func (r TransactionLinkArray) RawJSON() string { return r.JSON.raw }
func (r *TransactionLinkArray) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TransactionLinkRead struct {
	ID         string                        `json:"id" api:"required"`
	Attributes TransactionLinkReadAttributes `json:"attributes" api:"required"`
	Links      ObjectLink                    `json:"links" api:"required"`
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
func (r TransactionLinkRead) RawJSON() string { return r.JSON.raw }
func (r *TransactionLinkRead) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TransactionLinkReadAttributes struct {
	// The inward transaction transaction_journal_id for the link. This becomes the 'is
	// paid by' transaction of the set.
	InwardID string `json:"inward_id" api:"required"`
	// The link type ID to use. You can also use the link_type_name field.
	LinkTypeID string `json:"link_type_id" api:"required"`
	// The outward transaction transaction_journal_id for the link. This becomes the
	// 'pays for' transaction of the set.
	OutwardID string    `json:"outward_id" api:"required"`
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// The link type name to use. You can also use the link_type_id field.
	LinkTypeName string `json:"link_type_name"`
	// Optional. Some notes.
	Notes     string    `json:"notes" api:"nullable"`
	UpdatedAt time.Time `json:"updated_at" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		InwardID     respjson.Field
		LinkTypeID   respjson.Field
		OutwardID    respjson.Field
		CreatedAt    respjson.Field
		LinkTypeName respjson.Field
		Notes        respjson.Field
		UpdatedAt    respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TransactionLinkReadAttributes) RawJSON() string { return r.JSON.raw }
func (r *TransactionLinkReadAttributes) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TransactionLinkSingle struct {
	Data TransactionLinkRead `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TransactionLinkSingle) RawJSON() string { return r.JSON.raw }
func (r *TransactionLinkSingle) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TransactionLinkNewParams struct {
	// The inward transaction transaction_journal_id for the link. This becomes the 'is
	// paid by' transaction of the set.
	InwardID string `json:"inward_id" api:"required"`
	// The link type ID to use. You can also use the link_type_name field.
	LinkTypeID string `json:"link_type_id" api:"required"`
	// The outward transaction transaction_journal_id for the link. This becomes the
	// 'pays for' transaction of the set.
	OutwardID string `json:"outward_id" api:"required"`
	// Optional. Some notes.
	Notes param.Opt[string] `json:"notes,omitzero"`
	// The link type name to use. You can also use the link_type_id field.
	LinkTypeName param.Opt[string] `json:"link_type_name,omitzero"`
	XTraceID     param.Opt[string] `header:"X-Trace-Id,omitzero" format:"uuid" json:"-"`
	paramObj
}

func (r TransactionLinkNewParams) MarshalJSON() (data []byte, err error) {
	type shadow TransactionLinkNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TransactionLinkNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TransactionLinkGetParams struct {
	XTraceID param.Opt[string] `header:"X-Trace-Id,omitzero" format:"uuid" json:"-"`
	paramObj
}

type TransactionLinkUpdateParams struct {
	// Optional. Some notes. If you submit an empty string the current notes will be
	// removed
	Notes param.Opt[string] `json:"notes,omitzero"`
	// The inward transaction transaction_journal_id for the link. This becomes the 'is
	// paid by' transaction of the set.
	InwardID param.Opt[string] `json:"inward_id,omitzero"`
	// The link type ID to use. Use this field OR use the link_type_name field.
	LinkTypeID param.Opt[string] `json:"link_type_id,omitzero"`
	// The link type name to use. Use this field OR use the link_type_id field.
	LinkTypeName param.Opt[string] `json:"link_type_name,omitzero"`
	// The outward transaction transaction_journal_id for the link. This becomes the
	// 'pays for' transaction of the set.
	OutwardID param.Opt[string] `json:"outward_id,omitzero"`
	XTraceID  param.Opt[string] `header:"X-Trace-Id,omitzero" format:"uuid" json:"-"`
	paramObj
}

func (r TransactionLinkUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow TransactionLinkUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TransactionLinkUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TransactionLinkListParams struct {
	// Number of items per page. The default pagination is per 50 items.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Page number. The default pagination is per 50 items.
	Page     param.Opt[int64]  `query:"page,omitzero" json:"-"`
	XTraceID param.Opt[string] `header:"X-Trace-Id,omitzero" format:"uuid" json:"-"`
	paramObj
}

// URLQuery serializes [TransactionLinkListParams]'s query parameters as
// `url.Values`.
func (r TransactionLinkListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type TransactionLinkDeleteParams struct {
	XTraceID param.Opt[string] `header:"X-Trace-Id,omitzero" format:"uuid" json:"-"`
	paramObj
}
