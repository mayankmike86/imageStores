package images

import "errors"

func (s *ImageService) GetImageByID(imageID, albumID string) (*Image, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	for _, image := range s.images {
		if image.ID == imageID && image.AlbumID == albumID {
			return &image, nil
		}
	}

	return nil, errors.New("image not found")
}

func (s *ImageService) GetAllImagesInAlbum(albumID string) []Image {
	s.mu.Lock()
	defer s.mu.Unlock()

	var albumImages []Image
	for _, image := range s.images {
		if image.AlbumID == albumID {
			albumImages = append(albumImages, image)
		}
	}

	return albumImages
}
