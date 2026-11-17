# fyne_lucide

[![CI](https://github.com/timzifer/fyne_lucide/actions/workflows/ci.yml/badge.svg)](https://github.com/timzifer/fyne_lucide/actions/workflows/ci.yml)

[Lucide](https://lucide.dev/icons) (v1.74.0) for [Fyne](https://fyne.io) apps, embedded into your
binary and colorable at runtime. Requires Go 1.21+ and Fyne 2.3+.

## Install

```sh
go get github.com/timzifer/fyne_lucide
```

## Usage

```go
import (
	"image/color"

	"fyne.io/fyne/v2/widget"

	"github.com/timzifer/fyne_lucide"
)

// stained with the current theme foreground color
icon := widget.NewIcon(fyne_lucide.Icon(fyne_lucide.IconActivity))

// stained with a custom color (alpha is respected)
red := fyne_lucide.Icon(fyne_lucide.IconAlarmCheck, color.NRGBA{R: 0xd0, A: 0xff})

// re-color an existing icon
blue := red.(fyne_lucide.StainableResource).MustStain(color.NRGBA{B: 0xd0, A: 0xff})

// raw / stained SVG bytes, or a rasterized 64×64 PNG
svg := fyne_lucide.StainedSource(fyne_lucide.IconActivity, color.Black)
png, _ := fyne_lucide.PNG(fyne_lucide.IconActivity, 64, color.Black)

// names from configuration (as listed on https://lucide.dev/icons)
if i, ok := fyne_lucide.Lookup("arrow-up"); ok {
	icon.SetResource(fyne_lucide.Icon(i))
}
```

Icons are values of an unexported type that only the generated `Icon*`
variables and `Lookup` produce – arbitrary strings do not compile, so `Icon`
cannot fail. Since the type is unexported, store `fyne.Resource`s or names
(`i.String()`) rather than icon values in your own structs.

The default stain color is taken when the icon is requested; request the icon
again (or call `MustStain`) after a theme change.

`fyne_lucide.LucideVersion` holds the embedded Lucide release.

Built on [fyne_iconkit](https://github.com/timzifer/fyne_iconkit).

## Renamed icons

When Lucide renames an icon, the old `Icon*` variable stays available as a
deprecated alias of the new one (and `Lookup` accepts the old name). Like
Lucide's own packages, aliases are kept until the next major release of
Lucide (aliases from Lucide 1.x are removed with Lucide 2). See
[CHANGELOG.md](CHANGELOG.md) for renamed, removed, added and changed icons
per release.

## Versioning

Releases are tagged `v1.<minor>.<patch>`; the embedded Lucide release is
named in the release title, in [CHANGELOG.md](CHANGELOG.md) and in
`fyne_lucide.LucideVersion`. Every Lucide update merged into `main` is
released automatically as the next minor version, with its changelog section
as release notes ([release workflow](.github/workflows/release.yml)).

> [!IMPORTANT]
> Unlike strict semantic versioning, **minor releases may remove `Icon*`
> variables**: deprecated aliases once Lucide has a new major release, and
> icons that Lucide drops without a successor. Every removal is announced in
> CHANGELOG.md, and deprecated aliases are marked `// Deprecated:` (reported
> by `staticcheck` and gopls) for the whole Lucide major. Replace deprecated
> identifiers when you see them, or pin the version in your `go.mod` and read
> the changelog before upgrading.

## Updating the icons

A [daily workflow](.github/workflows/update-icons.yml) opens a pull request
for every new Lucide release, with the release notes as description.

Locally, `go generate ./...` does the same: it downloads the latest release
from npm (`lucide-static`), replaces the SVGs, license and the version above,
updates the aliases and `CHANGELOG.md`, and regenerates the `Icon*` variables.

```sh
go generate ./...                                    # latest release
go run ./internal/lucideupdate -version 1.74.0 && \
  go generate -run iconconst ./...                   # specific release
go run ./internal/lucideupdate -offline && \
  go generate -run iconconst ./...                   # regenerate without downloading
```

Renames are taken from the aliases `lucide-static` exports, from names
(`cart-x` → `shopping-cart-x`) and from identical shapes.
[`lucide.json`](lucide.json) records the imported releases and aliases;
there you can

- add an alias by hand (`"source": "manual"`, `"since"` = current version),
  then run with `-offline`,
- add a hint to `renames` (`"old": "new"`) that is used when a future release
  removes `old`.

## License

Code: [MIT](LICENSE). Icons: ISC, © Lucide Contributors (partly MIT, © Cole Bemis / Feather) – see [LICENSE-lucide](LICENSE-lucide).
