package request

type NewItem struct {
	Name string `json:"name" binding:"required,max=255"`
}

type DeleteItem struct {
	Id int `uri:"id" binding:"required"`
}

type SendItem struct {
	ItemId    int    `json:"id" binding:"required"`
	UserLogin string `json:"login" binding:"required,max=60"`
}

type Confirm struct {
	Link string `json:"temp_link" binding:"required"`
}
