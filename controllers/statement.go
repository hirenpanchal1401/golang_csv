package controllers

import (
	"fmt"
	requestresponse "statement_parser/request_response"
	"statement_parser/utils"
	"strconv"
	"strings"
)

// TotalByDateAndCurrency : returns total payment by currency for each date
func TotalByDateAndCurrency(input []requestresponse.Statement, typeOfPayment string) (total map[string]map[string]float64, err error) {
	total = map[string]map[string]float64{}

	for _, r := range input {
		if strings.HasPrefix(r.Narrative1, typeOfPayment) {
			_, ok := total[r.Date]
			if !ok {
				total[r.Date] = map[string]float64{}
			}

			amount := r.Credit
			if amount == 0 {
				amount = r.Debit
			}

			_, ok = total[r.Date][r.Currency]
			if !ok {
				total[r.Date][r.Currency] = amount
			} else {
				total[r.Date][r.Currency] += amount
			}
		}
	}

	return
}

// TotalByCurrency ...
func TotalByCurrency(req requestresponse.RequestBody) (res requestresponse.Response, err error) {
	parsedCSVRecords, err := utils.ParseCSV(req.FileName)
	if err != nil {
		fmt.Println("Error in ParseCSV:", err)
		return
	}

	if len(parsedCSVRecords) < 2 {
		fmt.Println("empty CSV file")
		return res, fmt.Errorf("empty CSV file")
	}

	statementData := []requestresponse.Statement{}

	for _, r := range parsedCSVRecords[1:] {
		if len(r) < 10 {
			fmt.Println("invalid CSV file")
			return res, fmt.Errorf("invalid CSV file")
		}

		credit := 0.0
		if r[7] != "" {
			credit, err = strconv.ParseFloat(r[7], 64)
			if err != nil {
				fmt.Println("Error in Credit amount parsing")
				return res, err
			}
		}

		debit := 0.0
		if r[8] != "" {
			debit, err = strconv.ParseFloat(r[8], 64)
			if err != nil {
				fmt.Println("Error in Debit amount parsing")
				return res, err
			}
		}

		statementData = append(statementData, requestresponse.Statement{
			Date:       r[0],
			Narrative1: r[1],
			Narrative2: r[2],
			Narrative3: r[3],
			Narrative4: r[4],
			Narrative5: r[5],
			Type:       r[6],
			Credit:     credit,
			Debit:      debit,
			Currency:   r[9],
		})
	}

	total, err := TotalByDateAndCurrency(statementData, req.TypeOfPayment)
	if err != nil {
		fmt.Println("Error in TotalByDateAndCurrency:", err)
		return
	}

	return total[req.Date], err
}
