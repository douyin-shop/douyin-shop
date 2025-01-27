# douyin-shop


```bash
# 生成用户微服务
cwgo server  --type RPC  --idl user.proto  --server_name user --registry NACOS  --module github.com/douyin-shop/douyin-shop/app/user -I ../../idl 
# 生成授权微服务
cwgo server  --type RPC  --idl auth.proto  --server_name auth --registry NACOS  --module github.com/douyin-shop/douyin-shop/app/auth -I ../../idl
cwgo client  --type RPC  --idl user.proto  --server_name user --registry NACOS  --module github.com/douyin-shop/douyin-shop/app/user -I ../../idl 
cwgo server  --type HTTP  --idl ../../idl/frontend.proto  --server_name frontend --registry NACOS  --module github.com/douyin-shop/douyin-shop/app/frontend 
```