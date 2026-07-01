


## 用法

来源：[PGHydro README](https://github.com/pghydro/pghydro/blob/master/README.md)，[pgh_output_en_au SQL](https://github.com/pghydro/pghydro/blob/master/pgh_output_en-au--6.6.sql)

`pgh_output_en_au` 提供澳大利亚英语命名的 PgHydro 输出 schema。它创建本地化的排水线、排水区和排水点导出表，并提供从基础 PgHydro 网络刷新这些表的函数。

### 启用

```sql
CREATE EXTENSION postgis;
CREATE EXTENSION pghydro;
CREATE EXTENSION pgh_output_en_au;
```

### 输出表

扩展创建 `pgh_output_en_au` schema，其中包括：

| 表 | 用途 |
|----|------|
| `geoft_bho_drainage_line` | 排水线导出表 |
| `geoft_bho_drainage_area` | 排水区导出表 |
| `geoft_bho_drainage_point` | 排水点导出表 |

### 刷新输出

```sql
SELECT pgh_output_en_au.pghfn_updateexporttables();
```

重新计算或编辑源 PgHydro 对象后，运行刷新函数更新输出表。

### 注意事项

当导出 schema 需要使用澳大利亚英语命名约定，同时数据仍然来自 PgHydro 排水网络时，可以使用该扩展。
