## 用法

来源：

- [README.md](https://api.pgxn.org/src/pg_promise_guard/pg_promise_guard-0.2.8/README.md)
- [pg_promise_guard.control](https://api.pgxn.org/src/pg_promise_guard/pg_promise_guard-0.2.8/pg_promise_guard.control)
- [CHANGELOG.md](https://api.pgxn.org/src/pg_promise_guard/pg_promise_guard-0.2.8/CHANGELOG.md)

`pg_promise_guard` 0.2.8 扫描系统目录，寻找已声明却不再完整执行的保证。它依赖 `pg_living_assertions`，以纯 SQL 实现，无需预加载。

### 核心用法

```sql
CREATE EXTENSION pg_living_assertions;
CREATE EXTENSION pg_promise_guard;

SELECT * FROM promise_guard.promise_breaks;
SELECT promise_guard.promises_kept();
SELECT promise_guard.promises_kept('app');
SELECT * FROM promise_guard.check_promises('app');
```

```sql
SELECT promise_guard.watch('public');
SELECT * FROM living_assertions.status;
```

### 运行边界

`promise_guard.check_promises()` 返回保证破坏与缺口，`promise_guard.promises_kept()` 仅在存在破坏时返回假，`promise_guard.watch()` 将结论写入持续断言。只接受 READ COMMITTED，不存在的模式会报错。被禁用的触发器、未执行的约束及失效的策略或规则可判为破坏；仅在复制模式触发的触发器、未验证的历史数据、已就绪但无效的唯一索引可判为缺口，仍允许重复值的失败索引则属于破坏。仅排除超级用户所属扩展的对象。0.2.7 之前的监测应退役后重建，以采用系统目录优先的名称解析。扫描不能证明物理完整性，也不审计全部权限。PGXN 声明 PostgreSQL 13+，文档测试覆盖 18.6 与 19 beta。Apache-2.0 从 0.2.3 开始适用。
