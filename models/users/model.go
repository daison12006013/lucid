package users

import (
	databaseconfig "github.com/daison12006013/gorvel/constants/databases.config"
	"github.com/daison12006013/gorvel/internals/database/sqlite"
)

func RawQuery(stmt string) []interface{} {
	var attr Attributes

	driver := sqlite.Make(databaseconfig.DB_DATABASE)
	records := driver.Query(stmt, &attr)

	return records
}

func First() Attributes {
	var record Attributes
	driver := sqlite.Make(databaseconfig.DB_DATABASE)
	driver.First(Table, PrimaryKey, &record)
	return record
}

func Last() Attributes {
	var record Attributes
	driver := sqlite.Make(databaseconfig.DB_DATABASE)
	driver.Last(Table, PrimaryKey, &record)
	return record
}
