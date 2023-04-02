package model

type Product struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	Price      int    `json:"price"`
	CategoryID string `json:"category_id"`
}
type Category struct {
	ID            string `json:"id"`
	Name          string `json:"name"`
	ParentID      string `json:"parent_id"`
	SubCategories any    `json:"sub_categories"`
}
type Comission struct {
	Balance int `json:"balance"`
}
type ShopCart struct {
	ID        string `json:"id"`
	ProductID string `json:"productId"`
	UserID    string `json:"userID"`
	Count     int    `json:"count"`
	Status    bool   `json:"status"`
	Date      string `json:"date"`
}
type User struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Balance int    `json:"balance"`
}

type Products struct{
	Prod []Product
}