package controllers 

import (
	"net/http"
	"fmt"
	"encoding/json"
)


type CreditCard struct {
	Name 		string
	CardNumber 	string 
	Expiry 		string
	CVC 		int
	Valid		string	
}

func DecodeJsonCreditCard(r *http.Request) (*CreditCard, error){
	creditCard := &CreditCard{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	err := d.Decode(&creditCard); if err != nil {
		return nil, err
	}
	
	return creditCard, nil
}


func VerifyCreditCard(w http.ResponseWriter, r *http.Request) {
	creditCard, err := DecodeJsonCreditCard(r); if err != nil {
		return
	}
	
	fmt.Println(creditCard)

}

func (creditCard *CreditCard) lengthValidate() error {
	return nil
}


func (creditCard *CreditCard) luhnValidate() error {
	return nil 
}

