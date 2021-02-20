package controllers 

import (
	"net/http"
	"fmt"
	"encoding/json"
	"regexp"
	"strings"
	u "github.com/sgrumley/CheckingCreditCards/utils"
)


type CreditCardVerification interface {
	lengthValidate()	error
	luhnValidate()		error
	resultBuilder() 	string
}


type CardProvider string
const (
	AMEX CardProvider 	= "AMEX"
	Discover 			= "Discover"
	MasterCard 			= "MasterCard"
	Visa 				= "VISA"
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


// Decodes incoming body into struct
func DecodeJsonCreditCard(r *http.Request) (*CreditCard, error){
	creditCard := &CreditCard{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	err := d.Decode(&creditCard); if err != nil {
		return nil, err
	}
	
	return creditCard, nil
}


// Handles the whole verification process
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


// Implements the interface to ensure all functions are implemented
func verify(ccv CreditCardVerification) map[string]interface{}{
	ccv.lengthValidate()
	ccv.luhnValidate()
	
	return u.Message(true, ccv.resultBuilder())
}


// Validates the card provider using starting number conditions and lengths
func (creditCard *CreditCard) lengthValidate() error {
	cardNumber := strings.ReplaceAll(creditCard.CardNumber, " ", "")
	// starts with 4 and ensures all values are numerical
	var visa 		= regexp.MustCompile(`4[0-9]{12,15}`)
	// starts with 51,52,53,54 or 55 and ensures all values are numerical
	var masterCard 	= regexp.MustCompile(`5[1-5][0-9]{14}`)
	// starts with 6011 and ensures all values are numerical
	var discover 	= regexp.MustCompile(`6011[0-9]{12}`)
	// starts with 34 or 37 and ensures all values are numerical
	var amex 		= regexp.MustCompile(`3[47][0-9]{13}`)

	// match regex conditions and catch lengths
	switch {
	case visa.MatchString(cardNumber) && (len(cardNumber) == 13 || len(cardNumber) == 16):	
		creditCard.Provider = Visa
		
	case masterCard.MatchString(cardNumber) && len(cardNumber) == 16: 
		creditCard.Provider = MasterCard
	
	case discover.MatchString(cardNumber) && len(cardNumber) == 16:
		creditCard.Provider = Discover

	case amex.MatchString(cardNumber) && len(cardNumber) == 15:
		creditCard.Provider = AMEX

	default:
		creditCard.Provider = Unknown
	}

	return nil
}


// sums indivual digits e.g 107 = 1 + 0 + 7
func charSum(digit int) int {
	sum := 0

	for digit > 0 {
		sum 	+= digit % 10
		digit 	= digit / 10
	}

	return sum
}


/*
	1. Starting with the next to last digit and continuing with every other digit going back to the
	beginning of the card, double the digit.

	2. Sum all doubled and untouched digits in the number. For digits greater than 9 you will need to
	split them and sum them independently (i.e. "10", 1 + 0).

	3. If that total is a multiple of 10, the number is valid.
*/
func (creditCard *CreditCard) luhnValidate() error {
	cardNumber 	:= strings.ReplaceAll(creditCard.CardNumber, " ", "")
	sum, count 		:= 0, 1

	// iterate backwards through cc num
	for i := len(cardNumber) - 1; i >= 0; i-- {
		val 	:= int(cardNumber[i]) - 48

		// Iterating over every second value of the credit card number
		if count % 2 == 0 {
			sum += charSum(val * 2)
		} else {
			sum += val
		}

		count++
	}

	if sum % 10 == 0 {
		creditCard.Valid = "valid"
	} else {
		creditCard.Valid = "invalid"
	}
	
	return nil 
}

// Builds the response string
func (creditCard *CreditCard) resultBuilder() string {
	return fmt.Sprintf("%v: %v (%v)", creditCard.Provider, creditCard.CardNumber, creditCard.Valid)
}
