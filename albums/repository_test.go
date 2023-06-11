package albums

import (
	"testing"
)

func TestGetAlbumByID(t *testing.T) {
	service := NewAlbumService()

	// Test case: album found
	album := Album{
		ID:   "album1",
		Name: "Album 1",
	}
	service.CreateAlbum(album)

	result, err := service.GetAlbumByID("album1")
	if err != nil {
		t.Errorf("Expected album found, got error: %v", err)
	}
	if result == nil || result.ID != "album1" {
		t.Errorf("Expected album with ID 'album1', got: %v", result)
	}

	// Test case: album not found
	result, err = service.GetAlbumByID("album2")
	if err == nil {
		t.Errorf("Expected album not found, got nil")
	}
	if result != nil {
		t.Errorf("Expected nil result, got: %v", result)
	}
}

func TestGetAllAlbums(t *testing.T) {
	service := NewAlbumService()

	// Test case: empty albums list
	result := service.GetAllAlbums()
	if len(result) != 0 {
		t.Errorf("Expected empty albums list, got: %v", result)
	}

	// Test case: non-empty albums list
	album1 := Album{
		ID:   "album1",
		Name: "Album 1",
	}
	album2 := Album{
		ID:   "album2",
		Name: "Album 2",
	}
	service.CreateAlbum(album1)
	service.CreateAlbum(album2)

	result = service.GetAllAlbums()
	if len(result) != 2 {
		t.Errorf("Expected albums list length 2, got: %v", result)
	}
}
