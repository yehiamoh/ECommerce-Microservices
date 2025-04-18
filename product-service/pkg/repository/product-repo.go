package repository

import (
	"database/sql"
	"fmt"
	pb "product-service/gen/product"
	"strconv"
)

type ProductRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB)*ProductRepository{
	return &ProductRepository{db: db}
}

func (p *ProductRepository) CreateProduct(name string, description string, price float32) (*pb.ProductResponse, error) {
    // Check if repository or db is nil (defensive programming)
    if p == nil || p.db == nil {
        return nil, fmt.Errorf("repository or database connection is nil")
    }

    tx, err := p.db.Begin()
    if err != nil {
        fmt.Println("error begin tx:", err)
        return nil, err
    }
    defer tx.Rollback()

    query := `
        INSERT INTO product(name, description, price)
        VALUES ($1, $2, $3)
        RETURNING id
    `
    
    var ID int
    // FIX: Pass pointer to ID with &ID
    err = tx.QueryRow(query, name, description, price).Scan(&ID)
    if err != nil {
        fmt.Println("error in query:", err)
        return nil, err
    }

    err = tx.Commit()
    if err != nil {
        fmt.Println("error committing tx:", err)
        return nil, err
    }

    product := &pb.Product{
        Id:          strconv.Itoa(ID),
        Name:        name,
        Description: description,
        Price:       price,
    }
    
    return &pb.ProductResponse{
        Product: product,
    }, nil
}
func (p *ProductRepository) GetProductByID(id int) (*pb.ProductResponse, error) {
	product := &pb.Product{}

	query := `
		SELECT id, name, description, price FROM product
		WHERE id = $1
	`
	err := p.db.QueryRow(query, id).Scan(&product.Id, &product.Name, &product.Description, &product.Price)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &pb.ProductResponse{
		Product: product,
	}, nil
}
func (p *ProductRepository) UpdateProduct(id int, name *string, description *string, price *float32) (*pb.ProductResponse, error) {
	if p == nil || p.db == nil {
		return nil, fmt.Errorf("repository or db is nil")
	}

	// Step 1: Get the existing product
	existing, err := p.GetProductByID(id)
	if err != nil {
		return nil, err
	}
	if existing == nil {
		return nil, fmt.Errorf("product with ID %d not found", id)
	}

	// Step 2: Update only the provided fields
	product := existing.Product

	if name != nil {
		product.Name = *name
	}
	if description != nil {
		product.Description = *description
	}
	if price != nil {
		product.Price = *price
	}

	// Step 3: Run the update
	query := `
		UPDATE product
		SET name = $1, description = $2, price = $3
		WHERE id = $4
	`

	_, err = p.db.Exec(query, product.Name, product.Description, product.Price, id)
	if err != nil {
		return nil, err
	}

	return &pb.ProductResponse{Product: product}, nil
}

func (p *ProductRepository) DeleteProduct(id int) (*pb.ProductResponse, error) {
	if p == nil || p.db == nil {
		return nil, fmt.Errorf("repository or database connection is nil")
	}

	// Get the product before deleting it
	existingProduct, err := p.GetProductByID(id)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch product before delete: %w", err)
	}
	if existingProduct == nil {
		return nil, fmt.Errorf("no product found with ID %d", id)
	}

	tx, err := p.db.Begin()
	if err != nil {
		fmt.Println("error starting transaction:", err)
		return nil, err
	}
	defer tx.Rollback()

	query := `DELETE FROM product WHERE id = $1`
	result, err := tx.Exec(query, id)
	if err != nil {
		fmt.Println("error deleting product:", err)
		return nil, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		fmt.Println("error fetching rows affected:", err)
		return nil, err
	}
	if rowsAffected == 0 {
		return nil, fmt.Errorf("no product found with ID %d", id)
	}

	err = tx.Commit()
	if err != nil {
		fmt.Println("error committing transaction:", err)
		return nil, err
	}

	// Return the deleted product info
	return existingProduct, nil
}
