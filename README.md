# Installation

## Clone the repository:
git clone https://github.com/example/image-store-service.git
cd image-store-service
## Build and run the application locally:
go build
sudo docker build -t image-store-service .
sudo docker run -p 8080:8080 image-store-service

## The application will be accessible at http://localhost:8080.

# Usage
## Create an album:

### To create image album. Send a POST request to http://localhost:8080/albums with the album data in the request body.
1. curl -X POST -H "Content-Type: application/json" -d '{"id":"albumid1","name":"myalbum1"}' http://localhost:8080/albums -v

### To delete image album Send a DELETE request to http://localhost:8080/albums/{albumID} to deleet the album by ID.
2. curl -X DELETE http://localhost:8080/albums/albumid1 -v

### To create a new image in image album Send a POST request to http://localhost:8080/albums/{albumID}/images with the image data in the request body.

3. curl -X POST -H "Content-Type: application/json" -d '{"id":"imageid1","albumId":"albumid1","url":"https://example.com/image1.jpg"}' http://localhost:8080/albums/{albumID}/images -v

### To delete an image id from given albumId Send a DELETE request to http://localhost:8080/albums/{albumID}/images/{imageID}
4. curl -X DELETE http://localhost:8080/albums/albumid1/images/imageid1 -v

### To get a single image in an album: Send a GET request to http://localhost:8080/albums/{albumID}/images/{imageID}.
5. curl http://localhost:8080/albums/albumid1/images/imageid1 -v

### Get all images in an album: Send a GET request to http://localhost:8080/albums/{albumID}/images.
6. curl http://localhost:8080/albums/albumid1/images -v


# Logs
sudo docker logs -f <container-id>