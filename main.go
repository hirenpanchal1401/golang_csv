package main

import (
	"fmt"
	"statement_parser/controllers"
	requestresponse "statement_parser/request_response"
)

func main() {
	/*
		Notice: we can take this inputs as a HTTP request
	*/

	// sample request body
	sampleRequest := requestresponse.RequestBody{
		FileName:      "statement.csv",
		Date:          "06/03/2011",
		TypeOfPayment: "PAY",
	}

	// get total by currency and payment type
	output, err := controllers.TotalByCurrency(sampleRequest)
	if err != nil {
		fmt.Println("Error in TotalByCurrency: ", err)
		return
	}

	/*
		we can return this output data as HTTP response
	*/

	// print output
	for k, v := range output {
		fmt.Println(k, ": ", v)
	}
}
