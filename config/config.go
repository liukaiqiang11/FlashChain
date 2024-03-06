package config

import (
	"fchain/common/signer"
	"fmt"
	"github.com/spf13/viper"
)

// GetClientConf 获取客户端的信息配置
func GetClientConf() (*signer.Config, error) {
	viper.SetDefault("client.MSPID", "*.lkq.com")
	viper.SetDefault("client.IdentityPath", "cert/client.pem")
	viper.SetDefault("client.KeyPath", "cert/client.key")

	viper.AddConfigPath(configFilePath)
	viper.SetConfigName("config")
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	return &signer.Config{
		MSPID:        viper.GetString("client.MSPID"),
		IdentityPath: ProjectPath + viper.GetString("client.IdentityPath"),
		KeyPath:      ProjectPath + viper.GetString("client.KeyPath"),
	}, nil
}

// GetPeerConf 获取Peer节点的信息配置
func GetPeerConf() (*signer.Config, error) {
	viper.SetDefault("peer.MSPID", "*.lkq.com")
	viper.SetDefault("peer.IdentityPath", "cert/peer.pem")
	viper.SetDefault("peer.KeyPath", "cert/peer.key")

	viper.AddConfigPath(configFilePath)
	viper.SetConfigName("config")
	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	return &signer.Config{
		MSPID:        viper.GetString("peer.MSPID"),
		IdentityPath: ProjectPath + viper.GetString("peer.IdentityPath"),
		KeyPath:      ProjectPath + viper.GetString("peer.KeyPath"),
	}, nil
}

// GetOrderConf 获取Order节点的配置信息
func GetOrderConf() (*signer.Config, error) {
	viper.SetDefault("order.MSPID", "*.lkq.com")
	viper.SetDefault("order.IdentityPath", "cert/order.pem")
	viper.SetDefault("order.KeyPath", "cert/order.key")

	viper.AddConfigPath(configFilePath)
	viper.SetConfigName("config")
	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	return &signer.Config{
		MSPID:        viper.GetString("order.MSPID"),
		IdentityPath: ProjectPath + viper.GetString("order.IdentityPath"),
		KeyPath:      ProjectPath + viper.GetString("order.KeyPath"),
	}, nil
}

func GetOrganizations() ([]Organization, error) {

	viper.AddConfigPath(configFilePath)
	viper.SetConfigName("config")
	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}
	// 创建一个切片来存储组织数据
	var organizations []Organization

	// 使用Viper将YAML数据映射到结构体切片
	if err := viper.UnmarshalKey("organizations", &organizations); err != nil {
		return nil, fmt.Errorf("Error unmarshaling data: %v\n", err)
	}

	return organizations, nil
}
