package db

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/mattfranciswork0/go/graph/model"
)

func GetAlbums(ctx context.Context) ([]model.Album, error) {
	rows, err := Pool.Query(ctx, "select title from albums LIMIT 1")

	// Get values as []any
	var albums []model.Album
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}
	for rows.Next() {
		var album model.Album
		if err := rows.Scan(&album.Title); err != nil {
			log.Fatal("Fatal error!!", err)
			return nil, err
		}
		albums = append(albums, album)
	}
	fmt.Println("Albums", albums)

	return albums, nil
}

func CreateAlbum(ctx context.Context, title string, artist string, price float64) (*model.Album, error) {
	var album model.Album
	fmt.Println("Artists", artist, &album.Artist)
	err := Pool.QueryRow(ctx,
		"INSERT INTO albums  (title, artist, price) VALUES ($1, $2, 11.1) RETURNING title, artist, price",
		title, artist, price).Scan(&album.Title, &album.Artist, &album.Price)

	if err != nil {
		return nil, err
	}
	return &album, nil
}
