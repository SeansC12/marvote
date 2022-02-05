package model

type CharacterInfo struct {
	Id    string `bson:"_id,omitempty" json:"id,omitempty"`
	Name  string `json:"name"`
	Aka   string `json:"aka"`
	Votes int64  `json:"votes,omitempty"`
}
