package tracing

import (
	"fmt"

	"github.com/upikoth/starter-go/internal/pkg/functions"
)

func GetHandlerTraceName() string {
	functionName := functions.GetFunctionName(functions.WithSkip(1))
	packageName := functions.GetPackageName(functions.WithSkip(1))

	return fmt.Sprintf("controller.http.%s.%s", packageName, functionName)
}

func GetServiceTraceName() string {
	functionName := functions.GetFunctionName(functions.WithSkip(1))
	packageName := functions.GetPackageName(functions.WithSkip(1))

	return fmt.Sprintf("services.%s.%s", packageName, functionName)
}

func GetRepositoryTraceName() string {
	functionName := functions.GetFunctionName(functions.WithSkip(1))
	packageName := functions.GetPackageName(functions.WithSkip(1))

	return fmt.Sprintf("repository.%s.%s", packageName, functionName)
}

func GetRepositoryYDBTraceName() string {
	functionName := functions.GetFunctionName(functions.WithSkip(1))
	packageName := functions.GetPackageName(functions.WithSkip(1))

	return fmt.Sprintf("repository.ydb.%s.%s", packageName, functionName)
}
