package child_workflow

import (
	"go.temporal.io/sdk/workflow"
	"time"
)

// @@@SNIPSTART samples-go-child-workflow-example-child-workflow-definition
// SampleChildWorkflow is a Workflow Definition
func SampleChildWorkflow(ctx workflow.Context, name string) (string, error) {
	logger := workflow.GetLogger(ctx)
	greeting := "Hello " + name + "!"
	logger.Info("Child workflow execution: " + greeting)
	workflow.Sleep(ctx, 15*time.Second)
	return greeting, nil
}

// @@@SNIPEND
