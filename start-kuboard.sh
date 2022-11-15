#!/bin/bash
#
# Start deploy Kuboard v3 to manage multi Kubernetes clusters
# according to https://kuboard.cn/install/v3/install.html
# Created by lhua <hualongfeiyyy@163.com> on 2022-11-15.

sudo docker run -d \
  --restart=unless-stopped \
  --name=kuboard -p 80:80/tcp -p 10081:10081/tcp \
  -e KUBOARD_ENDPOINT="http://192.168.110.80:80" \
  -e KUBOARD_AGENT_SERVER_TCP_PORT="10081" \
  -v /home/godev/backup/kuboard-data:/data \
  eipwork/kuboard:v3
