package datastruct

import ug "github.com/jonashiltl/sessions-backend/packages/grpc/user"

type AggregatedComment struct {
	Id        string      `json:"id"`
	PartyId   string      `json:"party_id"`
	Author    *ug.Profile `json:"author,omitempty"`
	Body      string      `json:"body"`
	CreatedAt string      `json:"created_at"`
}

type PagedAggregatedComment struct {
	Comments []AggregatedComment `json:"comments,omitempty"`
	NextPage string              `json:"nextPage"`
}

type AggregatedReply struct {
	Id        string      `json:"id"`
	CommentId string      `json:"comment_id"`
	Author    *ug.Profile `json:"author,omitempty"`
	Body      string      `json:"body"`
	CreatedAt string      `json:"created_at"`
}

type PagedAggregatedReply struct {
	Replies  []AggregatedReply `json:"replies,omitempty"`
	NextPage string            `json:"nextPage"`
}
