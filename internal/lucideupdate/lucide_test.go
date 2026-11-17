package main

import "testing"

func TestParseAliases(t *testing.T) {
	dts := []byte(`/**
 * @name house
 */
declare const House: string;
/**
 * @name arrow-down-0-1
 */
declare const ArrowDown01: string;
export { House };
export { ArrowDown01, SquareActivity as ActivitySquare, House as Home, House };
`)
	icons := []string{"house", "home", "arrow-down-0-1", "arrow-down-01", "square-activity", "activity-square"}
	got, err := parseAliases(dts, icons)
	if err != nil {
		t.Fatal(err)
	}
	if len(got) != 3 || got["Home"] != "House" || got["ActivitySquare"] != "SquareActivity" || got["arrow-down-01"] != "arrow-down-0-1" {
		t.Errorf("aliases = %v", got)
	}
	if _, err := parseAliases([]byte("declare const House: string;"), icons); err == nil {
		t.Error("typings without export accepted")
	}
}

func TestClean(t *testing.T) {
	got := string(clean([]byte("<!-- @license lucide-static v1.74.0 - ISC -->\n<svg\n  class=\"lucide\"\n></svg>\n")))
	if got != "<svg\n  class=\"lucide\"\n></svg>\n" {
		t.Errorf("clean = %q", got)
	}
	if got := string(clean([]byte("<svg></svg>"))); got != "<svg></svg>" {
		t.Errorf("clean(plain) = %q", got)
	}
}

func TestCompareURL(t *testing.T) {
	if got := compareURL("0.93.0", "1.74.0"); got != "https://github.com/lucide-icons/lucide/compare/v0.93.0...1.74.0" {
		t.Errorf("compareURL = %s", got)
	}
}
