package boot

import (
	_ "github.com/gogf/gf-empty/packed"
	_ "github.com/mattn/go-sqlite3"

	"github.com/gogf/gf-empty/library/utils"
	"github.com/gogf/gf/os/gres"
)

func init() {
	gres.Dump()
	utils.MemoryFileToLocal("./db", "sqlite3.db", "db/sqlite3.db")
}
