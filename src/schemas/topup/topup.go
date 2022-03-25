package topup

type TopupSchema struct {
	PhoneNumber          string `json:"phone_number"`
	RecipientPhoneNumber string `json:"recipient_phone_number"`
	Amount               int    `json:"amount"`
	Brand                string `json:"brand"`
	PaymentMethod        string `json:"payment_method"`
	Status               string `json:"status"`
	Description          string `json:"description"`
	TopupTime            string `json:"topup_time"`
}

type GetTopupDetailsResponse struct {
	Total  int64          `json:"total"`
	Limit  int            `json:"limit"`
	Offset int            `json:"offset"`
	Body   []*TopupSchema `json:"body"`
}

type TopupGMVResponse struct {
	GMV int `json:"gmv"`
}
