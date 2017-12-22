package tool

import (
	"flag"
	"github.com/Terry-Mao/goconf"
	"time"
	"runtime"
)

var (
	gconf *goconf.Config
	Conf *Config
	confFile string
)

func init() {
	flag.StringVar(&confFile, "c", "./im.conf", " set im config file path")
}

type Config struct {
	// base section
	PidFile          string `goconf:"base:pidfile"`
	Dir              string        `goconf:"base:dir"`
	Log              string        `goconf:"base:log"`
	MaxProc          int `goconf:"base:maxproc"`
	PprofAddrs       []string      `goconf:"base:pprof.addrs:,"`
	RPCAddrs         []string      `goconf:"base:rpc.addrs:,"`
	HTTPAddrs        []string      `goconf:"base:http.addrs:,"`
	HTTPReadTimeout  time.Duration `goconf:"base:http.read.timeout:time"`
	HTTPWriteTimeout time.Duration `goconf:"base:http.write.timeout:time"`
	// logic Rpc
	LogicRpcAddrs    map[string]string `-`
	// kafka
	KafkaAddrs       []string `goconf:"kafka:addrs"`
	// monitor
	MonitorOpen      bool `goconf:"monitor:open"`
	MonitorAddrs     []string `goconf:"monitor:addrs:,"`
	// mysql
	RPCConf          RpcConfig
}

type RpcConfig struct {
	MysqlAddr         string
	MysqlName         string
	MysqlMaxIdleConns int64
	MysqlMaxOpenConns int64
}

func NewConfig() *Config {
	return &Config{
		// base section
		PidFile: "/tmp/goim-im.pid",
		Dir: "./",
		Log: "./im-log.xml",
		MaxProc: runtime.NumCPU(),
		PprofAddrs: []string{"local:7471"},
		HTTPAddrs:      []string{"7472"},
		LogicRpcAddrs: make(map[string]string),
	}
}

// InitConfig init the global config.
func InitConfig(configFile string) (*Config, error) {
	if configFile != "" {
		confFile = configFile
	}
	var err error
	Conf = NewConfig()
	gconf = goconf.New()
	if err = gconf.Parse(confFile); err != nil {
		return nil, err
	}
	if err := gconf.Unmarshal(Conf); err != nil {
		return nil, err
	}
	for _, serverID := range gconf.Get("logic.addrs").Keys() {
		addr, err := gconf.Get("logic.addrs").String(serverID)
		if err != nil {
			return nil, err
		}
		Conf.LogicRpcAddrs[serverID] = addr
	}
	mysqlName, _ := gconf.Get("mysql").String("name")
	mysqlAddr, _ := gconf.Get("mysql").String("addr")
	mysqlMaxIdleConns, _ := gconf.Get("mysql").Int("maxIdleConns")
	mysqlMaxOpenConns, _ := gconf.Get("mysql").Int("maxOpenConns")
	Conf.RPCConf = RpcConfig{
		MysqlAddr:mysqlAddr,
		MysqlName:mysqlName,
		MysqlMaxIdleConns:mysqlMaxIdleConns,
		MysqlMaxOpenConns:mysqlMaxOpenConns,
	}
	return Conf, nil
}

func ReloadConfig() (*Config, error) {
	conf := NewConfig()
	ngconf, err := gconf.Reload()
	if err != nil {
		return nil, err
	}
	if err := ngconf.Unmarshal(conf); err != nil {
		return nil, err
	}
	gconf = ngconf
	return conf, nil
}