package main

import (
	"fmt"
	"study2/feature1"
	"study2/feature2"
	simpleconnection "study2/feature_postgres/simple_connection"
)

func main() {
	fmt.Println("Hello, Git!")

	feature1.Feature1()

	feature2.Feature2()

	simpleconnection.SimpleConnection()
}
