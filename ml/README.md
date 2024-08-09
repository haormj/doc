## 机器学习

### 概述

本文档记录一些机器学习相关内容

### 笔记

- transfomer
  - 由 encoder 与 decoder 组成
  - encoder
    - 比较核心的内容是注意力机制，对于一个有 5 个 token 的输入 x1, x2, x3, x4, x5
    - 首先会进行 position embedding, 得到最终 x1, x2, x3, x4, x5
    - 然后每个都有对应的 q1, k1, v1, 然后这里对应的权重矩阵分别乘以 q1, k1, v1
    - 然后当前词的 q 会分别与之前以及当前的 k 进行点积，对得到的结果 softmax 得到最后的注意力
    - 然后由之前的 v 分别与注意力相乘得到最后的 v，最后的 v 会进入 mlp 得到最后的结果

### 资料

- https://github.com/poloclub/transformer-explainer
