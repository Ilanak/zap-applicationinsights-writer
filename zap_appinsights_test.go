package zapappinsigths

import (
	"encoding/json"
	"testing"

	"github.com/Microsoft/ApplicationInsights-Go/appinsights"
	"github.com/Microsoft/ApplicationInsights-Go/appinsights/contracts"
)

func TestWriteIntegration(t *testing.T) {
	appInsightsConfig := AppInsightsConfig{
		client:  appinsights.NewTelemetryClient("00000000-0000-0000-0000-000000000000"),
		filters: make(map[string]func(interface{}) interface{}),
	}

	message := "hello world"
	msg := "{\"source\": \"test\", \"msg\": \"" + message + "\", \"level\":\"Information\"}"

	n, err := appInsightsConfig.Write([]byte(msg))
	if n != len(message) {
		t.Errorf("Expected for n value %v but got value %v", len(msg), n)
	}
	if err != nil {
		t.Errorf("Expected no error but got %v", err)
	}
}

func TestBuildTrace(t *testing.T) {
	message := "hello world"
	level := contracts.Information
	source := "test"
	msg := "{\"source\": \"" + source + "\", \"msg\": \"" + message + "\", \"level\":\"" + level.String() + "\"}"
	msgbyte := []byte(msg)
	var data map[string]interface{}
	json.Unmarshal(msgbyte, &data)
	trace := BuildTrace(data)
	if trace.Message != message {
		t.Errorf("Expected for value %v but got value %v", message, trace.Message)
	}
	if trace.SeverityLevel != level {
		t.Errorf("Expected for value %v but got value %v", level, trace.SeverityLevel)
	}
	if trace.BaseTelemetry.Properties["source"] != source {
		t.Errorf("Expected for value %v but got value %v", source, trace.Properties["source"])
	}
}
