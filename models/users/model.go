package users

import (
	"math"

	databaseconfig "github.com/daison12006013/gorvel/constants/databases.config"
	adapter "github.com/daison12006013/gorvel/internals/dbadapter"
	"github.com/daison12006013/gorvel/internals/query"
)

func Lists(currentPage int, perPage int, orderByCol string, orderBySort string) (*Paginate, error) {
	var err error

	// connect to our database
	db := query.Connect(adapter.SQLite(databaseconfig.DB_DATABASE))

	selectStmt := query.Interpreter().
		Table(Table).
		OrderBy(orderByCol, orderBySort).
		Limit(perPage).
		Offset((currentPage - 1) * perPage).
		ToSql()

	countStmt := query.Interpreter().
		Table(Table).
		CountSql()

	// query the total count
	var total int
	err = db.Select(countStmt).Find(&total)
	if err != nil {
		return nil, err
	}

	// query the data
	var data []Attributes
	err = db.Select(selectStmt).Find(&data)
	if err != nil {
		return nil, err
	}

	var paginated Paginate
	paginated.PerPage = perPage
	paginated.CurrentPage = currentPage
	paginated.Data = data
	paginated.Total = total
	paginated.LastPage = int(math.Ceil(float64(total) / float64(perPage)))

	return &paginated, nil
}
