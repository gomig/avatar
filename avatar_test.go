package avatar_test

import (
	"fmt"
	"testing"

	"github.com/gomig/avatar"
)

func TestLetter(t *testing.T) {
	avatar.RegisterExtraPalettes()
	avatar.RegisterPersiaShapes()

	persian := avatar.NewTextAvatar("مینا")
	persian.RandomizePalette()
	persian.RandomizeShape()
	if err := persian.SaveAs("./demo/persian.svg"); err != nil {
		t.Fatal(err)
	}

	john := avatar.NewTextAvatar("John Doe")
	john.RandomizePalette()
	john.RandomizeShape()
	if err := john.SaveAs("./demo/john.svg"); err != nil {
		t.Fatal(err)
	}

	mike := avatar.NewTextAvatar("Mike")
	mike.RandomizePalette()
	mike.RandomizeShape()
	if err := mike.SaveAs("./demo/mike.svg"); err != nil {
		t.Fatal(err)
	}
}

func TestSticker(t *testing.T) {
	avatar.RegisterExtraPalettes()
	avatar.AddSticker("apple", `<path fill="{text}" d="M91.3,52.6c-0.4,0.3-7.7,4.5-7.7,13.6c0,10.6,9.3,14.4,9.6,14.5c0,0.2-1.5,5.1-4.9,10.2c-3.1,4.4-6.3,8.8-11.1,8.8c-4.9,0-6.1-2.8-11.7-2.8c-5.5,0-7.4,2.9-11.9,2.9s-7.6-4.1-11.1-9.1c-4.1-5.9-7.5-15-7.5-23.7c0-13.9,9-21.3,18-21.3c4.7,0,8.7,3.1,11.6,3.1c2.8,0,7.2-3.3,12.6-3.3C79.1,45.5,86.5,45.7,91.3,52.6L91.3,52.6z M74.5,39.6c2.2-2.6,3.8-6.3,3.8-10c0-0.5,0-1-0.1-1.4c-3.6,0.1-7.9,2.4-10.5,5.4c-2,2.3-3.9,6-3.9,9.7c0,0.6,0.1,1.1,0.1,1.3c0.2,0,0.6,0.1,1,0.1C68.1,44.7,72.2,42.5,74.5,39.6L74.5,39.6z"/>`)

	apple := avatar.NewStickerAvatar("apple")
	apple.RandomizeShape()
	apple.RandomizePalette()
	if err := apple.SaveAs("./demo/apple.svg"); err != nil {
		t.Fatal(err)
	}

	star := avatar.NewStickerAvatar(avatar.Default)
	star.RandomizeShape()
	star.RandomizePalette()
	if err := star.SaveAs("./demo/star.svg"); err != nil {
		t.Fatal(err)
	}
}

func TestMale(t *testing.T) {
	avatar.RegisterExtraPalettes()
	avatar.RegisterExtraMaleHair()
	avatar.RegisterExtraMaleDress()
	avatar.RegisterExtraMaleGlass()
	avatar.RegisterExtraFacialHair()
	avatar.RegisterExtraMaleAccessory()

	for i := 1; i < 6; i++ {
		av := avatar.NewPersonAvatar(true)
		av.RandomizeSkinColor()
		av.RandomizePalette()
		av.RandomizeShape()
		if err := av.SaveAs(fmt.Sprintf("./demo/male%d.svg", i)); err != nil {
			t.Fatal(err)
		}
	}

}

func TestFemale(t *testing.T) {
	avatar.RegisterExtraPalettes()
	avatar.RegisterExtraFemaleHair()
	avatar.RegisterExtraFemaleDress()
	avatar.RegisterExtraFemaleGlass()
	avatar.RegisterExtraFemaleAccessory()

	for i := 1; i < 6; i++ {
		av := avatar.NewPersonAvatar(false)
		av.RandomizeSkinColor()
		av.RandomizePalette()
		av.RandomizeShape()
		if err := av.SaveAs(fmt.Sprintf("./demo/female%d.svg", i)); err != nil {
			t.Fatal(err)
		}
	}
}
