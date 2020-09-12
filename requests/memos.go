package requests

type CreateMemo struct {
	Content string `json:"content"`
}

type UpdateMemo struct {
	Content string `json:"content"`
}
