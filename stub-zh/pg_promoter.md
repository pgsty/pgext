## 用法

来源：

- [官方上游 README](https://github.com/masahikosawada/pg_promoter/blob/9e70f65508e2bfe8c1e30631dc709353127c9ce0/README.md)
- [官方扩展控制文件 (pg_promoter.control)](https://github.com/masahikosawada/pg_promoter/blob/9e70f65508e2bfe8c1e30631dc709353127c9ce0/pg_promoter.control)
- [官方实现源码](https://github.com/masahikosawada/pg_promoter/blob/9e70f65508e2bfe8c1e30631dc709353127c9ce0/pg_promoter.c)

`pg_promoter` 是一个早期的后台工作原型，仅在备用节点上运行，定期检查主节点，并在配置次数的失败探测后本地进行提升。它不提供完整的共识、票数或隔离系统。

### 核心工作流

将库安装在临时备用节点上，配置为预加载并重启：

```ini
shared_preload_libraries = 'pg_promoter'
pg_promoter.keepalives_time = 5
pg_promoter.keepalives_count = 3
pg_promoter.primary_conninfo = 'host=primary port=5432 dbname=postgres'
```

源代码拼写形式为 `keepalives_time` 和 `keepalives_count`；即使 README 中的示例使用单数名称，也请使用这些复数形式。工作进程在连续连接失败达到配置次数后进行提升。

### 重要设置

- `pg_promoter.primary_conninfo` 是用于探测主节点的 libpq 连接字符串。
- `pg_promoter.keepalives_time` 是轮询间隔（秒）。
- `pg_promoter.keepalives_count` 是在达到提升阈值前的失败探测次数。

### 要求与注意事项

- 经审核的控制文件、注册表或目录证据标识版本为 `1.0`。
- 控制文件将扩展标记为可重定位。
- 这是一个仅用于预加载的工作进程；经审核的仓库未提供扩展 SQL。
- 网络故障并不能证明旧主节点已下线。在没有外部隔离的情况下，此设计可能会导致脑裂和数据丢失。
- 代码使用了过时的触发文件提升路径和 PostgreSQL 内部实现。请将其视为历史原型代码，并仅在与编写它的服务器版本对应的环境中进行测试。
