package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/RichardKnop/machinery/v1"
	"github.com/RichardKnop/machinery/v1/config"
	"github.com/RichardKnop/machinery/v1/tasks"
)

func main() {
	ts := GetMachineryServer()

	go StartWorker(ts)

	go func() {
		for {
			time.Sleep(10 * time.Second)
			SendTask(ts)
		}
	}()

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	<-ch
}

func GetMachineryServer() *machinery.Server {
	ts, err := machinery.NewServer(&config.Config{
		Broker:        "redis://localhost:6379",
		ResultBackend: "redis://localhost:6379",
	})
	if err != nil {
		log.Fatal(err.Error())
	}

	err = ts.RegisterTasks(map[string]interface{}{
		"send_email": SendMail,
	})
	if err != nil {
		log.Fatal(err.Error())
	}

	return ts
}

func StartWorker(ts *machinery.Server) error {
	worker := ts.NewWorker("machinery_worker", 10)

	if err := worker.Launch(); err != nil {
		return err
	}

	return nil
}

func SendTask(ts *machinery.Server) {
	p := SendMailPayload{
		Email: "gopher@golang.org",
	}

	reqJSON, err := json.Marshal(p)
	if err != nil {
		log.Println(err.Error())
	}

	b64EncodedReq := base64.StdEncoding.EncodeToString(reqJSON)
	task := tasks.Signature{
		Name: "send_email",
		Args: []tasks.Arg{
			{
				Type:  "string",
				Value: b64EncodedReq,
			},
		},
	}

	res, err := ts.SendTask(&task)
	if err != nil {
		log.Println(err.Error())
	}

	fmt.Println("task_uuid:", res.GetState().TaskUUID)
}