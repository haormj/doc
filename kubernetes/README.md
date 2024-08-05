## kubernetes

### 概述

本文档主要记录个人对 kubernetes 的一些理解

### 技术点

- 容器
  - 容器的核心由 cgroup、namespace、rootfs 构成
  - cgroup 全称是 linux control group，主要是用于控制进程对资源的使用，包括 cpu，内存，磁盘，网络等，具体可以在 /sys/fs/cgroup 下面查看
  - 对于基于 systemd 做用户进程管理的系统，可以在上述文件夹下面看到 system.slice user.slice 文件夹，下面存放着 systemd 管理的进程的 cgroup 文件
  - namespace 是 linux 的隔离机制，目前支持 pid namespace, mount namespace, network namespace, ipc namespace, cgroup namespace
  - 有了 namespace 之后，不同进程之间就有了天然的隔离
  - rootfs，其实就是将系统目录，不包括内核的进行打包，然后在进程有了 cgroup namespace 这些约束之后，基于 chroot 类似的系统调用，来改变 rootfs
  - docker 镜像，是由很多层组成的，目前分为三个类别 readonly whiteout, init, readwrite， 其中 whiteout 主要就是支持通过在读写层创建空白文件来实现删除效果
  - init 层主要是不会被 docker commit 提供，里面会包含 /etc/hostname /etc/resolv.conf 文件
  - 最后将不同层之间联合挂载到一个目录，这个目录就是进程运行的rootfs
  - 而将宿主机目录挂载到容器中，主要是基于绑定挂载来实现
- k8s 架构
  - master：api-server，controller-manager， schedule
  - node：kubelet
- k8s 基本概念
  - pod: 是k8s最基本的调度单位，一个pod中可以由多个容器，多个容器会共享命名空间，比如 network namespace，pod中有一个 infra 容器
    - 从根本上来讲，其实 pod 就是一组共享命名空间的进程
