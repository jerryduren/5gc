package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"net/http"
)

type myhandler struct {
}

func (h *myhandler) ServeHTTP(w http.ResponseWriter,
	r *http.Request) {
	fmt.Fprintf(w,
		"Hi, This is an example of https service in golang!\n")
}

func main() {
	pool := x509.NewCertPool()
	caCertPath := "ca.crt" // 设置CA机构自己的数字证书的路径和文件名，这个是所有认证的根证书

	caCrt, err := ioutil.ReadFile(caCertPath)
	if err != nil {
		fmt.Println("ReadFile err:", err)
		return
	}
	pool.AppendCertsFromPEM(caCrt) // 将自己的CA机构的根证书加入到证书池里面

	// 配置服务器采用tls加密，并且要认证客户端的身份
	// server如何认证client的身份：在建立tls的时候要求client提供从CA机构申请的数字证书
	// 然后用CA根证书里面的公钥解密client数字证书里面的签名，把解码后的结果与client证书的摘要比较，一致则合法
	// 所以server需要加载CA的根证书，其实就是读取里面的公钥和算法
	s := &http.Server{
		Addr:    ":8081",
		Handler: &myhandler{},
		TLSConfig: &tls.Config{
			ClientCAs:  pool,
			ClientAuth: tls.RequireAndVerifyClientCert,
		},
	}

	// server采用smf的数字证书
	// 因为是双向认证，所以这个也要设置server的证书，在tls协商的时候需要传给client，以便client认证server
	err = s.ListenAndServeTLS("smf.crt", "smf.key")
	if err != nil {
		fmt.Println("ListenAndServeTLS err:", err)
	}
}
