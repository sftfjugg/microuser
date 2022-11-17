package config

import (
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"gopkg.in/yaml.v2"
	"log"
)

type Conf struct {
	App App `yaml:"app"`
}

type App struct {
	Name string `yaml:"name"`
	Port int64  `yaml:"port"`
}

var GlobalConf Conf

func InitSetting() {
	clientConfig := constant.ClientConfig{
		NamespaceId:         "9dbcec67-96b5-45c9-8e80-737342fac52d",
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              "D:/alidata/logs/nacos/log",
		CacheDir:            "D:/alidata/logs/nacos/cache",
		RotateTime:          "1h",
		MaxAge:              3,
		LogLevel:            "debug",
	}
	// At least one ServerConfig
	serverConfigs := []constant.ServerConfig{
		{
			IpAddr:      "127.0.0.1",
			ContextPath: "/nacos",
			Port:        8848,
			Scheme:      "http",
		},
	}
	configClient, err := clients.NewConfigClient(
		vo.NacosClientParam{
			ClientConfig:  &clientConfig,
			ServerConfigs: serverConfigs,
		},
	)
	content, err := configClient.GetConfig(vo.ConfigParam{
		DataId: "microuser",
		Group:  "DEFAULT_GROUP"})

	if err != nil {
		log.Fatalf("获取配置信息出错，", err)
	}
	err = MapTo(content)
	if err != nil {
		log.Fatalf("yaml解析配置文件出错，", err)
	}
	// 动态监听
	err = configClient.ListenConfig(vo.ConfigParam{
		DataId: "microuser",
		Group:  "DEFAULT_GROUP",
		OnChange: func(namespace, group, dataId, data string) {
			err = MapTo(data)
			if err != nil {
				log.Printf("监听配置出错，", err)
			}
		},
	})
	if err != nil {
		log.Fatalf("动态监听配置文件出错，", err)
	}
}

func MapTo(content string) (err error) {
	err = yaml.Unmarshal([]byte(content), &GlobalConf)
	return
}
