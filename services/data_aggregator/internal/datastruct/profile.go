package datastruct

type FriendshipStatus struct {
	IsFriend        bool `json:"is_friend"`
	OutgoingRequest bool `json:"outgoing_request"`
}

type AggregatedProfile struct {
	Id               string           `json:"id,omitempty"`
	Username         string           `json:"username,omitempty"`
	Firstname        string           `json:"firstname,omitempty"`
	Lastname         string           `json:"lastname,omitempty"`
	Avatar           string           `json:"avatar,omitempty"`
	FriendCount      int64            `json:"friend_count,omitempty"`
	FriendshipStatus FriendshipStatus `json:"friendship_status,omitempty"`
}
