package main

import (
	"context"
	"fmt"
	"hello-world-temporal/app"
	"log"

	"go.temporal.io/sdk/client"
)

func main() {
	c, err := client.Dial(client.Options{})

	if err != nil {
		log.Fatalln("unable to create Temporal client", err)
	}

	defer c.Close()

	options := client.StartWorkflowOptions{
		ID:        "greeting-workflow",
		TaskQueue: app.GreetingTaskQueue,
	}

	name := "Piter Albeiro"
	we, err := c.ExecuteWorkflow(context.Background(), options, app.GreetingWorkFlow, name)

	if err != nil {
		log.Fatalln("unable to complete Workflow result", err)
	}

	var greeting string

	err = we.Get(context.Background(), &greeting)
	if err != nil {
		log.Fatalln("unable to get Workflow result")
	}

	printResults(greeting, we.GetID(), we.GetRunID())
}

func printResults(greeting, workflowID, runID string) {
	fmt.Printf("\nWorkflowID: %s RunID: %s\n", workflowID, runID)
	fmt.Printf("\n%s\n\n ", greeting)
}
