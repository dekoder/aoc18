package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strconv"
)

const inputFilePath = "input.txt"

var claimFormat = regexp.MustCompile(`^#([0-9]+) @ ([0-9]+),([0-9]+): ([0-9]+)x([0-9]+)$`)

type position struct {
	x uint64
	y uint64
}

type claim struct {
	ID     string
	width  uint64
	height uint64
	x      uint64
	y      uint64
}

func parseClaims(inputPath string) []claim {
	content, err := ioutil.ReadFile(inputPath)
	if err != nil {
		log.Fatal(fmt.Sprint("Unable to open input file:", err))
	}

	lines := bytes.Split(content, []byte("\n"))

	var claims []claim
	for _, line := range lines {
		matches := claimFormat.FindAllStringSubmatch(string(line), -1)
		if matches == nil {
			continue
		}

		if len(matches[0]) != 6 {
			fmt.Printf("Wrong number of matches: Issue with line %q: %d matches\n", line, len(matches[0]))
			continue
		}

		x, err := strconv.ParseUint(matches[0][2], 10, 64)
		if err != nil {
			log.Fatal(fmt.Sprint("Unable to parse claim:", err))
		}

		y, err := strconv.ParseUint(matches[0][3], 10, 64)
		if err != nil {
			log.Fatal(fmt.Sprint("Unable to parse claim:", err))
		}

		width, err := strconv.ParseUint(matches[0][4], 10, 64)
		if err != nil {
			log.Fatal(fmt.Sprint("Unable to parse claim:", err))
		}

		height, err := strconv.ParseUint(matches[0][5], 10, 64)
		if err != nil {
			log.Fatal(fmt.Sprint("Unable to parse claim:", err))
		}

		claims = append(claims, claim{
			ID:     matches[0][1],
			x:      x,
			y:      y,
			width:  width,
			height: height,
		})
	}

	log.Printf("Successfully parsed %d claims\n", len(claims))

	return claims
}

func computeOverlap(claims []claim) uint {
	surfaceClaimed := make(map[position][]string)

	// Compute surface claimed
	for _, claim := range claims {
		for x := claim.x; x < claim.x+claim.width; x++ {
			for y := claim.y; y < claim.y+claim.height; y++ {
				surfaceClaimed[position{x: x, y: y}] = append(surfaceClaimed[position{x: x, y: y}], claim.ID)
			}
		}
	}

	// Compute overlap
	var overlap uint
	for _, IDs := range surfaceClaimed {
		if len(IDs) > 1 {
			overlap++
		}
	}

	return overlap
}

func solveExercise(inputPath string) uint {
	claims := parseClaims(inputPath)

	return computeOverlap(claims)
}

func main() {
	log.Println("Beginning day03ex01...")

	overlap := solveExercise(inputFilePath)

	log.Println("Overlapping claims successfully computed")
	log.Printf("Inches covered by overlapping claims: %d\n", overlap)
}
