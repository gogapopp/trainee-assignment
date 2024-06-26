// Package handler provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen/v2 version v2.1.0 DO NOT EDIT.
package handler

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"net/http"
	"net/url"
	"path"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/go-chi/chi/v5"
	"github.com/oapi-codegen/runtime"
)

// GetBannerParams defines parameters for GetBanner.
type GetBannerParams struct {
	FeatureId *int `form:"feature_id,omitempty" json:"feature_id,omitempty"`
	TagId     *int `form:"tag_id,omitempty" json:"tag_id,omitempty"`
	Limit     *int `form:"limit,omitempty" json:"limit,omitempty"`
	Offset    *int `form:"offset,omitempty" json:"offset,omitempty"`

	// Token Токен админа
	Token *string `json:"token,omitempty"`
}

// PostBannerJSONBody defines parameters for PostBanner.
type PostBannerJSONBody struct {
	// Content Содержимое баннера
	Content *map[string]interface{} `json:"content,omitempty"`

	// FeatureId Идентификатор фичи
	FeatureId *int `json:"feature_id,omitempty"`

	// IsActive Флаг активности баннера
	IsActive *bool `json:"is_active,omitempty"`

	// TagIds Идентификаторы тэгов
	TagIds *[]int `json:"tag_ids,omitempty"`
}

// PostBannerParams defines parameters for PostBanner.
type PostBannerParams struct {
	// Token Токен админа
	Token *string `json:"token,omitempty"`
}

// DeleteBannerFeatureIdIdParams defines parameters for DeleteBannerFeatureIdId.
type DeleteBannerFeatureIdIdParams struct {
	// Token Токен админа
	Token *string `json:"token,omitempty"`
}

// DeleteBannerIdParams defines parameters for DeleteBannerId.
type DeleteBannerIdParams struct {
	// Token Токен админа
	Token *string `json:"token,omitempty"`
}

// PatchBannerIdJSONBody defines parameters for PatchBannerId.
type PatchBannerIdJSONBody struct {
	// Content Содержимое баннера
	Content *map[string]interface{} `json:"content"`

	// FeatureId Идентификатор фичи
	FeatureId *int `json:"feature_id"`

	// IsActive Флаг активности баннера
	IsActive *bool `json:"is_active"`

	// TagIds Идентификаторы тэгов
	TagIds *[]int `json:"tag_ids"`
}

// PatchBannerIdParams defines parameters for PatchBannerId.
type PatchBannerIdParams struct {
	// Token Токен админа
	Token *string `json:"token,omitempty"`
}

// GetUserBannerParams defines parameters for GetUserBanner.
type GetUserBannerParams struct {
	TagId           int   `form:"tag_id" json:"tag_id"`
	FeatureId       int   `form:"feature_id" json:"feature_id"`
	UseLastRevision *bool `form:"use_last_revision,omitempty" json:"use_last_revision,omitempty"`

	// Token Токен пользователя
	Token *string `json:"token,omitempty"`
}

// PostBannerJSONRequestBody defines body for PostBanner for application/json ContentType.
type PostBannerJSONRequestBody PostBannerJSONBody

// PatchBannerIdJSONRequestBody defines body for PatchBannerId for application/json ContentType.
type PatchBannerIdJSONRequestBody PatchBannerIdJSONBody

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Получение всех баннеров c фильтрацией по фиче и/или тегу
	// (GET /banner)
	GetBanner(w http.ResponseWriter, r *http.Request, params GetBannerParams)
	// Создание нового баннера
	// (POST /banner)
	PostBanner(w http.ResponseWriter, r *http.Request, params PostBannerParams)
	// Удаление баннера по идентификатору
	// (DELETE /banner/feature_id/{id})
	DeleteBannerFeatureIdId(w http.ResponseWriter, r *http.Request, id int, params DeleteBannerFeatureIdIdParams)
	// Удаление баннера по идентификатору
	// (DELETE /banner/{id})
	DeleteBannerId(w http.ResponseWriter, r *http.Request, id int, params DeleteBannerIdParams)
	// Обновление содержимого баннера
	// (PATCH /banner/{id})
	PatchBannerId(w http.ResponseWriter, r *http.Request, id int, params PatchBannerIdParams)
	// Получение баннера для пользователя
	// (GET /user_banner)
	GetUserBanner(w http.ResponseWriter, r *http.Request, params GetUserBannerParams)
}

// Unimplemented server implementation that returns http.StatusNotImplemented for each endpoint.

type Unimplemented struct{}

// Получение всех баннеров c фильтрацией по фиче и/или тегу
// (GET /banner)
func (_ Unimplemented) GetBanner(w http.ResponseWriter, r *http.Request, params GetBannerParams) {
	w.WriteHeader(http.StatusNotImplemented)
}

// Создание нового баннера
// (POST /banner)
func (_ Unimplemented) PostBanner(w http.ResponseWriter, r *http.Request, params PostBannerParams) {
	w.WriteHeader(http.StatusNotImplemented)
}

// Удаление баннера по идентификатору
// (DELETE /banner/feature_id/{id})
func (_ Unimplemented) DeleteBannerFeatureIdId(w http.ResponseWriter, r *http.Request, id int, params DeleteBannerFeatureIdIdParams) {
	w.WriteHeader(http.StatusNotImplemented)
}

// Удаление баннера по идентификатору
// (DELETE /banner/{id})
func (_ Unimplemented) DeleteBannerId(w http.ResponseWriter, r *http.Request, id int, params DeleteBannerIdParams) {
	w.WriteHeader(http.StatusNotImplemented)
}

// Обновление содержимого баннера
// (PATCH /banner/{id})
func (_ Unimplemented) PatchBannerId(w http.ResponseWriter, r *http.Request, id int, params PatchBannerIdParams) {
	w.WriteHeader(http.StatusNotImplemented)
}

// Получение баннера для пользователя
// (GET /user_banner)
func (_ Unimplemented) GetUserBanner(w http.ResponseWriter, r *http.Request, params GetUserBannerParams) {
	w.WriteHeader(http.StatusNotImplemented)
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler            ServerInterface
	HandlerMiddlewares []MiddlewareFunc
	ErrorHandlerFunc   func(w http.ResponseWriter, r *http.Request, err error)
}

type MiddlewareFunc func(http.Handler) http.Handler

// GetBanner operation middleware
func (siw *ServerInterfaceWrapper) GetBanner(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params GetBannerParams

	// ------------- Optional query parameter "feature_id" -------------

	err = runtime.BindQueryParameter("form", true, false, "feature_id", r.URL.Query(), &params.FeatureId)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "feature_id", Err: err})
		return
	}

	// ------------- Optional query parameter "tag_id" -------------

	err = runtime.BindQueryParameter("form", true, false, "tag_id", r.URL.Query(), &params.TagId)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "tag_id", Err: err})
		return
	}

	// ------------- Optional query parameter "limit" -------------

	err = runtime.BindQueryParameter("form", true, false, "limit", r.URL.Query(), &params.Limit)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "limit", Err: err})
		return
	}

	// ------------- Optional query parameter "offset" -------------

	err = runtime.BindQueryParameter("form", true, false, "offset", r.URL.Query(), &params.Offset)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "offset", Err: err})
		return
	}

	headers := r.Header

	// ------------- Optional header parameter "token" -------------
	if valueList, found := headers[http.CanonicalHeaderKey("token")]; found {
		var Token string
		n := len(valueList)
		if n != 1 {
			siw.ErrorHandlerFunc(w, r, &TooManyValuesForParamError{ParamName: "token", Count: n})
			return
		}

		err = runtime.BindStyledParameterWithOptions("simple", "token", valueList[0], &Token, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationHeader, Explode: false, Required: false})
		if err != nil {
			siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "token", Err: err})
			return
		}

		params.Token = &Token

	}

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetBanner(w, r, params)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// PostBanner operation middleware
func (siw *ServerInterfaceWrapper) PostBanner(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params PostBannerParams

	headers := r.Header

	// ------------- Optional header parameter "token" -------------
	if valueList, found := headers[http.CanonicalHeaderKey("token")]; found {
		var Token string
		n := len(valueList)
		if n != 1 {
			siw.ErrorHandlerFunc(w, r, &TooManyValuesForParamError{ParamName: "token", Count: n})
			return
		}

		err = runtime.BindStyledParameterWithOptions("simple", "token", valueList[0], &Token, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationHeader, Explode: false, Required: false})
		if err != nil {
			siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "token", Err: err})
			return
		}

		params.Token = &Token

	}

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.PostBanner(w, r, params)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// DeleteBannerFeatureIdId operation middleware
func (siw *ServerInterfaceWrapper) DeleteBannerFeatureIdId(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "id" -------------
	var id int

	err = runtime.BindStyledParameterWithOptions("simple", "id", chi.URLParam(r, "id"), &id, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "id", Err: err})
		return
	}

	// Parameter object where we will unmarshal all parameters from the context
	var params DeleteBannerFeatureIdIdParams

	headers := r.Header

	// ------------- Optional header parameter "token" -------------
	if valueList, found := headers[http.CanonicalHeaderKey("token")]; found {
		var Token string
		n := len(valueList)
		if n != 1 {
			siw.ErrorHandlerFunc(w, r, &TooManyValuesForParamError{ParamName: "token", Count: n})
			return
		}

		err = runtime.BindStyledParameterWithOptions("simple", "token", valueList[0], &Token, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationHeader, Explode: false, Required: false})
		if err != nil {
			siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "token", Err: err})
			return
		}

		params.Token = &Token

	}

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.DeleteBannerFeatureIdId(w, r, id, params)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// DeleteBannerId operation middleware
func (siw *ServerInterfaceWrapper) DeleteBannerId(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "id" -------------
	var id int

	err = runtime.BindStyledParameterWithOptions("simple", "id", chi.URLParam(r, "id"), &id, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "id", Err: err})
		return
	}

	// Parameter object where we will unmarshal all parameters from the context
	var params DeleteBannerIdParams

	headers := r.Header

	// ------------- Optional header parameter "token" -------------
	if valueList, found := headers[http.CanonicalHeaderKey("token")]; found {
		var Token string
		n := len(valueList)
		if n != 1 {
			siw.ErrorHandlerFunc(w, r, &TooManyValuesForParamError{ParamName: "token", Count: n})
			return
		}

		err = runtime.BindStyledParameterWithOptions("simple", "token", valueList[0], &Token, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationHeader, Explode: false, Required: false})
		if err != nil {
			siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "token", Err: err})
			return
		}

		params.Token = &Token

	}

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.DeleteBannerId(w, r, id, params)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// PatchBannerId operation middleware
func (siw *ServerInterfaceWrapper) PatchBannerId(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "id" -------------
	var id int

	err = runtime.BindStyledParameterWithOptions("simple", "id", chi.URLParam(r, "id"), &id, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "id", Err: err})
		return
	}

	// Parameter object where we will unmarshal all parameters from the context
	var params PatchBannerIdParams

	headers := r.Header

	// ------------- Optional header parameter "token" -------------
	if valueList, found := headers[http.CanonicalHeaderKey("token")]; found {
		var Token string
		n := len(valueList)
		if n != 1 {
			siw.ErrorHandlerFunc(w, r, &TooManyValuesForParamError{ParamName: "token", Count: n})
			return
		}

		err = runtime.BindStyledParameterWithOptions("simple", "token", valueList[0], &Token, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationHeader, Explode: false, Required: false})
		if err != nil {
			siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "token", Err: err})
			return
		}

		params.Token = &Token

	}

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.PatchBannerId(w, r, id, params)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// GetUserBanner operation middleware
func (siw *ServerInterfaceWrapper) GetUserBanner(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params GetUserBannerParams

	// ------------- Required query parameter "tag_id" -------------

	if paramValue := r.URL.Query().Get("tag_id"); paramValue != "" {

	} else {
		siw.ErrorHandlerFunc(w, r, &RequiredParamError{ParamName: "tag_id"})
		return
	}

	err = runtime.BindQueryParameter("form", true, true, "tag_id", r.URL.Query(), &params.TagId)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "tag_id", Err: err})
		return
	}

	// ------------- Required query parameter "feature_id" -------------

	if paramValue := r.URL.Query().Get("feature_id"); paramValue != "" {

	} else {
		siw.ErrorHandlerFunc(w, r, &RequiredParamError{ParamName: "feature_id"})
		return
	}

	err = runtime.BindQueryParameter("form", true, true, "feature_id", r.URL.Query(), &params.FeatureId)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "feature_id", Err: err})
		return
	}

	// ------------- Optional query parameter "use_last_revision" -------------

	err = runtime.BindQueryParameter("form", true, false, "use_last_revision", r.URL.Query(), &params.UseLastRevision)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "use_last_revision", Err: err})
		return
	}

	headers := r.Header

	// ------------- Optional header parameter "token" -------------
	if valueList, found := headers[http.CanonicalHeaderKey("token")]; found {
		var Token string
		n := len(valueList)
		if n != 1 {
			siw.ErrorHandlerFunc(w, r, &TooManyValuesForParamError{ParamName: "token", Count: n})
			return
		}

		err = runtime.BindStyledParameterWithOptions("simple", "token", valueList[0], &Token, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationHeader, Explode: false, Required: false})
		if err != nil {
			siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "token", Err: err})
			return
		}

		params.Token = &Token

	}

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetUserBanner(w, r, params)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

type UnescapedCookieParamError struct {
	ParamName string
	Err       error
}

func (e *UnescapedCookieParamError) Error() string {
	return fmt.Sprintf("error unescaping cookie parameter '%s'", e.ParamName)
}

func (e *UnescapedCookieParamError) Unwrap() error {
	return e.Err
}

type UnmarshalingParamError struct {
	ParamName string
	Err       error
}

func (e *UnmarshalingParamError) Error() string {
	return fmt.Sprintf("Error unmarshaling parameter %s as JSON: %s", e.ParamName, e.Err.Error())
}

func (e *UnmarshalingParamError) Unwrap() error {
	return e.Err
}

type RequiredParamError struct {
	ParamName string
}

func (e *RequiredParamError) Error() string {
	return fmt.Sprintf("Query argument %s is required, but not found", e.ParamName)
}

type RequiredHeaderError struct {
	ParamName string
	Err       error
}

func (e *RequiredHeaderError) Error() string {
	return fmt.Sprintf("Header parameter %s is required, but not found", e.ParamName)
}

func (e *RequiredHeaderError) Unwrap() error {
	return e.Err
}

type InvalidParamFormatError struct {
	ParamName string
	Err       error
}

func (e *InvalidParamFormatError) Error() string {
	return fmt.Sprintf("Invalid format for parameter %s: %s", e.ParamName, e.Err.Error())
}

func (e *InvalidParamFormatError) Unwrap() error {
	return e.Err
}

type TooManyValuesForParamError struct {
	ParamName string
	Count     int
}

func (e *TooManyValuesForParamError) Error() string {
	return fmt.Sprintf("Expected one value for %s, got %d", e.ParamName, e.Count)
}

// Handler creates http.Handler with routing matching OpenAPI spec.
func Handler(si ServerInterface) http.Handler {
	return HandlerWithOptions(si, ChiServerOptions{})
}

type ChiServerOptions struct {
	BaseURL          string
	BaseRouter       chi.Router
	Middlewares      []MiddlewareFunc
	ErrorHandlerFunc func(w http.ResponseWriter, r *http.Request, err error)
}

// HandlerFromMux creates http.Handler with routing matching OpenAPI spec based on the provided mux.
func HandlerFromMux(si ServerInterface, r chi.Router) http.Handler {
	return HandlerWithOptions(si, ChiServerOptions{
		BaseRouter: r,
	})
}

func HandlerFromMuxWithBaseURL(si ServerInterface, r chi.Router, baseURL string) http.Handler {
	return HandlerWithOptions(si, ChiServerOptions{
		BaseURL:    baseURL,
		BaseRouter: r,
	})
}

// HandlerWithOptions creates http.Handler with additional options
func HandlerWithOptions(si ServerInterface, options ChiServerOptions) http.Handler {
	r := options.BaseRouter

	if r == nil {
		r = chi.NewRouter()
	}
	if options.ErrorHandlerFunc == nil {
		options.ErrorHandlerFunc = func(w http.ResponseWriter, r *http.Request, err error) {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	}
	wrapper := ServerInterfaceWrapper{
		Handler:            si,
		HandlerMiddlewares: options.Middlewares,
		ErrorHandlerFunc:   options.ErrorHandlerFunc,
	}

	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/banner", wrapper.GetBanner)
	})
	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/banner", wrapper.PostBanner)
	})
	r.Group(func(r chi.Router) {
		r.Delete(options.BaseURL+"/banner/feature_id/{id}", wrapper.DeleteBannerFeatureIdId)
	})
	r.Group(func(r chi.Router) {
		r.Delete(options.BaseURL+"/banner/{id}", wrapper.DeleteBannerId)
	})
	r.Group(func(r chi.Router) {
		r.Patch(options.BaseURL+"/banner/{id}", wrapper.PatchBannerId)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/user_banner", wrapper.GetUserBanner)
	})

	return r
}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/+xZ3U4bRxR+ldVcb4JJ6I0v06oVrdREqnpVKmvwjmHS/cvMGAVZlsCmTaTQoFa9qFQ1",
	"FWkfYDE4cSBeXuHMG1VnxsRevAs2BJIgX9k7u3POmfPzfXNmGoSHtYiUG0Rx5TNSJrALXb0BHejpTQf2",
	"IIE+9M1QCh3ikjUmJI9CUibzt0u3S6TpkihmIY05KZO7ZsglMVWrEqXOLdMwZAL/rjCFP1HMBFU8Chc9",
	"UiZfMXXPfoGTBA2YYkKS8g8N4jFZFTxWVhm8hBQOoQt9BxI4gLfQgz4kxCUcX68y6hkhIQ1wFSr6iYXE",
	"JbK6ygKKetljGsRmhdQLeFg5+UKtxzgoleDhCmk23YaV+KjOxPpQYI1RVReswr2M1FNG/gkHaKJuQU9v",
	"QQ8OIdEtSPWGg4/6CfSGGnmo2AoTxSoVXbm4uhZ0Yd/4Z1J1Pg+4OkvbX9BDt+vWFEKjWk2yM6W+0Ft6",
	"S29CN1/ujy4RTMZRKJlJqDulEv5Uo1Cx0CQUjWOfV01KzT2UKLQxoo0rFpiJscDMU9yKsWmJ7p3Gq5ly",
	"yPOtmzHM8zjKpP6DEd1K1Jl7WuUupKhUb8Ar4+QUuuPKhhncWLLlukTKzhKRUcAqg2fXWSKKPVajb8wj",
	"vqgLf2TcPDWHi4iWH7KqMmsQjCrmVajK8c4fxhuJozchhddwYIzs6Z1xe2uRCFAE8ahitxQP2Hi5uaOF",
	"9Z7qySVcVmhV8TWWI/I/OIIE9hFFDlEudKAPqd7E/4UBXo4in9EQZduqlNMYq585uqV/hf0BhL7LyXHL",
	"ByNUCLqOz/XYOz8UkMIeLgI6cIQWXDwczZxsyJqEn2TNuP8NfrVQms+x8B9I4Uhvw2tjXGJQ6UhvO2gW",
	"RqBjHQS9ky+gb4XdnVoYlk0XUcSBAxtP3YZjSFDeZ1OCRhYrmBCRGAlXsbPGnQO/Q1+3dUtvmMD09Q6G",
	"JtVPoQd7mCGOgT7kWxsnFCHrQUDF+nDNbf3ExhUX2jEzfh7jZqdqigJdYvQl+hecAG8cOIb0pGDQU3P4",
	"FfQGJKHb6KI4kjn0/CCSHxs/G0J4VGdS3Yu89UuE9SYg9Qw8i8Ezvzwxd7hgng1vc2x3MX+JjLrYpiJD",
	"pIji+5Cev9OYBHw+tzxuEfWDIeDf0IVDE0nEQEybvn6GVTNYMT7MGOTKGGR3dJuGy7Qbhdw0w6mDjm1u",
	"CC1zDe41rTd9ptg4S3xhxi1PfGmnLXqL3jhlGCrA3nBIBKbByRblhRqeCQrmw1JWBmXu5Kztt+ECHN3W",
	"m3AMXf0Uw+Xotgmg2dzNqvny1bxQWjgnAHAAR3rnhHmSgcw+JPDGpuCnDQr/DvPJgkKmeuyGEXqFlNzO",
	"IMU08DBDhTNQYWGGCjNU+KRRwZz8VldzGkkcnkHArJd918uGdd+nyyjY2nYFvW2BiivpdQt0XVvvW6D/",
	"Ur1wadzWk8O+GdNcMdPcLGZ5cfqAGrr23CMLP0UtaV0yUTn/JvF7yUTRaeWZF2wT08xLLEZDhDmR35ni",
	"XixznfgeWO4C14t1ySo+laoi2Bo3V7pZzTVa9xUp16gvx6ljeC6d6BbmvAFJ3cb9g942yfIca6Cvt0xB",
	"vbXH0fp5zrHg2QRc7Oupydik0eTb8emKbXLe/fq7+9/eghQjB3uY5PCqYMt1fUfJOSWfgaOiKMzI4Pra",
	"jhvGCTlXW6f6jcGyC3Ov2Wz+HwAA///X4owtwCIAAA==",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %w", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}

	return buf.Bytes(), nil
}

var rawSpec = decodeSpecCached()

// a naive cached of a decoded swagger spec
func decodeSpecCached() func() ([]byte, error) {
	data, err := decodeSpec()
	return func() ([]byte, error) {
		return data, err
	}
}

// Constructs a synthetic filesystem for resolving external references when loading openapi specifications.
func PathToRawSpec(pathToFile string) map[string]func() ([]byte, error) {
	res := make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	resolvePath := PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		pathToFile := url.String()
		pathToFile = path.Clean(pathToFile)
		getSpec, ok := resolvePath[pathToFile]
		if !ok {
			err1 := fmt.Errorf("path not found: %s", pathToFile)
			return nil, err1
		}
		return getSpec()
	}
	var specData []byte
	specData, err = rawSpec()
	if err != nil {
		return
	}
	swagger, err = loader.LoadFromData(specData)
	if err != nil {
		return
	}
	return
}
