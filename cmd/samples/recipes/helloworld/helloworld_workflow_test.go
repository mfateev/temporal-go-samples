package main

import (
	"github.com/stretchr/testify/mock"
	"go.temporal.io/temporal/worker"
	"testing"

	"github.com/stretchr/testify/require"
	"go.temporal.io/temporal/testsuite"
)

func Test_Workflow(t *testing.T) {
	testSuite := &testsuite.WorkflowTestSuite{}
	env := testSuite.NewTestWorkflowEnvironment()
	env.RegisterActivity(helloWorldActivity)

	// Mock activity
	env.OnActivity(helloWorldActivity, mock.Anything, "World").Return("Test Hello World!", nil)

	env.ExecuteWorkflow(HelloWorldWorkflow, "World")

	require.True(t, env.IsWorkflowCompleted())
	require.NoError(t, env.GetWorkflowError())
	var result string
	require.NoError(t, env.GetWorkflowResult(&result))
	require.Equal(t, "Test Hello World!", result)
}

// This replay test is the recommended way to make sure changing workflow code is backward compatible
// without non-deterministic errors.
// "helloworld.json" can be downloaded from cadence CLI:
//      cadence --do <domain> wf show -w <workflowID> --output_filename <output>.json
// Or from Cadence Web UI.
func TestReplayWorkflowHistoryFromFile(t *testing.T) {
	r := worker.NewWorkflowReplayer()
	r.RegisterWorkflow(HelloWorldWorkflow)
	err := r.ReplayWorkflowHistoryFromJSONFile(nil, "helloworld.json")
	require.NoError(t, err)
}
