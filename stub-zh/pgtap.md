## 用法

来源：

- [v1.3.4 README](https://github.com/theory/pgtap/blob/v1.3.4/README.md)
- [v1.3.4 release history](https://github.com/theory/pgtap/blob/v1.3.4/Changes)
- [Control file](https://github.com/theory/pgtap/blob/v1.3.4/pgtap.control)
- [SQL definitions](https://github.com/theory/pgtap/blob/v1.3.4/sql/pgtap.sql.in)

`pgtap` 是一个 PostgreSQL 单元测试框架，输出 TAP（Test Anything Protocol）格式的结果，提供数百个断言函数用于测试数据库对象和查询结果。

```sql
CREATE EXTENSION pgtap;
```

### 测试结构

```sql
BEGIN;
SELECT plan(3);  -- declare how many tests to run

SELECT ok(1 = 1, 'one equals one');
SELECT is(1 + 1, 2, 'addition works');
SELECT isnt(1, 2, 'one is not two');

SELECT * FROM finish();
ROLLBACK;
```

当测试数量未知时使用 `no_plan()`：

```sql
BEGIN;
SELECT * FROM no_plan();
SELECT ok(2 > 1, 'comparison works');
SELECT * FROM finish();
ROLLBACK;
```

### 基本断言

```sql
SELECT ok(expression, description);           -- boolean test
SELECT is(got, expected, description);         -- equality test
SELECT isnt(got, unexpected, description);     -- inequality test
SELECT matches(value, regex, description);     -- regex match
```

### 模式测试

```sql
SELECT has_table('users');
SELECT has_table('myschema', 'users', 'users table exists');
SELECT has_column('users', 'email');
SELECT col_type_is('users', 'email', 'text');
SELECT col_not_null('users', 'id');
SELECT col_has_default('users', 'created_at');
SELECT has_function('calculate_total');
SELECT has_function('calculate_total', ARRAY['integer', 'numeric']);
SELECT has_index('users', 'users_email_idx');
SELECT has_pk('users');
SELECT has_fk('orders');
```

### 错误测试

```sql
SELECT lives_ok('INSERT INTO t(id) VALUES (1)', 'insert succeeds');
SELECT throws_ok(
  'SELECT 1/0',
  '22012',          -- SQLSTATE for division by zero
  'division by zero'
);
```

### 查询结果测试

```sql
-- Compare ordered result sets
SELECT results_eq(
  'SELECT * FROM active_users()',
  'SELECT * FROM users WHERE active',
  'active_users returns correct rows'
);

-- Compare unordered result sets
SELECT set_eq(
  'SELECT * FROM active_ids()',
  ARRAY[2, 3, 4, 5]
);

-- Check query returns no rows
SELECT is_empty('SELECT * FROM users WHERE id = -1');

-- Compare bag (multiset) results
SELECT bag_eq(
  'SELECT color FROM items',
  $$VALUES ('red'), ('blue'), ('red')$$
);
```

### 使用 pg_prove 运行测试

```bash
pg_prove -d mydb tests/*.sql
pg_prove -d mydb --ext .sql --recurse tests/
```

### xUnit 风格

```sql
CREATE FUNCTION test_my_feature() RETURNS SETOF text AS $$
BEGIN
  RETURN NEXT ok(1 = 1, 'basic check');
  RETURN NEXT is(abs(-1), 1, 'absolute value works');
END;
$$ LANGUAGE plpgsql;

SELECT * FROM runtests('test_my_feature');
```

### 1.3.4 版本与测试边界

1.3.4 新增 `index_is_partial()`，为 `has_composite()` 和 `hasnt_composite()` 增加 name/name 重载，并修复若干旧版本升级路径。先安装匹配的脚本，再执行 `ALTER EXTENSION pgtap UPDATE TO '1.3.4'`。控制文件要求 `plpgsql`，设置 `superuser = false`，允许重定位，且无需共享预加载。角色仍须拥有数据库 CREATE 权限，以及测试所涉及对象的适当权限。

应使用可丢弃的测试数据或隔离测试库；事务回滚无法撤销被测试函数执行的外部动作。pg_prove 客户端需要单独安装，安装扩展本身不会提供这个测试执行器。
