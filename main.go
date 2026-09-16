package main

import (
	"fmt"
	"study/feature1"
	simpleconnection "study/feature_postgres/simple_connection"
)

func main() {
	fmt.Println("Hello, Git!")
	feature1.Feature1()

	simpleconnection.CheckConnection()
}
