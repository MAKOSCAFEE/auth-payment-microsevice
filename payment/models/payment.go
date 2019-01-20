package models

// Payment : represent payment structure
type Payment struct {
	ID       int
	Ref      string
	Amount   float64
	Currency string
}

// Result : represent result structure
type Result struct {
	Message string
}
