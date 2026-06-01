// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package firefly

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/stainless-sdks/emcees-prod-testing-5-go/internal/requestconfig"
	"github.com/stainless-sdks/emcees-prod-testing-5-go/option"
	"github.com/stainless-sdks/emcees-prod-testing-5-go/packages/param"
)

// Endpoints to manage the currencies in Firefly III. Depending on the user&#039;s
// role you can also disable and enable them, or add new ones.
//
// CurrencyPrimaryService contains methods and other services that help with
// interacting with the firefly API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCurrencyPrimaryService] method instead.
type CurrencyPrimaryService struct {
	options []option.RequestOption
}

// NewCurrencyPrimaryService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewCurrencyPrimaryService(opts ...option.RequestOption) (r CurrencyPrimaryService) {
	r = CurrencyPrimaryService{}
	r.options = opts
	return
}

// Get the primary currency of the current administration. This replaces what was
// called "the user's default currency" although they are essentially the same.
func (r *CurrencyPrimaryService) Get(ctx context.Context, query CurrencyPrimaryGetParams, opts ...option.RequestOption) (res *CurrencySingle, err error) {
	if !param.IsOmitted(query.XTraceID) {
		opts = append(opts, option.WithHeader("X-Trace-Id", fmt.Sprintf("%v", query.XTraceID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	path := "v1/currencies/primary"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Make this currency the primary currency for the current financial
// administration. If the currency is not enabled, it will be enabled as well.
func (r *CurrencyPrimaryService) MakePrimary(ctx context.Context, code string, body CurrencyPrimaryMakePrimaryParams, opts ...option.RequestOption) (res *CurrencySingle, err error) {
	if !param.IsOmitted(body.XTraceID) {
		opts = append(opts, option.WithHeader("X-Trace-Id", fmt.Sprintf("%v", body.XTraceID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/vnd.api+json")}, opts...)
	if code == "" {
		err = errors.New("missing required code parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/currencies/%s/primary", url.PathEscape(code))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

type CurrencyPrimaryGetParams struct {
	XTraceID param.Opt[string] `header:"X-Trace-Id,omitzero" format:"uuid" json:"-"`
	paramObj
}

type CurrencyPrimaryMakePrimaryParams struct {
	XTraceID param.Opt[string] `header:"X-Trace-Id,omitzero" format:"uuid" json:"-"`
	paramObj
}
