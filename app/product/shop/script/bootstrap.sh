#! /usr/bin/env bash
CURDIR=$(cd $(dirname $0); pwd)
echo "$CURDIR/bin/shop"
exec "$CURDIR/bin/shop"
