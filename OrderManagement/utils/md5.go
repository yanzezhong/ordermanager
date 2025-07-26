package utils

import (
	"crypto/md5"
	"encoding/hex"
)

func GetMD5(target string) string {
	// 生成店名的 MD5 值
	hash := md5.Sum([]byte(target))
	shopNameMD5 := hex.EncodeToString(hash[:])
	return shopNameMD5
}
