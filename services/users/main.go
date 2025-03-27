package main

import (
	"encoding/json"
	"fmt"
	"github.com/titsex/another-messenger/shared/proto/generated/go/user/v1"
)

func main() {
	user := &userpb.User{
		Id:    "1",
		Name:  "titsex",
		Email: "titsex@titsex.dev",
		Abc:   "test field",
	}

	userJSON, err := json.MarshalIndent(user, "", "  ")
	if err != nil {
		fmt.Printf("Failed to marshal user to JSON: %v\n", err)
		return
	}

	fmt.Printf("User: %+v\n", string(userJSON))
}
