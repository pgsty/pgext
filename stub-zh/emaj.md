## 用法

来源：

- [E-Maj 5.0.0 README](https://github.com/dalibo/emaj/blob/v5.0.0/README.md)
- [E-Maj 5.0.0 变更日志](https://github.com/dalibo/emaj/blob/v5.0.0/CHANGES.md)
- [E-Maj 快速入门](https://github.com/dalibo/emaj/blob/v5.0.0/docs/en/quickStart.rst)
- [E-Maj 升级指南](https://github.com/dalibo/emaj/blob/v5.0.0/docs/en/upgrade.rst)
- [E-Maj 设置指南](https://github.com/dalibo/emaj/blob/v5.0.0/docs/en/setup.rst)

规范扩展名是 `emaj`；E-Maj 为一个协调表组记录表与序列变更，并可把整个表组回滚到命名标记。它适用于可重复测试、批处理保存点、变更检查与定向恢复，但 E-Maj 回滚不能替代 PostgreSQL 事务回滚或备份。

### 核心流程

```sql
CREATE EXTENSION emaj CASCADE;
GRANT emaj_adm TO app_admin;

SELECT emaj.emaj_create_group('my_group', true);
SELECT emaj.emaj_assign_table('app', 'orders', 'my_group');
SELECT emaj.emaj_assign_sequences('app', '.*', '', 'my_group');

SELECT emaj.emaj_start_group('my_group', 'mark_1');
-- Run application changes.
SELECT emaj.emaj_set_mark_group('my_group', 'mark_2');
-- Run more application changes.

SELECT emaj.emaj_rollback_group('my_group', 'mark_1');
SELECT emaj.emaj_stop_group('my_group');
SELECT emaj.emaj_drop_group('my_group');
```

可回滚表组可以包含多个 schema 中的表与序列，但每张表必须有主键。仅审计表组可记录不可回滚对象的变更。启动和停止表组会锁定其应用表，因此要结合并发流量安排这些操作。

### 重要对象

- `emaj_create_group` 与各类分配函数用于定义表组。
- `emaj_start_group`、`emaj_set_mark_group` 与 `emaj_stop_group` 管理日志会话和标记。
- `emaj_rollback_group` 执行不记日志的回滚；`emaj_logged_rollback_group` 会记录补偿变更。
- 多组变体可在同一时间点操作由组名组成的数组。
- 统计与变更导出函数可检查两个标记之间的变更，或生成用于重放的 SQL。
- `emaj_set_param` 无需直接写内部参数表即可修改或重置 E-Maj 参数。
- `emaj_drop_extension()` 是受支持的完整移除辅助函数。

### 5.0 版本升级

如果 E-Maj 以扩展方式安装且版本不低于 2.3.1，先安装新软件包文件，再执行：

```sql
ALTER EXTENSION emaj UPDATE;
```

文档给出的扩展升级会保留日志，而且表组可以继续处于 LOGGING 状态。切换前应检查以下 5.0 兼容性变化：

- 支持 PostgreSQL 14 到 19；不再支持 PostgreSQL 12 与 13。
- 对 `emaj_param` 的直接 `INSERT`、`UPDATE` 或 `DELETE` 必须改为调用 `emaj_set_param`。
- 幂等启动和停止调用增加了允许已启动或允许已空闲参数；使用命名参数的调用方必须检查参数重命名。
- PHP 命令行客户端与 `emaj_uninstall.sql` 已移除。

通过独立 SQL 脚本安装的环境没有相同的就地扩展升级路径；应遵循官方删除并重装流程。

### 要求与注意事项

标准 `CREATE EXTENSION` 路径需要超级用户权限，并通过 `CASCADE` 安装 `dblink` 与 `btree_gist`。E-Maj 也支持受限的非超级用户脚本安装，其能力限制取决于安装角色。

只有并行回滚客户端需要 `max_prepared_transactions`，其值必须不小于计划使用的会话数；修改后需要重启。大型表组也可能需要更高的 `max_locks_per_transaction`。应把 E-Maj 日志表视为运维数据：明确规划保留策略、监控增长，并继续使用普通备份进行灾难恢复。
