// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package firefly

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/steel-gareth/firefly-go/internal/apijson"
	"github.com/steel-gareth/firefly-go/internal/apiquery"
	"github.com/steel-gareth/firefly-go/internal/requestconfig"
	"github.com/steel-gareth/firefly-go/option"
	"github.com/steel-gareth/firefly-go/packages/param"
	"github.com/steel-gareth/firefly-go/packages/respjson"
)

// The &quot;charts&quot; endpoints deliver optimised data for charts and graphs.
//
// ChartAccountService contains methods and other services that help with
// interacting with the firefly API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewChartAccountService] method instead.
type ChartAccountService struct {
	options []option.RequestOption
}

// NewChartAccountService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewChartAccountService(opts ...option.RequestOption) (r ChartAccountService) {
	r = ChartAccountService{}
	r.options = opts
	return
}

// This endpoint returns the data required to generate a chart with basic asset
// account balance information. This is used on the dashboard.
func (r *ChartAccountService) GetOverview(ctx context.Context, params ChartAccountGetOverviewParams, opts ...option.RequestOption) (res *[]ChartDataSet, err error) {
	if !param.IsOmitted(params.XTraceID) {
		opts = append(opts, option.WithHeader("X-Trace-Id", fmt.Sprintf("%v", params.XTraceID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	path := "v1/chart/account/overview"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

type ChartDataSet struct {
	// The currency code of the currency associated with this object.
	CurrencyCode          string `json:"currency_code"`
	CurrencyDecimalPlaces int64  `json:"currency_decimal_places"`
	// The currency ID of the currency associated with this object.
	CurrencyID string `json:"currency_id"`
	// The currency name of the currency associated with this object.
	CurrencyName   string    `json:"currency_name"`
	CurrencySymbol string    `json:"currency_symbol"`
	Date           time.Time `json:"date" format:"date-time"`
	EndDate        time.Time `json:"end_date" format:"date-time"`
	// The actual entries for this data set. They 'key' value is the label for the data
	// point. The value is the actual (numerical) value.
	Entries any `json:"entries"`
	// This is the title of the current set. It can refer to an account, a budget or
	// another object (by name).
	Label string `json:"label"`
	// The actual entries for this data set. They 'key' value is the label for the data
	// point. The value is the actual (numerical) value.
	PcEntries any `json:"pc_entries"`
	// Period of the chart.
	//
	// Any of "1D", "1W", "1M", "3M", "1Y", "custom".
	Period ChartDataSetPeriod `json:"period"`
	// The currency code of the administration's primary currency.
	PrimaryCurrencyCode string `json:"primary_currency_code"`
	// The currency decimal places of the administration's primary currency.
	PrimaryCurrencyDecimalPlaces int64 `json:"primary_currency_decimal_places"`
	// The currency ID of the administration's primary currency.
	PrimaryCurrencyID string `json:"primary_currency_id"`
	// The currency name of the administration's primary currency.
	PrimaryCurrencyName string `json:"primary_currency_name"`
	// The currency symbol of the administration's primary currency.
	PrimaryCurrencySymbol string    `json:"primary_currency_symbol"`
	StartDate             time.Time `json:"start_date" format:"date-time"`
	// Indicated the type of chart that is expected to be rendered. You can safely
	// ignore this if you want.
	Type string `json:"type"`
	// Used to indicate the Y axis for this data set. Is usually between 0 and 1 (left
	// and right side of the chart).
	YAxisID int64 `json:"yAxisID"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CurrencyCode                 respjson.Field
		CurrencyDecimalPlaces        respjson.Field
		CurrencyID                   respjson.Field
		CurrencyName                 respjson.Field
		CurrencySymbol               respjson.Field
		Date                         respjson.Field
		EndDate                      respjson.Field
		Entries                      respjson.Field
		Label                        respjson.Field
		PcEntries                    respjson.Field
		Period                       respjson.Field
		PrimaryCurrencyCode          respjson.Field
		PrimaryCurrencyDecimalPlaces respjson.Field
		PrimaryCurrencyID            respjson.Field
		PrimaryCurrencyName          respjson.Field
		PrimaryCurrencySymbol        respjson.Field
		StartDate                    respjson.Field
		Type                         respjson.Field
		YAxisID                      respjson.Field
		ExtraFields                  map[string]respjson.Field
		raw                          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChartDataSet) RawJSON() string { return r.JSON.raw }
func (r *ChartDataSet) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Period of the chart.
type ChartDataSetPeriod string

const (
	ChartDataSetPeriod1D     ChartDataSetPeriod = "1D"
	ChartDataSetPeriod1W     ChartDataSetPeriod = "1W"
	ChartDataSetPeriod1M     ChartDataSetPeriod = "1M"
	ChartDataSetPeriod3M     ChartDataSetPeriod = "3M"
	ChartDataSetPeriod1Y     ChartDataSetPeriod = "1Y"
	ChartDataSetPeriodCustom ChartDataSetPeriod = "custom"
)

type ChartAccountGetOverviewParams struct {
	// A date formatted YYYY-MM-DD.
	End time.Time `query:"end" api:"required" format:"date" json:"-"`
	// A date formatted YYYY-MM-DD.
	Start    time.Time         `query:"start" api:"required" format:"date" json:"-"`
	XTraceID param.Opt[string] `header:"X-Trace-Id,omitzero" format:"uuid" json:"-"`
	// Optional period to group the data by. If not provided, it will default to '1M'
	// or whatever is deemed relevant for the range provided.
	//
	// If you want to know which periods are available, see the enums or get the
	// configuration value: `GET /api/v1/configuration/firefly.valid_view_ranges`
	//
	// Any of "1D", "1W", "1M", "3M", "6M", "1Y".
	Period ChartAccountGetOverviewParamsPeriod `query:"period,omitzero" json:"-"`
	// Optional set of preselected accounts to limit the chart to. This may be easier
	// than submitting all asset accounts manually, for example. If you want to know
	// which selection are available, see the enums here or get the configuration
	// value: `GET /api/v1/configuration/firefly.preselected_accounts`
	//
	// - `empty`: do not do a pre-selection
	// - `all`: select all asset and all liability accounts
	// - `assets`: select all asset accounts
	// - `liabilities`: select all liability accounts
	//
	// If no accounts are found, the user's "frontpage accounts" preference will be
	// used. If that is empty, all asset accounts will be used.
	//
	// Any of "empty", "all", "assets", "liabilities".
	Preselected ChartAccountGetOverviewParamsPreselected `query:"preselected,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ChartAccountGetOverviewParams]'s query parameters as
// `url.Values`.
func (r ChartAccountGetOverviewParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Optional period to group the data by. If not provided, it will default to '1M'
// or whatever is deemed relevant for the range provided.
//
// If you want to know which periods are available, see the enums or get the
// configuration value: `GET /api/v1/configuration/firefly.valid_view_ranges`
type ChartAccountGetOverviewParamsPeriod string

const (
	ChartAccountGetOverviewParamsPeriod1D ChartAccountGetOverviewParamsPeriod = "1D"
	ChartAccountGetOverviewParamsPeriod1W ChartAccountGetOverviewParamsPeriod = "1W"
	ChartAccountGetOverviewParamsPeriod1M ChartAccountGetOverviewParamsPeriod = "1M"
	ChartAccountGetOverviewParamsPeriod3M ChartAccountGetOverviewParamsPeriod = "3M"
	ChartAccountGetOverviewParamsPeriod6M ChartAccountGetOverviewParamsPeriod = "6M"
	ChartAccountGetOverviewParamsPeriod1Y ChartAccountGetOverviewParamsPeriod = "1Y"
)

// Optional set of preselected accounts to limit the chart to. This may be easier
// than submitting all asset accounts manually, for example. If you want to know
// which selection are available, see the enums here or get the configuration
// value: `GET /api/v1/configuration/firefly.preselected_accounts`
//
// - `empty`: do not do a pre-selection
// - `all`: select all asset and all liability accounts
// - `assets`: select all asset accounts
// - `liabilities`: select all liability accounts
//
// If no accounts are found, the user's "frontpage accounts" preference will be
// used. If that is empty, all asset accounts will be used.
type ChartAccountGetOverviewParamsPreselected string

const (
	ChartAccountGetOverviewParamsPreselectedEmpty       ChartAccountGetOverviewParamsPreselected = "empty"
	ChartAccountGetOverviewParamsPreselectedAll         ChartAccountGetOverviewParamsPreselected = "all"
	ChartAccountGetOverviewParamsPreselectedAssets      ChartAccountGetOverviewParamsPreselected = "assets"
	ChartAccountGetOverviewParamsPreselectedLiabilities ChartAccountGetOverviewParamsPreselected = "liabilities"
)
