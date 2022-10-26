package db

import (
	"database/sql"
	"fmt"
	"path"

	"github.com/nleof/goyesql"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

const (
	dbName      = "logger_fitness_db"
	queriesRoot = "bacend/db/queries"
	driverName  = "postgres"
)

var (
	queryFiles = []string{}
)

// Client contains the mongo connection and custom functions.
type Db struct {
	Queries goyesql.Queries
	Client  *sql.DB
}

type DbConfigOpts struct {
	User   string
	Pass   string
	Port   string
	Host   string
	DbName string
}

// TODO: Check if schema is present, if not create???

// Connect for connecting to the mongo serer.
func Connect(creds DbConfigOpts) (*Db, error) {
	db, err := sql.Open(driverName, fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable",
		creds.Host, creds.User, creds.Pass, creds.DbName))

	if err != nil {
		log.WithError(err).Fatal("failed to create database client")
	}

	// read sql queries in from file
	queries := make(goyesql.Queries)
	for _, file := range queryFiles {
		qs, err := goyesql.ParseFile(path.Join(queriesRoot, file))
		if err != nil {
			return nil, errors.Wrap(err, fmt.Sprintf("failed to parse file: '%s'", path.Join(queriesRoot, file)))
		}

		for k, v := range qs {
			queries[k] = v
		}
	}

	return &Db{
		Queries: queries,
		Client:  db,
	}, nil
}
