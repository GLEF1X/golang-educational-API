package serializers

import "fmt"

type SerializationError struct {
	msg string
}

func (err SerializationError) Error() string {
	return fmt.Sprintf("Cannot serialize/deserialize object🤧")
}
