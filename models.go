package avatar

import (
	"fmt"
	"strings"
)

// Constants
type Gender int

const (
	Male   Gender = 1
	Female Gender = 2
	Both   Gender = 3
)

type ColorType int

const (
	Hex        ColorType = 0
	Definition ColorType = 1
)

type AvatarType int

const (
	MALE    AvatarType = 1
	FEMALE  AvatarType = 2
	LETTER  AvatarType = 3
	STICKER AvatarType = 4
)

// Models
type Color struct {
	Mode  ColorType
	Color string
}

func (c Color) validate(method string, field string) {
	if c.Mode == Definition && !strings.Contains(c.Color, `id="{id}"`) {
		panic(fmt.Sprintf(`[%s] definition color "%s" must contains id="{id}" placeholder attribute.`, method, field))
	}
}

type paletteType struct {
	shape       Color
	text        Color
	dress       Color
	dressShadow Color
	decorator   Color
}

type shapeType struct {
	shape string
	mask  string
}

type hairColorType struct {
	base      Color
	shadow    Color
	highlight Color
}

type skinColorType struct {
	skin   Color
	shadow Color
}
