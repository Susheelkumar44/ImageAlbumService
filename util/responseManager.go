/**
 * @author Susheelkumar S
 * @version V1
 * @create date 2020-10-11
 * @desc REST API Response Mapper
 */

package util

import (
	"encoding/json"
	"log"
	"net/http"

	"ImageAlbumService/model"

	"github.com/labstack/echo"
)

//ResponseMapper to handle response
func ResponseMapper(code int, message string, c echo.Context) error {

	log.Println(message)
	response := model.APIResponse{}

	response.Code = code
	response.Type = http.StatusText(code)
	response.Message = message

	c.Response().WriteHeader(code)
	return json.NewEncoder(c.Response()).Encode(response)
}
