## 用法

来源：

- [已复核 commit 的 pg_dropbuffers README](https://github.com/rjuju/pg_dropbuffers/blob/8328aed1a7ffc3c1a322fa2decf1be7c63024842/README.md)
- [已复核 commit 的 pg_dropbuffers 0.0.1 安装 SQL](https://github.com/rjuju/pg_dropbuffers/blob/8328aed1a7ffc3c1a322fa2decf1be7c63024842/pg_dropbuffers--0.0.1.sql)
- [已复核 commit 的 pg_dropbuffers C 实现](https://github.com/rjuju/pg_dropbuffers/blob/8328aed1a7ffc3c1a322fa2decf1be7c63024842/pg_dropbuffers.c)

`pg_dropbuffers` 0.0.1 是仅用于测试的缓存清除工具。`pg_drop_current_db_buffers` 无需重启集群即可移除当前数据库的缓冲区。它只能用于已停止并发写入的隔离、可丢弃集群。

```sql
CREATE EXTENSION pg_dropbuffers;

CHECKPOINT;
SELECT pg_drop_current_db_buffers();
```

`pg_drop_system_cache` 通过 sudo 调用下列命令，而 `pg_drop_caches` 会依次调用两个缓存清除函数：

```text
/usr/bin/sudo /sbin/sysctl vm.drop_caches=3
```

### 注意事项

- 绝不能在生产环境使用。实现明确警告：移除时缓冲区仍可能是脏页，因此即使先请求检查点，也可能丢失数据。
- `pg_drop_system_cache` 仅适用于 Linux，并要求 PostgreSQL 操作系统用户通过免密 sudo 执行准确的 sysctl 命令；这是一项重要的主机级权限。
- 清除共享缓存和操作系统缓存会严重干扰其他负载。基准测试应在专用主机执行，并记录缓存重置方法。
- 扩展暴露会影响集群的函数，却没有内置环境保护；必须显式收紧 EXECUTE 权限。
