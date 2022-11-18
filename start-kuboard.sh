#!/bin/bash
#
# Start deploy Kuboard v3 to manage multi Kubernetes clusters
# according to https://kuboard.cn/install/v3/install.html
# Created by lhua <hualongfeiyyy@163.com> on 2022-11-15.

KUBOARD_IP=192.168.110.80
KUBOARD_PORT=80
KUBOARD_AGENT_SERVER_TCP_PORT=10081
KUBOARD_VOLUME=/home/godev/backup/kuboard-data

echo "---> Start Kuboard Dashboard..."
sudo docker run -d \
  --restart=unless-stopped \
  --name=kuboard \
  -p ${KUBOARD_PORT}:${KUBOARD_PORT}/tcp -p ${KUBOARD_AGENT_SERVER_TCP_PORT}:${KUBOARD_AGENT_SERVER_TCP_PORT}/tcp \
  -e KUBOARD_ENDPOINT="http://${KUBOARD_IP}:${KUBOARD_PORT}" \
  -e KUBOARD_AGENT_SERVER_TCP_PORT="${KUBOARD_AGENT_SERVER_TCP_PORT}" \
  -v ${KUBOARD_VOLUME}:/data \
  eipwork/kuboard:v3
[[ $? -eq 0 ]] && echo "---> Successful, next access http://${KUBOARD_IP}:${KUBOARD_PORT}"    
