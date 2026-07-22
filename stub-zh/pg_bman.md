## 用法

来源：

- [官方上游文档](https://github.com/s-hironobu/pg_bman/blob/3d812f26abed778c44645259eaac2ede6e398a9d/README.md)
- [官方扩展 SQL](https://github.com/s-hironobu/pg_bman/blob/3d812f26abed778c44645259eaac2ede6e398a9d/pg_bman--1.0.sql)
- [官方扩展实现](https://github.com/s-hironobu/pg_bman/blob/3d812f26abed778c44645259eaac2ede6e398a9d/pg_bman.c)

`pg_bman` 1.0 是 2014 年前后远程备份工具包的一个组件。其可选服务器扩展允许独立的 `pg_archivebackup` 客户端通过 libpq 列出并获取归档 WAL 段；`pg_bman` shell 脚本则协调完整与增量备份仓库。它不是当前 PostgreSQL 备份方案，只能用于理解或恢复固定的旧部署。

### 扩展 API

SQL 扩展安装两个函数，撤销 `PUBLIC` 的默认访问，并在 C 层再次检查超级用户身份：

```sql
CREATE EXTENSION pg_bman;

SELECT *
FROM pg_show_archives('/var/lib/postgresql/data/archives');

SELECT octet_length(
  pg_get_archive(
    '/var/lib/postgresql/data/archives/000000010000000000000001'
  )
);
```

`pg_show_archives(text)` 返回看似 24 位十六进制 WAL 段的名称。`pg_get_archive(text)` 将最多一个 WAL 段读入 `bytea`。它们是内部传输原语，不是备份一致性或恢复验证 API。

### 旧式备份流程

上游流程配置 PostgreSQL 归档，创建备份服务器仓库，并将外部脚本指向 `pg_basebackup`、归档目录与 libpq 连接。命令包括 `pg_bman BACKUP FULL`、`pg_bman BACKUP INCREMENTAL`、`pg_bman SHOW` 和 `pg_bman RESTORE`。恢复输出还必须通过外部工具复制回已经停止的 PostgreSQL 数据目录。

### 未验证时不得用于现代 PostgreSQL

文档配置面向 PostgreSQL 9.3 时代的 WAL 与恢复行为，包括 `recovery.conf` 和手工复制归档。现代 PostgreSQL 的恢复信号、WAL 布局、权限、校验和、时间线、表空间与备份清单都有实质差异。源码中的文件名检查同样较旧；即使调用要求超级用户，也不能把它视为通用安全文件 API。

对于仍保留的旧部署，应隔离备份角色与网络，保持撤销 `PUBLIC`，验证仓库权限，并完整演练恢复到独立主机。`BACKUP` 命令成功或 WAL 读取成功，都不能证明集群可恢复。受支持的 PostgreSQL 应选择有当前版本恢复测试、保留、加密、时间线处理与完整性验证的受维护物理备份工具。
