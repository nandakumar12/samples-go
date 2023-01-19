// @@@SNIPSTART go-samples-yourapp-your-activity-definition
package yourapp

import (
	"context"

	"go.temporal.io/sdk/activity"
)

// YourActivityParam is the struct passed to your Activity.
// Use a struct so that your function signature remains compatible if fields change.
type YourActivityParam struct {
	ActivityParamX string
	ActivityParamY int
}

// YourActivityResultObject is the struct returned from your Activity.
// Use a struct so that you can return multiple values of different types.
// Additionally, your function signature remains compatible if the fields change.
type YourActivityResultObject struct {
	ResultFieldX string
	ResultFieldY int
}

// YourActivityObject is the struct that maintains shared state across Activities.
// If the Worker crashes this Activity object loses its state.
type YourActivityObject struct {
	SharedMessageState *string
	SharedCounterState *int
}

// YourActivityDefinition is your custom Activity Definition.
// An Activity Definiton is an exportable function.
func (a *YourActivityObject) YourActivityDefinition(ctx context.Context, param YourActivityParam) (YourActivityResultObject, error) {
	// Use Acivities for computations or calling external APIs.
	// This is just an example of appending to text and incrementing a counter.
	message := param.ActivityParamX + " World!"
	counter := param.ActivityParamY + 1
	a.SharedMessageState = &message
	a.SharedCounterState = &counter
	result := YourActivityResultObject{
		ResultFieldX: *a.SharedMessageState,
		ResultFieldY: *a.SharedCounterState,
	}
	// Return the results back to the Workflow Execution.
	// The results persist within the Event History of the Workflow Execution.
	return result, nil
}

// PrintSharedState is another custom Activity Definition.
func (a *YourActivityObject) PrintSharedSate(ctx context.Context) error {
	logger := activity.GetLogger(ctx)
	logger.Info("The current message is:", *a.SharedMessageState)
	logger.Info("The current counter is:", *a.SharedCounterState)
	return nil
}

// YourSimpleActivityDefinition is a basic Activity Definiton.
func YourSimpleActivityDefinition(ctx context.Context) error {
	return nil
}

// @@@SNIPEND
