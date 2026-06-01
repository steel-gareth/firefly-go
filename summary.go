// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package firefly

import (
	"context"
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

// These endpoints deliver summaries, like sums, lists of numbers and other
// processed information. Mainly used for the main dashboard and pretty specific
// for Firefly III itself.
//
// SummaryService contains methods and other services that help with interacting
// with the firefly API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSummaryService] method instead.
type SummaryService struct {
	options []option.RequestOption
}

// NewSummaryService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewSummaryService(opts ...option.RequestOption) (r SummaryService) {
	r = SummaryService{}
	r.options = opts
	return
}

// Returns basic sums of the users data, like the net worth, spent and earned
// amounts. It is multi-currency, and is used in Firefly III to populate the
// dashboard.
func (r *SummaryService) GetBasic(ctx context.Context, params SummaryGetBasicParams, opts ...option.RequestOption) (res *SummaryGetBasicResponse, err error) {
	if !param.IsOmitted(params.XTraceID) {
		opts = append(opts, option.WithHeader("X-Trace-Id", fmt.Sprintf("%v", params.XTraceID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/vnd.api+json")}, opts...)
	path := "v1/summary/basic"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

type SummaryGetBasicResponse map[string]SummaryGetBasicResponseItem

type SummaryGetBasicResponseItem struct {
	CurrencyCode string `json:"currency_code"`
	// Number of decimals for the associated currency.
	CurrencyDecimalPlaces int64 `json:"currency_decimal_places"`
	// The currency ID of the associated currency.
	CurrencyID     string `json:"currency_id"`
	CurrencySymbol string `json:"currency_symbol"`
	// This is a reference to the type of info shared, not influenced by translations
	// or user preferences. The EUR value is a reference to the currency code.
	// Possibilities are: balance-in-ABC, spent-in-ABC, earned-in-ABC,
	// bills-paid-in-ABC, bills-unpaid-in-ABC, left-to-spend-in-ABC and
	// net-worth-in-ABC.
	Key string `json:"key"`
	// Reference to a font-awesome icon without the fa- part.
	LocalIcon string `json:"local_icon"`
	// The amount as a float.
	MonetaryValue float64 `json:"monetary_value"`
	// True if there are no available budgets available.
	NoAvailableBudgets bool `json:"no_available_budgets"`
	// A short explanation of the amounts origin. Already formatted according to the
	// locale of the user or translated, if relevant.
	SubTitle string `json:"sub_title"`
	// A translated title for the information shared.
	Title string `json:"title"`
	// The amount formatted according to the users locale
	ValueParsed string `json:"value_parsed"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CurrencyCode          respjson.Field
		CurrencyDecimalPlaces respjson.Field
		CurrencyID            respjson.Field
		CurrencySymbol        respjson.Field
		Key                   respjson.Field
		LocalIcon             respjson.Field
		MonetaryValue         respjson.Field
		NoAvailableBudgets    respjson.Field
		SubTitle              respjson.Field
		Title                 respjson.Field
		ValueParsed           respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SummaryGetBasicResponseItem) RawJSON() string { return r.JSON.raw }
func (r *SummaryGetBasicResponseItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SummaryGetBasicParams struct {
	// A date formatted YYYY-MM-DD.
	End time.Time `query:"end" api:"required" format:"date" json:"-"`
	// A date formatted YYYY-MM-DD.
	Start time.Time `query:"start" api:"required" format:"date" json:"-"`
	// A currency code like EUR or USD, to filter the result.
	CurrencyCode param.Opt[string] `query:"currency_code,omitzero" json:"-"`
	XTraceID     param.Opt[string] `header:"X-Trace-Id,omitzero" format:"uuid" json:"-"`
	paramObj
}

// URLQuery serializes [SummaryGetBasicParams]'s query parameters as `url.Values`.
func (r SummaryGetBasicParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
