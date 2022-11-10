## Example for Kubernetes NFS-Client Privisioner

- YAML files edited and deployed in Kubernetes v1.22.1 cluster.

- According to `{00..04}-{resource-definied-name}.yml` to create object.

- 👉 If use yaml files to create resources on OpenShift cluster, you can get more specified doc on [kubernetes-retired / external-storage](https://github.com/kubernetes-retired/external-storage/tree/master/nfs-client).

> ✨ Note: `kubernetes-retired / external-storage` deprecated to [kubernetes-sigs / nfs-subdir-external-provisioner](https://github.com/kubernetes-sigs/nfs-subdir-external-provisioner)!

- In fortunately `deployment` and `storageclass` are also be supported in Red Hat OpenShift Container Platform v3.9 (OCP v3.9).

- Also Kubernetes NFS-Client Privisioner demo has been tested on OCP v3.9 and it works OK.