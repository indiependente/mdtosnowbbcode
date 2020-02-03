package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/indiependente/mdtosnowbbcode/converter"
)

func main() {
	err := run()
	if err != nil {
		log.Fatalf("Error during conversion: %v", err)
	}

}

func run() error {
	conv := converter.MDToSnow{}
	r, err := conv.Convert(os.Stdin)
	if err != nil {
		return fmt.Errorf("could not convert: %w", err)
	}
	bbcode, err := ioutil.ReadAll(r)
	if err != nil {
		return fmt.Errorf("could not read converted content: %w", err)
	}
	_, err = os.Stdout.Write(bbcode)
	if err != nil {
		return fmt.Errorf("could not write converted content: %w", err)
	}
	fmt.Println()
	return nil
}
