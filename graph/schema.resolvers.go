package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"github.com/vadlakun/gql-demo-server/graph/api/db"
	"github.com/vadlakun/gql-demo-server/graph/api/errors"
	"time"

	"github.com/vadlakun/gql-demo-server/graph/generated"
	"github.com/vadlakun/gql-demo-server/graph/model"
)


func (r *mutationResolver) CreateVideo(ctx context.Context, input model.NewVideo) (*model.Video, error) {
	newVideo := model.Video{
		URL: input.URL,
		Description: input.Description,
		Name: input.Name,
		CreatedAt: time.Now().UTC(),
	}
	rows, err := db.LogAndQuery(r.db, "INSERT INTO videos (name, description, url, created_at) VALUES($1, $2, $3, $4) RETURNING id",
		input.Name, input.Description, input.URL, newVideo.CreatedAt)
	if err != nil || !rows.Next() {
		return &model.Video{}, err
	}
	defer rows.Close()

	if err := rows.Scan(&newVideo.ID) ; err != nil {
		errors.DebugPrintf(err)
		if errors.IsForeignKeyError(err) {
			return &model.Video{}, errors.UserNotExist
		}
		return &model.Video{}, errors.InternalServerError
	}

	//for _, observer := range videoPublishedChannel {
	//	observer <- newVideo
	//}
	return &newVideo, nil
}

func (r *queryResolver) Videos(ctx context.Context, limit *int, offset *int) ([]*model.Video, error) {
	var video *model.Video
	var videos []*model.Video

	rows, err := db.LogAndQuery(r.db,
		"SELECT id, name, description, url, created_at FROM videos " +
		"ORDER BY created_at desc limit $1 offset $2", limit, offset)
	defer rows.Close()
	if err != nil {
		errors.DebugPrintf(err)
		return nil, errors.InternalServerError
	}
	for rows.Next() {
		if err := rows.Scan(&video.ID, &video.Name, &video.URL, &video.CreatedAt); err != nil {
			errors.DebugPrintf(err)
			return nil, errors.InternalServerError
		}
		videos = append(videos, video)
	}
	return videos, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }


type queryResolver struct{ *Resolver }


/*
func UserFromContext(ctx context.Context) (int) {
	userIDStr, _ := ctx.Value(UserIDCtxKey).(string)
	userID, _ := strconv.Atoi(userIDStr)
	return userID
}
*/