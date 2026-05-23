package logger

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/aws/aws-lambda-go/lambdacontext"
)

type loggerLevel string

const (
	loggerLevelError loggerLevel = "error"
	loggerLevelInfo  loggerLevel = "success"
)

var (
	commitID = "unknown"
)

type logger struct {
	Event   string        `json:"event"`
	Level   loggerLevel   `json:"level"`
	Context loggerGlobal  `json:"context"`
	Data    loggerDataMap `json:"data"`
}

type loggerGlobal struct {
	CommitID     string `json:"commit_id"`
	FunctionName string `json:"function_name"`
	RequestID    string `json:"request_id"`
}

type loggerDataMap map[string]map[string]any

type LoggerData interface {
	GetKey() string
	GetData() map[string]any
}

func buildLoggerData(data []LoggerData) loggerDataMap {
	objData := make(loggerDataMap, len(data))

	for _, data := range data {
		objData[data.GetKey()] = data.GetData()
	}

	return objData
}

func Notify(ctx context.Context, event string, level loggerLevel, data ...LoggerData) {
	lc, _ := lambdacontext.FromContext(ctx)

	log, err := json.Marshal(logger{
		Level: level,
		Event: event,
		Context: loggerGlobal{
			CommitID:     commitID,
			FunctionName: lambdacontext.FunctionName,
			RequestID:    lc.AwsRequestID,
		},
		Data: buildLoggerData(data),
	})

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(log))
}

func Info(ctx context.Context, event string, data ...LoggerData) {
	Notify(ctx, event, loggerLevelInfo, data...)
}

func Error(ctx context.Context, event string, err error, data ...LoggerData) {
	Notify(ctx, event, loggerLevelError, append(data, ObjError(err))...)
}

func Auto(ctx context.Context, event string, err error, data ...LoggerData) {
	if err != nil {
		Error(ctx, event, err, data...)
	} else {
		Info(ctx, event, data...)
	}
}
