package helpers

import (
	"math/rand"
	"time"

	"github.com/axrav/SysAnalytics/backend/db"
)

const charset = "abcdefghijklmnopqrstuvwxyz" +
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var seededRand *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

func GenerateServerToken() string {
	b := make([]byte, 10)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return "SYSTO-" + string(b)
}

func SaveServerToken(ip string, token string) bool {
	if err := db.RedisClient.Set(db.Ctx, ip, token, 0); err != nil {
		return false
	}
	return true
}
