# Avatar

Advanced customizable svg avatar generator for golang.

| Unicode                                                                           | Letter                                                                      | Sticker                                                                       | Female                                                                           | Female                                                                           | Female                                                                           | Male                                                                         | Male                                                                         | Male                                                                         |
| --------------------------------------------------------------------------------- | --------------------------------------------------------------------------- | ----------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| ![Unicode](https://github.com/gomig/avatar/blob/master/demo/persian.svg?raw=true) | ![John](https://github.com/gomig/avatar/blob/master/demo/john.svg?raw=true) | ![Apple](https://github.com/gomig/avatar/blob/master/demo/apple.svg?raw=true) | ![Female](https://github.com/gomig/avatar/blob/master/demo/female1.svg?raw=true) | ![Female](https://github.com/gomig/avatar/blob/master/demo/female3.svg?raw=true) | ![Female](https://github.com/gomig/avatar/blob/master/demo/female5.svg?raw=true) | ![Male](https://github.com/gomig/avatar/blob/master/demo/male1.svg?raw=true) | ![Male](https://github.com/gomig/avatar/blob/master/demo/male3.svg?raw=true) | ![Male](https://github.com/gomig/avatar/blob/master/demo/male5.svg?raw=true) |

## Installation

```bash
go get github.com/gomig/avatar
```

## Generate Avatar

### Avatar For Text

If letter shape or letter map not founded for first character of string fallback (`Default`) letter shape will be used.

```go
import "github.com/gomig/avatar"

func main(){
    john := avatar.NewTextAvatar("John") // make avatar with J letter
    htmlTag := john.InlineSVG()
    svg := john.SVG()
    dataImg := john.Base64()
    john.SaveAs("john.svg")
}
```

### Avatar For Sticker

```go
import "github.com/gomig/avatar"

func main(){
    // Register sticker shape globally
    avatar.AddSticker("apple", `<path fill="{text}" d="..."/>`)
    // Create avatar using apple sticker
    apple := avatar.NewStickerAvatar("apple")
    apple.SaveAs("apple.svg")
}
```

### Avatar For Person

```go
import "github.com/gomig/avatar"

func main(){
    mia := avatar.NewPersonAvatar(false) // pass true for male, false for female
    mia.RandomizeShape(avatar.Circle, "polygon") // get random background shape
    mia.SaveAs("mia.svg")
}
```

## Avatar Instance Methods

### InlineSVG

Generate svg content for embeded (html svg tag) element.

### SVG

Get svg file content.

### Base64

Get svg as base64 data image.

### SaveAs

Save svg file.

### Randomize Avatar Parameters

You can pass list off registered parts to generate random avatar. If no parameters passed random character created from all registered values.

**Note:** in PersonAvatar random values get from registered value based on gender.

#### Available Constant

If you want use predefined elements you can access them by constant.

- **Default (All except FacialHair, Glass And Accessory)** default element
- **None (FacialHair, Glass And Accessory)** allow select no item.
- **BrownHair (HairColor)** `#3a292f` hair color.
- **LightBrownHair (HairColor)** `#744819` hair color.
- **DarkBrownHair (HairColor)** `#432818` hair color.
- **WhiteSkin (Skin)** `#fbdad0` skin color.
- **BrownSkin (Skin)** `#ceaa82` skin color.
- **BlackSkin (Skin)** `#b06f51` skin color.
- **Fill (Shape)** fill background of avatar
- **Circle (Shape)** circular background
- **Polygon (Shape)** polygon background
- **Purple (Palette)** purple palette (only available if `RegisterExtraPalettes` called)
- **Green (Palette)** green palette (only available if `RegisterExtraPalettes` called)
- **Blue (Palette)** blue palette (only available if `RegisterExtraPalettes` called)
- **Yellow (Palette)** yellow palette (only available if `RegisterExtraPalettes` called)
- **Orange (Palette)** orange palette (only available if `RegisterExtraPalettes` called)
- **Red (Palette)** red palette (only available if `RegisterExtraPalettes` called)
- **Teal (Palette)** teal palette (only available if `RegisterExtraPalettes` called)
- **Pink (Palette)** pink palette (only available if `RegisterExtraPalettes` called)
- **Short (Hair)** short hair style (only available if `RegisterExtraMaleHair|RegisterExtraFemaleHair` called)
- **Wavy (Hair)** wavy hair style (only available if `RegisterExtraMaleHair|RegisterExtraFemaleHair` called)
- **Medium (MaleHair)** medium hair style (only available if `RegisterExtraMaleHair` called)
- **Curly (FemaleHair)** curly hair style (only available if `RegisterExtraFemaleHair` called)
- **Suit (Dress)** short hair style (only available if `RegisterExtraMaleDress|RegisterExtraFemaleDress` called)
- **Shirt (Dress)** short hair style (only available if `RegisterExtraMaleDress|RegisterExtraFemaleDress` called)
- **TShirt (Dress)** short hair style (only available if `RegisterExtraMaleDress|RegisterExtraFemaleDress` called)
- **Mustach (FacialHair)** mustach facial (only available if `RegisterExtraFacialHair` called)
- **MustachFancy (FacialHair)** fancy mustach (only available if `RegisterExtraFacialHair` called)
- **Beard (FacialHair)** short beard (only available if `RegisterExtraFacialHair` called)
- **BeardMedium (FacialHair)** medium beard (only available if `RegisterExtraFacialHair` called)
- **BeardLong (FacialHair)** long beard (only available if `RegisterExtraFacialHair` called)
- **Prescription (Glass)** colored prescription glass (only available if `RegisterExtraMaleGlass|RegisterExtraFemaleGlass` called)
- **PrescriptionRound (Glass)** rounded prescription glass (only available if `RegisterExtraMaleGlass|RegisterExtraFemaleGlass` called)
- **Sunglass (Glass)** black sunglass (only available if `RegisterExtraMaleGlass|RegisterExtraFemaleGlass` called)
- **SunglassRound (Glass)** rounded black sunglass (only available if `RegisterExtraMaleGlass|RegisterExtraFemaleGlass` called)
- **Necklace (FemaleAccessory)** normal necklace (only available if `RegisterExtraFemaleAccessory` called)
- **Choker (FemaleAccessory)** choker necklace (only available if `RegisterExtraFemaleAccessory` called)

#### RandomizeShape

Randomize shape from registered ones.

#### RandomizePalette

Randomize avatar color palette from registered ones.

#### RandomizeSticker

Randomize sticker from registered ones (this method works on sticker avatar only).

#### RandomizeSkinColor

Randomize skin from registered ones (this method works on person avatar only).

#### RandomizeHairColor

Randomize hair color from registered ones (this method works on person avatar only).

#### RandomizeFacialHair

Randomize facial hair from registered ones (this method works on male person avatar only).

#### RandomizeHair

Randomize hair from registered ones (this method works on person avatar only).

#### RandomizeDress

Randomize dress from registered ones (this method works on person avatar only).

#### RandomizeEye

Randomize eye from registered ones (this method works on person avatar only).

#### RandomizeEyebrow

Randomize eyebrow from registered ones (this method works on person avatar only).

#### RandomizeMouth

Randomize mouth from registered ones (this method works on person avatar only).

#### RandomizeGlass

Randomize glass from registered ones (this method works on person avatar only).

#### RandomizeAccessory

Randomize accessory from registered ones (this method works on person avatar only).

## Use Extra Elements

By default no extra element are exists and all generated avatar are equal! for using predefined package element you must call register functions.

### RegisterExtraPalettes

Register extra color palettes (`purple`, `green`, `blue`, `yellow`, `orange`, `red`, `teal` and `pink`).

### RegisterPersianTransform

Register transform map for transforming persian letter to english.

### RegisterPersiaShapes

Register letter shape for persian letters (if `RegisterPersianTransform` called this method not work and persian letter transformed and shown as english letter)

### RegisterExtraMaleHair

Register extra male hairs (`short`, `medium` and `wavy`).

### RegisterExtraMaleDress

Register extra male dresses (`t-shirt`, `suit` and `shirt`).

### RegisterExtraFacialHair

Register extra male facial hairs (`mustach`, `mustach-fancy`, `beard`, `beard-medium` and `beard-long`).

### RegisterExtraMaleGlass

Register extra male glasses (`prescription`, `prescription-round`, `sunglass` and `sunglass-round`).

### RegisterExtraMaleAccessory

Currently no male accessory are designed built-in!

### RegisterExtraFemaleHair

Register extra female hairs (`short`, `wavy` and `curly`).

### RegisterExtraFemaleDress

Register extra female dresses (`t-shirt`, `suit` and `shirt`).

### RegisterExtraFemaleGlass

Register extra female glasses (`prescription`, `prescription-round`, `sunglass` and `sunglass-round`).

### RegisterExtraFemaleAccessory

Register extra female accessory (`necklace` and `choker`).

## Custom Avatar And Extending

Avatar generation is based on separated svg elements. You can customize current avatar style or add new styles to avatar package using helpers method.

You can work on [template.ai](https://github.com/gomig/avatar/blob/master/template.ai) file and save your work to svg then add your generated elements parts to avatar packages. Each shape must have some placeholder for replace colors or id in final result.

### Colors

Elements color are based on palette. Color value can be simple hex (`Hex`) or gradient (`Definition`).

There are three type of colors:

- **HairColor:** contains _base_, _shadow_ and _highlight_ color.
- **SkinColor:** contains _skin_ and _shadow_ color.
- **Palette:** contains _shape_, _text_, _dress_, _dressShadow_ and _decorator_ color.

**Caution:** for gradient colors, gradient element must contains `id="{id}"` placeholder attribute. otherwise color not work!

#### ResetHairColor

Clear all registered hair color and set hair colors to pre-defined (`Default`, `brown`, `light_brown` and `dark_brown`) value.

#### AddHairColor

Add new hair color.

```go
import "github.com/gomig/avatar"

func main(){
    avatar.AddHairColor(
        "brown",
        avatar.Color{avatar.Hex, "#432818"},
        avatar.Color{avatar.Hex, "#140b06"},
        avatar.Color{avatar.Hex, "#884139"},
    )
}
```

#### SetDefaultHairColor

Change default hair color.

#### ResetSkinColor

Clear all registered skin color and set skin colors to pre-defined (`Default`, `white`, `brown` and `black`) value.

#### AddSkinColor

Add new skin color.

```go
import "github.com/gomig/avatar"

func main(){
    avatar.AddHairColor(
        "blond",
        avatar.Color{
            avatar.Definition,
            `<linearGradient id="{id}" gradientTransform="rotate(90)">
                <stop offset="5%" stop-color="#fffa00" />
                <stop offset="95%" stop-color="#cc8800" />
            </linearGradient>`,
        },
        avatar.Color{avatar.Hex, "#84706d"})
}
```

#### SetDefaultSkinColor

Change default skin color.

#### ResetPalette

Clear all registered palette and keep `Default` palette only.

#### AddPalette

Add new palette.

```go
import "github.com/gomig/avatar"

func main(){
    avatar.AddPalette(
        "blue",
        avatar.Color{avatar.Hex, "#a2cced"},
        avatar.Color{avatar.Hex, "#162835"},
        avatar.Color{avatar.Hex, "#006fc8"},
        avatar.Color{avatar.Hex, "#045389"},
        avatar.Color{avatar.Hex, "#121b21"},
    )
}
```

#### SetDefaultPalette

Change default palette.

### Available Placeholders

You can use following placeholder in your element to replace with dynamic generated data in avatar.

- `{id}` must included in gradient colors to pass dynamic id.
- `{skin}` person skin color (used in body shape).
- `{skin_shadow}` dark skin color (used in shadow parts of body shape).
- `{hair}` hair color (used in hair, eyebrow and facialHair).
- `{hair_shadow}` dark parts of hair (used in hair, eyebrow and facialHair).
- `{hair_highlight}` highlight parts of hair (used in hair, eyebrow and facialHair).
- `{shape}` background shape color (used in background shape fill).
- `{text}` text and sticker color.
- `{dress}` main dress color.
- `{dress_shadow}` darken parts color of dress.
- `{decorator}` decorator color (tie, glass frame, necklace, etc.).

### Background Shape

**Note:** You can set mask for each background shape. mask shape must has `white` fill.

#### ClearShapes

Clear all registered shapes.

#### AddShape

Register new background shape.

```go
import "github.com/gomig/avatar"

func main(){
    avatar.AddShape(
        "circle",
        "<circle fill="{shape}" cx="64" cy="64" r="64"/>",
        "<path fill="white" d="M0,0v64c0,35.3,28.7,64,64,64s64-28.7,64-64V0H0z"/>")
}
```

### Letter

You can register shape for each unicode letters or you can transform letters to english.

#### AddLetterMap

Define new unicode letter for transform to english letter or letter with shape.

```go
import "github.com/gomig/avatar"

func main(){
    avatar.AddLetterMap("ج", `j`)
}
```

#### AddLetterShape

Define shape for letter. Letter must have one character length! By default english letter shapes are registered but you can override them by re-set letter shape.

```go
import "github.com/gomig/avatar"

func main(){
    avatar.AddLetterShape("a", `<path fill="{text}" d="..."/>`)
}
```

#### SetDefaultLetter

Set fallback default letter shape. If no shape or transform for character found this shape used!

### Sticker

#### ResetSticker

Clear all registered stickers.

```go
import "github.com/gomig/avatar"

func main(){
    avatar.ResetSticker()
}
```

#### AddSticker

Register new sticker.

```go
import "github.com/gomig/avatar"

func main(){
    avatar.AddSticker("star", `<path fill="{text}" d="..."/>`)
}
```

### Person Avatar

Person avatar can be male or female. all type of data can registered for `Male`, `Female` or `Both`.

#### ResetBody

Set character body shape to default.

#### SetBody

Set character body shape.

```go
import "github.com/gomig/avatar"

func main(){
    avatar.SetBody(`<path fill="{skin}" ...>`)
}
```

#### ResetFacialHairs

Clear all registered facial hairs.

#### AddFacialHair

Add new facial hair.

```go
import "github.com/gomig/avatar"

func main(){
    avatar.SetBody("wavy", `<path fill="{hair}" ...>`)
}
```

#### ResetMaleHairs

Clear all registered male hairs (except `Default`).

#### ResetFemaleHairs

Clear all registered female hairs (except `Default`).

#### AddHair

Register new hair style.

```go
import "github.com/gomig/avatar"

func main(){
    avatar.SetBody("wavy", avatar.Male, `<path fill="{hair}" ...>`)
}
```

#### SetDefaultHair

Set default hair.

#### ResetMaleDresses

Clear all registered male dresses (except `Default`).

#### ResetFemaleDresses

Clear all registered female dresses (except `Default`).

#### AddDress

Register new dress.

```go
import "github.com/gomig/avatar"

func main(){
    avatar.AddCloth(
        "polo",
        avatar.Both,
        `<g>
            <path fill="{dress}" ...>
            <path fill="{dress_shadow}" ...>
        </g>`,
    )
}
```

#### SetDefaultDress

Set default dress shape.

#### ResetMaleEyes

Clear all registered male eye (except `Default`).

#### ResetFemaleEyes

Clear all registered female eye (except `Default`).

#### AddEye

Register new eye.

```go
import "github.com/gomig/avatar"

func main(){
    avatar.SetBody("asian", avatar.Male, `<path ...>`)
}
```

#### SetDefaultEye

Set default eye shpae.

#### ResetMaleEyebrows

Clear all registered male eyebrow (except `Default`).

#### ResetFemaleEyebrows

Clear all registered female eyebrow (except `Default`).

#### AddEyebrow

Register new eyebrow.

```go
import "github.com/gomig/avatar"

func main(){
    avatar.SetBody("wavy", avatar.Male, `<path fill="{hair_shadow}" ...>`)
}
```

#### SetDefaultEyebrow

Set default eyebrow shape.

#### ResetMaleMouths

Clear all registered male mouths (except `Default`).

#### ResetFemaleMouths

Clear all registered female mouths (except `Default`).

#### AddMouth

Register new mouth.

```go
import "github.com/gomig/avatar"

func main(){
    avatar.AddMouth("red_lip", avatar.Female, `<path ...>`)
}
```

#### SetDefaultMouth

Set default mouth.

#### ResetMaleGlasses

Delete all registered male glasses.

#### ResetFemaleGlasses

Delete all registered female glasses.

#### AddGlass

Register new glass.

#### ResetMaleAccessories

Clear all registered male accessory.

#### ResetFemaleAccessories

Clear all registered female accessory.

#### AddAccessory

Add new accessory.

```go
import "github.com/gomig/avatar"

func main(){
    avatar.SetBody("necklace", avatar.Female, `<path fill="{decorator}" ...>`)
}
```
