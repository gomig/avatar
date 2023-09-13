package avatar

import (
	"encoding/base64"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/google/uuid"
)

// model
type Avatar struct {
	_id     string
	_type   AvatarType
	shape   string
	palette string
	// text
	name string
	// sticker
	sticker string
	// character
	skinColor  string
	hairColor  string
	facialHair string
	hair       string
	dress      string
	eye        string
	eyebrow    string
	mouth      string
	glass      string
	accessory  string
}

func (avatar *Avatar) init() {
	avatar._id = uuid.NewString()
	avatar.RandomizeShape("square")
	avatar.RandomizePalette(Default)

	if avatar._type == MALE || avatar._type == FEMALE {
		avatar.RandomizeSkinColor()
		avatar.RandomizeHair()
		avatar.RandomizeFacialHair()
		avatar.RandomizeDress()
		avatar.RandomizeEye()
		avatar.RandomizeEyebrow()
		avatar.RandomizeMouth()
		avatar.RandomizeGlass()
		avatar.RandomizeAccessory()
	}
}

func (Avatar) randTo(length int) int {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	return rnd.Intn(length)
}

func (avatar Avatar) id(key string) string {
	return key + "-" + avatar._id
}

// InlineSVG generate svg content for embeded (html svg tag) element
func (avatar Avatar) InlineSVG() string {
	data := make([]string, 0)
	data = append(data, `<svg viewBox="0 0 128 128" xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink">`)
	data = append(data, `<desc>Created with https://github.com/gomig/avatar</desc>`)

	// generate color definitions
	skinColors := resolveSkinColor(avatar.skinColor)
	hairColors := resolveHairColor(avatar.hairColor)
	palette := resolvePalette(avatar.palette)

	data = append(data, "<defs>")
	skinColor := skinColors.skin.Color
	if skinColors.skin.Mode == Definition {
		_id := avatar.id("skin-color")
		data = append(data, strings.ReplaceAll(skinColors.skin.Color, "{id}", _id))
		skinColor = fmt.Sprintf("url(#%s)", _id)
	}

	skinShadow := skinColors.shadow.Color
	if skinColors.shadow.Mode == Definition {
		_id := avatar.id("skin-shadow")
		data = append(data, strings.ReplaceAll(skinColors.shadow.Color, "{id}", _id))
		skinShadow = fmt.Sprintf("url(#%s)", _id)
	}

	hairColor := hairColors.base.Color
	if hairColors.base.Mode == Definition {
		_id := avatar.id("hair-color")
		data = append(data, strings.ReplaceAll(hairColors.base.Color, "{id}", _id))
		hairColor = fmt.Sprintf("url(#%s)", _id)
	}

	hairShadow := hairColors.shadow.Color
	if hairColors.shadow.Mode == Definition {
		_id := avatar.id("hair-shadow")
		data = append(data, strings.ReplaceAll(hairColors.shadow.Color, "{id}", _id))
		hairShadow = fmt.Sprintf("url(#%s)", _id)
	}

	hairHighlight := hairColors.highlight.Color
	if hairColors.highlight.Mode == Definition {
		_id := avatar.id("hair-highlight")
		data = append(data, strings.ReplaceAll(hairColors.highlight.Color, "{id}", _id))
		hairHighlight = fmt.Sprintf("url(#%s)", _id)
	}

	shapeColor := palette.shape.Color
	if palette.shape.Mode == Definition {
		_id := avatar.id("shape-color")
		data = append(data, strings.ReplaceAll(palette.shape.Color, "{id}", _id))
		shapeColor = fmt.Sprintf("url(#%s)", _id)
	}

	textColor := palette.text.Color
	if palette.text.Mode == Definition {
		_id := avatar.id("text-color")
		data = append(data, strings.ReplaceAll(palette.text.Color, "{id}", _id))
		textColor = fmt.Sprintf("url(#%s)", _id)
	}

	dressColor := palette.dress.Color
	if palette.dress.Mode == Definition {
		_id := avatar.id("dress-color")
		data = append(data, strings.ReplaceAll(palette.dress.Color, "{id}", _id))
		dressColor = fmt.Sprintf("url(#%s)", _id)
	}

	dressShadow := palette.dressShadow.Color
	if palette.dressShadow.Mode == Definition {
		_id := avatar.id("dress-shadow")
		data = append(data, strings.ReplaceAll(palette.dressShadow.Color, "{id}", _id))
		dressShadow = fmt.Sprintf("url(#%s)", _id)
	}

	decoratorColor := palette.decorator.Color
	if palette.decorator.Mode == Definition {
		_id := avatar.id("decorator-color")
		data = append(data, strings.ReplaceAll(palette.decorator.Color, "{id}", _id))
		decoratorColor = fmt.Sprintf("url(#%s)", _id)
	}
	data = append(data, "</defs>")

	// utils
	render := func(svg string) string {
		svg = strings.ReplaceAll(svg, "{skin}", skinColor)
		svg = strings.ReplaceAll(svg, "{skin_shadow}", skinShadow)
		svg = strings.ReplaceAll(svg, "{hair}", hairColor)
		svg = strings.ReplaceAll(svg, "{hair_shadow}", hairShadow)
		svg = strings.ReplaceAll(svg, "{hair_highlight}", hairHighlight)
		svg = strings.ReplaceAll(svg, "{shape}", shapeColor)
		svg = strings.ReplaceAll(svg, "{text}", textColor)
		svg = strings.ReplaceAll(svg, "{dress}", dressColor)
		svg = strings.ReplaceAll(svg, "{dress_shadow}", dressShadow)
		svg = strings.ReplaceAll(svg, "{decorator}", decoratorColor)
		return svg
	}

	// add shape
	shape, mask := resolveShape(avatar.shape)
	if shape != "" {
		data = append(data, render(shape))
	}

	if mask != "" {
		data = append(data, fmt.Sprintf(`<mask id="%s">`, avatar.id("mask")))
		data = append(data, mask)
		data = append(data, `</mask>`)
		data = append(data, fmt.Sprintf(`<g mask="url(#%s)">`, avatar.id("mask")))
	} else {
		data = append(data, "<g>")
	}

	// text avatar
	if avatar._type == LETTER {
		shape := resolveLetterShape(avatar.name)
		data = append(data, render(shape))
	}

	// sticker avatar
	if avatar._type == STICKER {
		shape := resolveSticker(avatar.sticker)
		data = append(data, render(shape))
	}

	// character avatar
	if avatar._type == MALE {
		data = append(data, render(bodyShape))
		data = append(data, render(resolveMaleDress(avatar.dress)))
		data = append(data, render(resolveMaleEye(avatar.eye)))
		data = append(data, render(resolveMaleEyebrow(avatar.eyebrow)))
		data = append(data, render(resolveMaleMouth(avatar.mouth)))
		data = append(data, render(resolveMaleGlass(avatar.glass)))
		data = append(data, render(resolveMaleAccessory(avatar.accessory)))
		data = append(data, render(resolveMaleHair(avatar.hair)))
		data = append(data, render(resolveFacialHair(avatar.facialHair)))
	}

	if avatar._type == FEMALE {
		data = append(data, render(bodyShape))
		data = append(data, render(resolveFemaleDress(avatar.dress)))
		data = append(data, render(resolveFemaleEye(avatar.eye)))
		data = append(data, render(resolveFemaleEyebrow(avatar.eyebrow)))
		data = append(data, render(resolveFemaleMouth(avatar.mouth)))
		data = append(data, render(resolveFemaleGlass(avatar.glass)))
		data = append(data, render(resolveFemaleAccessory(avatar.accessory)))
		data = append(data, render(resolveFemaleHair(avatar.hair)))
	}

	data = append(data, `</g>`)
	data = append(data, `</svg>`)
	return strings.Join(data, "")
}

// SVG get svg file content
func (avatar Avatar) SVG() string {
	return `<?xml version="1.0" encoding="utf-8"?><!DOCTYPE svg PUBLIC "-//W3C//DTD SVG 1.1//EN" "http://www.w3.org/Graphics/SVG/1.1/DTD/svg11.dtd">` + avatar.InlineSVG()
}

// Base64 get svg as base64 data image
func (avatar Avatar) Base64() string {
	return "data:image/svg+xml;base64," + string(base64.StdEncoding.EncodeToString([]byte(avatar.SVG())))
}

// SaveAs save svg file
func (avatar Avatar) SaveAs(path string) error {
	return os.WriteFile(path, []byte(avatar.SVG()), 0644)
}

// #region randomize
// RandomizeShape randomize shape from registered ones
func (avatar *Avatar) RandomizeShape(only ...string) *Avatar {
	if len(only) == 0 {
		only = shapeKeys()
	}
	only = filterKeys(shapeKeys(), only)
	if len(only) == 1 {
		avatar.shape = only[0]
	} else if len(only) > 1 {
		avatar.shape = only[avatar.randTo(len(only))]
	} else {
		avatar.shape = ""
	}
	return avatar
}

// RandomizePalette randomize palette from registered ones
func (avatar *Avatar) RandomizePalette(only ...string) *Avatar {
	if len(only) == 0 {
		only = paletteKeys()
	}
	only = filterKeys(paletteKeys(), only)
	if len(only) == 1 {
		avatar.palette = only[0]
	} else if len(only) > 1 {
		avatar.palette = only[avatar.randTo(len(only))]
	} else {
		avatar.palette = ""
	}
	return avatar
}

// RandomizeSticker randomize sticker from registered ones
func (avatar *Avatar) RandomizeSticker(only ...string) *Avatar {
	if len(only) == 0 {
		only = stickerKeys()
	}
	only = filterKeys(stickerKeys(), only)
	if len(only) == 1 {
		avatar.sticker = only[0]
	} else if len(only) > 1 {
		avatar.sticker = only[avatar.randTo(len(only))]
	} else {
		avatar.sticker = ""
	}
	return avatar
}

// RandomizeSkinColor randomize skin from registered ones
func (avatar *Avatar) RandomizeSkinColor(only ...string) *Avatar {
	if len(only) == 0 {
		only = skinColorKeys()
	}
	only = filterKeys(skinColorKeys(), only)
	if len(only) == 1 {
		avatar.skinColor = only[0]
	} else if len(only) > 1 {
		avatar.skinColor = only[avatar.randTo(len(only))]
	} else {
		avatar.skinColor = ""
	}
	return avatar
}

// RandomizeHairColor randomize hair color from registered ones
func (avatar *Avatar) RandomizeHairColor(only ...string) *Avatar {
	if len(only) == 0 {
		only = hairColorKeys()
	}
	only = filterKeys(hairColorKeys(), only)
	if len(only) == 1 {
		avatar.hairColor = only[0]
	} else if len(only) > 1 {
		avatar.hairColor = only[avatar.randTo(len(only))]
	} else {
		avatar.hairColor = ""
	}
	return avatar
}

// RandomizeFacialHair randomize facial hair from registered ones
func (avatar *Avatar) RandomizeFacialHair(only ...string) *Avatar {
	if len(only) == 0 {
		only = facialHairKeys()
	}
	only = filterKeys(facialHairKeys(), only)
	if len(only) == 1 {
		avatar.facialHair = only[0]
	} else if len(only) > 1 {
		avatar.facialHair = only[avatar.randTo(len(only))]
	} else {
		avatar.facialHair = ""
	}
	return avatar
}

// RandomizeHair randomize hair from registered ones
func (avatar *Avatar) RandomizeHair(only ...string) *Avatar {
	if avatar._type == FEMALE {
		if len(only) == 0 {
			only = femaleHairKeys()
		}
		only = filterKeys(femaleHairKeys(), only)
	} else {
		if len(only) == 0 {
			only = maleHairKeys()
		}
		only = filterKeys(maleHairKeys(), only)
	}
	if len(only) == 1 {
		avatar.hair = only[0]
	} else if len(only) > 1 {
		avatar.hair = only[avatar.randTo(len(only))]
	} else {
		avatar.hair = ""
	}
	return avatar
}

// RandomizeDress randomize dress from registered ones
func (avatar *Avatar) RandomizeDress(only ...string) *Avatar {
	if avatar._type == FEMALE {
		if len(only) == 0 {
			only = femaleDressKeys()
		}
		only = filterKeys(femaleDressKeys(), only)
	} else {
		if len(only) == 0 {
			only = maleDressKeys()
		}
		only = filterKeys(maleDressKeys(), only)
	}
	if len(only) == 1 {
		avatar.dress = only[0]
	} else if len(only) > 1 {
		avatar.dress = only[avatar.randTo(len(only))]
	} else {
		avatar.dress = ""
	}
	return avatar
}

// RandomizeEye randomize eye from registered ones
func (avatar *Avatar) RandomizeEye(only ...string) *Avatar {
	if avatar._type == FEMALE {
		if len(only) == 0 {
			only = femaleEyeKeys()
		}
		only = filterKeys(femaleEyeKeys(), only)
	} else {
		if len(only) == 0 {
			only = maleEyeKeys()
		}
		only = filterKeys(maleEyeKeys(), only)
	}
	if len(only) == 1 {
		avatar.eye = only[0]
	} else if len(only) > 1 {
		avatar.eye = only[avatar.randTo(len(only))]
	} else {
		avatar.eye = ""
	}
	return avatar
}

// RandomizeEyebrow randomize eyebrow from registered ones
func (avatar *Avatar) RandomizeEyebrow(only ...string) *Avatar {
	if avatar._type == FEMALE {
		if len(only) == 0 {
			only = femaleEyebrowKeys()
		}
		only = filterKeys(femaleEyebrowKeys(), only)
	} else {
		if len(only) == 0 {
			only = maleEyebrowKeys()
		}
		only = filterKeys(maleEyebrowKeys(), only)
	}
	if len(only) == 1 {
		avatar.eyebrow = only[0]
	} else if len(only) > 1 {
		avatar.eyebrow = only[avatar.randTo(len(only))]
	} else {
		avatar.eyebrow = ""
	}
	return avatar
}

// RandomizeMouth randomize mouth from registered ones
func (avatar *Avatar) RandomizeMouth(only ...string) *Avatar {
	if avatar._type == FEMALE {
		if len(only) == 0 {
			only = femaleMouthKeys()
		}
		only = filterKeys(femaleMouthKeys(), only)
	} else {
		if len(only) == 0 {
			only = maleMouthKeys()
		}
		only = filterKeys(maleMouthKeys(), only)
	}
	if len(only) == 1 {
		avatar.mouth = only[0]
	} else if len(only) > 1 {
		avatar.mouth = only[avatar.randTo(len(only))]
	} else {
		avatar.mouth = ""
	}
	return avatar
}

// RandomizeGlass randomize glass from registered ones
func (avatar *Avatar) RandomizeGlass(only ...string) *Avatar {
	if avatar._type == FEMALE {
		if len(only) == 0 {
			only = femaleGlassKeys()
		}
		only = filterKeys(femaleGlassKeys(), only)
	} else {
		if len(only) == 0 {
			only = maleGlassKeys()
		}
		only = filterKeys(maleGlassKeys(), only)
	}
	if len(only) == 1 {
		avatar.glass = only[0]
	} else if len(only) > 1 {
		avatar.glass = only[avatar.randTo(len(only))]
	} else {
		avatar.glass = ""
	}
	return avatar
}

// RandomizeAccessory randomize accessory from registered ones
func (avatar *Avatar) RandomizeAccessory(only ...string) *Avatar {
	if avatar._type == FEMALE {
		if len(only) == 0 {
			only = femaleAccessoryKeys()
		}
		only = filterKeys(femaleAccessoryKeys(), only)
	} else {
		if len(only) == 0 {
			only = maleAccessoryKeys()
		}
		only = filterKeys(maleAccessoryKeys(), only)
	}
	if len(only) == 1 {
		avatar.accessory = only[0]
	} else if len(only) > 1 {
		avatar.accessory = only[avatar.randTo(len(only))]
	} else {
		avatar.accessory = ""
	}
	return avatar
}

// #endregion
