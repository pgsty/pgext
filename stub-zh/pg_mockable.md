


## 用法

> 来源：[pg_mockable upstream README](https://github.com/bigsmoke/pg_mockable/blob/v1.1.0/README.md)、[v1.1.0 tag](https://github.com/bigsmoke/pg_mockable/tree/v1.1.0)、[PGXN pg_mockable](https://pgxn.org/dist/pg_mockable/)、本地源码归档 `pg_mockable-1.1.0.tar.gz`。

`pg_mockable` 用于为 PostgreSQL 函数创建可 Mock 的包装函数。它主要面向数据库测试：应用代码调用稳定的包装函数，而测试可以临时替换包装函数的返回值。

```sql
CREATE EXTENSION pg_mockable CASCADE;
```

该扩展会安装到固定的 `mockable` schema 中，且不可重定位。

### Mock 内置时间函数

`mockable.now()` 已经预先创建，因为 Mock `now()` 也会覆盖该扩展暴露的一组当前时间包装函数。

```sql
SELECT mockable.now();

SELECT mockable.mock(
  'pg_catalog.now()',
  '2026-06-17 10:00:00+08'::timestamptz
);

SELECT mockable.now();
SELECT mockable.current_timestamp();
SELECT mockable.current_date();

CALL mockable.unmock('pg_catalog.now()');
```

`mockable.mock(regprocedure, anyelement)` 会保存 mock 值并返回它。`mockable.unmock(regprocedure)` 会清除 mock，并让包装函数恢复为调用原始函数。

### 包装应用函数

使用 `mockable.wrap_function()` 在 `mockable` schema 中创建一个薄包装函数：

```sql
CREATE SCHEMA app;

CREATE FUNCTION app.answer()
RETURNS int
LANGUAGE sql
RETURN 42;

SELECT mockable.wrap_function('app.answer()');

SELECT mockable.answer();                 -- 42
SELECT mockable.mock('app.answer()', 7);   -- 7
SELECT mockable.answer();                 -- 7

CALL mockable.unmock('app.answer()');
SELECT mockable.answer();                 -- 42
```

第一个参数是 `regprocedure`，因此函数存在重载时需要写出参数类型：

```sql
SELECT mockable.wrap_function('pg_catalog.current_setting(text, boolean)');
```

如果自动生成包装函数不够用，可以把精确的 `CREATE OR REPLACE FUNCTION` 语句作为第二个参数传入：

```sql
SELECT mockable.wrap_function(
  'app.answer()',
  $$
  CREATE OR REPLACE FUNCTION mockable.answer()
  RETURNS int
  LANGUAGE sql
  RETURN app.answer();
  $$
);
```

版本 1.1.0 还通过 `mockable.wrap_function(...)` 的 `raise_debug_messages$` 参数和 `mock_memory.raise_debug_messages` 列，为 wrapped/mockable routines 增加可选 debug logging。

### Mock 生命周期

默认 mock 生命周期是事务级的。若值需要跨 dump/restore 或后续事务保留，可以在创建包装函数时使用持久生命周期：

```sql
SELECT mockable.wrap_function(
  'app.answer()',
  mock_duration$ => 'PERSISTENT'
);
```

测试夹具不再需要持久 mock 时，应显式清除：

```sql
CALL mockable.unmock('app.answer()');
```

### Search Path 注意事项

应用代码必须实际调用包装函数，例如 `mockable.now()` 或 `mockable.answer()`，mock 才会生效。某些 PL/pgSQL 代码可以通过调整 `search_path` 重定向，但表默认值等表达式会被编译为函数 OID；事后再把 `mockable` 加进 `search_path` 不会改写这些引用。需要可测试的代码应优先显式调用 `mockable.*`。

### 注意事项

- 版本 1.1.0 支持 PostgreSQL 14-18。它是 SQL 扩展，不需要 `shared_preload_libraries`。
- `pg_mockable` 拥有 `mockable` schema；control 文件不支持把它安装到其他 schema。
- 包装函数权限来自被包装的原函数。上游测试会验证：包装一个私有函数不会把执行权限授予原本无法调用该函数的角色。
