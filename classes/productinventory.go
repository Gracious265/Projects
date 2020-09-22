package classes


type Product struct {
	qty int
	productId string
	price float64
}

type Inventory struct {
	name string
	products []Product
}

func (inventory *Inventory) addProduct(qty int, id string, price float64) {
	product := Product{productId:id, qty: qty, price: price}
	inventory.products = append(inventory.products, product)	
}

func (inventory *Inventory) removeProduct(id string){
	for index, value := range inventory.products{
		if value.productId == id {
			inventory.products = append(inventory.products[:index], inventory.products[index+1:]...)
			break
		}
	}

}

// TODO: Make this work. and TEst it all.
func (product *Product) update() {

}