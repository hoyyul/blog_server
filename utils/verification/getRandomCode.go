package verification

import (
	"fmt"
	"math/rand"
	"time"
)

var stringCode = ""

func GetRandomCode() string {
	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)

	return fmt.Sprintf("%04d", random.Intn(10000))
}
