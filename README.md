# template [![Software License](https:img.shields.io/:license-MIT-blue.svg?style=plastic)](LICENSE.md) [![Go Report Card](https:goreportcard.com/badge/github.com/GoLangsam/template?style=plastic)](https:goreportcard.com/report/github.com/GoLangsam/template) [![GoDoc](https:godoc.org/github.com/GoLangsam/template?status.svg&style=plastic)](https:godoc.org/github.com/GoLangsam/template) [![Build Status](https:travis-ci.org/GoLangsam/template.svg?branch=master&style=plastic)](https:travis-ci.org/GoLangsam/template)

Use standard "text/template" &amp; "html/template" together under a common umbrella

## package docu

Package `template` implements the interface common to the standard packages
`"text/template"` and `"html/template"` as type `Template`.

Thus, it exploits the fact, that

	  "package `html/template` provides the same interface as package `text/template`
	  and should be used instead of `text/template` whenever the output is HTML
	  as it automatically secures HTML output against certain attacks."

as said in `go doc text/template` and `go doc html/template`.

All methods and all package level functions are forewarded.

Instead of a single `New(name)` this package unsurprisingly provides
two constructors:
- `template.Text(name)` &
- `template.Html(name)`.

Note: The package level functions `ParseFiles` & `ParseGlob` became
- `ParseTextFiles` / `ParseHtmlFile` resp.
- `ParseTextGlob` / `ParseHtmlGlob`.

Thus, the exported type `Template` represents the template used,
be it `html` or `text`.

Also the type `FuncMap` is forwarded.

Note: Clients in need to access any other type
- such as `ExecError` (from `"text/template"`) or
- data types such as `HTML`, `CSS`, `JS` and friends
- as well as `Error` and `ErrorCode` (from `"html/template"`)

are requested to use the respective standard package directly for access to the error and data types.

For example `escape_test.go` uses

	  import( data "html/template" )

and refers to

	  data.HTML, data.CSS, data.JS ...

later.

## Implementation notes

`template.go` defines the interface (and the type `FuncMap` and the conveniencewrapper `Must`)
`foreward.go` provides package level function, and is simply taken from `"html/template"` as noted inside.

`wrapTextTemplate.go` and `wrapHtmlTemplate.go` define respective (internal) types and foreward the methods;
intentionally they differ as little as possible.

## Examples and test

(Almost) all `*_test.go` files from both standard packages are used.

Adjusted are just things such as:
- imports
  - `"text/template"` resp. `"html/template"` => `"github.com/GoLangsam/template"`
- references to type Template
  - `*Template` => `Template`
- constructors
  - `New` => `Text` resp. `Html`
- the global Parse-functions
  - `template.Parse...` => `template.ParseText...` resp. `template.ParseHtml...` 
- data types from `"html/template"`
  - import as `data` and use the types accordingly, e.g.
  - `template.HTML` => `data.HTML` 
- few portions are deactivated / commented, as they require internals of the underlying package 

As of now, no additional tests are provided. The author could not think of anything reasonable yet.

Your suggestions and contributions are very welcome ;-)

---
Think deep - code happy - be simple - see clear :-)