package demo

// DemoConst be available to packages other than the current one via import since it starts with a capital character
const DemoConst = 21 // Basically like `public final int DemoConst = 1` for those familiar with Java

// Since this const starts with a lower-case character, it'll only be visible within this package
const constForOurselves = 42

// PublicStruct will be available to other packages that import this package since it starts with a capital character
type PublicStruct struct {
	// We still need to export public fields (this is important for things like JSON decoding/encoding)
	FieldWeWantPeopleToSee string

	// Even though the struct is exported, we can still have package-only fields
	// This field will not be decoded/encoded in JSON because the encoding/json package won't be able to see it
	fieldWeDontWantPeopleToSee string
}

// This is an unexported struct because it starts with a lower-case character
type demoStruct struct {
	// We can still have exported fields (this is important for things like JSON decoding/encoding)
	FieldWeWantPeopleToSee string
	// This field will not be decoded/encoded in JSON because the encoding/json package won't be able to see it
	fieldWeDontWantPeopleToSee string
}
