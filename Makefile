# 用于安装和卸载smartapi服务

# 服务信息
SERVICE = smartapi
# 相关变量
#VERSION=`git tag | tail -1`
VERSION=`git for-each-ref --sort=creatordate --format='%(refname)' | grep refs/tags | sed -e 's:refs/tags/::g' | tail -1`
BUILD=`date +"%Y-%m-%d %H:%M:%S"`
COMMITSHA1=`git rev-parse HEAD`
LDFLAGS=-ldflags "-s -w -X main.VERSION=${VERSION} -X 'main.BUILD=${BUILD}' -X main.COMMITSHA1=${COMMITSHA1}"

build:
	##测试构建环境参数


install: deploy_env
	##安装服务相关配置


uninstall: deploy_env
	##停止相关服务进程


linux:
	go build $(LDFLAGS)

arm:
	CGO_ENABLED=0 GOOS=linux GOARCH=arm GOARM=7 \
	go build -v -a $(LDFLAGS)
