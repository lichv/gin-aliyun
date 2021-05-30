package services

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/services/sts"
)

func GetAccess() (response sts.Credentials, err error){
	code := "dg54df3g"
	configOne, err := GetAliyunConfigOne(map[string]interface{}{"code": code}, "code desc")
	if err != nil {
		return sts.Credentials{"","","",""},nil
	}
	client, err := sts.NewClientWithAccessKey(configOne.Region,configOne.AccessId,configOne.AccessKey)
	request := sts.CreateAssumeRoleRequest()
	request.Scheme = "https"
	request.RoleArn = configOne.Rolearn
	request.RoleSessionName = configOne.Rolesessionrole
	resp, err :=client.AssumeRole(request)
	if err != nil {
		return sts.Credentials{"","","",""},nil
	}
	return (*resp).Credentials, nil
}