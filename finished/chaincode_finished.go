/*
Licensed to the Apache Software Foundation (ASF) under one
or more contributor license agreements.  See the NOTICE file
distributed with this work for additional information
regarding copyright ownership.  The ASF licenses this file
to you under the Apache License, Version 2.0 (the
"License"); you may not use this file except in compliance
with the License.  You may obtain a copy of the License at

  http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing,
software distributed under the License is distributed on an
"AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
KIND, either express or implied.  See the License for the
specific language governing permissions and limitations
under the License.
*/

package main

import (
	"errors"
	"fmt"
//	"strconv"
	"encoding/json"
//	"time"
//	"strings"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

// SimpleChaincode example simple Chaincode implementation
type SimpleChaincode struct {
}

type Adminlogin struct{
	Userid string `json:"userid"`					//User login for system Admin
	Password string `json:"password"`
}

// ============================================================================================================================
// Main
// ============================================================================================================================
func main() {
	err := shim.Start(new(SimpleChaincode))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}

// ============================================================================================================================
// Init - reset all the things
// ============================================================================================================================
func (t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	var err error

	if len(args) != 2 {
	   return nil, errors.New("Incorrect number of arguments. Expecting 1")
	}
    //Write the User Id "mail Id" arg[0] and password arg[1]
	userid := args[0]															//argument for UserID
	password := args[1]  	//argument for password
	str := `{"userid": "` + userid+ `", "password": "` + password + `"}`
	
	err = stub.PutState(userid, []byte(str))								//Put the userid and password in blockchain
	
	if err != nil {
		return nil, err
	}
	
	return nil, nil
}

// ============================================================================================================================
// Run - Our entry point for Invocations - [LEGACY] obc-peer 4/25/2016
// ============================================================================================================================
//func (t *SimpleChaincode) Run(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
//	fmt.Println("run is running " + function)
//	return t.Invoke(stub, function, args)
//}

// ============================================================================================================================
// Invoke - Our entry point for Invocations
// ============================================================================================================================
func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
return nil, nil

}

// ============================================================================================================================
// Query - Our entry point for Queries
// ============================================================================================================================
func (t *SimpleChaincode) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
    fmt.Println("query is running " + function)

	// Handle different functions
	if function == "read" { //read a variable
		return t.read(stub, args)
	}
	fmt.Println("query did not find func: " + function)

	return nil, errors.New("Received unknown function query: " + function)
}
//==============================================================================================================================
// read - query function to read key/value pair
//===============================================================================================================================
func (t *SimpleChaincode) read(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	var err error

	if len(args) != 1 {
		return nil, errors.New("Incorrect number of arguments. Expecting name of the key to query")
	}

	userid := args[0]
	PassAsbytes, err := stub.GetState(userid)
	
	if err != nil {
		jsonResp := "{\"Error\":\"Failed to get state for " + userid + "\"}"
		return nil, errors.New(jsonResp)
	}
	
	res := Adminlogin{}
	json.Unmarshal(PassAsbytes,&res)
	
	if res.Userid == userid{
	   fmt.Println("The userid password matched: " +res.Userid + res.Password);
	   }
	fmt.Println("The value for Userid is: " +res.Userid + res.Password);
	
	return PassAsbytes, nil
}
