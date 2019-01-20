## kubernates

### 安装kubernates

安装docker,并成功启动,可以直接从软件源安装

按照kubernetes官方提供的方案安装

```shell
CNI_VERSION="v0.6.0"
mkdir -p /opt/cni/bin
curl -L "https://github.com/containernetworking/plugins/releases/download/${CNI_VERSION}/cni-plugins-amd64-${CNI_VERSION}.tgz" | tar -C /opt/cni/bin -xz

CRICTL_VERSION="v1.11.1"
mkdir -p /opt/bin
curl -L "https://github.com/kubernetes-incubator/cri-tools/releases/download/${CRICTL_VERSION}/crictl-${CRICTL_VERSION}-linux-amd64.tar.gz" | tar -C /opt/bin -xz

#RELEASE="$(curl -sSL https://dl.k8s.io/release/stable.txt)"
RELEASE="v1.11.3"
mkdir -p /opt/bin
cd /opt/bin
curl -L --remote-name-all https://storage.googleapis.com/kubernetes-release/release/${RELEASE}/bin/linux/amd64/{kubeadm,kubelet,kubectl}
chmod +x {kubeadm,kubelet,kubectl}

curl -sSL "https://raw.githubusercontent.com/kubernetes/kubernetes/${RELEASE}/build/debs/kubelet.service" | sed "s:/usr/bin:/opt/bin:g" > /etc/systemd/system/kubelet.service
mkdir -p /etc/systemd/system/kubelet.service.d
curl -sSL "https://raw.githubusercontent.com/kubernetes/kubernetes/${RELEASE}/build/debs/10-kubeadm.conf" | sed "s:/usr/bin:/opt/bin:g" > /etc/systemd/system/kubelet.service.d/10-kubeadm.conf

systemctl enable kubelet && systemctl start kubelet
```

kubeadm.yaml

```yaml
apiVersion: kubeadm.k8s.io/v1alpha1
kind: MasterConfiguration
controllerManagerExtraArgs:
  horizontal-pod-autoscaler-use-rest-clients: "true"
  horizontal-pod-autoscaler-sync-period: "10s"
  node-monitor-grace-period: "10s"
apiServerExtraArgs:
  runtime-config: "api/all=true"
kubernetesVersion: "stable-1.11"
```

部署master节点

```shell
kubeadm init --config kubeadm.yaml

mkdir -p $HOME/.kube
sudo cp -i /etc/kubernetes/admin.conf $HOME/.kube/config
sudo chown $(id -u):$(id -g) $HOME/.kube/config

# 部署网络插件weave
kubectl apply -f https://git.io/weave-kube-1.6
```

部署worker节点

```shell
# 同样需要先要按照kubernates官方方式安装kubelet/kubeadm
# 之后执行部署完成master之后打印出的kubeadm join
```

通过 Taint/Toleration 调整 Master 执行 Pod 的策略

```shell
kubectl taint nodes node1 foo=bar:NoSchedule
# Taints:             node-role.kubernetes.io/master:NoSchedule
kubectl describe node
# 我们在node-role.kubernetes.io/master这个键后面加上了一个短横线“-”，
# 这个格式就意味着移除所有以node-role.kubernetes.io/master为键的 Taint。
kubectl taint nodes --all node-role.kubernetes.io/master-
```

部署dashboard可视化插件

```shell
kubectl apply -f https://raw.githubusercontent.com/kubernetes/dashboard/master/src/deploy/recommended/kubernetes-dashboard.yaml
```

部署容器存储插件

```shell
kubectl apply -f https://raw.githubusercontent.com/rook/rook/master/cluster/examples/kubernetes/ceph/operator.yaml

kubectl apply -f https://raw.githubusercontent.com/rook/rook/master/cluster/examples/kubernetes/ceph/cluster.yaml

kubectl get pods -n rook-ceph-system
kubectl get pods -n rook-ceph
```

### 通过kubernetes部署程序

nginx-deployment.yaml

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: nginx-deployment
spec:
  selector:
    matchLabels:
      app: nginx
  replicas: 2
  template:
    metadata:
      labels:
        app: nginx
    spec:
      containers:
      - name: nginx
        image: nginx:1.8
        ports:
        - containerPort: 80
        volumeMounts:
        - mountPath: "/usr/share/nginx/html"
          name: nginx-vol
      volumes:
      - name: nginx-vol
        emptyDir: {}
```

基本操作

```shell
kubectl create -f nginx-deployment.yaml
kubectl get pods -l app=nginx
kubectl describe pod xxx
kubectl replace -f nginx-deployment.yaml
# 推荐之后就直接使用 apply/delete
kubectl apply -f nginx-deployment.yaml
kubectl delete -f nginx-deployment.yaml
```

### 参考文档

[Bootstrapping Clusters with kubeadm](https://kubernetes.io/docs/setup/independent/install-kubeadm/)

[kubeadm](https://kubernetes.io/docs/reference/setup-tools/kubeadm/kubeadm/)





