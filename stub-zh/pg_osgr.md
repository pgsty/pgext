## 用法

来源：

- [已复核 commit 的上游 README](https://github.com/Edwards-R/pg_osgr/blob/253fa94dea4ae2e13ae3ec353c3fa9c42459dc25/readme.md)
- [1.1.0 版本 SQL API](https://github.com/Edwards-R/pg_osgr/blob/253fa94dea4ae2e13ae3ec353c3fa9c42459dc25/pg_osgr--1.1.0.sql)
- [网格参考转换文档](https://github.com/Edwards-R/pg_osgr/blob/253fa94dea4ae2e13ae3ec353c3fa9c42459dc25/docs/function/to_gridref.md)
- [扩展 control 文件](https://github.com/Edwards-R/pg_osgr/blob/253fa94dea4ae2e13ae3ec353c3fa9c42459dc25/pg_osgr.control)

`pg_osgr` 是纯 PL/pgSQL 工具集，用于把 Ordnance Survey 网格参考转换为 easting、northing、datum 与精度，也可从这些分量构造参考值。

```sql
CREATE EXTENSION pg_osgr;
SELECT osgr_find_datum('SU387148');
SELECT osgr_process_easting('SU387148'), osgr_process_northing('SU387148');
SELECT osgr_to_gridref(438700, 114800, 100, 27700);
```

通用函数能够检测或接收英国的 EPSG 27700、爱尔兰与北爱尔兰的 EPSG 29901，以及海峡群岛的 EPSG 32630。分辨率只能是十的幂，且不会低于一米。

该扩展负责解析和格式化网格参考，不执行一般坐标系转换或 datum 变换。存储坐标时应保留 EPSG 代码、验证覆盖边界处的参考值，并在坐标参考系之间转换时使用 PostGIS/PROJ。上游表示规则截至 2023-08-01 正确，受监管数据还需核实后续规范变化。
