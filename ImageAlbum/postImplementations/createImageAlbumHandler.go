/**
 * @author Susheelkumar S
 * @version V1
 * @create date 2020-10-11
 * @desc Create Image Album Router
 */

package postImplementations

import (
	"encoding/json"
	"net/http"

	"github.com/labstack/echo"

	"ImageAlbumService/interfaces"
	model "ImageAlbumService/model"
	util "ImageAlbumService/util"
)

//Constants Declaration

const (
	decodeJSONErr       string = "Invalid Input: Error in Decoding Input JSON"
	successfulOperation string = "Successful operation"
	imageAlbumEmpty     string = "Image Album Name is Empty"
	albumCreationError  string = "Error in creating Image Album"
)

//PostHandler binder
type PostHandler struct{}

/*
* @desc 	CreateImageAlbumTable Handler router to Create Album Query
* @input 	context echo context object
* @output	error, if any
 */
func (postApi *PostHandler) CreateImageAlbumTable(context echo.Context) error {
	imageAlbumModel := model.ImageModel{}

	//JSON Reader
	errDecode := json.NewDecoder(context.Request().Body).Decode(&imageAlbumModel)

	defer context.Request().Body.Close()
	context.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)

	//validating the format of JSON provided
	if errDecode != nil {
		return util.ResponseMapper(http.StatusMethodNotAllowed, decodeJSONErr, context)
	}

	if imageAlbumModel.ImageAlbumName == "" {
		return util.ResponseMapper(http.StatusMethodNotAllowed, imageAlbumEmpty, context)
	}

	//Database Call
	err := interfaces.DBClient.CreateTable(imageAlbumModel.ImageAlbumName)
	if err != nil {
		return util.ResponseMapper(http.StatusMethodNotAllowed, albumCreationError, context)
	}

	return util.ResponseMapper(http.StatusOK, successfulOperation, context)

}
