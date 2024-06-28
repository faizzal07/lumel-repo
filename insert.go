package main

import (
	"context"
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/jackc/pgx/v4/pgxpool"
)

type Customer struct {
	ID      int
	Name    string
	Email   string
	Address string
}

type Product struct {
	ID       int
	Name     string
	Category string
}

type Order struct {
	ID              int
	CustomerID      int
	ProductID       int
	ProductName     string
	Category        string
	Region          string
	DateOfSale      string
	QuantitySold    int
	UnitPrice       float64
	Discount        float64
	ShippingCost    float64
	PaymentMethod   string
	CustomerName    string
	CustomerEmail   string
	CustomerAddress string
}

func main() {
	connString := "postgres://khalith@localhost:5432/lumeldb?sslmode=disable"
	poolConfig, err := pgxpool.ParseConfig(connString)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to parse connection string: %v\n", err)
		os.Exit(1)
	}

	pool, err := pgxpool.ConnectConfig(context.Background(), poolConfig)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer pool.Close()

	filePath := "/Users/khalith/Documents/lumel/data.csv"
	csvFile, err := os.Open(filePath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening CSV file: %v\n", err)
		os.Exit(1)
	}
	defer csvFile.Close()

	reader := csv.NewReader(csvFile)
	reader.FieldsPerRecord = -1
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading CSV file: %v\n", err)
		os.Exit(1)
	}

	var customers []Customer
	var products []Product
	var orders []Order

	for i, record := range records {
		if i == 0 {
			// Skip header row
			continue
		}

		order := Order{
			ID:              parseInt(record[0]),
			ProductID:       parseInt(record[1]),
			CustomerID:      parseInt(record[2]),
			ProductName:     record[3],
			Category:        record[4],
			Region:          record[5],
			DateOfSale:      record[6],
			QuantitySold:    parseInt(record[7]),
			UnitPrice:       parseFloat(record[8]),
			Discount:        parseFloat(record[9]),
			ShippingCost:    parseFloat(record[10]),
			PaymentMethod:   record[11],
			CustomerName:    record[12],
			CustomerEmail:   record[13],
			CustomerAddress: record[14],
		}

		orders = append(orders, order)

		// Collect unique customers and products
		customers = appendUniqueCustomer(customers, Customer{
			ID:      order.CustomerID,
			Name:    order.CustomerName,
			Email:   order.CustomerEmail,
			Address: order.CustomerAddress,
		})

		products = appendUniqueProduct(products, Product{
			ID:       order.ProductID,
			Name:     order.ProductName,
			Category: order.Category,
		})
	}

	// Insert customers into database
	for _, customer := range customers {
		_, err := pool.Exec(context.Background(), `
            INSERT INTO customers (customer_id, customer_name, customer_email, customer_address)
            VALUES ($1, $2, $3, $4)
            ON CONFLICT (customer_id) DO NOTHING;
        `, customer.ID, customer.Name, customer.Email, customer.Address)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error inserting customer: %v\n", err)
			continue
		}
	}

	// Insert products into database
	for _, product := range products {
		_, err := pool.Exec(context.Background(), `
            INSERT INTO products (product_id, product_name, category)
            VALUES ($1, $2, $3)
            ON CONFLICT (product_id) DO NOTHING;
        `, product.ID, product.Name, product.Category)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error inserting product: %v\n", err)
			continue
		}
	}

	// Insert orders into database
	for _, order := range orders {
		_, err := pool.Exec(context.Background(), `
            INSERT INTO orders (order_id, customer_id, product_id, date_of_sale, quantity_sold, unit_price, discount, shipping_cost, payment_method, region)
            VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10);
        `, order.ID, order.CustomerID, order.ProductID, order.DateOfSale, order.QuantitySold, order.UnitPrice, order.Discount, order.ShippingCost, order.PaymentMethod, order.Region)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error inserting order: %v\n", err)
			continue
		}
	}
}

func parseInt(s string) int {
	i, _ := strconv.Atoi(strings.TrimSpace(s))
	return i
}

func parseFloat(s string) float64 {
	f, _ := strconv.ParseFloat(strings.TrimSpace(s), 64)
	return f
}

func appendUniqueCustomer(customers []Customer, customer Customer) []Customer {
	for _, c := range customers {
		if c.ID == customer.ID {
			return customers
		}
	}
	return append(customers, customer)
}

func appendUniqueProduct(products []Product, product Product) []Product {
	for _, p := range products {
		if p.ID == product.ID {
			return products
		}
	}
	return append(products, product)
}
