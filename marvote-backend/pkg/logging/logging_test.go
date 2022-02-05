package logging

import (
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"go.uber.org/zap"
	"go.uber.org/zap/zaptest/observer"
)

func init() {

}

type LoggingTestSuite struct {
	suite.Suite
}

func (ts *LoggingTestSuite) TestDebugLog() {
	// Given
	observedZapCore, observedLogs := observer.New(zap.DebugLevel)
	zapLog = zap.New(observedZapCore)

	// When
	Debug("debug log")

	// Then
	require.Equal(ts.T(), 1, observedLogs.Len())
	assert.Equal(ts.T(), "debug log", observedLogs.All()[0].Message, "Must have the same logs")

}

func (ts *LoggingTestSuite) TestDebugfLog() {
	// Given
	observedZapCore, observedLogs := observer.New(zap.DebugLevel)
	zapLog = zap.New(observedZapCore)

	// When
	Debugf("debug log ==> %d", 1)

	// Then
	require.Equal(ts.T(), 1, observedLogs.Len())
	assert.Equal(ts.T(), "debug log ==> 1", observedLogs.All()[0].Message, "Must have the same logs")

}

func (ts *LoggingTestSuite) TestInfoLog() {
	// Given
	observedZapCore, observedLogs := observer.New(zap.InfoLevel)
	zapLog = zap.New(observedZapCore)

	// When
	Info("info log")

	// Then
	require.Equal(ts.T(), 1, observedLogs.Len())
	assert.Equal(ts.T(), "info log", observedLogs.All()[0].Message, "Must have the same logs")

}

func (ts *LoggingTestSuite) TestInfofLog() {
	// Given
	observedZapCore, observedLogs := observer.New(zap.InfoLevel)
	zapLog = zap.New(observedZapCore)

	// When
	Infof("info log ==> %d", 1)

	// Then
	require.Equal(ts.T(), 1, observedLogs.Len())
	assert.Equal(ts.T(), "info log ==> 1", observedLogs.All()[0].Message, "Must have the same logs")

}
func (ts *LoggingTestSuite) TestErrorLog() {
	// Given
	observedZapCore, observedLogs := observer.New(zap.ErrorLevel)
	zapLog = zap.New(observedZapCore)

	// When
	Error("error")

	// Then
	require.Equal(ts.T(), 1, observedLogs.Len())
	assert.Equal(ts.T(), "error", observedLogs.All()[0].Message, "Must have the same logs")

}

func (ts *LoggingTestSuite) TestErrorfLog() {
	// Given
	observedZapCore, observedLogs := observer.New(zap.ErrorLevel)
	zapLog = zap.New(observedZapCore)

	// When
	Errorf("%v", errors.New("error is here"))

	// Then
	require.Equal(ts.T(), 1, observedLogs.Len())
	assert.Equal(ts.T(), "error is here", observedLogs.All()[0].Message, "Must have the same logs")

}

func TestLogging(t *testing.T) {
	suite.Run(t, new(LoggingTestSuite))
}
