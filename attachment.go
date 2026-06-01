// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package emceesprodtesting5

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/stainless-sdks/emcees-prod-testing-5-go/internal/apiform"
	"github.com/stainless-sdks/emcees-prod-testing-5-go/internal/apijson"
	"github.com/stainless-sdks/emcees-prod-testing-5-go/internal/apiquery"
	"github.com/stainless-sdks/emcees-prod-testing-5-go/internal/requestconfig"
	"github.com/stainless-sdks/emcees-prod-testing-5-go/option"
	"github.com/stainless-sdks/emcees-prod-testing-5-go/packages/param"
	"github.com/stainless-sdks/emcees-prod-testing-5-go/packages/respjson"
)

// Endpoints to manage the attachments of the authenticated user, including up- and
// downloading of the files.
//
// AttachmentService contains methods and other services that help with interacting
// with the emcees-prod-testing-5 API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAttachmentService] method instead.
type AttachmentService struct {
	options []option.RequestOption
}

// NewAttachmentService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAttachmentService(opts ...option.RequestOption) (r AttachmentService) {
	r = AttachmentService{}
	r.options = opts
	return
}

// Creates a new attachment. The data required can be submitted as a JSON body or
// as a list of parameters. You cannot use this endpoint to upload the actual file
// data (see below). This endpoint only creates the attachment object.
func (r *AttachmentService) New(ctx context.Context, params AttachmentNewParams, opts ...option.RequestOption) (res *AttachmentSingle, err error) {
	if !param.IsOmitted(params.XTraceID) {
		opts = append(opts, option.WithHeader("X-Trace-Id", fmt.Sprintf("%v", params.XTraceID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/vnd.api+json")}, opts...)
	path := "v1/attachments"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Get a single attachment. This endpoint only returns the available metadata for
// the attachment. Actual file data is handled in two other endpoints (see below).
func (r *AttachmentService) Get(ctx context.Context, id string, query AttachmentGetParams, opts ...option.RequestOption) (res *AttachmentSingle, err error) {
	if !param.IsOmitted(query.XTraceID) {
		opts = append(opts, option.WithHeader("X-Trace-Id", fmt.Sprintf("%v", query.XTraceID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/vnd.api+json")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/attachments/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Update the meta data for an existing attachment. This endpoint does not allow
// you to upload or download data. For that, see below.
func (r *AttachmentService) Update(ctx context.Context, id string, params AttachmentUpdateParams, opts ...option.RequestOption) (res *AttachmentSingle, err error) {
	if !param.IsOmitted(params.XTraceID) {
		opts = append(opts, option.WithHeader("X-Trace-Id", fmt.Sprintf("%v", params.XTraceID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/vnd.api+json")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/attachments/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return res, err
}

// This endpoint lists all attachments.
func (r *AttachmentService) List(ctx context.Context, params AttachmentListParams, opts ...option.RequestOption) (res *AttachmentArray, err error) {
	if !param.IsOmitted(params.XTraceID) {
		opts = append(opts, option.WithHeader("X-Trace-Id", fmt.Sprintf("%v", params.XTraceID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/vnd.api+json")}, opts...)
	path := "v1/attachments"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

// With this endpoint you delete an attachment, including any stored file data.
func (r *AttachmentService) Delete(ctx context.Context, id string, body AttachmentDeleteParams, opts ...option.RequestOption) (err error) {
	if !param.IsOmitted(body.XTraceID) {
		opts = append(opts, option.WithHeader("X-Trace-Id", fmt.Sprintf("%v", body.XTraceID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return err
	}
	path := fmt.Sprintf("v1/attachments/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// This endpoint allows you to download the binary content of a transaction. It
// will be sent to you as a download, using the content type
// "application/octet-stream" and content disposition "attachment;
// filename=example.pdf".
func (r *AttachmentService) Download(ctx context.Context, id string, query AttachmentDownloadParams, opts ...option.RequestOption) (res *http.Response, err error) {
	if !param.IsOmitted(query.XTraceID) {
		opts = append(opts, option.WithHeader("X-Trace-Id", fmt.Sprintf("%v", query.XTraceID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/octet-stream")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/attachments/%s/download", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Use this endpoint to upload (and possible overwrite) the file contents of an
// attachment. Simply put the entire file in the body as binary data.
func (r *AttachmentService) Upload(ctx context.Context, id string, body io.Reader, params AttachmentUploadParams, opts ...option.RequestOption) (err error) {
	if !param.IsOmitted(params.XTraceID) {
		opts = append(opts, option.WithHeader("X-Trace-Id", fmt.Sprintf("%v", params.XTraceID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*"), option.WithRequestBody("application/octet-stream", body)}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return err
	}
	path := fmt.Sprintf("v1/attachments/%s/upload", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, nil, opts...)
	return err
}

// The object class to which the attachment must be linked.
type AttachableType string

const (
	AttachableTypeAccount            AttachableType = "Account"
	AttachableTypeBudget             AttachableType = "Budget"
	AttachableTypeBill               AttachableType = "Bill"
	AttachableTypeTransactionJournal AttachableType = "TransactionJournal"
	AttachableTypePiggyBank          AttachableType = "PiggyBank"
	AttachableTypeTag                AttachableType = "Tag"
)

type AttachmentRead struct {
	ID         string                   `json:"id" api:"required"`
	Attributes AttachmentReadAttributes `json:"attributes" api:"required"`
	Links      ObjectLink               `json:"links" api:"required"`
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
func (r AttachmentRead) RawJSON() string { return r.JSON.raw }
func (r *AttachmentRead) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AttachmentReadAttributes struct {
	// ID of the model this attachment is linked to.
	AttachableID string `json:"attachable_id"`
	// The object class to which the attachment must be linked.
	//
	// Any of "Account", "Budget", "Bill", "TransactionJournal", "PiggyBank", "Tag".
	AttachableType AttachableType `json:"attachable_type"`
	CreatedAt      time.Time      `json:"created_at" format:"date-time"`
	DownloadURL    string         `json:"download_url"`
	Filename       string         `json:"filename"`
	// Hash of the file for basic duplicate detection.
	Hash      string    `json:"hash"`
	Mime      string    `json:"mime"`
	Notes     string    `json:"notes" api:"nullable"`
	Size      int64     `json:"size"`
	Title     string    `json:"title" api:"nullable"`
	UpdatedAt time.Time `json:"updated_at" format:"date-time"`
	UploadURL string    `json:"upload_url"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AttachableID   respjson.Field
		AttachableType respjson.Field
		CreatedAt      respjson.Field
		DownloadURL    respjson.Field
		Filename       respjson.Field
		Hash           respjson.Field
		Mime           respjson.Field
		Notes          respjson.Field
		Size           respjson.Field
		Title          respjson.Field
		UpdatedAt      respjson.Field
		UploadURL      respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AttachmentReadAttributes) RawJSON() string { return r.JSON.raw }
func (r *AttachmentReadAttributes) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AttachmentSingle struct {
	Data AttachmentRead `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AttachmentSingle) RawJSON() string { return r.JSON.raw }
func (r *AttachmentSingle) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ObjectLink struct {
	Number0 ObjectLink0 `json:"0"`
	Self    string      `json:"self" format:"uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Number0     respjson.Field
		Self        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ObjectLink) RawJSON() string { return r.JSON.raw }
func (r *ObjectLink) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ObjectLink0 struct {
	Rel string `json:"rel"`
	Uri string `json:"uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Rel         respjson.Field
		Uri         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ObjectLink0) RawJSON() string { return r.JSON.raw }
func (r *ObjectLink0) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AttachmentNewParams struct {
	// ID of the model this attachment is linked to.
	AttachableID string `json:"attachable_id" api:"required"`
	// The object class to which the attachment must be linked.
	//
	// Any of "Account", "Budget", "Bill", "TransactionJournal", "PiggyBank", "Tag".
	AttachableType AttachableType    `json:"attachable_type,omitzero" api:"required"`
	Filename       string            `json:"filename" api:"required"`
	Notes          param.Opt[string] `json:"notes,omitzero"`
	Title          param.Opt[string] `json:"title,omitzero"`
	XTraceID       param.Opt[string] `header:"X-Trace-Id,omitzero" format:"uuid" json:"-"`
	paramObj
}

func (r AttachmentNewParams) MarshalJSON() (data []byte, err error) {
	type shadow AttachmentNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AttachmentNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AttachmentGetParams struct {
	XTraceID param.Opt[string] `header:"X-Trace-Id,omitzero" format:"uuid" json:"-"`
	paramObj
}

type AttachmentUpdateParams struct {
	Notes    param.Opt[string] `json:"notes,omitzero"`
	Filename param.Opt[string] `json:"filename,omitzero"`
	Title    param.Opt[string] `json:"title,omitzero"`
	XTraceID param.Opt[string] `header:"X-Trace-Id,omitzero" format:"uuid" json:"-"`
	paramObj
}

func (r AttachmentUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow AttachmentUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AttachmentUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AttachmentListParams struct {
	// Number of items per page. The default pagination is per 50 items.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Page number. The default pagination is per 50 items.
	Page     param.Opt[int64]  `query:"page,omitzero" json:"-"`
	XTraceID param.Opt[string] `header:"X-Trace-Id,omitzero" format:"uuid" json:"-"`
	paramObj
}

// URLQuery serializes [AttachmentListParams]'s query parameters as `url.Values`.
func (r AttachmentListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AttachmentDeleteParams struct {
	XTraceID param.Opt[string] `header:"X-Trace-Id,omitzero" format:"uuid" json:"-"`
	paramObj
}

type AttachmentDownloadParams struct {
	XTraceID param.Opt[string] `header:"X-Trace-Id,omitzero" format:"uuid" json:"-"`
	paramObj
}

type AttachmentUploadParams struct {
	XTraceID param.Opt[string] `header:"X-Trace-Id,omitzero" format:"uuid" json:"-"`
	paramObj
}

func (r AttachmentUploadParams) MarshalMultipart() (data []byte, contentType string, err error) {
	buf := bytes.NewBuffer(nil)
	writer := multipart.NewWriter(buf)
	err = apiform.MarshalRoot(r, writer)
	if err == nil {
		err = apiform.WriteExtras(writer, r.ExtraFields())
	}
	if err != nil {
		writer.Close()
		return nil, "", err
	}
	err = writer.Close()
	if err != nil {
		return nil, "", err
	}
	return buf.Bytes(), writer.FormDataContentType(), nil
}
