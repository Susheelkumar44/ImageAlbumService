/**
 * @author Susheelkumar S
 * @version V1
 * @create date 2020-10-11
 * @desc Create Image Album Query
 */

package database

import (
	model "ImageAlbumService/model"
)

//CreateTable function creates the tables in the database
func (dc *DBRepo) CreateTable(tableName string) error {

	err := dc.GormDB.Table(tableName).AutoMigrate(&model.ImageAlbumModel{}).Error

	return err

}
