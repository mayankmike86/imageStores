package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mayankmike86/imageStores/albums"
	"github.com/mayankmike86/imageStores/images"
)

type Handler struct {
	albumService *albums.AlbumService
	imageService *images.ImageService
}

func NewHandler(albumService *albums.AlbumService, imageService *images.ImageService) *Handler {
	return &Handler{
		albumService: albumService,
		imageService: imageService,
	}
}

func (h *Handler) CreateAlbum(w http.ResponseWriter, r *http.Request) {
	var album albums.Album
	err := json.NewDecoder(r.Body).Decode(&album)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	//album.ID = uuid.New().String()
	msg, err := h.albumService.CreateAlbum(album)
	if err != nil {
		log.Printf("err while creating new image album: %v , %v", err, msg)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(msg)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(album)
}

func (h *Handler) DeleteAlbum(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	albumID := params["albumID"]

	err := h.albumService.DeleteAlbum(albumID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *Handler) CreateImage(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	albumID := params["albumID"]
	var image images.Image
	err := json.NewDecoder(r.Body).Decode(&image)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	//check if albumId passed in request body is not matching as in path params.
	if albumID != image.AlbumID {
		msg := "albumId passed in request body is not matching as in path params"
		log.Printf("err while creating new image in album: %v", msg)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(msg)
		return
	}

	//image.ID = uuid.New().String()
	msg, err := h.imageService.CreateImage(image)
	if err != nil {
		log.Printf("err while creating new image in album: %v , %v", err, msg)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(msg)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(image)
}

func (h *Handler) DeleteImage(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	imageID := params["imageID"]
	albumID := params["albumID"]

	err := h.imageService.DeleteImage(imageID, albumID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *Handler) GetImage(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	imageID := params["imageID"]
	albumID := params["albumID"]

	image, err := h.imageService.GetImageByID(imageID, albumID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(image)
}

func (h *Handler) GetImagesInAlbum(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	albumID := params["albumID"]

	images := h.imageService.GetAllImagesInAlbum(albumID)
	if len(images) == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(images)
}
