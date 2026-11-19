## 用法

来源：

- [Official source](https://github.com/diviyank/pg_reflex/blob/c1e787fe390ad2398bcdd438e51e4c38bdb8978c/README.md)
- [Official source](https://github.com/diviyank/pg_reflex/blob/c1e787fe390ad2398bcdd438e51e4c38bdb8978c/Cargo.toml)
- [Official source](https://github.com/diviyank/pg_reflex/blob/c1e787fe390ad2398bcdd438e51e4c38bdb8978c/pg_reflex.control)
- [Official source](https://github.com/diviyank/pg_reflex/blob/c1e787fe390ad2398bcdd438e51e4c38bdb8978c/docs/api/create_reflex_ivm.md)
- [Official source](https://github.com/diviyank/pg_reflex/blob/c1e787fe390ad2398bcdd438e51e4c38bdb8978c/docs/getting-started/upgrading.md)
- [Official source](https://github.com/diviyank/pg_reflex/blob/c1e787fe390ad2398bcdd438e51e4c38bdb8978c/sql/pg_reflex--1.12.4--1.12.5.sql)
- [Official source](https://github.com/diviyank/pg_reflex/releases/tag/1.12.5)

`pg_reflex` 1.12.5 在 PostgreSQL 15–18 上通过源表触发器维护查询结果。安装需要超级用户；扩展不是 trusted，且不可迁移模式。官方启用流程使用 `CREATE EXTENSION`，无需共享预加载。

### 创建并检查维护结果

```sql
CREATE EXTENSION pg_reflex;
CREATE TABLE sales (id bigint PRIMARY KEY, region text, amount numeric);
SELECT create_reflex_ivm(
    'sales_by_region',
    'SELECT region, SUM(amount) AS total FROM sales GROUP BY region',
    NULL, 'LOGGED', 'IMMEDIATE'
);
INSERT INTO sales VALUES (1, 'EU', 150);
SELECT * FROM sales_by_region;
SELECT * FROM reflex_ivm_status();
```

结果由普通表和自动生成的维护对象承载。查询可以包含受支持的聚合、连接和拆分后的 CTE，但并非所有 SELECT 都可维护。当前创建检查会拒绝不支持的查询形式；部分允许的形式在写入时仍会重算整个结果，不能假定更新耗时恒定。

示例明确选择 `LOGGED`。默认值为 `UNLOGGED`，其数据不具备崩溃持久性。默认维护模式为 `IMMEDIATE`；`DEFERRED` 将工作合并到 COMMIT。直接投影结果的删除维护需要唯一键，可从合适的源表主键推断。使用 `ignore_sources` 会主动允许结果陈旧，直到执行文档规定的校正或修复操作。

### 管理与权限

分别使用 `reflex_reconcile`、`reflex_rebuild_imv` 和 `drop_reflex_ivm` 执行文档规定的校正、重建和移除。生成的触发器与内部登记表属于维护协议，不应直接修改。创建时必须具备源表与目标模式所需权限，再按需授予结果访问权。安装包含事件触发器以及范围有限的 security-definer 分区辅助函数；能够安装扩展并不等于已建立通用应用安全边界。

### 升级边界

替换库后，应及时在各数据库执行 `ALTER EXTENSION pg_reflex UPDATE TO '1.12.5'`。迁移会新增辅助函数并重写源表增量触发器，应检查未匹配函数体的警告。从 1.12.0 之前升级，还需要按指南执行 SQL 更新、验证定义并重建每个 IMV；仅复制库文件会保留旧的生成维护代码。

1.12.5 修复管理操作与提交刷新之间的锁顺序、同一语句改变连接两侧源表时的维护，以及非空分区删除。应按升级指南校正此前受影响的连接结果。源表 DDL、分区变化、延迟维护失败和崩溃恢复都需要操作检查，不能只验证 SELECT 成功。
