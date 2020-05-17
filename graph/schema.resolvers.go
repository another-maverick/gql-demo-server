package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/another-maverick/gql-demo-server/graph/api/db"
	gqlerrors "github.com/another-maverick/gql-demo-server/graph/api/errors"
	"github.com/another-maverick/gql-demo-server/graph/generated"
	"github.com/another-maverick/gql-demo-server/graph/model"
)



func NewRootResolvers(db *sql.DB) generated.Config {
	c := generated.Config{
		Resolvers: &Resolver{
			db: db,
		},
	}
	return c
}


func (r *mutationResolver) CreateVideo(ctx context.Context, input model.NewVideo) (*model.Video, error) {
	newVideo := model.Video{
		URL:         input.URL,
		Description: input.Description,
		Name:        input.Name,
		CreatedAt:   time.Now().String(),
	}

	dbCon, err := db.Connect()
	if err != nil {
		panic(err)
	}


	fmt.Printf("%T", r.db)
	rows, err := db.LogAndQuery(dbCon, "INSERT INTO videos (name, description, url, created_at) VALUES($1, $2, $3, $4) RETURNING id",
		input.Name, input.Description, input.URL, newVideo.CreatedAt)
	if err != nil || !rows.Next() {
		return &model.Video{}, err
	}
	defer rows.Close()

	if err := rows.Scan(&newVideo.ID); err != nil {
		gqlerrors.DebugPrintf(err)
		if gqlerrors.IsForeignKeyError(err) {
			return &model.Video{}, gqlerrors.UserNotExist
		}
		return &model.Video{}, gqlerrors.InternalServerError
	}

	return &newVideo, nil
}

func (r *queryResolver) Videos(ctx context.Context, limit *int, offset *int) ([]*model.Video, error) {
	var videos []*model.Video

	dbCon, err := db.Connect()
	if err != nil {
		panic(err)
	}

	fmt.Printf("%T", dbCon)
	fmt.Println("About to execute search query..")
	rows, err := db.LogAndQuery(dbCon,
		"SELECT id, name, description, url, created_at FROM videos "+
			"ORDER BY created_at desc limit $1 offset $2", limit, offset)
	fmt.Println("executed search query..")
	defer rows.Close()
	if err != nil {
		gqlerrors.DebugPrintf(err)
		return nil, gqlerrors.InternalServerError
	}
	for rows.Next() {
		var video model.Video
		if err := rows.Scan(&video.ID, &video.Name, &video.Description, &video.URL, &video.CreatedAt); err != nil {
			gqlerrors.DebugPrintf(err)
			return nil, gqlerrors.InternalServerError
		}
		fmt.Println(video.ID, video.Name, video.Description, video.URL, video.CreatedAt)
		videos = append(videos, &video)
	}
	return videos, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{
	*Resolver }
type queryResolver struct{ *Resolver }
