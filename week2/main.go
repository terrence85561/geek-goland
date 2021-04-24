package main

import (
	"database/sql"
	"fmt"
	"github.com/pkg/errors"
)

// queryFromDao mock query in DAO
func queryFromDao() error {
	return sql.ErrNoRows
}

func mockService() error {
	err := queryFromDao()
	if errors.Is(err, sql.ErrNoRows) {
		return errors.Wrap(err, "some error info...")
	}
	return nil
}

func main() {
	err := mockService()
	fmt.Printf("%+v\n", err)
}
