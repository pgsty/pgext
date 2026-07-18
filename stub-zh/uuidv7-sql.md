## 用法

来源：

- [上游 README](https://github.com/dverite/postgres-uuidv7-sql/blob/396a44433e6e0eb63b1d9d1517e9098256d97351/README.md)
- [1.0 版生成函数](https://github.com/dverite/postgres-uuidv7-sql/blob/396a44433e6e0eb63b1d9d1517e9098256d97351/sql/uuidv7-sql--1.0.sql)
- [1.1 版加密函数](https://github.com/dverite/postgres-uuidv7-sql/blob/396a44433e6e0eb63b1d9d1517e9098256d97351/sql/uuidv7-sql--1.0--1.1.sql)
- [1.1 版控制文件](https://github.com/dverite/postgres-uuidv7-sql/blob/396a44433e6e0eb63b1d9d1517e9098256d97351/uuidv7-sql.control)
- [PostgreSQL 18 UUID 函数](https://www.postgresql.org/docs/18/functions-uuid.html)

uuidv7-sql 1.1 是 UUIDv7 生成、时间戳提取、范围边界、亚毫秒变体，以及可逆 UUIDv4 或 UUIDv8 转换的纯 SQL 实现。

### 生成与检查

应安装到显式模式，以便 PostgreSQL 18 及以上版本中的调用保持无歧义：

```sql
CREATE EXTENSION "uuidv7-sql" WITH SCHEMA public;
SELECT public.uuidv7() AS id;
SELECT public.uuidv7_extract_timestamp(public.uuidv7());
SELECT public.uuidv7_boundary('2026-01-01 00:00:00+00');
```

PostgreSQL 18 在 pg_catalog 中内置 uuidv7 函数，普通名称解析时会优先匹配。常规生成应优先使用内置函数，或者像示例一样以模式限定方式调用本扩展版本。

### 注意事项

- UUIDv7 会暴露创建时间戳。不要在近似事件时间敏感的场景中公开原始值。
- 同一毫秒内生成的值包含随机排序位，不保证严格递增。亚毫秒变体也不提供全局单调序列。
- uuidv7_extract_timestamp 不验证输入版本或变体；它会把任何 UUID 的前 48 位解释为毫秒数。
- uuidv7_boundary 是确定性的非随机值。它只能用作范围边界，绝不能作为行标识符。
- 1.1 版 XTEA 转换是可逆、无认证的 64 位分组结构，不是通用认证加密。它不能检测篡改，没有定义密钥存储或轮换，也不验证输入是否 UUIDv7；通过 SQL 提供的密钥还可能出现在日志和活动视图中。
- 生成依赖 gen_random_uuid，该函数从 PostgreSQL 13 起内置；更早版本需要 pgcrypto，但控制文件没有声明该依赖。
- 仓库没有自动化测试。依赖前必须验证 RFC 一致性、时间戳边界、时区、支持日期范围、加密往返与从 1.0 升级。
