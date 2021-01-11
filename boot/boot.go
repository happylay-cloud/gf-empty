package boot

import (
	_ "github.com/gogf/gf-empty/packed"
	_ "github.com/mattn/go-sqlite3"

	"github.com/gogf/gf/os/gres"
)

func init() {
	gres.Dump()
}

