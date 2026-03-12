

## 用法

> [anon: PostgreSQL 匿名化与数据脱敏](https://gitlab.com/dalibo/postgresql_anonymizer/)

`postgresql_anonymizer`（扩展名：`anon`）使用声明式方法对个人身份信息（PII）进行脱敏或替换。脱敏规则直接通过安全标签在数据库模式中定义。

```sql
CREATE EXTENSION IF NOT EXISTS anon CASCADE;
SELECT anon.init();
```

### 声明脱敏规则

```sql
SECURITY LABEL FOR anon ON COLUMN player.name
  IS 'MASKED WITH FUNCTION anon.fake_last_name()';

SECURITY LABEL FOR anon ON COLUMN player.id
  IS 'MASKED WITH VALUE NULL';
```

### 静态脱敏

永久替换数据库中的 PII：

```sql
SECURITY LABEL FOR anon ON COLUMN customer.full_name
  IS 'MASKED WITH FUNCTION anon.fake_first_name() || '' '' || anon.fake_last_name()';

SECURITY LABEL FOR anon ON COLUMN customer.birth
  IS 'MASKED WITH FUNCTION anon.random_date_between(''1920-01-01''::DATE, now())';

SELECT anon.anonymize_database();
-- 也可使用：anon.anonymize_table()、anon.anonymize_column()
```

### 动态脱敏

对特定角色隐藏 PII，其他角色可见原始数据：

```sql
SELECT anon.start_dynamic_masking();

CREATE ROLE skynet LOGIN;
SECURITY LABEL FOR anon ON ROLE skynet IS 'MASKED';

SECURITY LABEL FOR anon ON COLUMN people.lastname
  IS 'MASKED WITH FUNCTION anon.fake_last_name()';

SECURITY LABEL FOR anon ON COLUMN people.phone
  IS 'MASKED WITH FUNCTION anon.partial(phone, 2, $$******$$, 2)';
```

当 `skynet` 查询该表时，会自动返回脱敏后的数据。

### 匿名导出

```bash
pg_dump_anon.sh -h localhost -p 5432 -U bob bob_db > dump.sql
```

### 常用脱敏函数

| 函数 | 描述 |
|----------|-------------|
| `anon.fake_first_name()` | 随机名字 |
| `anon.fake_last_name()` | 随机姓氏 |
| `anon.fake_company()` | 随机公司名 |
| `anon.random_date_between(d1, d2)` | 范围内随机日期 |
| `anon.random_zip()` | 随机邮政编码 |
| `anon.partial(value, prefix, padding, suffix)` | 部分混淆 |
| `anon.random_string(n)` | 长度为 n 的随机字符串 |
| `anon.random_int_between(i1, i2)` | 范围内随机整数 |
