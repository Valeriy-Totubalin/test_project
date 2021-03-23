package response

type TempLink struct {
	Link string `json:"temp_link" binding:"required"`
}
