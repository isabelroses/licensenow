package main

import _ "embed"

//go:embed templates/mit.tmpls
var Mit string

//go:embed templates/cc-by-nc-sa-40.tmpls
var Ccbyncsa40 string

//go:embed templates/gplv3.tmpls
var Gplv3 string

//go:embed templates/eupl-12.tmpls
var Eupl12 string

var licenses = map[string]string{
	"MIT":             Mit,
	"GPLv3":           Gplv3,
	"cc by-nc-sa 4.0": Ccbyncsa40,
	"EUPL-1.2":        Eupl12,
}

var licenseNames = []string{
	"MIT",
	"GPLv3",
	"cc by-nc-sa 4.0",
	"EUPL-1.2",
}
