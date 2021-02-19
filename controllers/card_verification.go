package controllers 

import (
	"net/http"
	"fmt"
	"encoding/json"
	u "github.com/sgrumley/CheckingCreditCards/utils"
)


type CreditCardVerification interface {
	lengthValidate()	error
	luhnValidate()		error
}


type CardProvider string
const (
	AMEX CardProvider 	= "AMEX"
	Discover 			= "Discover"
	MasterCard 			= "MasterCard"
	Visa 				= "Visa"
	Unknown 			= "Unknown"
)


type CreditCard struct {
	Name 		string
	CardNumber 	string 
	Expiry 		string
	CVC 		int
	Provider 	CardProvider
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
		u.Respond(w, http.StatusBadRequest, u.Message(false, "Invalid JSON request body"))
		return
	}
	
	response := verify(creditCard)
	fmt.Println(creditCard)

	if response["status"].(bool) == false {
		u.Respond(w, http.StatusBadRequest, response)
	} else {
		u.Respond(w, http.StatusOK, response)
	}

	return
}


func verify(ccv CreditCardVerification) map[string]interface{}{
	ccv.lengthValidate()
	ccv.luhnValidate()
	
	return u.Message(true, "Success")
}

func (creditCard *CreditCard) lengthValidate() error {
	return nil
}


func (creditCard *CreditCard) luhnValidate() error {
	return nil 
}

