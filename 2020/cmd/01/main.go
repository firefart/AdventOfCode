package main

import (
	"strconv"

	"github.com/firefart/adventofcode/internal"
	log "github.com/sirupsen/logrus"
)

func main() {
	input, err := internal.ReadFile("cmd/01/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	l := len(input)

	for i := 0; i < l; i++ {
		for j := 0; j < l; j++ {
			i2, err := strconv.Atoi(input[i])
			if err != nil {
				log.Fatal(err)
			}
			j2, err := strconv.Atoi(input[j])
			if err != nil {
				log.Fatal(err)
			}
			if i2+j2 == 2020 {
				log.Infof("Found solution 1 match! %d + %d = 2020, Solution: %d", i2, j2, i2*j2)
			}

			for k := 0; k < l; k++ {
				k2, err := strconv.Atoi(input[k])
				if err != nil {
					log.Fatal(err)
				}
				if i2+j2+k2 == 2020 {
					log.Infof("Found solution 2 match! %d + %d + %d = 2020, Solution: %d", i2, j2, k2, i2*j2*k2)
				}
			}
		}
	}
}
