## linux proc 伪文件系统

### 介绍

proc是一个伪文件系统,提供了系统内核的数据结构,一般被挂载到/proc目录下,它下面文件是只读的,也有一些文件是可以被修改的.

### 文件和目录

#### /proc/[pid]/fd/

这个子目录下包含了该进程打开的所有文件,名称为文件描述符,通过符号链接指向实际文件.`0`标准输入,`1`标准输出,`2`标准错误.

> 文件描述符在形式上是一个非负整数。实际上，它是一个索引值，指向[内核](https://zh.wikipedia.org/wiki/%E5%86%85%E6%A0%B8)为每一个[进程](https://zh.wikipedia.org/wiki/%E8%BF%9B%E7%A8%8B)所维护的该进程打开文件的记录表。当程序打开一个现有文件或者创建一个新文件时，内核向进程返回一个文件描述符。
>
> **inode**是指在许多“类[Unix](https://zh.wikipedia.org/wiki/Unix)[文件系统](https://zh.wikipedia.org/wiki/%E6%96%87%E4%BB%B6%E7%B3%BB%E7%BB%9F)”中的一种[数据结构](https://zh.wikipedia.org/wiki/%E6%95%B0%E6%8D%AE%E7%BB%93%E6%9E%84)。每个inode保存了文件系统中的一个**文件系统对象**（包括[文件](https://zh.wikipedia.org/wiki/%E8%AE%A1%E7%AE%97%E6%9C%BA%E6%96%87%E4%BB%B6)、[目录](https://zh.wikipedia.org/wiki/%E7%9B%AE%E5%BD%95_(%E6%96%87%E4%BB%B6%E7%B3%BB%E7%BB%9F))、[设备文件](https://zh.wikipedia.org/wiki/%E8%AE%BE%E5%A4%87%E6%96%87%E4%BB%B6)、[socket](https://zh.wikipedia.org/wiki/Unix%E5%9F%9F%E5%A5%97%E6%8E%A5%E5%AD%97)、[管道](https://zh.wikipedia.org/wiki/%E7%AE%A1%E9%81%93_(Unix)), 等等）的元信息数据，但不包括数据内容或者文件名[[1\]](https://zh.wikipedia.org/zh-hans/Inode#cite_note-1)。

如果文件描述符为`pipes`和`sockets`,符号链接的内容为`type:[inode]`,例如:`socket:[2248868]`

如果文件描述符没有对应的`inode`,符号链接的内容为`anon_inode:<file-type>`,例如:`anon_inode:[eventpoll]`



### 参考

1. [文件描述符](https://zh.wikipedia.org/wiki/%E6%96%87%E4%BB%B6%E6%8F%8F%E8%BF%B0%E7%AC%A6)
2. `man proc`
3. [inode](https://zh.wikipedia.org/zh-hans/Inode)

