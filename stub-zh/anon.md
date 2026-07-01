## 用法

> 来源：[overview](https://postgresql-anonymizer.readthedocs.io/en/stable/)、[static masking](https://postgresql-anonymizer.readthedocs.io/en/stable/static_masking/)、[dynamic masking](https://postgresql-anonymizer.readthedocs.io/en/stable/dynamic_masking/)、[anonymous dumps](https://postgresql-anonymizer.readthedocs.io/en/stable/anonymous_dumps/)、[masking functions](https://postgresql-anonymizer.readthedocs.io/en/stable/masking_functions/)、[release 3.1.1](https://gitlab.com/dalibo/postgresql_anonymizer/-/releases/3.1.1)

`anon` 使用 `SECURITY LABEL FOR anon` 声明脱敏规则。官方文档描述了六种脱敏方法：anonymous dumps、static masking、dynamic masking、replica masking、masking views 和 masking data wrappers。

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

静态脱敏会原地改写已存储的数据：

```sql
SELECT anon.anonymize_database();
-- See also: anon.anonymize_table(), anon.anonymize_column()
-- For larger databases: anon.anonymize_database_parallel(worker_count)
```

静态脱敏文档还覆盖 shuffling、噪声注入，以及面向较大数据集的并行脱敏。并行静态脱敏受 `anon.max_bg_workers` 和服务器 `max_worker_processes` 限制。

### 动态脱敏

动态脱敏只对标记为 masked 的角色隐藏值：

```sql
ALTER DATABASE demo SET session_preload_libraries = 'anon';
ALTER DATABASE demo SET anon.transparent_dynamic_masking TO true;

CREATE ROLE skynet LOGIN;
SECURITY LABEL FOR anon ON ROLE skynet IS 'MASKED';
GRANT pg_read_all_data TO skynet;

SECURITY LABEL FOR anon ON COLUMN people.lastname
IS 'MASKED WITH FUNCTION anon.dummy_last_name()';
```

当 `skynet` 查询该表时，会返回脱敏值而不是原始值。

### 匿名导出与伪名化

当前文档推荐通过 masked role 和 `pg_dump` 做透明匿名导出。较早的辅助工具 `pg_dump_anon.sh` 和 `pg_dump_anon` 已明确标记为 deprecated。

对于 PostgreSQL 17 及之后版本，导出示例使用 `--exclude-extension="anon"` 搭配 `--no-security-labels`；较旧的 `pg_dump` 版本需要另一种 extension 选择方式，例如 `--extension plpgsql`。

对于导出中的稳定键重映射，文档列出了：

- `anon.pseudo_shift(bigint)`
- `anon.pseudo_xor(bigint)`
- `anon.set_shift()`

### 常用函数与注意事项

函数目录中的常用脱敏辅助函数包括：

- `anon.dummy_first_name()`
- `anon.dummy_last_name()`
- `anon.dummy_company_name()`
- `anon.random_zip()`
- `anon.random_date_between(date, date)`
- `anon.partial(value, prefix, mask, suffix)`

官方文档中的注意事项：

- dynamic masking 需要在 masked role 会话使用前完成 preload/configuration
- static masking 会破坏原始值
- pseudonymization 不等同于 anonymization
- 版本 3.1.1 是针对 CVE-2026-11945 的 security release，并从 import/export 行为中移除 conditional randomization；应尽快从 3.0.x 升级。
