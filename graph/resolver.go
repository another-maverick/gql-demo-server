package graph

import (
	"database/sql"
	"github.com/vadlakun/gql-demo-server/graph/model"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.


var videoPublishedChannel map[string]chan model.Video

func init() {
	videoPublishedChannel = map[string]chan model.Video{}
}

type contextKey string

type Resolver struct{
	db *sql.DB
}
