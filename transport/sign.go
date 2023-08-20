package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
)

func VerifySignature(c *fiber.Ctx) bool {
	sign := c.Get("X-Hub-Signature-256")
	body := c.Body()

	code := Sha2Encode(body, viper.GetString("secret"))
	return fmt.Sprintf("sha256=%s", code) == sign
}

func Sha2Encode(data []byte, secret string) string {
	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write(data)
	return hex.EncodeToString(mac.Sum(nil))
}
