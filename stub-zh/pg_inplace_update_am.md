## 用法

来源：

- [Official usage](https://github.com/igor-suhorukov/pg_inplace_update_am/blob/7b868ca7aee12af6faf2c1c3804964d7aff81a10/.github/README.md)
- [Extension control](https://github.com/igor-suhorukov/pg_inplace_update_am/blob/7b868ca7aee12af6faf2c1c3804964d7aff81a10/pg_inplace_update_am.control)
- [Versioned SQL](https://github.com/igor-suhorukov/pg_inplace_update_am/blob/7b868ca7aee12af6faf2c1c3804964d7aff81a10/sql/pg_inplace_update_am--0.1.sql)
- [Initialization and settings](https://github.com/igor-suhorukov/pg_inplace_update_am/blob/7b868ca7aee12af6faf2c1c3804964d7aff81a10/src/iheap_handler.c)

`pg_inplace_update_am` 0.1 在未经修改的 PostgreSQL 19 上提供 `iheap` 表访问方法与 `dmbtree` 索引访问方法。行可以原位更新，undo 记录为旧快照与回滚保留历史版本。这是早期开发中的存储代码；用于重要数据前，应验证恢复、复制和升级流程。

### 启用与使用

超级用户须预加载模块并重启 PostgreSQL，以注册其 WAL 资源管理器：

```ini
shared_preload_libraries = 'pg_inplace_update_am'
```

随后在目标数据库安装 SQL 对象：

```sql
CREATE EXTENSION pg_inplace_update_am;
CREATE TABLE accounts (id int PRIMARY KEY, branch int, balance numeric)
    USING iheap;
CREATE INDEX ON accounts (branch);
INSERT INTO accounts VALUES (1, 1, 100);
UPDATE accounts SET balance = balance + 1 WHERE id = 1;
UPDATE accounts SET branch = 2 WHERE id = 1;
```

示例中的普通索引默认使用 `dmbtree`；主键与唯一索引默认使用 `btree`。使用 `dmbtree` 时，修改索引列仍可原位更新；普通 `btree` 索引则需要移动该行。

### 设置与检查

- `iheap.enable_inplace` 控制原位更新；关闭后，每次更新都会移动行。
- `iheap.default_index_am` 默认为 `dmbtree`；`iheap.unique_index_am` 默认为 `btree`。
- `iheap.foreground_undo_limit` 默认为 1024 kB。更大的事务中止会把部分回滚工作留给后续访问与 VACUUM。
- `iheap.logical_decoding` 默认为 `on`；设为 `ignore` 会在解码时跳过这些表的变化。
- `iheap_relation_stats`、`iheap_metapage_info` 和 `dmbtree_index_stats` 用于检查存储；`iheap_check`、`dmbtree_check` 与 `dmbtree_parent_check` 用于验证其结构。
- `iheap_stat_tables` 展示统计信息。`iheap_discard` 丢弃所有快照都不再需要的 undo；`dmbtree_copy_btree_opclasses` 复制安装后新增的 B-tree 操作符类。

安装 SQL 撤销了多数底层检查、验证和维护函数的公共执行权限。运维流程需要包含 undo 清理与 VACUUM；上游提供测试套件，并不等于已达到生产可用标准。
