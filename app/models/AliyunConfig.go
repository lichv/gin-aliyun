package models

import (
	"fmt"
	"gin-aliyun/utils"
	"github.com/jinzhu/gorm"
)

type AliyunConfig struct {
	gorm.Model
	Code string `json:"code" form:"code" gorm:"code;PRIMARY_KEY"`
	Group string `json:"group" form:"group" gorm:"group"`
	Company string `json:"company" form:"company" gorm:"company"`
	AccessId string `json:"access_id" form:"access_id" gorm:"access_id"`
	AccessKey string `json:"access_key" form:"access_key" gorm:"access_key"`
	Region string `json:"region" form:"region" gorm:"region"`
	Rolesessionrole string `json:"rolesessionrole" form:"rolesessionrole" gorm:"rolesessionrole"`
	Rolearn string `json:"rolearn" form:"rolearn" gorm:"rolearn"`
	Tags string `json:"tags" form:"tags" gorm:"tags"`
	Flag bool `json:"flag" form:"flag" gorm:"flag"`
	State bool `json:"state" form:"state" gorm:"state"`
}

func FindAliyunConfigByCode( code string) ( aliyunConfig *AliyunConfig, err error) {
	err = db.Model(&AliyunConfig{}).Where("code = ? ",code).First(&aliyunConfig).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return &AliyunConfig{},err
	}
	return aliyunConfig, err
}

func GetAliyunConfigOne( query map[string]interface{},orderBy interface{}) ( *AliyunConfig,error) {
	var aliyunConfig AliyunConfig
	model := db.Model(&AliyunConfig{})
	for key, value := range query {
		b,err := utils.In ([]string{"code", "group", "company", "access_id", "access_key", "region", "rolesessionrole", "rolearn", "tags", "flag", "state"},key)
		if  err == nil && b{
			model = model.Where(key + "= ?", value)
		}
	}
	err := model.First(&aliyunConfig).Error
	fmt.Println(aliyunConfig)
	if err != nil && err != gorm.ErrRecordNotFound{
		return &AliyunConfig{},nil
	}
	fmt.Println(aliyunConfig)
	return &aliyunConfig, nil
}

func GetAliyunConfigPages( query map[string]interface{},orderBy interface{},pageNum int,pageSize int) ( []*AliyunConfig, []error) {
	var AliyunConfigs []*AliyunConfig
	var errs []error
	model := db.Where("state=?",true)
	for key, value := range query {
		b,err := utils.In ([]string{"code", "group", "company", "access_id", "access_key", "region", "rolesessionrole", "rolearn", "tags", "flag", "state"},key)
		if  err == nil && b{
			model = model.Where(key + "= ?", value)
		}
	}
	model =model.Offset(pageNum).Limit(pageSize).Order(orderBy).Find(&AliyunConfigs)
	errs = model.GetErrors()
	//err = model.Offset(pageNum).Limit(pageSize).Order(orderBy).Find(&AliyunConfigs).Error

	return AliyunConfigs, errs
}

func ExistAliyunConfigByCode(code string) (b bool,err error) {
	var aliyunConfig AliyunConfig
	err = db.Model(&AliyunConfig{}).Select("code").Where("code = ? ",code).First(&aliyunConfig).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false,err
	}
	return true, err
}

func GetAliyunConfigTotal(maps interface{}) (count int,err error) {
	err = db.Model(&AliyunConfig{}).Where("state = ?",true).Count(&count).Error
	if err != nil {
		return 0,err
	}
	return count, err
}

func AddAliyunConfig( data map[string]interface{}) error {
	AliyunConfig := AliyunConfig{
		Code:data["code"].(string),
		Group:data["group"].(string),
		Company:data["company"].(string),
		AccessId:data["access_id"].(string),
		AccessKey:data["access_key"].(string),
		Region:data["region"].(string),
		Rolesessionrole:data["rolesessionrole"].(string),
		Rolearn:data["rolearn"].(string),
		Tags:data["tags"].(string),
		Flag:data["flag"].(bool),
		State:data["state"].(bool),
	}
	if err:= db.Create(&AliyunConfig).Error;err != nil{
		return err
	}
	return nil
}

func EditAliyunConfig( code string,data map[string]interface{}) error {
	if err:= db.Model(&AliyunConfig{}).Where("code=?",code).Updates(data).Error;err != nil{
		return err
	}
	return nil
}

func DeleteAliyunConfig(code string) error {
	if err := db.Where("code=?",code).Delete(AliyunConfig{}).Error;err != nil{
		return err
	}
	return nil
}

func DeleteAliyunConfigs(maps map[string]interface{}) error{
	model := db
	for key, value := range maps {
		b,err := utils.In ([]string{"code", "group", "company", "access_id", "access_key", "region", "rolesessionrole", "rolearn", "tags", "flag", "state"},key)
		if  err == nil && b{
			model = model.Where(key + "= ?", value)
		}
	}
	if err :=model.Delete(&AliyunConfig{}).Error;err != nil{
		return err
	}
	return nil
}

func ClearAllAliyunConfig() error {
	if err := db.Unscoped().Delete(&AliyunConfig{}).Error; err != nil {
		return err
	}
	return nil
}
