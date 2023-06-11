package images

import (
	"testing"
)

func TestGetImageByID1(t *testing.T) {
	albumService := &MockAlbumService{}

	service := NewImageService(albumService)

	// Test case: image found
	image1 := Image{
		ID:      "image1",
		AlbumID: "album1",
	}
	service.CreateImage(image1)

	result, err := service.GetImageByID("image1", "album1")
	if err != nil {
		t.Errorf("Expected image found, got error: %v", err)
	}
	if result == nil || result.ID != "image1" || result.AlbumID != "album1" {
		t.Errorf("Expected image with ID 'image1' and AlbumID 'album1', got: %v", result)
	}

	// Test case: image not found
	result, err = service.GetImageByID("image2", "album1")
	if err == nil {
		t.Errorf("Expected image not found, got nil")
	}
	if result != nil {
		t.Errorf("Expected nil result, got: %v", result)
	}
}

func TestGetAllImagesInAlbum(t *testing.T) {
	albumService := &MockAlbumService{}

	service := NewImageService(albumService)

	// Test case: empty album images list
	result := service.GetAllImagesInAlbum("album1")
	if len(result) != 0 {
		t.Errorf("Expected empty album images list, got: %v", result)
	}

	// Test case: non-empty album images list
	image1 := Image{
		ID:      "image1",
		AlbumID: "album1",
	}
	image2 := Image{
		ID:      "image2",
		AlbumID: "album2",
	}
	image3 := Image{
		ID:      "image3",
		AlbumID: "album1",
	}
	service.CreateImage(image1)
	service.CreateImage(image2)
	service.CreateImage(image3)

	result = service.GetAllImagesInAlbum("album1")
	if len(result) != 2 {
		t.Errorf("Expected album images list length 2, got: %v", result)
	}
}
