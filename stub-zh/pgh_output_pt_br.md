


## 用法

来源：[PGHydro README](https://github.com/pghydro/pghydro/blob/master/README.md)，[pgh_output_pt_br SQL](https://github.com/pghydro/pghydro/blob/master/pgh_output_pt_br--6.6.sql)

`pgh_output_pt_br` 提供巴西葡萄牙语命名的 PgHydro 输出 schema。它创建本地化的排水网络产品导出表，并提供从基础 PgHydro 模型刷新这些表的函数。

### 启用

```sql
CREATE EXTENSION postgis;
CREATE EXTENSION pghydro;
CREATE EXTENSION pgh_output_pt_br;
```

### 输出表

扩展创建 `pgh_output_pt_br` schema，其中包括：

| 表 | 用途 |
|----|------|
| `geoft_bho_trecho_drenagem` | 排水线导出表 |
| `geoft_bho_area_drenagem` | 排水区导出表 |
| `geoft_bho_curso_dagua` | 水道导出表 |
| `geoft_bho_ponto_drenagem` | 排水点导出表 |
| `geoft_bho_linha_costa` | 海岸线导出表 |

### 刷新输出

```sql
SELECT pgh_output_pt_br.pghfn_updateexporttables();
```

### 注意事项

当导出的 PgHydro 产品需要使用巴西葡萄牙语对象名时，可以使用该扩展。源排水几何、拓扑或编码变化后，应重新生成输出表。
