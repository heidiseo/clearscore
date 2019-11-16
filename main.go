package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

//UserInfo represents information received from post request to creditcards endpoint
type UserInfo struct {
	FirstName   string `json:"firstname"`
	LastName    string `json:"lastname"`
	DOB         string `json:"dob"`
	CreditScore int    `json:"credit-score"`
	EmpStatus   string `json:"employment-status"`
	Salary      int    `json:"salary"`
}

//CreditCard represents the response of creditcards endpoint
type CreditCard struct {
	Provider  string   `json:"provider"`
	Name      string   `json:"name"`
	ApplyURL  string   `json:"apply-url"`
	Apr       float32  `json:"apr"`
	Features  []string `json:"features"`
	CardScore float32  `json:"card-score"`
}

//CreditCards contains all credit cards
type CreditCards []CreditCard

func main() {
	// handleRequest()
	userInfo := UserInfo{
		FirstName:   "John",
		LastName:    "Smith",
		DOB:         "1991/04/18",
		CreditScore: 500,
		EmpStatus:   "FULL_TIME",
		Salary:      28000,
	}
	userInfo.csCards()
	userInfo.scoredCards()
}

func allCreditCards(w http.ResponseWriter, r *http.Request) {
	creditCards := CreditCards{
		CreditCard{Provider: "Test Provider", Name: "Test NAme", ApplyURL: "http://www.example.com/apply", Apr: 34.3, Features: []string{"supports ApplyPay", "free interest for 10 years"}, CardScore: 0.2},
	}
	fmt.Println("Endpoint Hit: All Credit Cards Endpoint")
	json.NewEncoder(w).Encode(creditCards)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Homepage Endpoint Hit")
}

// func handleRequest() {
// 	router := mux.NewRouter().StrictSlash(true)
// 	router.HandleFunc("/", homePage)
// 	router.HandleFunc("/creditcard", getUserInfo).Methods("POST")
// 	router.HandleFunc("/creditcards", allCreditCards).Methods("GET")
// 	log.Fatal(http.ListenAndServe(":8081", router))
// }

func getUserInfo(w http.ResponseWriter, r *http.Request) *UserInfo {
	var newUserInfo UserInfo
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "please enter user info")
	}
	json.Unmarshal(reqBody, &newUserInfo)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(newUserInfo)
	return &newUserInfo
}

func (userInfo *UserInfo) csCards() {
	url := "https://y4xvbk1ki5.execute-api.us-west-2.amazonaws.com/CS/v1/cards"
	fmt.Println("URL:>", url)

	var jsonStr = []byte(fmt.Sprintf(`{
		"fullName": "%s %s",
		"dateOfBirth": "%s",
		"creditScore": %d
	}`, userInfo.FirstName, userInfo.LastName, userInfo.DOB, userInfo.CreditScore))
	myString := string(jsonStr)
	fmt.Printf("XXXXXXX: %v\n", myString)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))
}

func (userInfo *UserInfo) scoredCards() {
	url := "https://m33dnjs979.execute-api.us-west-2.amazonaws.com/CS/v2/creditcards"
	fmt.Println("URL:>", url)

	var jsonStr = []byte(fmt.Sprintf(`{
		"first-name": "%s",
		"last-name": "%s",
		"date-of-birth": "%s",
		"score": %d,
		"employment-status": "%s",
		"salary": %d
	}`, userInfo.FirstName, userInfo.LastName, userInfo.DOB, userInfo.CreditScore, userInfo.EmpStatus, userInfo.Salary))
	myString := string(jsonStr)
	fmt.Printf("XXXXXXX: %v\n", myString)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))
}
