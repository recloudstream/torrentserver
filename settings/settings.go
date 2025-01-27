package settings

import (
	"path/filepath"

	"server/log"
)

var (
	tdb      *TDB
	Path     string
	Port     string
	ReadOnly bool
	HttpAuth bool
	SearchWA bool
	PubIPv4  string
	PubIPv6  string
	TorAddr  string
)

func InitSets(readOnly, searchWA bool) bool {
	ReadOnly = readOnly
	SearchWA = searchWA
	tdb = NewTDB()
	if tdb == nil {
		log.TLogln("Error open db:", filepath.Join(Path, "config.db"))
		return false
	}
	loadBTSets()
	Migrate()
	return true
}

func CloseDB() {
	tdb.CloseDB()
}
