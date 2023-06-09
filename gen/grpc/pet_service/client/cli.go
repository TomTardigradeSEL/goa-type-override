// Code generated by goa v3.11.3, DO NOT EDIT.
//
// Pet Service gRPC client CLI support package
//
// Command:
// $ goa gen petsvc/design

package client

import (
	"encoding/json"
	"fmt"
	pet_servicepb "petsvc/gen/grpc/pet_service/pb"
	petservice "petsvc/gen/pet_service"
	"petsvc/interfaces/pets"
)

// BuildSendCatPayload builds the payload for the Pet Service SendCat endpoint
// from CLI flags.
func BuildSendCatPayload(petServiceSendCatMessage string) (*petservice.CatPayload, error) {
	var err error
	var message pet_servicepb.SendCatRequest
	{
		if petServiceSendCatMessage != "" {
			err = json.Unmarshal([]byte(petServiceSendCatMessage), &message)
			if err != nil {
				return nil, fmt.Errorf("invalid JSON for message, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"cat\": \"Consequuntur doloribus quibusdam.\"\n   }'")
			}
		}
	}
	v := &petservice.CatPayload{
		Cat: pets.Cat(message.Cat),
	}

	return v, nil
}
