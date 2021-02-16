package packer

import (
	"encoding/json"

	msgpack "github.com/vmihailenco/msgpack/v5"
	"go.mongodb.org/mongo-driver/bson"
)

// BsonMarshal ...
func BsonMarshal(v interface{}) ([]byte, error) {
	return bson.Marshal(v)
}

// BsonUnmarshal ...
func BsonUnmarshal(data []byte, v interface{}) error {
	return bson.Unmarshal(data, v)
}

// JSONMarshal ...
func JSONMarshal(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}

// JSONUnmarshal ...
func JSONUnmarshal(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}

// MessagePackMarshal ...
func MessagePackMarshal(v interface{}) ([]byte, error) {
	return msgpack.Marshal(v)
}

// MessagePackUnmarshal ...
func MessagePackUnmarshal(data []byte, v interface{}) error {
	return msgpack.Unmarshal(data, v)
}
