#https单向认证

创建CA:
    生成CA私钥:openssl genrsa -out ca.key 2048
    自签发证书(主要参数:-subj "/CN=fightcoder.com"):openssl req -x509 -new -nodes -key ca.key -days 365 -out ca.crt

由CA签发服务端证书:
    生成私钥:openssl genrsa -out server.key 2048
    生成证书请求:openssl req -new -key server.key -out server.csr
    CA签发证书(主要参数:-subj "/CN=localhost"):openssl x509 -req -in server.csr -CA ca.crt -CAkey ca.key -CAcreateserial -out server.crt -days 365