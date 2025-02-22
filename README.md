# douyin-shop


```bash
# 生成用户微服务
cwgo server  --type RPC  --idl user.proto  --server_name user --registry NACOS  --module github.com/douyin-shop/douyin-shop/app/user -I ../../idl 
# 生成授权微服务
cwgo server  --type RPC  --idl auth.proto  --server_name auth --registry NACOS  --module github.com/douyin-shop/douyin-shop/app/auth -I ../../idl
cwgo server  --type RPC  --idl order.proto  --server_name order --registry NACOS  --module github.com/douyin-shop/douyin-shop/app/order -I ../../idl
cwgo server  --type RPC  --idl payment.proto  --server_name payment  --module github.com/douyin-shop/douyin-shop/app/payment -I ../../idl
cwgo server  --type RPC  --idl checkout.proto  --server_name checkout  --module github.com/douyin-shop/douyin-shop/app/checkout -I ../../idl
cwgo server  --type RPC  --idl cart.proto  --server_name cart  --module github.com/douyin-shop/douyin-shop/app/cart -I ../../idl

cwgo client  --type RPC  --idl user.proto  --server_name user --registry NACOS  --module github.com/douyin-shop/douyin-shop/app/user -I ../../idl 
```

```bash
# 生成前端用户微服务
cwgo server  --type HTTP -I ../../idl --idl ../../idl/frontend/user_page.proto  --server_name frontend  --module github.com/douyin-shop/douyin-shop/app/frontend 
# 生成前端购物车微服务
cwgo server  --type HTTP -I ../../idl --idl ../../idl/frontend/cart_page.proto  --server_name frontend  --module github.com/douyin-shop/douyin-shop/app/frontend
# 生成前端结算微服务
cwgo server  --type HTTP -I ../../idl --idl ../../idl/frontend/checkout_page.proto  --server_name frontend  --module github.com/douyin-shop/douyin-shop/app/frontend
# 生成前端商品微服务
cwgo server  --type HTTP -I ../../idl --idl ../../idl/frontend/product_page.proto  --server_name frontend  --module github.com/douyin-shop/douyin-shop/app/frontend
# 生成前端订单微服务
cwgo server  --type HTTP -I ../../idl --idl ../../idl/frontend/order_page.proto  --server_name frontend  --module github.com/douyin-shop/douyin-shop/app/frontend

```