/**
 * @author Susheelkumar S
 * @version V1
 * @create date 2020-10-11
 * @desc Delete Image Album Query
 */

package database

import (
	model "ImageAlbumService/model"
)

//DeleteTable function delete the tables in the database
func (dc *DBRepo) DeleteTable(tableName string) error {

	err := dc.GormDB.Table(tableName).DropTable(&model.ImageAlbumModel{}).Error

	return err

}
