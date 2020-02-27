package main

import (
	"context"
	"fmt"
	"go.temporal.io/temporal/workflow"
	"time"

	"go.temporal.io/temporal-proto/workflowservice"
	"go.temporal.io/temporal/client"
	"go.temporal.io/temporal/worker"
	"google.golang.org/grpc"
)

const localServiceAddress = "127.0.0.1:7233"
const domain = "sample"
const taskList = "HelloWorld"

func HelloWorldWorkflow(ctx workflow.Context, name string) (string, error) {
	ctx = workflow.WithActivityOptions(ctx, workflow.ActivityOptions{
		ScheduleToCloseTimeout: time.Second * 5,
		ScheduleToStartTimeout: time.Second * 5,
		StartToCloseTimeout:    time.Second * 5,
	})
	var result string
	err := workflow.ExecuteActivity(ctx, helloWorldActivity, name).Get(ctx, &result)
	if err != nil {
		return "", fmt.Errorf("failure executing helloWorldActivity: %w", err)
	}
	return result, nil
}

func helloWorldActivity(ctx context.Context, name string) (string, error) {
	return "Hello " + name + "!", nil
}

func main() {
	connection, err := grpc.Dial(localServiceAddress, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	serviceClient := workflowservice.NewWorkflowServiceClient(connection)

	// Worker hosts both workflow and activity code
	worker := worker.New(serviceClient, domain, taskList, worker.Options{})
	worker.RegisterWorkflow(HelloWorldWorkflow)
	worker.RegisterActivity(helloWorldActivity)
	err = worker.Start()
	if err != nil {
		panic(err)
	}

	// Start workflow execution. In the majority of real applications the start is called from a different process.
	tClient := client.NewClient(serviceClient, domain, nil)
	startOptions := client.StartWorkflowOptions{
		ExecutionStartToCloseTimeout: time.Minute,
		TaskList:                     taskList,
	}
	run, err := tClient.ExecuteWorkflow(context.Background(), startOptions, HelloWorldWorkflow, "World")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Started HelloWorld workflow with ID=%v and RunID=%v\n", run.GetID(), run.GetRunID())

	// Block until workflow is completed.
	var result string
	if err = run.Get(context.Background(), &result); err != nil {
		panic(err)
	}
	fmt.Printf("Workflow result: %v\n", result)
}
