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

// func GetUsers(ctx context.Context) ([]*model.DBUser, error) {
// 	// rows, err := Pool.Query(ctx, "SELECT id, name, email FROM users")
// 	// if err != nil {
// 	// 	return nil, err
// 	// }
// 	// defer rows.Close()

// 	// var users []*model.DBUser
// 	// for rows.Next() {
// 	// 	var u model.DBUser
// 	// 	if err := rows.Scan(&u.ID, &u.Name, &u.Email); err != nil {
// 	// 		return nil, err
// 	// 	}
// 	// 	users = append(users, &u)
// 	// }
// 	var title string
// 	// var weight int64
// 	// err = conn.QueryRow(context.Background(), "select name, weight from widgets where id=$1", 42).Scan(&name, &weight)
// 	err = conn.QueryRow(context.Background(), "select title from albums LIMIT 1").Scan(&title)
// 	fmt.Println(title)
// 	return title, nil
// }

// func GetUserByID(ctx context.Context, id int) (*model.DBUser, error) {
// 	var u model.DBUser
// 	err := Pool.QueryRow(ctx, "SELECT id, name, email FROM users WHERE id=$1", id).
// 		Scan(&u.ID, &u.Name, &u.Email)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &u, nil
// }

// func CreateUser(ctx context.Context, name, email string) (*model.DBUser, error) {
// 	var u model.DBUser
// 	err := Pool.QueryRow(ctx,
// 		"INSERT INTO users (name, email) VALUES ($1, $2) RETURNING id, name, email",
// 		name, email,
// 	).Scan(&u.ID, &u.Name, &u.Email)

// 	if err != nil {
// 		return nil, err
// 	}
// 	return &u, nil
// }
