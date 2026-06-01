// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package firefly

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

// These endpoints deliver general system information, version- and meta
// information.
//
// CronService contains methods and other services that help with interacting with
// the firefly API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCronService] method instead.
type CronService struct {
	options []option.RequestOption
}

// NewCronService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewCronService(opts ...option.RequestOption) (r CronService) {
	r = CronService{}
	r.options = opts
	return
}

// Firefly III has one endpoint for its various cron related tasks. Send a GET to
// this endpoint to run the cron. The cron requires the CLI token to be present.
// The cron job will fire for all users.
func (r *CronService) Run(ctx context.Context, cliToken string, params CronRunParams, opts ...option.RequestOption) (res *CronRunResponse, err error) {
	if !param.IsOmitted(params.XTraceID) {
		opts = append(opts, option.WithHeader("X-Trace-Id", fmt.Sprintf("%v", params.XTraceID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	if cliToken == "" {
		err = errors.New("missing required cliToken parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/cron/%s", url.PathEscape(cliToken))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

type CronResultRow struct {
	// If the cron job ran into some kind of an error, this value will be true.
	JobErrored bool `json:"job_errored" api:"nullable"`
	// This value tells you if this specific cron job actually fired. It may not fire.
	// Some cron jobs only fire every 24 hours, for example.
	JobFired bool `json:"job_fired" api:"nullable"`
	// This value tells you if this specific cron job actually did something. The job
	// may fire but not change anything.
	JobSucceeded bool `json:"job_succeeded" api:"nullable"`
	// If the cron job ran into some kind of an error, this value will be the error
	// message. The success message if the job actually ran OK.
	Message string `json:"message" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		JobErrored   respjson.Field
		JobFired     respjson.Field
		JobSucceeded respjson.Field
		Message      respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CronResultRow) RawJSON() string { return r.JSON.raw }
func (r *CronResultRow) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CronRunResponse struct {
	AutoBudgets           CronResultRow `json:"auto_budgets"`
	RecurringTransactions CronResultRow `json:"recurring_transactions"`
	Telemetry             CronResultRow `json:"telemetry"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AutoBudgets           respjson.Field
		RecurringTransactions respjson.Field
		Telemetry             respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CronRunResponse) RawJSON() string { return r.JSON.raw }
func (r *CronRunResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CronRunParams struct {
	// A date formatted YYYY-MM-DD. This can be used to make the cron job pretend it's
	// running on another day.
	Date param.Opt[time.Time] `query:"date,omitzero" format:"date" json:"-"`
	// Forces the cron job to fire, regardless of whether it has fired before. This may
	// result in double transactions or weird budgets, so be careful.
	Force    param.Opt[bool]   `query:"force,omitzero" json:"-"`
	XTraceID param.Opt[string] `header:"X-Trace-Id,omitzero" format:"uuid" json:"-"`
	paramObj
}

// URLQuery serializes [CronRunParams]'s query parameters as `url.Values`.
func (r CronRunParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
