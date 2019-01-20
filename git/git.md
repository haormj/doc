### git

合并其他分支的部分提交点

```shell
# https://www.jianshu.com/p/08c3f1804b36
# 首先切换到要合并的分支
# 现在有v1和v2分支,需要将v2的修改合并到v1
# 则首先切换到v1分支,然后使用git cherry-pick合并v2的一些提交点
git checkout v1
git cherry-pick commit_id
```

