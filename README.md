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

Built on [fyne_iconkit](https://github.com/timzifer/fyne_iconkit). After
updating the SVGs, run `go generate ./...` to refresh the `Icon*` variables.

## License

Code: [MIT](LICENSE). Icons: ISC, © Lucide Contributors (partly MIT, © Cole Bemis / Feather) – see [LICENSE-lucide](LICENSE-lucide).
