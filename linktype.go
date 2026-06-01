// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package emceesprodtesting5

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

// Endpoints to manage links between transactions, and manage the type of links
// available.
//
// LinkTypeService contains methods and other services that help with interacting
// with the emcees-prod-testing-5 API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewLinkTypeService] method instead.
type LinkTypeService struct {
	options []option.RequestOption
}

// NewLinkTypeService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewLinkTypeService(opts ...option.RequestOption) (r LinkTypeService) {
	r = LinkTypeService{}
	r.options = opts
	return
}

// Creates a new link type. The data required can be submitted as a JSON body or as
// a list of parameters (in key=value pairs, like a webform).
func (r *LinkTypeService) New(ctx context.Context, params LinkTypeNewParams, opts ...option.RequestOption) (res *LinkTypeSingle, err error) {
	if !param.IsOmitted(params.XTraceID) {
		opts = append(opts, option.WithHeader("X-Trace-Id", fmt.Sprintf("%v", params.XTraceID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/vnd.api+json")}, opts...)
	path := "v1/link-types"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Returns a single link type by its ID.
func (r *LinkTypeService) Get(ctx context.Context, id string, query LinkTypeGetParams, opts ...option.RequestOption) (res *LinkTypeSingle, err error) {
	if !param.IsOmitted(query.XTraceID) {
		opts = append(opts, option.WithHeader("X-Trace-Id", fmt.Sprintf("%v", query.XTraceID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/vnd.api+json")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/link-types/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Used to update a single link type. All fields that are not submitted will be
// cleared (set to NULL). The model will tell you which fields are mandatory. You
// cannot update some of the system provided link types, indicated by the
// editable=false flag when you list it.
func (r *LinkTypeService) Update(ctx context.Context, id string, params LinkTypeUpdateParams, opts ...option.RequestOption) (res *LinkTypeSingle, err error) {
	if !param.IsOmitted(params.XTraceID) {
		opts = append(opts, option.WithHeader("X-Trace-Id", fmt.Sprintf("%v", params.XTraceID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/vnd.api+json")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/link-types/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return res, err
}

// List all the link types the system has. These include the default ones as well
// as any new ones.
func (r *LinkTypeService) List(ctx context.Context, params LinkTypeListParams, opts ...option.RequestOption) (res *LinkTypeListResponse, err error) {
	if !param.IsOmitted(params.XTraceID) {
		opts = append(opts, option.WithHeader("X-Trace-Id", fmt.Sprintf("%v", params.XTraceID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/vnd.api+json")}, opts...)
	path := "v1/link-types"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

// Will permanently delete a link type. The links between transactions will be
// removed. The transactions themselves remain. You cannot delete some of the
// system provided link types, indicated by the editable=false flag when you list
// it.
func (r *LinkTypeService) Delete(ctx context.Context, id string, body LinkTypeDeleteParams, opts ...option.RequestOption) (err error) {
	if !param.IsOmitted(body.XTraceID) {
		opts = append(opts, option.WithHeader("X-Trace-Id", fmt.Sprintf("%v", body.XTraceID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return err
	}
	path := fmt.Sprintf("v1/link-types/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// List all transactions under this link type, both the inward and outward
// transactions.
func (r *LinkTypeService) ListTransactions(ctx context.Context, id string, params LinkTypeListTransactionsParams, opts ...option.RequestOption) (res *TransactionArray, err error) {
	if !param.IsOmitted(params.XTraceID) {
		opts = append(opts, option.WithHeader("X-Trace-Id", fmt.Sprintf("%v", params.XTraceID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/link-types/%s/transactions", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

type LinkType struct {
	Inward   string `json:"inward" api:"required"`
	Name     string `json:"name" api:"required"`
	Outward  string `json:"outward" api:"required"`
	Editable bool   `json:"editable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Inward      respjson.Field
		Name        respjson.Field
		Outward     respjson.Field
		Editable    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LinkType) RawJSON() string { return r.JSON.raw }
func (r *LinkType) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this LinkType to a LinkTypeParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// LinkTypeParam.Overrides()
func (r LinkType) ToParam() LinkTypeParam {
	return param.Override[LinkTypeParam](json.RawMessage(r.RawJSON()))
}

// The properties Inward, Name, Outward are required.
type LinkTypeParam struct {
	Inward  string `json:"inward" api:"required"`
	Name    string `json:"name" api:"required"`
	Outward string `json:"outward" api:"required"`
	paramObj
}

func (r LinkTypeParam) MarshalJSON() (data []byte, err error) {
	type shadow LinkTypeParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LinkTypeParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LinkTypeRead struct {
	ID         string     `json:"id" api:"required"`
	Attributes LinkType   `json:"attributes" api:"required"`
	Links      ObjectLink `json:"links" api:"required"`
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
func (r LinkTypeRead) RawJSON() string { return r.JSON.raw }
func (r *LinkTypeRead) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LinkTypeSingle struct {
	Data LinkTypeRead `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LinkTypeSingle) RawJSON() string { return r.JSON.raw }
func (r *LinkTypeSingle) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LinkTypeListResponse struct {
	Data  []LinkTypeRead `json:"data" api:"required"`
	Links PageLink       `json:"links" api:"required"`
	Meta  Meta           `json:"meta" api:"required"`
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
func (r LinkTypeListResponse) RawJSON() string { return r.JSON.raw }
func (r *LinkTypeListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LinkTypeNewParams struct {
	LinkType LinkTypeParam
	XTraceID param.Opt[string] `header:"X-Trace-Id,omitzero" format:"uuid" json:"-"`
	paramObj
}

func (r LinkTypeNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.LinkType)
}
func (r *LinkTypeNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LinkTypeGetParams struct {
	XTraceID param.Opt[string] `header:"X-Trace-Id,omitzero" format:"uuid" json:"-"`
	paramObj
}

type LinkTypeUpdateParams struct {
	Inward   param.Opt[string] `json:"inward,omitzero"`
	Name     param.Opt[string] `json:"name,omitzero"`
	Outward  param.Opt[string] `json:"outward,omitzero"`
	XTraceID param.Opt[string] `header:"X-Trace-Id,omitzero" format:"uuid" json:"-"`
	paramObj
}

func (r LinkTypeUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow LinkTypeUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LinkTypeUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LinkTypeListParams struct {
	// Number of items per page. The default pagination is per 50 items.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Page number. The default pagination is per 50 items.
	Page     param.Opt[int64]  `query:"page,omitzero" json:"-"`
	XTraceID param.Opt[string] `header:"X-Trace-Id,omitzero" format:"uuid" json:"-"`
	paramObj
}

// URLQuery serializes [LinkTypeListParams]'s query parameters as `url.Values`.
func (r LinkTypeListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type LinkTypeDeleteParams struct {
	XTraceID param.Opt[string] `header:"X-Trace-Id,omitzero" format:"uuid" json:"-"`
	paramObj
}

type LinkTypeListTransactionsParams struct {
	// A date formatted YYYY-MM-DD, to limit the results.
	End param.Opt[time.Time] `query:"end,omitzero" format:"date" json:"-"`
	// Number of items per page. The default pagination is per 50 items.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Page number. The default pagination is per 50 items.
	Page param.Opt[int64] `query:"page,omitzero" json:"-"`
	// A date formatted YYYY-MM-DD, to limit the results.
	Start    param.Opt[time.Time] `query:"start,omitzero" format:"date" json:"-"`
	XTraceID param.Opt[string]    `header:"X-Trace-Id,omitzero" format:"uuid" json:"-"`
	// Any of "all", "withdrawal", "withdrawals", "expense", "deposit", "deposits",
	// "income", "transfer", "transfers", "opening_balance", "reconciliation",
	// "special", "specials", "default".
	Type TransactionTypeFilter `query:"type,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [LinkTypeListTransactionsParams]'s query parameters as
// `url.Values`.
func (r LinkTypeListTransactionsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
