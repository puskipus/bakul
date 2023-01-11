package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/hanifhahn/bakul/controllers/produkcontroller"
)

func ProdukRoutes(r *gin.RouterGroup) {
	r.POST("/produk", produkcontroller.TambahProduk)
	r.GET("/produk", produkcontroller.LihatProduk)
	r.PATCH("/produk/:id", produkcontroller.UpdateProduk)
	r.DELETE("/produk/:id", produkcontroller.HapusProduk)

	r.PATCH("/archiveProduk/:id", produkcontroller.ArchiveProduk)
	r.PATCH("/restoreProduk/:id", produkcontroller.RestoreProduk)
}
