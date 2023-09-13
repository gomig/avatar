package avatar

var stickers map[string]string

func init() {
	ResetSticker()
}

// ResetSticker clear registered stickers
func ResetSticker() {
	stickers = map[string]string{Default: defaultSticker}
}

// AddSticker register new sticker
func AddSticker(name string, shape string) {
	stickers[name] = shape
}
