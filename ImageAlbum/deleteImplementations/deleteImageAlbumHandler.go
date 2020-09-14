/**
 * @author Susheelkumar S
 * @version V1
 * @create date 2020-10-11
 * @desc Delete Image Album Router
 */

package deleteImplementations

import (
	"net/http"

	"github.com/labstack/echo"

	"ImageAlbumService/interfaces"
	util "ImageAlbumService/util"
)

//Constants Declaration

const (
	successfulOperation string = "Successful operation"
	imageAlbumNameEmpty string = "Image Album Name is Empty"
	albumDeletionError  string = "Error in deleting Image Album"
)

//DeleteHandler Binder
type DeleteHandler struct {
}

/*
* @desc 	DeleteImageAlbumTable Handler router to Delete Album Query
* @input 	context echo context object
* @output	error, if any
 */
func (deleteApi *DeleteHandler) DeleteImageAlbumTable(context echo.Context) error {

	//Path Parameter reader
	imageAlbumName := context.Param("imageAlbumName")

	defer context.Request().Body.Close()
	context.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)

	if imageAlbumName == "" {
		return util.ResponseMapper(http.StatusMethodNotAllowed, imageAlbumNameEmpty, context)
	}

	//Database call
	err := interfaces.DBClient.DeleteTable(imageAlbumName)
	if err != nil {
		return util.ResponseMapper(http.StatusMethodNotAllowed, albumDeletionError, context)
	}

	return util.ResponseMapper(http.StatusOK, successfulOperation, context)

}
