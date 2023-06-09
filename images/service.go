package images

import (
	"errors"
	"fmt"
	"sync"

	"github.com/google/uuid"
	"github.com/mayankmike86/imageStores/albums"
)

type Image struct {
	ID      string `json:"id"`
	AlbumID string `json:"albumId"`
	URL     string `json:"url"`
}

type AlbumService interface {
	GetAlbumByID(id string) (*albums.Album, error)
}

type ImageService struct {
	mu           sync.Mutex
	images       []Image
	albumService AlbumService
}

func NewImageService(albumService AlbumService) *ImageService {
	return &ImageService{
		images:       make([]Image, 0),
		albumService: albumService,
	}
}

func (s *ImageService) CreateImage(image Image) (string, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	// Check if the album ID already exists
	_, err := s.albumService.GetAlbumByID(image.AlbumID)
	if err != nil {
		return "album ID does not exis", fmt.Errorf("album ID does not exist")
	}

	// Check if the image ID already exists
	for _, existingImage := range s.images {
		if existingImage.ID == image.ID {
			return "image ID already exists", fmt.Errorf("image ID already exists")
		}
	}

	// Generate a new unique ID if ID is not provided
	if image.ID == "" {
		image.ID = uuid.New().String()
	}

	s.images = append(s.images, image)
	return image.ID, nil
}

func (s *ImageService) DeleteImage(imageID, albumID string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	// Check if the album ID already exists
	_, err := s.albumService.GetAlbumByID(albumID)
	if err != nil {
		return fmt.Errorf("album ID does not exist")
	}

	index := -1
	for i, image := range s.images {
		if image.ID == imageID {
			index = i
			break
		}
	}

	if index == -1 {
		return errors.New("image not found")
	}

	s.images = append(s.images[:index], s.images[index+1:]...)
	return nil
}
