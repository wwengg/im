// @Title
// @Description
// @Author  Wangwengang  2023/12/21 23:44
// @Update  Wangwengang  2023/12/21 23:44
package model

import (
	"errors"
	"github.com/wwengg/simple/core/store"
	"github.com/wwengg/simple/proto/pbbase"
	"gorm.io/gorm"
)

type ServerType int32

var DBUpms *gorm.DB

const (
	ServerType_GateHttp ServerType = iota + 1
	ServerType_GateTcp
	ServerType_GateWebsocket
	ServerType_Logic
	ServerType_RPCX_service
)

type ServerInfo struct {
	store.BASE_MODEL
	Name       string     `json:"name" form:"name" gorm:"column:name;comment:服务名称;type:varchar(20);size:20;"`
	Ip         string     `json:"ip" form:"ip" gorm:"column:ip;comment:服务名称;type:varchar(20);size:20;"`
	Port       string     `json:"port" form:"prot" gorm:"column:port;comment:端口号;type:varchar(8);size:8"`
	ServerType ServerType `json:"serverType" form:"serverType" gorm:"column:server_type;comment:服务类型;type:smallint(6);size:6;"`
}

func (model *ServerInfo) Proto() *ServerInfo {
	return &ServerInfo{}
}

func CreateServerInfo(serverInfo *ServerInfo) error {
	return DBUpms.Create(&serverInfo).Error
}

func DeleteServerInfo(id int64) error {
	serverInfo := ServerInfo{
		BASE_MODEL: store.BASE_MODEL{ID: id},
	}
	return DBUpms.Delete(&serverInfo).Error
}

func UpdateServerInfo(serverInfo *ServerInfo) error {
	return DBUpms.Save(&serverInfo).Error
}

func GetServerInfoById(id int64) (serverInfo ServerInfo, err error) {
	if err = DBUpms.Where("id = ?", id).First(&serverInfo).Error; err == nil {
		return
	} else {

		return
	}
}

func GetServerInfoList(info pbbase.PageInfo) (list []ServerInfo, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := DBUpms.Model(&ServerInfo{})
	var ServerInfoList []ServerInfo
	// 此处增加查询条件
	//if info.Keyword != "" {
	//	db.Where("keywaord = ?", info.Keyword)
	//}
	err = db.Count(&total).Error
	if err != nil {
		return ServerInfoList, total, err
	} else {
		err = db.Limit(int(limit)).Offset(int(offset)).Find(&ServerInfoList).Error
	}
	return ServerInfoList, total, err
}
func GetServerInfoByIpAndPortAndServerType(info *ServerInfo) error {
	if err := DBUpms.Where("port = ?", info.Port).Where("ip = ?", info.Ip).Where("server_type = ?", info.ServerType).First(&info).Error; err == nil {
		return nil
	} else {
		return err
	}
}

func ServerInit(info *ServerInfo) error {
	if err := GetServerInfoByIpAndPortAndServerType(info); err == nil {
		return nil
	} else {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return CreateServerInfo(info)
		} else {
			return err
		}
	}
}
