package test

import (
    "testing"
    functions"example-mod/functions"
	"time"

    zap "go.uber.org/zap"
	zapcore "go.uber.org/zap/zapcore"

)

func TestAdd(t *testing.T) {
    // Test cases
	loggerConfig := zap.NewProductionConfig()
	loggerConfig.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout(time.RFC3339)
	logger, _ := loggerConfig.Build()
	defer logger.Sync()

    testCases := []struct {
        a, b     int
        expected int
    }{
        {1, 2, 3},
        {0, 0, 0},
        {-1, 1, 0},
        {10, -5, 5},
    }

    for _, tc := range testCases {
        result := functions.Add(tc.a, tc.b)
        if result != tc.expected {
            t.Errorf("Add(%d, %d) = %d; expected %d", tc.a, tc.b, result, tc.expected)
        }
    }
}
