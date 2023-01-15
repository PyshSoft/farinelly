package entity

type Profile struct {
	Id                string
	UserName          string
	UserAge           int
	Sex               Sex
	Location          string
	About             string
	Pictures          []Picture
	Interests         []string
	SearchConditions  SearchConditions
	ViewedProfileIds  []string
	MatchedProfileIds []string
}

type Sex uint8

const (
	Male Sex = iota
	Female
)

func (g Sex) String() string {
	if g == 0 {
		return "MALE"
	}
	return "FEMALE"
}

type Picture []byte

type SearchConditions struct {
	Sex Sex
}
