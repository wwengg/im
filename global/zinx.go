// @Title
// @Description
// @Author  Wangwengang  2023/12/21 22:37
// @Update  Wangwengang  2023/12/21 22:37
package global

import "github.com/aceld/zinx/zconf"

type Zinx struct {
	/*
		Server
	*/
	Host    string `json:"host" yaml:"host" mapstructure:"host"`            // The IP address of the current server. (当前服务器主机IP)
	TCPPort int    `json:"tcpPort" yaml:"tcp-port" mapstructure:"tcp-port"` // The port number on which the server listens for TCP connections.(当前服务器主机监听端口号)
	WsPort  int    `json:"wsPort" yaml:"ws-port" mapstructure:"ws-port"`    // The port number on which the server listens for WebSocket connections.(当前服务器主机websocket监听端口)
	Name    string `json:"name" yaml:"name" mapstructure:"name"`            // The name of the current server.(当前服务器名称)
	KcpPort int    `json:"kcpPort" yaml:"kcp-port" mapstructure:"kcp-port"` // he port number on which the server listens for KCP connections.(当前服务器主机监听端口号)

	/*
		Zinx
	*/
	Version          string `json:"version" yaml:"version" mapstructure:"version"`                                  // The version of the Zinx framework.(当前Zinx版本号)
	MaxPacketSize    uint32 `json:"maxPacketSize" yaml:"max-packet-size" mapstructure:"max-packet-size"`            // The maximum size of the packets that can be sent or received.(读写数据包的最大值)
	MaxConn          int    `json:"maxConn" yaml:"max-conn" mapstructure:"max-conn"`                                // The maximum number of connections that the server can handle.(当前服务器主机允许的最大链接个数)
	WorkerPoolSize   uint32 `json:"workerPoolSize" yaml:"worker-pool-size" mapstructure:"worker-pool-size"`         // The number of worker pools in the business logic.(业务工作Worker池的数量)
	MaxWorkerTaskLen uint32 `json:"maxWorkerTaskLen" yaml:"max-worker-task-len" mapstructure:"max-worker-task-len"` // The maximum number of tasks that a worker pool can handle.(业务工作Worker对应负责的任务队列最大任务存储数量)
	WorkerMode       string `json:"workerMode" yaml:"worker-mode" mapstructure:"worker-mode"`                       // The way to assign workers to connections.(为链接分配worker的方式)
	MaxMsgChanLen    uint32 `json:"maxMsgChanLen" yaml:"max-msg-chan-len" mapstructure:"max-msg-chan-len"`          // The maximum length of the send buffer message queue.(SendBuffMsg发送消息的缓冲最大长度)
	IOReadBuffSize   uint32 `json:"IOReadBuffSize" yaml:"io-read-buff-size" mapstructure:"io-read-buff-size"`       // The maximum size of the read buffer for each IO operation.(每次IO最大的读取长度)

	//The server mode, which can be "tcp" or "websocket". If it is empty, both modes are enabled.
	//"tcp":tcp监听, "websocket":websocket 监听 为空时同时开启
	Mode string `json:"mode" yaml:"mode" mapstructure:"mode"`

	// A boolean value that indicates whether the new or old version of the router is used. The default value is false.
	// 路由模式 false为旧版本路由，true为启用新版本的路由 默认使用旧版本
	RouterSlicesMode bool `json:"routerSlicesMode" yaml:"router-slices-mode" mapstructure:"router-slices-mode"`

	/*
		logger
	*/
	LogDir string `json:"logDir" yaml:"log-dir" mapstructure:"log-dir"` // The directory where log files are stored. The default value is "./log".(日志所在文件夹 默认"./log")

	// The name of the log file. If it is empty, the log information will be printed to stderr.
	// (日志文件名称   默认""  --如果没有设置日志文件，打印信息将打印至stderr)
	LogFile string `json:"logFile" yaml:"log-file" mapstructure:"log-file"`

	LogSaveDays int   `json:"logSaveDays" yaml:"log-save-days" mapstructure:"log-save-days"` // 日志最大保留天数
	LogFileSize int64 `json:"logFileSize" yaml:"log-file-size" mapstructure:"log-file-size"` // 日志单个日志最大容量 默认 64MB,单位：字节，记得一定要换算成MB（1024 * 1024）
	LogCons     bool  `json:"logCons" yaml:"log-cons" mapstructure:"log-cons"`               // 日志标准输出  默认 false

	// The level of log isolation. The values can be 0 (all open), 1 (debug off), 2 (debug/info off), 3 (debug/info/warn off), and so on.
	// 日志隔离级别  -- 0：全开 1：关debug 2：关debug/info 3：关debug/info/warn ...
	LogIsolationLevel int `json:"logIsolationLevel" yaml:"log-isolation-level" mapstructure:"log-isolation-level"`

	/*
		Keepalive
	*/
	// The maximum interval for heartbeat detection in seconds.
	// 最长心跳检测间隔时间(单位：秒),超过改时间间隔，则认为超时，从配置文件读取
	HeartbeatMax int `json:"heartbeatMax" yaml:"heartbeat-max" mapstructure:"heartbeat-max"`

	/*
		TLS
	*/
	CertFile       string `json:"certFile" yaml:"cert-file" mapstructure:"cert-file"` // The name of the certificate file. If it is empty, TLS encryption is not enabled.(证书文件名称 默认"")
	PrivateKeyFile string `json:"privateKeyFile" yaml:"private-key-file" yaml:"private-key-file"`
}

func (z *Zinx) ToZinxConfig() *zconf.Config {
	return &zconf.Config{
		Host:              z.Host,
		TCPPort:           z.TCPPort,
		WsPort:            z.WsPort,
		Name:              z.Name,
		KcpPort:           z.KcpPort,
		Version:           z.Version,
		MaxPacketSize:     z.MaxPacketSize,
		MaxConn:           z.MaxConn,
		WorkerPoolSize:    z.WorkerPoolSize,
		MaxWorkerTaskLen:  z.MaxWorkerTaskLen,
		WorkerMode:        z.WorkerMode,
		MaxMsgChanLen:     z.MaxMsgChanLen,
		IOReadBuffSize:    z.IOReadBuffSize,
		Mode:              z.Mode,
		RouterSlicesMode:  z.RouterSlicesMode,
		LogDir:            z.LogDir,
		LogFile:           z.LogFile,
		LogSaveDays:       z.LogSaveDays,
		LogFileSize:       z.LogFileSize,
		LogCons:           z.LogCons,
		LogIsolationLevel: z.LogIsolationLevel,
		HeartbeatMax:      z.HeartbeatMax,
		CertFile:          z.CertFile,
		PrivateKeyFile:    z.PrivateKeyFile,
	}
}
