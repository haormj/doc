## python

### config

- pip config --editor vi edit
    ```
    [global]
    index-url = https://mirrors.aliyun.com/pypi/simple/ 
    ```

### FAQ

#### conda/miniconda/anaconda 区别
- [[https://github.com/conda/conda][Conda]] 是一个跨平台，语言无关的二进制包管理器。
- [[https://www.anaconda.com/distribution/][Anaconda]] Anaconda 是一个用于科学计算的 Python 发行版，支持 Linux, Mac, Windows, 包含了众多流行的科学计算、数据分析的 Python 包。
- [[https://conda.io/en/latest/miniconda.html][Miniconda]] 是一个 Anaconda 的轻量级替代，默认只包含了 python 和 conda，但是可以通过 pip 和 conda 来安装所需要的包。

#### jupyter notebook选择conda环境

```
conda install nb_conda_kernels
```
