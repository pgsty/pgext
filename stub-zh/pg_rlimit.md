## 用法

来源：

- [上游 README](https://github.com/harukat/pg_rlimit/blob/51aec672a347f92fe3e08ed2be75c8ee6b3427e2/README.md)
- [扩展 SQL](https://github.com/harukat/pg_rlimit/blob/51aec672a347f92fe3e08ed2be75c8ee6b3427e2/pg_rlimit--1.0.sql)
- [资源限制实现](https://github.com/harukat/pg_rlimit/blob/51aec672a347f92fe3e08ed2be75c8ee6b3427e2/pg_rlimit.c)

`pg_rlimit` 版本 `1.0` 向每个 PostgreSQL 后端暴露 POSIX 地址空间软限制。它可以限制会话的虚拟地址空间，但不是单条查询的内存配额。

### 核心流程

预加载该库以启用启动钩子和 `pg_rlimit.v` GUC，然后重启 PostgreSQL。

```conf
shared_preload_libraries = 'pg_rlimit'
pg_rlimit.v = '1GB'
```

```sql
CREATE EXTENSION pg_rlimit;
SELECT pg_getrlimit('v');
SELECT pg_setrlimit('v', 1073741824);
SET pg_rlimit.v = '768MB';
```

唯一支持的资源选择符是 `v`，对应 `RLIMIT_AS`。`pg_getrlimit` 以字节返回当前软限制，其中 `-1` 表示无限制。`pg_setrlimit` 设置当前后端的软限制并返回实际值。后端不能把软限制提高到操作系统硬限制以上。

### 运行注意事项

该 GUC 以千字节计量，并接受 PostgreSQL 内存单位。客户端认证钩子会在会话启动时应用配置值；后续 `SET` 修改当前后端。

地址空间包括映射区域与 PostgreSQL 共享内存，并不只是私有分配或 `work_mem`。过低的限制可能让无关操作出现内存不足。应在目标操作系统上测试认证、连接池、并行工作进程、大型映射和错误恢复。上游 SQL 函数没有声明为仅限超级用户，因此若普通会话不应修改限制，需要管理 EXECUTE 权限。
