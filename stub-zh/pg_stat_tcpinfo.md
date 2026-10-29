## 用法

来源：

- [README](https://github.com/jakubwartakEDB/pg_stat_tcpinfo/blob/v1.1/README.md)
- [安装 SQL](https://github.com/jakubwartakEDB/pg_stat_tcpinfo/blob/v1.1/pg_stat_tcpinfo--1.1.sql)
- [扩展控制文件](https://github.com/jakubwartakEDB/pg_stat_tcpinfo/blob/v1.1/pg_stat_tcpinfo.control)
- [1.1 版本发布](https://github.com/jakubwartakEDB/pg_stat_tcpinfo/releases/tag/v1.1)
- [许可证](https://github.com/jakubwartakEDB/pg_stat_tcpinfo/blob/v1.1/LICENSE)

`pg_stat_tcpinfo` 1.1 通过 SQL 提供 Linux TCP 套接字统计，利用往返延迟、重传次数、拥塞控制和队列大小，帮助排查缓慢的客户端连接、复制链路与外部服务器连接，无需登录主机 Shell。

扩展要求 Linux 4.18 或更新版本，能够访问 `/proc` 和 `NETLINK_INET_DIAG` 接口。须由超级用户安装，无需预加载或重启数据库。

### 核心用法

```sql
CREATE EXTENSION pg_stat_tcpinfo;

SELECT a.pid, a.usename, a.application_name,
       t.dst_addr, t.dst_port, t.state, t.recvq, t.sendq,
       t.tcpinfo->>'rtt' AS rtt_ms,
       t.tcpinfo->>'total_retrans' AS retransmissions,
       t.tcpinfo->>'congestion' AS congestion_control
FROM pg_stat_activity a
JOIN pg_stat_tcpinfo t ON t.pid = a.pid
WHERE t.state = 'ESTABLISHED';
```

### 对象与指标

`pg_stat_tcpinfo` 视图调用 `pg_stat_get_tcpinfo()`。每行包含进程与操作系统用户 ID、本地和远端地址及端口、连接状态、收发队列，以及 `tcpinfo` JSONB 对象。支持 IPv4 和 IPv6 套接字；`rtt` 以毫秒为单位，`total_retrans` 为累计重传计数。

视图包含数据库服务端操作系统用户拥有的全部 TCP 套接字，也可能包含同一账号下其他数据库集群或程序的连接。按后端 PID 关联可将结果缩小到当前集群可见的连接。Unix 域连接不属于 TCP 套接字，不会显示；TIME-WAIT 套接字没有属主，也会被排除。旧内核不支持的指标可能为空。

### 权限

安装脚本撤销函数的公共执行权限，并将函数和视图访问权授予 `pg_read_all_stats`，`pg_monitor` 继承这些权限。普通应用角色默认不能读取指标。额外授权会暴露服务端操作系统用户的全部套接字，并不限于调用者自己的数据库连接。

### 运行边界

每次调用都会扫描进程描述符并查询内核，提供当前状态，不保存历史采样。容器的 PID、网络命名空间和 procfs 访问限制决定可见的套接字与进程。应根据进程和连接数量选择采样间隔。

使用与服务端 PostgreSQL 主版本匹配的软件包。1.1 版本包含 1.0 和 1.1 的新安装 SQL，但未提供两者之间的升级脚本，因此现有 1.0 安装没有上游提供的 `ALTER EXTENSION pg_stat_tcpinfo UPDATE` 升级路径。
