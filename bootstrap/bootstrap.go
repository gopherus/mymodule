package bootstrap

import (
	"github.com/gopherus/mymodule/checksum"
	"github.com/viant/xdatly/types/core"
	"reflect"
)

var PackageName = "bootstrap"

func init() {
	core.RegisterType(PackageName, "Bootstrap", reflect.TypeOf(struct{}{}), checksum.GeneratedTime)
}
