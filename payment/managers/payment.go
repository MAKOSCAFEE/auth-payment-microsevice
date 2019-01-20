package managers

import "auth-payment-microservice/payment/models"

// GetPayments : Return list of users.
func GetPayments() []models.Payment {
	payments := []models.Payment{
		models.Payment{ID: 1, Ref: "test1", Amount: 65.4, Currency: "GBP"},
		models.Payment{ID: 2, Ref: "test2", Amount: 50.9, Currency: "GBP"},
		models.Payment{ID: 3, Ref: "test3", Amount: 60.7, Currency: "GBP"},
	}
	return payments
}

// GetPaymentByID : return user by ID
func GetPaymentByID(id int) models.Payment {

	for _, item := range GetPayments() {
		if item.ID == id {
			return item
		}
	}

	return models.Payment{}
}
