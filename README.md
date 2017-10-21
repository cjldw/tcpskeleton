### Golang TCP 服务框架, 快速搭建TCP服务

代码是fork[gansidui](https://github.com/gansidui)的 gotcp. 根据自己理解做了部分修改.

### 如何编写代码

框架定义好了3个loop, 是要实现下面接口即可

1. 实现 `Callback`接口
2. 实现 `Protocol`接口
3. 实现 `Packet` 接口
4. 启动服务, 完事.

### 安装

使用golang的官方安装凡是

```bash
    $ go get github.com/vvotm/tcpskeleton
```

### 简单的例子

* echo实现[echo](https://github.com/vvotm/tcpskeleton/tree/master/examples/echo)
* telnet实现[telnet](https://github.com/vvotm/tcpskeleton/tree/master/examples/telnet)
* 自定协议实现[diyproto](https://github.com/vvotm/tcpskeleton/tree/master/examples/diyproto)

### 文档

[Doc](http://godoc.org/github.com/vvotm/tcpskeleton)


