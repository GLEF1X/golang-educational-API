package serializers

import jsoniter "github.com/json-iterator/go"

var json = jsoniter.ConfigCompatibleWithStandardLibrary

type JsonSerializer struct{}

func (s JsonSerializer) Serialize(data interface{}) ([]byte, error) {
	result, err := json.Marshal(data)
	if err != nil {
		return nil, SerializationError{msg: err.Error()}
	}
	return result, nil
}

func (s JsonSerializer) Deserialize(data []byte, structure interface{}) error {
	err := json.Unmarshal(data, structure)
	if err != nil {
		return SerializationError{msg: err.Error()}
	}
	return nil
}
