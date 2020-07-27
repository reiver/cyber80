package verboten

import (
	"github.com/reiver/cyber80/os/export"

	"math/rand"
	"reflect"
)

const (
	pkg = "math/rand"
)

func init() {
	if nil == export.Symbols[pkg] {
		export.Symbols[pkg] = map[string]reflect.Value{}
	}

	export.Symbols[pkg]["ExpFloat64"]  = reflect.ValueOf(rand.ExpFloat64)
	export.Symbols[pkg]["Float32"]     = reflect.ValueOf(rand.Float32)
	export.Symbols[pkg]["Float64"]     = reflect.ValueOf(rand.Float64)
	export.Symbols[pkg]["Int"]         = reflect.ValueOf(rand.Int)
	export.Symbols[pkg]["Int31"]       = reflect.ValueOf(rand.Int31)
	export.Symbols[pkg]["Int31n"]      = reflect.ValueOf(rand.Int31n)
	export.Symbols[pkg]["Int63"]       = reflect.ValueOf(rand.Int63)
	export.Symbols[pkg]["Int63n"]      = reflect.ValueOf(rand.Int63n)
	export.Symbols[pkg]["Intn"]        = reflect.ValueOf(rand.Intn)
	export.Symbols[pkg]["NormFloat64"] = reflect.ValueOf(rand.NormFloat64)
	export.Symbols[pkg]["Perm"]        = reflect.ValueOf(rand.Perm)
	export.Symbols[pkg]["Read"]        = reflect.ValueOf(rand.Read)
	export.Symbols[pkg]["Seed"]        = reflect.ValueOf(rand.Seed)
	export.Symbols[pkg]["Shuffle"]     = reflect.ValueOf(rand.Shuffle)
	export.Symbols[pkg]["Uint32"]      = reflect.ValueOf(rand.Uint32)
	export.Symbols[pkg]["Uint64"]      = reflect.ValueOf(rand.Uint64)
//	export.Symbols[pkg][""]   = reflect.ValueOf(rand.)
}
