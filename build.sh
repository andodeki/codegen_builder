VERSION=`git symbolic-ref -q --short HEAD || git describe --tags --exact-match`
# OOS=linux go build -ldflags="-s -w -X main.buildVersion=$VERSION -X main.buildTime=`date -Is`"
# usage: go build [-o output] [build flags] [packages]
OOS=linux go build -o buildProj/codegen_builder -ldflags="-s -w -X main.BuildVersion=$VERSION -X main.buildTime=`date -Is`"
# cd ..
upx -qqq buildProj/codegen_builder
# cd ..
pwd
./buildProj/codegen_builder
# ./codegen_builder -c ./config.yaml

# Run the program
# chmod +x build.sh
# ./build.sh

apk --update-cache --repository http://dl-3.alpinelinux.org/alpine/edge/testing/ add android-tools