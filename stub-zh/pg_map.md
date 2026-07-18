## 用法

来源：

- [已复核 commit 的 pg_map README](https://github.com/semenikhind/pg_map/blob/f1a9c628625e20df148afa92a3df57bb8523bd40/README.md)
- [已复核 commit 的 pg_map 安装 SQL](https://github.com/semenikhind/pg_map/blob/f1a9c628625e20df148afa92a3df57bb8523bd40/pg_map--1.0.sql)
- [已复核 commit 的 pg_map C 实现](https://github.com/semenikhind/pg_map/blob/f1a9c628625e20df148afa92a3df57bb8523bd40/pg_map.c)

`pg_map` 将一个单参数 PostgreSQL 函数应用到数组的每个元素，并返回由函数结果组成的数组。它的两个重载分别用 OID 或文本标识待映射函数；带括号的文本签名按特定 `regprocedure` 解析，裸函数名按 `regproc` 解析。

### 在数组上映射函数

```sql
CREATE EXTENSION pg_map;

SELECT pg_map(
  'upper(text)',
  ARRAY['alpha', 'beta', 'gamma']::text[]
);
```

本例将 `upper(text)` 应用于每个元素，并返回大写文本数组。函数存在重载时，给出完整签名可避免歧义。

### 注意事项

- 上游只声明与 PostgreSQL 9.6devel 兼容，没有说明支持现代 PostgreSQL 版本。
- 1.0 版 C 实现使用 PostgreSQL 内部数组、catalog 和函数管理器 API。使用前应针对确切服务器源码构建并运行回归测试。
- 被映射函数必须接收一个参数。数组元素类型不同于函数参数类型时，实现会尝试做元素类型转换；找不到转换函数时会报错。
- 上游没有发布许可证或发行版兼容矩阵。
