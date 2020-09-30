package main

import (
	"encoding/json"
	"fmt"
	"github.com/lucasb-eyer/go-colorful"
)

type hexColor struct {
	Name string `json:"name"`
	Hex  string `json:"hex"`
}

// GetColorName looks up the name for a specific hex color code.
func GetColorName(myHexColorCode string) string {
	var colorData []hexColor

	// allColorsJSONString is defined in `all-colors.go` (generated by gen.go).
	err := json.Unmarshal([]byte(allColorsJSONString), &colorData)

	if err != nil {
		fmt.Println("error:", err)
	}

	myHashedHexColorCode := fmt.Sprintf("#%s", myHexColorCode)
	myReferenceColor, err := colorful.Hex(myHashedHexColorCode)

	if err != nil {
		fmt.Println("error:", err)
	}

	var closestColorIndex int
	var closestColorDistance float64 = 1.0 // Default to 1.0, so subsequent colors will be closer.

	for i := range colorData {
		// Exit early if there's an exact hex code match.
		if colorData[i].Hex == myHashedHexColorCode {
			fmt.Println("Exact match for index:", i)
			closestColorIndex = i
			break
		}

		// Otherwise, we do a color distance comparison.
		currentColor, err := colorful.Hex(colorData[i].Hex)

		if err != nil {
			fmt.Println("error:", err)
		}

		// We use the LAB color space to find the closest color based on human perceptions.
		currentColorDistance := myReferenceColor.DistanceLab(currentColor)

		if currentColorDistance < closestColorDistance {
			closestColorDistance = currentColorDistance
			closestColorIndex = i
		}
	}

	return colorData[closestColorIndex].Name
}
