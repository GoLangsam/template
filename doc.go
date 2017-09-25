// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
Package template implements the interface common to the standard packages
"text/template" and "html/template" as type Template.

Thus, it exploits the fact, that

	  package html/template provides the same interface as package text/template
	  and should be used instead of text/template whenever the output is HTML
	  as it automatically secures HTML output against certain attacks.

as said in go doc text/template and go doc html/template.

All methods and all package level functions are forewarded.

Instead of a single New(name) this package unsurprisingly provides
two constructors:

- template.Text(name) &

- template.Html(name).

Note: The package level functions ParseFiles & ParseGlob became

- ParseTextFiles / ParseHtmlFile resp.

- ParseTextGlob / ParseHtmlGlob.

Thus, the exported type Template represents the template used,
be it html or text.

Also the type FuncMap is forwarded.

Note: Clients in need to access any other type

- such as ExecError (from "text/template") or

- data types such as HTML, CSS, JS and friends

- as well as Error and ErrorCode (from "html/template")

are requested to use the respective standard package directly for access to the error and data types.

For example escape_test.go uses

	  import( data "html/template" )

and refers to

	  data.HTML, data.CSS, data.JS ...

later.
*/
package template
