


## 用法

> 来源：[概览](https://postgresql-anonymizer.readthedocs.io/en/stable/)、[静态脱敏](https://postgresql-anonymizer.readthedocs.io/en/stable/static_masking/)、[动态脱敏](https://postgresql-anonymizer.readthedocs.io/en/stable/dynamic_masking/)、[匿名导出](https://postgresql-anonymizer.readthedocs.io/en/stable/anonymous_dumps/)、[脱敏函数](https://postgresql-anonymizer.readthedocs.io/en/stable/masking_functions/)、[3.1.1 发行版](https://gitlab.com/dalibo/postgresql_anonymizer/-/releases/3.1.1)

`anon` 使用 `SECURITY LABEL FOR anon` 声明脱敏规则。官方文档描述了六种脱敏方法：匿名导出、静态脱敏、动态脱敏、副本脱敏、脱敏视图和脱敏数据包装器。

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

静态脱敏文档还覆盖随机重排、噪声注入，以及面向较大数据集的并行脱敏。并行静态脱敏受 `anon.max_bg_workers` 和服务器 `max_worker_processes` 限制。

### 动态脱敏

动态脱敏只对标记为已脱敏的角色隐藏值：

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

当前文档推荐通过脱敏角色和 `pg_dump` 做透明匿名导出。较早的辅助工具 `pg_dump_anon.sh` 和 `pg_dump_anon` 已明确标记为弃用。

对于 PostgreSQL 17 及之后版本，导出示例使用 `--exclude-extension="anon"` 搭配 `--no-security-labels`；较旧的 `pg_dump` 版本需要另一种扩展选择方式，例如 `--extension plpgsql`。

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

- 动态脱敏需要在脱敏角色会话使用前完成预加载和配置
- 静态脱敏会破坏原始值
- 假名化不等同于匿名化
- 版本 3.1.1 是针对 CVE-2026-11945 的安全发行版，并从导入/导出行为中移除条件随机化；应尽快从 3.0.x 升级。
