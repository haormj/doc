## mac 个人使用记录

### karabiner element

- 这个配置文件生成有依赖 https://github.com/haormj/code-generator
- 个人还是习惯双control，双alt，双win/command，所以将cherry xaga 右ctrl左边键修改为right command按键
- 本文档 karabiner element下面存放了一些自定义的快捷键
    - 使用的时候，将这些json文件直接复制到 ~/.config/karabiner/assets/complex_modifications 下，然后 karabiner 界面中点击complex modifications中enable对应的功能即可
    - win/command + j/k 来将桌面进行左右切换
    - ctrl + alt + t 打开终端
    - ctrl + d 关闭终端
    - alt + tab 切换不同应用
- karabiner在修改按键的时是支持限制为特定的application，也就是可以通过condition来限制
    - 为了方便查看 bundle_identifiers ，可以通过 karabiner eventViewer来进行

### 常见问题

- mac中的command control和之前用的ctrl和alt的区别，以及如何能够继续使用之前的一套按键
- mac中的按键
    - 因为之前一直使用的ubuntu，对于很小一部分按键还是有需求的
    - 打开终端 ctrl + shift + t， ctrl + d 可以直接关闭掉当前终端
- xcrun: error: invalid active developer path (/Library/Developer/CommandLineTools), missing xcrun at: /Library/Developer/CommandLineTools/usr/bin/xcrun
    - 升级macbook之后可能会遇到上述问题，执行下面命令即可
    - xcode-select --install
- 还是需要解决这个window和mac按键的问题，还有就是编辑器方面的
    - 首先mac的触摸板肯定是好的，感觉比其他笔记本好
    - 也不是它的按键不好，主要是习惯了window，去习惯例外一个就比较麻烦
    - 键盘重新买一个之前的青轴太响了，办公环境不适合，还是要买一个cherry 87键 红轴
