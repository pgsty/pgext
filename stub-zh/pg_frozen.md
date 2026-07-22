## 用法

来源：

- [官方 README](https://github.com/lawrencejones/pg_frozen/blob/68004b4f21b01041487e2a080658565dc4cdc937/README.md)
- [0.0.1 版本扩展 SQL](https://github.com/lawrencejones/pg_frozen/blob/68004b4f21b01041487e2a080658565dc4cdc937/pg_frozen--0.0.1.sql)
- [C 实现](https://github.com/lawrencejones/pg_frozen/blob/68004b4f21b01041487e2a080658565dc4cdc937/pg_frozen.c)

`pg_frozen` 暴露一个底层诊断函数，用于检查堆元组头并报告其插入事务 ID 是否已冻结。它适合小范围 VACUUM 实验，不应作为持久元组属性，也不能替代 PostgreSQL 的清理统计。

### 核心流程

```sql
CREATE EXTENSION pg_frozen;

CREATE TABLE messages (id bigint GENERATED ALWAYS AS IDENTITY, msg text);
INSERT INTO messages (msg) VALUES ('hello');

SELECT id, ctid, xmin, frozen(tableoid, ctid)
FROM messages;

VACUUM (FREEZE) messages;

SELECT id, ctid, xmin, frozen(tableoid, ctid)
FROM messages;
```

可见元组通常在冻结前报告 `0`，其元组头被标记为冻结后报告 `1`。

### 对象索引

- `frozen(tableoid oid, tid tid) RETURNS integer` 获取由关系 OID 和 `ctid` 标识的物理元组；它返回元组头冻结标志，无法获取元组时返回 `3`。

### 运维说明

版本 `0.0.1` 可重定位，且未声明依赖或预加载要求。函数为 `STRICT` 和 `PARALLEL SAFE`，但它通过 PostgreSQL 内部接口读取物理堆状态。

始终在同一条语句中从同一个当前行取得 `tableoid` 和 `ctid`。`ctid` 不是稳定标识符：UPDATE、VACUUM、剪枝和并发活动都可能替换或删除对应元组。扩展使用 `SnapshotAny`，因此结果是实现层观察，可能包含正常 MVCC 可见性之外的元组。不要将其用于应用逻辑，并应针对实际使用的 PostgreSQL 大版本测试。
