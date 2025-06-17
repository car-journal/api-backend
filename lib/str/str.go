package str

import (
	"math/rand"
	"strconv"
	"time"
)

const DefaultCharsetAlphaNumeric = "abcdefghijklmnopqrstuvwxyz" + "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789" + "0123456789"
const DefaultCharset = "abcdefghijklmnopqrstuvwxyz" + "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func Random(length int, charset string) string {
	seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))
	bt := make([]byte, length)
	for i := range bt {
		bt[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(bt)
}

func StringToInt64(str string) (int64, error) {
	val, err := strconv.ParseInt(str, 10, 64)
	if nil != err {
		return 0, err
	}
	return val, nil
}
