// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package emceesprodtesting5

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/stainless-sdks/emcees-prod-testing-5-go/internal/apiquery"
	"github.com/stainless-sdks/emcees-prod-testing-5-go/internal/requestconfig"
	"github.com/stainless-sdks/emcees-prod-testing-5-go/option"
	"github.com/stainless-sdks/emcees-prod-testing-5-go/packages/param"
)

// The &quot;data&quot;-endpoints manage generic Firefly III and user-specific
// data.
//
// DataService contains methods and other services that help with interacting with
// the emcees-prod-testing-5 API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDataService] method instead.
type DataService struct {
	options []option.RequestOption
	// The &quot;data&quot;-endpoints manage generic Firefly III and user-specific
	// data.
	Bulk DataBulkService
	// The &quot;data&quot;-endpoints manage generic Firefly III and user-specific
	// data.
	Export DataExportService
}

// NewDataService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewDataService(opts ...option.RequestOption) (r DataService) {
	r = DataService{}
	r.options = opts
	r.Bulk = NewDataBulkService(opts...)
	r.Export = NewDataExportService(opts...)
	return
}

// A call to this endpoint deletes the requested data type. Use it with care and
// always with user permission. The demo user is incapable of using this endpoint.
func (r *DataService) Destroy(ctx context.Context, params DataDestroyParams, opts ...option.RequestOption) (err error) {
	if !param.IsOmitted(params.XTraceID) {
		opts = append(opts, option.WithHeader("X-Trace-Id", fmt.Sprintf("%v", params.XTraceID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "v1/data/destroy"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, params, nil, opts...)
	return err
}

// A call to this endpoint purges all previously deleted data. Use it with care and
// always with user permission. The demo user is incapable of using this endpoint.
func (r *DataService) Purge(ctx context.Context, body DataPurgeParams, opts ...option.RequestOption) (err error) {
	if !param.IsOmitted(body.XTraceID) {
		opts = append(opts, option.WithHeader("X-Trace-Id", fmt.Sprintf("%v", body.XTraceID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "v1/data/purge"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

type DataDestroyParams struct {
	// The type of data that you wish to destroy. You can only use one at a time.
	//
	// Any of "not_assets_liabilities", "budgets", "bills", "piggy_banks", "rules",
	// "recurring", "categories", "tags", "object_groups", "accounts",
	// "asset_accounts", "expense_accounts", "revenue_accounts", "liabilities",
	// "transactions", "withdrawals", "deposits", "transfers".
	Objects  DataDestroyParamsObjects `query:"objects,omitzero" api:"required" json:"-"`
	XTraceID param.Opt[string]        `header:"X-Trace-Id,omitzero" format:"uuid" json:"-"`
	paramObj
}

// URLQuery serializes [DataDestroyParams]'s query parameters as `url.Values`.
func (r DataDestroyParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// The type of data that you wish to destroy. You can only use one at a time.
type DataDestroyParamsObjects string

const (
	DataDestroyParamsObjectsNotAssetsLiabilities DataDestroyParamsObjects = "not_assets_liabilities"
	DataDestroyParamsObjectsBudgets              DataDestroyParamsObjects = "budgets"
	DataDestroyParamsObjectsBills                DataDestroyParamsObjects = "bills"
	DataDestroyParamsObjectsPiggyBanks           DataDestroyParamsObjects = "piggy_banks"
	DataDestroyParamsObjectsRules                DataDestroyParamsObjects = "rules"
	DataDestroyParamsObjectsRecurring            DataDestroyParamsObjects = "recurring"
	DataDestroyParamsObjectsCategories           DataDestroyParamsObjects = "categories"
	DataDestroyParamsObjectsTags                 DataDestroyParamsObjects = "tags"
	DataDestroyParamsObjectsObjectGroups         DataDestroyParamsObjects = "object_groups"
	DataDestroyParamsObjectsAccounts             DataDestroyParamsObjects = "accounts"
	DataDestroyParamsObjectsAssetAccounts        DataDestroyParamsObjects = "asset_accounts"
	DataDestroyParamsObjectsExpenseAccounts      DataDestroyParamsObjects = "expense_accounts"
	DataDestroyParamsObjectsRevenueAccounts      DataDestroyParamsObjects = "revenue_accounts"
	DataDestroyParamsObjectsLiabilities          DataDestroyParamsObjects = "liabilities"
	DataDestroyParamsObjectsTransactions         DataDestroyParamsObjects = "transactions"
	DataDestroyParamsObjectsWithdrawals          DataDestroyParamsObjects = "withdrawals"
	DataDestroyParamsObjectsDeposits             DataDestroyParamsObjects = "deposits"
	DataDestroyParamsObjectsTransfers            DataDestroyParamsObjects = "transfers"
)

type DataPurgeParams struct {
	XTraceID param.Opt[string] `header:"X-Trace-Id,omitzero" format:"uuid" json:"-"`
	paramObj
}
