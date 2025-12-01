package components

import (
	"fmt"

	"github.com/benbjohnson/hashfs"
)

// StaticSys will be set by the web package
var StaticSys *hashfs.FS

func StaticPath(format string, args ...any) string {
	return "/" + StaticSys.HashName(fmt.Sprintf("static/"+format, args...))
}
