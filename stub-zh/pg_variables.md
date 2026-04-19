## 用法

- 来源：[README](https://github.com/postgrespro/pg_variables/blob/master/README.md)，[repository tags](https://github.com/postgrespro/pg_variables/tags)，[control file](https://github.com/postgrespro/pg_variables/blob/master/pg_variables.control)

`pg_variables` 提供按名称 package 分组的会话级变量。变量只存在于当前 session；除非创建时指定 `is_transactional := true`，否则默认不参与事务。

### 基本读写

```sql
CREATE EXTENSION pg_variables;

SELECT pgv_set('vars', 'int1', 101);
SELECT pgv_get('vars', 'int1', NULL::int);
```

事务性变量会参与 savepoint 和 rollback：

```sql
BEGIN;
SELECT pgv_set('vars', 'trans_int', 101, true);
SAVEPOINT sp1;
SELECT pgv_set('vars', 'trans_int', 102, true);
ROLLBACK TO sp1;
COMMIT;
```

### 核心 API

README 记录了通用标量与数组 API：

- `pgv_set(package, name, value, is_transactional default false)`
- `pgv_get(package, name, NULL::type, strict default true)`

同时也记录了面向 record 的 API：

- `pgv_insert()`
- `pgv_update()`
- `pgv_delete()`
- `pgv_select()`

常用的管理辅助函数包括 `pgv_exists()`、`pgv_remove()`、`pgv_free()`、`pgv_list()` 和 `pgv_stats()`。

### 错误与 strict 行为

`pgv_get()` 会同时检查变量是否存在以及类型是否匹配。README 展示了 package 不存在、变量不存在或类型不匹配时会抛错；如果指定 `strict := false`，则在值缺失时返回 `NULL`。

### 已废弃辅助函数与版本说明

上游仍然提供一些已废弃的类型专用辅助函数，例如 `pgv_set_int()` / `pgv_get_int()` 和 `pgv_set_jsonb()` / `pgv_get_jsonb()`，但推荐使用通用的 `pgv_set()` / `pgv_get()`。

仓库 tag 已到 `v1.2.5`，而当前 `pg_variables.control` 仍声明 `default_version = '1.2'`。这与打包说明一致，即 release tag 已前进，但 SQL 扩展版本字符串没有同步变更。
