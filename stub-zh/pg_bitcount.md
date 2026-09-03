## 用法

来源：

- [官方文档](https://github.com/thehyve/pg_bitcount/blob/77d4fa8e26dd46166a9d906920e8b341fb8fb643/README.md)
- [扩展控制文件](https://github.com/thehyve/pg_bitcount/blob/77d4fa8e26dd46166a9d906920e8b341fb8fb643/pg_bitcount.control)
- [官方仓库](https://github.com/thehyve/pg_bitcount)

`pg_bitcount` 面向整数与位串的位计数标量和聚合函数。

### 启用

为目标服务器安装文件，然后在目标数据库中创建 `pg_bitcount`：

```sql
CREATE EXTENSION pg_bitcount;
```

### 核心流程

以下示例取自经审查的上游文档。使用前应调整对象名、路径、凭据与工作负载参数。

```sql
-- Register the extension in PostgreSQL
create extension pg_bitcount version '0.0.3';

-- Use the pg_bitcount function
select public.pg_bitcount(127::bit(8)); -- 7
select public.pg_bitcount(B'101010101'); -- 5
select public.pg_bitcount((17^15)::bigint::bit(128) << 64 | (17^14)::bigint::bit(128)); -- 58

-- Use the pg_int_to_bit_agg aggregate using a bit string of size 24
select public.pg_bitcount(public.pg_int_to_bit_agg(i::int, 24)) from (select generate_series(2, 8) as i) data; -- 7
```

### 主要对象

经审查的安装接口包含以下主要对象：

| 对象 | 类型 | 运维作用 |
|---|---|---|
| `public.pg_bitcount` | FUNCTION | 经审查安装接口中的可调用函数。 |
| `public.pg_int_to_bit_agg` | AGGREGATE | 扩展提供的聚合函数。 |
| `public.pg_int_to_bit_agg_transfn` | FUNCTION | 经审查安装接口中的可调用函数。 |

### 运维与边界

- 已验证的 PostgreSQL 主版本证据覆盖 11, 12, 13；不要推断未列出的主版本。
- 目录生命周期为 archived；生产使用前应测试升级、备份恢复与服务器兼容性。
