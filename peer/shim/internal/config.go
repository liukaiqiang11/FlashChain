package internal

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"os"

	"fchain/config"
)

func LoadConfig() (*tls.Config, error) {

	conf := config.ClientConf
	// TLS连接
	//从证书相关文件中读取和解析信息，得到证书公钥、密钥对
	cert, err := tls.LoadX509KeyPair(conf.IdentityPath, conf.KeyPath)
	if err != nil {
		return nil, fmt.Errorf("failed to LoadX509KeyPair: %s", err)
	}
	certPool := x509.NewCertPool()

	ca, err := os.ReadFile(config.CaCrt)
	if err != nil {
		return nil, fmt.Errorf("failed to read private key file: %s", err)
	}
	certPool.AppendCertsFromPEM(ca)

	return &tls.Config{
		Certificates: []tls.Certificate{cert}, //客户端证书
		RootCAs:      certPool,
	}, nil
}

func LoadTLSConfig() (*tls.Config, error) {

	//获取Peer相关的配置
	conf := config.PeerConf

	// TLS认证
	//从证书相关文件中读取和解析信息，得到证书公钥、密钥对
	cert, err := tls.LoadX509KeyPair(conf.IdentityPath, conf.KeyPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read certificate: %s", err)
	}
	certPool := x509.NewCertPool() //初始化一个CertPool

	ca, err := os.ReadFile(config.CaCrt)
	if err != nil {
		return nil, fmt.Errorf("failed to read CA certificate: %s", err)
	}
	certPool.AppendCertsFromPEM(ca) //解析传入的证书，解析成功会将其加到池子中
	tlscfg := &tls.Config{ //构建基于TLS的TransportCredentials选项
		Certificates: []tls.Certificate{cert},        //服务端证书链，可以有多个
		ClientAuth:   tls.RequireAndVerifyClientCert, //要求必须验证客户端证书
		ClientCAs:    certPool,                       //设置根证书的集合，校验方式使用 ClientAuth 中设定的模式
	}

	return tlscfg, nil
}
