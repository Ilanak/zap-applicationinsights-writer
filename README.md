# zap-applicationinsights-writer

Writes log messages from go.uber.org/zap to application insights as traces.</br>

A complete example for setting up zap with application insights integreation:

    package main

    import (
        "time"

        "go.uber.org/zap"

        "github.com/ilanak/zap-applicationinsights-writer"
    )

    var baseLogger *zap.Logger

    func init() {
        core, fieldsOption, _ := zapappinsigths.NewAppInsightsCore(zapappinsigths.Config{
            InstrumentationKey: "Enter your Ikey",
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

## Notes

* currently only custom dimension supported for traces (custom metrics are not supported)
