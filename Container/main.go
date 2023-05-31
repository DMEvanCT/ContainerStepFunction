package main

import (
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sfn"
	"log"
	"os"
)

type EmployeeInfo struct {
	EmployeeId int    `json:"employeeId"`
	FirstName  string `json:"firstName"`
	LastName   string `json:"lastName"`
	City       string `json:"city"`
	State      string `json:"state"`
}

type Employee struct {
	EmployeeInfo EmployeeInfo `json:"employeeInfo"`
}

type TaskFailure struct {
	Error string `json:"error"`
	Cause string `json:"cause"`
}

func main() {
	TaskToken := os.Getenv("TASK_TOKEN")
	employeeData := os.Getenv("EMPLOYEE_JSON_ENV")
	region := os.Getenv("AWS_REGION")

	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(region),
	})
	if err != nil {
		fmt.Println("Error creating session: ", err)

	}
	stepFunction := sfn.New(sess)

	var emp Employee

	jsonError := json.Unmarshal([]byte(employeeData), &emp)
	if jsonError != nil {
		_, err := stepFunction.SendTaskFailure(&sfn.SendTaskFailureInput{
			Cause:     aws.String("JSON Parse Error"),
			Error:     aws.String("Invalid JSON"),
			TaskToken: aws.String(TaskToken),
		})
		if err != nil {
			log.Fatal(err)
		}
	}

	fmt.Printf("Parsed data: %+v\n", emp)
	_, err = stepFunction.SendTaskSuccess(&sfn.SendTaskSuccessInput{
		Output:    aws.String(employeeData),
		TaskToken: aws.String(TaskToken),
	})
	if err != nil {
		log.Fatal(err)
	}
}
