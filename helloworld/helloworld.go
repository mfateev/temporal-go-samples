package helloworld

import (
	"context"
	"time"

	"go.temporal.io/temporal/activity"
	"go.temporal.io/temporal/workflow"
	"go.uber.org/zap"
)

// Workflow is a Hello World workflow definition.
func Workflow(ctx workflow.Context, name string) (*Customer, error) {
	ao := workflow.ActivityOptions{
		ScheduleToStartTimeout: time.Minute,
		StartToCloseTimeout:    time.Minute,
	}
	ctx = workflow.WithActivityOptions(ctx, ao)

	logger := workflow.GetLogger(ctx)
	logger.Info("HelloWorld workflow started", zap.String("Name", name))

	var result Customer
	err := workflow.ExecuteActivity(ctx, Activity, name).Get(ctx, &result)
	if err != nil {
		logger.Error("Activity failed.", zap.Error(err))
		return nil, err
	}

	logger.Info("HelloWorld workflow completed.", zap.String("result", result.Name))

	return &result, nil
}

type Customer struct {
	Name string
}

func Activity(ctx context.Context, name string) (*Customer, error) {
	logger := activity.GetLogger(ctx)
	logger.Info("Activity", zap.String("Name", name))
	return &Customer{Name: "Hello " + name + "!"}, nil
}
