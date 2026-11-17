package main

import (
	"context"
	"fmt"
	"regexp"
	"strings"
)

// typings declares all icons of lucide-static, ending with one export
// statement that also lists the aliases.
const typings = "package/dist/lucide-static.d.ts"

var (
	exportBlock = regexp.MustCompile(`export \{([^}]*)\}`)
	aliasExport = regexp.MustCompile(`^(\w+) as (\w+)$`)
	declaration = regexp.MustCompile(`(?m)@name ([a-z0-9-]+)$`)
	// licenseComment precedes every SVG and names the release, which would
	// make every icon change with every release.
	licenseComment = regexp.MustCompile(`^<!--[^>]*-->\s*`)
)

// aliases returns the renames of lucide-static (old name → current name).
// Exported aliases ("export { House as Home }") are returned as Go
// identifiers without prefix, e.g. "Home" → "House"; the update tool matches
// them to icon names by identifier. lucide-static also ships an SVG file for
// every alias, but declares only current icons ("@name house"): undeclared
// files whose identifier equals a current icon's ("arrow-down-01" for
// "arrow-down-0-1") are returned by name.
func aliases(_ context.Context, _ string, files map[string][]byte, icons []string) (map[string]string, error) {
	return parseAliases(files[typings], icons)
}

func parseAliases(dts []byte, icons []string) (map[string]string, error) {
	blocks := exportBlock.FindAllSubmatch(dts, -1)
	if blocks == nil {
		return nil, fmt.Errorf("%s: no export statement", typings)
	}
	renames := map[string]string{}
	for _, export := range strings.Split(string(blocks[len(blocks)-1][1]), ",") {
		if m := aliasExport.FindStringSubmatch(strings.TrimSpace(export)); m != nil {
			renames[m[2]] = m[1]
		}
	}

	declared := map[string]string{} // identifier → name
	for _, m := range declaration.FindAllSubmatch(dts, -1) {
		declared[pascal(string(m[1]))] = string(m[1])
	}
	if len(declared) == 0 {
		return nil, fmt.Errorf("%s: no icon declarations", typings)
	}
	for _, name := range icons {
		id := pascal(name)
		if current, ok := declared[id]; ok && current != name {
			if _, exported := renames[id]; !exported {
				renames[name] = current
			}
		}
	}
	return renames, nil
}

// pascal converts an icon name like Lucide does for its exports
// ("arrow-down-0-1" → "ArrowDown01").
func pascal(name string) string {
	var b strings.Builder
	for _, s := range strings.Split(name, "-") {
		if s != "" {
			b.WriteString(strings.ToUpper(s[:1]) + s[1:])
		}
	}
	return b.String()
}

func clean(svg []byte) []byte {
	return licenseComment.ReplaceAll(svg, nil)
}

// compareURL links the Lucide releases; tags have a "v" prefix before 1.0.
func compareURL(from, to string) string {
	tag := func(version string) string {
		if strings.HasPrefix(version, "0.") {
			return "v" + version
		}
		return version
	}
	return "https://github.com/lucide-icons/lucide/compare/" + tag(from) + "..." + tag(to)
}
