package avatar

var hairColors map[string]hairColorType
var skinColors map[string]skinColorType
var palettes map[string]paletteType

func init() {
	ResetHairColor()
	ResetSkinColor()
	ResetPalette()
}

// ResetHairColor clear all hair colors except default
func ResetHairColor() {
	hairColors = map[string]hairColorType{
		Default:        defaultHairColor,
		BrownHair:      {Color{Hex, "#3a292f"}, Color{Hex, "#191215"}, Color{Hex, "#664954"}},
		LightBrownHair: {Color{Hex, "#744819"}, Color{Hex, "#563517"}, Color{Hex, "#825329"}},
		DarkBrownHair:  {Color{Hex, "#432818"}, Color{Hex, "#140b06"}, Color{Hex, "#884139"}},
	}
}

// AddHairColor add new hair color
func AddHairColor(name string, base Color, shadow Color, highlight Color) {
	base.validate("SetHairColor", "base")
	shadow.validate("SetHairColor", "shadow")
	highlight.validate("SetHairColor", "highlight")
	hairColors[name] = hairColorType{base, shadow, highlight}
}

// SetDefaultHairColor set default hair color
func SetDefaultHairColor(base Color, shadow Color, highlight Color) {
	AddHairColor(Default, base, shadow, highlight)
}

// ResetSkinColor clear all skin colors except default
func ResetSkinColor() {
	skinColors = map[string]skinColorType{
		Default:   defaultSkinColor,
		WhiteSkin: {Color{Hex, "#fbdad0"}, Color{Hex, "#84706d"}},
		BrownSkin: {Color{Hex, "#ceaa82"}, Color{Hex, "#896f52"}},
		BlackSkin: {Color{Hex, "#b06f51"}, Color{Hex, "#3d261d"}},
	}
}

// AddSkinColor add new skin color
func AddSkinColor(name string, base Color, highlight Color) {
	base.validate("SetSkinColor", "base")
	highlight.validate("SetSkinColor", "highlight")
	skinColors[name] = skinColorType{base, highlight}
}

// SetDefaultSkinColor set default skin color
func SetDefaultSkinColor(base Color, highlight Color) {
	AddSkinColor(Default, base, highlight)
}

// ResetPalette clear all palette except default
func ResetPalette() {
	palettes = map[string]paletteType{Default: defaultPalette}
}

// AddPalette add new palette
func AddPalette(name string, shape Color, text Color, dress Color, dressShadow Color, decorator Color) {
	shape.validate("SetPaletteColor", "shape")
	text.validate("SetPaletteColor", "text")
	dress.validate("SetPaletteColor", "dress")
	dressShadow.validate("SetPaletteColor", "dressShadow")
	decorator.validate("SetPaletteColor", "decorator")
	palettes[name] = paletteType{
		shape:       shape,
		text:        text,
		dress:       dress,
		dressShadow: dressShadow,
		decorator:   decorator,
	}
}

// SetDefaultPalette set default palette value
func SetDefaultPalette(shape Color, text Color, dress Color, dressShadow Color, decorator Color) {
	AddPalette(Default, shape, text, dress, dressShadow, decorator)
}
