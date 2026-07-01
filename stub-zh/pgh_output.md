


## 用法

来源：[PGHydro README](https://github.com/pghydro/pghydro/blob/master/README.md)，[pgh_output SQL](https://github.com/pghydro/pghydro/blob/master/pgh_output--6.6.sql)

`pgh_output` 为 PgHydro 排水网络产品创建导出表。它把排水线、排水区、水道、排水点和海岸线物化到输出 schema，便于后续报表或 GIS 交换。

### 启用

```sql
CREATE EXTENSION postgis;
CREATE EXTENSION pghydro;
CREATE EXTENSION pgh_output;
```

### 导出表

扩展创建 `pgh_output` schema，其中包括：

| 表 | 用途 |
|----|------|
| `geoft_bho_drainage_line` | 导出的排水线要素 |
| `geoft_bho_drainage_area` | 导出的排水区要素 |
| `geoft_bho_watercourse` | 导出的水道要素 |
| `geoft_bho_drainage_point` | 导出的排水点要素 |
| `geoft_bho_shoreline` | 导出的海岸线要素 |

### 刷新输出

```sql
SELECT pgh_output.pghfn_updateexporttables();
```

加载或重新计算基础 PgHydro 网络后，运行刷新函数更新导出表。

### 注意事项

条件允许时，先校验 PgHydro 网络一致性。导出表是派生产品，上游排水几何、拓扑或编码变化后都应重新生成。
