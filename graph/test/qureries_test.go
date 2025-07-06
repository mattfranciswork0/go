package graph

import (
	"context"
	"testing"

	"github.com/mattfranciswork0/go/db"
)

func TestGetAlbums(t *testing.T) {
	ctx := context.Background()

	// Setup: Insert a test album if needed for non-empty test
	_, err := db.CreateAlbum(ctx, "Test Title", "Test Artist", 10.99)
	if err != nil {
		t.Fatalf("Failed to insert test album: %v", err)
	}

	albums, err := db.GetAlbums(ctx)
	if err != nil {
		t.Errorf("GetAlbums returned error: %v", err)
	}

	if len(albums) == 0 {
		t.Error("Expected at least one album, got none")
	}

	if albums[0].Title == "" {
		t.Error("Album title should not be empty")
	}
}

func TestCreateAlbum(t *testing.T) {
	ctx := context.Background()
	title := "New Album"
	artist := "New Artist"
	price := 12.34

	album, err := db.CreateAlbum(ctx, title, artist, price)
	if err != nil {
		t.Errorf("CreateAlbum returned error: %v", err)
	}

	if album.Title != title || album.Artist != artist || album.Price != price {
		t.Errorf("CreateAlbum returned unexpected result: %+v", album)
	}
}
