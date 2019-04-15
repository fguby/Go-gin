package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Tag struct {
	Base_model
	Name       string `json:"name" form:"name"`
	CreatedBy  string `json:"created_by" form:"created_by"`
	ModifiedBy string `json:"modified_by"`
	State      int    `json:"state" form:"state"`
}

//根据条件分页获取标签
func GetTags(pageNum int, pageSize int, maps interface{}) ([]Tag, error) {
	var (
		tags []Tag
		err  error
	)
	err = db.Where(maps).Find(&tags).Offset(pageNum).Limit(pageSize).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return tags, nil
}

//根据条件获取标签总数
func GetTagTotal(maps interface{}) (int, error) {
	var count int
	if err := db.Model(&Tag{}).Where(maps).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

//根据标签id删除标签
func DeleteTagById(id int) error {
	if err := db.Where("id = ?", id).Delete(Tag{}).Error; err != nil {
		return err
	}
	return nil
}

//修改标签
func UpdateTag(id int, data interface{}) error {
	if err := db.Model(&Tag{}).Where("id = ? AND deleted_on = ?", id, 0).
		Updates(data).Error; err != nil {
		return err
	}
	return nil
}

//添加标签
func AddTag(data interface{}) error {
	if err := db.Model(&Tag{}).Create(data).Error; err != nil {
		return err
	}
	return nil
}

//添加创建前的回调
func (tag *Tag) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("CreatedOn", time.Now().Unix())

	return nil
}

//添加修改前的回调
func (tag *Tag) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("ModifiedOn", time.Now().Unix())

	return nil
}
