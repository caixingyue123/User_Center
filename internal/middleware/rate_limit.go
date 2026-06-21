// internal/middleware/rate_limit.go
package middleware

import (
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

type rateLimitRecord struct {
	count      int
	windowFrom time.Time
}

func RateLimit(limit int, window time.Duration) gin.HandlerFunc {
	var mu sync.Mutex
	records := make(map[string]*rateLimitRecord)

	return func(ctx *gin.Context) {
		ip := ctx.ClientIP()
		now := time.Now()

		mu.Lock()

		record, exists := records[ip]
		if !exists || now.Sub(record.windowFrom) >= window {
			records[ip] = &rateLimitRecord{
				count:      1,
				windowFrom: now,
			}
			mu.Unlock()
			ctx.Next()
			return
		}

		if record.count >= limit {
			retryAfter := int(window.Seconds() - now.Sub(record.windowFrom).Seconds())
			if retryAfter < 1 {
				retryAfter = 1
			}

			mu.Unlock()

			ctx.Header("Retry-After", strconv.Itoa(retryAfter))
			ctx.JSON(http.StatusTooManyRequests, gin.H{
				"code":    429,
				"message": "请求过于频繁，请稍后再试",
			})
			ctx.Abort()
			return
		}

		record.count++
		mu.Unlock()

		ctx.Next()
	}
}
