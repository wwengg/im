// Code generated by protoc-gen-simple. DO NOT EDIT.
// versions:
// - protoc-gen-simple v0.0.4
// - protoc          v3.17.2
// source: pbcmdService/pbcmdService.proto

package model

import (
	"github.com/wwengg/im/proto/pbcmdService"
	"github.com/wwengg/im/proto/pbcommon"
	store "github.com/wwengg/simple/core/store"
	"time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = store.TODO

// CmdService Model
type CmdService struct {
	store.BASE_MODEL

	Cmd         int64  `json:"cmd" gorm:"column:cmd;comment: ;type:bigint(20);size:20;"`
	ServicePath string `json:"servicePath" gorm:"column:service_path;comment: ;type:varchar(20);size:20;"`
	ServiceName string `json:"serviceName" gorm:"column:service_name;comment: ;type:varchar(20);size:20;"`
}

func (model *CmdService) Proto() *pbcmdService.CmdServiceModel {
	return &pbcmdService.CmdServiceModel{
		Id:        model.ID,
		CreatedAt: model.CreatedAt.Format(time.DateTime),
		UpdatedAt: model.UpdatedAt.Format(time.DateTime),

		Cmd:         model.Cmd,
		ServicePath: model.ServicePath,
		ServiceName: model.ServiceName,
	}
}

func CmdServiceProtoToModel(proto *pbcmdService.CmdServiceModel) *CmdService {
	cmdService := &CmdService{
		BASE_MODEL: store.BASE_MODEL{
			ID: proto.Id,
		},

		Cmd:         proto.Cmd,
		ServicePath: proto.ServicePath,
		ServiceName: proto.ServiceName,
	}
	if createdAt, err := time.Parse(time.DateTime, proto.CreatedAt); err == nil {
		cmdService.CreatedAt = createdAt
	}
	if updatedAt, err := time.Parse(time.DateTime, proto.UpdatedAt); err == nil {
		cmdService.CreatedAt = updatedAt
	}
	return cmdService
}

// CreateCmdService Func 创建
func CreateCmdService(a CmdService) (err error) {
	err = DBUpms.Create(&a).Error
	return err
}

// DeleteCmdService  删除
func DeleteCmdService(a CmdService) (err error) {
	err = DBUpms.Delete(&a).Error
	return err
}

// UpdateCmdService 修改
func UpdateCmdService(a *CmdService) (err error) {
	err = DBUpms.Save(a).Error
	return err
}

// UpdateCmdService 查询
func GetCmdService(id int64) (result CmdService, err error) {
	err = DBUpms.Where("id = ?", id).First(&result).Error
	return
}

// UpdateCmdService 查询
func GetCmdServiceByCmd(cmd int32) (result CmdService, err error) {
	err = DBUpms.Where("cmd = ?", cmd).First(&result).Error
	return
}

// 分页查询
func GetCmdServiceList(info pbcommon.PageInfo) (list []CmdService, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := DBUpms.Model(&CmdService{})
	var CmdServiceList []CmdService
	// 此处增加查询条件
	//if info.Keyword != "" {
	//	db.Where("keywaord = ?", info.Keyword)
	//}
	err = db.Count(&total).Error
	if err != nil {
		return CmdServiceList, total, err
	} else {
		err = db.Limit(int(limit)).Offset(int(offset)).Find(&CmdServiceList).Error
	}
	return CmdServiceList, total, err
}