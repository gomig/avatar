package avatar

// NewTextAvatar generate new text based avatar instance
func NewTextAvatar(name string) *Avatar {
	avatar := new(Avatar)
	avatar._type = LETTER
	avatar.name = name
	avatar.init()
	return avatar
}

// NewStickerAvatar generate new sticker based avatar instance
func NewStickerAvatar(sticker string) *Avatar {
	avatar := new(Avatar)
	avatar._type = STICKER
	avatar.sticker = sticker
	avatar.init()
	return avatar
}

// NewPersonAvatar generate new character based avatar instance
func NewPersonAvatar(isMale bool) *Avatar {
	avatar := new(Avatar)
	if isMale {
		avatar._type = MALE
	} else {
		avatar._type = FEMALE
	}
	avatar.init()
	return avatar
}
