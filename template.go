// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package template implements the interface common to the standard packages
// `"text/template"` and `"html/template"` as type `Template`.
//
// Thus, it exploits the fact, that
//   "package `html/template` provides the same interface as package `text/template`
//   and should be used instead of `text/template` whenever the output is HTML
//   as it automatically secures HTML output against certain attacks."
// as said in `go doc text/template` and `go doc html/template`.
//
// All methods and all package level functions are forewarded.
//
// Instead of a single `New(name)` this package unsurprisingly provides
// two constructors: `template.Text(name)` & `template.Html(name)`.
//
// Note: The package level functions `ParseFiles` & `ParseGlob` became
// `ParseTextFiles` / `ParseHtmlFile` resp. `ParseTextGlob` / `ParseHtmlGlob`.
//
// Thus, the exported type `Template` represents the template used,
// be it `html` or `text`.
//
// Also the type `FuncMap` is forwarded.
//
// Note: Clients in need to access any other type such as `ExecError` (from `"text/template"`)
// or data types such as `HTML`, `CSS`, `JS` and friends as well as `Error` and `ErrorCode` (from `"html/template"`)
// are requested to use the respective standard package directly for access to the error and data types.
//
// Hint: `escape_test.go` for example uses
//   import( data "html/template" )
// and refers to
//   data.HTML, data.CSS, data.JS ...
// later.
package template

import (
	"io"
	"text/template/parse"
)

// Template represents the template used (html or text)
//
// For documentation of these methods please refer to the
// respective underlying standard package,
// e.g. `go doc text/template Parse`
// or `go doc html/template ExecuteTemplate`.
type Template interface {
	AddParseTree(name string, tree *parse.Tree) (Template, error)
	Clone() (Template, error)
	DefinedTemplates() string
	Delims(left, right string) Template
	Execute(wr io.Writer, data interface{}) error
	ExecuteTemplate(wr io.Writer, name string, data interface{}) error
	Funcs(funcMap map[string]interface{}) Template
	Lookup(name string) Template
	Name() string
	New(name string) Template
	Option(opt ...string) Template
	Parse(text string) (Template, error)
	ParseFiles(filenames ...string) (Template, error)
	ParseGlob(pattern string) (Template, error)
	Templates() []Template
}

// FuncMap is the type of the map defining the mapping from names to functions.
// Each function must have either a single return value, or two return values of
// which the second has type error. In that case, if the second (error)
// return value evaluates to non-nil during execution, execution terminates and
// Execute returns that error.
//
// When template execution invokes a function with an argument list, that list
// must be assignable to the function's parameter types. Functions meant to
// apply to arguments of arbitrary type can use parameters of type interface{} or
// of type reflect.Value. Similarly, functions meant to return a result of arbitrary
// type can return interface{} or reflect.Value.
type FuncMap map[string]interface{}

// Must is a helper that wraps a call to a function returning (Template, error)
// and panics if the error is non-nil. It is intended for use in variable
// initializations such as
//	var t = template.Must(template.NewText|HtmlTemplate("name").Parse("text"))
func Must(t Template, err error) Template {
	if err != nil {
		panic(err)
	}
	return t
}
