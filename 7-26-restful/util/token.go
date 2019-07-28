package util

import (
	"7-26-restful/encryption"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"
)

var (
	//jwt头部
	header = map[string]string{
		"typ": "JWT",
		"alg": "ASE",
	}
	//jwt载体
	payload = map[string]interface{}{
		"iss": "hangiangai", //签发者
		"sub": "all",        //登录用户名
		"aud": "test",       //登录端
		"exp": 0,            //过期时间(30分钟)
	}
)

//生成token
func CreateToken(username string) (string, string) {
	//将头部转化为json
	header_json, err := json.Marshal(header)
	//对头部使用base64编码
	header_base64 := base64.StdEncoding.EncodeToString(header_json)
	//将载体转化为json
	payload["sub"] = username
	payload["exp"] = time.Now().Unix() + 30*60
	payload_json, err := json.Marshal(payload)
	ErrProcess(err, 153)
	//对载体进行base64编码
	payload_base64 := base64.StdEncoding.EncodeToString(payload_json)
	//生成秘钥
	secret_key := base64.StdEncoding.EncodeToString([]byte(strconv.FormatInt(time.Now().UnixNano(), 10)))
	aeskey := []byte(secret_key[0:16])
	//以.连接头部和载体
	encodedString := header_base64 + "." + payload_base64
	pass := []byte(encodedString)
	xpass, err := encryption.AesEncrypt(pass, aeskey)
	ErrProcess(err, 168)
	//base64加密
	pass64 := base64.StdEncoding.EncodeToString(xpass)
	return pass64, secret_key
}

//获取token信息
func Get(token string, secret_key string) (map[string]string, map[string]interface{}) {
	fmt.Println(token, secret_key)
	//base64解密
	base64_, err := base64.StdEncoding.DecodeString(token)
	ErrProcess(err, 53)
	//ase解密
	aes, err := encryption.AesDecrypt(base64_, []byte(secret_key[0:16]))
	ErrProcess(err, 57)
	//获得base64加密的头部和载体
	get_jwt := strings.Split(string(aes), ".")
	//对头部进行base64解密
	jwt_header, err := base64.StdEncoding.DecodeString(get_jwt[0])
	get_header := make(map[string]string)
	ErrProcess(err, 67)
	ErrProcess(json.Unmarshal([]byte(jwt_header), &get_header), 63)
	//对载体进行base64为解密
	jwt_payload, err := base64.StdEncoding.DecodeString(get_jwt[1])
	ErrProcess(err, 70)
	get_payload := make(map[string]interface{})
	ErrProcess(json.Unmarshal([]byte(jwt_payload), &get_payload), 63)
	return get_header, get_payload
}
