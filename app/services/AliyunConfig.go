package services

import (
	"fmt"
	"gin-aliyun/app/models"
	"gin-aliyun/utils/jwt"
)

type AliyunConfig struct {
	Code string `json:"code" form:"code"`
	Group string `json:"group" form:"group"`
	Company string `json:"company" form:"company"`
	AccessId string `json:"access_id" form:"access_id"`
	AccessKey string `json:"access_key" form:"access_key"`
	Region string `json:"region" form:"region"`
	Rolesessionrole string `json:"rolesessionrole" form:"rolesessionrole"`
	Rolearn string `json:"rolearn" form:"rolearn"`
	Tags string `json:"tags" form:"tags"`	
	Flag bool `json:"flag" form:"flag"`
	State bool `json:"state" form:"state"`
}

func ExistAliyunConfigByCode(code string) (b bool,err error) {
	b,err = models.ExistAliyunConfigByCode(code)
	return b, err
}

func GetAliyunConfigTotal(maps interface{}) (count int,err error) {
	count,err = models.GetAliyunConfigTotal(map[string]interface{}{})
	return count, err
}
func GetAliyunConfigOne( query map[string]interface{},orderBy interface{}) (AliyunConfig *AliyunConfig, err error) {
	var nu *models.AliyunConfig
	nu,err = models.GetAliyunConfigOne(query,orderBy)
	fmt.Println(nu)
	return TransferAliyunConfigModel(nu),nil
}

func GetAliyunConfigPages( query map[string]interface{},orderBy interface{},pageNum int,pageSize int) (AliyunConfigs []*AliyunConfig, total int, errs []error) {
	count,err := models.GetAliyunConfigTotal(query)
	fmt.Println(count)
	if err != nil {
		return nil,0,errs
	}
	us,errs := models.GetAliyunConfigPages(query,orderBy,pageNum,pageSize)
	AliyunConfigs = TransferAliyunConfigs(us)
	return AliyunConfigs,total,nil
}

func AddAliyunConfig( data map[string]interface{}) (err error ){
	err = models.AddAliyunConfig(data)
	return err
}

func EditAliyunConfig( code string,data map[string]interface{}) (err error) {
	err = models.EditAliyunConfig(code,data)
	return err
}

func DeleteAliyunConfig(maps map[string]interface{}) (err error) {
	err = models.DeleteAliyunConfigs(maps)
	return nil
}

func ClearAllAliyunConfig() (err error) {
	err = models.ClearAllAliyunConfig()
	return err
}

func TransferAliyunConfigModel(u *models.AliyunConfig)(aliyunConfig *AliyunConfig){
	aliyunConfig =  &AliyunConfig{
		Code:u.Code,
		Group:u.Group,
		Company:u.Company,
		AccessId:u.AccessId,
		AccessKey:u.AccessKey,
		Region:u.Region,
		Rolesessionrole:u.Rolesessionrole,
		Rolearn:u.Rolearn,
		Tags:u.Tags,
		Flag:u.Flag,
		State:u.State,
	}
	return
}
func TransferAliyunConfigs(us []*models.AliyunConfig) (AliyunConfigs []*AliyunConfig) {
	for _,value := range us {
		AliyunConfig := TransferAliyunConfigModel(value)
		AliyunConfigs = append(AliyunConfigs, AliyunConfig)
	}
	return AliyunConfigs
}

func GenerateToken(code, AliyunConfigname string) (string, error) {
	return jwt.GenerateToken(code,AliyunConfigname)
}
