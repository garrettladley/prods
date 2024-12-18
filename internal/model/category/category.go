package category

type Category string

const (
	Electronics    Category = "electronics"
	Apparel        Category = "apparel"
	HomeGoods      Category = "home goods"
	Sports         Category = "sports"
	Beauty         Category = "beauty"
	Grocery        Category = "grocery"
	OfficeSupplies Category = "office supplies"
	Outdoor        Category = "outdoor"
	Toys           Category = "toys"
	Health         Category = "health"
	Automotive     Category = "automotive"
	Luxury         Category = "luxury"
	Books          Category = "books"
)

var Categories = [13]Category{
	Electronics,
	Apparel,
	HomeGoods,
	Sports,
	Beauty,
	Grocery,
	OfficeSupplies,
	Outdoor,
	Toys,
	Health,
	Automotive,
	Luxury,
	Books,
}
