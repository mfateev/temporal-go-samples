package child_workflow

import (
	"context"
	"errors"
	"go.temporal.io/sdk/temporal"
	"go.temporal.io/sdk/workflow"
	"time"
)

func ChildActivity(ctx context.Context) error {
	// This activity does nothing
	return temporal.NewNonRetryableApplicationError("simulated activity error", "Simulated", nil)
}

// @@@SNIPSTART samples-go-child-workflow-example-child-workflow-definition
// SampleChildWorkflow is a Workflow Definition
func SampleChildWorkflow(ctx workflow.Context, name string) (string, error) {
	logger := workflow.GetLogger(ctx)
	greeting := "Hello " + name + "!"
	logger.Info("Child workflow execution: " + greeting)
	err := workflow.ExecuteActivity(workflow.WithActivityOptions(ctx, workflow.ActivityOptions{StartToCloseTimeout: time.Minute}), ChildActivity).Get(ctx, nil)
	if err != nil {
		var applicationErr *temporal.ApplicationError
		if errors.As(err, &applicationErr) {
			if applicationErr.NonRetryable() {
				return "", temporal.NewNonRetryableApplicationError("activity failed", "NonRetryableFailure", err)
			}
		}
	}
	return greeting, nil
}

// @@@SNIPEND
