package services

import (
	"fmt"
	pb "inventory-service/gen/inventory"
	"inventory-service/pkg/repository"
)
const SellerRole = "SELLER"

type InventoryService struct {
	inventoryRepo *repository.InventoryRepository
}
func NewInvenotryservice(invRepo *repository.InventoryRepository)*InventoryService{
	return &InventoryService{
		inventoryRepo: invRepo,
	}
}
func (s *InventoryService) isAuthorized(user_role string) bool {
	return user_role == SellerRole
}
func(s *InventoryService)CheckStockService(product_id string)(*pb.CheckStockResponse,error){
	if product_id==""{
		return nil,fmt.Errorf("product id must be provided")
	}
	res,err:= s.inventoryRepo.GetStockOfProduct(product_id)
	if err!=nil{
		return nil,fmt.Errorf("failed to get the sotck of product %s :%w ",product_id,err)
	}
	return res,nil
}
func(s *InventoryService)ModifyStockService(user_id string,user_role string,product_id string,quantity int32)(*pb.ModifyStockResponse,error){
	// Input validation
	switch {
	case user_id == "":
		return nil, fmt.Errorf("user id must be provided")
	case user_role == "":
		return nil, fmt.Errorf("user role must be provided")
	case product_id == "":
		return nil, fmt.Errorf("product id must be provided")
	case quantity < 0:
		return nil, fmt.Errorf("quantity must be positive")
	}

	// Authorization check
	if !s.isAuthorized(user_role) {
		return nil, fmt.Errorf("unauthorized: only SELLER role can modify stock, got %s", user_role)
	}

	res, err := s.inventoryRepo.ModifyStock(product_id, quantity)
	if err != nil {
		return nil, fmt.Errorf("failed to modify stock of product %s: %w", product_id, err)
	}
	
	return res, nil
}
func(s *InventoryService)AddStockService(user_id string,user_role string,product_id string,quantity int32)(*pb.AddStockResponse,error){
		// Input validation
		switch {
		case user_id == "":
			return nil, fmt.Errorf("user id must be provided")
		case user_role == "":
			return nil, fmt.Errorf("user role must be provided")
		case product_id == "":
			return nil, fmt.Errorf("product id must be provided")
		case quantity < 0:
			return nil, fmt.Errorf("quantity must be positive")
		}
	
		// Authorization check
		if !s.isAuthorized(user_role) {
			return nil, fmt.Errorf("unauthorized: only SELLER role can modify stock, got %s", user_role)
		}
	
		res, err := s.inventoryRepo.CreateNewStock(product_id, quantity)
		if err != nil {
			return nil, fmt.Errorf("failed to modify stock of product %s: %w", product_id, err)
		}
		
		return res, nil
}
func(s *InventoryService)DecreaseStockService(product_id string,quantity int32)(*pb.DecreaseStockResponse,error){
	switch{
	case product_id=="":
		return nil,fmt.Errorf("product id not provided ")
	case quantity<0:
		return nil,fmt.Errorf("quantity is less than zero")
	}
	res,err:=s.inventoryRepo.DecreaseStockOfProduct(product_id,quantity)
	if err!=nil{
		return nil,fmt.Errorf("error in decreasing stock of product %v : %v",product_id,err)
	}
	return res,nil
}