#!/usr/bin/env bash
RUN_NAME="order"
mkdir -p output/bin output/conf
cp script/* output/
cp -r conf/* output/conf
cp .env output/.env
chmod +x output/bootstrap.sh
go build -o output/bin/${RUN_NAME}
