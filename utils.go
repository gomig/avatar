package avatar

import "github.com/gomig/utils"

func filterKeys(legalItems []string, items []string) []string {
	res := make([]string, 0)

	for _, i := range items {
		if utils.Contains(legalItems, i) {
			res = append(res, i)
		}
	}

	return res
}

// shapeKeys get registered shape names
func shapeKeys() []string {
	keys := make([]string, 0)
	for k := range shapes {
		keys = append(keys, k)
	}
	return keys
}

// resolveShape get registered shape data
//
// @return shape, mask
func resolveShape(name string) (string, string) {
	if v, ok := shapes[name]; ok {
		return v.shape, v.mask
	}
	return "", ""
}

// resolveLetterShape get letter shape
// return fallback shape if letter not found
//
// @return shape
func resolveLetterShape(text string) string {
	letter := transformLetter(text)
	if v, ok := letters[letter]; ok {
		return v
	}
	if v, ok := letters[Default]; ok {
		return v
	}
	return ""
}

// stickerKeys get registered sticker names
func stickerKeys() []string {
	keys := make([]string, 0)
	for k := range stickers {
		keys = append(keys, k)
	}
	return keys
}

// resolveSticker get registered sticker
// return default sticker if not found
//
// @return sticker
func resolveSticker(name string) string {
	if v, ok := stickers[name]; ok {
		return v
	}
	if v, ok := stickers[Default]; ok {
		return v
	}
	return ""
}

// hairColorKeys get registered hair color names
func hairColorKeys() []string {
	keys := make([]string, 0)
	for k := range hairColors {
		keys = append(keys, k)
	}
	return keys
}

// resolveHairColor get registered hair color
// return default hair color if not found
//
// @return base, highlight
func resolveHairColor(name string) hairColorType {
	if v, ok := hairColors[name]; ok {
		return v
	}
	if v, ok := hairColors[Default]; ok {
		return v
	}
	return defaultHairColor
}

// skinColorKeys get registered skin color names
func skinColorKeys() []string {
	keys := make([]string, 0)
	for k := range skinColors {
		keys = append(keys, k)
	}
	return keys
}

// resolveSkinColor get registered skin color
// return default skin color if not found
//
// @return skinColor
func resolveSkinColor(name string) skinColorType {
	if v, ok := skinColors[name]; ok {
		return v
	}
	if v, ok := skinColors[Default]; ok {
		return v
	}
	return defaultSkinColor
}

// paletteKeys get registered palette names
func paletteKeys() []string {
	keys := make([]string, 0)
	for k := range palettes {
		keys = append(keys, k)
	}
	return keys
}

// resolvePalette get registered palette
// return default palette if not found
//
// @return palette
func resolvePalette(name string) paletteType {
	if v, ok := palettes[name]; ok {
		return v
	}
	if v, ok := palettes[Default]; ok {
		return v
	}
	return defaultPalette
}

// facialHairKeys get registered facial hair names
func facialHairKeys() []string {
	keys := make([]string, 0)
	for k := range facialHairs {
		keys = append(keys, k)
	}
	return keys
}

// resolveFacialHair get registered facial hair
//
// @return facial hair
func resolveFacialHair(name string) string {
	if v, ok := facialHairs[name]; ok {
		return v
	}
	return ""
}

// maleHairKeys get registered hair name for mans
func maleHairKeys() []string {
	keys := make([]string, 0)
	for k := range maleHairs {
		keys = append(keys, k)
	}
	return keys
}

// resolveMaleHair get registered hair for man
//
// @return hair
func resolveMaleHair(name string) string {
	if v, ok := maleHairs[name]; ok {
		return v
	}
	return ""
}

// femaleHairKeys get registered hair name for womans
func femaleHairKeys() []string {
	keys := make([]string, 0)
	for k := range femaleHairs {
		keys = append(keys, k)
	}
	return keys
}

// resolveFemaleHair get registered hair for woman
//
// @return hair
func resolveFemaleHair(name string) string {
	if v, ok := femaleHairs[name]; ok {
		return v
	}
	return ""
}

// maleDressKeys get registered dress name for mans
func maleDressKeys() []string {
	keys := make([]string, 0)
	for k := range maleDresses {
		keys = append(keys, k)
	}
	return keys
}

// resolveMaleDress get registered dress for man
// return default dress if not found
//
// @return dress
func resolveMaleDress(name string) string {
	if v, ok := maleDresses[name]; ok {
		return v
	}
	if v, ok := maleDresses[Default]; ok {
		return v
	}
	return ""
}

// femaleDressKeys get registered dress name for womans
func femaleDressKeys() []string {
	keys := make([]string, 0)
	for k := range femaleDresses {
		keys = append(keys, k)
	}
	return keys
}

// resolveFemaleDress get registered dress for woman
// return default dress if not found
//
// @return dress
func resolveFemaleDress(name string) string {
	if v, ok := femaleDresses[name]; ok {
		return v
	}
	if v, ok := femaleDresses[Default]; ok {
		return v
	}
	return ""
}

// maleEyeKeys get registered eye name for mans
func maleEyeKeys() []string {
	keys := make([]string, 0)
	for k := range maleEyes {
		keys = append(keys, k)
	}
	return keys
}

// resolveMaleEye get registered eye for man
// return default eye if not found
//
// @return eye
func resolveMaleEye(name string) string {
	if v, ok := maleEyes[name]; ok {
		return v
	}
	if v, ok := maleEyes[Default]; ok {
		return v
	}
	return ""
}

// femaleEyeKeys get registered eye name for womans
func femaleEyeKeys() []string {
	keys := make([]string, 0)
	for k := range femaleEyes {
		keys = append(keys, k)
	}
	return keys
}

// resolveFemaleEye get registered eye for woman
// return default eye if not found
//
// @return eye
func resolveFemaleEye(name string) string {
	if v, ok := femaleEyes[name]; ok {
		return v
	}
	if v, ok := femaleEyes[Default]; ok {
		return v
	}
	return ""
}

// maleEyebrowKeys get registered eyebrow name for mans
func maleEyebrowKeys() []string {
	keys := make([]string, 0)
	for k := range maleEyebrows {
		keys = append(keys, k)
	}
	return keys
}

// resolveMaleEyebrow get registered eyebrow for man
// return default eyebrow if not found
//
// @return eyebrow
func resolveMaleEyebrow(name string) string {
	if v, ok := maleEyebrows[name]; ok {
		return v
	}
	if v, ok := maleEyebrows[Default]; ok {
		return v
	}
	return ""
}

// femaleEyebrowKeys get registered eyebrow name for womans
func femaleEyebrowKeys() []string {
	keys := make([]string, 0)
	for k := range femaleEyebrows {
		keys = append(keys, k)
	}
	return keys
}

// resolveFemaleEyebrow get registered eyebrow for woman
// return default eyebrow if not found
//
// @return eyebrow
func resolveFemaleEyebrow(name string) string {
	if v, ok := femaleEyebrows[name]; ok {
		return v
	}
	if v, ok := femaleEyebrows[Default]; ok {
		return v
	}
	return ""
}

// maleMouthKeys get registered mouth name for mans
func maleMouthKeys() []string {
	keys := make([]string, 0)
	for k := range maleMouths {
		keys = append(keys, k)
	}
	return keys
}

// resolveMaleMouth get registered mouth for man
// return default mouth if not found
//
// @return mouth
func resolveMaleMouth(name string) string {
	if v, ok := maleMouths[name]; ok {
		return v
	}
	if v, ok := maleMouths[Default]; ok {
		return v
	}
	return ""
}

// femaleMouthKeys get registered mouth name for womans
func femaleMouthKeys() []string {
	keys := make([]string, 0)
	for k := range femaleMouths {
		keys = append(keys, k)
	}
	return keys
}

// resolveFemaleMouth get registered mouth for woman
// return default mouth if not found
//
// @return mouth
func resolveFemaleMouth(name string) string {
	if v, ok := femaleMouths[name]; ok {
		return v
	}
	if v, ok := femaleMouths[Default]; ok {
		return v
	}
	return ""
}

// maleGlassKeys get registered glass name for mans
func maleGlassKeys() []string {
	keys := make([]string, 0)
	for k := range maleGlasses {
		keys = append(keys, k)
	}
	return keys
}

// resolveMaleGlass get registered glass for man
// return default glass if not found
//
// @return glass
func resolveMaleGlass(name string) string {
	if v, ok := maleGlasses[name]; ok {
		return v
	}
	return ""
}

// femaleGlassKeys get registered glass name for womans
func femaleGlassKeys() []string {
	keys := make([]string, 0)
	for k := range femaleGlasses {
		keys = append(keys, k)
	}
	return keys
}

// resolveFemaleGlass get registered glass for woman
// return default glass if not found
//
// @return glass
func resolveFemaleGlass(name string) string {
	if v, ok := femaleGlasses[name]; ok {
		return v
	}
	return ""
}

// maleAccessoryKeys get registered accessory name for mans
func maleAccessoryKeys() []string {
	keys := make([]string, 0)
	for k := range maleAccessories {
		keys = append(keys, k)
	}
	return keys
}

// resolveMaleAccessory get registered accessory for man
// return default accessory if not found
//
// @return accessory
func resolveMaleAccessory(name string) string {
	if v, ok := maleAccessories[name]; ok {
		return v
	}
	return ""
}

// femaleAccessoryKeys get registered accessory name for womans
func femaleAccessoryKeys() []string {
	keys := make([]string, 0)
	for k := range femaleAccessories {
		keys = append(keys, k)
	}
	return keys
}

// resolveFemaleAccessory get registered accessory for woman
// return default accessory if not found
//
// @return accessory
func resolveFemaleAccessory(name string) string {
	if v, ok := femaleAccessories[name]; ok {
		return v
	}
	return ""
}
