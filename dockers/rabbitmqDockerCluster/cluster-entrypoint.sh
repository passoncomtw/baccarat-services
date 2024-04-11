#!/bin/bash
# 在背景啟動 rabbitmq service
/usr/local/bin/docker-entrypoint.sh rabbitmq-server -detached
# 等待 rabbitmq 啟動完畢
sleep 5s
# 停止 rabbitmq application
rabbitmqctl stop_app
# 加入 rabbimtq cluster
rabbitmqctl join_cluster rabbit@rabbitmq-1
# 啟動 rabbitmq application 來 sync 資料
rabbitmqctl start_app
# 停止 rabbitmq service
rabbitmqctl stop
# 等待 rabbitmq service 完全停止
sleep 2s
# 將 rabbitmq service 啟動在前景
rabbitmq-server