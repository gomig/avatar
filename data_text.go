package avatar

import (
	"strings"
	"unicode/utf8"
)

var lettersMap map[string]string
var letters map[string]string

func init() {
	letters = map[string]string{
		Default: `<path fill="{text}" d="M56.06,73.83v-14.55l2.75,.08c5.35,0,8.41-1.81,8.41-7.63,0-2.2-.39-6.68-3.46-6.68s-3.38,4.4-3.38,6.76h-15.41v-1.18c-.24-11.17,7.71-18.79,18.63-18.79s19.42,8.02,19.42,19.34c0,7.47-3.54,14.7-11.72,16.51v6.13h-15.25Zm-1.1,13.21c0-4.8,4.01-8.88,8.65-8.88,5.11,0,9.36,3.85,9.36,9.12,0,4.95-4.09,8.89-9.04,8.89s-8.96-4.09-8.96-9.12Z"/>`,
		"a":     `<path fill="{text}" d="M70.6,93.64l-1.18-7.63h-10.77l-1.34,7.63h-16.51l14.23-59.29h19.58l12.58,59.29h-16.59Zm-3.07-20.21l-1.73-13.76c-.47-3.7-.79-7.47-1.18-11.17h-.16c-.47,3.7-.87,7.47-1.49,11.17l-2.2,13.76h6.76Z"/>`,
		"b":     `<path fill="{text}" d="M43.87,34.36h18.87c9.44,0,20.44,4.09,20.44,16.91,0,6.68-4.32,11.87-11.09,11.87v.16c7.86,.63,12.03,5.5,12.03,13.45,0,7.31-4.4,16.91-18.71,16.91h-21.55V34.36Zm16.98,23.04c3.3,0,6.37-1.34,6.37-5.19,0-3.54-2.44-5.03-5.74-5.03h-1.57v10.22h.94Zm.39,23.43c3.54,0,6.68-1.73,6.68-5.74s-3.07-5.66-6.76-5.66h-1.26v11.4h1.34Z"/>`,
		"c":     `<path fill="{text}" d="M80.75,51.73c-2.28-2.2-4.17-3.46-7.47-3.46-8.33,0-10.38,9.2-10.38,15.73,0,6.92,2.04,15.73,10.06,15.73,3.62,0,5.9-1.42,8.49-3.77l-.94,16.59c-3.22,1.42-6.13,2.12-9.67,2.12-13.13,0-24.3-9.99-24.3-30.27,0-26.58,18.24-31.06,26.03-31.06,2.83,0,5.58,.39,8.18,1.42v16.98Z"/>`,
		"d":     `<path fill="{text}" d="M43.52,34.36h12.89c18.24,0,28.07,11.48,28.07,29.41,0,27.36-17.22,29.88-26.97,29.88h-14V34.36Zm16.04,44.98c7.31-.08,8.57-9.2,8.57-14.78,0-6.45-1.26-15.41-8.57-15.41v30.19Z"/>`,
		"e":     `<path fill="{text}" d="M49.1,34.36h29.8v14h-13.76v8.26h12.34v14h-12.34v9.04h13.76v14h-29.8V34.36Z"/>`,
		"f":     `<path fill="{text}" d="M49.85,93.64V34.36h28.31v14h-12.27v8.57h11.48v14h-11.48v22.72h-16.04Z"/>`,
		"g":     `<path fill="{text}" d="M86.17,59.36l.08,2.52c0,14.63-2.83,32.79-21.39,32.79-17.46,0-23.12-15.33-23.12-30.35s5.58-30.98,22.96-30.98c9.59,0,16.75,6.05,20.21,14.63l-15.49,5.74c-.16-2.36-1.73-6.05-4.56-6.05-7.39,0-6.76,13.6-6.76,18.09,.24,4.25-.08,14.63,6.13,14.63,4.25,0,5.58-5.03,5.5-8.49h-5.9v-12.5h22.33Z"/>`,
		"h":     `<path fill="{text}" d="M67.81,93.64v-24.06h-7.63v24.06h-16.04V34.36h16.04v21.23h7.63v-21.23h16.04v59.29h-16.04Z"/>`,
		"i":     `<path fill="{text}" d="M55.98,93.64V34.36h16.04v59.29h-16.04Z"/>`,
		"j":     `<path fill="{text}" d="M50.4,75.68c1.49,2.04,3.85,3.46,6.53,3.46,4.48,0,4.64-4.8,4.64-8.26V33.85h16.04v40.18c0,7.94-1.02,20.13-16.59,20.13-3.7,0-7.16-.71-10.61-1.89v-16.59Z"/>`,
		"k":     `<path fill="{text}" d="M84.17,34.36l-12.82,28.62,14.31,30.67h-17.22l-9.99-27.21h-.08v27.21h-16.04V34.36h16.04v25.24h.08l8.81-25.24h16.91Z"/>`,
		"l":     `<path fill="{text}" d="M49.96,93.64V34.36h16.04v45.29h12.03v14h-28.07Z"/>`,
		"m":     `<path fill="{text}" d="M78.59,93.64l-1.49-37.27h-.16c-.71,4.32-1.26,8.65-2.04,12.97l-4.56,24.3h-13.05l-5.58-37.27h-.16c-.08,4.4,0,8.73-.24,13.13l-1.34,24.14h-17.22l6.92-59.29h19.97l4.56,31.53h.16c.47-3.7,.86-7.31,1.57-11.01l3.85-20.52h19.74l5.74,59.29h-16.67Z"/>`,
		"n":     `<path fill="{text}" d="M43.08,34.36h15.25l12.82,35.7,.16-.16c-.63-6.84-1.81-13.6-1.81-20.44v-15.1h15.41v59.29h-15.41l-11.95-34.2-.16,.16c.39,5.58,1.1,11.09,1.1,16.67v17.38h-15.41V34.36Z"/>`,
		"o":     `<path fill="{text}" d="M64,33.33c16.28,0,22.65,16.2,22.65,30.67s-6.37,30.67-22.65,30.67-22.65-16.2-22.65-30.67,6.37-30.67,22.65-30.67Zm0,46.39c6.06,0,6.29-12.03,6.29-15.73s-.24-15.73-6.29-15.73-6.29,12.03-6.29,15.73,.24,15.73,6.29,15.73Z"/>`,
		"p":     `<path fill="{text}" d="M44.3,93.64V34.36h14.63c14,0,24.77,4.56,24.77,20.13s-9.44,20.13-23.35,19.74v19.42h-16.04Zm17.22-46.55l-1.18,.08v13.68c4.4,.24,7.16-2.2,7.16-6.6,0-3.77-2.04-7.16-5.98-7.16Z"/>`,
		"q":     `<path fill="{text}" d="M73.95,98.28l-5.5-7.94c-1.49,.47-3.07,.71-4.64,.71-16.28,0-22.65-16.2-22.65-30.67s6.37-30.67,22.65-30.67,22.65,16.2,22.65,30.67c-.16,8.41-1.89,16.67-6.61,22.65l7,8.18-12.89,7.08Zm-16.43-37.9c0,3.7,.24,15.73,6.29,15.73s6.29-12.03,6.29-15.73-.24-15.73-6.29-15.73-6.29,12.03-6.29,15.73Z"/>`,
		"r":     `<path fill="{text}" d="M67.73,93.64l-8.89-24.69-.16,.16c.08,2.91,.24,5.82,.24,8.73v15.8h-16.04V34.36h15.49c13.92,0,23.98,4.09,23.98,19.74,0,6.53-2.83,12.34-9.2,14.86l11.95,24.69h-17.38Zm-7.78-31.85c4.25,0,6.53-3.7,6.53-7.55,0-5.19-3.14-7.39-7.55-7.16v14.63l1.02,.08Z"/>`,
		"s":     `<path fill="{text}" d="M48.67,73.67c2.2,3.3,6.05,6.6,10.14,6.6,2.2,0,4.72-1.34,4.72-3.93,0-1.42-.47-2.44-1.26-3.3-.71-.86-1.65-1.57-2.67-2.28-3.14-2.44-5.9-4.72-7.94-7.55-1.97-2.75-3.22-6.05-3.22-10.38,0-6.92,4.4-19.5,18.09-19.5,3.85,0,8.02,1.18,11.4,2.99v18.01c-1.89-2.99-5.74-6.37-9.44-6.37-1.89,0-4.17,1.34-4.17,3.62,0,1.26,.71,2.36,1.57,3.3,.86,.94,2.04,1.73,2.91,2.44,3.38,2.36,6.05,4.56,7.86,7.23,1.89,2.67,2.91,5.82,2.91,10.3,0,11.01-7.71,19.82-18.87,19.82-4.09,0-8.26-.87-12.03-2.44v-18.56Z"/>`,
		"t":     `<path fill="{text}" d="M55.78,93.64V48.82h-8.81v-14.47h34.05v14.47h-9.2v44.82h-16.04Z"/>`,
		"u":     `<path fill="{text}" d="M84.13,33.85v38.76c0,7.16-1.57,21.55-20.13,21.55s-20.13-14.39-20.13-21.55V33.85h16.04v37.27c0,1.57-.08,3.62,.39,5.19,.47,1.65,1.49,2.91,3.7,2.91s3.3-1.26,3.69-2.83c.47-1.65,.39-3.7,.39-5.19V33.85h16.04Z"/>`,
		"v":     `<path fill="{text}" d="M54.56,93.64l-13.6-59.29h16.83l3.77,24.38c.71,4.8,1.1,9.51,1.73,14.31h.16c.71-4.8,.94-9.51,1.81-14.31l4.4-24.38h17.38l-16.12,59.29h-16.36Z"/>`,
		"w":     `<path fill="{text}" d="M70.05,93.64l-4.33-24.06c-.86-4.8-1.1-8.96-1.65-13.05h-.16c-.39,4.09-.55,8.26-1.26,13.05l-3.38,24.06h-15.57l-13.21-59.29h15.8l3.54,23.9c.63,4.4,1.1,8.73,1.65,13.13h.16c.47-4.4,.55-8.73,1.34-13.13l4.17-23.9h13.76l4.09,23.9c.71,4.32,1.18,8.57,1.81,12.9h.16c.47-3.7,.63-8.26,1.34-12.9l3.54-23.9h15.65l-11.95,59.29h-15.49Z"/>`,
		"x":     `<path fill="{text}" d="M69.23,93.64l-3.3-9.2c-.63-2.2-1.26-4.48-1.89-6.68h-.16c-.63,2.2-1.26,4.48-1.89,6.68l-3.54,9.2h-17.46l14.39-30.74-13.6-28.54h17.85l2.59,7.16c.79,2.2,1.18,4.56,1.81,6.84h.16c.55-2.28,.94-4.56,1.65-6.84l2.12-7.16h18.01l-13.52,28.54,14.55,30.74h-17.77Z"/>`,
		"y":     `<path fill="{text}" d="M56.37,93.64v-25.16l-14.63-34.13h16.67l5.5,18.32h.16l5.58-18.32h16.59l-13.84,34.13v25.16h-16.04Z"/>`,
		"z":     `<path fill="{text}" d="M42.61,93.64l16.2-37.66c1.1-2.59,2.91-5.43,4.56-8.1-3.14,.16-6.68,.47-10.14,.47h-7.08v-14h39.24l-17.53,40.57c-.79,1.89-1.73,3.62-2.52,5.43,2.91-.24,5.9-.71,8.96-.71h9.12v14H42.61Z"/>`,
		"0":     `<path fill="{text}" d="M84.99,64c0,13.13-4.48,30.67-21,30.67s-20.99-17.53-20.99-30.67,4.48-30.67,20.99-30.67,21,17.53,21,30.67Zm-26.18,0c0,6.68,.31,16.98,5.19,16.98s5.19-10.3,5.19-16.98-.31-16.98-5.19-16.98-5.19,10.3-5.19,16.98Z"/>`,
		"1":     `<path fill="{text}" d="M59.16,93.64V48.04h-6.37v-13.68h22.41v59.29h-16.04Z"/>`,
		"2":     `<path fill="{text}" d="M46.03,54.37c0-11.87,8.02-20.52,20.05-20.52,11.24,0,18.71,8.49,18.71,19.42s-7.78,20.52-15.49,27.6l.16,.16c2.67-.16,5.35-.55,8.02-.55h6.61v13.68H43.2l15.65-20.84c4.4-5.9,10.3-13.37,10.3-20.99,0-2.28-.94-5.27-3.7-5.27-3.46,0-3.85,4.8-3.85,7.23v2.2h-15.49l-.08-2.12Z"/>`,
		"3":     `<path fill="{text}" d="M59.88,77.21c.24,3.22,1.97,4.56,3.85,4.56,3.15,0,4.4-2.52,4.4-5.35,0-4.88-3.54-6.53-7.08-6.53h-2.2v-12.5h1.97c4.25,0,6.53-1.42,6.53-6.05,0-1.97-.71-5.11-3.3-5.11-3.3,0-3.22,4.01-3.22,4.56v.39l-15.49,.08c.71-10.85,8.41-17.93,19.34-17.93,9.91,0,18.32,7.16,18.32,17.38,0,5.11-2.59,9.83-7.16,12.27,5.43,2.83,8.02,7.94,8.02,14,0,11.4-9.83,17.69-20.36,17.69s-19.74-7.08-19.34-18.48h15.65l.08,1.02Z"/>`,
		"4":     `<path fill="{text}" d="M42.69,81.61v-12.5l17.85-34.75h19.66v34.52h5.11v12.74h-5.11v12.03h-15.96v-12.03h-21.55Zm22.1-22.88c0-2.52,.39-5.03,.63-7.55l-.16-.16-9.51,17.85h9.04v-10.14Z"/>`,
		"5":     `<path fill="{text}" d="M46.43,75.52c3.62,2.2,8.26,4.25,12.58,4.25,4.09,0,8.34-2.2,8.34-6.76,0-5.11-4.95-6.68-9.28-6.68-3.77,0-7.16,.71-10.61,2.12l4.56-34.6h29.57v13.68h-16.91l-1.26,6.92c.79-.08,1.49-.16,2.28-.16,9.99,0,17.77,8.26,17.77,18.16,0,12.66-9.91,21.7-22.41,21.7-5.58,0-11.72-1.73-16.51-4.64l1.89-14Z"/>`,
		"6":     `<path fill="{text}" d="M63.88,55.23l.16,.16c1.73-1.1,3.3-1.65,5.5-1.65,9.51,0,14.55,10.46,14.55,18.79,0,12.5-7.94,21.62-20.68,21.62s-19.5-9.75-19.5-21.55c0-9.28,5.9-21.07,10.22-29.33l4.95-9.44h17.22l-12.42,21.39Zm4.88,17.61c0-2.91-.79-7.55-4.56-7.55s-4.48,4.88-4.48,7.78,.39,8.34,4.48,8.34,4.56-5.58,4.56-8.57Z"/>`,
		"7":     `<path fill="{text}" d="M42.02,93.64l21.94-45.61h-20.13v-13.68h42.15l-26.73,59.29h-17.22Z"/>`,
		"8":     `<path fill="{text}" d="M43.79,77.05c0-6.92,3.54-12.03,9.99-14.23v-.16c-5.98-1.73-8.88-6.29-8.88-12.5,0-10.38,8.88-16.83,19.11-16.83s19.11,6.45,19.11,16.83c0,6.21-2.91,10.77-8.89,12.5v.16c6.45,2.2,9.99,7.31,9.99,14.23,0,10.77-9.83,17.61-20.21,17.61s-20.21-6.84-20.21-17.61Zm24.22-2.04c0-2.36-1.02-5.98-4.01-5.98s-4.01,3.62-4.01,5.98c0,3.38,.86,7,4.01,7s4.01-3.62,4.01-7Zm-7.47-24.06c0,2.2,.63,5.43,3.46,5.43s3.46-3.22,3.46-5.43c0-2.04-.71-4.95-3.46-4.95s-3.46,2.91-3.46,4.95Z"/>`,
		"9":     `<path fill="{text}" d="M64.2,72.77l-.16-.16c-1.73,1.02-3.3,1.65-5.5,1.65-9.59,0-14.63-10.46-14.63-18.79,0-12.5,7.94-21.62,20.76-21.62s19.42,9.67,19.42,21.54c0,9.28-5.9,21.07-10.14,29.25l-5.03,9.51h-17.14l12.42-21.39Zm-4.88-17.61c0,2.83,.79,7.55,4.56,7.55s4.33-4.88,4.33-7.78c0-3.07-.31-8.41-4.4-8.41s-4.48,5.66-4.48,8.65Z"/>`,
	}

	lettersMap = map[string]string{}
}

func transformLetter(str string) string {
	str = strings.TrimSpace(strings.ToLower(str))
	if len(str) > 0 {
		first, _ := utf8.DecodeRuneInString(str)
		if _, ok := letters[string(first)]; ok {
			return string(first)
		} else if v, ok := lettersMap[string(first)]; ok && v != "" {
			return v
		}
	}
	return Default
}

// AddLetterMap define new unicode letter for transform to english letter
func AddLetterMap(letter string, replacement string) {
	lettersMap[letter] = strings.ToLower(replacement)
}

// AddLetterShape define shape for letter
func AddLetterShape(letter string, shape string) {
	letters[strings.ToLower(letter)] = shape
}

// SetDefaultLetter set fallback default letter shape
func SetDefaultLetter(shape string) {
	letters[Default] = shape
}

// RegisterPersianTransform add transform map to transforming persian letter to english
func RegisterPersianTransform() {
	AddLetterMap("ا", "a")
	AddLetterMap("ب", "b")
	AddLetterMap("پ", "p")
	AddLetterMap("ت", "t")
	AddLetterMap("ث", "s")
	AddLetterMap("ج", "j")
	AddLetterMap("چ", "c")
	AddLetterMap("ح", "h")
	AddLetterMap("خ", "k")
	AddLetterMap("د", "d")
	AddLetterMap("ذ", "z")
	AddLetterMap("ر", "r")
	AddLetterMap("ز", "z")
	AddLetterMap("ژ", "z")
	AddLetterMap("س", "s")
	AddLetterMap("ش", "s")
	AddLetterMap("ص", "s")
	AddLetterMap("ض", "z")
	AddLetterMap("ط", "t")
	AddLetterMap("ظ", "z")
	AddLetterMap("ع", "a")
	AddLetterMap("غ", "g")
	AddLetterMap("ف", "f")
	AddLetterMap("ق", "g")
	AddLetterMap("ک", "k")
	AddLetterMap("گ", "g")
	AddLetterMap("ل", "l")
	AddLetterMap("م", "m")
	AddLetterMap("ن", "n")
	AddLetterMap("و", "v")
	AddLetterMap("ه", "h")
	AddLetterMap("ی", "i")
	AddLetterMap("۰", "0")
	AddLetterMap("۱", "1")
	AddLetterMap("۲", "2")
	AddLetterMap("۳", "3")
	AddLetterMap("۴", "4")
	AddLetterMap("۵", "5")
	AddLetterMap("۶", "6")
	AddLetterMap("۷", "7")
	AddLetterMap("۸", "8")
	AddLetterMap("۹", "9")
}

// RegisterPersiaShapes add persian letters shape
func RegisterPersiaShapes() {
	AddLetterMap("۰", "0")
	AddLetterShape("ا", `<path fill="{text}" d="M59.63,44.21h8.74v40.92h-8.74V44.21Z"/>`)
	AddLetterShape("ب", `<path fill="{text}" d="M47.94,73.93c-2.35-1.38-4.21-3.24-5.59-5.59-1.38-2.35-2.06-4.92-2.06-7.71v-7.29h8.74v6.8c0,1.82,.65,3.39,1.94,4.71,1.3,1.32,2.85,1.97,4.67,1.97h19.31c1.17,0,2.13-.39,2.88-1.18,.75-.79,1.12-1.77,1.12-2.94v-9.35h8.74v9.35c0,2.43-.56,4.65-1.67,6.68-1.11,2.02-2.64,3.63-4.58,4.83s-4.11,1.79-6.5,1.79h-19.31c-2.79,0-5.36-.69-7.71-2.06Zm21.49,16.32h-7.77v-7.77h7.77v7.77Z"/>`)
	AddLetterShape("پ", `<path fill="{text}" d="M47.94,73.93c-2.35-1.38-4.21-3.24-5.59-5.59-1.38-2.35-2.06-4.92-2.06-7.71v-7.29h8.74v6.8c0,1.82,.65,3.39,1.94,4.71,1.3,1.32,2.85,1.97,4.67,1.97h19.31c1.17,0,2.13-.39,2.88-1.18,.75-.79,1.12-1.77,1.12-2.94v-9.35h8.74v9.35c0,2.43-.56,4.65-1.67,6.68-1.11,2.02-2.64,3.63-4.58,4.83s-4.11,1.79-6.5,1.79h-19.31c-2.79,0-5.36-.69-7.71-2.06Zm8.2,10.13h7.77v7.77h-7.77v-7.77Zm5.52,10.87h7.83v7.77h-7.83v-7.77Zm5.34-10.87h7.83v7.77h-7.83v-7.77Z"/>`)
	AddLetterShape("ت", `<path fill="{text}" d="M47.94,77.85c-2.35-1.38-4.21-3.24-5.59-5.59-1.38-2.35-2.06-4.92-2.06-7.71v-7.29h8.74v6.8c0,1.82,.65,3.39,1.94,4.71,1.3,1.32,2.85,1.97,4.67,1.97h19.31c1.17,0,2.13-.39,2.88-1.18,.75-.79,1.12-1.77,1.12-2.94v-9.35h8.74v9.35c0,2.43-.56,4.65-1.67,6.68-1.11,2.02-2.64,3.63-4.58,4.83s-4.11,1.79-6.5,1.79h-19.31c-2.79,0-5.36-.69-7.71-2.06Zm6.8-28.42h7.77v7.83h-7.77v-7.83Zm10.93,0h7.77v7.83h-7.77v-7.83Z"/>`)
	AddLetterShape("ث", `<path fill="{text}" d="M47.94,83.28c-2.35-1.38-4.21-3.24-5.59-5.59-1.38-2.35-2.06-4.92-2.06-7.71v-7.29h8.74v6.8c0,1.82,.65,3.39,1.94,4.71,1.3,1.32,2.85,1.97,4.67,1.97h19.31c1.17,0,2.13-.39,2.88-1.18,.75-.79,1.12-1.77,1.12-2.94v-9.35h8.74v9.35c0,2.43-.56,4.65-1.67,6.68-1.11,2.02-2.64,3.63-4.58,4.83s-4.11,1.79-6.5,1.79h-19.31c-2.79,0-5.36-.69-7.71-2.06Zm14.57-28.42v7.83h-7.77v-7.83h7.77Zm5.53-10.87v7.83h-7.77v-7.83h7.77Zm5.46,10.87v7.83h-7.83v-7.83h7.83Z"/>`)
	AddLetterShape("ج", `<path fill="{text}" d="M52.95,87.89c-2.33-1.38-4.17-3.23-5.53-5.55-1.36-2.33-2.03-4.83-2.03-7.5,0-2.27,.5-4.43,1.49-6.5,.99-2.07,2.38-3.83,4.16-5.28l15.6-11.47-6.5-3.1c-.65-.32-1.23-.49-1.76-.49-1.01,0-1.98,.43-2.91,1.27l-4.55,4.92-6.07-5.28,4.61-5.53c1.13-1.37,2.48-2.39,4.04-3.04,1.56-.65,3.17-.97,4.83-.97,1.82,0,3.56,.45,5.22,1.34l13.11,6.13c.32,.08,.63,.12,.91,.12h5.59v8.56h-7.35l-18.76,13.78c-1.66,1.34-2.49,3.12-2.49,5.34,0,1.82,.66,3.38,1.97,4.68,1.31,1.29,2.86,1.94,4.64,1.94h17.61v8.68h-18.28c-2.71,0-5.23-.69-7.56-2.06Zm13.15-20.1h7.77v7.77h-7.77v-7.77Z"/>`)
	AddLetterShape("چ", `<path fill="{text}" d="M52.95,87.89c-2.33-1.38-4.17-3.23-5.53-5.55-1.36-2.33-2.03-4.83-2.03-7.5,0-2.27,.5-4.43,1.49-6.5,.99-2.07,2.38-3.83,4.16-5.28l15.6-11.47-6.5-3.1c-.65-.32-1.23-.49-1.76-.49-1.01,0-1.98,.43-2.91,1.27l-4.55,4.92-6.07-5.28,4.61-5.53c1.13-1.37,2.48-2.39,4.04-3.04,1.56-.65,3.17-.97,4.83-.97,1.82,0,3.56,.45,5.22,1.34l13.11,6.13c.32,.08,.63,.12,.91,.12h5.59v8.56h-7.35l-18.76,13.78c-1.66,1.34-2.49,3.12-2.49,5.34,0,1.82,.66,3.38,1.97,4.68,1.31,1.29,2.86,1.94,4.64,1.94h17.61v8.68h-18.28c-2.71,0-5.23-.69-7.56-2.06Zm11.99-22.46h6.44v6.5h-6.44v-6.5Zm3.76,7.77h6.44v5.89h-6.44v-5.89Zm3.95-7.77h6.5v6.5h-6.5v-6.5Z"/>`)
	AddLetterShape("ح", `<path fill="{text}" d="M52.95,87.89c-2.33-1.38-4.17-3.23-5.53-5.55-1.36-2.33-2.03-4.83-2.03-7.5,0-2.27,.5-4.43,1.49-6.5,.99-2.07,2.38-3.83,4.16-5.28l15.6-11.47-6.5-3.1c-.65-.32-1.23-.49-1.76-.49-1.01,0-1.98,.43-2.91,1.27l-4.55,4.92-6.07-5.28,4.61-5.53c1.13-1.37,2.48-2.39,4.04-3.04,1.56-.65,3.17-.97,4.83-.97,1.82,0,3.56,.45,5.22,1.34l13.11,6.13c.32,.08,.63,.12,.91,.12h5.59v8.56h-7.35l-18.76,13.78c-1.66,1.34-2.49,3.12-2.49,5.34,0,1.82,.66,3.38,1.97,4.68,1.31,1.29,2.86,1.94,4.64,1.94h17.61v8.68h-18.28c-2.71,0-5.23-.69-7.56-2.06Z"/>`)
	AddLetterShape("خ", `<path fill="{text}" d="M52.95,94.15c-2.33-1.38-4.17-3.23-5.53-5.55-1.36-2.33-2.03-4.83-2.03-7.5,0-2.27,.5-4.43,1.49-6.5,.99-2.07,2.38-3.83,4.16-5.28l15.6-11.47-6.5-3.1c-.65-.32-1.23-.49-1.76-.49-1.01,0-1.98,.43-2.91,1.27l-4.55,4.92-6.07-5.28,4.61-5.53c1.13-1.37,2.48-2.39,4.04-3.04,1.56-.65,3.17-.97,4.83-.97,1.82,0,3.56,.45,5.22,1.34l13.11,6.13c.32,.08,.63,.12,.91,.12h5.59v8.56h-7.35l-18.76,13.78c-1.66,1.34-2.49,3.12-2.49,5.34,0,1.82,.66,3.38,1.97,4.68,1.31,1.29,2.86,1.94,4.64,1.94h17.61v8.68h-18.28c-2.71,0-5.23-.69-7.56-2.06Zm1.43-61.02h7.77v7.77h-7.77v-7.77Z"/>`)
	AddLetterShape("د", `<path fill="{text}" d="M50.58,70.98h15.18c.89,0,1.61-.26,2.15-.79,.55-.53,.82-1.19,.82-2,0-.57-.17-1.06-.52-1.49-.34-.43-.96-.94-1.85-1.55l-12.75-8.74,4.92-7.22,14.21,9.84c1.54,1.17,2.7,2.46,3.49,3.86s1.18,3.09,1.18,5.07c0,2.27-.56,4.33-1.67,6.19-1.11,1.86-2.62,3.33-4.52,4.4-1.9,1.07-3.97,1.61-6.19,1.61h-14.45v-9.17Z"/>`)
	AddLetterShape("ذ", `<path fill="{text}" d="M50.58,77.24h15.18c.89,0,1.61-.26,2.15-.79,.55-.53,.82-1.19,.82-2,0-.57-.17-1.06-.52-1.49-.34-.43-.96-.94-1.85-1.55l-12.75-8.74,4.92-7.22,14.21,9.84c1.54,1.17,2.7,2.46,3.49,3.86s1.18,3.09,1.18,5.07c0,2.27-.56,4.33-1.67,6.19-1.11,1.86-2.62,3.33-4.52,4.4-1.9,1.07-3.97,1.61-6.19,1.61h-14.45v-9.17Zm4.31-34.3h7.77v7.83h-7.77v-7.83Z"/>`)
	AddLetterShape("ر", `<path fill="{text}" d="M61.18,73.26c1.21-1.01,2.02-2.1,2.43-3.28,.4-1.17,.61-2.87,.61-5.1v-20.16h8.74v20.16c0,3.4-.43,6.19-1.28,8.38-.85,2.19-2.31,4.14-4.37,5.86-2.07,1.72-5.06,3.55-8.99,5.49l-3.28-7.35c2.87-1.66,4.92-3,6.13-4.01Z"/>`)
	AddLetterShape("ز", `<path fill="{text}" d="M61.18,79.76c1.21-1.01,2.02-2.1,2.43-3.28,.4-1.17,.61-2.87,.61-5.1v-20.16h8.74v20.16c0,3.4-.43,6.19-1.28,8.38-.85,2.19-2.31,4.14-4.37,5.86-2.07,1.72-5.06,3.55-8.99,5.49l-3.28-7.35c2.87-1.66,4.92-3,6.13-4.01Zm3.52-41.53h7.77v7.77h-7.77v-7.77Z"/>`)
	AddLetterShape("ژ", `<path fill="{text}" d="M58.84,85.22c1.21-1.01,2.02-2.1,2.43-3.28,.4-1.17,.61-2.87,.61-5.1v-20.16h8.74v20.16c0,3.4-.43,6.19-1.28,8.38-.85,2.19-2.31,4.14-4.37,5.86-2.07,1.72-5.06,3.55-8.99,5.49l-3.28-7.35c2.87-1.66,4.92-3,6.13-4.01Zm-1.82-41.53h7.77v7.77h-7.77v-7.77Zm5.28-10.93h7.77v7.83h-7.77v-7.83Zm5.22,10.93h7.77v7.77h-7.77v-7.77Z"/>`)
	AddLetterShape("س", `<path fill="{text}" d="M40.14,79.61c-3-1.76-5.37-4.14-7.13-7.13-1.76-3-2.64-6.25-2.64-9.77v-11.35h8.74v10.87c0,1.94,.49,3.74,1.46,5.4,.97,1.66,2.29,2.97,3.95,3.95,1.66,.97,3.46,1.46,5.4,1.46s3.82-.49,5.49-1.46,3.01-2.29,3.98-3.95c.97-1.66,1.46-3.46,1.46-5.4v-13.3h8.44v13.78c0,3.52-.86,6.78-2.58,9.77-1.72,3-4.06,5.37-7.01,7.13-2.96,1.76-6.21,2.64-9.78,2.64s-6.78-.88-9.77-2.64Zm26.08-10.99c-1.15-2.04-1.73-4.85-1.73-8.41v-11.29h5.16v10.75c0,.81,.27,1.49,.82,2.03,.55,.55,1.2,.82,1.97,.82,.81,0,1.48-.27,2-.82,.53-.55,.79-1.22,.79-2.03v-10.75h6.98v11.29c0,3.32-.85,6.06-2.55,8.23-1.7,2.17-4.11,3.25-7.23,3.25s-5.07-1.02-6.22-3.07Zm13.33-.3c-1.7-2.17-2.55-4.91-2.55-8.23v-11.17h6.98v10.63c0,.81,.26,1.49,.79,2.03,.53,.55,1.17,.82,1.94,.82h2.19v-15.3h8.74v24.47h-10.93c-3.08,0-5.46-1.08-7.16-3.25Z"/>`)
	AddLetterShape("ش", `<path fill="{text}" d="M40.14,90.57c-3-1.76-5.37-4.14-7.13-7.13-1.76-3-2.64-6.25-2.64-9.77v-11.35h8.74v10.87c0,1.94,.49,3.74,1.46,5.4,.97,1.66,2.29,2.97,3.95,3.95,1.66,.97,3.46,1.46,5.4,1.46s3.82-.49,5.49-1.46,3.01-2.29,3.98-3.95c.97-1.66,1.46-3.46,1.46-5.4v-13.3h8.44v13.78c0,3.52-.86,6.78-2.58,9.77-1.72,3-4.06,5.37-7.01,7.13-2.96,1.76-6.21,2.64-9.78,2.64s-6.78-.88-9.77-2.64Zm26.08-10.99c-1.15-2.04-1.73-4.85-1.73-8.41v-11.29h5.16v10.75c0,.81,.27,1.49,.82,2.03,.55,.55,1.2,.82,1.97,.82,.81,0,1.48-.27,2-.82,.53-.55,.79-1.22,.79-2.03v-10.75h6.98v11.29c0,3.32-.85,6.06-2.55,8.23-1.7,2.17-4.11,3.25-7.23,3.25s-5.07-1.02-6.22-3.07Zm4.04-32.57h7.83v7.77h-7.83v-7.77Zm5.53-10.87h7.83v7.77h-7.83v-7.77Zm3.76,43.14c-1.7-2.17-2.55-4.91-2.55-8.23v-11.17h6.98v10.63c0,.81,.26,1.49,.79,2.03,.53,.55,1.17,.82,1.94,.82h2.19v-15.3h8.74v24.47h-10.93c-3.08,0-5.46-1.08-7.16-3.25Zm1.7-32.27h7.77v7.77h-7.77v-7.77Z"/>`)
	AddLetterShape("ص", `<path fill="{text}" d="M36.92,80.88c-3-1.76-5.37-4.14-7.13-7.13-1.76-3-2.64-6.25-2.64-9.77v-11.35h8.74v10.87c0,1.94,.49,3.74,1.46,5.4,.97,1.66,2.29,2.97,3.95,3.95,1.66,.97,3.46,1.46,5.4,1.46s3.82-.49,5.49-1.46,3.01-2.29,3.98-3.95c.97-1.66,1.46-3.46,1.46-5.4v-13.3h8.44v13.78c0,3.52-.85,6.78-2.55,9.77s-4.04,5.37-7.01,7.13c-2.97,1.76-6.24,2.64-9.81,2.64s-6.78-.88-9.77-2.64Zm28.81-10.14c-2.04-1.4-3.6-3.29-4.67-5.68-1.07-2.39-1.61-5.04-1.61-7.95v-6.92h6.98v8.14c0,1.46,.52,2.71,1.55,3.76,1.03,1.05,2.26,1.58,3.67,1.58h15.79c1.21,0,2.26-.44,3.13-1.34,.87-.89,1.31-1.94,1.31-3.16,0-1.29-.44-2.41-1.31-3.34-.87-.93-1.97-1.4-3.31-1.4-.57,0-1.19,.13-1.88,.39-.69,.26-1.6,.76-2.73,1.49l-16.58,9.96-3.1-8.01,16.15-9.59c1.66-1.01,3.07-1.74,4.22-2.19,1.15-.44,2.52-.67,4.1-.67,2.39,0,4.61,.61,6.68,1.82,2.06,1.21,3.7,2.86,4.92,4.95,1.21,2.08,1.82,4.34,1.82,6.77s-.61,4.62-1.82,6.71c-1.21,2.08-2.86,3.73-4.95,4.95-2.09,1.21-4.32,1.82-6.71,1.82h-14.39c-2.79,0-5.21-.7-7.26-2.09Z"/>`)
	AddLetterShape("ض", `<path fill="{text}" d="M36.92,87.14c-3-1.76-5.37-4.14-7.13-7.13-1.76-3-2.64-6.25-2.64-9.77v-11.35h8.74v10.87c0,1.94,.49,3.74,1.46,5.4,.97,1.66,2.29,2.97,3.95,3.95,1.66,.97,3.46,1.46,5.4,1.46s3.82-.49,5.49-1.46,3.01-2.29,3.98-3.95c.97-1.66,1.46-3.46,1.46-5.4v-13.3h8.44v13.78c0,3.52-.85,6.78-2.55,9.77s-4.04,5.37-7.01,7.13c-2.97,1.76-6.24,2.64-9.81,2.64s-6.78-.88-9.77-2.64Zm28.81-10.14c-2.04-1.4-3.6-3.29-4.67-5.68-1.07-2.39-1.61-5.04-1.61-7.95v-6.92h6.98v8.14c0,1.46,.52,2.71,1.55,3.76,1.03,1.05,2.26,1.58,3.67,1.58h15.79c1.21,0,2.26-.44,3.13-1.34,.87-.89,1.31-1.94,1.31-3.16,0-1.29-.44-2.4-1.31-3.31-.87-.91-1.91-1.37-3.13-1.37-1.38,0-3.04,.53-4.98,1.58l-16.58,9.9-2.91-7.71,17.43-10.44c1.01-.65,2.14-1.14,3.4-1.49,1.25-.34,2.47-.52,3.64-.52,2.39,0,4.61,.61,6.68,1.82,2.06,1.21,3.7,2.86,4.92,4.95,1.21,2.08,1.82,4.34,1.82,6.77s-.61,4.62-1.82,6.71c-1.21,2.08-2.86,3.73-4.95,4.95-2.09,1.21-4.32,1.82-6.71,1.82h-14.39c-2.79,0-5.21-.7-7.26-2.09Zm17.82-37.43h7.77v7.77h-7.77v-7.77Z"/>`)
	AddLetterShape("ط", `<path fill="{text}" d="M45.51,75.96h23.56c1.21,0,2.26-.44,3.13-1.34,.87-.89,1.31-1.94,1.31-3.16,0-1.29-.45-2.41-1.34-3.34-.89-.93-1.98-1.4-3.28-1.4-.57,0-1.19,.14-1.88,.42-.69,.28-1.6,.79-2.73,1.52l-13.48,7.89-4.13-7.29,14.02-8.32c1.7-1.01,3.13-1.74,4.28-2.19,1.15-.44,2.52-.67,4.1-.67,2.39,0,4.61,.61,6.68,1.82,2.07,1.21,3.7,2.86,4.92,4.95s1.82,4.34,1.82,6.77-.61,4.62-1.82,6.71-2.86,3.73-4.95,4.95c-2.09,1.21-4.32,1.82-6.71,1.82h-23.5v-9.17Zm6.98-11.17c.24-.85,.36-1.86,.36-3.04v-17.55h8.68v10.2c0,1.86-.16,3.48-.49,4.86-.33,1.38-.79,2.33-1.4,2.85l-8.62,4.55c.73-.4,1.21-1.03,1.46-1.88Z"/>`)
	AddLetterShape("ظ", `<path fill="{text}" d="M45.51,75.96h23.56c1.21,0,2.26-.44,3.13-1.34,.87-.89,1.31-1.94,1.31-3.16,0-1.29-.45-2.41-1.34-3.34-.89-.93-1.98-1.4-3.28-1.4-.57,0-1.19,.14-1.88,.42-.69,.28-1.6,.79-2.73,1.52l-14.45,8.56-3.16-7.95,14.02-8.32c1.7-1.01,3.13-1.74,4.28-2.19,1.15-.44,2.52-.67,4.1-.67,2.39,0,4.61,.61,6.68,1.82,2.07,1.21,3.7,2.86,4.92,4.95s1.82,4.34,1.82,6.77-.61,4.62-1.82,6.71-2.86,3.73-4.95,4.95c-2.09,1.21-4.32,1.82-6.71,1.82h-23.5v-9.17Zm6.98-11.17c.24-.85,.36-1.86,.36-3.04v-17.55h8.68v10.2c0,1.9-.11,3.44-.33,4.61s-.62,2.02-1.18,2.55l-8.99,5.1c.73-.4,1.21-1.03,1.46-1.88Zm13.42-19.19h7.83v7.77h-7.83v-7.77Z"/>`)
	AddLetterShape("ع", `<path fill="{text}" d="M58.47,89.26c-2.71-1.58-4.87-3.72-6.47-6.44s-2.4-5.67-2.4-8.86c0-2.91,.68-5.52,2.03-7.8,1.36-2.29,3.21-4.12,5.56-5.49,2.35-1.38,5-2.25,7.95-2.61,.28-.04,.63-.38,1.03-1.03,.49-.65,.87-.97,1.15-.97h8.56v9.17h-8.56c-1.54,0-2.97,.38-4.28,1.15-1.32,.77-2.36,1.82-3.13,3.16-.77,1.34-1.15,2.77-1.15,4.31s.38,3.03,1.15,4.34c.77,1.31,1.81,2.36,3.13,3.13,1.31,.77,2.74,1.15,4.28,1.15h11.05v9.17h-11.05c-3.2,0-6.15-.79-8.86-2.37Zm1.97-27.44c-2.17-1.01-3.88-2.43-5.13-4.25-1.26-1.82-1.88-3.95-1.88-6.37,0-2.55,.63-4.85,1.88-6.89,1.25-2.04,2.97-3.65,5.16-4.83,2.19-1.17,4.57-1.76,7.16-1.76h6.68v8.56h-6.68c-1.5,0-2.71,.45-3.64,1.37-.93,.91-1.4,2.08-1.4,3.49,0,1.5,.47,2.69,1.4,3.58,.93,.89,2.14,1.34,3.64,1.34v7.29c-2.63,0-5.03-.5-7.19-1.52Z"/>`)
	AddLetterShape("غ", `<path fill="{text}" d="M58.47,95.6c-2.71-1.58-4.87-3.72-6.47-6.44s-2.4-5.67-2.4-8.86c0-2.91,.68-5.52,2.03-7.8,1.36-2.29,3.21-4.12,5.56-5.49,2.35-1.38,5-2.25,7.95-2.61,.28-.04,.63-.38,1.03-1.03,.49-.65,.87-.97,1.15-.97h8.56v9.17h-8.56c-1.54,0-2.97,.38-4.28,1.15-1.32,.77-2.36,1.82-3.13,3.16-.77,1.34-1.15,2.77-1.15,4.31s.38,3.03,1.15,4.34c.77,1.31,1.81,2.36,3.13,3.13,1.31,.77,2.74,1.15,4.28,1.15h11.05v9.17h-11.05c-3.2,0-6.15-.79-8.86-2.37Zm1.97-27.44c-2.17-1.01-3.88-2.43-5.13-4.25-1.26-1.82-1.88-3.95-1.88-6.37,0-2.55,.63-4.85,1.88-6.89,1.25-2.04,2.97-3.65,5.16-4.83,2.19-1.17,4.57-1.76,7.16-1.76h6.68v8.56h-6.68c-1.5,0-2.71,.45-3.64,1.37-.93,.91-1.4,2.08-1.4,3.49,0,1.5,.47,2.69,1.4,3.58,.93,.89,2.14,1.34,3.64,1.34v7.29c-2.63,0-5.03-.5-7.19-1.52Zm2.64-36.79h7.77v7.77h-7.77v-7.77Z"/>`)
	AddLetterShape("ف", `<path fill="{text}" d="M47.94,86.42c-2.35-1.38-4.21-3.24-5.59-5.59-1.38-2.35-2.06-4.92-2.06-7.71v-7.29h8.74v6.8c0,1.82,.65,3.39,1.94,4.71,1.3,1.32,2.85,1.97,4.67,1.97h20.58c.81,0,1.47-.26,1.97-.79,.51-.53,.76-1.19,.76-2v-.61h-2.85c-2.23,0-4.28-.52-6.16-1.55s-3.38-2.44-4.49-4.22c-1.11-1.78-1.67-3.74-1.67-5.89,0-2.27,.57-4.35,1.7-6.25,1.13-1.9,2.64-3.4,4.52-4.49,1.88-1.09,3.92-1.64,6.1-1.64s4.16,.55,5.92,1.64c1.76,1.09,3.15,2.59,4.16,4.49,1.01,1.9,1.52,3.99,1.52,6.25v12.75c0,2.06-.51,3.98-1.52,5.74-1.01,1.76-2.39,3.16-4.13,4.19-1.74,1.03-3.66,1.55-5.77,1.55h-20.64c-2.79,0-5.36-.69-7.71-2.06Zm31.57-18.58v-3.89c0-1.05-.36-1.94-1.09-2.67-.73-.73-1.6-1.09-2.61-1.09-1.09,0-2,.36-2.73,1.09s-1.09,1.62-1.09,2.67,.37,2.01,1.09,2.76c.73,.75,1.64,1.12,2.73,1.12h3.7Zm-7.29-28.54h7.77v7.77h-7.77v-7.77Z"/>`)
	AddLetterShape("ق", `<path fill="{text}" d="M59.42,55.05c1.19-2.04,2.81-3.66,4.86-4.86,2.04-1.19,4.26-1.79,6.65-1.79s4.54,.6,6.47,1.79c1.92,1.2,3.43,2.8,4.52,4.83,1.09,2.02,1.64,4.25,1.64,6.68v12.69c0,3.52-.87,6.78-2.61,9.77-1.74,3-4.11,5.37-7.1,7.13-3,1.76-6.27,2.64-9.84,2.64s-6.78-.88-9.77-2.64c-3-1.76-5.37-4.14-7.13-7.13-1.76-3-2.64-6.25-2.64-9.77v-11.35h8.74v10.87c0,1.94,.49,3.74,1.46,5.4,.97,1.66,2.29,2.98,3.95,3.98,1.66,.99,3.46,1.49,5.4,1.49s3.82-.5,5.49-1.49c1.68-.99,3.01-2.32,3.98-3.98,.97-1.66,1.46-3.46,1.46-5.4v.43h-4.01c-2.43,0-4.65-.55-6.68-1.64-2.02-1.09-3.63-2.6-4.83-4.52-1.19-1.92-1.79-4.08-1.79-6.47s.6-4.6,1.79-6.65Zm2.16-19.64h7.77v7.77h-7.77v-7.77Zm6.41,28.63c.79,.75,1.77,1.12,2.94,1.12h4.01v-3.95c0-1.17-.39-2.15-1.15-2.94-.77-.79-1.72-1.18-2.85-1.18s-2.15,.39-2.94,1.18-1.18,1.77-1.18,2.94,.39,2.08,1.18,2.82Zm4.52-28.63h7.77v7.77h-7.77v-7.77Z"/>`)
	AddLetterShape("ک", `<path fill="{text}" d="M47.94,85.19c-2.35-1.38-4.21-3.24-5.59-5.59-1.38-2.35-2.06-4.92-2.06-7.71v-7.22h8.74v6.74c0,1.82,.65,3.39,1.94,4.71,1.3,1.32,2.85,1.97,4.67,1.97h19.31c1.09,0,2.03-.36,2.82-1.09,.79-.73,1.18-1.7,1.18-2.91,0-1.46-.71-2.61-2.13-3.46l-13.66-9.05v-6.62l12.08-12.87,5.95,5.65-8.99,9.41,10.08,6.62c1.7,1.09,3.02,2.56,3.98,4.4,.95,1.84,1.43,3.88,1.43,6.1,0,2.35-.58,4.51-1.73,6.5-1.15,1.98-2.7,3.56-4.64,4.74-1.94,1.17-4.09,1.76-6.44,1.76h-19.25c-2.79,0-5.36-.69-7.71-2.06Z"/>`)
	AddLetterShape("گ", `<path fill="{text}" d="M47.94,87.04c-2.35-1.38-4.21-3.24-5.59-5.59-1.38-2.35-2.06-4.92-2.06-7.71v-7.22h8.74v6.74c0,1.82,.65,3.39,1.94,4.71,1.3,1.32,2.85,1.97,4.67,1.97h19.31c1.09,0,2.03-.36,2.82-1.09,.79-.73,1.18-1.7,1.18-2.91,0-1.46-.71-2.61-2.13-3.46l-13.66-9.05v-6.62l12.08-12.87,5.95,5.65-8.99,9.41,10.08,6.62c1.7,1.09,3.02,2.56,3.98,4.4,.95,1.84,1.43,3.88,1.43,6.1,0,2.35-.58,4.51-1.73,6.5-1.15,1.98-2.7,3.56-4.64,4.74-1.94,1.17-4.09,1.76-6.44,1.76h-19.25c-2.79,0-5.36-.69-7.71-2.06Zm8.68-35.52l10.69-11.29,4.19,4.25-10.69,11.29-4.19-4.25Z"/>`)
	AddLetterShape("ل", `<path fill="{text}" d="M54.16,87.83c-3-1.76-5.37-4.14-7.13-7.13-1.76-3-2.64-6.25-2.64-9.77v-11.35h8.74v10.87c0,1.94,.49,3.74,1.46,5.4,.97,1.66,2.29,2.97,3.95,3.95,1.66,.97,3.46,1.46,5.4,1.46s3.82-.49,5.49-1.46,3.01-2.29,3.98-3.95c.97-1.66,1.46-3.46,1.46-5.4v-31.57h8.74v32.06c0,3.52-.88,6.78-2.64,9.77-1.76,3-4.15,5.37-7.16,7.13-3.02,1.76-6.31,2.64-9.87,2.64s-6.78-.88-9.77-2.64Z"/>`)
	AddLetterShape("م", `<path fill="{text}" d="M45.88,50.4c1.11-1.9,2.62-3.42,4.52-4.55,1.9-1.13,3.97-1.7,6.19-1.7h13.9c2.43,0,4.65,.55,6.68,1.64,2.02,1.09,3.63,2.6,4.83,4.52,1.19,1.92,1.79,4.08,1.79,6.47s-.6,4.6-1.79,6.65c-1.19,2.04-2.81,3.66-4.86,4.86-2.04,1.19-4.26,1.79-6.65,1.79s-4.54-.6-6.47-1.79c-1.92-1.19-3.43-2.8-4.52-4.83-1.09-2.02-1.64-4.25-1.64-6.68v-3.95h-1.21c-1.01,0-1.88,.37-2.61,1.12-.73,.75-1.09,1.63-1.09,2.64v28.6h-8.74v-28.6c0-2.23,.56-4.29,1.67-6.19Zm27.56,9.32c.79-.79,1.18-1.77,1.18-2.94s-.39-2.08-1.18-2.85c-.79-.77-1.77-1.15-2.94-1.15h-4.01v4.01c0,1.17,.38,2.15,1.15,2.94,.77,.79,1.72,1.18,2.85,1.18s2.15-.39,2.94-1.18Z"/>`)
	AddLetterShape("ن", `<path fill="{text}" d="M54.16,81.52c-3-1.76-5.37-4.14-7.13-7.13-1.76-3-2.64-6.25-2.64-9.77v-11.35h8.74v10.87c0,1.94,.49,3.74,1.46,5.4,.97,1.66,2.29,2.97,3.95,3.95,1.66,.97,3.46,1.46,5.4,1.46s3.82-.49,5.49-1.46,3.01-2.29,3.98-3.95c.97-1.66,1.46-3.46,1.46-5.4v-13.3h8.74v13.78c0,3.52-.88,6.78-2.64,9.77-1.76,3-4.15,5.37-7.16,7.13-3.02,1.76-6.31,2.64-9.87,2.64s-6.78-.88-9.77-2.64Zm5.71-36.34h7.83v7.77h-7.83v-7.77Z"/>`)
	AddLetterShape("و", `<path fill="{text}" d="M63.7,74.78c1.76-.97,2.96-1.88,3.61-2.73,.65-.85,.97-1.86,.97-3.04h-3.95c-2.43,0-4.66-.55-6.68-1.64-2.02-1.09-3.63-2.6-4.83-4.52-1.19-1.92-1.79-4.08-1.79-6.47s.6-4.65,1.79-6.68c1.19-2.02,2.8-3.63,4.83-4.83,2.02-1.19,4.25-1.79,6.68-1.79s4.49,.6,6.44,1.79c1.94,1.2,3.46,2.8,4.55,4.83s1.64,4.25,1.64,6.68v8.86c0,3.76-.58,6.81-1.73,9.14-1.15,2.33-3.02,4.35-5.59,6.07-2.57,1.72-6.32,3.65-11.26,5.8l-2.73-7.71c3.6-1.54,6.28-2.79,8.05-3.76Zm4.58-14.39v-4.07c0-1.17-.38-2.14-1.12-2.91-.75-.77-1.69-1.15-2.82-1.15s-2.16,.39-2.95,1.15c-.79,.77-1.18,1.74-1.18,2.91s.39,2.15,1.18,2.91c.79,.77,1.77,1.15,2.95,1.15h3.95Z"/>`)
	AddLetterShape("ه", `<path fill="{text}" d="M51.86,73.66c-.61-1.5-.91-3.06-.91-4.67,0-1.74,.35-3.45,1.06-5.13,.71-1.68,1.73-3.19,3.07-4.52l2.79-2.79-2.73-2.67,6.25-6.07,11.6,11.54c1.34,1.34,2.35,2.84,3.04,4.52,.69,1.68,1.03,3.39,1.03,5.13,0,1.62-.3,3.18-.91,4.67-.61,1.5-1.52,2.85-2.73,4.07-1.3,1.26-2.75,2.21-4.37,2.85-1.62,.65-3.3,.97-5.04,.97s-3.42-.32-5.04-.97c-1.62-.65-3.08-1.6-4.37-2.85-1.21-1.21-2.13-2.57-2.73-4.07Zm12.2-1.21c1.13,0,2.08-.43,2.85-1.27,.77-.73,1.15-1.62,1.15-2.67,0-1.13-.43-2.12-1.27-2.97l-2.79-2.85-2.79,2.73c-.85,.85-1.27,1.84-1.27,2.97s.4,2.02,1.21,2.79c.81,.85,1.78,1.27,2.91,1.27Z"/>`)
	AddLetterShape("ی", `<path fill="{text}" d="M50.89,82.7c-2.35-1.38-4.21-3.24-5.59-5.59-1.38-2.35-2.06-4.92-2.06-7.71v-14.75h8.74v14.27c0,1.82,.65,3.39,1.94,4.71,1.3,1.31,2.85,1.97,4.67,1.97h13.9c.97,0,1.76-.31,2.37-.94,.61-.63,.91-1.43,.91-2.4,0-.73-.22-1.4-.67-2-.45-.61-1.05-.99-1.82-1.15l-15.66-3.58v-7.95c0-2.31,.6-4.46,1.79-6.47,1.19-2,2.81-3.59,4.86-4.77,2.04-1.17,4.24-1.76,6.59-1.76h11.29v8.68h-11.41c-1.17,0-2.2,.44-3.07,1.31s-1.31,1.89-1.31,3.07v.97c2.27,.49,3.88,.84,4.83,1.06,.95,.22,1.77,.42,2.46,.58l2.43,.55c2.55,.77,4.63,2.25,6.25,4.43,1.62,2.19,2.43,4.64,2.43,7.35,0,2.23-.55,4.27-1.64,6.13-1.09,1.86-2.58,3.34-4.46,4.43s-3.94,1.64-6.16,1.64h-13.9c-2.79,0-5.36-.69-7.71-2.07Z"/>`)
	AddLetterShape("۱", `<path fill="{text}" d="M60.48,63.33c0-3.72-.11-7.08-.33-10.08-.22-2.99-.68-5.85-1.37-8.56l7.95-1.34c1.66,3.85,2.49,12.2,2.49,25.08v17.55h-8.74v-22.65Z"/>`)
	AddLetterShape("۲", `<path fill="{text}" d="M68.74,43.72v10.81c0,1.17-.38,2.16-1.15,2.94s-1.72,1.18-2.85,1.18-2.16-.39-2.95-1.18c-.56-.56-.91-1.22-1.07-1.98-.39-5.73-1.1-9.78-2.11-12.14l-7.95,1.34c.69,2.71,1.14,5.57,1.37,8.56,.22,3,.33,6.36,.33,10.08v22.65h8.74v-17.55c0-.38,0-.74,0-1.11,1.11,.37,2.32,.56,3.65,.56,2.39,0,4.54-.6,6.47-1.79,1.92-1.19,3.43-2.8,4.52-4.83,1.09-2.02,1.64-4.25,1.64-6.68v-10.87h-8.62Z"/>`)
	AddLetterShape("۳", `<path fill="{text}" d="M74.5,43.72v12.57c0,.81-.26,1.49-.79,2.03-.53,.55-1.19,.82-2,.82s-1.43-.27-1.97-.82c-.55-.55-.82-1.22-.82-2.03v-10.81h-8.68v10.81c0,.81-.27,1.49-.82,2.03-.55,.55-1.21,.82-1.97,.82-.81,0-1.48-.27-2-.82-.53-.55-.79-1.22-.79-2.03h-.02c-.38-6.15-1.1-10.47-2.16-12.93l-7.95,1.34c.69,2.71,1.14,5.57,1.37,8.56,.22,3,.33,6.36,.33,10.08v22.65h8.74v-17.55c0-.13,0-.26,0-.4,.78,.18,1.6,.27,2.49,.27,3.04,0,5.42-1.06,7.13-3.18,1.72,2.12,4.09,3.18,7.13,3.18,2.14,0,4.12-.52,5.92-1.55,1.8-1.03,3.23-2.43,4.28-4.19,1.05-1.76,1.58-3.67,1.58-5.74v-13.11h-8.99Z"/>`)
	AddLetterShape("۴", `<path fill="{text}" d="M65.46,76.57c-1.17,.08-2.18-.28-3.01-1.09-.83-.81-1.24-1.84-1.24-3.1,0-.93,.31-1.74,.94-2.43,.63-.69,1.45-1.13,2.46-1.34l10.62-1.88-1.52-8.5-7.47,1.25c-.22,.02-.46,.03-.72,.03-.97,0-1.78-.35-2.43-1.06-.65-.71-.97-1.59-.97-2.64,0-.97,.33-1.8,1-2.49,.67-.69,1.51-1.03,2.52-1.03h7.53v-8.74h-7.53c-2.23,0-4.28,.56-6.16,1.67-1.88,1.11-3.37,2.6-4.46,4.46-1.09,1.86-1.64,3.91-1.64,6.13s.55,4.34,1.64,6.22c.3,.51,.63,.98,.98,1.44-1.02,.95-1.87,2.04-2.53,3.27-.95,1.78-1.43,3.68-1.43,5.71,0,.77,.06,1.52,.18,2.25,.36,2.15,1.19,4.08,2.49,5.8,1.3,1.72,2.89,3.05,4.8,3.98,1.9,.93,3.9,1.38,6.01,1.34h10.44v-9.23h-10.5Z"/>`)
	AddLetterShape("۵", `<path fill="{text}" d="M81.7,65.91c-1.03-2.35-2.49-4.63-4.37-6.86-1.88-2.23-4.67-5.16-8.35-8.8-1.13-1.13-2.69-2.71-4.67-4.74-1.98-2.02-3.46-3.56-4.43-4.61l-6.37,5.77c.12,.12,1.59,1.58,4.4,4.37,.05,.05,.09,.09,.14,.14-3.21,3.24-5.7,5.89-7.43,7.94-1.88,2.23-3.33,4.5-4.34,6.83-1.01,2.33-1.52,4.93-1.52,7.8,0,4.53,1.21,8.12,3.64,10.75,2.43,2.63,5.89,3.95,10.38,3.95h10.44c4.49,0,7.95-1.31,10.38-3.95,2.43-2.63,3.64-6.21,3.64-10.75,0-2.87-.52-5.48-1.55-7.83Zm-9.05,13.05c-.91,.81-1.87,1.21-2.88,1.21h-11.54c-1.01,0-1.97-.4-2.88-1.21-.91-.81-1.37-2.16-1.37-4.07s.27-3.59,.82-5.19c.55-1.6,1.62-3.47,3.22-5.62,1.38-1.86,3.37-4.11,5.95-6.75,2.59,2.66,4.59,4.91,5.98,6.75,1.62,2.15,2.7,4.02,3.25,5.62,.55,1.6,.82,3.33,.82,5.19s-.45,3.26-1.37,4.07Z"/>`)
	AddLetterShape("۶", `<path fill="{text}" d="M77.05,76.54c-.22-3.02-.33-6.41-.33-10.17v-22.77h-13.42c-2.03,0-4.3-.09-6.83-.27-2.53-.18-4.54-.37-6.04-.58l-.85,8.44c4.69,.73,9.13,1.09,13.3,1.09h5.1v9.17c0,12.91,.83,21.29,2.49,25.14l7.95-1.4c-.69-2.75-1.14-5.64-1.37-8.65Z"/>`)
	AddLetterShape("۷", `<path fill="{text}" d="M73.08,43.06l-6.98,19.06c-.77,2.09-1.47,4.05-2.1,5.91-.63-1.85-1.32-3.82-2.09-5.91l-6.98-19.06-8.2,2.97,6.86,18.82,.73,1.88c1.66,4.57,2.95,8.42,3.89,11.54,.93,3.12,1.4,5.79,1.4,8.01h8.8c0-2.23,.47-4.9,1.4-8.01,.93-3.12,2.23-6.96,3.89-11.54l.73-1.88,6.86-18.82-8.2-2.97Z"/>`)
	AddLetterShape("۸", `<path fill="{text}" d="M74.41,64.52l-.73-2c-1.58-4.21-2.85-7.95-3.83-11.23-.97-3.28-1.46-6.03-1.46-8.26h-8.8c0,2.27-.5,5.06-1.49,8.38-.99,3.32-2.26,7.02-3.79,11.11l-.73,2-6.86,18.82,8.2,2.97,6.98-19.13c.77-2.09,1.46-4.05,2.09-5.91,.63,1.85,1.33,3.82,2.1,5.91l6.98,19.13,8.2-2.97-6.86-18.82Z"/>`)
	AddLetterShape("۹", `<path fill="{text}" d="M67.95,78.3c-.39-3.26-.58-6.18-.58-8.77v-.61h-3.89c-2.43,0-4.66-.55-6.68-1.64-2.02-1.09-3.63-2.6-4.83-4.52-1.19-1.92-1.79-4.08-1.79-6.47s.6-4.65,1.79-6.68c1.19-2.02,2.8-3.63,4.83-4.83,2.02-1.19,4.25-1.79,6.68-1.79s4.49,.6,6.44,1.79c1.94,1.19,3.46,2.8,4.55,4.83,1.09,2.02,1.64,4.25,1.64,6.68v9.84c0,3.76,.11,7.16,.33,10.2,.22,3.04,.68,5.93,1.37,8.68l-7.95,1.34c-.89-2.1-1.53-4.79-1.91-8.04Zm-.58-18v-4.01c0-1.17-.37-2.16-1.09-2.94-.73-.79-1.66-1.18-2.79-1.18s-2.15,.39-2.94,1.18-1.18,1.77-1.18,2.94,.39,2.08,1.18,2.85c.79,.77,1.77,1.15,2.94,1.15h3.89Z"/>`)
}