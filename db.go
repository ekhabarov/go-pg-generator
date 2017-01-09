package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Column struct {
	Name     string
	Type     string
	Default  *string
	IsNull   bool
	IsArray  bool
	Category string
}

func dbConnect(opts Options) (db *sql.DB, err error) {
	db, err = sql.Open("postgres",
		fmt.Sprintf(
			"host=%s port=%d user=%s dbname=%s sslmode=%s\n",
			opts.Host,
			opts.Port,
			opts.User,
			opts.Database,
			opts.SSLMode,
		),
	)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func columnList(db *sql.DB, tab string) (cols []*Column, err error) {
	rows, err := db.Query(COLUMNS_QUERY, tab)
	if err != nil {
		return nil, fmt.Errorf("unable to query data: %s", err)
	}

	for rows.Next() {
		c := &Column{}
		err := rows.Scan(&c.Name, &c.Type, &c.Default, &c.IsNull, &c.IsArray, &c.Category)
		if err != nil {
			return nil, fmt.Errorf("unable to scan columns data: %s", err)
		}
		cols = append(cols, c)
	}
	return cols, nil
}
