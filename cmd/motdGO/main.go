package main

import (
	"flag"
	"fmt"
	"github.com/kevinmidboe/motdGO"
	"log"
	"os"
	"strings"
)

var (
	str      *string = flag.String("str", "", "String to be converted with FIGlet")
	font     *string = flag.String("font", "", "Font name to use")
	fontpath *string = flag.String("fontpath", "", "Font path to load fonts from")
	colors   *string = flag.String("colors", "", "Character colors separated by ';'\n\tPossible colors: black, red, green, yellow, blue, magenta, cyan, white, or any hexcode (f.e. '885DBA')")
	parser   *string = flag.String("parser", "terminal", "Parser to use\tPossible parsers: terminal, html, motd")
	file     *string = flag.String("file", "", "File to write to")
)

func main() {
	// Parse the flags
	flag.Parse()

	// Validate and log the error
	validate()

	// Create objects
	ascii := motdGO.NewAsciiRender()
	options := motdGO.NewRenderOptions()

	// Load fonts
	if *fontpath != "" {
		ascii.LoadFont(*fontpath)
	}

	// Set the font
	options.FontName = *font

	// Set the parser
	p, err := motdGO.GetParser(*parser)
	if err != nil {
		p, _ = motdGO.GetParser("terminal")
	}
	options.Parser = *p

	// Set colors
	if *colors != "" {
		options.FontColor = getColorSlice(*colors)
	}

	// Render the string
	renderStr, err := ascii.RenderOpts(*str, options)
	formattedStr := strings.ReplaceAll(renderStr, "L", "_")
	if err != nil {
		log.Fatal(err)
	}

	// Write to file if given
	if *file != "" {
		// Create file
		f, err := os.Create(*file)
		defer f.Close()
		if err != nil {
			log.Fatal(err)
		}
		// Write to file
		b, err := f.WriteString(formattedStr)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Wrote %d bytes to %s\n", b, *file)
		return
	}

	// Default is printing
	fmt.Print(formattedStr)
}

// Get a slice with colors to give to the RenderOptions
// Splits the given string with the separator ";"
func getColorSlice(colorStr string) []motdGO.Color {

	givenColors := strings.Split(colorStr, ";")

	colors := make([]motdGO.Color, len(givenColors))

	for i, c := range givenColors {
		switch c {
		case "black":
			colors[i] = motdGO.ColorBlack
		case "red":
			colors[i] = motdGO.ColorRed
		case "green":
			colors[i] = motdGO.ColorGreen
		case "yellow":
			colors[i] = motdGO.ColorYellow
		case "blue":
			colors[i] = motdGO.ColorBlue
		case "magenta":
			colors[i] = motdGO.ColorMagenta
		case "cyan":
			colors[i] = motdGO.ColorCyan
		case "white":
			colors[i] = motdGO.ColorWhite
		default:
			// Try to parse the TrueColor from the string
			color, err := motdGO.NewTrueColorFromHexString(c)
			if err != nil {
				log.Fatal(err)
			}
			colors[i] = color
		}
	}

	return colors
}

// Validate if all required options are given
// flag.Parse() must be called before this
func validate() {
	if *str == "" {
		flag.Usage()
		os.Exit(1)
	}
}
