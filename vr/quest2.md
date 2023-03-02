## quest2

### FAQ

- 调整时区
    - 执行下面命令开启android系统设置，然后再里面找到时间设置，就可以调整时区了
    ```
    adb shell am start -a android.intent.action.VIEW -d com.oculus.tv -e uri com.android.settings/.DevelopmentSettings com.oculus.vrshell/.MainActivity
    ```
