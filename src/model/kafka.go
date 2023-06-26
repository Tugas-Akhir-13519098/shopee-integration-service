package model

type Method int

const (
	CREATE Method = iota
	UPDATE
	DELETE
)

type KafkaProductMessage struct {
	Method            Method  `json:"method"`
	ID                string  `json:"id"`
	Name              string  `json:"name"`
	Price             int     `json:"price"`
	Weight            float32 `json:"weight"`
	Stock             int     `json:"stock"`
	Image             string  `json:"image"`
	Description       string  `json:"description"`
	ShopeeProductID   int     `json:"shopee_product_id"`
	ShopeePartnerID   int     `json:"shopee_partner_id"`
	ShopeeShopID      int     `json:"shopee_shop_id"`
	ShopeeAccessToken string  `json:"shopee_access_token"`
	ShopeeSign        string  `json:"shopee_sign"`
}

type KafkaErrorMessage struct {
	Method      string `json:"method"`
	Url         string `json:"url"`
	RequestBody string `json:"request_body"`
	Error       string `json:"error"`
	Status      string `json:"status"`
	RequestTime string `json:"request_time"`
}
