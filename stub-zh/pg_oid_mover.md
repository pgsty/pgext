## 用法

来源：

- [官方资料](https://github.com/e-buzyurkin/pg_oid_mover/blob/f5a02b2162fa356dff49312980d1b60ee29d2496/README.md)
- [官方资料](https://github.com/e-buzyurkin/pg_oid_mover/blob/f5a02b2162fa356dff49312980d1b60ee29d2496/pg_oid_mover.control)
- [官方资料](https://github.com/e-buzyurkin/pg_oid_mover/blob/f5a02b2162fa356dff49312980d1b60ee29d2496/pg_oid_mover--1.0.sql)
- [官方资料](https://github.com/e-buzyurkin/pg_oid_mover/blob/f5a02b2162fa356dff49312980d1b60ee29d2496/pg_oid_mover.c)

pg_oid_mover 是供一次性集群使用的测试扩展，用于向前移动 PostgreSQL 集群级 OID 分配位置，以测试高 OID 情况。该变化不随 SQL 事务回滚，禁止用于生产集群。

### 测试接口

```sql
CREATE EXTENSION pg_oid_mover;
SELECT oid_mover.advance_to(3000000000::bigint);
```

仅由超级用户在一次性测试集群中选择合适目标后运行。advance_to(bigint) 返回此前的下一个 OID，只接受有效 OID 范围内的向前移动。所有数据库共用此分配器，开启并回滚事务不会撤销推进。

### 边界

扩展不可迁移，安装在 oid_mover 模式中，并撤销 PUBLIC 执行权限；恢复状态、只读事务和并行执行都会被拒绝。无需预加载。必须匹配服务端 ABI；源码没有确立通用主版本支持矩阵或许可证。
