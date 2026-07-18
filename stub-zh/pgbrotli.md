## 用法

来源：

- [1.0 版本 SQL 对象](https://github.com/johto/pgbrotli/blob/21c84a7d3dc0777c60db1d6266736e5e8f4e5906/pgbrotli--1.0.sql)
- [C 实现](https://github.com/johto/pgbrotli/blob/21c84a7d3dc0777c60db1d6266736e5e8f4e5906/pgbrotli.c)
- [扩展 control 文件](https://github.com/johto/pgbrotli/blob/21c84a7d3dc0777c60db1d6266736e5e8f4e5906/pgbrotli.control)

`pgbrotli` 在 PostgreSQL 内提供 Brotli 压缩。`pgbrotli_compress(text)` 与 `pgbrotli_compress(bytea)` 返回压缩后的 `bytea`；`pgbrotli_decompress(bytea)` 返回原始字节。

```sql
CREATE EXTENSION pgbrotli;
WITH packed AS (
  SELECT pgbrotli_compress('hello world'::text) AS value
)
SELECT convert_from(pgbrotli_decompress(value), 'UTF8')
FROM packed;
```

解压在数据库后端内进行，小输入可能膨胀成很大输出。调用前应拒绝不可信或过大负载，设置语句与资源限制，并避免扫描时反复解压。应明确处理 `bytea` 结果；字节不符合指定编码时，`convert_from` 会失败。

这个已放弃的 1.0 项目没有当前 PostgreSQL/Brotli 兼容矩阵，除源码与测试外也没有格式契约。应验证畸形输入、截断流、膨胀上限、空值、架构、库 ABI、内存耗尽和并行执行。PostgreSQL TOAST 已可能压缩大值，因此增加应用级 Brotli 前，应测量端到端存储与 CPU。
