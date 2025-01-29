#!/bin/bash


# 定义变量
PROTOC_VERSION="3.19.4"
ARCH="linux-x86_64"
ZIP_FILE="protoc-${PROTOC_VERSION}-${ARCH}.zip"
DOWNLOAD_URL="https://github.com/protocolbuffers/protobuf/releases/download/v${PROTOC_VERSION}/${ZIP_FILE}"
EXTRACT_DIR="protoc3"

#下载cwgo与Kitex
go install github.com/cloudwego/cwgo@latest
go install github.com/cloudwego/kitex/tool/cmd/kitex@latest


# 下载 protoc
echo "正在下载 protoc ${PROTOC_VERSION}..."
wget -O "${ZIP_FILE}" "${DOWNLOAD_URL}"

# 检查下载是否成功
if [ $? -ne 0 ]; then
    echo "下载失败，请检查网络连接或版本号。"
    exit 1
fi

# 解压文件到临时目录
echo "解压文件..."
unzip "${ZIP_FILE}" -d "${EXTRACT_DIR}"

# 检查解压是否成功
if [ $? -ne 0 ]; then
    echo "解压失败，请检查压缩包完整性。"
    rm -f "${ZIP_FILE}"
    exit 1
fi

# 将 protoc 移动到 /usr/local/bin
echo "安装 protoc 到 /usr/local/bin..."
cp "${EXTRACT_DIR}/bin/protoc" /usr/local/bin/

# 将 include 文件夹复制到 /usr/local/include
echo "安装 protoc 包含文件到 /usr/local/include..."
mkdir -p /usr/local/include/google
cp -r "${EXTRACT_DIR}/include/google" /usr/local/include/

# 清理临时文件
echo "清理临时文件..."
rm -rf "${ZIP_FILE}" "${EXTRACT_DIR}"

# 验证安装
echo "验证安装..."
protoc --version

# 提示完成
echo "protoc 安装完成。"