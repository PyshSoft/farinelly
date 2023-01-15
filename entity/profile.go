package entity

type Profile struct {
	Id string
	UserName string
	UserAge string
	Gender Gender
	Location string
	About string
	Pictures []Picture
	Interests []string
	SearchConditions SearchConditions
	ViewedProfileIds []string
	MatchedProfileIds []string
}

type Gender uint8
const (
	Male Gender = iota
	Female
)
func (g Gender) String() string {
	if g == 0 { return "MALE" }
	return "FEMALE"
}

type Picture []byte

type SearchConditions struct {
	Gender Gender
}