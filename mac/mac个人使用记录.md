## mac 个人使用记录

### 常见问题

- mac中的command control和之前用的ctrl和alt的区别，以及如何能够继续使用之前的一套按键
- mac中的按键
    - 因为之前一直使用的ubuntu，对于很小一部分按键还是有需求的
    - 打开终端 ctrl + shift + t， ctrl + d 可以直接关闭掉当前终端
- xcrun: error: invalid active developer path (/Library/Developer/CommandLineTools), missing xcrun at: /Library/Developer/CommandLineTools/usr/bin/xcrun
    - 升级macbook之后可能会遇到上述问题，执行下面命令即可
    - xcode-select --install
