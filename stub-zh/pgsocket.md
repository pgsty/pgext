## 用法

来源：

- [官方 README](https://github.com/AbdulYadi/pgsocket/blob/master/README.md)
- [扩展控制文件](https://github.com/AbdulYadi/pgsocket/blob/master/pgsocket.control)
- [1.0 版扩展 SQL](https://github.com/AbdulYadi/pgsocket/blob/master/pgsocket--1.0.sql)
- [1.0 版 C 实现](https://github.com/AbdulYadi/pgsocket/blob/master/pgsocket.c)

pgsocket 从 PostgreSQL 后端发起出站 IPv4 TCP 连接。它可以发送原始字节、交换由 STX 与 ETX 控制字节包围的消息，或使用应用专用的长度前缀图像协议。它不提供 TLS、身份认证、连接池或异步 I/O。

### 核心流程

1.0 版把所有函数明确安装到 public 模式。基础发送函数接收数字 IPv4 地址、端口、以秒计的发送超时和字节载荷。

```sql
CREATE EXTENSION pgsocket;

SELECT public.pgsocketsend(
    '192.0.2.10',
    9090,
    30,
    convert_to('Hello', 'UTF8')
);
```

如果服务端回复以字节 0x02 开始并以字节 0x03 结束，可使用带帧的请求/响应函数：

```sql
SELECT public.pgsocketsendrcvstxetx(
    '192.0.2.10',
    9090,
    30,
    40,
    convert_to('Hello', 'UTF8')
);
```

返回的字节数组不包含帧字节。图像辅助函数实现另一种带本机整数域的二进制协议，只能和为该准确格式构建的对端一起使用。

### 实现边界

虽然 README 给出了主机名示例，但目录中的 C 源码会把地址传给数字 IPv4 转换函数，并不执行 DNS 解析。应使用数字地址，并从数据库主机验证连通性。

每次调用都会新建连接，并阻塞一个 PostgreSQL 后端直到完成或套接字超时。扩展 SQL 把函数声明为稳定函数，但它们实际上执行网络 I/O，因此不能依赖该易变性标签获得安全的优化或重放语义。调用不具备事务性：之后的 SQL 回滚无法撤销已经发送的字节。

### 安全与运维

从 SQL 发起出站套接字会形成服务端请求与数据外泄面。应撤销公共角色的执行权限，通过主机防火墙只允许固定目标，并避免传入用户可控的地址、端口或载荷。协议是明文，也不认证对端。

控制文件声称扩展可重定位，但 SQL 硬编码了 public 模式，因此这些函数实际上无法随模式重定位。扩展没有声明预加载或重启要求。源码使用 POSIX 套接字，必须针对准确的 PostgreSQL 与操作系统目标构建和测试。
