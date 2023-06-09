package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mayankmike86/imageStores/albums"
	api "github.com/mayankmike86/imageStores/apis"
	"github.com/mayankmike86/imageStores/images"
)

func main() {
	albumService := albums.NewAlbumService()
	imageService := images.NewImageService(albumService)

	handler := api.NewHandler(albumService, imageService)

	router := mux.NewRouter()
	router.HandleFunc("/albums", handler.CreateAlbum).Methods(http.MethodPost)
	router.HandleFunc("/albums/{albumID}", handler.DeleteAlbum).Methods(http.MethodDelete) //TODO delete all the related images also
	router.HandleFunc("/albums/{albumID}/images", handler.CreateImage).Methods(http.MethodPost)
	router.HandleFunc("/albums/{albumID}/images/{imageID}", handler.DeleteImage).Methods(http.MethodDelete)
	router.HandleFunc("/albums/{albumID}/images/{imageID}", handler.GetImage).Methods(http.MethodGet)
	router.HandleFunc("/albums/{albumID}/images", handler.GetImagesInAlbum).Methods(http.MethodGet)

	log.Fatal(http.ListenAndServe(":8080", router))
}
