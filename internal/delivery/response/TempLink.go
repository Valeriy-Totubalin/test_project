package response

type TempLink struct {
	Link string `json:"link" binding:"required"`
}
