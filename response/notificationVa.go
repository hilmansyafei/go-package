package response

// NotificationVa : hold data NotificationVa
type NotificationVa struct {
	Request   DataRequest `json:"pid" bson:"pid"`
	Signature string      `json:"signature" bson:"signature"`
}

// DataRequest : hold data DataRequest
type DataRequest struct {
	Data   DataNotif  `json:"data" bson:"data"`
	Result DataResult `json:"result" bson:"result"`
}

// DataResult : hold data DataResult
type DataResult struct {
	Code    string `json:"code" bson:"code"`
	Message string `json:"message" bson:"message"`
	Status  string `json:"status" bson:"status"`
}

// DataNotif : hold data DataNotif
type DataNotif struct {
	Amount        int    `json:"amount" bson:"amount"`
	Bank          string `json:"bank" bson:"bank"`
	Number        string `json:"number" bson:"number"`
	OrderID       string `json:"orderID" bson:"orderID"`
	Status        string `json:"status" bson:"status"`
	Time          Time   `json:"time" bson:"time"`
	TransactionID string `json:"transactionID" bson:"transactionID"`
}

// Time : hold data Time
type Time struct {
	Created string `json:"created" bson:"created"`
	Updated string `json:"updated" bson:"updated"`
}
