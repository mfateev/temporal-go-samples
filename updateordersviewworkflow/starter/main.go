package main

import (
	"context"
	"go.temporal.io/api/enums/v1"
	"log"

	"go.temporal.io/sdk/client"

	"github.com/temporalio/samples-go/updateordersviewworkflow"
)

func main() {
	// The client is a heavyweight object that should be created once per process.
	c, err := client.NewClient(client.Options{})
	if err != nil {
		log.Fatalln("Unable to create client", err)
	}
	defer c.Close()

	workflowOptions := client.StartWorkflowOptions{
		ID:                                       "1",
		TaskQueue:                                "ordersViewQueue",
		WorkflowExecutionErrorWhenAlreadyStarted: false,
		CronSchedule:                             "* * * * *",
		WorkflowIDReusePolicy:                    enums.WORKFLOW_ID_REUSE_POLICY_ALLOW_DUPLICATE,
	}

	we, err := c.ExecuteWorkflow(context.Background(), workflowOptions, updateordersviewworkflow.UpdateOrdersViewWorkflow,
		&updateordersviewworkflow.ActivitiesOptions{}, "ViewName1")
	if err != nil {
		log.Fatalln("Unable to execute workflow", err)
	}

	log.Println("Started workflow", "WorkflowID", we.GetID(), "RunID", we.GetRunID())

}
