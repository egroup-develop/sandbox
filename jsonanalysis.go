package main

import (
	"encoding/json"
	"fmt"
	//"log"
	"os"
)

type image2 struct {
	Height	int		`json:"height"`
	Link	string		`json:"link"`
	Title	string		`json:"title"`
	Url	string		`json:"url"`
	Width	int		`json:"width"`
}

type provider struct {
	Link	string		`json:"link"`
	Name	string		`json:"name"`
}

type copy struct {
	Image		image2	`json:"image"`
	Link 		string	`json:"link"`
	Provider	[]provider
	Title		string	`json:"title"`
}

type des struct {
	PublicTime	string	`json:"publicTime"`
	Text		string	`json:"text"`
}

type images struct {
	heiget	int		`json:"height"`
	Title	string		`json:"title"`
	Url	string		`json:"url"`
	Width	int		`json:"width"`
}

type min struct {
        Celsius         string  `json:"celsius"`
        Fahrenheit      string  `json:"fahrenheit"`
}


type max struct {
	Celsius		int	`json:"celsius"`
	Fahrenheit	float64	`json:"fahrenheit"`
}

type temp struct {
	Max	max		`json:"max"`
	Min	min		`json:"min"`
}

type forecasts struct {
	Date	string		`json:"date"`
	DateLabel string	`json:"dateLabel"`
	Image 	images		`json:"image"`
	Telop	string		`json:"telop"`
	Temperature temp	`json:"temperature"`

}

type locate struct {
	Area	string		`json:"area"`
	City	string		`json:"city"`
	Prefecture string	`json:"prefecture"`
}

type point struct {
	Link	string	`json:"link"`
	Name	string	`json:"name"`
}

type data struct {
	Copyright	copy	`json:"copyright"`
	Description	des	`json:"description"`
	Forecasts	[]forecasts
	Link		string	`json:"link"`
	Location	locate	`json:"location"`
	PinpointLocations []point
	PublicTime	string	
	Title		string
} 

func main() {
    dec := json.NewDecoder(os.Stdin)
    var d data
    dec.Decode(&d)
    fmt.Printf("%+v\n", d)
}
