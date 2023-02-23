package schema

import "time"

// Structs for the API
type Center struct {
	Id            string
	Longitude     float32
	Latitude      float32
	Speed         int32
	Temperature   int32
	Address       string
	StopsNormal   int32
	StopsIregular int32
	Camera        string
	State         bool
	Panic         bool
}

// Structs for the Response
type Response struct {
	Status  string
	Message string
}

// Structs for the Task
type Task struct {
	Body      Center
	Start     time.Time
	Emergency []string
	Warning   []string
	Endpoints []string
	Elapsed   time.Duration
}

// Structs for the DynamoDB
type TableStruct struct {
	Id            string   `dynamodbav:"Id"`
	CreationDate  string   `dynamodbav:"CreationDate"`
	Longitude     float32  `dynamodbav:"Longitude"`
	Latitude      float32  `dynamodbav:"Latitude"`
	Speed         int32    `dynamodbav:"Speed"`
	Temperature   int32    `dynamodbav:"Temperature"`
	Address       string   `dynamodbav:"Address"`
	StopsNormal   int32    `dynamodbav:"StopsNormal"`
	StopsIregular int32    `dynamodbav:"StopsIregular"`
	Camera        string   `dynamodbav:"Camera"`
	State         bool     `dynamodbav:"State"`
	Panic         bool     `dynamodbav:"Panic"`
	Elapsed       int64    `dynamodbav:"Elapsed"`
	Warnings      []string `dynamodbav:"Warnings"`
	Emergencies   []string `dynamodbav:"Emergencies"`
}
