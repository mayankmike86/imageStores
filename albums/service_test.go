package albums

import (
	"testing"
)

func TestCreateAlbum(t *testing.T) {
	service := NewAlbumService()

	// Test creating a new album
	album := Album{
		ID:   "album1",
		Name: "Album 1",
	}
	_, err := service.CreateAlbum(album)
	if err != nil {
		t.Errorf("Error creating album: %v", err)
	}

	// Test creating an album with an existing ID
	duplicateAlbum := Album{
		ID:   "album1",
		Name: "Duplicate Album",
	}
	_, err = service.CreateAlbum(duplicateAlbum)
	if err == nil {
		t.Errorf("Expected error for creating duplicate album, got nil")
	}
}

func TestDeleteAlbum(t *testing.T) {
	service := NewAlbumService()
	album := Album{
		ID:   "album1",
		Name: "Album 1",
	}
	_, err := service.CreateAlbum(album)
	if err != nil {
		t.Errorf("Error creating album: %v", err)
	}
	// Test deleting an existing album
	err = service.DeleteAlbum("album1")
	if err != nil {
		t.Errorf("Error deleting album: %v", err)
	}

	// Test deleting a non-existent album
	err = service.DeleteAlbum("album2")
	if err == nil {
		t.Errorf("Expected error for deleting non-existent album, got nil")
	}
}
