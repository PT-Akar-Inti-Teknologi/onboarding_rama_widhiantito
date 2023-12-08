package product

type Product struct {
	ID    uint   `gorm:"primaryKey" json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
	Stock int    `json:"stock"`
}

type ProductList struct {
	ProductID uint `json:"id_product"`
	Quantity  int  `json:"quantity"`
}
