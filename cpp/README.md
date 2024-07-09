## cpp

### 概述

本文档记录一些 c++ 相关常见问题

### 问题

- ERROR: flag 'flagfile' was defined more than once
  - 个人目前在mac上遇到，是因为 conda 与 brew 都存在 gflags 冲突导致，可以通过 conda deactivate 后，重新编译即可
