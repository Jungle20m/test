package mobilecard

type MobileCardSchema struct {
	CustomerID   string `json:"customer_id"`
	PhoneNumber  string `json:"phone_number"`
	Name         string `json:"name"`
	Brand        string `json:"brand"`
	Price        int    `json:"price"`
	Point        int    `json:"point"`
	Quantity     int    `json:"quantity"`
	ExchangeTime string `json:"exchange_time"`
}

type GetAllDataResponse struct {
	Total  int64               `json:"total"`
	Limit  int                 `json:"limit"`
	Offset int                 `json:"offset"`
	Body   []*MobileCardSchema `json:"body"`
}

type MobileCardGMVSchema struct {
	Brand string `json:"brand"`
	GMV   int    `json:"gmv"`
}

type MobileCardGMVResponse struct {
	Body []*MobileCardGMVSchema `json:"body"`
}
