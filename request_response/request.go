package requestresponse

type RequestBody struct {
	FileName      string `json:"fileName"`
	Date          string `json:"date"`
	TypeOfPayment string `json:"typeOfPayment"`
}
