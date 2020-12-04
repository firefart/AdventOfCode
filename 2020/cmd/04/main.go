package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/firefart/adventofcode/internal"
	log "github.com/sirupsen/logrus"
)

// byr (Birth Year)
// iyr (Issue Year)
// eyr (Expiration Year)
// hgt (Height)
// hcl (Hair Color)
// ecl (Eye Color)
// pid (Passport ID)
// cid (Country ID)

type Passport struct {
	BirthYear      string
	IssueYear      string
	ExpirationYear string
	Height         string
	HairColor      string
	EyeColor       string
	PassportID     string
	CountryID      string
}

func (p Passport) ValidPart1() bool {
	if p.BirthYear != "" && p.ExpirationYear != "" && p.EyeColor != "" && p.HairColor != "" && p.Height != "" && p.IssueYear != "" && p.PassportID != "" {
		return true
	}
	return false
}

func (p Passport) ValidPart2() bool {
	// byr (Birth Year) - four digits; at least 1920 and at most 2002.
	// iyr (Issue Year) - four digits; at least 2010 and at most 2020.
	// eyr (Expiration Year) - four digits; at least 2020 and at most 2030.
	// hgt (Height) - a number followed by either cm or in:
	// If cm, the number must be at least 150 and at most 193.
	// If in, the number must be at least 59 and at most 76.
	// hcl (Hair Color) - a # followed by exactly six characters 0-9 or a-f.
	// ecl (Eye Color) - exactly one of: amb blu brn gry grn hzl oth.
	// pid (Passport ID) - a nine-digit number, including leading zeroes.
	// cid (Country ID) - ignored, missing or not.
	if p.BirthYear != "" {
		byr, err := strconv.Atoi(p.BirthYear)
		if err != nil {
			log.Errorf("Invalid birthyear %v", err)
			return false
		}
		if byr < 1920 || byr > 2002 {
			// log.Infof("byr invalid: %d", byr)
			return false
		}
	} else {
		return false
	}

	if p.IssueYear != "" {
		iyr, err := strconv.Atoi(p.IssueYear)
		if err != nil {
			log.Errorf("Invalid issueyear %v", err)
			return false
		}
		if iyr < 2010 || iyr > 2020 {
			// log.Infof("iyr invalid: %d", iyr)
			return false
		}
	} else {
		return false
	}

	if p.ExpirationYear != "" {
		eyr, err := strconv.Atoi(p.ExpirationYear)
		if err != nil {
			log.Errorf("Invalid expirationyear %v", err)
			return false
		}
		if eyr < 2020 || eyr > 2030 {
			// log.Infof("eyr invalid: %d", eyr)
			return false
		}
	} else {
		return false
	}

	if p.Height != "" {
		re := regexp.MustCompile(`^(\d+)(cm|in)$`)
		match := re.FindStringSubmatch(p.Height)
		if match == nil {
			return false
		}
		h, err := strconv.Atoi(match[1])
		if err != nil {
			log.Errorf("Invalid height %v", err)
			return false
		}

		switch match[2] {
		case "cm":
			if h < 150 || h > 193 {
				return false
			}
		case "in":
			if h < 59 || h > 76 {
				return false
			}
		default:
			log.Fatalf("Invalid height: %s", match[2])
		}
	} else {
		return false
	}

	if p.HairColor != "" {
		re := regexp.MustCompile(`^#[0-=9a-f]{6}$`)
		match := re.FindStringSubmatch(p.HairColor)
		if match == nil {
			// log.Infof("haircolor invalid: %s", p.HairColor)
			return false
		}
	} else {
		return false
	}

	if p.EyeColor != "" {
		if p.EyeColor != "amb" && p.EyeColor != "blu" && p.EyeColor != "brn" && p.EyeColor != "gry" && p.EyeColor != "grn" && p.EyeColor != "hzl" && p.EyeColor != "oth" {
			// log.Infof("eyecolor invalid: %s", p.EyeColor)
			return false
		}
	} else {
		return false
	}

	// pid (Passport ID) - a nine-digit number, including leading zeroes.
	if p.PassportID != "" {
		re := regexp.MustCompile(`^\d{9}$`)
		match := re.FindStringSubmatch(p.PassportID)
		if match == nil {
			return false
		}
	} else {
		return false
	}

	return true
}

func parsePassports(passports []string) ([]Passport, error) {
	var ret []Passport
	for _, passport := range passports {
		var p Passport
		fields := strings.Split(passport, " ")
		for _, field := range fields {
			kv := strings.Split(field, ":")
			key := kv[0]
			value := kv[1]
			switch key {
			case "byr":
				p.BirthYear = value
			case "iyr":
				p.IssueYear = value
			case "eyr":
				p.ExpirationYear = value
			case "hgt":
				p.Height = value
			case "hcl":
				p.HairColor = value
			case "ecl":
				p.EyeColor = value
			case "pid":
				p.PassportID = value
			case "cid":
				p.CountryID = value
			default:
				return nil, fmt.Errorf("invalid field %s", key)
			}
		}
		ret = append(ret, p)
	}

	return ret, nil
}

func main() {
	input, err := internal.ReadFile("cmd/04/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	var passports []string
	currentPassport := ""
	for _, line := range input {
		line = strings.TrimSpace(line)
		if line == "" {
			// new passport
			passports = append(passports, strings.TrimSpace(currentPassport))
			currentPassport = ""
			continue
		}

		currentPassport += line
		currentPassport += " "
	}

	log.Infof("Got %d passports", len(passports))
	ps, err := parsePassports(passports)
	if err != nil {
		log.Fatal(err)
	}

	valid := 0
	for _, p := range ps {
		if p.ValidPart1() {
			valid += 1
		}
	}
	log.Infof("Part1: %d valid passports", valid)

	valid = 0
	for _, p := range ps {
		if p.ValidPart2() {
			valid += 1
		}
	}
	log.Infof("Part2: %d valid passports", valid)
}
