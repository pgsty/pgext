

## 用法

> [pgtap: PostgreSQL 单元测试框架](https://github.com/theory/pgtap)

`pgtap` 是一个 PostgreSQL 单元测试框架，输出 TAP（Test Anything Protocol）格式的结果，提供数百个断言函数用于测试数据库对象和查询结果。

```sql
CREATE EXTENSION pgtap;
```

### 测试结构

```sql
BEGIN;
SELECT plan(3);  -- 声明要运行的测试数量

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
-- ... 测试 ...
SELECT * FROM finish();
ROLLBACK;
```

### 基本断言

```sql
SELECT ok(expression, description);           -- 布尔测试
SELECT is(got, expected, description);         -- 相等测试
SELECT isnt(got, unexpected, description);     -- 不等测试
SELECT matches(value, regex, description);     -- 正则匹配
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
  '22012',          -- 除零错误的 SQLSTATE
  'division by zero'
);
```

### 查询结果测试

```sql
-- 比较有序结果集
SELECT results_eq(
  'SELECT * FROM active_users()',
  'SELECT * FROM users WHERE active',
  'active_users returns correct rows'
);

-- 比较无序结果集
SELECT set_eq(
  'SELECT * FROM active_ids()',
  ARRAY[2, 3, 4, 5]
);

-- 检查查询返回空结果
SELECT is_empty('SELECT * FROM users WHERE id = -1');

-- 比较多重集结果
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
  RETURN NEXT ok(1 = 1, 'basic check');
  RETURN NEXT is(my_func(1), 42, 'function works');
$$ LANGUAGE plpgsql;

SELECT * FROM runtests('test_my_feature');
```
