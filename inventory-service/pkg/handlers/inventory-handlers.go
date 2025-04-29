package handlers

import (
	"context"
	pb "inventory-service/gen/inventory"
	"inventory-service/pkg/services"
)

type InventoryHandlers struct {
    inventoryService *services.InventoryService
    pb.UnimplementedInventoryServiceServer
}

func NewInventoryHandlers(invSvc *services.InventoryService) *InventoryHandlers {
    return &InventoryHandlers{
        inventoryService: invSvc,
    }
}

func (h *InventoryHandlers) AddStock(ctx context.Context, req *pb.AddStockRequest) (*pb.AddStockResponse, error) {
    return h.inventoryService.AddStockService(req.UserId, req.UserRole, req.ProductId, req.Quantity)
}


func (h *InventoryHandlers) ModifyStock(ctx context.Context, req *pb.ModifyStockRequest) (*pb.ModifyStockResponse, error) {
    return h.inventoryService.ModifyStockService(req.UserId, req.UserRole, req.StockId, req.NewQuantity)
}


func (h *InventoryHandlers) CheckStock(ctx context.Context, req *pb.CheckStockRequest) (*pb.CheckStockResponse, error) {
    return h.inventoryService.CheckStockService(req.ProductId)
}


func (h *InventoryHandlers) DecreaseStock(ctx context.Context, req *pb.DecreaseStockRequest) (*pb.DecreaseStockResponse, error) {
    return h.inventoryService.DecreaseStockService(req.ProductId, req.Quantity)
}