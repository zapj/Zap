package str

import (
	"math/rand"
	"regexp"
	"strings"
	"time"
)

const uCharactersPool = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var charactersPool = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()_+"
var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")

// 生成无符号字符串
func RandUString(length int) string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	b := make([]byte, length)
	for i := range b {
		b[i] = uCharactersPool[r.Intn(len(uCharactersPool))]
	}
	return string(b)
}

func RandString(length int) string {

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	b := make([]byte, length)
	b[0] = charactersPool[r.Intn(len(uCharactersPool))]
	for i := 1; i < length; i++ {
		b[i] = charactersPool[r.Intn(len(charactersPool))]
	}
	return string(b)
}

func ToSnakeCase(str string) string {
	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}
