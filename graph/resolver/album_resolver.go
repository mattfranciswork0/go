package graph

import (
	"context"
	"fmt"
	"os"

	"github.com/mattfranciswork0/go/db"
	"github.com/mattfranciswork0/go/graph/model"
)

func (r *queryResolver) Albums(ctx context.Context) ([]*model.Album, error) {
	albums, err := db.GetAlbums(ctx)

	if err != nil {
		return nil, err
	}

	var gqlAlbums []*model.Album
	for _, u := range albums {
		gqlAlbums = append(gqlAlbums, &model.Album{
			ID: u.ID,
		})
	}

	return gqlAlbums, nil
}

func (r *mutationResolver) CreateAlbum(ctx context.Context, title string, artist string, price float64) (*model.Album, error) {
	albums, err := db.CreateAlbum(ctx, title, artist, price)
	fmt.Println("Albums", albums)
	if err != nil {
		return nil, err
	}
	fmt.Fprintf(os.Stdout, "Album!!: %+v\n", &model.Album{
		Title:  title,
		Artist: artist,
		Price:  price,
	})

	return &model.Album{
		Title:  title,
		Artist: artist,
		Price:  price,
	}, nil
}
