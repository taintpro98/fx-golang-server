package storage

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"fx-golang-server/module/core/model"

	"github.com/elastic/go-elasticsearch/v7"
)

type IElasticStorage interface {
	AddUsers(ctx context.Context, users []model.UserModel) error
}

type elasticStorage struct {
	es *elasticsearch.Client
}

func NewElasticStorage(
	es *elasticsearch.Client,
) IElasticStorage {
	fmt.Println("fuck storage", es)
	return &elasticStorage{
		es: es,
	}
}

func (v *elasticStorage) AddUsers(ctx context.Context, users []model.UserModel) error {
	var buf bytes.Buffer

	for _, user := range users {
		// Tạo hành động chèn cho mỗi người dùng
		action := map[string]interface{}{
			"index": map[string]interface{}{
				"_index": "users_index",
				"_id":    user.ID,
			},
		}

		// Ghi hành động và dữ liệu vào buffer
		if err := json.NewEncoder(&buf).Encode(action); err != nil {
			return fmt.Errorf("failed to encode action for user %s: %w", user.Name, err)
		}
		if err := json.NewEncoder(&buf).Encode(user); err != nil {
			return fmt.Errorf("failed to encode user %s: %w", user.Name, err)
		}
	}

	// Gửi yêu cầu Bulk
	res, err := v.es.Bulk(bytes.NewReader(buf.Bytes()), v.es.Bulk.WithIndex("users_index"))
	if err != nil {
		return fmt.Errorf("failed to execute bulk request: %w", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		return fmt.Errorf("error response from Elasticsearch: %s", res.Status)
	}
	return nil
}
