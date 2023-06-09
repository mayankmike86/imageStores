package albums

import (
	"errors"
	"fmt"
	"sync"

	"github.com/google/uuid"
)

type Album struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type AlbumService struct {
	mu     sync.Mutex
	albums []Album
}

func NewAlbumService() *AlbumService {
	return &AlbumService{}
}

func (s *AlbumService) CreateAlbum(album Album) (string, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	// Check if album ID already exists
	for _, existingAlbum := range s.albums {
		if existingAlbum.ID == album.ID {
			return "album ID already exists", fmt.Errorf("album ID already exists")
		}
	}

	// Generate a new unique ID if ID is not provided
	if album.ID == "" {
		album.ID = uuid.New().String()
	}

	// Create the album
	s.albums = append(s.albums, album)

	return album.ID, nil
}

func (s *AlbumService) DeleteAlbum(albumID string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	index := -1
	for i, album := range s.albums {
		if album.ID == albumID {
			index = i
			break
		}
	}

	if index == -1 {
		return errors.New("album not found")
	}

	s.albums = append(s.albums[:index], s.albums[index+1:]...)
	return nil
}
