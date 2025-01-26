// @Author Adrian.Wang 2025/1/21 14:46:00
package nacos

import (
	"github.com/cloudwego/kitex/pkg/discovery"
	register "github.com/cloudwego/kitex/pkg/registry"
	"github.com/douyin-shop/douyin-shop/common/conf"
	"github.com/kitex-contrib/registry-nacos/v2/registry"
	"github.com/kitex-contrib/registry-nacos/v2/resolver"
	"github.com/nacos-group/nacos-sdk-go/v2/clients"
	"github.com/nacos-group/nacos-sdk-go/v2/clients/naming_client"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
	"log"
)

func GetNacosConfigClient() (iClient naming_client.INamingClient, err error) {

	nacos := conf.GetConf().Nacos

	sc := []constant.ServerConfig{
		*constant.NewServerConfig(nacos.Address, nacos.Port),
	}

	cc := constant.ClientConfig{
		NamespaceId:         nacos.Namespace,
		TimeoutMs:           nacos.TimeoutMs,
		NotLoadCacheAtStart: nacos.NotLoadCacheAtStart,
		LogDir:              nacos.LogDir,
		CacheDir:            nacos.CacheDir,
		LogLevel:            nacos.LogLevel,
		Username:            nacos.Username,
		Password:            nacos.Password,
	}

	client, err := clients.NewNamingClient(
		vo.NacosClientParam{
			ClientConfig:  &cc,
			ServerConfigs: sc,
		},
	)

	if err != nil {
		log.Fatalf("create nacos client error_code: %v", err)
		return nil, err
	}

	return client, nil

}

func GetNacosRegistry() register.Registry {

	// 获取nacos配置
	client, err := GetNacosConfigClient()

	if err != nil {
		log.Fatalf("get nacos client error_code: %v", err)
		return nil
	}

	return registry.NewNacosRegistry(client)

}

// GetNacosResolver 获取nacos resolver
func GetNacosResolver() discovery.Resolver {

	// 获取nacos配置
	client, err := GetNacosConfigClient()

	if err != nil {
		log.Fatalf("get nacos client error_code: %v", err)
		return nil
	}

	return resolver.NewNacosResolver(client)

}
