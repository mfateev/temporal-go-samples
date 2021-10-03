package updateordersviewworkflow

import (
	"context"
	"time"

	"go.temporal.io/sdk/workflow"
)

type ActivitiesOptions struct {

}

func (o *ActivitiesOptions) NewActivityOptionsForOrdersViewQueue() workflow.ActivityOptions {
	return workflow.ActivityOptions{TaskQueue:  "ordersViewQueue", StartToCloseTimeout: time.Minute}
}

type NextCronStart struct {

}

func UpdateOrdersViewWorkflow(ctx workflow.Context, defaultOptions *ActivitiesOptions, viewName string) (*NextCronStart, error) {

	var a *UpdateViewActivities
	for i := 0; i < 1000; i++ {
		opts := defaultOptions.NewActivityOptionsForOrdersViewQueue()

		ctx = workflow.WithActivityOptions(ctx, opts)
		if err := workflow.ExecuteActivity(ctx, a.ApplyEventActivity, viewName, nil).Get(ctx, nil); err != nil {
			return nil, err
		}
	}
	return nil, nil
}

type UpdateViewActivities struct {

}

type EventEntry struct {

}

func (a *UpdateViewActivities) ApplyEventActivity(ctx context.Context, viewName string, eventEntry *EventEntry) error {
	return nil
}
