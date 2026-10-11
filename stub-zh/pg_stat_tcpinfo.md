## 用法

来源：

- [README.md](https://github.com/jakubwartakEDB/pg_stat_tcpinfo/blob/46be6d5f55330a6f890a093963110bf477c90fdc/README.md)
- [pg_stat_tcpinfo--1.0.sql](https://github.com/jakubwartakEDB/pg_stat_tcpinfo/blob/46be6d5f55330a6f890a093963110bf477c90fdc/pg_stat_tcpinfo--1.0.sql)
- [LICENSE](https://github.com/jakubwartakEDB/pg_stat_tcpinfo/blob/46be6d5f55330a6f890a093963110bf477c90fdc/LICENSE)
- [pg_stat_tcpinfo.control](https://github.com/jakubwartakEDB/pg_stat_tcpinfo/blob/46be6d5f55330a6f890a093963110bf477c90fdc/pg_stat_tcpinfo.control)

`pg_stat_tcpinfo` 展示服务端操作系统用户拥有的 TCP 连接，包括客户端、复制和外部服务器连接。1.0 版要求 Linux 4.18 或更新版本，并能访问 `/proc` 与 `NETLINK_INET_DIAG`。

### 核心用法

```sql
CREATE EXTENSION pg_stat_tcpinfo;
SELECT pid, dst_addr, dst_port, state, recvq, sendq, tcpinfo
FROM pg_stat_tcpinfo;
```

### 运行边界

视图调用 `pg_stat_get_tcpinfo()`，返回进程与用户 ID、本地和远端地址及端口、状态、队列大小，以及 `tcpinfo` JSON 对象。原生扩展须由超级用户安装。访问权限授予 `pg_read_all_stats`，`pg_monitor` 通过继承获得权限；额外授权须考虑会暴露服务端操作系统用户的全部套接字。TIME-WAIT 套接字没有属主，因此不显示。上游未公布 PostgreSQL 主版本支持矩阵，也未要求预加载。
