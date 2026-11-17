// Command lucideupdate imports a release of Lucide from the npm registry
// (package lucide-static) into this package, using fyne_iconkit/iconupdate.
//
// Renames are taken from the aliases lucide-static exports
// ("export { House as Home }"), so old names stay available as deprecated
// aliases until the next major release of Lucide.
//
// It runs as part of go generate, before iconconst regenerates the Icon*
// variables:
//
//	//go:generate go run ./internal/lucideupdate
package main

import "github.com/timzifer/fyne_iconkit/iconupdate"

func main() {
	iconupdate.Main(iconupdate.Source{
		Name:       "Lucide",
		URL:        "https://lucide.dev/icons",
		Package:    "lucide-static",
		Dirs:       map[string]string{"package/icons/": ""},
		License:    "package/LICENSE",
		Files:      []string{typings},
		CompareURL: compareURL,
		Aliases:    aliases,
		Clean:      clean,

		Command:      "go run ./internal/lucideupdate",
		GoPackage:    "fyne_lucide",
		VersionConst: "LucideVersion",
		StateFile:    "lucide.json",
		Out:          "lucide_gen.go",
		LicenseFile:  "LICENSE-lucide",
	})
}
