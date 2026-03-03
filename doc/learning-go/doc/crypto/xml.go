package main

import (
	"bufio"
	"encoding/xml"
	"fmt"
	"log"
	"os"
)

func XMLDemo() {
	var file, err = os.Open("./testdata/Person.xml")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var decoder = xml.NewDecoder(bufio.NewReader(file))
	for token, err := decoder.Token(); err == nil; token, err = decoder.Token() {
		switch token := token.(type) {
		case xml.StartElement:
			name := token.Name.Local
			fmt.Printf("Token name: %s\n", name)
			for _, attr := range token.Attr {
				attrName := attr.Name.Local
				attrValue := attr.Value
				fmt.Printf("An attribute is: %s %s\n", attrName, attrValue)
			}
		case xml.EndElement:
			fmt.Println("End of token")
		case xml.CharData:
			content := string([]byte(token))
			fmt.Printf("This is the content: %v\n", content)
		}
	}
}
