/**
 * @author Susheelkumar S
 * @version V1
 * @create date 2020-10-11
 * @desc REST API Service Router
 */

package ImageAlbum

import (
	deleteService "ImageAlbumService/ImageAlbum/deleteImplementations"
	postService "ImageAlbumService/ImageAlbum/postImplementations"
)

//APIServiceHandler REST API Method Handler structure
type APIServiceHandler struct {
	PostHandler   *postService.PostHandler
	DeleteHandler *deleteService.DeleteHandler
}

//Initalization initalizes APIServiceHandler
func Initalization() *APIServiceHandler {
	APIServiceHandler := new(APIServiceHandler)
	APIServiceHandler.PostHandler = new(postService.PostHandler)
	APIServiceHandler.DeleteHandler = new(deleteService.DeleteHandler)
	return APIServiceHandler
}
