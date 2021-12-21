package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"log"
	"strconv"
)

type FizzBuzz struct{
	Fizzbuzz  []string `json:"fizzBuzz"`
}
// a utility for better reading and understanding

func isFullyDivisible(number int, divider int) bool{
	return (number % divider == 0)
}

func fizzBuzzHandler(c*fiber.Ctx) error {
	count, err := strconv.Atoi(c.Params("count")) //convert string count to int to for in.
	if err != nil {
		return err
	}
	var response FizzBuzz
	for i:= 1 ; i <= count; i++ {
		if !isFullyDivisible(i,3) && !isFullyDivisible(i,5){
			response.Fizzbuzz = append(response.Fizzbuzz,strconv.Itoa(i))
		} else if isFullyDivisible(i,3) && !isFullyDivisible(i,5){
			response.Fizzbuzz = append(response.Fizzbuzz,"fizz")
		} else if !isFullyDivisible(i,3) && isFullyDivisible(i,5){
			response.Fizzbuzz = append(response.Fizzbuzz,"buzz")
		} else if isFullyDivisible(i,3) && isFullyDivisible(i,5){
			response.Fizzbuzz = append(response.Fizzbuzz,"fizzbuzz")
		}
	}
	return c.JSON(response)
}
func StartServer(port int ,UseDefaultCors bool) error{
	app:= fiber.New()
	if UseDefaultCors{
		app.Use(cors.New())
	}
	app.Get("/fizzbuzz/:count",fizzBuzzHandler)
	err := app.Listen(fmt.Sprintf(":%d",port))
	return err
}
func main() {
	err := StartServer(4000,true)
	if err != nil {
		log.Fatal(err)
	}
}
