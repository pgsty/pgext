## 用法

来源：

- [官方上游文档](https://github.com/dsarafan/pg_dirty_hands/blob/3d5bb01c39e987bc394fc0f093c37d673744ba7e/README.md)
- [官方扩展 SQL](https://github.com/dsarafan/pg_dirty_hands/blob/3d5bb01c39e987bc394fc0f093c37d673744ba7e/pg_dirty_hands--1.0.sql)
- [官方实现](https://github.com/dsarafan/pg_dirty_hands/blob/3d5bb01c39e987bc394fc0f093c37d673744ba7e/src/pg_dirty_hands.c)

`pg_dirty_hands` 暴露可重写单个堆元组可见性元数据的底层函数。它复制了 PostgreSQL 9.6 时代的清理与冻结内部实现，并有意绕过正常 MVCC 保护。它只能由专家在一次性副本上用于取证或恢复实验；错误的关系或过期 CTID 可能破坏数据、索引、备份和副本。

### 锁定安装权限

C 代码不会独立检查超级用户身份，而扩展函数可能允许 `PUBLIC` 执行。安装后应立即撤销访问，并只授予专用应急角色：

```sql
CREATE EXTENSION pg_dirty_hands;

REVOKE EXECUTE ON FUNCTION freeze_tuple(regclass, tid, boolean) FROM PUBLIC;
REVOKE EXECUTE ON FUNCTION freeze_tuple_unlogged(regclass, tid) FROM PUBLIC;
GRANT EXECUTE ON FUNCTION freeze_tuple(regclass, tid, boolean)
  TO emergency_dba;
```

正常使用绝不能授予 `freeze_tuple_unlogged(regclass, tid)`。

### 重要对象

- `freeze_tuple(regclass, tid, boolean)` 按物理 CTID 冻结一个元组。使用默认的 `false` 时执行源码中有限的资格检查；使用 `true` 时会强制把插入事务标为冻结，并使删除或更新元数据失效。
- `freeze_tuple_unlogged(regclass, tid)` 修改元组头，但既不将缓冲区标为脏，也不写 WAL。尽管名称如此，实现并不会验证关系是否为 unlogged。

### 极高风险边界

上游只报告测试过 PostgreSQL 9.6.3 至 9.6.6。实现与该世代的堆页面布局、缓冲处理、WAL 记录和事务头位紧密耦合。绝不能假定它与其他 PostgreSQL 发布在二进制或语义上兼容。任何调用前都应停止写入、保留物理副本、用取证工具核对准确元组和页面，并演练恢复。绝不能在生产主库上使用 `true` 强制路径或 unlogged 函数。布尔成功结果不能证明逻辑可见性、索引、副本或后续清理仍然有效。
