#!/bin/bash

cd app

# 遍历所有服务目录
for service in */; do
    service=${service%/}  # 移除路径末尾的斜杠

    echo "Generating code for service: $service"

    # 进入服务目录
    cd "$service" || continue

    # 执行 cwgo 命令
    cwgo server \
        --type RPC \
        --idl "${service}.proto" \
        --server_name "$service" \
        --module "github.com/douyin-shop/douyin-shop/app/${service}" \
        -I ../../idl

    cd ..
done