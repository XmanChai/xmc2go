package gen

import (
	"crypto/md5"
	"encoding/binary"
	"encoding/hex"
	"time"
)

const (
	HASH_MIN    int64 = 16
	HASH_NORMAL int64 = 32
	HASH_LARGE  int64 = 64
	HASH_MAX    int64 = 128
)

func GetToken(bit int64) string {
	//纳秒级别的
	t := time.Now().UnixNano()
	h := md5.New()
	b := make([]byte, bit)

	binary.BigEndian.PutUint64(b, uint64(t))
	h.Write([]byte(b))

	return hex.EncodeToString(h.Sum(nil))
}
