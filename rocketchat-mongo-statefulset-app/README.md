## ☸️ Kubernetes 中部署 Rocket.Chat 与 MongoDB 数据库实时交流平台

### 部署环境说明：

- Kubernetes 版本：`v1.22.1`

- Rocket.Chat 容器镜像版本：`docker.io/rocketchat/rocket.chat:3.18.7`

- MongoDB 容器镜像版本：`docker.io/library/mongo:4.0`

- Kubernetes NFS-Client Privisioner 容器镜像版本：`quay.io/external_storage/nfs-client-provisioner:latest`

> 👉 若无法拉取指定的容器镜像，以上镜像均可从 quay.io/user/alberthua 中拉取下载。

### 部署方式及步骤：

- 💥 该应用后端的 MongoDB 集群使用 NFS 作为动态 PV 的提供者，需提前配置 NFS 服务器节点用于提供 PV。

- MongoDB 集群使用 `StatefulSet` 部署，而该资源需使用 `StorageClass` 实现卷声明模板（`volumeClaimTemplates`）。

- Kubernetes 中未集成 NFS 类型的内部调配者（`internal privisioner`），因此需使用 `nfs-client-provisioner` 将外部 NFS 调配至集群以支持 PV 动态分配。

- nfs-client-provisioner 在集群中的部署可参考该 [链接](https://github.com/Alberthua-Perl/go-kubernetes-learn-path/tree/hotfixes/nfs-provisioned-storageclass)。

> 注意：nfs-client-provisioner pod 与应用 pod 部署于同一命名空间（namespace）中。

- 因此，该应用的部署方式如下所示：

  ```bash
  $ kubectl create namespace rocketchat-mongodb-app
  $ kubectl apply -f \
    00-nfs-provisioned-rbac.yml 01-nfs-provisioned-deployment.yml 02-nfs-provisioned-class.yml -n rocketchat-mongodb-app
  $ kubectl apply -f 03-mongodb-internal-headless-svc.yml -n rocketchat-mongodb-app
  $ kubectl apply -f 04-mongodb-statefulset.yml -n rocketchat-mongodb-app
  # 该资源创建完成后并未实现 MongoDB 的 ReplicaSet 模式集群，需登录至其中的一个节点实现集群的初始化及 mongo 节点的添加。
  ```
  ```bash
  $ kubectl exec -it rocketmongo-0 -n rocketchat-mongodb-app -- mongo
    # 进入 mongo 节点进行集群的初始化与配置
    ...
    > rs.initiate()  # 初始化集群
    > var config = rs.conf()
    > config.members[0].host="rocketmongo-0.mongodb-internal:27017"  
      # 通过 headless service 指向 mongo 节点，将该节点配置为 primary 节点。
    > rs.reconfig(config)  # 刷新集群配置
    > rs.add("rocketmongo-1.mongodb-internal:27017")
    > rs.add("rocketmongo-2.mongodb-internal:27017")  # 添加额外的 mongo 节点
    > rs.status()  # 查看集群的状态
    > rs.isMaster()  # 确认当前 mongo 节点是否为 primary 节点 
    > exit  # 退出 MongoDB Shell
    ...
  ```
  ```bash
  $ kubectl apply -f 05-rockerchat-deployment.yml -n rocketchat-mongodb-app
  # 部署前端 Rocket.Chat 应用
  ```

  🤘 如下所示，刷新 Rocket.Chat pod 日志可确认其与 MongoDB 集群成功连接：

  ![](https://github.com/Alberthua-Perl/go-kubernetes-learn-path/blob/hotfixes/rocketchat-mongo-statefulset-app/images/rocketchat-mongo-connect-successfully.png)

### 确认应用资源与登录认证：

- 该应用所涉及的资源对象如下所示：

  ![](https://github.com/Alberthua-Perl/go-kubernetes-learn-path/blob/hotfixes/rocketchat-mongo-statefulset-app/images/rocketchat-mongo-app-resources.png)

- 可通过 Rocket.Chat pod 日志中的 URL 链接登录应用并注册账户使用。

  ![](https://github.com/Alberthua-Perl/go-kubernetes-learn-path/blob/hotfixes/rocketchat-mongo-statefulset-app/images/rocketchat-login.png)
  
### 参考链接：

- [Running MongoDB on Kubernetes with StatefulSets](https://kubernetes.io/blog/2017/01/running-mongodb-on-kubernetes-with-statefulsets/)

- [Running Rocket.Chat and MongoDB on Kubernetes with StatefulSets](https://dev.to/jmarhee/running-rocketchat-and-mongodb-on-kubernetes-with-statefulsets-431o)

- [Deploy Rocket chat server using Kubernetes](https://ajorloo.medium.com/deploy-rocket-chat-server-using-kubernetes-2d6c4228853)
  
