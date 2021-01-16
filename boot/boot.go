package boot

import (
	_ "github.com/gogf/gf-empty/packed"
	_ "github.com/mattn/go-sqlite3"

	"github.com/gogf/gf/os/gres"
	"github.com/happylay-cloud/gf-extend/web/gfboot"
)

func init() {
	gres.Dump()
	gfboot.SingleFileMemoryToLocal("./db", "sqlite3.db", "db/sqlite3.db")
}
