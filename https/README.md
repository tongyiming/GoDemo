生成私钥:openssl genrsa -out server.key 2048

生成自签名证书(没有根证书,验证不通过):openssl req -new -x509 -key server.key -out server.crt -days 365

获取页面信息(-k:允许不使用证书到SSL站点):curl -k https://localhost:8080