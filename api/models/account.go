package models

type Account struct {
	Id        string  `json:"id"`
	Name      string  `json:"name"`
	Phone     string  `json:"phone"`
	Balance   float64 `json:"balance"`
	CreatedAt string  `json:"created_at"`
	UpdatedAt string  `json:"updated_at"`
}

type CreateAccount struct {
	Name    string  `json:"name"`
	Phone   string  `json:"phone"`
	Balance float64 `json:"balance"`
}

type Deposit struct {
	Id     string  `json:"id"`
	Amount float64 `json:"amount"`
}

type Withdraw struct {
	Id     string  `json:"id"`
	Amount float64 `json:"amount"`
}
