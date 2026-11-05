## 用法

来源：

- [上游来源 1](https://api.pgxn.org/src/dist_emaj/dist_emaj-0.9.0/README.md)
- [上游来源 2](https://api.pgxn.org/src/dist_emaj/dist_emaj-0.9.0/dist_emaj.control)
- [上游来源 3](https://api.pgxn.org/src/dist_emaj/dist_emaj-0.9.0/sql/dist_emaj--0.9.0.sql)
- [上游来源 4](https://api.pgxn.org/src/dist_emaj/dist_emaj-0.9.0/client/distEmaj.pl)

`dist_emaj` 0.9.0 协调跨数据库的 E-Maj 操作，使已设置标记的组能够一致回滚。

### 核心流程

在协调库中安装扩展及其 `dblink` 依赖。用 `dist_emaj.create_database` 注册连接定义，通过 `dist_emaj.create_cluster` 定义逻辑集群，再用 `dist_emaj.assign_group` 分配 E-Maj 组。协调操作前运行 `dist_emaj.verify_cluster`。分发中的 `distEmaj.pl` 与 `distEmajRollback.pl` 客户端负责分布式标记和回滚。

```sql
CREATE EXTENSION dist_emaj CASCADE;
SELECT extversion FROM pg_extension WHERE extname = 'dist_emaj';
```

### 运行边界

每个受管数据库均需兼容的 E-Maj 及获授权的远程连接。分布式回滚使用两阶段提交，需配置预备事务容量，并保留中断操作的恢复流程。连接秘密应按特权数据管理，调用时遵循对应版本的 SQL 函数签名。本扩展不会为其他应用写入自动提供分布式事务语义。
