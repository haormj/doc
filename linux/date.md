## Date

将当前时间拷贝到剪切板

```shell
date "+%Y-%m-%d %H:%M:%S" | xsel -b
```

将特定日期转为时间戳

```shell
date --date "2023-06-20 10:28:00" +%s
```

