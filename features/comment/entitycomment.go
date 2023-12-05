package comment

type Comment struct {
	Posting_ID uint   `json:"posting_id"`
	User_ID    uint   `json:"user_id"`
	Comment    string `json:"comment"`
}
