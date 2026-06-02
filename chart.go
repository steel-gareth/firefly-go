// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package firefly

import (
	"github.com/steel-gareth/firefly-go/option"
)

// ChartService contains methods and other services that help with interacting with
// the firefly API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewChartService] method instead.
type ChartService struct {
	options []option.RequestOption
	// The &quot;charts&quot; endpoints deliver optimised data for charts and graphs.
	Account ChartAccountService
	// The &quot;charts&quot; endpoints deliver optimised data for charts and graphs.
	Balance ChartBalanceService
	// The &quot;charts&quot; endpoints deliver optimised data for charts and graphs.
	Budget ChartBudgetService
	// The &quot;charts&quot; endpoints deliver optimised data for charts and graphs.
	Category ChartCategoryService
}

// NewChartService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewChartService(opts ...option.RequestOption) (r ChartService) {
	r = ChartService{}
	r.options = opts
	r.Account = NewChartAccountService(opts...)
	r.Balance = NewChartBalanceService(opts...)
	r.Budget = NewChartBudgetService(opts...)
	r.Category = NewChartCategoryService(opts...)
	return
}
