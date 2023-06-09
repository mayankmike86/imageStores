package albums

import "errors"

func (s *AlbumService) GetAlbumByID(albumID string) (*Album, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	for _, album := range s.albums {
		if album.ID == albumID {
			return &album, nil
		}
	}

	return nil, errors.New("album not found")
}

func (s *AlbumService) GetAllAlbums() []Album {
	s.mu.Lock()
	defer s.mu.Unlock()

	return s.albums
}
