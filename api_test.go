package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"testing"
)

//Created some data structures to use while testing
var id int

type cUserResp struct {
	Status string
	User   User
}

type testmsg struct {
	Message string
}

// TestCreateUser() => Testing function to check the post request for creating user

func TestCreateUser(t *testing.T) {

	//Create local user for uploading to create new user
	user := User{Name: "testuser", Address: "Utopia"}

	json_data, err := json.Marshal(user) //Convert sturct to json format

	if err != nil {
		t.Fatal(err)
	}

	//Make post request to the api with json data
	resp, err := http.Post("https://userapi-1.herokuapp.com/create", "application/json",
		bytes.NewBuffer(json_data))

	if err != nil {
		t.Fatal(err)
	}

	defer resp.Body.Close() //Close body after the end of function

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}

	var gotUser cUserResp //Create object to store recieved data

	json.Unmarshal(body, &gotUser) //Bind recieved data to the local object

	expectedStatus := "User created successfully."

	//Compare the returned data with the expected data
	if gotUser.Status != expectedStatus || gotUser.User.Name != "testuser" || gotUser.User.Address != "Utopia" {
		t.Errorf("Got wrong result: Expected %v Got %v", user, gotUser)
	}

	id = gotUser.User.Id //Store the id of newly created user for testing other fucntions

}

// TestGetUser() => Testing function to check get request for getting user data

func TestGetUser(t *testing.T) {

	//Make Get request to the api with id previously stored from the created user
	resp, err := http.Get("https://userapi-1.herokuapp.com/get/" + strconv.Itoa(id))
	if err != nil {
		t.Fatal(err)
	}

	defer resp.Body.Close() //Close body after the end of function

	//If status not ok fail test with error
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Returned wrong status code. Expected %v, Got %v", http.StatusOK, resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var gotUser User

	json.Unmarshal(body, &gotUser)

	//Create expected value of user
	var expectedUser = User{Id: id, Name: "testuser", Address: "Utopia"}

	//Check returned value and expected value
	if gotUser != expectedUser {
		t.Errorf("Returned Unexpected Body: \nGot \n%v \nWant \n%v", gotUser, expectedUser)
	}
}

// TestUpdateUser() => Testing function to check the put request for updating user

func TestUpdateUser(t *testing.T) {

	//Create a client
	client := &http.Client{}

	user := User{Name: "Hehe"} //Create data to update in user

	// marshal User to json
	jsond, err := json.Marshal(user)
	if err != nil {
		t.Fatal(err)
	}

	// set the HTTP method, url, and request body
	req, err := http.NewRequest(http.MethodPut, "http://userapi-1.herokuapp.com/update/"+strconv.Itoa(id), bytes.NewBuffer(jsond))
	if err != nil {
		t.Fatal(err)
	}

	// set the request header Content-Type for json
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	resp, err := client.Do(req)
	if err != nil {
		t.Fatal(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}

	var msg testmsg

	//Unmarsahl the body data to local msg struct
	json.Unmarshal(body, &msg)

	expected := "User updated successfully."

	//Check the returned values and expected values
	if msg.Message != expected {
		t.Errorf("Got wrong result: Expected %v Got %v", expected, msg.Message)
	}
}

// TestDeleteUser() => Testing function to check the delete request for removing user

func TestDeleteUser(t *testing.T) {

	// Create client
	client := &http.Client{}

	// Create request
	req, err := http.NewRequest("DELETE", "http://userapi-1.herokuapp.com/delete/"+strconv.Itoa(id), nil)
	if err != nil {
		t.Error(err)
		return
	}

	// Fetch Request
	resp, err := client.Do(req)
	if err != nil {
		t.Error(err)
		return
	}
	defer resp.Body.Close()

	// Read Response Body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Error(err)
		return
	}

	// Check the returned status
	if resp.StatusCode != http.StatusAccepted {
		t.Errorf("Got wrong Status: Expectd %v Got %v", http.StatusAccepted, resp.StatusCode)
	}

	var msg testmsg
	//Unmarshal the returned data to local msg struct
	json.Unmarshal(body, &msg)

	expected := "User deleted successfully."

	//Check the return value and the expected value
	if msg.Message != expected {
		t.Errorf("Got wrong result: Expected %v Got %v", expected, msg.Message)
	}
}
