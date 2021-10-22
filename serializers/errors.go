package serializers

type SerializationError struct {
	msg string
}

func (err SerializationError) Error() string {
	return "Cannot serialize/deserialize object🤧"
}
