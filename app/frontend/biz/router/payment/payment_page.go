// Code generated by hertz generator. DO NOT EDIT.

package payment

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	payment "github.com/douyin-shop/douyin-shop/app/frontend/biz/handler/payment"
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
		_payment := root.Group("/payment", _paymentMw()...)
		_payment.GET("/callback", append(_paymentcallbackMw(), payment.PaymentCallback)...)
	}
}
