# fyne_lucide

[![CI](https://github.com/timzifer/fyne_lucide/actions/workflows/ci.yml/badge.svg)](https://github.com/timzifer/fyne_lucide/actions/workflows/ci.yml)

[Lucide](https://lucide.dev/icons) for [Fyne](https://fyne.io) apps, embedded into your
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
icon := widget.NewIcon(fyne_lucide.MustIcon(fyne_lucide.IconActivity))

// stained with a custom color (alpha is respected)
red := fyne_lucide.MustIcon("alarm-check", color.NRGBA{R: 0xd0, A: 0xff})

// re-color an existing icon
blue := red.(fyne_lucide.StainableResource).MustStain(color.NRGBA{B: 0xd0, A: 0xff})

// raw / stained SVG bytes, or a rasterized 64×64 PNG
svg, _ := fyne_lucide.StainedSource("activity", color.Black)
png, _ := fyne_lucide.PNG("activity", 64, color.Black)
```

Icon names match the file names on https://lucide.dev/icons (without `.svg`); the generated
`Icon*` constants list all of them. `Icon` returns an error wrapping
`fs.ErrNotExist` for unknown names; `MustIcon` logs via `fyne.LogError` and
returns `theme.ErrorIcon()` instead.

The default stain color is taken when the icon is requested; request the icon
again (or call `MustStain`) after a theme change.

Built on [fyne_iconkit](https://github.com/timzifer/fyne_iconkit). After
updating the SVGs, run `go generate ./...` to refresh the constants.

## License

Code: [MIT](LICENSE). Icons: ISC, © Lucide Contributors (partly MIT, © Cole Bemis / Feather) – see [LICENSE-lucide](LICENSE-lucide).
