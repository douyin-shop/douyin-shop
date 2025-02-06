// Package utils @Author Adrian.Wang 2025/1/25 23:53:00
package utils

import "fmt"

func GenerateTokenKey(userId int32) string {
	return fmt.Sprintf("token:%d", userId)
}
