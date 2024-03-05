package domain

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPayload(t *testing.T) {
	jsonMessage := `{
		"schema": {
			"type": "struct",
			"fields": [
				{
					"type": "int32",
					"optional": false,
					"field": "id"
				},
				{
					"type": "string",
					"optional": false,
					"field": "uuid"
				},
				{
					"type": "string",
					"optional": false,
					"field": "user_id"
				},
				{
					"type": "string",
					"optional": false,
					"field": "transaction_code"
				},
				{
					"type": "string",
					"optional": false,
					"field": "json"
				},
				{
					"type": "int64",
					"optional": true,
					"name": "org.apache.kafka.connect.data.Timestamp",
					"version": 1,
					"field": "created_at"
				},
				{
					"type": "int64",
					"optional": true,
					"name": "org.apache.kafka.connect.data.Timestamp",
					"version": 1,
					"field": "updated_at"
				}
			],
			"optional": false,
			"name": "outbox"
		},
		"payload": {
			"id": 7,
			"uuid": "618842c7-23b6-4769-ad95-dd5015d7834b",
			"user_id": "434baec3-c230-4d20-9907-d532d8483541",
			"transaction_code": "112dsfdfdsdq",
			"json": "{\"name\":\"John\",\"surname\":\"Doe\",\"details\":{\"code\":123,\"operation_code\":456,\"transaction_code\":\"sdfsdfsdfs\"}}",
			"created_at": 1709593948361,
			"updated_at": 1709593948361
		}
	}`

	var outbox Outbox

	err := json.Unmarshal([]byte(jsonMessage), &outbox)
	assert.NoError(t, err)
}
