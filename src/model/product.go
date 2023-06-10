package model

type Method int

const (
	CREATE Method = iota
	UPDATE
	DELETE
)

type ProductMessage struct {
	Method             Method
	ID                 string
	Name               string
	Price              int
	Weight             float32
	Stock              int
	Image              string
	Description        string
	TokopediaProductID int
	ShopeeProductID    int
}

type Image struct {
	ImageIDList []string `json:"image_id_list" binding:"required"`
}

type Stock struct {
	Stock int `json:"stock" binding:"required"`
}

type LogisticInfo struct {
	Enabled    bool `json:"enabled" binding:"required"`
	LogisticId int  `json:"logistic_id" binding:"required"`
}

type CreateItemRequest struct {
	OriginalPrice float32        `json:"original_price" binding:"required"`
	Description   string         `json:"description" binding:"required"`
	Weight        float32        `json:"weight" binding:"required"`
	ItemName      string         `json:"item_name" binding:"required"`
	LogisticInfo  []LogisticInfo `json:"logistic_info" binding:"required"`
	CategoryID    int            `json:"category_id" binding:"required"`
	Image         Image          `json:"image" binding:"required"`
	SellerStock   []Stock        `json:"seller_stock" binding:"required"`
}

type UpdateItemRequest struct {
	ItemID        int     `json:"item_id" binding:"required"`
	OriginalPrice float32 `json:"original_price" binding:"required"`
	Description   string  `json:"description" binding:"required"`
	Weight        float32 `json:"weight" binding:"required"`
	ItemName      string  `json:"item_name" binding:"required"`
	Image         Image   `json:"image" binding:"required"`
	SellerStock   []Stock `json:"seller_stock" binding:"required"`
}

type DeleteItemRequest struct {
	ItemID []int `json:"item_id" binding:"required"`
}

type ItemResponse struct {
	Error     string       `json:"error" binding:"required"`
	Message   string       `json:"message" binding:"required"`
	Warning   string       `json:"warning" binding:"required"`
	RequestID string       `json:"request_id" binding:"required"`
	Response  ResponseData `json:"response"`
}

type ResponseData struct {
	ItemID int `json:"item_id"`
}

type UpdateProductIdRequest struct {
	ShopeeProductID int `json:"shopee_product_id" binding:"required"`
}
