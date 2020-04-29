# api-gin-web
[![Go Report Card]()]()
[![build status]()]()
[![coverage report]()]()


## 基本配置

- go v1.12 以上
- mysql v5.7.17 以上
- redis v3.2.9 以上

## 系统环境变量设置

- `export GOSUMDB=off`  用来配置验证包有效性的变量，默认是GOSUMDB=sum.golang.org，这个网址被墙了，所以一般关闭掉这个变量
- `export GO111MODULE=on` 设置go mod拉取依赖包
- `export ENV_GO=dev` 访问内网配置变量`conf/config-dev.yaml`

## 后台接口

- 启动：`./restart.sh`

## 环境配置

- 测试环境 
    - api-gin-web
- 预发布环境
    - api-gin-web
- 正式环境
    - api-gin-web
