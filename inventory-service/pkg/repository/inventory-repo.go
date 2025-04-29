package repository

import (
	"database/sql"
	"fmt"
	pb "inventory-service/gen/inventory"
)

type InventoryRepository struct {
	db *sql.DB
}
func NewInventoryRepository(db *sql.DB)*InventoryRepository{
	return &InventoryRepository{
		db: db,
	}
}
func(r *InventoryRepository)GetStockOfProduct(product_id string)(*pb.CheckStockResponse,error){
	tx,err:=r.db.Begin()
	if err!=nil{
		return nil,err
	}
	defer tx.Rollback()
	query:=`
	select product_id,quantity from inventory
	where product_id=$1
	`
	res:=&pb.CheckStockResponse{}

	err=tx.QueryRow(query,product_id).Scan(&res.ProductId,&res.AvailableQuantity)
	if err!=nil{
		return nil,err
	}

	if err=tx.Commit();err!=nil{
		return nil,err
	}
	return res,nil
}
func(r *InventoryRepository)CreateNewStock(product_id string,quantityToAdd int32)(*pb.AddStockResponse,error){
	tx,err:=r.db.Begin()
	if err!=nil{
		return nil,err
	}
	defer tx.Rollback()
	query:=`
	INSERT INTO inventory (product_id,quantity) Values($1,$2)
	`
	_,err=tx.Exec(query,product_id,quantityToAdd)
	if err != nil {
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return &pb.AddStockResponse{
		Message: "Stock added successfully",
	}, nil
}
func (r *InventoryRepository)DecreaseStockOfProduct(product_id string,qunatityToDecrease int32)(*pb.DecreaseStockResponse,error){
	tx,err:=r.db.Begin()
	if err!=nil{
		return nil,err
	}
	defer tx.Rollback()
	var currentQuantity int32
	// FOR UPDATE lock the row to prevent race condition 
	err=tx.QueryRow(`SELECT quantity from inventory where product_id=$1 FOR UPDATE`,product_id).Scan(&currentQuantity)
	if err!=nil{
		return nil,err
	}

	if currentQuantity<qunatityToDecrease{
		return nil,fmt.Errorf("insufficient stock : request:%v ,available:%v ",currentQuantity,qunatityToDecrease)
	}
	var remainingQuantity int32
	err=tx.QueryRow(`UPDATE inventory SET quantity = quantity-$1 WHERE product_id=$2 RETURNING quantity`,qunatityToDecrease,product_id).Scan(&remainingQuantity)
	if err!=nil{
		return nil,err
	}
	if err:=tx.Commit();err!=nil{
		return nil,err
	}
	return &pb.DecreaseStockResponse{
		Message: "Stock decreased successfully",
		RemainingQuantity: remainingQuantity,
	}, nil
}	
func (r *InventoryRepository)ModifyStock(product_id string,newQuantity int32)(*pb.ModifyStockResponse,error){
	tx,err:=r.db.Begin()
	if err!=nil{
		return nil,err
	}
	defer tx.Rollback()
	//var newQuantity int32
	_,err=tx.Exec(`UPDATE inventory SET quantity =$1 WHERE product_id=$2`,newQuantity,product_id)
	if err!=nil{
		return nil,err
	}
	if err := tx.Commit(); err != nil {
        return nil, err
    }
	return &pb.ModifyStockResponse{
		Message: "Stock Updated successfully",
	},nil
}