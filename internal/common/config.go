package common

import "encoding/json"

type Config struct {
	LogConfig   *LogConfig   `json:"log_config"`
	MysqlConfig *MysqlConfig `json:"db_config"`
	PprofAddr   string       `json:"pprof_addr"`
	ServicePort int          `json:"service_port"`
	AESKey      string       `json:"aes_key"`
	Mode        string       `json:"mode"`
	PidFile     string       `json:"pid_file"`
}

type MysqlConfig struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Username string `json:"user"`
	Password string `json:"password"`
	Database string `json:"database"`
}

// LogConfig 实现了internal/log的配置属性
type LogConfig struct {
	RollType string `json:"roll_type"`
	Dir      string `json:"dir"`
	File     string `json:"file"`
	Count    int32  `json:"count"`
	Size     int64  `json:"size"`
	Uint     string `json:"uint"`
	Level    string `json:"level"`
	Compress int64  `json:"compress"`
}

//// RedisConfig 定义了 redis 连接的配置
//type RedisConfig struct {
//	Address     string `json:"address"`
//	Password    string `json:"password"`
//	Timeout     int    `json:"timeout"`
//	MaxIdle     int    `json:"max_idle"`
//	IdleTimeout int    `json:"idle_timeout"`
//}

//LoadConfig will load config info from []byte,[]byte may read from the config file
func LoadConfig(bs []byte) (*Config, error) {
	cfg := Config{}
	err := json.Unmarshal(bs, &cfg)
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}
