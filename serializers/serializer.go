package serializers

type Serializer interface {
	Serialize(data interface{}) ([]byte, error)
	Deserialize(data []byte, structure interface{}) error
}
