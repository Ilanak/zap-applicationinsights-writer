package main

import (
	"time"

	"github.com/ilanak/zap-applicationinsights-writer"
	"go.uber.org/zap"
)

var baseLogger *zap.Logger

func init() {
	core, fieldsOption, _ := zapappinsigths.NewAppInsightsCore(zapappinsigths.Config{
		InstrumentationKey: "enter your iKey",
		MaxBatchSize:       10,              // optional
		MaxBatchInterval:   time.Second * 5, // optional
	})
	baseLogger = zap.New(core, fieldsOption)
	defer baseLogger.Sync()

}

func main() {
	for i := 0; i < 10; i++ {
		baseLogger.Info("new trace", zap.Int("MessageId", i), zap.String("source", "main func"))
	}
	time.Sleep(5000 * time.Millisecond)
}
