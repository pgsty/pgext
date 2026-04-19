
## 用法

> 来源：[overview](https://postgresql-anonymizer.readthedocs.io/en/stable/), [static masking](https://postgresql-anonymizer.readthedocs.io/en/stable/static_masking/), [dynamic masking](https://postgresql-anonymizer.readthedocs.io/en/stable/dynamic_masking/), [anonymous dumps](https://postgresql-anonymizer.readthedocs.io/en/stable/anonymous_dumps/), [masking functions](https://postgresql-anonymizer.readthedocs.io/en/stable/masking_functions/)

`anon` 通过 `SECURITY LABEL FOR anon` 以声明式方式定义脱敏规则。官方文档主要围绕三条用户侧流程展开：永久脱敏、masked role，以及匿名导出。

### 初始化与声明规则

```sql
CREATE EXTENSION IF NOT EXISTS anon CASCADE;
SELECT anon.init();

SECURITY LABEL FOR anon ON COLUMN customer.full_name
IS 'MASKED WITH FUNCTION anon.dummy_name()';

SECURITY LABEL FOR anon ON COLUMN customer.employer
IS 'MASKED WITH FUNCTION anon.dummy_company_name()';

SECURITY LABEL FOR anon ON COLUMN customer.phone
IS 'MASKED WITH FUNCTION anon.partial(phone, 2, $$******$$, 2)';
```

### 静态脱敏

静态脱敏会直接原地改写已存储的数据：

```sql
SELECT anon.anonymize_database();
-- See also: anon.anonymize_table(), anon.anonymize_column()
```

静态脱敏文档还覆盖了 `shuffling`、噪声注入，以及面向更大数据集的并行脱敏。

### 动态脱敏

动态脱敏只会对被标记为 masked 的角色隐藏值：

```sql
ALTER DATABASE demo SET session_preload_libraries = 'anon';
ALTER DATABASE demo SET anon.transparent_dynamic_masking TO true;

CREATE ROLE skynet LOGIN;
SECURITY LABEL FOR anon ON ROLE skynet IS 'MASKED';
GRANT pg_read_all_data TO skynet;

SECURITY LABEL FOR anon ON COLUMN people.lastname
IS 'MASKED WITH FUNCTION anon.dummy_last_name()';
```

当 `skynet` 查询该表时，返回的是脱敏值而不是原始值。

### 匿名导出与假名化

当前文档推荐通过 masked role 配合 `pg_dump` 进行透明匿名导出。旧的辅助脚本 `pg_dump_anon.sh` 和 `pg_dump_anon` 已被明确标记为 deprecated。

对于导出中的稳定键重映射，文档特别列出了：

- `anon.pseudo_shift(bigint)`
- `anon.pseudo_xor(bigint)`
- `anon.set_shift()`

### 常用函数与注意事项

函数目录中常见的 masking helper 包括：

- `anon.dummy_first_name()`
- `anon.dummy_last_name()`
- `anon.dummy_company_name()`
- `anon.random_zip()`
- `anon.random_date_between(date, date)`
- `anon.partial(value, prefix, mask, suffix)`

官方文档中的注意事项：

- dynamic masking 在 masked-role 会话使用前需要完成 preload 和相关配置
- static masking 会销毁原始值
- pseudonymization 不等于 anonymization
