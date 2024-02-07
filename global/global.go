package global

import (
	"github.com/wwengg/im/model"
	"github.com/wwengg/simple/core/slog"
	"github.com/wwengg/simple/core/srpc"
	"github.com/wwengg/simple/core/store"
	clientv3 "go.etcd.io/etcd/client/v3"
	"gorm.io/gorm"
)

var (
	CONFIG    *Config
	LOG       slog.Slog
	SRPC      srpc.SRPC
	DBList    map[string]*gorm.DB
	ClientV3  *clientv3.Client
	RedisBase *store.RedisBase
)

func InitSlog() {
	// 初始化日志
	LOG = slog.NewZapLog(&CONFIG.Slog)
}

func InitSRPC() {
	// 初始化SRPC
	SRPC = srpc.NewSRPCClients(&CONFIG.RPC)
}

func InitDB() {
	// 初始化DBList
	DBList = store.DBList(&CONFIG.DBList)

	model.DBUpms = DBList["upms"]

	model.DBUpms.AutoMigrate(model.CmdService{})
}

func InitEtcdV3() {
	cfg := clientv3.Config{
		Endpoints: CONFIG.RPC.RegisterAddr,
	}
	c, err := clientv3.New(cfg)
	if err != nil {
		LOG.Errorf("clientv3.New err = %s", err.Error())
		panic(err)
		return
	}
	ClientV3 = c
}

func InitRedisBase() {
	RedisBase = store.NewCache(CONFIG.Redis)
}
