package main

import (
	"github.com/temporalio/samples-go/updateordersviewworkflow"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
	"log"
)

func main() {
	// The client and worker are heavyweight objects that should be created once per process.
	c, err := client.NewClient(client.Options{})
	if err != nil {
		log.Fatalln("Unable to create client", err)
	}
	defer c.Close()

	w := worker.New(c, "ordersViewQueue", worker.Options{})

	w.RegisterWorkflow(updateordersviewworkflow.UpdateOrdersViewWorkflow)
	w.RegisterActivity(&updateordersviewworkflow.UpdateViewActivities{})

	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("Unable to start worker", err)
	}
}
