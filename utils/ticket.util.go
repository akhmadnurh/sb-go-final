package utils

import (
	"fmt"
	"time"
)

func GenerateTicketCode() (code string) {
	code = "OTX-"

	now := time.Now().Unix()

	code += fmt.Sprintf("%d", now)

	return
}
