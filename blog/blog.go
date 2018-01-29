/**
 * A Blog App based on block chain
 */
package main

import (
	"bytes"
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
	"time"
)

type SimpleConstract struct {
}

type Article struct {
	Title     string `json:"title"`
	Content   string `json:"content"`
	IsDeleted bool   `json:"is_deleted"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
}

// Init ...
func (s *SimpleConstract) Init(stub shim.ChaincodeStubInterface) peer.Response {
	// do something
	return shim.Success(nil)
}

// Invoke for user
func (s *SimpleConstract) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
	fn, args := stub.GetFunctionAndParameters()

	var err error
	if fn == "write" {
		err = write(stub, args)
	} else if fn == "update" {
		err = update(stub, args)
	} else if fn == "query" {
		err = query(stub, args)
	} else if fn == "delete" {
		err = delete(stub, args)
	} else if fn == "comment" {
		err = comment(stub, args)
	}

	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(nil)
}

// write blog
func write(stub shim.ChaincodeStubInterface, args []string) error {
	author, err := getAuthor(stub)
	if err != nil {
		return fmt.Errorf("Failed to get creator: %s", err.Error())
	}

	article := Article{
		Title:     args[0],
		Content:   args[1],
		IsDeleted: false,
		CreatedAt: time.Now().Unix(),
	}

	articleJson, err := json.Marshal(article)
	if err != nil {
		return fmt.Errorf("Failed to convert article to string:%s", err.Error())
	}

	err = stub.PutState(author, articleJson)
	if err != nil {
		return fmt.Errorf("Failed to write article:%s", err.Error())
	}

	return nil

}

// update article
func update(stub shim.ChaincodeStubInterface, args []string) error {
	return nil
}

// query article
func query(stub shim.ChaincodeStubInterface, args []string) error {
	return nil
}

// delete article
func delete(stub shim.ChaincodeStubInterface, args []string) error {
	return nil
}

// comment article
func comment(stub shim.ChaincodeStubInterface, args []string) error {
	return nil
}

// getAuthor name
func getAuthor(stub shim.ChaincodeStubInterface) (string, error) {
	creatorByte, err := stub.GetCreator()
	certStart := bytes.IndexAny(creatorByte, "-----BEGIN")
	if certStart == -1 {
		return "", fmt.Errorf("No certificate found")
	}

	certText := creatorByte[certStart:]
	certPem, _ := pem.Decode(certText)
	if certPem == nil {
		return "", fmt.Errorf("Failed to decode pem")
	}

	cert509, err := x509.ParseCertificate(certPem.Bytes)
	if err != nil {
		return "", fmt.Errorf("Failed to ParseCertificate")
	}

	return cert509.Subject.CommonName, nil
}

// main function starts up the chaincode in the container during instantiate
func main() {
	if err := shim.Start(new(SimpleConstract)); err != nil {
		fmt.Printf("Error starting blog chaincode: %s", err)
	}
}
