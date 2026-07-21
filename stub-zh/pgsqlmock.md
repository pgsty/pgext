## 用法

来源：

- [pgSQLMock 1.0.1 文档](https://github.com/v-maliutin/pgSQLMock/blob/Release_v1.0.1/doc/pgsqlmock.md)
- [pgSQLMock 说明文档](https://github.com/v-maliutin/pgSQLMock/blob/Release_v1.0.1/README.md)
- [pgSQLMock 控制文件](https://github.com/v-maliutin/pgSQLMock/blob/Release_v1.0.1/pgsqlmock.control)
- [pgSQLMock 1.0.1 发行版](https://github.com/v-maliutin/pgSQLMock/releases/tag/Release_v1.0.1)

`pgsqlmock` 扩展了 pgTAP，增加了表模拟、函数和视图模拟、调用次数断言以及调试辅助功能。其辅助功能会修改或替换真实的数据库对象，因此需要在基于事务的测试上下文中使用它们，以便在测试结束后回滚这些更改。

```sql
CREATE EXTENSION pgtap;
CREATE EXTENSION pgsqlmock;
```

### 模拟表

`fake_table(text[], ...)` 可以隔离测试与外键、主键、`NOT NULL` 约束、分区或预存行之间的关系。将模式限定的表名作为 `text[]` 传递：

```sql
SELECT plan(2);

SELECT fake_table(
  _table_ident       => ARRAY['app.accounts', 'app.transactions'],
  _make_table_empty  => true,
  _leave_primary_key => false,
  _drop_not_null     => true
);

INSERT INTO app.transactions(account_id, amount)
VALUES (999, 42.00);

SELECT is(
  (SELECT sum(amount) FROM app.transactions WHERE account_id = 999),
  42.00::numeric,
  'transaction logic is isolated from account fixtures'
);

SELECT * FROM finish();
```

重要选项包括 `make_table_empty`、`leave_primary_key`、`drop_not_null`、`drop_collation` 和 `drop_partitions`。在测试中保留主键同时删除参与列的 `NOT NULL` 约束是矛盾的；对于这种测试形状，需要显式地移除或重新创建该键。

### 模拟函数

`mock_func(schema, name, signature, ...)` 临时替换一个例行程序，但保持其身份不变。提供标量值或 SQL/准备语句文本作为结果集：

```sql
CREATE OR REPLACE FUNCTION app.current_business_time()
RETURNS time LANGUAGE sql AS $$ SELECT current_time $$;

SELECT mock_func(
  'app',
  'current_business_time',
  '()',
  _return_scalar_value => '13:00'::time
);

SELECT is(app.current_business_time(), '13:00'::time, 'clock is deterministic');
```

对于返回集合的例行程序，请传递 `_return_set_value` 作为 SQL 查询或已准备语句的名称。使用 `get_routine_signature()` 来确定在重载或默认参数使存储签名模糊时的存储签名。

### 模拟视图

`mock_view(schema, view_name, return_set_sql)` 替换一个带有受控行数的视图：

```sql
SELECT mock_view(
  'app',
  'active_accounts',
  $$SELECT * FROM (VALUES (1, 'test')) AS v(id, name)$$
);

SELECT results_eq(
  'SELECT id, name FROM app.active_accounts',
  $$VALUES (1, 'test')$$,
  'view consumer sees only the fixture'
);
```

### 调用次数和诊断信息

在使用 `track_functions = 'all'` 断言例行程序被调用的次数之前，请设置 `call_count()`：

```sql
SET LOCAL track_functions = 'all';

SELECT call_count(
  1,
  'app',
  'current_business_time',
  '()'
);
```

`print_table_as_json()` 和 `print_query_as_json()` 通过 `NOTICE` 发出可重复的 SQL/JSON 样式快照，这对于当 pgTAP 的回滚会隐藏失败测试期间创建的状态时非常有用。

### 注意事项

- 只在隔离的测试事务中运行模拟和表模拟；它们会发出真实的 `ALTER`、`DROP` 和替换 DDL。
- pgSQLMock 依赖于 PL/pgSQL 和 pgTAP。在运行其断言之前，请加载 pgTAP。
- `call_count()` 依赖于 PostgreSQL 函数统计信息，因此需要设置 `track_functions = 'all'`。
- 发行版 1.0.1 修复了 `fake_table()` 在没有主键的表上删除 `NOT NULL` 约束的问题。
