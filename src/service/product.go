package service

import (
	"context"
	"encoding/json"
	"fmt"
	"shopee-integration-service/config"
	"shopee-integration-service/src/model"
	"shopee-integration-service/src/util"

	"github.com/segmentio/kafka-go"
)

type ProductService interface {
	ConsumeProductMessages()
}

type productService struct{}

func NewProductService() ProductService {
	return &productService{}
}

func (ps *productService) ConsumeProductMessages() {
	// Set up the Kafka reader for product topic
	cfg := config.Get()
	config := kafka.ReaderConfig{
		Brokers: []string{fmt.Sprintf("%s:%s", cfg.KafkaHost, cfg.KafkaPort)},
		Topic:   cfg.KafkaProductTopic,
		GroupID: cfg.KafkaProductConsumerGroup,
	}
	reader := kafka.NewReader(config)

	// Continuously read messages from Kafka
	for {
		msg, err := reader.ReadMessage(context.Background())
		if err != nil {
			fmt.Println("Error reading message from Kafka:", err.Error())
			continue
		}

		// Change kafka message from byte to struct
		var productMessage *model.KafkaProductMessage
		err = json.Unmarshal(msg.Value, &productMessage)
		if err != nil {
			fmt.Println("Can't unmarshal the kafka message")
			continue
		}

		if productMessage.Method == model.CREATE {
			createProductBody := util.ConvertProductToCreateProductRequest(productMessage)
			url := cfg.ShopeeURL + fmt.Sprintf("add_item?partner_id=%d&shop_id=%d&access_token=%s&sign=%s",
				productMessage.ShopeePartnerID, productMessage.ShopeeShopID, productMessage.ShopeeAccessToken, productMessage.ShopeeSign)
			resp, _ := util.SendPostRequest(createProductBody, url, "")
			util.AfterHTTPRequestHandler(createProductBody.String(), resp, "CREATE", "POST", productMessage.ID, url)

		} else if productMessage.Method == model.UPDATE {
			updateProductBody := util.ConvertProductToUpdateItemRequest(productMessage)
			url := cfg.ShopeeURL + fmt.Sprintf("update_item?partner_id=%d&shop_id=%d&access_token=%s&sign=%s",
				productMessage.ShopeePartnerID, productMessage.ShopeeShopID, productMessage.ShopeeAccessToken, productMessage.ShopeeSign)
			resp, _ := util.SendPostRequest(updateProductBody, url, "")
			util.AfterHTTPRequestHandler(updateProductBody.String(), resp, "UPDATE", "POST", string(msg.Key), url)

		} else { // productMessage.Method == model.DELETE
			deleteProductBody := util.ConvertProductToDeleteItemRequest(productMessage)
			url := cfg.ShopeeURL + fmt.Sprintf("delete_item?partner_id=%d&shop_id=%d&access_token=%s&sign=%s",
				productMessage.ShopeePartnerID, productMessage.ShopeeShopID, productMessage.ShopeeAccessToken, productMessage.ShopeeSign)
			resp, _ := util.SendPostRequest(deleteProductBody, url, "")
			util.AfterHTTPRequestHandler(deleteProductBody.String(), resp, "DELETE", "POST", string(msg.Key), url)
		}
	}
}
