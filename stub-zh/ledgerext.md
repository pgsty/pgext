## 用法

来源：

- [Official source](https://github.com/geovedi/ledgerext/blob/34c13c0ce3fbb6b1095aac5823fd578d8d994632/README.md)
- [Official source](https://github.com/geovedi/ledgerext/blob/34c13c0ce3fbb6b1095aac5823fd578d8d994632/ledgerext/ledgerext.control)
- [Official source](https://github.com/geovedi/ledgerext/blob/34c13c0ce3fbb6b1095aac5823fd578d8d994632/ledgerext/ledgerext--0.1.sql)
- [Official source](https://github.com/geovedi/ledgerext/blob/34c13c0ce3fbb6b1095aac5823fd578d8d994632/LICENSE)
- [Official source](https://github.com/geovedi/ledgerext/blob/34c13c0ce3fbb6b1095aac5823fd578d8d994632/docs/installation.md)
- [Official source](https://github.com/geovedi/ledgerext/blob/34c13c0ce3fbb6b1095aac5823fd578d8d994632/docs/usage.md)
- [Official source](https://github.com/geovedi/ledgerext/blob/34c13c0ce3fbb6b1095aac5823fd578d8d994632/docs/operations.md)

`ledgerext` 0.1 是面向 PostgreSQL 18 的实验性复式账本。账户及转账状态保存在共享内存中，由 PostgreSQL WAL 之外的独立日志持久化。它不是普通 SQL 表账本，备份、复制、事务和升级流程都必须考虑独立存储。

### 启用与检查

先预加载库并重启，再由管理员创建扩展。默认内存网格约占 1192 MB，应在启动前配置容量并检查主机实际预算。显式大页是可选配置。

```ini
shared_preload_libraries = 'ledgerext'
```

```sql
CREATE EXTENSION ledgerext;
SELECT ledger_capacity();
SELECT ledger_journal_count();
SELECT ledger_expiry_due();
```

管理员在 `ledger_tenants` 中将应用角色绑定到账簿编号，仅授予必要函数。安装时会撤销扩展函数的 PUBLIC 执行权限。租户映射决定账本访问边界，客户端不能在批次中自行传入租户。部分日志及容量管理函数可展示整个集群的信息，应保持受限。

### 账户、批次与状态

`ledger_create_account` 创建单个账户；`ledger_create_accounts` 和 `ledger_apply` 接收打包的 BYTEA 批次。账户、转账的代码及账本编号必须非零。UUID 由客户端提供，转账 ID 支持幂等重试。金额采用无符号 128 位整数单位；`ledger_amount` 将 NUMERIC 金额编码为文档规定的二进制格式。

批次调用将操作排队，直到预提交阶段处理。提交后检查 `ledger_last_statuses`，并用 `ledger_status_name` 或 `ledger_account_status_name` 解码；返回数量并不证明每个事件均已成功。`ledger_balance`、`ledger_lookup_transfer` 及日志查询函数用于检查状态和历史。关联、待处理、入账及撤销标志含义不同，应遵循官方用法指南中的准确批次布局和示例。

### 持久化与维护

`$PGDATA/ledgerext.journal` 是持久账本，`$PGDATA/ledgerext.snap` 用于加速重放。SQL 转储或 PostgreSQL WAL 归档本身不能完整保存该状态。须按上游流程一致备份日志及快照，并演练恢复。日志丢失可能导致重新启动一个空账本。只追加日志不会自动清理，应监测磁盘空间以及固定账户、转账容量。

复制单独传送日志记录，不构成仲裁协议。上游设计不会隔离旧主节点，提升副本须配合外部隔离流程。不能因提供 SQL API 就假定支持 PostgreSQL PITR、逻辑复制、普通表 MVCC 或跨存储原子性。容量耗尽会返回明确状态，工作集不能溢写到磁盘。

本版本没有提供升级脚本，不兼容的日志及快照版本会被拒绝。替换已部署库之前须审查迁移与恢复方案。
