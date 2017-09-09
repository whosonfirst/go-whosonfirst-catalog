package flags

import (
	"errors"
	"github.com/whosonfirst/go-whosonfirst-catalog"
	"github.com/whosonfirst/go-whosonfirst-catalog/index"
	"strconv"
	"strings"
)

type PgSQLFlags struct {
	catalog.Flags
	flags []string
}

func (fl *PgSQLFlags) String() string {
	return strings.Join(fl.flags, "\n")
}

func (fl *PgSQLFlags) Set(value string) error {
	fl.flags = append(fl.flags, value)
	return nil
}

func (fl PgSQLFlags) ToIndexes() ([]catalog.Index, error) {

	indexes := make([]catalog.Index, 0)

	for _, dsn := range fl.flags {

		host := "localhost"
		port := 5432
		user := "postgres"
		password := ""
		dbname := "whosonfirst"
		maxconns := 10

		parts := strings.Split(dsn, " ")

		for _, pair := range parts {

			kv := strings.Split(pair, "=")

			if len(kv) != 2 {
				return nil, errors.New("Invalid DSN")
			}

			key := kv[0]
			value := kv[1]

			switch key {
			case "host":
				host = value
			case "port":

				p, err := strconv.Atoi(value)

				if err != nil {
					return nil, err
				}

				port = p
			case "user":
				user = value
			case "password":
				password = value
			case "dbname":
				dbname = value
			default:
				// pass
			}
		}

		idx, err := index.NewPgSQLIndex(host, port, user, password, dbname, maxconns)

		if err != nil {
			return nil, err
		}

		indexes = append(indexes, idx)
	}

	return indexes, nil
}
