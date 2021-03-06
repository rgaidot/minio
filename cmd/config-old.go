package cmd

import (
	"os"
	"path/filepath"
	"sync"

	"github.com/minio/minio/pkg/quick"
)

/////////////////// Config V1 ///////////////////
type configV1 struct {
	Version   string `json:"version"`
	AccessKey string `json:"accessKeyId"`
	SecretKey string `json:"secretAccessKey"`
}

// loadConfigV1 load config
func loadConfigV1() (*configV1, error) {
	configPath, err := getConfigPath()
	if err != nil {
		return nil, err
	}
	configFile := filepath.Join(configPath, "fsUsers.json")
	if _, err = os.Stat(configFile); err != nil {
		return nil, err
	}
	c := &configV1{}
	c.Version = "1"
	qc, err := quick.New(c)
	if err != nil {
		return nil, err
	}
	if err := qc.Load(configFile); err != nil {
		return nil, err
	}
	return c, nil
}

/////////////////// Config V2 ///////////////////
type configV2 struct {
	Version     string `json:"version"`
	Credentials struct {
		AccessKey string `json:"accessKeyId"`
		SecretKey string `json:"secretAccessKey"`
		Region    string `json:"region"`
	} `json:"credentials"`
	MongoLogger struct {
		Addr       string `json:"addr"`
		DB         string `json:"db"`
		Collection string `json:"collection"`
	} `json:"mongoLogger"`
	SyslogLogger struct {
		Network string `json:"network"`
		Addr    string `json:"addr"`
	} `json:"syslogLogger"`
	FileLogger struct {
		Filename string `json:"filename"`
	} `json:"fileLogger"`
}

// loadConfigV2 load config version '2'.
func loadConfigV2() (*configV2, error) {
	configFile, err := getConfigFile()
	if err != nil {
		return nil, err
	}
	if _, err = os.Stat(configFile); err != nil {
		return nil, err
	}
	c := &configV2{}
	c.Version = "2"
	qc, err := quick.New(c)
	if err != nil {
		return nil, err
	}
	if err := qc.Load(configFile); err != nil {
		return nil, err
	}
	return c, nil
}

/////////////////// Config V3 ///////////////////

// backendV3 type.
type backendV3 struct {
	Type  string   `json:"type"`
	Disk  string   `json:"disk,omitempty"`
	Disks []string `json:"disks,omitempty"`
}

// syslogLogger v3
type syslogLoggerV3 struct {
	Enable bool   `json:"enable"`
	Addr   string `json:"address"`
	Level  string `json:"level"`
}

// loggerV3 type.
type loggerV3 struct {
	Console struct {
		Enable bool   `json:"enable"`
		Level  string `json:"level"`
	}
	File struct {
		Enable   bool   `json:"enable"`
		Filename string `json:"fileName"`
		Level    string `json:"level"`
	}
	Syslog struct {
		Enable bool   `json:"enable"`
		Addr   string `json:"address"`
		Level  string `json:"level"`
	} `json:"syslog"`
	// Add new loggers here.
}

// configV3 server configuration version '3'.
type configV3 struct {
	Version string `json:"version"`

	// Backend configuration.
	Backend backendV3 `json:"backend"`

	// http Server configuration.
	Addr string `json:"address"`

	// S3 API configuration.
	Credential credential `json:"credential"`
	Region     string     `json:"region"`

	// Additional error logging configuration.
	Logger loggerV3 `json:"logger"`
}

// loadConfigV3 load config version '3'.
func loadConfigV3() (*configV3, error) {
	configFile, err := getConfigFile()
	if err != nil {
		return nil, err
	}
	if _, err = os.Stat(configFile); err != nil {
		return nil, err
	}
	c := &configV3{}
	c.Version = "3"
	qc, err := quick.New(c)
	if err != nil {
		return nil, err
	}
	if err := qc.Load(configFile); err != nil {
		return nil, err
	}
	return c, nil
}

// logger type representing version '4' logger config.
type loggerV4 struct {
	Console struct {
		Enable bool   `json:"enable"`
		Level  string `json:"level"`
	} `json:"console"`
	File struct {
		Enable   bool   `json:"enable"`
		Filename string `json:"fileName"`
		Level    string `json:"level"`
	} `json:"file"`
	Syslog struct {
		Enable bool   `json:"enable"`
		Addr   string `json:"address"`
		Level  string `json:"level"`
	} `json:"syslog"`
}

// configV4 server configuration version '4'.
type configV4 struct {
	Version string `json:"version"`

	// S3 API configuration.
	Credential credential `json:"credential"`
	Region     string     `json:"region"`

	// Additional error logging configuration.
	Logger loggerV4 `json:"logger"`
}

// loadConfigV4 load config version '4'.
func loadConfigV4() (*configV4, error) {
	configFile, err := getConfigFile()
	if err != nil {
		return nil, err
	}
	if _, err = os.Stat(configFile); err != nil {
		return nil, err
	}
	c := &configV4{}
	c.Version = "4"
	qc, err := quick.New(c)
	if err != nil {
		return nil, err
	}
	if err := qc.Load(configFile); err != nil {
		return nil, err
	}
	return c, nil
}

// logger type representing version '5' logger config.
type loggerV5 struct {
	Console struct {
		Enable bool   `json:"enable"`
		Level  string `json:"level"`
	} `json:"console"`
	File struct {
		Enable   bool   `json:"enable"`
		Filename string `json:"fileName"`
		Level    string `json:"level"`
	} `json:"file"`
	Syslog struct {
		Enable bool   `json:"enable"`
		Addr   string `json:"address"`
		Level  string `json:"level"`
	} `json:"syslog"`
	AMQP struct {
		Enable       bool   `json:"enable"`
		Level        string `json:"level"`
		URL          string `json:"url"`
		Exchange     string `json:"exchange"`
		RoutingKey   string `json:"routingKey"`
		ExchangeType string `json:"exchangeType"`
		Mandatory    bool   `json:"mandatory"`
		Immediate    bool   `json:"immediate"`
		Durable      bool   `json:"durable"`
		Internal     bool   `json:"internal"`
		NoWait       bool   `json:"noWait"`
		AutoDeleted  bool   `json:"autoDeleted"`
	} `json:"amqp"`
	ElasticSearch struct {
		Enable bool   `json:"enable"`
		Level  string `json:"level"`
		URL    string `json:"url"`
		Index  string `json:"index"`
	} `json:"elasticsearch"`
	Redis struct {
		Enable   bool   `json:"enable"`
		Level    string `json:"level"`
		Addr     string `json:"address"`
		Password string `json:"password"`
		Key      string `json:"key"`
	} `json:"redis"`
}

// configV5 server configuration version '5'.
type configV5 struct {
	Version string `json:"version"`

	// S3 API configuration.
	Credential credential `json:"credential"`
	Region     string     `json:"region"`

	// Additional error logging configuration.
	Logger loggerV5 `json:"logger"`
}

// loadConfigV5 load config version '5'.
func loadConfigV5() (*configV5, error) {
	configFile, err := getConfigFile()
	if err != nil {
		return nil, err
	}
	if _, err = os.Stat(configFile); err != nil {
		return nil, err
	}
	c := &configV5{}
	c.Version = "5"
	qc, err := quick.New(c)
	if err != nil {
		return nil, err
	}
	if err := qc.Load(configFile); err != nil {
		return nil, err
	}
	return c, nil
}

type loggerV6 struct {
	Console consoleLogger  `json:"console"`
	File    fileLogger     `json:"file"`
	Syslog  syslogLoggerV3 `json:"syslog"`
}

// configV6 server configuration version '6'.
type configV6 struct {
	Version string `json:"version"`

	// S3 API configuration.
	Credential credential `json:"credential"`
	Region     string     `json:"region"`

	// Additional error logging configuration.
	Logger loggerV6 `json:"logger"`

	// Notification queue configuration.
	Notify notifierV1 `json:"notify"`
}

// loadConfigV6 load config version '6'.
func loadConfigV6() (*configV6, error) {
	configFile, err := getConfigFile()
	if err != nil {
		return nil, err
	}
	if _, err = os.Stat(configFile); err != nil {
		return nil, err
	}
	c := &configV6{}
	c.Version = "6"
	qc, err := quick.New(c)
	if err != nil {
		return nil, err
	}
	if err := qc.Load(configFile); err != nil {
		return nil, err
	}
	return c, nil
}

// Notifier represents collection of supported notification queues in version
// 1 without NATS streaming.
type notifierV1 struct {
	AMQP          map[string]amqpNotify          `json:"amqp"`
	NATS          map[string]natsNotifyV1        `json:"nats"`
	ElasticSearch map[string]elasticSearchNotify `json:"elasticsearch"`
	Redis         map[string]redisNotify         `json:"redis"`
	PostgreSQL    map[string]postgreSQLNotify    `json:"postgresql"`
	Kafka         map[string]kafkaNotify         `json:"kafka"`
}

// Notifier represents collection of supported notification queues in version 2
// with NATS streaming but without webhook.
type notifierV2 struct {
	AMQP          map[string]amqpNotify          `json:"amqp"`
	NATS          map[string]natsNotify          `json:"nats"`
	ElasticSearch map[string]elasticSearchNotify `json:"elasticsearch"`
	Redis         map[string]redisNotify         `json:"redis"`
	PostgreSQL    map[string]postgreSQLNotify    `json:"postgresql"`
	Kafka         map[string]kafkaNotify         `json:"kafka"`
}

// configV7 server configuration version '7'.
type serverConfigV7 struct {
	Version string `json:"version"`

	// S3 API configuration.
	Credential credential `json:"credential"`
	Region     string     `json:"region"`

	// Additional error logging configuration.
	Logger loggerV6 `json:"logger"`

	// Notification queue configuration.
	Notify notifierV1 `json:"notify"`

	// Read Write mutex.
	rwMutex *sync.RWMutex
}

// loadConfigV7 load config version '7'.
func loadConfigV7() (*serverConfigV7, error) {
	configFile, err := getConfigFile()
	if err != nil {
		return nil, err
	}
	if _, err = os.Stat(configFile); err != nil {
		return nil, err
	}
	c := &serverConfigV7{}
	c.Version = "7"
	qc, err := quick.New(c)
	if err != nil {
		return nil, err
	}
	if err := qc.Load(configFile); err != nil {
		return nil, err
	}
	return c, nil
}

// serverConfigV8 server configuration version '8'. Adds NATS notifier
// configuration.
type serverConfigV8 struct {
	Version string `json:"version"`

	// S3 API configuration.
	Credential credential `json:"credential"`
	Region     string     `json:"region"`

	// Additional error logging configuration.
	Logger loggerV6 `json:"logger"`

	// Notification queue configuration.
	Notify notifierV1 `json:"notify"`

	// Read Write mutex.
	rwMutex *sync.RWMutex
}

// loadConfigV8 load config version '8'.
func loadConfigV8() (*serverConfigV8, error) {
	configFile, err := getConfigFile()
	if err != nil {
		return nil, err
	}
	if _, err = os.Stat(configFile); err != nil {
		return nil, err
	}
	c := &serverConfigV8{}
	c.Version = "8"
	qc, err := quick.New(c)
	if err != nil {
		return nil, err
	}
	if err := qc.Load(configFile); err != nil {
		return nil, err
	}
	return c, nil
}

// serverConfigV9 server configuration version '9'. Adds PostgreSQL
// notifier configuration.
type serverConfigV9 struct {
	Version string `json:"version"`

	// S3 API configuration.
	Credential credential `json:"credential"`
	Region     string     `json:"region"`

	// Additional error logging configuration.
	Logger loggerV6 `json:"logger"`

	// Notification queue configuration.
	Notify notifierV1 `json:"notify"`

	// Read Write mutex.
	rwMutex *sync.RWMutex
}

func loadConfigV9() (*serverConfigV9, error) {
	configFile, err := getConfigFile()
	if err != nil {
		return nil, err
	}
	if _, err = os.Stat(configFile); err != nil {
		return nil, err
	}
	srvCfg := &serverConfigV9{}
	srvCfg.Version = "9"
	srvCfg.rwMutex = &sync.RWMutex{}
	qc, err := quick.New(srvCfg)
	if err != nil {
		return nil, err
	}
	if err := qc.Load(configFile); err != nil {
		return nil, err
	}
	return srvCfg, nil
}

// serverConfigV10 server configuration version '10' which is like
// version '9' except it drops support of syslog config, and makes the
// RWMutex global (so it does not exist in this struct).
type serverConfigV10 struct {
	Version string `json:"version"`

	// S3 API configuration.
	Credential credential `json:"credential"`
	Region     string     `json:"region"`

	// Additional error logging configuration.
	Logger logger `json:"logger"`

	// Notification queue configuration.
	Notify notifierV1 `json:"notify"`
}

func loadConfigV10() (*serverConfigV10, error) {
	configFile, err := getConfigFile()
	if err != nil {
		return nil, err
	}
	if _, err = os.Stat(configFile); err != nil {
		return nil, err
	}
	srvCfg := &serverConfigV10{}
	srvCfg.Version = "10"
	qc, err := quick.New(srvCfg)
	if err != nil {
		return nil, err
	}
	if err := qc.Load(configFile); err != nil {
		return nil, err
	}
	return srvCfg, nil
}

// natsNotifyV1 - structure was valid until config V 11
type natsNotifyV1 struct {
	Enable       bool   `json:"enable"`
	Address      string `json:"address"`
	Subject      string `json:"subject"`
	Username     string `json:"username"`
	Password     string `json:"password"`
	Token        string `json:"token"`
	Secure       bool   `json:"secure"`
	PingInterval int64  `json:"pingInterval"`
}

// serverConfigV11 server configuration version '11' which is like
// version '10' except it adds support for Kafka notifications.
type serverConfigV11 struct {
	Version string `json:"version"`

	// S3 API configuration.
	Credential credential `json:"credential"`
	Region     string     `json:"region"`

	// Additional error logging configuration.
	Logger logger `json:"logger"`

	// Notification queue configuration.
	Notify notifierV1 `json:"notify"`
}

func loadConfigV11() (*serverConfigV11, error) {
	configFile, err := getConfigFile()
	if err != nil {
		return nil, err
	}
	if _, err = os.Stat(configFile); err != nil {
		return nil, err
	}
	srvCfg := &serverConfigV11{}
	srvCfg.Version = "11"
	qc, err := quick.New(srvCfg)
	if err != nil {
		return nil, err
	}
	if err := qc.Load(configFile); err != nil {
		return nil, err
	}
	return srvCfg, nil
}

// serverConfigV12 server configuration version '12' which is like
// version '11' except it adds support for NATS streaming notifications.
type serverConfigV12 struct {
	Version string `json:"version"`

	// S3 API configuration.
	Credential credential `json:"credential"`
	Region     string     `json:"region"`

	// Additional error logging configuration.
	Logger logger `json:"logger"`

	// Notification queue configuration.
	Notify notifierV2 `json:"notify"`
}

func loadConfigV12() (*serverConfigV12, error) {
	configFile, err := getConfigFile()
	if err != nil {
		return nil, err
	}
	if _, err = os.Stat(configFile); err != nil {
		return nil, err
	}
	srvCfg := &serverConfigV12{}
	srvCfg.Version = "12"
	qc, err := quick.New(srvCfg)
	if err != nil {
		return nil, err
	}
	if err := qc.Load(configFile); err != nil {
		return nil, err
	}
	return srvCfg, nil
}
