#!/bin/bash
# -------------------------------------
# 服务编译重启脚本
#
#
# @author lg1024
# @date 2018年09月25日11:48:25
# -------------------------------------

filename=`pwd`"/api-gin-web"
pidFile=`pwd`"/api-gin-web.pid"

# build
echo -n "正在编译 ... "
go clean
go build -a

# stop
echo -n "正在关闭 ... "
pid=$(ps x | grep $filename | grep -v grep | awk '{print $1}')
echo $pid
kill -9 $pid

sleep 1

# start
echo -n "正在启动 ... "

if [ $ENV_GO == "dev" ]
then
  nohup $filename &
else
  nohup $filename > /dev/null 2>&1 &
fi
echo $! > $pidFile
echo "成功, 进程ID:" $(cat $pidFile)
