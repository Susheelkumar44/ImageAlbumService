/**
 * @author Susheelkumar S
 * @version V1
 * @create date 2020-10-11
 * @desc DB Interface methods
 */

package interfaces

import (
	config "ImageAlbumService/config"
)

//DBClient - DBInteractions Interface Object
var DBClient DBInteractions

//DBInteractions Interface
type DBInteractions interface {
	DBConnect(config config.DBConfig) error
	CreateTable(tableName string) error
	DeleteTable(tableName string) error
}
