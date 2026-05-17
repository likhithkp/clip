package dto

type CreateUrlDto struct {
	Title   string `json:"title"`
	LongUrl string `json:"longUrl"`
	Code    string `json:"code"`
}
