package main

import (
	"fmt"
	"github.com/nacos-group/nacos-sdk-go/v2/clients"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
	"io/ioutil"
)

/*
   @Auth: menah3m
   @Desc:
*/

func main() {
	// 设置server配置项
	sc := []constant.ServerConfig{
		*constant.NewServerConfig("192.168.108.210", 8848, constant.WithContextPath("/path")),
	}

	// 设置client配置项
	cc := *constant.NewClientConfig(constant.WithNamespaceId(""), constant.WithTimeoutMs(5000), constant.WithNotLoadCacheAtStart(true), constant.WithLogDir("/tmp/nacos/log"), constant.WithCacheDir("/tmp/nacos/cache"),
		constant.WithLogLevel("debug"),
	)

	// 创建client实例
	client, err := clients.NewConfigClient(vo.NacosClientParam{
		ClientConfig:  &cc,
		ServerConfigs: sc,
	})

	if err != nil {
		fmt.Println(err)
	}

	// 获取配置内容
	buf, err := ioutil.ReadFile("./config.yaml")
	if err != nil {
		fmt.Println(err)
	}
	ct := string(buf)

	// 添加配置
	_, err = client.PublishConfig(vo.ConfigParam{
		DataId:  "test-data2",
		Group:   "",
		Content: ct,
	})

	if err != nil {
		fmt.Println("publishing config err :%v", err)
	}

	// 删除配置
	// success, err := client.DeleteConfig(vo.ConfigParam{
	//	DataId: "test-data2",
	//	Group:  "",
	// })
	// if err != nil {
	//	panic(err)
	// }
	// if success {
	//	fmt.Println("delete success")
	// }

}
