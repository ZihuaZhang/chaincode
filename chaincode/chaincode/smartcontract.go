package chaincode

import (
	"encoding/json"
	"strconv"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// SmartContract provides functions for managing an Asset
type SmartContract struct {
	contractapi.Contract
}

type EHR struct {
	UserID       string `json:"UserID"`
	LocationIPFS string `json:"LocationIPFS"`
	KeyCipher    string `json:"KeyCipher"`
}

var totalNum = 0

func (s *SmartContract) uploadEHR(ctx contractapi.TransactionContextInterface, isRedactable string, pk string, msp string, locationIPFS string, keyCiper string) (string, error) {
	if isRedactable == "" || pk == "" || msp == "" {
		return "", nil
	}
	aEHR := EHR{
		UserID:       strconv.Itoa(totalNum),
		LocationIPFS: locationIPFS,
		KeyCipher:    keyCiper,
	}
	totalNum++
	aEHRJSON, err := json.Marshal(aEHR)
	if err != nil {
		return "", err
	}
	err = ctx.GetStub().PutState(aEHR.UserID, aEHRJSON)
	return aEHR.UserID, err
}

func (s *SmartContract) queryEHR(ctx contractapi.TransactionContextInterface, userID string) (*EHR, error) {
	aEHRJSON, err := ctx.GetStub().GetState(userID)
	if err != nil {
		return nil, err
	}
	var aEHR EHR
	err = json.Unmarshal(aEHRJSON, &aEHR)
	if err != nil {
		return nil, err
	}

	return &aEHR, nil
}