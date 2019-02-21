package models

// Payment : represent payment structure
type Payment struct {
	ID       int     `json:"id"`
	Ref      string  `json:"reference"`
	Amount   float64 `json:"amount"`
	Currency string  `json:"currency"`
}

// Result : represent result structure
type Result struct {
	Message string `json:"message"`
}
