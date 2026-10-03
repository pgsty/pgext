## 用法

来源：

- [pg_dynamic.control](https://github.com/JoshInnis/pg_dynamic/blob/921415dec1750ce48b8011a1cead46a8820cb7a6/pg_dynamic.control)
- [README.md](https://github.com/JoshInnis/pg_dynamic/blob/921415dec1750ce48b8011a1cead46a8820cb7a6/README.md)
- [pg_dynamic--0.1.0.sql](https://github.com/JoshInnis/pg_dynamic/blob/921415dec1750ce48b8011a1cead46a8820cb7a6/pg_dynamic--0.1.0.sql)
- [Makefile](https://github.com/JoshInnis/pg_dynamic/blob/921415dec1750ce48b8011a1cead46a8820cb7a6/Makefile)
- [regress/sql/integer.sql](https://github.com/JoshInnis/pg_dynamic/blob/921415dec1750ce48b8011a1cead46a8820cb7a6/regress/sql/integer.sql)

`pg_dynamic` 增加 `dynamic` 类型，用于承载部分 PostgreSQL 类型并提供类型转换及重载运算。0.1.0 仍是实验实现，其支持所有类型的项目目标超出了现有 SQL 接口。

### 核心用法

```sql
CREATE EXTENSION pg_dynamic;
SELECT (42::bigint)::dynamic;
SELECT ((42::bigint)::dynamic)::bigint;
```

### 运行边界

需要超级用户安装，共享库不要求服务器预加载。所核对 SQL 包含 bigint、inet、box 转换，以及算术、几何和部分数学函数。应仅使用该修订实际导出的转换与函数，不能假定已支持任意扩展类型或通用自动分派。所核对文档未明确 PostgreSQL 主版本支持范围。
