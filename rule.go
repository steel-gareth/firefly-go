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

	"github.com/steel-gareth/firefly-go/internal/apijson"
	"github.com/steel-gareth/firefly-go/internal/apiquery"
	"github.com/steel-gareth/firefly-go/internal/requestconfig"
	"github.com/steel-gareth/firefly-go/option"
	"github.com/steel-gareth/firefly-go/packages/param"
	"github.com/steel-gareth/firefly-go/packages/respjson"
)

// These endpoints can be used to manage all of the user&#039;s rules. Also
// includes triggers to execute or test rules and individual triggers.
//
// RuleService contains methods and other services that help with interacting with
// the firefly API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRuleService] method instead.
type RuleService struct {
	options []option.RequestOption
}

// NewRuleService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewRuleService(opts ...option.RequestOption) (r RuleService) {
	r = RuleService{}
	r.options = opts
	return
}

// Creates a new rule. The data required can be submitted as a JSON body or as a
// list of parameters.
func (r *RuleService) New(ctx context.Context, params RuleNewParams, opts ...option.RequestOption) (res *RuleSingle, err error) {
	if !param.IsOmitted(params.XTraceID) {
		opts = append(opts, option.WithHeader("X-Trace-Id", fmt.Sprintf("%v", params.XTraceID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/vnd.api+json")}, opts...)
	path := "v1/rules"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Get a single rule.
func (r *RuleService) Get(ctx context.Context, id string, query RuleGetParams, opts ...option.RequestOption) (res *RuleSingle, err error) {
	if !param.IsOmitted(query.XTraceID) {
		opts = append(opts, option.WithHeader("X-Trace-Id", fmt.Sprintf("%v", query.XTraceID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/vnd.api+json")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/rules/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Update existing rule.
func (r *RuleService) Update(ctx context.Context, id string, params RuleUpdateParams, opts ...option.RequestOption) (res *RuleSingle, err error) {
	if !param.IsOmitted(params.XTraceID) {
		opts = append(opts, option.WithHeader("X-Trace-Id", fmt.Sprintf("%v", params.XTraceID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/vnd.api+json")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/rules/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return res, err
}

// List all rules.
func (r *RuleService) List(ctx context.Context, params RuleListParams, opts ...option.RequestOption) (res *RuleArray, err error) {
	if !param.IsOmitted(params.XTraceID) {
		opts = append(opts, option.WithHeader("X-Trace-Id", fmt.Sprintf("%v", params.XTraceID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/vnd.api+json")}, opts...)
	path := "v1/rules"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

// Delete an rule.
func (r *RuleService) Delete(ctx context.Context, id string, body RuleDeleteParams, opts ...option.RequestOption) (err error) {
	if !param.IsOmitted(body.XTraceID) {
		opts = append(opts, option.WithHeader("X-Trace-Id", fmt.Sprintf("%v", body.XTraceID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return err
	}
	path := fmt.Sprintf("v1/rules/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// Test which transactions would be hit by the rule. No changes will be made. Limit
// the result if you want to.
func (r *RuleService) Test(ctx context.Context, id string, params RuleTestParams, opts ...option.RequestOption) (res *TransactionArray, err error) {
	if !param.IsOmitted(params.XTraceID) {
		opts = append(opts, option.WithHeader("X-Trace-Id", fmt.Sprintf("%v", params.XTraceID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/vnd.api+json")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/rules/%s/test", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

// Fire the rule group on your transactions. Changes will be made by the rules in
// the group. Limit the result if you want to.
func (r *RuleService) Trigger(ctx context.Context, id string, params RuleTriggerParams, opts ...option.RequestOption) (err error) {
	if !param.IsOmitted(params.XTraceID) {
		opts = append(opts, option.WithHeader("X-Trace-Id", fmt.Sprintf("%v", params.XTraceID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return err
	}
	path := fmt.Sprintf("v1/rules/%s/trigger", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, nil, opts...)
	return err
}

// The type of thing this action will do. A limited set is possible.
type RuleActionKeyword string

const (
	RuleActionKeywordUserAction            RuleActionKeyword = "user_action"
	RuleActionKeywordSetCategory           RuleActionKeyword = "set_category"
	RuleActionKeywordClearCategory         RuleActionKeyword = "clear_category"
	RuleActionKeywordSetBudget             RuleActionKeyword = "set_budget"
	RuleActionKeywordClearBudget           RuleActionKeyword = "clear_budget"
	RuleActionKeywordAddTag                RuleActionKeyword = "add_tag"
	RuleActionKeywordRemoveTag             RuleActionKeyword = "remove_tag"
	RuleActionKeywordRemoveAllTags         RuleActionKeyword = "remove_all_tags"
	RuleActionKeywordSetDescription        RuleActionKeyword = "set_description"
	RuleActionKeywordAppendDescription     RuleActionKeyword = "append_description"
	RuleActionKeywordPrependDescription    RuleActionKeyword = "prepend_description"
	RuleActionKeywordSetSourceAccount      RuleActionKeyword = "set_source_account"
	RuleActionKeywordSetDestinationAccount RuleActionKeyword = "set_destination_account"
	RuleActionKeywordSetNotes              RuleActionKeyword = "set_notes"
	RuleActionKeywordAppendNotes           RuleActionKeyword = "append_notes"
	RuleActionKeywordPrependNotes          RuleActionKeyword = "prepend_notes"
	RuleActionKeywordClearNotes            RuleActionKeyword = "clear_notes"
	RuleActionKeywordLinkToBill            RuleActionKeyword = "link_to_bill"
	RuleActionKeywordConvertWithdrawal     RuleActionKeyword = "convert_withdrawal"
	RuleActionKeywordConvertDeposit        RuleActionKeyword = "convert_deposit"
	RuleActionKeywordConvertTransfer       RuleActionKeyword = "convert_transfer"
	RuleActionKeywordDeleteTransaction     RuleActionKeyword = "delete_transaction"
)

type RuleRead struct {
	ID         string             `json:"id" api:"required"`
	Attributes RuleReadAttributes `json:"attributes" api:"required"`
	Links      ObjectLink         `json:"links" api:"required"`
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
func (r RuleRead) RawJSON() string { return r.JSON.raw }
func (r *RuleRead) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RuleReadAttributes struct {
	Actions []RuleReadAttributesAction `json:"actions" api:"required"`
	// ID of the rule group under which the rule must be stored. Either this field or
	// rule_group_title is mandatory.
	RuleGroupID string `json:"rule_group_id" api:"required"`
	Title       string `json:"title" api:"required"`
	// Which action is necessary for the rule to fire? Use either store-journal,
	// update-journal or manual-activation.
	//
	// Any of "store-journal", "update-journal", "manual-activation".
	Trigger  RuleTriggerType             `json:"trigger" api:"required"`
	Triggers []RuleReadAttributesTrigger `json:"triggers" api:"required"`
	// Whether or not the rule is even active. Default is true.
	Active      bool      `json:"active"`
	CreatedAt   time.Time `json:"created_at" format:"date-time"`
	Description string    `json:"description"`
	Order       int64     `json:"order"`
	// Title of the rule group under which the rule must be stored. Either this field
	// or rule_group_id is mandatory.
	RuleGroupTitle string `json:"rule_group_title"`
	// If this value is true and the rule is triggered, other rules after this one in
	// the group will be skipped. Default value is false.
	StopProcessing bool `json:"stop_processing"`
	// If the rule is set to be strict, ALL triggers must hit in order for the rule to
	// fire. Otherwise, just one is enough. Default value is true.
	Strict    bool      `json:"strict"`
	UpdatedAt time.Time `json:"updated_at" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Actions        respjson.Field
		RuleGroupID    respjson.Field
		Title          respjson.Field
		Trigger        respjson.Field
		Triggers       respjson.Field
		Active         respjson.Field
		CreatedAt      respjson.Field
		Description    respjson.Field
		Order          respjson.Field
		RuleGroupTitle respjson.Field
		StopProcessing respjson.Field
		Strict         respjson.Field
		UpdatedAt      respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RuleReadAttributes) RawJSON() string { return r.JSON.raw }
func (r *RuleReadAttributes) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RuleReadAttributesAction struct {
	// The type of thing this action will do. A limited set is possible.
	//
	// Any of "user_action", "set_category", "clear_category", "set_budget",
	// "clear_budget", "add_tag", "remove_tag", "remove_all_tags", "set_description",
	// "append_description", "prepend_description", "set_source_account",
	// "set_destination_account", "set_notes", "append_notes", "prepend_notes",
	// "clear_notes", "link_to_bill", "convert_withdrawal", "convert_deposit",
	// "convert_transfer", "delete_transaction".
	Type RuleActionKeyword `json:"type" api:"required"`
	// The accompanying value the action will set, change or update. Can be empty, but
	// for some types this value is mandatory.
	Value string `json:"value" api:"required"`
	ID    string `json:"id"`
	// If the action is active. Defaults to true.
	Active    bool      `json:"active"`
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Order of the action
	Order int64 `json:"order"`
	// When true, other actions will not be fired after this action has fired. Defaults
	// to false.
	StopProcessing bool      `json:"stop_processing"`
	UpdatedAt      time.Time `json:"updated_at" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type           respjson.Field
		Value          respjson.Field
		ID             respjson.Field
		Active         respjson.Field
		CreatedAt      respjson.Field
		Order          respjson.Field
		StopProcessing respjson.Field
		UpdatedAt      respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RuleReadAttributesAction) RawJSON() string { return r.JSON.raw }
func (r *RuleReadAttributesAction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RuleReadAttributesTrigger struct {
	// The type of thing this trigger responds to. A limited set is possible
	//
	// Any of "from_account_starts", "from_account_ends", "from_account_is",
	// "from_account_contains", "to_account_starts", "to_account_ends",
	// "to_account_is", "to_account_contains", "amount_less", "amount_exactly",
	// "amount_more", "description_starts", "description_ends", "description_contains",
	// "description_is", "transaction_type", "category_is", "budget_is", "tag_is",
	// "currency_is", "has_attachments", "has_no_category", "has_any_category",
	// "has_no_budget", "has_any_budget", "has_no_tag", "has_any_tag",
	// "notes_contains", "notes_starts", "notes_end", "notes_are", "no_notes",
	// "any_notes", "source_account_is", "destination_account_is",
	// "source_account_starts".
	Type RuleTriggerKeyword `json:"type" api:"required"`
	// The accompanying value the trigger responds to. This value is often mandatory,
	// but this depends on the trigger.
	Value string `json:"value" api:"required"`
	ID    string `json:"id"`
	// If the trigger is active. Defaults to true.
	Active    bool      `json:"active"`
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Order of the trigger
	Order int64 `json:"order"`
	// If 'prohibited' is true, this rule trigger will be negated. 'Description is'
	// will become 'Description is NOT' etc.
	Prohibited bool `json:"prohibited"`
	// When true, other triggers will not be checked if this trigger was triggered.
	// Defaults to false.
	StopProcessing bool      `json:"stop_processing"`
	UpdatedAt      time.Time `json:"updated_at" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type           respjson.Field
		Value          respjson.Field
		ID             respjson.Field
		Active         respjson.Field
		CreatedAt      respjson.Field
		Order          respjson.Field
		Prohibited     respjson.Field
		StopProcessing respjson.Field
		UpdatedAt      respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RuleReadAttributesTrigger) RawJSON() string { return r.JSON.raw }
func (r *RuleReadAttributesTrigger) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RuleSingle struct {
	Data RuleRead `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RuleSingle) RawJSON() string { return r.JSON.raw }
func (r *RuleSingle) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The type of thing this trigger responds to. A limited set is possible
type RuleTriggerKeyword string

const (
	RuleTriggerKeywordFromAccountStarts    RuleTriggerKeyword = "from_account_starts"
	RuleTriggerKeywordFromAccountEnds      RuleTriggerKeyword = "from_account_ends"
	RuleTriggerKeywordFromAccountIs        RuleTriggerKeyword = "from_account_is"
	RuleTriggerKeywordFromAccountContains  RuleTriggerKeyword = "from_account_contains"
	RuleTriggerKeywordToAccountStarts      RuleTriggerKeyword = "to_account_starts"
	RuleTriggerKeywordToAccountEnds        RuleTriggerKeyword = "to_account_ends"
	RuleTriggerKeywordToAccountIs          RuleTriggerKeyword = "to_account_is"
	RuleTriggerKeywordToAccountContains    RuleTriggerKeyword = "to_account_contains"
	RuleTriggerKeywordAmountLess           RuleTriggerKeyword = "amount_less"
	RuleTriggerKeywordAmountExactly        RuleTriggerKeyword = "amount_exactly"
	RuleTriggerKeywordAmountMore           RuleTriggerKeyword = "amount_more"
	RuleTriggerKeywordDescriptionStarts    RuleTriggerKeyword = "description_starts"
	RuleTriggerKeywordDescriptionEnds      RuleTriggerKeyword = "description_ends"
	RuleTriggerKeywordDescriptionContains  RuleTriggerKeyword = "description_contains"
	RuleTriggerKeywordDescriptionIs        RuleTriggerKeyword = "description_is"
	RuleTriggerKeywordTransactionType      RuleTriggerKeyword = "transaction_type"
	RuleTriggerKeywordCategoryIs           RuleTriggerKeyword = "category_is"
	RuleTriggerKeywordBudgetIs             RuleTriggerKeyword = "budget_is"
	RuleTriggerKeywordTagIs                RuleTriggerKeyword = "tag_is"
	RuleTriggerKeywordCurrencyIs           RuleTriggerKeyword = "currency_is"
	RuleTriggerKeywordHasAttachments       RuleTriggerKeyword = "has_attachments"
	RuleTriggerKeywordHasNoCategory        RuleTriggerKeyword = "has_no_category"
	RuleTriggerKeywordHasAnyCategory       RuleTriggerKeyword = "has_any_category"
	RuleTriggerKeywordHasNoBudget          RuleTriggerKeyword = "has_no_budget"
	RuleTriggerKeywordHasAnyBudget         RuleTriggerKeyword = "has_any_budget"
	RuleTriggerKeywordHasNoTag             RuleTriggerKeyword = "has_no_tag"
	RuleTriggerKeywordHasAnyTag            RuleTriggerKeyword = "has_any_tag"
	RuleTriggerKeywordNotesContains        RuleTriggerKeyword = "notes_contains"
	RuleTriggerKeywordNotesStarts          RuleTriggerKeyword = "notes_starts"
	RuleTriggerKeywordNotesEnd             RuleTriggerKeyword = "notes_end"
	RuleTriggerKeywordNotesAre             RuleTriggerKeyword = "notes_are"
	RuleTriggerKeywordNoNotes              RuleTriggerKeyword = "no_notes"
	RuleTriggerKeywordAnyNotes             RuleTriggerKeyword = "any_notes"
	RuleTriggerKeywordSourceAccountIs      RuleTriggerKeyword = "source_account_is"
	RuleTriggerKeywordDestinationAccountIs RuleTriggerKeyword = "destination_account_is"
	RuleTriggerKeywordSourceAccountStarts  RuleTriggerKeyword = "source_account_starts"
)

// Which action is necessary for the rule to fire? Use either store-journal,
// update-journal or manual-activation.
type RuleTriggerType string

const (
	RuleTriggerTypeStoreJournal     RuleTriggerType = "store-journal"
	RuleTriggerTypeUpdateJournal    RuleTriggerType = "update-journal"
	RuleTriggerTypeManualActivation RuleTriggerType = "manual-activation"
)

type RuleNewParams struct {
	Actions []RuleNewParamsAction `json:"actions,omitzero" api:"required"`
	// ID of the rule group under which the rule must be stored. Either this field or
	// rule_group_title is mandatory.
	RuleGroupID string `json:"rule_group_id" api:"required"`
	Title       string `json:"title" api:"required"`
	// Which action is necessary for the rule to fire? Use either store-journal,
	// update-journal or manual-activation.
	//
	// Any of "store-journal", "update-journal", "manual-activation".
	Trigger  RuleTriggerType        `json:"trigger,omitzero" api:"required"`
	Triggers []RuleNewParamsTrigger `json:"triggers,omitzero" api:"required"`
	// Whether or not the rule is even active. Default is true.
	Active      param.Opt[bool]   `json:"active,omitzero"`
	Description param.Opt[string] `json:"description,omitzero"`
	Order       param.Opt[int64]  `json:"order,omitzero"`
	// Title of the rule group under which the rule must be stored. Either this field
	// or rule_group_id is mandatory.
	RuleGroupTitle param.Opt[string] `json:"rule_group_title,omitzero"`
	// If this value is true and the rule is triggered, other rules after this one in
	// the group will be skipped. Default value is false.
	StopProcessing param.Opt[bool] `json:"stop_processing,omitzero"`
	// If the rule is set to be strict, ALL triggers must hit in order for the rule to
	// fire. Otherwise, just one is enough. Default value is true.
	Strict   param.Opt[bool]   `json:"strict,omitzero"`
	XTraceID param.Opt[string] `header:"X-Trace-Id,omitzero" format:"uuid" json:"-"`
	paramObj
}

func (r RuleNewParams) MarshalJSON() (data []byte, err error) {
	type shadow RuleNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RuleNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Type, Value are required.
type RuleNewParamsAction struct {
	// The accompanying value the action will set, change or update. Can be empty, but
	// for some types this value is mandatory.
	Value param.Opt[string] `json:"value,omitzero" api:"required"`
	// The type of thing this action will do. A limited set is possible.
	//
	// Any of "user_action", "set_category", "clear_category", "set_budget",
	// "clear_budget", "add_tag", "remove_tag", "remove_all_tags", "set_description",
	// "append_description", "prepend_description", "set_source_account",
	// "set_destination_account", "set_notes", "append_notes", "prepend_notes",
	// "clear_notes", "link_to_bill", "convert_withdrawal", "convert_deposit",
	// "convert_transfer", "delete_transaction".
	Type RuleActionKeyword `json:"type,omitzero" api:"required"`
	// If the action is active. Defaults to true.
	Active param.Opt[bool] `json:"active,omitzero"`
	// Order of the action
	Order param.Opt[int64] `json:"order,omitzero"`
	// When true, other actions will not be fired after this action has fired. Defaults
	// to false.
	StopProcessing param.Opt[bool] `json:"stop_processing,omitzero"`
	paramObj
}

func (r RuleNewParamsAction) MarshalJSON() (data []byte, err error) {
	type shadow RuleNewParamsAction
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RuleNewParamsAction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Type, Value are required.
type RuleNewParamsTrigger struct {
	// The type of thing this trigger responds to. A limited set is possible
	//
	// Any of "from_account_starts", "from_account_ends", "from_account_is",
	// "from_account_contains", "to_account_starts", "to_account_ends",
	// "to_account_is", "to_account_contains", "amount_less", "amount_exactly",
	// "amount_more", "description_starts", "description_ends", "description_contains",
	// "description_is", "transaction_type", "category_is", "budget_is", "tag_is",
	// "currency_is", "has_attachments", "has_no_category", "has_any_category",
	// "has_no_budget", "has_any_budget", "has_no_tag", "has_any_tag",
	// "notes_contains", "notes_starts", "notes_end", "notes_are", "no_notes",
	// "any_notes", "source_account_is", "destination_account_is",
	// "source_account_starts".
	Type RuleTriggerKeyword `json:"type,omitzero" api:"required"`
	// The accompanying value the trigger responds to. This value is often mandatory,
	// but this depends on the trigger.
	Value string `json:"value" api:"required"`
	// If the trigger is active. Defaults to true.
	Active param.Opt[bool] `json:"active,omitzero"`
	// Order of the trigger
	Order param.Opt[int64] `json:"order,omitzero"`
	// If 'prohibited' is true, this rule trigger will be negated. 'Description is'
	// will become 'Description is NOT' etc.
	Prohibited param.Opt[bool] `json:"prohibited,omitzero"`
	// When true, other triggers will not be checked if this trigger was triggered.
	// Defaults to false.
	StopProcessing param.Opt[bool] `json:"stop_processing,omitzero"`
	paramObj
}

func (r RuleNewParamsTrigger) MarshalJSON() (data []byte, err error) {
	type shadow RuleNewParamsTrigger
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RuleNewParamsTrigger) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RuleGetParams struct {
	XTraceID param.Opt[string] `header:"X-Trace-Id,omitzero" format:"uuid" json:"-"`
	paramObj
}

type RuleUpdateParams struct {
	// Whether or not the rule is even active. Default is true.
	Active      param.Opt[bool]   `json:"active,omitzero"`
	Description param.Opt[string] `json:"description,omitzero"`
	Order       param.Opt[int64]  `json:"order,omitzero"`
	// ID of the rule group under which the rule must be stored. Either this field or
	// rule_group_title is mandatory.
	RuleGroupID param.Opt[string] `json:"rule_group_id,omitzero"`
	// If this value is true and the rule is triggered, other rules after this one in
	// the group will be skipped. Default value is false.
	StopProcessing param.Opt[bool] `json:"stop_processing,omitzero"`
	// If the rule is set to be strict, ALL triggers must hit in order for the rule to
	// fire. Otherwise, just one is enough. Default value is true.
	Strict   param.Opt[bool]          `json:"strict,omitzero"`
	Title    param.Opt[string]        `json:"title,omitzero"`
	XTraceID param.Opt[string]        `header:"X-Trace-Id,omitzero" format:"uuid" json:"-"`
	Actions  []RuleUpdateParamsAction `json:"actions,omitzero"`
	// Which action is necessary for the rule to fire? Use either store-journal,
	// update-journal or manual-activation.
	//
	// Any of "store-journal", "update-journal", "manual-activation".
	Trigger  RuleTriggerType           `json:"trigger,omitzero"`
	Triggers []RuleUpdateParamsTrigger `json:"triggers,omitzero"`
	paramObj
}

func (r RuleUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow RuleUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RuleUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RuleUpdateParamsAction struct {
	// The accompanying value the action will set, change or update. Can be empty, but
	// for some types this value is mandatory.
	Value param.Opt[string] `json:"value,omitzero"`
	// If the action is active.
	Active param.Opt[bool] `json:"active,omitzero"`
	// Order of the action
	Order param.Opt[int64] `json:"order,omitzero"`
	// When true, other actions will not be fired after this action has fired.
	StopProcessing param.Opt[bool] `json:"stop_processing,omitzero"`
	// The type of thing this action will do. A limited set is possible.
	//
	// Any of "user_action", "set_category", "clear_category", "set_budget",
	// "clear_budget", "add_tag", "remove_tag", "remove_all_tags", "set_description",
	// "append_description", "prepend_description", "set_source_account",
	// "set_destination_account", "set_notes", "append_notes", "prepend_notes",
	// "clear_notes", "link_to_bill", "convert_withdrawal", "convert_deposit",
	// "convert_transfer", "delete_transaction".
	Type RuleActionKeyword `json:"type,omitzero"`
	paramObj
}

func (r RuleUpdateParamsAction) MarshalJSON() (data []byte, err error) {
	type shadow RuleUpdateParamsAction
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RuleUpdateParamsAction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RuleUpdateParamsTrigger struct {
	// If the trigger is active.
	Active param.Opt[bool] `json:"active,omitzero"`
	// Order of the trigger
	Order param.Opt[int64] `json:"order,omitzero"`
	// When true, other triggers will not be checked if this trigger was triggered.
	StopProcessing param.Opt[bool] `json:"stop_processing,omitzero"`
	// The accompanying value the trigger responds to. This value is often mandatory,
	// but this depends on the trigger. If the rule trigger is something like 'has any
	// tag', submit the string 'true'.
	Value param.Opt[string] `json:"value,omitzero"`
	// The type of thing this trigger responds to. A limited set is possible
	//
	// Any of "from_account_starts", "from_account_ends", "from_account_is",
	// "from_account_contains", "to_account_starts", "to_account_ends",
	// "to_account_is", "to_account_contains", "amount_less", "amount_exactly",
	// "amount_more", "description_starts", "description_ends", "description_contains",
	// "description_is", "transaction_type", "category_is", "budget_is", "tag_is",
	// "currency_is", "has_attachments", "has_no_category", "has_any_category",
	// "has_no_budget", "has_any_budget", "has_no_tag", "has_any_tag",
	// "notes_contains", "notes_starts", "notes_end", "notes_are", "no_notes",
	// "any_notes", "source_account_is", "destination_account_is",
	// "source_account_starts".
	Type RuleTriggerKeyword `json:"type,omitzero"`
	paramObj
}

func (r RuleUpdateParamsTrigger) MarshalJSON() (data []byte, err error) {
	type shadow RuleUpdateParamsTrigger
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RuleUpdateParamsTrigger) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RuleListParams struct {
	// Number of items per page. The default pagination is per 50 items.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Page number. The default pagination is per 50 items.
	Page     param.Opt[int64]  `query:"page,omitzero" json:"-"`
	XTraceID param.Opt[string] `header:"X-Trace-Id,omitzero" format:"uuid" json:"-"`
	paramObj
}

// URLQuery serializes [RuleListParams]'s query parameters as `url.Values`.
func (r RuleListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RuleDeleteParams struct {
	XTraceID param.Opt[string] `header:"X-Trace-Id,omitzero" format:"uuid" json:"-"`
	paramObj
}

type RuleTestParams struct {
	// A date formatted YYYY-MM-DD, to limit the transactions the test will be applied
	// to. Both the start date and the end date must be present.
	End param.Opt[time.Time] `query:"end,omitzero" format:"date" json:"-"`
	// A date formatted YYYY-MM-DD, to limit the transactions the test will be applied
	// to. Both the start date and the end date must be present.
	Start    param.Opt[time.Time] `query:"start,omitzero" format:"date" json:"-"`
	XTraceID param.Opt[string]    `header:"X-Trace-Id,omitzero" format:"uuid" json:"-"`
	// Limit the testing of the rule to these asset accounts or liabilities. Only asset
	// accounts and liabilities will be accepted. Other types will be silently dropped.
	Accounts []int64 `query:"accounts,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [RuleTestParams]'s query parameters as `url.Values`.
func (r RuleTestParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RuleTriggerParams struct {
	// A date formatted YYYY-MM-DD, to limit the transactions the actions will be
	// applied to. If the end date is not present, it will be set to today. If you use
	// this field, both the start date and the end date must be present.
	End param.Opt[time.Time] `query:"end,omitzero" format:"date" json:"-"`
	// A date formatted YYYY-MM-DD, to limit the transactions the actions will be
	// applied to. If the start date is not present, it will be set to one year ago. If
	// you use this field, both the start date and the end date must be present.
	Start    param.Opt[time.Time] `query:"start,omitzero" format:"date" json:"-"`
	XTraceID param.Opt[string]    `header:"X-Trace-Id,omitzero" format:"uuid" json:"-"`
	// Limit the triggering of the rule to these asset accounts or liabilities. Only
	// asset accounts and liabilities will be accepted. Other types will be silently
	// dropped.
	Accounts []int64 `query:"accounts,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [RuleTriggerParams]'s query parameters as `url.Values`.
func (r RuleTriggerParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
