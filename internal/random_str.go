package internal

import (
	"math/rand"
	"time"
)

//RandomStr 随机生成字符串
func RandomStr(length int) string {
	str := `0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ`
	bts := []byte(str)
	var result []byte
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < length; i++ {
		result = append(result, bts[r.Intn(len(bts))])
	}
	return string(result)
}
