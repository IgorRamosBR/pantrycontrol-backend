package entities

type Product struct {
	tableName struct{} `pg:"products,alias:p"`
	Id        int      `json:"id"`
	Name      string   `json:"name"`
	Unit      string   `json:"unit"`
	Brand     string   `json:"brand"`
	Category  string   `json:"category"`
}
