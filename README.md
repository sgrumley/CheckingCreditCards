# Checking Credit Cards
A small API to verify a credit card number written in Golang.

### Server Endpoints
|Method|Endpoint|Purpose|Authentication Required|Fields|
|---|---|---|---|---|
POST| `/api/VerifyCreditCard`|Checks the credit card number is valid and checks which provider the card belongs to|No|Name (string), CardNumber (string), Expiry (string), CVC (int)  

### Features
- Uses the [Luhn Algorithm](https://en.wikipedia.org/wiki/Luhn_algorithm) to validate the card number
- Finds and verifies the card number with the provider of the card (Visa, AMEX, MasterCard, Discover)

### How To Run
in the root directory run
```
go run main.go
```
For an easy way to use the API, I would reccomend using Postman. Make a POST request to  [http://localhost:8080](http://localhost:8080). In the body tab select 'raw' then on the dropdown to the right select JSON. Below is an example of a valid input
```
{
    "Name" :		"Samuel Grumley",
	"CardNumber" : 	"4408 0412 3456 7893",
	"Expiry" : 		"07/24",
	"CVC" :		789
}
```
### How To Test
in `CheckingCreditCards/controllers` run
```
go test
```