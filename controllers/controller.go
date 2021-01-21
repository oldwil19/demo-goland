package controllers

import (
	"github.com/gofiber/fiber/v2"
	"demo/utils"
)

//Satellite struct
type Satellite struct{
	//name
	Name string `json:"name"`
	
	//distance
	Distance float32 `json:"distance"`
	
	//message
	Message []string `json:"message"`
}

//Satellites struct
type Satellites struct{
	//satellites
	Satellites []Satellite `json:"satellites"`
}

//SatelliteSplit stuct
type SatelliteSplit struct {
		
	//distance
	Distance float32 `json:"distance"`
	
	//message
	Message []string `json:"message"`
}

//Position response
type Position struct {
	// X
	X float32`json:"x"`

	// Y
	Y float32`json:"y"`
}


//TopSecret Returns the location of the ship and the message
func TopSecret(c *fiber.Ctx) error {
	//Request
	Request := new(Satellites)
	if err := c.BodyParser(request); err != nil {
		return err
	}
	utils.Solved(Request)
	return c.Status(fiber.StatusOK).JSON(Request)
}

//TopSecretByOne return location of ship and message given a satellite information
func TopSecretByOne(c *fiber.Ctx) error {
	request := new(SatelliteSplit)
	if err := c.BodyParser(request); err != nil {
		return err
	}
	satelliteName := c.Params("satelliteName")
	data := [] interface{}{satelliteName,request.Distance,request.Message}
	return c.Status(fiber.StatusOK).JSON(data)
}
