package main

import (
	"bufio"
	"exercise/httpmetrics"
	"log"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		l, err := httpmetrics.Parse(line)
		if err != nil {
			log.Printf("Failed to parse line %v", err)
			continue
		}
		log.Printf("Parsed line %#v", l)
	}
}
