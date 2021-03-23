package request

type NewItem struct {
	Name string `json:"name" binding:"required,max=255"`
}

type DeleteItem struct {
	Id int `uri:"id" binding:"required"`
}

type SendItem struct {
	ItemId    int    `json:"item_id" binding:"required"`
	UserLogin string `json:"user_login" binding:"required,max=60"`
}

type Confirm struct {
	Link string `uri:"link" binding:"required"`
}
