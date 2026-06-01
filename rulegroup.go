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

// Manage all of the user&#039;s groups of rules and trigger the execution of
// entire groups.
//
// RuleGroupService contains methods and other services that help with interacting
// with the firefly API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRuleGroupService] method instead.
type RuleGroupService struct {
	options []option.RequestOption
}

// NewRuleGroupService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewRuleGroupService(opts ...option.RequestOption) (r RuleGroupService) {
	r = RuleGroupService{}
	r.options = opts
	return
}

// Creates a new rule group. The data required can be submitted as a JSON body or
// as a list of parameters.
func (r *RuleGroupService) New(ctx context.Context, params RuleGroupNewParams, opts ...option.RequestOption) (res *RuleGroupSingle, err error) {
	if !param.IsOmitted(params.XTraceID) {
		opts = append(opts, option.WithHeader("X-Trace-Id", fmt.Sprintf("%v", params.XTraceID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/vnd.api+json")}, opts...)
	path := "v1/rule-groups"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Get a single rule group. This does not include the rules. For that, see below.
func (r *RuleGroupService) Get(ctx context.Context, id string, query RuleGroupGetParams, opts ...option.RequestOption) (res *RuleGroupSingle, err error) {
	if !param.IsOmitted(query.XTraceID) {
		opts = append(opts, option.WithHeader("X-Trace-Id", fmt.Sprintf("%v", query.XTraceID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/vnd.api+json")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/rule-groups/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Update existing rule group.
func (r *RuleGroupService) Update(ctx context.Context, id string, params RuleGroupUpdateParams, opts ...option.RequestOption) (res *RuleGroupSingle, err error) {
	if !param.IsOmitted(params.XTraceID) {
		opts = append(opts, option.WithHeader("X-Trace-Id", fmt.Sprintf("%v", params.XTraceID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/vnd.api+json")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/rule-groups/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return res, err
}

// Delete a rule group.
func (r *RuleGroupService) Delete(ctx context.Context, id string, body RuleGroupDeleteParams, opts ...option.RequestOption) (err error) {
	if !param.IsOmitted(body.XTraceID) {
		opts = append(opts, option.WithHeader("X-Trace-Id", fmt.Sprintf("%v", body.XTraceID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return err
	}
	path := fmt.Sprintf("v1/rule-groups/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// List all rule groups.
func (r *RuleGroupService) ListAll(ctx context.Context, params RuleGroupListAllParams, opts ...option.RequestOption) (res *RuleGroupListAllResponse, err error) {
	if !param.IsOmitted(params.XTraceID) {
		opts = append(opts, option.WithHeader("X-Trace-Id", fmt.Sprintf("%v", params.XTraceID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/vnd.api+json")}, opts...)
	path := "v1/rule-groups"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

// List rules in this rule group.
func (r *RuleGroupService) ListRules(ctx context.Context, id string, params RuleGroupListRulesParams, opts ...option.RequestOption) (res *RuleArray, err error) {
	if !param.IsOmitted(params.XTraceID) {
		opts = append(opts, option.WithHeader("X-Trace-Id", fmt.Sprintf("%v", params.XTraceID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/vnd.api+json")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/rule-groups/%s/rules", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

// Test which transactions would be hit by the rule group. No changes will be made.
// Limit the result if you want to.
func (r *RuleGroupService) TestTransactions(ctx context.Context, id string, params RuleGroupTestTransactionsParams, opts ...option.RequestOption) (res *TransactionArray, err error) {
	if !param.IsOmitted(params.XTraceID) {
		opts = append(opts, option.WithHeader("X-Trace-Id", fmt.Sprintf("%v", params.XTraceID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/vnd.api+json")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/rule-groups/%s/test", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

// Fire the rule group on your transactions. Changes will be made by the rules in
// the rule group. Limit the result if you want to.
func (r *RuleGroupService) TriggerRules(ctx context.Context, id string, params RuleGroupTriggerRulesParams, opts ...option.RequestOption) (err error) {
	if !param.IsOmitted(params.XTraceID) {
		opts = append(opts, option.WithHeader("X-Trace-Id", fmt.Sprintf("%v", params.XTraceID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return err
	}
	path := fmt.Sprintf("v1/rule-groups/%s/trigger", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, nil, opts...)
	return err
}

type RuleGroupRead struct {
	ID         string                  `json:"id" api:"required"`
	Attributes RuleGroupReadAttributes `json:"attributes" api:"required"`
	Links      ObjectLink              `json:"links" api:"required"`
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
func (r RuleGroupRead) RawJSON() string { return r.JSON.raw }
func (r *RuleGroupRead) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RuleGroupReadAttributes struct {
	Title       string    `json:"title" api:"required"`
	Active      bool      `json:"active"`
	CreatedAt   time.Time `json:"created_at" format:"date-time"`
	Description string    `json:"description" api:"nullable"`
	Order       int64     `json:"order"`
	UpdatedAt   time.Time `json:"updated_at" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Title       respjson.Field
		Active      respjson.Field
		CreatedAt   respjson.Field
		Description respjson.Field
		Order       respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RuleGroupReadAttributes) RawJSON() string { return r.JSON.raw }
func (r *RuleGroupReadAttributes) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RuleGroupSingle struct {
	Data RuleGroupRead `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RuleGroupSingle) RawJSON() string { return r.JSON.raw }
func (r *RuleGroupSingle) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RuleGroupListAllResponse struct {
	Data  []RuleGroupRead `json:"data" api:"required"`
	Links PageLink        `json:"links" api:"required"`
	Meta  Meta            `json:"meta" api:"required"`
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
func (r RuleGroupListAllResponse) RawJSON() string { return r.JSON.raw }
func (r *RuleGroupListAllResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RuleGroupNewParams struct {
	Title       string            `json:"title" api:"required"`
	Description param.Opt[string] `json:"description,omitzero"`
	Active      param.Opt[bool]   `json:"active,omitzero"`
	Order       param.Opt[int64]  `json:"order,omitzero"`
	XTraceID    param.Opt[string] `header:"X-Trace-Id,omitzero" format:"uuid" json:"-"`
	paramObj
}

func (r RuleGroupNewParams) MarshalJSON() (data []byte, err error) {
	type shadow RuleGroupNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RuleGroupNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RuleGroupGetParams struct {
	XTraceID param.Opt[string] `header:"X-Trace-Id,omitzero" format:"uuid" json:"-"`
	paramObj
}

type RuleGroupUpdateParams struct {
	Description param.Opt[string] `json:"description,omitzero"`
	Active      param.Opt[bool]   `json:"active,omitzero"`
	Order       param.Opt[int64]  `json:"order,omitzero"`
	Title       param.Opt[string] `json:"title,omitzero"`
	XTraceID    param.Opt[string] `header:"X-Trace-Id,omitzero" format:"uuid" json:"-"`
	paramObj
}

func (r RuleGroupUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow RuleGroupUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RuleGroupUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RuleGroupDeleteParams struct {
	XTraceID param.Opt[string] `header:"X-Trace-Id,omitzero" format:"uuid" json:"-"`
	paramObj
}

type RuleGroupListAllParams struct {
	// Number of items per page. The default pagination is per 50 items.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Page number. The default pagination is per 50 items.
	Page     param.Opt[int64]  `query:"page,omitzero" json:"-"`
	XTraceID param.Opt[string] `header:"X-Trace-Id,omitzero" format:"uuid" json:"-"`
	paramObj
}

// URLQuery serializes [RuleGroupListAllParams]'s query parameters as `url.Values`.
func (r RuleGroupListAllParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RuleGroupListRulesParams struct {
	// Number of items per page. The default pagination is per 50 items.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Page number. The default pagination is per 50 items.
	Page     param.Opt[int64]  `query:"page,omitzero" json:"-"`
	XTraceID param.Opt[string] `header:"X-Trace-Id,omitzero" format:"uuid" json:"-"`
	paramObj
}

// URLQuery serializes [RuleGroupListRulesParams]'s query parameters as
// `url.Values`.
func (r RuleGroupListRulesParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RuleGroupTestTransactionsParams struct {
	// A date formatted YYYY-MM-DD, to limit the transactions the test will be applied
	// to. Both the start date and the end date must be present.
	End param.Opt[time.Time] `query:"end,omitzero" format:"date" json:"-"`
	// Number of items per page. The default pagination is per 50 items.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Page number. The default pagination is per 50 items.
	Page param.Opt[int64] `query:"page,omitzero" json:"-"`
	// Maximum number of transactions Firefly III will try. Don't set this too high, or
	// it will take Firefly III very long to run the test. I suggest a max of 200.
	SearchLimit param.Opt[int64] `query:"search_limit,omitzero" json:"-"`
	// A date formatted YYYY-MM-DD, to limit the transactions the test will be applied
	// to. Both the start date and the end date must be present.
	Start param.Opt[time.Time] `query:"start,omitzero" format:"date" json:"-"`
	// Maximum number of transactions the rule group can actually trigger on, before
	// Firefly III stops. I would suggest setting this to 10 or 15. Don't go above the
	// user's page size, because browsing to page 2 or 3 of a test result would fire
	// the test again, making any navigation efforts very slow.
	TriggeredLimit param.Opt[int64]  `query:"triggered_limit,omitzero" json:"-"`
	XTraceID       param.Opt[string] `header:"X-Trace-Id,omitzero" format:"uuid" json:"-"`
	// Limit the testing of the rule group to these asset accounts or liabilities. Only
	// asset accounts and liabilities will be accepted. Other types will be silently
	// dropped.
	Accounts []int64 `query:"accounts,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [RuleGroupTestTransactionsParams]'s query parameters as
// `url.Values`.
func (r RuleGroupTestTransactionsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RuleGroupTriggerRulesParams struct {
	// A date formatted YYYY-MM-DD, to limit the transactions the actions will be
	// applied to. Both the start date and the end date must be present.
	End param.Opt[time.Time] `query:"end,omitzero" format:"date" json:"-"`
	// A date formatted YYYY-MM-DD, to limit the transactions the actions will be
	// applied to. Both the start date and the end date must be present.
	Start    param.Opt[time.Time] `query:"start,omitzero" format:"date" json:"-"`
	XTraceID param.Opt[string]    `header:"X-Trace-Id,omitzero" format:"uuid" json:"-"`
	// Limit the triggering of the rule group to these asset accounts or liabilities.
	// Only asset accounts and liabilities will be accepted. Other types will be
	// silently dropped.
	Accounts []int64 `query:"accounts,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [RuleGroupTriggerRulesParams]'s query parameters as
// `url.Values`.
func (r RuleGroupTriggerRulesParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
