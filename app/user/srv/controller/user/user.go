package user

import (
	v1 "github.com/coderi421/goshop/api/user/v1"
	srv1 "github.com/coderi421/goshop/app/user/srv/service/v1"
)

type userServer struct {
	srv srv1.ServiceFactory
}

var _ v1.UserServer = &userServer{}

func NewUserServer(srv srv1.ServiceFactory) *userServer {
	return &userServer{srv: srv}
}

func DTOToResponse(userDTO srv1.UserDTO) *v1.UserInfoResponse {
	//在grpc的message中字段有默认值，你不能随便赋值nil进去，容易出错
	//这里要搞清，哪些字段是由默认值的
	userInfoRsp := &v1.UserInfoResponse{
		Id:       userDTO.ID,
		PassWord: userDTO.Password,
		Mobile:   userDTO.Mobile,
		NickName: userDTO.NickName,
		Gender:   userDTO.Gender,
		Role:     int32(userDTO.Role),
	}
	if userDTO.Birthday != nil {
		userInfoRsp.BirthDay = uint64(userDTO.Birthday.Unix())
	}
	//底层Mutex不能copy，应该使用指针
	return userInfoRsp
}
