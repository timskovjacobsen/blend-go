package main

import (
	"fmt"
	"github.com/gookit/color"
	"os"
	"strconv"
	"strings"
)

type RGBColor struct {
	R int
	G int
	B int
}

func RGBToHex(rgbColor RGBColor) string {
	return fmt.Sprintf("#%02x%02x%02x", rgbColor.R, rgbColor.G, rgbColor.B)
}

func HexToRGB(hex string) (RGBColor, error) {
	if !strings.HasPrefix(hex, "#") {
		errMsg := fmt.Sprintf(`couldn't convert hex value %s to RGB, 
							   expected hex value to start with a #`, hex)
		return RGBColor{}, fmt.Errorf(errMsg)
	}
	if len(hex) != 7 {
		errMsg := fmt.Sprintf(`couldn't convert hex value %s to RGB, 
							   expected # followed by 6 chars`, hex)
		return RGBColor{}, fmt.Errorf(errMsg)
	}
	// Convert hex values to decimal
	red, err := strconv.ParseInt(hex[1:3], 16, 0)
	if err != nil {
		return RGBColor{}, fmt.Errorf("error converting red hex value %s: %s", hex, err)
	}
	green, err := strconv.ParseInt(hex[3:5], 16, 0)
	if err != nil {
		return RGBColor{}, fmt.Errorf("error converting green hex value %s: %s", hex, err)
	}
	blue, err := strconv.ParseInt(hex[5:7], 16, 0)
	if err != nil {
		return RGBColor{}, fmt.Errorf("error converting blue hex value %s: %s", hex, err)
	}
	return RGBColor{R: int(red), G: int(green), B: int(blue)}, nil
}

// Blend color1 into color2 with an alpha value
//
// If alhpa=0.0, returns color2
// If alpha=1.0, returns color1
// If alpha=0.6, return color1 mixed 60% into color2
func BlendColors(color1, color2 RGBColor, alpha float64) RGBColor {
	// Ensure alpha is between 0 and 1
	if alpha < 0 {
		alpha = 0
	} else if alpha > 1 {
		alpha = 1
	}

	return RGBColor{
		R: int(float64(color1.R)*alpha + float64(color2.R)*(1-alpha)),
		G: int(float64(color1.G)*alpha + float64(color2.G)*(1-alpha)),
		B: int(float64(color1.B)*alpha + float64(color2.B)*(1-alpha)),
	}
}

func main() {
	color1, err := HexToRGB(fmt.Sprintf("#%s", os.Args[1]))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	color2, err := HexToRGB(fmt.Sprintf("#%s", os.Args[2]))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	alpha, _ := strconv.ParseFloat(os.Args[3], 8)
	c1 := color.HEX(os.Args[1], true)
	c2 := color.HEX(os.Args[2], true)
	c1Show := fmt.Sprintf("%s #%s", c1.Sprintf(" "), os.Args[1])
	c2Show := fmt.Sprintf("%s #%s", c2.Sprintf(" "), os.Args[2])
	fmt.Printf("Blending (%s) into (%s) with alpha=%f\n\n", c1Show, c2Show, alpha)
	color3 := BlendColors(color1, color2, alpha)
	c3Hex := RGBToHex(color3)
	c3 := color.HEX(c3Hex, true)

	c3Show := fmt.Sprintf("%s %s", c3.Sprintf(" "), c3Hex)

	fmt.Printf("%s\n", c1Show)
	fmt.Printf("%s\n", c2Show)
	fmt.Printf("%s\n", c3Show)
}
