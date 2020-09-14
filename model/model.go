/**
 * @author Susheelkumar S
 * @version V1
 * @create date 2020-10-11
 * @desc JSON-to-DB-Interaction Models
 */

package model

import "time"

//ImageModel For Album Creation
type ImageModel struct {
	ImageAlbumName string `json:"ImageAlbumName"`
}

//ImageAlbumModel field particulars
type ImageAlbumModel struct {
	ImageID           int       `json:"imageId" gorm:"column:ImageID;primary_key;AUTO_INCREMENT"`
	ImageName         string    `json:"imageName" gorm:"column:ImageName"`
	ImageBase64Format string    `json:"imageBase64Format" gorm:"column:ImageBase64Format"`
	ImageCreatedAt    time.Time `json:"createdAt" gorm:"column:CreatedTime"`
}

//APIResponse struct defines structure for the responses generated
type APIResponse struct {
	Code    int    `json:"code"`
	Type    string `json:"type"`
	Message string `json:"message"`
}
