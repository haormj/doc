## ssh FAQ

- unable to negotiate with x.x.x.x port 22: no matching host key type found. Their offer: ssh-rsa
  - ref [link](https://stackoverflow.com/questions/69875520/unable-to-negotiate-with-40-74-28-9-port-22-no-matching-host-key-type-found-th)
  - ~/.ssh/config
      ```
      Host ssh.haormj.xyz
      User git
      PubkeyAcceptedAlgorithms +ssh-rsa
      HostkeyAlgorithms +ssh-rsa
      ```
