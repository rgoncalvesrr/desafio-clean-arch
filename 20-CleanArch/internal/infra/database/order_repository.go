package database

import (
	"database/sql"

	"github.com/devfullcycle/20-CleanArch/internal/entity"
)

type OrderRepository struct {
	Db *sql.DB
}

func NewOrderRepository(db *sql.DB) *OrderRepository {
	return &OrderRepository{Db: db}
}

func (r *OrderRepository) Save(order *entity.Order) error {
	stmt, err := r.Db.Prepare("INSERT INTO orders (id, price, tax, final_price) VALUES (?, ?, ?, ?)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(order.ID, order.Price, order.Tax, order.FinalPrice)
	if err != nil {
		return err
	}
	return nil
}

func (r *OrderRepository) GetTotal() (int, error) {
	var total int
	err := r.Db.QueryRow("Select count(*) from orders").Scan(&total)
	if err != nil {
		return 0, err
	}
	return total, nil
}

func (r *OrderRepository) FindAll() ([]entity.Order, error) {

	orders := []entity.Order{}

	rows, err := r.Db.Query(`
		select 
			id, price, tax
		from
			orders`)

	if err != nil {
		return orders, err
	}

	defer rows.Close()

	for rows.Next() {

		var id string
		var price float64
		var tax float64

		if err := rows.Scan(&id, &price, &tax); err != nil {
			return orders, err
		}

		order, err := entity.NewOrder(id, price, tax)

		if err != nil {
			return orders, err
		}

		if err = order.CalculateFinalPrice(); err != nil {
			return orders, err
		}

		orders = append(orders, *order)
	}

	return orders, nil
}
