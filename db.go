package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type column struct {
	Name     string
	Type     string
	Default  *string
	IsNull   bool
	IsArray  bool
	Category string
}

func dbConnect(opts Options) (*sql.DB, error) {
	db, err := sql.Open("postgres",
		fmt.Sprintf(
			"postgresql://%s:%s@%s:%d/%s?connect_timeout=10&application_name=gogen&sslmode=%s",
			opts.User,
			opts.Password,
			opts.Host,
			opts.Port,
			opts.Database,
			opts.SSLMode,
		),
	)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func columnList(db *sql.DB, tab string) (cols []*column, err error) {
	rows, err := db.Query(ColumnsQuery, tab)
	if err != nil {
		return nil, fmt.Errorf("unable to query data: %s", err)
	}

	for rows.Next() {
		c := &column{}
		err := rows.Scan(&c.Name, &c.Type, &c.Default, &c.IsNull, &c.IsArray, &c.Category)
		if err != nil {
			return nil, fmt.Errorf("unable to scan columns data: %s", err)
		}
		cols = append(cols, c)
	}
	return cols, nil
}
