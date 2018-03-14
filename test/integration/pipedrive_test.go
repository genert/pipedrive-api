package integration

import (
	"math/rand"
	"os"
	"time"

	"github.com/genert/pipedrive-api/pipedrive"
)

var (
	client *pipedrive.Client

	// Random string generator
	src = rand.NewSource(time.Now().UnixNano())
)

// Random string generator
const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

func RandomString(n int) string {
	b := make([]byte, n)

	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return string(b)
}

func init() {
	token := os.Getenv("PIPEDRIVE_API_TOKEN")

	if token == "" {
		print("No API key found. Integration tests won't run!\n\n")
		os.Exit(1)
	} else {
		config := &pipedrive.Config{
			APIKey: token,
		}

		client = pipedrive.NewClient(config)
	}
}
