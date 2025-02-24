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

cd ..
## 下载 IK 分词器插件方便挂载到docker容器中
# 定义变量
TOOLS_DIR="./tools"
IK_ANALYZER_VERSION="8.17.1"
PLUGIN_NAME="elasticsearch-analysis-ik-$IK_ANALYZER_VERSION"
PLUGIN_ZIP="$PLUGIN_NAME.zip"
PLUGIN_DIR="$TOOLS_DIR/$PLUGIN_NAME"

# 创建 tools 目录（如果不存在）
mkdir -p "$TOOLS_DIR"

# 检查插件目录是否存在
if [ ! -d "$PLUGIN_DIR" ]; then
    echo "IK Analyzer plugin not found in $PLUGIN_DIR, starting download..."

    # 下载 IK 分词器插件
    echo "Downloading IK Analyzer plugin..."
    wget "https://release.infinilabs.com/analysis-ik/stable/$PLUGIN_ZIP" -P "$TOOLS_DIR"
    
    # 检查下载是否成功
    if [ $? -ne 0 ]; then
        echo "Failed to download the IK Analyzer plugin."
        exit 1
    fi

    # 解压插件
    echo "Extracting IK Analyzer plugin..."
    unzip "$TOOLS_DIR/$PLUGIN_ZIP" -d "$TOOLS_DIR"

    # 删除压缩包
    rm "$TOOLS_DIR/$PLUGIN_ZIP"

    echo "IK Analyzer plugin installation completed."
else
    echo "IK Analyzer plugin already exists in $PLUGIN_DIR, skipping download and extraction."
fi