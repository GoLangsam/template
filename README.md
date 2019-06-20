# template
Use both [`"text/template"`](https://godoc.org/github.com/golang/go/src/text/template) &amp; [`"html/template"`](https://godoc.org/github.com/golang/go/src/html/template) together under [a common umbrella](https://github.com/GoLangsam/template).

[![Software License](https:img.shields.io/:license-MIT-blue.svg?style=plastic)](LICENSE.md) [![Go Report Card](https:goreportcard.com/badge/github.com/GoLangsam/template?style=plastic)](https:goreportcard.com/report/github.com/GoLangsam/template) [![GoDoc](https:godoc.org/github.com/GoLangsam/template?status.svg&style=plastic)](https:godoc.org/github.com/GoLangsam/template) [![Build Status](https:travis-ci.org/GoLangsam/template.svg?branch=master&style=plastic)](https:travis-ci.org/GoLangsam/template)

## background
Shall be written another day :-)

## package docu

Package [`template`](https://github.com/GoLangsam/template) implements the interface common to the standard packages
[`"text/template"`](https://github.com/golang/go/tree/master/src/text/template) and [`"html/template"`](https://github.com/golang/go/tree/master/src/html/template) as type [`Template`](https://godoc.org/github.com/GoLangsam/template#Template).

Thus, it exploits the fact, that

	  package `html/template` provides the same interface as package `text/template`
	  and should be used instead of `text/template` whenever the output is HTML
	  as it automatically secures HTML output against certain attacks.

as said in `go doc text/template` and `go doc html/template`.

All methods and all package level functions are forewarded.

Instead of a single `New(name)` this package unsurprisingly provides
***two*** constructors:

- `template.Text(name)` &

- `template.HTML(name)`.

Note: The package level functions `ParseFiles` & `ParseGlob` became

- `ParseTextFiles` / `ParseHTMLfile` resp.

- `ParseTextGlob` / `ParseHTMLglob`.

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

- `doc.go` just documents the package (as quoted above)
  for [`go doc github.com/GoLangsam/template`](https:godoc.org/github.com/GoLangsam/template)
- `template.go` defines the interface type `Template`
  (and the convenient wrapper `Must` and the ubiquous type `FuncMap`)
- `forward.go` forwards package level functions.
  It'is simply taken from `"html/template"` as noted [inside](https://github.com/GoLangsam/template/blob/master/forward.go).
- `wrapTemplateText.go` and `wrapTemplateHTML.go` define respective (private) implementation types,
  foreward the methods (if need) and the `Parse...` functions.
  Intentionally they are as similar possible.

## Examples and test

***All*** `*_test.go` files from ***both*** standard packages are used! _(Except as noted below.)_

Adjusted are just things such as:
- package declarations
  - `template` => `template_test`
  - in order to avoid spurious bug: duplicate flag definition `debug` 
- imports
  - `"text/template"` resp. `"html/template"` => `"github.com/GoLangsam/template"`
- references to type `Template`
  - `*Template` => `Template`
- constructors
  - `New` => `Text` resp. `HTML`
- the global `Parse`-functions
  - `template.Parse...` => `template.ParseText...` resp. `template.ParseHTML...` 
- data types from `"html/template"`
  - imported as `data`, and types used accordingly, e.g.
  - `template.HTML` => `data.HTML` 
- few portions are deactivated / commented, as they require internals of the underlying package
  - `css_test.go`, `html_test.go`, `js_test.go` `url_test.go` and from `"html/template"` are entirely omitted for same reason
- misspellings: *Cincinatti* => *Cincinnati*
  - in order to make [Go Report Card](https:goreportcard.com/report/github.com/GoLangsam/template) more happy :-)


As of now, no additional tests are provided. The author could not think of anything reasonable yet.

Your suggestions, remarks, questions and/or contributions are welcome ;-)

---
## Think deep - code happy - be simple - see clear :-)

---
## Support on Beerpay
Hey dude! Help me out for a couple of :beers:!

[![Beerpay](https://beerpay.io/GoLangsam/template/badge.svg?style=beer-square)](https://beerpay.io/GoLangsam/template)  [![Beerpay](https://beerpay.io/GoLangsam/template/make-wish.svg?style=flat-square)](https://beerpay.io/GoLangsam/template?focus=wish)