package main

import (
	"fmt"

	"github.com/eug48/fhir/utils"
)

func main() {
	processed, err := utils.ProcessDirectory("../../models", "models")
	if err != nil {
		fmt.Printf("Error: %s", err.Error())
	}
	fmt.Printf("processed: %v\n", processed)
}
