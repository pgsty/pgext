## 用法

来源：

- [已复核 commit 的上游 README](https://github.com/pgbouncer/pg_pgbouncer/blob/e71619b01d391316c98eb8456e5b0b37c76d96d9/README.md)
- [扩展模式与引导数据](https://github.com/pgbouncer/pg_pgbouncer/blob/e71619b01d391316c98eb8456e5b0b37c76d96d9/pg_pgbouncer.sql)
- [进程管理器实现](https://github.com/pgbouncer/pg_pgbouncer/blob/e71619b01d391316c98eb8456e5b0b37c76d96d9/src/child.rs)

`pg_pgbouncer` 是一个 alpha 质量的 pgrx 扩展：它运行 PostgreSQL 后台工作进程，生成配置并监管外部 `pgbouncer` 进程。上游明确不建议用于生产，并提醒 API 可能发生重大变化。

```conf
shared_preload_libraries = 'pg_pgbouncer'
pg_pgbouncer.database = 'pg_pgbouncer'
```

重启 PostgreSQL，在该精确数据库中创建 `pg_pgbouncer`，并确保操作系统账户可以找到和执行 `pgbouncer` 二进制文件。配置存储于 `pgbouncer.groups`、`pgbouncer.databases`、`pgbouncer.users`、`pgbouncer.peers`、`pgbouncer.auth`、`pgbouncer.hba` 和 `pgbouncer.settings`；管理器写出 INI 文件，并启动子进程或向其发送信号。

引导 SQL 会从 `pg_shadow` 复制密码值，把凭据存进普通表，启用 `trust` 认证，并加入示例用户。应把它视为不安全的示例状态：隔离在可丢弃集群，撤销 `PUBLIC` 的模式和表权限，替换所有默认值，保护生成文件，并审计进程所有权和信号处理。

工作进程作用于整个集群，并以超级用户身份通过 SPI 连接。数据库不可用、畸形表行、可执行文件缺失、过期 PID/套接字或子进程反复崩溃，都可能扰乱监管。生产环境应使用操作系统服务管理器和 PgBouncer 官方支持的配置工作流。
