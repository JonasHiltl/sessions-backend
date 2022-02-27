package dto

type Story struct {
	Id            string   `json:"id"`
	PId           string   `json:"partyId"`
	UId           string   `json:"userId"`
	Lat           float64  `json:"lat"`
	Long          float64  `json:"long"`
	Url           string   `json:"url"`
	TaggedFriends []string `json:"tagged_friends"`
}
