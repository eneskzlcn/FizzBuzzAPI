package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/stretchr/testify/assert" // add Testify package
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)


func TestFizzBuzzApi(t *testing.T) {
	tests := []struct {
	description  string // description of the test case
	route        string // route path to test
	expectedCode int    // expected HTTP status code
	value        FizzBuzz
	}{
		{
			description:  "get http status 200, when succesfully get ordinal ascending numbers array by count 2",
			route:        "/fizzbuzz/2",
			expectedCode: http.StatusOK,
			value:        FizzBuzz{
				Fizzbuzz: []string{"1","2"},
			},
		},
		{
			description:  "get http status 200, when succesfully a 'fizz' on number '3' in array by count 3",
			route:        "/fizzbuzz/3",
			expectedCode: http.StatusOK,
			value:        FizzBuzz{
				Fizzbuzz: []string{"1","2","fizz"},
			},
		},
		{
			description:  "get http status 200, when succesfully get a 'fizz' on number '3' and get a 'buzz' on number '5' by count 5",
			route:        "/fizzbuzz/5",
			expectedCode: http.StatusOK,
			value:        FizzBuzz{
				Fizzbuzz: []string{"1","2","fizz","4","buzz"},
			},
		},
		{
			description:  "get http status 200, when succesfully get 'fizz' on numbers divisible by 3,get 'buzz' on numbers divisible by 5, get 'fizzbuzz' on numbers divisible by both of 3 and 5  by count 15",
			route:        "/fizzbuzz/5",
			expectedCode: http.StatusOK,
			value:        FizzBuzz{
				Fizzbuzz: []string{"1","2","fizz","4","buzz","fizz","7","8","fizz","buzz","11","fizz","13","14","fizzbuzz"},
			},
		},

	}
	app:= fiber.New()
	app.Use(cors.New())
	app.Get("/fizzbuzz/:count", func(c *fiber.Ctx) error {
		count, err := strconv.Atoi(c.Params("count")) //convert string count to int to for in.
		if err != nil{
			t.Fatal(err)
		}

		var numbers []string
		for i:= 1 ; i <= count; i++ {

			if !isFullyDivisible(i,3) && !isFullyDivisible(i,5){
				numbers = append(numbers,strconv.Itoa(i))
			} else if isFullyDivisible(i,3) && !isFullyDivisible(i,5){
				numbers = append(numbers,"fizz")
			} else if !isFullyDivisible(i,3) && isFullyDivisible(i,5){
				numbers = append(numbers,"buzz")
			} else if isFullyDivisible(i,3) && isFullyDivisible(i,5){
				numbers = append(numbers,"fizzbuzz")
			}
		}

		var response = FizzBuzz{Fizzbuzz: numbers}
		return c.JSON(response)
	})

	for _, test := range tests {
		req := httptest.NewRequest("GET",test.route,nil)
		resp,_ := app.Test(req,1)

		assert.Equalf(t, test.expectedCode,resp.StatusCode,test.description)
	}

}
