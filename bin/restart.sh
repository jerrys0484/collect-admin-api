#!/bin/bash

# 如果进程存在
if pgrep -f "collect-admin-api" > /dev/null; then
  echo "强制杀死进程..."
  pkill -f "collect-admin-api"
fi

# 启动接口服务
nohup ./bin/collect-admin-api --gf.gcfg.file="./bin/collect-admin-api.yaml" > /dev/null 2>&1 &

echo "重启完成"
