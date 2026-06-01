// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package firefly

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/stainless-sdks/emcees-prod-testing-5-go/internal/apiquery"
	"github.com/stainless-sdks/emcees-prod-testing-5-go/internal/requestconfig"
	"github.com/stainless-sdks/emcees-prod-testing-5-go/option"
	"github.com/stainless-sdks/emcees-prod-testing-5-go/packages/param"
)

// The &quot;charts&quot; endpoints deliver optimised data for charts and graphs.
//
// ChartBalanceService contains methods and other services that help with
// interacting with the firefly API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewChartBalanceService] method instead.
type ChartBalanceService struct {
	options []option.RequestOption
}

// NewChartBalanceService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewChartBalanceService(opts ...option.RequestOption) (r ChartBalanceService) {
	r = ChartBalanceService{}
	r.options = opts
	return
}

// This endpoint returns the data required to generate a chart with balance
// information.
func (r *ChartBalanceService) GetBalance(ctx context.Context, params ChartBalanceGetBalanceParams, opts ...option.RequestOption) (res *[]ChartDataSet, err error) {
	if !param.IsOmitted(params.XTraceID) {
		opts = append(opts, option.WithHeader("X-Trace-Id", fmt.Sprintf("%v", params.XTraceID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	path := "v1/chart/balance/balance"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

type ChartBalanceGetBalanceParams struct {
	// A date formatted YYYY-MM-DD.
	End time.Time `query:"end" api:"required" format:"date" json:"-"`
	// A date formatted YYYY-MM-DD.
	Start    time.Time         `query:"start" api:"required" format:"date" json:"-"`
	XTraceID param.Opt[string] `header:"X-Trace-Id,omitzero" format:"uuid" json:"-"`
	// Limit the chart to these asset accounts or liabilities. Only asset accounts and
	// liabilities will be accepted. Other types will be silently dropped.
	//
	// This list of accounts will be OVERRULED by the `preselected` parameter.
	Accounts []int64 `query:"accounts,omitzero" json:"-"`
	// Optional period to group the data by. If not provided, it will default to '1M'
	// or whatever is deemed relevant for the range provided.
	//
	// If you want to know which periods are available, see the enums or get the
	// configuration value: `GET /api/v1/configuration/firefly.valid_view_ranges`
	//
	// Any of "1D", "1W", "1M", "3M", "6M", "1Y".
	Period ChartBalanceGetBalanceParamsPeriod `query:"period,omitzero" json:"-"`
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
	Preselected ChartBalanceGetBalanceParamsPreselected `query:"preselected,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ChartBalanceGetBalanceParams]'s query parameters as
// `url.Values`.
func (r ChartBalanceGetBalanceParams) URLQuery() (v url.Values, err error) {
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
type ChartBalanceGetBalanceParamsPeriod string

const (
	ChartBalanceGetBalanceParamsPeriod1D ChartBalanceGetBalanceParamsPeriod = "1D"
	ChartBalanceGetBalanceParamsPeriod1W ChartBalanceGetBalanceParamsPeriod = "1W"
	ChartBalanceGetBalanceParamsPeriod1M ChartBalanceGetBalanceParamsPeriod = "1M"
	ChartBalanceGetBalanceParamsPeriod3M ChartBalanceGetBalanceParamsPeriod = "3M"
	ChartBalanceGetBalanceParamsPeriod6M ChartBalanceGetBalanceParamsPeriod = "6M"
	ChartBalanceGetBalanceParamsPeriod1Y ChartBalanceGetBalanceParamsPeriod = "1Y"
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
type ChartBalanceGetBalanceParamsPreselected string

const (
	ChartBalanceGetBalanceParamsPreselectedEmpty       ChartBalanceGetBalanceParamsPreselected = "empty"
	ChartBalanceGetBalanceParamsPreselectedAll         ChartBalanceGetBalanceParamsPreselected = "all"
	ChartBalanceGetBalanceParamsPreselectedAssets      ChartBalanceGetBalanceParamsPreselected = "assets"
	ChartBalanceGetBalanceParamsPreselectedLiabilities ChartBalanceGetBalanceParamsPreselected = "liabilities"
)
