package main

import (
	"fmt"
	marathon "github.com/gambol99/go-marathon"
	"log"
	"time"
)

func main() {
	fmt.Println("Marathon Test - Started")

	marathonURL := "http://1.1.1.1:8080"
	config := marathon.NewDefaultConfig()
	config.URL = marathonURL
	config.EventsTransport = marathon.EventsTransportSSE
	client, err := marathon.NewClient(config)
	if err != nil {
		log.Fatalf("Failed to create a client for marathon, error: %s", err)
	}

	go listenMarathonEvents(client)
	time.Sleep(5 * time.Second)
	createNewMarathonApplication(client)


	/*for  true   {
		application, err := client.Application("kamaltest")
		if err != nil {
			log.Fatalf("Failed to get application details: %s", err)
		}

		fmt.Println(application.ID)
		fmt.Println(application.TasksRunning)
		fmt.Println(application.AllTaskRunning())

		if (application.TasksRunning > 0) {
			break
		}

		time.Sleep(5 * time.Second)

	}*/

	deleteMarathonApplication(client)
	time.Sleep(30 * time.Second)

	fmt.Println("Marathon Test - Ended")

}

func createNewMarathonApplication(client marathon.Marathon) {
	fmt.Println("Application Creation - Started")

	application := marathon.NewDockerApplication().Name("kamaltest").CPU(1).Memory(128).Storage(0.0).Count(1)

	var healthCheck marathon.HealthCheck;
	healthCheck.Protocol = "TCP"
	var portIndex int = 0
	healthCheck.PortIndex = &portIndex
	application.AddHealthCheck(healthCheck)

	application.Container.Docker.Container("nginx").Bridged().Expose(80)

	if _, err := client.CreateApplication(application); err != nil {
		log.Fatalf("Failed to create application: %s, error: %s", application, err)
	} else {
		log.Printf("Created the application.")

	}

	fmt.Println("Application Creation - Ended")

}

func deleteMarathonApplication(client marathon.Marathon) {
	if deploymentId, err := client.DeleteApplication("kamaltest"); err != nil {
		log.Fatalf("Failed to create application: %s, error: %s", "kamaltest", err)
	} else {
		log.Printf("Deleting the application: %s", deploymentId)
	}
}

func listenMarathonEvents (client marathon.Marathon) {

	fmt.Println("Event Listener - Started")

	// Register for events
	applicationEvents := make(marathon.EventsChannel, 5)
	apiEvents := make(marathon.EventsChannel, 5)

	err := client.AddEventsListener(applicationEvents, marathon.EventIDApplications)
	err = client.AddEventsListener(apiEvents, marathon.EventIDAPIRequest)

	if err != nil {
		log.Fatalf("Failed to register for events, %s", err)
	}

	timer := time.After(60 * time.Second)
	done := false

	// Receive events from channel for 60 seconds
	for {
		if done {
			break
		}
		select {
		case <-timer:
			log.Printf("Exiting the loop")
			done = true
		case event := <-applicationEvents:
			log.Printf("Recieved Application event: %s", event)
		case event := <-apiEvents:
			log.Printf("Recieved API event: %s", event.Name)
		}
	}

	// Unsubscribe from Marathon events
	client.RemoveEventsListener(applicationEvents)

	fmt.Println("Event Listener - Ended")

}