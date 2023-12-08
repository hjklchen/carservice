package data

import "github.com/jmoiron/sqlx"

type RepoFactory interface {
	UserRelated() UserRepo
	PingRelated() PingRepo
}

type datastore struct {
	dbc *sqlx.DB
}

func NewDatastore(dbc *sqlx.DB) RepoFactory {
	return &datastore{dbc}
}

func (ds *datastore) UserRelated() UserRepo {
	return newUser(ds.dbc)
}

func (ds *datastore) PingRelated() PingRepo {
	return newPing(ds.dbc)
}