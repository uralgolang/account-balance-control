package request

type PayingType uint8

const (
	Visa       PayingType = 1
	MasterCard            = 2
)

type DepositAccountRequest struct {
	AccountId  string     `json:"accountId"`
	PayingType PayingType `json:"payingType"`
	Amount     int        `json:"amount"`
}
