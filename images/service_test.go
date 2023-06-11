package images

import (
	"testing"

	"github.com/mayankmike86/imageStores/albums"
)

type MockAlbumService struct{}

func (m *MockAlbumService) GetAlbumByID(id string) (*albums.Album, error) {
	// Implement the mock behavior here
	return nil, nil
}
func TestCreateImage(t *testing.T) {
	albumService := &MockAlbumService{}
	imageService := NewImageService(albumService)

	// Test creating a new image
	image := Image{
		ID:      "1",
		URL:     "ImageURL",
		AlbumID: "album123",
	}
	_, err := imageService.CreateImage(image)
	if err != nil {
		t.Errorf("Error creating image: %v", err)
	}

	// Test creating an image with an existing ID
	duplicateImage := Image{
		ID:      "1",
		URL:     "duplicateImageURL",
		AlbumID: "album456",
	}
	_, err = imageService.CreateImage(duplicateImage)
	if err == nil {
		t.Errorf("Expected error for creating image with duplicate ID, got nil")
	}
}

func TestDeleteImage(t *testing.T) {
	albumService := &MockAlbumService{}
	imageService := NewImageService(albumService)

	// Create a test image
	image := Image{
		ID:      "1",
		URL:     "ImageURL",
		AlbumID: "album123",
	}
	_, err := imageService.CreateImage(image)
	if err != nil {
		t.Errorf("Error creating image: %v", err)
	}

	// Test deleting an existing image
	err = imageService.DeleteImage("1", "album123")
	if err != nil {
		t.Errorf("Error deleting image: %v", err)
	}

	// Test deleting a non-existent image
	err = imageService.DeleteImage("2", "album123")
	if err == nil {
		t.Errorf("Expected error for deleting non-existent image, got nil")
	}
}
