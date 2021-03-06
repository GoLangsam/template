// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package template

import (
	"html/template"
	"text/template/parse"
)

type htmlTmpl struct{ *template.Template }

// HTML returns a new "html/template" Template
func HTML(Name string) Template {
	return htmlTmpl{template.New(Name)}
}

// method wrappers - in alphabetical order

func (t htmlTmpl) AddParseTree(name string, tree *parse.Tree) (Template, error) {
	tmpl, err := t.Template.AddParseTree(name, tree)
	return htmlTmpl{tmpl}, err
}

func (t htmlTmpl) Clone() (Template, error) {
	tmpl, err := t.Template.Clone()
	return htmlTmpl{tmpl}, err
}

func (t htmlTmpl) Delims(left, right string) Template {
	return htmlTmpl{t.Template.Delims(left, right)}
}

/* inherited:
func (t htmlTmpl) Execute(wr io.Writer, data interface{}) error {
	return t.Template.Execute(wr, data)
}

func (t htmlTmpl) ExecuteTemplate(wr io.Writer, name string, data interface{}) error {
	return t.Template.ExecuteTemplate(wr, name, data)
}
*/

func (t htmlTmpl) Funcs(funcMap map[string]interface{}) Template {
	return htmlTmpl{t.Template.Funcs(template.FuncMap(funcMap))}
}

func (t htmlTmpl) Lookup(name string) Template {
	return htmlTmpl{t.Template.Lookup(name)}
}

/* inherited:
func (t htmlTmpl) Name() string {
	return t.Template.Name()
}
*/

func (t htmlTmpl) New(name string) Template {
	return htmlTmpl{t.Template.New(name)}
}

func (t htmlTmpl) Option(opt ...string) Template {
	return htmlTmpl{t.Template.Option(opt...)}
}

func (t htmlTmpl) Parse(text string) (Template, error) {
	tmpl, err := t.Template.Parse(text)
	return htmlTmpl{tmpl}, err
}

func (t htmlTmpl) ParseFiles(filenames ...string) (Template, error) {
	tmpl, err := t.Template.ParseFiles(filenames...)
	return htmlTmpl{tmpl}, err
}

func (t htmlTmpl) ParseGlob(pattern string) (Template, error) {
	tmpl, err := t.Template.ParseGlob(pattern)
	return htmlTmpl{tmpl}, err
}

func (t htmlTmpl) Templates() []Template {
	tmps := t.Template.Templates()
	news := make([]Template, 0, len(tmps))
	for i := range tmps {
		news = append(news, htmlTmpl{tmps[i]})
	}
	return news
}

// package functions

// ParseHTMLfiles wraps html/template.ParseFiles
func ParseHTMLfiles(filenames ...string) (Template, error) {
	tmpl, err := template.ParseFiles(filenames...)
	return htmlTmpl{tmpl}, err
}

// ParseHTMLglob wraps html/template.ParseGlob
func ParseHTMLglob(pattern string) (Template, error) {
	tmpl, err := template.ParseGlob(pattern)
	return htmlTmpl{tmpl}, err
}
