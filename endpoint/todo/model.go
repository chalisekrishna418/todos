package todo

//Todo struct
type Todo struct {
	Item   string `bson:"item"`
	Status string `bson:"status"`
	TodoID string `bson:"todoId"`
}
