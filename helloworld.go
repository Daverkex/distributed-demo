package distributed

import (
	"context"
	"strings"
	"time"

	"go.temporal.io/sdk/activity"
	"go.temporal.io/sdk/workflow"
)

// Workflow is a Hello World workflow definition.
func Workflow(ctx workflow.Context, text string) ([]string, error) {
	ao := workflow.ActivityOptions{
		StartToCloseTimeout: 30 * time.Second,
	}
	ctx = workflow.WithActivityOptions(ctx, ao)

	logger := workflow.GetLogger(ctx)
	logger.Info("HelloWorld workflow started", "text", text)

	var result []string
	err := workflow.ExecuteActivity(ctx, SplitText, text).Get(ctx, &result)
	if err != nil {
		logger.Error("SplitText failed.", "Error", err)
		return nil, err
	}

	logger.Info("HelloWorld workflow completed.", "result", result)

	for _, s := range result {
		err := workflow.ExecuteActivity(ctx, PrintText, s).Get(ctx, nil)
		if err != nil {
			logger.Error("PrintText failed.", "Error", err)
			return nil, err
		}
	}

	return result, nil
}

func SplitText(ctx context.Context, text string) ([]string, error) {
	logger := activity.GetLogger(ctx)
	s := strings.Split(text, " ")
	logger.Info("SplitText", "text", s)
	return s, nil
}

func PrintText(ctx context.Context, text string) (string, error) {
	logger := activity.GetLogger(ctx)
	// Simulate a long-running activity
	time.Sleep(10 * time.Second)
	logger.Info("PrintText", "text", text)
	return text, nil
}
