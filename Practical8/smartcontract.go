package main

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// SmartContract provides functions for managing a data marketplace
type SmartContract struct {
	contractapi.Contract
}

// Dataset describes a listed dataset asset
type Dataset struct {
	ID          string `json:"ID"`
	Name        string `json:"Name"`
	Description string `json:"Description"`
	Price       int    `json:"Price"`
	Owner       string `json:"Owner"`
}

// InitLedger seeds the ledger with a sample dataset
func (s *SmartContract) InitLedger(ctx contractapi.TransactionContextInterface) error {
	datasets := []Dataset{
		{ID: "dataset1", Name: "Customer Purchase Dataset", Description: "Age, income and purchase behavior", Price: 100, Owner: "Seller1"},
	}

	for _, dataset := range datasets {
		datasetJSON, err := json.Marshal(dataset)
		if err != nil {
			return err
		}
		if err := ctx.GetStub().PutState(dataset.ID, datasetJSON); err != nil {
			return fmt.Errorf("failed to put to world state: %v", err)
		}
	}

	return nil
}

// CreateDataset lists a new dataset on the ledger
func (s *SmartContract) CreateDataset(ctx contractapi.TransactionContextInterface, id string, name string, description string, price int, owner string) error {
	exists, err := s.DatasetExists(ctx, id)
	if err != nil {
		return err
	}
	if exists {
		return fmt.Errorf("dataset %s already exists", id)
	}

	dataset := Dataset{ID: id, Name: name, Description: description, Price: price, Owner: owner}
	datasetJSON, err := json.Marshal(dataset)
	if err != nil {
		return err
	}

	return ctx.GetStub().PutState(id, datasetJSON)
}

// ReadDataset returns a dataset by ID
func (s *SmartContract) ReadDataset(ctx contractapi.TransactionContextInterface, id string) (*Dataset, error) {
	datasetJSON, err := ctx.GetStub().GetState(id)
	if err != nil {
		return nil, fmt.Errorf("failed to read from world state: %v", err)
	}
	if datasetJSON == nil {
		return nil, fmt.Errorf("dataset %s does not exist", id)
	}

	var dataset Dataset
	if err := json.Unmarshal(datasetJSON, &dataset); err != nil {
		return nil, err
	}

	return &dataset, nil
}

// DatasetExists checks whether a dataset ID is already on the ledger
func (s *SmartContract) DatasetExists(ctx contractapi.TransactionContextInterface, id string) (bool, error) {
	datasetJSON, err := ctx.GetStub().GetState(id)
	if err != nil {
		return false, fmt.Errorf("failed to read from world state: %v", err)
	}

	return datasetJSON != nil, nil
}

// PurchaseDataset transfers access/ownership of a dataset to the buyer
func (s *SmartContract) PurchaseDataset(ctx contractapi.TransactionContextInterface, id string, buyer string) error {
	dataset, err := s.ReadDataset(ctx, id)
	if err != nil {
		return err
	}

	dataset.Owner = buyer

	datasetJSON, err := json.Marshal(dataset)
	if err != nil {
		return err
	}

	return ctx.GetStub().PutState(id, datasetJSON)
}

// GetAllDatasets returns every dataset currently on the ledger
func (s *SmartContract) GetAllDatasets(ctx contractapi.TransactionContextInterface) ([]*Dataset, error) {
	resultsIterator, err := ctx.GetStub().GetStateByRange("", "")
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	var datasets []*Dataset
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}

		var dataset Dataset
		if err := json.Unmarshal(queryResponse.Value, &dataset); err != nil {
			return nil, err
		}
		datasets = append(datasets, &dataset)
	}

	return datasets, nil
}
