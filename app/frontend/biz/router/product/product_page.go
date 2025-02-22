// Code generated by hertz generator. DO NOT EDIT.

package product

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	product "github.com/douyin-shop/douyin-shop/app/frontend/biz/handler/product"
)

/*
 This file will register all the routes of the services in the master idl.
 And it will update automatically when you use the "update" command for the idl.
 So don't modify the contents of the file, or your code will be deleted when it is updated.
*/

// Register register routes based on the IDL 'api.${HTTP Method}' annotation.
func Register(r *server.Hertz) {

	root := r.Group("/", rootMw()...)
	{
		_product := root.Group("/product", _productMw()...)
		_product.POST("/add", append(_addproductMw(), product.AddProduct)...)
		_product.POST("/delete", append(_deleteproductMw(), product.DeleteProduct)...)
		_product.POST("/get", append(_getproductMw(), product.GetProduct)...)
		_product.POST("/list", append(_listproductsMw(), product.ListProducts)...)
		_product.POST("/search", append(_searchproductsMw(), product.SearchProducts)...)
		_product.POST("/update", append(_updateproductMw(), product.UpdateProduct)...)
	}
}
