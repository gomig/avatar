package avatar

var bodyShape string                    // common
var facialHairs map[string]string       // males
var maleHairs map[string]string         // males
var maleDresses map[string]string       // males
var maleEyes map[string]string          // males
var maleEyebrows map[string]string      // males
var maleMouths map[string]string        // males
var maleGlasses map[string]string       // males
var maleAccessories map[string]string   // males
var femaleHairs map[string]string       // females
var femaleDresses map[string]string     // females
var femaleEyes map[string]string        // females
var femaleEyebrows map[string]string    // females
var femaleMouths map[string]string      // females
var femaleGlasses map[string]string     // females
var femaleAccessories map[string]string // females

func init() {
	ResetBody()
	ResetFacialHairs()
	ResetMaleHairs()
	ResetFemaleHairs()
	ResetMaleDresses()
	ResetFemaleDresses()
	ResetMaleEyes()
	ResetFemaleEyes()
	ResetMaleEyebrows()
	ResetFemaleEyebrows()
	ResetMaleMouths()
	ResetFemaleMouths()
	ResetMaleGlasses()
	ResetFemaleGlasses()
	ResetMaleAccessories()
	ResetFemaleAccessories()
}

// ResetBody set character body to default
func ResetBody() {
	bodyShape = defaultBody
}

// SetBody set character body shape
func SetBody(shape string) {
	bodyShape = shape
}

// ResetFacialHairs delete all facial hairs
func ResetFacialHairs() {
	facialHairs = map[string]string{None: ""}
}

// AddFacialHair add new facial hair
func AddFacialHair(name string, shape string) {
	facialHairs[name] = shape
}

// ResetMaleHairs delete all male hairs except default
func ResetMaleHairs() {
	maleHairs = map[string]string{Default: maleDefaultHair}
}

// ResetFemaleHairs delete all female hairs except default
func ResetFemaleHairs() {
	femaleHairs = map[string]string{Default: femaleDefaultHair}
}

// AddHair add hairs
func AddHair(name string, gender Gender, shape string) {
	if gender == Male || gender == Both {
		maleHairs[name] = shape
	}
	if gender == Female || gender == Both {
		femaleHairs[name] = shape
	}
}

// SetDefaultHair set default hair
func SetDefaultHair(gender Gender, shape string) {
	AddHair(Default, gender, shape)
}

// ResetMaleDresses delete all male dress except default
func ResetMaleDresses() {
	maleDresses = map[string]string{Default: maleDefaultDress}
}

// ResetFemaleDresses delete all female dress except default
func ResetFemaleDresses() {
	femaleDresses = map[string]string{Default: femaleDefaultDress}
}

// AddDress add new dress
func AddDress(name string, gender Gender, shape string) {
	if gender == Male || gender == Both {
		maleDresses[name] = shape
	}
	if gender == Female || gender == Both {
		femaleDresses[name] = shape
	}
}

// SetDefaultDress set default dress
func SetDefaultDress(gender Gender, shape string) {
	AddDress(Default, gender, shape)
}

// ResetMaleEyes delete all male eye except default
func ResetMaleEyes() {
	maleEyes = map[string]string{Default: maleDefaultEye}
}

// ResetFemaleEyes delete all female eye except default
func ResetFemaleEyes() {
	femaleEyes = map[string]string{Default: femaleDefaultEye}
}

// AddEye add new eye
func AddEye(name string, gender Gender, shape string) {
	if gender == Male || gender == Both {
		maleEyes[name] = shape
	}
	if gender == Female || gender == Both {
		femaleEyes[name] = shape
	}
}

// SetDefaultEye set default eye
func SetDefaultEye(gender Gender, shape string) {
	AddEye(Default, gender, shape)
}

// ResetMaleEyebrows delete all male eyebrow except default
func ResetMaleEyebrows() {
	maleEyebrows = map[string]string{Default: maleDefaultEyebrow}
}

// ResetFemaleEyebrows delete all female eyebrow except default
func ResetFemaleEyebrows() {
	femaleEyebrows = map[string]string{Default: femaleDefaultEyebrow}
}

// AddEyebrow add new eyebrow
func AddEyebrow(name string, gender Gender, shape string) {
	if gender == Male || gender == Both {
		maleEyebrows[name] = shape
	}
	if gender == Female || gender == Both {
		femaleEyebrows[name] = shape
	}
}

// SetDefaultEyebrow set default eyebrow
func SetDefaultEyebrow(gender Gender, shape string) {
	AddEyebrow(Default, gender, shape)
}

// ResetMaleMouths delete all male mouth except default
func ResetMaleMouths() {
	maleMouths = map[string]string{Default: maleDefaultMouth}
}

// ResetFemaleMouths delete all female mouth except default
func ResetFemaleMouths() {
	femaleMouths = map[string]string{Default: femaleDefaultMouth}
}

// AddMouth add new mouth
func AddMouth(name string, gender Gender, shape string) {
	if gender == Male || gender == Both {
		maleMouths[name] = shape
	}
	if gender == Female || gender == Both {
		femaleMouths[name] = shape
	}
}

// SetDefaultMouth set default mouth
func SetDefaultMouth(gender Gender, shape string) {
	AddMouth(Default, gender, shape)
}

// ResetMaleGlasses delete all male glasses
func ResetMaleGlasses() {
	maleGlasses = map[string]string{None: ""}
}

// ResetFemaleGlasses delete all female glasses
func ResetFemaleGlasses() {
	femaleGlasses = map[string]string{None: ""}
}

// AddGlass add new glass
func AddGlass(name string, gender Gender, shape string) {
	if gender == Male || gender == Both {
		maleGlasses[name] = shape
	}
	if gender == Female || gender == Both {
		femaleGlasses[name] = shape
	}
}

// ResetMaleAccessories delete all male accessory
func ResetMaleAccessories() {
	maleAccessories = map[string]string{None: ""}
}

// ResetFemaleAccessories delete all female accessory
func ResetFemaleAccessories() {
	femaleAccessories = map[string]string{None: ""}
}

// AddAccessory add new accessory
func AddAccessory(name string, gender Gender, shape string) {
	if gender == Male || gender == Both {
		maleAccessories[name] = shape
	}
	if gender == Female || gender == Both {
		femaleAccessories[name] = shape
	}
}
