## 用法

来源：

- [1.0 版本 SQL 对象](https://github.com/johto/pgbrotli/blob/21c84a7d3dc0777c60db1d6266736e5e8f4e5906/pgbrotli--1.0.sql)
- [C 实现](https://github.com/johto/pgbrotli/blob/21c84a7d3dc0777c60db1d6266736e5e8f4e5906/pgbrotli.c)
- [扩展 control 文件](https://github.com/johto/pgbrotli/blob/21c84a7d3dc0777c60db1d6266736e5e8f4e5906/pgbrotli.control)

`pgbrotli` 1.0 在 PostgreSQL 后端内提供文本 Brotli 压缩。`pgbrotli_compress(text)` 返回压缩后的 `bytea`，`pgbrotli_decompress(bytea)` 以 `bytea` 返回原始字节。

### 核心流程

```sql
CREATE EXTENSION pgbrotli;
WITH packed AS (
  SELECT pgbrotli_compress('hello world'::text) AS value
)
SELECT convert_from(pgbrotli_decompress(value), 'UTF8')
FROM packed;
```

文本压缩器使用 Brotli 的最高质量、默认窗口和文本模式，这些设置不能通过 SQL 参数调整。虽然 1.0 版本声明了 `pgbrotli_compress(bytea)`，其 C 入口始终抛出 `not implemented yet`，因此不支持二进制输入。解压始终返回 `bytea`；只有原始字节符合指定编码时才能使用 `convert_from`。

### 资源与兼容性限制

压缩和解压都在数据库后端中运行。源码把输入或展开值上限设为 1 GiB，但很小的不可信压缩流仍可能触发反复分配并逼近该上限。调用前应拒绝不可信或过大负载，设置语句和资源限制，并避免扫描时反复解压。畸形或截断流会报错。

这个已放弃项目没有当前 PostgreSQL/Brotli 兼容矩阵，除源码外也没有格式契约。应在实际构建上验证空值、库 ABI、架构、畸形流、内存耗尽和并行执行。PostgreSQL TOAST 已可能压缩大值，因此增加应用级 Brotli 前，应测量端到端存储与 CPU。
