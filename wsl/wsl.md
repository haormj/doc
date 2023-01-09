## wsl

### nameserver

- /etc/wsl.conf
    ```
    [network]
    generateResolvConf = false 
    ```
- wsl.exe --shutdown
- /etc/resolv.conf
    ```
    sudo cat > /etc/resolv.conf <<EOF
    # This file was automatically generated by WSL. To stop automatic generation of this file, add the following entry to /etc/wsl.conf:
    # [network]
    # generateResolvConf = false
    nameserver 1.1.1.1
    nameserver 8.8.8.8
    nameserver 8.8.4.4
    # local dns, modify it to yours
    nameserver 192.168.1.1
    EOF
    ```