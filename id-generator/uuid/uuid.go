package uuid

import (
	"log"
	"strings"

	uuid "github.com/sliveryou/go-tool/id-generator/uuid/satori"
)

// NextV1 generates v1 uuid.
func NextV1() string {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("uuid: NextV1 panic: %v", r)
		}
	}()

	u, err := uuid.NewV1()
	if err != nil {
		panic(err)
	}

	return strings.ReplaceAll(u.String(), "-", "")
}

// NextV4 generates v4 uuid.
func NextV4() string {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("uuid: NextV4 panic: %v", r)
		}
	}()

	u, err := uuid.NewV4()
	if err != nil {
		panic(err)
	}

	return strings.ReplaceAll(u.String(), "-", "")
}

// Parse parses uuid.
func Parse(input string) (uuid.UUID, error) {
	return uuid.FromString(input)
}
