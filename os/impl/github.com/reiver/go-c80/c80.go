package verboten

import (
	"github.com/reiver/cyber80/os/export"

	"github.com/reiver/go-c80"

	"reflect"
)

const (
	pkg = "github.com/reiver/go-c80"
)

func init() {
	if nil == export.Symbols[pkg] {
                export.Symbols[pkg] = map[string]reflect.Value{}
	}

        export.Symbols[pkg]["ColorIndex"] = reflect.ValueOf(c80.ColorIndex)
        export.Symbols[pkg]["Colorize"]   = reflect.ValueOf(c80.Colorize)
        export.Symbols[pkg]["Draw"]       = reflect.ValueOf(c80.Draw)
        export.Symbols[pkg]["Dye"]        = reflect.ValueOf(c80.Dye)
        export.Symbols[pkg]["Map"]        = reflect.ValueOf(c80.Map)
        export.Symbols[pkg]["Pixel"]      = reflect.ValueOf(c80.Pixel)
        export.Symbols[pkg]["Relocate"]   = reflect.ValueOf(c80.Relocate)
        export.Symbols[pkg]["Serialize"]  = reflect.ValueOf(c80.Serialize)
        export.Symbols[pkg]["SetSprite"]  = reflect.ValueOf(c80.SetSprite)
        export.Symbols[pkg]["Sprite"]     = reflect.ValueOf(c80.Sprite)
	export.Symbols[pkg]["Terminal"]   = reflect.ValueOf(c80.Terminal)
}
