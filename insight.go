// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package firefly

import (
	"github.com/stainless-sdks/emcees-prod-testing-5-go/option"
)

// InsightService contains methods and other services that help with interacting
// with the firefly API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewInsightService] method instead.
type InsightService struct {
	options []option.RequestOption
	// The &quot;insight&quot; endpoints try to deliver sums, balances and insightful
	// information in the broadest sense of the word.
	Expense InsightExpenseService
	// The &quot;insight&quot; endpoints try to deliver sums, balances and insightful
	// information in the broadest sense of the word.
	Income InsightIncomeService
	// The &quot;insight&quot; endpoints try to deliver sums, balances and insightful
	// information in the broadest sense of the word.
	Transfer InsightTransferService
}

// NewInsightService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewInsightService(opts ...option.RequestOption) (r InsightService) {
	r = InsightService{}
	r.options = opts
	r.Expense = NewInsightExpenseService(opts...)
	r.Income = NewInsightIncomeService(opts...)
	r.Transfer = NewInsightTransferService(opts...)
	return
}
