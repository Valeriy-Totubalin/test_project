package request

type NewItem struct {
	Name string `json:"name" binding:"required,max=255"`
}

type DeleteItem struct {
	Id int `uri:"id" binding:"required,int"`
}
