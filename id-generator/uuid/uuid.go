package uuid

import (
	"log"
	"strings"

	"github.com/google/uuid"
)

// NextV1 generates v1 uuid.
func NextV1() string {
	u, err := uuid.NewUUID()
	if err != nil {
		log.Printf("uuid: NextV1 err: %v", err)
		return ""
	}

	return strings.ReplaceAll(u.String(), "-", "")
}

// NextV4 generates v4 uuid.
func NextV4() string {
	u, err := uuid.NewRandom()
	if err != nil {
		log.Printf("uuid: NextV4 err: %v", err)
		return ""
	}

	return strings.ReplaceAll(u.String(), "-", "")
}

// Parse parses uuid.
func Parse(input string) (uuid.UUID, error) {
	return uuid.Parse(input)
}
