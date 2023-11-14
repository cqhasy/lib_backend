package middlewares

import (
	"AILN/app/common"
	"go.uber.org/zap"
	"time"

	"github.com/gin-gonic/gin"
)

// zap接管gin日志(有需要自行加载)
func UseZapLogger(r *gin.Engine) {
	logger := common.LOG
	r.Use(func(ctx *gin.Context) {
		start := time.Now()
		ctx.Next()
		cost := time.Since(start)
		logger.Info(ctx.Request.URL.Path,
			zap.Int("status", ctx.Writer.Status()),
			zap.String("method", ctx.Request.Method),
			zap.String("path", ctx.Request.URL.Path),
			zap.String("query", ctx.Request.URL.RawQuery),
			zap.String("ip", ctx.ClientIP()),
			zap.String("user-agent", ctx.Request.UserAgent()),
			zap.String("errors", ctx.Errors.ByType(gin.ErrorTypePrivate).String()),
			zap.Duration("cost", cost),
		)
	})
}
