package engine

import "fmt"

func (cp *EngineProperties) QueryExecutor(plan interface{}) {
	fmt.Printf("Plan : %+v\n", plan)
}
