### awk

输出满足条件的行

```shell
awk '{if ($1 > 10) print $0}' filename
# input
# 1 hello
# 2 world
# 100 who are you
# output
# 100 who are you
```

