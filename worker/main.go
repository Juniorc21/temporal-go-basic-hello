package main

import (
	"hello-world-temporal/app"
	"log"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
)

func main() {
	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalln("unable to create Temporal client", err)
	}
	defer c.Close()

	w := worker.New(c, app.GreetingTaskQueue, worker.Options{})

	w.RegisterWorkflow(app.GreetingWorkFlow)
	w.RegisterActivity(app.ComposeGreeting)

	err = w.Run(worker.InterruptCh())

	if err != nil {
		log.Fatalln("unable to start worker")
	}
}
