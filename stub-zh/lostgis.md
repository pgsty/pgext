## 用法

来源：

- [官方扩展文档](https://github.com/gojuno/lostgis/blob/85f26871adef0eb5f8fed43704a9d28b2fb5b80b/doc/lostgis.md)
- [官方扩展 SQL 目录](https://github.com/gojuno/lostgis/tree/85f26871adef0eb5f8fed43704a9d28b2fb5b80b/sql)
- [官方 PGXN 版本](https://pgxn.org/dist/lostgis/1.0.2/)

`lostgis` 是一组 PostGIS 工具，覆盖位置/时间/速度轨迹、OpenStreetMap 风格营业时间、几何修复与叠加、Web Mercator 实际距离辅助计算、网格单元和地图泛化。版本 `1.0.2` 是 2017 年发布的纯 SQL 版本；应在验证各函数假设后按需使用，而不能把它视为完整 GIS 模型。

### 核心流程

尽管控制文件没有声明依赖，使用前仍必须已有 PostGIS 类型：

```sql
CREATE EXTENSION postgis;
CREATE EXTENSION lostgis;

SELECT ST_GridCell(
  ST_SetSRID(ST_Point(4180000, 7500000), 3857),
  500
);

SELECT overlaps(
  timestamp '2017-08-13 13:00',
  'Mo-Fr 05:00-15:00; Sa 05:00-12:00'::text::opening_hours
);
```

`overlaps()` 输入是普通 `timestamp`：检查营业时间前，应显式转换成业务地点的时区。它的解析器只支持简单的星期与时/分范围。

### 对象索引

- `tpv` 保存 Web Mercator 点、精度、方向、速度、时间戳、来源和 OSM 标识符。
- `opening_hours` 保存解析后的可读文本、有效性标志与每周分钟位图。
- `project_position()` 和 `ST_AddTime()` 处理移动位置与轨迹。
- `ST_Fast_Real_Buffer()`、`ST_Fast_Real_Length()` 和 `ST_RealOffsetCurve()` 为投影几何近似执行以米为单位的操作。
- `ST_Safe_Repair()`、`ST_Safe_Difference()` 和 `ST_Safe_Intersection()` 尝试绕过无效几何进行恢复。
- `ST_FilterSmallRings()`、`ST_LargestSubPolygon()`、`ST_LineAngleAtPoint()`、`ST_TimeLineMerge()`、`ST_AnglesEqual()`、`coslat()` 和 `median()` 提供其他辅助能力。

### 注意事项

多数函数假定特定几何类型或 SRID `3857`；使用前应确认单位、SRID、维度、空值/NULL 行为和有效性。“安全”叠加函数可能修复或对齐数据，因此会改变拓扑，而不保证精确保留输入。近似实际距离辅助函数不能替代精度敏感场景中的 PostGIS geography。应测试典型的极区/高纬数据、反子午线情况、无效多边形、涉及 DST 的营业时间、大数组，以及与目标 PostGIS 版本的兼容性。
