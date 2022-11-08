## Rocket.Chat 与 MongoDB 数据库集群部署 Kubernetes 实时交流平台

### 部署环境说明：

- Kubernetes 版本：`v1.22.1`

- Rocket.Chat 容器镜像版本：`docker.io/rocketchat/rocket.chat:3.18.7`

- MongoDB 容器镜像版本：`docker.io/library/mongo:4.0`

- Kubernetes NFS-Client Privisioner 容器镜像版本：`quay.io/external_storage/nfs-client-provisioner:latest`

> 👉 若无法拉取指定的容器镜像，以上镜像均可从 quay.io/user/alberthua 中拉取下载。

### 部署方式及步骤：

- 该应用后端的 MongoDB 集群使用 NFS 作为动态 PV 的提供者，需提前配置 NFS 服务器节点用于提供 PV。

- MongoDB 集群使用 `StatefulSet` 部署，而该资源需使用 `StorageClass` 资源实现卷声明模板（`volumeClaimTemplates`）。

- Kubernetes 中未集成 NFS 类型的内部调配者（`internal provisioner`），因此需要外部的 `nfs-client-provisioner` 作为对外部 NFS 向集群的引入调配以支持动态 PV 提供。

- nfs-client-provisioner 在集群中的部署可参考该 [链接](https://github.com/Alberthua-Perl/go-kubernetes-learn-path/tree/hotfixes/nfs-provisioned-storageclass)。

> 注意：nfs-client-provisioner pod 与应用 pod 部署于同一命名空间（namespace）中。

- 因此，该应用的部署方式如下所示：

  ```bash
  $ kubectl create namespace rocketchat-mongodb-app
  $ kubectl apply -f 00-nfs-provisioned-rbac.yml 01-nfs-provisioned-deployment.yml 02-nfs-provisioned-class.yml -n rocketchat-mongodb-app
  $ kubectl apply -f 03-mongodb-internal-headless-svc.yml -n rocketchat-mongodb-app
  $ kubectl apply -f 04-mongodb-statefulset.yml -n rocketchat-mongodb-app
  # 该资源创建完成后并未实现 MongoDB 的 ReplicaSet 模式集群，需登录至其中的一个节点实现集群的初始化及 mongo 节点的添加。
  ```
  
  
  
  
  
  
  
