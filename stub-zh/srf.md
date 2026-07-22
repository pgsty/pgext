## 用法

来源：

- [1.0 版本扩展 SQL](https://github.com/zombodb/srf/blob/350fb8962cf5fdd06d4bdca1ed1a01a7a798a498/srf--1.0.sql)
- [C 实现](https://github.com/zombodb/srf/blob/350fb8962cf5fdd06d4bdca1ed1a01a7a798a498/srf.c)
- [扩展控制文件](https://github.com/zombodb/srf/blob/350fb8962cf5fdd06d4bdca1ed1a01a7a798a498/srf.control)

`srf` 是一个已废弃的演示扩展，只包含一个 C 集合返回函数。`srf_c` 生成包含端点的 32 位整数序列，仅适合用于研究 PostgreSQL 的多次调用 C 函数 API。

### 核心流程

```sql
CREATE EXTENSION srf;

SELECT * FROM srf_c(3, 7);
```

查询返回 `3`、`4`、`5`、`6` 和 `7`。当 `start` 大于 `end` 时，不返回任何行。

### 对象索引

- `srf_c(start integer, "end" integer) RETURNS SETOF integer` 生成从 `start` 到 `end` 的每个整数，包含两端。

### 运维说明

版本 `1.0` 可重定位，且未声明依赖或预加载要求。SQL 函数没有易变性、严格性或并行安全标注，因此使用 PostgreSQL 默认值。

仓库没有用户文档或发行支持，目录将其标记为已废弃。不要把它作为生产环境中 `generate_series` 的替代品。应避免把 `end` 设为 `2147483647`：C 实现在返回端点后会递增有符号 32 位计数器，但没有溢出保护。应用查询应优先使用 PostgreSQL 内置的 `generate_series(integer, integer)`。
