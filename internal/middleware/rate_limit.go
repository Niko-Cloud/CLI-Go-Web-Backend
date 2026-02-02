package middleware

import (
	"time"

	"github.com/JGLTechnologies/gin-rate-limit"
	"github.com/gin-gonic/gin"
)

func RateLimiter() gin.HandlerFunc {
	store := ratelimit.InMemoryStore(&ratelimit.InMemoryOptions{
		Rate:  time.Second,
		Limit: 5,
	})

	return ratelimit.RateLimiter(store, &ratelimit.Options{
		ErrorHandler: func(c *gin.Context, info ratelimit.Info) {
			c.AbortWithStatusJSON(429, gin.H{
				"code":    "TOO_MANY_REQUESTS",
				"status":  429,
				"message": "Too many requests",
			})
		},
		KeyFunc: func(c *gin.Context) string {
			return c.ClientIP()
		},
	})
}
