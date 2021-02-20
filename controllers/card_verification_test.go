package controllers

import (
	"testing"
	"fmt"
	"net/http"
	"net/http/httptest"
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/assert"
)


func TestVerifyCreditCard(t *testing.T) {

	samples := []struct {
		inputJSON   string
		statusCode  int
		message		string
	}{
		{
			inputJSON:    	`{"Name":"Samuel Grumley", "CardNumber":"4111111111111111","Expiry":"07/24","CVC":789}`,
			statusCode:   	200,
			message:       	"VISA: 4111111111111111 (valid)",
		},
		{
			inputJSON:    	`{"Name":"Samuel Grumley", "CardNumber":"4111111111111","Expiry":"07/24","CVC":789}`,
			statusCode:   	200,
			message:       	"VISA: 4111111111111 (invalid)",
		},
		{
			inputJSON:    	`{"Name":"Samuel Grumley", "CardNumber":"4012888888881881","Expiry":"07/24","CVC":789}`,
			statusCode:   	200,
			message:       	"VISA: 4012888888881881 (valid)",
		},
		{
			inputJSON:    	`{"Name":"Samuel Grumley", "CardNumber":"378282246310005","Expiry":"07/24","CVC":789}`,
			statusCode:   	200,
			message:       	"AMEX: 378282246310005 (valid)",
		},
		{
			inputJSON:    	`{"Name":"Samuel Grumley", "CardNumber":"6011111111111117","Expiry":"07/24","CVC":789}`,
			statusCode:   	200,
			message:       	"Discover: 6011111111111117 (valid)",
		},
		{
			inputJSON:    	`{"Name":"Samuel Grumley", "CardNumber":"5105105105105100","Expiry":"07/24","CVC":789}`,
			statusCode:   	200,
			message:       	"MasterCard: 5105105105105100 (valid)",
		},
		{
			inputJSON:    	`{"Name":"Samuel Grumley", "CardNumber":"5105 1051 0510 5106","Expiry":"07/24","CVC":789}`,
			statusCode:   	200,
			message:       	"MasterCard: 5105 1051 0510 5106 (invalid)",
		},
		{
			inputJSON:    	`{"Name":"Samuel Grumley", "CardNumber":"9111111111111111","Expiry":"07/24","CVC":789}`,
			statusCode:   	200,
			message:       	"Unknown: 9111111111111111 (invalid)",
		},
		{
			inputJSON:    	`{"Name":"Samuel Grumley", "CardNumber":"9111111111111111","Expiry":"07/24","CVC":789, "Age":25}`,
			statusCode:   	400,
			message:       	"Invalid JSON request body",
		},
	}	
	for _, v := range samples {	
		req, err := http.NewRequest("POST", "/api/VerifyCreditCard", bytes.NewBufferString(v.inputJSON))
		if err != nil {
			t.Errorf("this is the error: %v\n", err)
		}
		// implements responseWriter to find the actual http status code
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(VerifyCreditCard)
		handler.ServeHTTP(rr, req)		
		responseMap := make(map[string]interface{})
		err = json.Unmarshal([]byte(rr.Body.String()), &responseMap)
		if err != nil {
			fmt.Printf("Cannot convert to json: %v", err)
		}

		assert.Equal(t, rr.Code, v.statusCode)
		// check to see if we have correctly matched message
		if v.statusCode == 200 {
			assert.Equal(t, responseMap["message"], v.message)
		}
		if v.statusCode == 400 {
			assert.Equal(t, responseMap["message"], v.message)
		}
	}
}

