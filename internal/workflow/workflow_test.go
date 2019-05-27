package workflow

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"testing"
)

func TestNewBuilder(t *testing.T) {

	logger := zap.Logger{}

	testBuilder := NewBuilder(&logger)

	testCadenceClient,err := testBuilder.BuildCadenceClient()

	//testCadenceClient.DescribeWorkflowExecution()

	if err != nil{
		assert.Fail(t,"Cannot build cadence client")
	}

	fmt.Println("Cadence client built " , testCadenceClient)
}