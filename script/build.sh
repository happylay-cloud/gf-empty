#!/bin/sh
##########################################################################
#Author :       happylay 安徽理工大学
#Created Time : 2021-01-11 15:33
#Environment :  darwin
##########################################################################

# 打包版本号（修改）
VERSION=1.0.0

# 应用名称（修改）
APP_NAME=app

cd ..

# 打包swagger接口文件（根据选择注释）
gfctl swagger --pack -y

# 打包静态资源文件及sqlite3数据库文件（根据选择注释）
gfctl pack public,template,db packed/data.go -y

cd ./script

# linux_amd64环境
gfctl build ../main.go --cgo --CC=x86_64-linux-musl-gcc --CGO_LDFLAGS="-static" --name $APP_NAME --arch amd64 --system linux --version $VERSION -p ../bin

# windows_amd64环境
gfctl build ../main.go --cgo --CC=x86_64-w64-mingw32-gcc --name $APP_NAME --arch amd64 --system windows --version $VERSION -p ../bin

# mac_amd64环境
gfctl build ../main.go --cgo --name $APP_NAME --arch amd64 --system darwin --version $VERSION -p ../bin

# 压缩应用
cd ../bin/$VERSION

tar -zcvf $APP_NAME.$VERSION-darwin-amd64.tar.gz darwin_amd64

tar -zcvf $APP_NAME.$VERSION-linux-amd64.tar.gz linux_amd64

zip -r $APP_NAME.$VERSION-windows-amd64.zip windows_amd64
