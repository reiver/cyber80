package verboten

import (
	"github.com/reiver/cyber80/os/export"

	"github.com/reiver/go-c80"

	"image"
	"image/color"
	"reflect"
)

const (
	pkg = "cyber80"
)

func init() {
	if nil == export.Symbols[pkg] {
                export.Symbols[pkg] = map[string]reflect.Value{}
	}

        export.Symbols[pkg]["ColorIndex"] = reflect.ValueOf(export_ColorIndex)
        export.Symbols[pkg]["Colorize"]   = reflect.ValueOf(export_Colorize)
        export.Symbols[pkg]["Draw"]       = reflect.ValueOf(export_Draw)
        export.Symbols[pkg]["Dye"]        = reflect.ValueOf(export_Dye)
        export.Symbols[pkg]["Map"]        = reflect.ValueOf(export_Map)
        export.Symbols[pkg]["Pixel"]      = reflect.ValueOf(export_Pixel)
        export.Symbols[pkg]["Relocate"]   = reflect.ValueOf(export_Relocate)
        export.Symbols[pkg]["SetSprite"]  = reflect.ValueOf(export_SetSprite)
        export.Symbols[pkg]["Sprite"]     = reflect.ValueOf(export_Sprite)
	export.Symbols[pkg]["Terminal"]   = reflect.ValueOf(c80.Terminal)
}

func export_ColorIndex(c color.Color) uint8 {
	return c80.ColorIndex(c)
}

func export_Colorize(a ...interface{}) error {
	return c80.Colorize(a...)
}

func export_Draw(img image.Image) error {
	return c80.Draw(img)
}

func export_Dye(index uint8) image.Image {
	return c80.Dye(index)
}

func export_Map() image.Image {
	return c80.Map()
}

func export_Pixel(x, y int, index uint8) image.Image {
	return c80.Pixel(x,y, index)
}

func export_Relocate(x, y int, img image.Image) image.Image {
	return c80.Relocate(x,y, img)
}

func export_SetSprite(kind string, id uint8, img image.Image) error {
	return c80.SetSprite(kind, id, img)
}

func export_Sprite(kind string, id uint8) image.Image {
	return c80.Sprite(kind, id)
}
