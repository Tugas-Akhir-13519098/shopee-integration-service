package util

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"shopee-integration-service/src/model"
)

func ConvertProductToCreateProductRequest(pm *model.ProductMessage) *bytes.Buffer {
	product := model.CreateItemRequest{
		OriginalPrice: float32(pm.Price),
		Description:   pm.Description,
		Weight:        pm.Weight,
		ItemName:      pm.Name,
		LogisticInfo:  []model.LogisticInfo{{Enabled: true, LogisticId: 1}},
		CategoryID:    1,
		Image:         model.Image{ImageIDList: []string{pm.Image}},
		SellerStock:   []model.Stock{{Stock: pm.Stock}},
	}
	body, _ := json.Marshal(product)
	responseBody := bytes.NewBuffer(body)

	return responseBody
}

func ConvertProductToUpdateItemRequest(pm *model.ProductMessage) *bytes.Buffer {
	product := model.UpdateItemRequest{
		ItemID:        pm.ShopeeProductID,
		OriginalPrice: float32(pm.Price),
		Description:   pm.Description,
		Weight:        pm.Weight,
		ItemName:      pm.Name,
		Image:         model.Image{ImageIDList: []string{pm.Image}},
		SellerStock:   []model.Stock{{Stock: pm.Stock}},
	}
	body, _ := json.Marshal(product)
	responseBody := bytes.NewBuffer(body)

	return responseBody
}

func ConvertProductToDeleteItemRequest(pm *model.ProductMessage) *bytes.Buffer {
	product := model.DeleteItemRequest{
		ItemID: []int{pm.ShopeeProductID},
	}
	body, _ := json.Marshal(product)
	responseBody := bytes.NewBuffer(body)

	return responseBody
}

func ConvertResponseToItemResponse(body io.ReadCloser) model.ItemResponse {
	respBody, _ := io.ReadAll(body)
	var itemResponse model.ItemResponse
	err := json.Unmarshal(respBody, &itemResponse)
	if err != nil {
		fmt.Println(err)
	}

	return itemResponse
}

func ConvertProductIdToUpdateProductIdRequest(productID int) *bytes.Buffer {
	request := model.UpdateProductIdRequest{
		ShopeeProductID: productID,
	}
	body, _ := json.Marshal(request)
	responseBody := bytes.NewBuffer(body)

	return responseBody
}
