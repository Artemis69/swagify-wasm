package main

import (
	"math/rand"
	"strings"
	"syscall/js"
	"time"
)

type Config struct {
	LetterReplacementChange int
	UpperCaseChance         int
	TripleChance            int
	MaxTags                 int
}

func reverse(s string) string {
	rns := []rune(s)

	for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 {
		rns[i], rns[j] = rns[j], rns[i]
	}

	return string(rns)
}

func randInt(min int, max int) int {
	return rand.Intn(max-min+1) + min
}

func replaceLetters(str string, config *Config) string {

	chars := strings.Split(str, "")

	for i, char := range chars {

		if randInt(0, 20) < config.LetterReplacementChange {
			continue
		}

		switch char {
		case "S":
			chars[i] = "$"
		case "s":
			chars[i] = "z"
		case "l", "i", "I":
			chars[i] = "1"
		case "E", "e":
			chars[i] = "3"
		case "A":
			chars[i] = "4"
		case "a":
			chars[i] = "@"
		case "0":
			chars[i] = "0"
		case "H":
			chars[i] = "#"
		case "z":
			chars[i] = "zzz"
		case "Z":
			chars[i] = "ZZZ"
		case "g":
			chars[i] = "ggg"
		case "G":
			chars[i] = "GGG"
		case "t":
			chars[i] = "+"
		case "D":
			chars[i] = "|)"
		default:
		}
	}

	return strings.Join(chars, "")
}

func randomizeCase(str string, config *Config) string {

	chars := strings.Split(str, "")

	for i, char := range chars {

		if randInt(0, 20) < config.UpperCaseChance {
			chars[i] = strings.ToUpper(char)
		}

	}

	return strings.Join(chars, "")
}

func tripleLetters(str string, config *Config) string {

	chars := strings.Split(str, "")

	for i, char := range chars {

		if randInt(0, 20) < config.TripleChance {
			chars[i] = strings.Repeat(char, 3)
		}

	}

	return strings.Join(chars, "")
}

func decorate(str string) string {

	decorations := [22]string{
		"x",
		"X",
		"xX",
		"xxx",
		"~",
		".-~",
		"xXx",
		"XxX",
		"xxX_",
		"|",
		"./|",
		"@@@",
		"$$$",
		"***",
		"+",
		"|420|",
		".::",
		".:",
		".-.",
		"|||",
		"--",
		"*--",
	}

	decoration := decorations[randInt(0, len(decorations) - 1)]

	return decoration + str + reverse(decoration)
}

func addTags(str string, config *Config) string {

	tags := [44]string{
		"SHOTS FIRED",
		"420",
		"LEGIT",
		"360",
		"Pr0",
		"NO$$$cop3z",
		"0SC0pe",
		"MLG",
		"h4xx0r",
		"M4X$W4G",
		"L3G1TZ",
		"3edgy5u",
		"2edgy4u",
		"nedgy(n+2)u",
		"s0b4s3d",
		"SWEG",
		"LEGIT",
		"WUBWUBWUB",
		"BLAZEIT",
		"b14Z3d",
		"[le]G1t",
		"60x7",
		"24x7BLAZEIT",
		"4.2*10^2",
		"literally",
		"[le]terally",
		"1337",
		"l33t",
		"31337",
		"Tr1Ck$h0t",
		"SCRUBLORD",
		"DR0PTH3B4$$",
		"w33d",
		"ev REE DAI",
		"MTNDEW",
		"WATCH OUT",
		"EDGY",
		"ACE DETECTIVE",
		"90s KID",
		"NO REGRETS",
		"THANKS OBAMA",
		"SAMPLE TEXT",
		"FAZE",
		"#nofilter",
	}

	tagsCount := randInt(0, config.MaxTags)

	for i := 0; i < tagsCount; i++ {
		str = "[" + tags[randInt(0, len(tags) - 1)] + "]" + str
	}

	return str

}

func swagify(this js.Value, args []js.Value) interface{} {

	if args == nil || len(args) < 1 {
		return nil
	}

	var letterReplacementChange int = 8
	var upperCaseChance int = 5
	var tripleChance int = 1
	var maxTags int = 3

	if len(args) == 2 && args[1].Type() == 6 {

		if args[1].Get("letterReplacementChange").Type() == 3 {
			letterReplacementChange = args[1].Get("letterReplacementChange").Int()
		}

		if args[1].Get("upperCaseChance").Type() == 3 {
			upperCaseChance = args[1].Get("upperCaseChance").Int()
		}

		if args[1].Get("tripleChance").Type() == 3 {
			tripleChance = args[1].Get("tripleChance").Int()
		}

		if args[1].Get("maxTags").Type() == 3 {
			maxTags = args[1].Get("maxTags").Int()
		}
		
	}

	var str string = ""

	if args[0].Type() == 4 {
		str = args[0].String()
	}

	config := &Config{LetterReplacementChange: letterReplacementChange, UpperCaseChance: upperCaseChance, TripleChance: tripleChance, MaxTags: maxTags}

	str = replaceLetters(str, config)

	str = randomizeCase(str, config)

	str = tripleLetters(str, config)

	str = decorate(str)

	str = addTags(str, config)

	return js.ValueOf(str)
}

func init() {
    rand.Seed(time.Now().UTC().UnixNano())
}

func main() {

	c := make(chan struct{}, 0)

	js.Global().Set("swagify", js.FuncOf(swagify))

	<-c
}
