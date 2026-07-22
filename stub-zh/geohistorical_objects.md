## 用法

来源：

- [官方 README](https://github.com/GeoHistoricalData/geohistorical_objects/blob/5e7ec83dc568a509a88a0a66e3add8310bba13b3/README.md)
- [1.0 版本控制文件](https://github.com/GeoHistoricalData/geohistorical_objects/blob/5e7ec83dc568a509a88a0a66e3add8310bba13b3/geohistorical_objects.control)
- [1.0 版本 SQL 实现](https://github.com/GeoHistoricalData/geohistorical_objects/blob/5e7ec83dc568a509a88a0a66e3add8310bba13b3/geohistorical_objects.sql)

`geohistorical_objects` 提供基于继承的数据模型，用于表示名称、几何、来源、日期与精度随历史变化的地理对象。它在固定的 `geohistorical_object` 模式中创建对象，并依赖 `postgis`、`pgsfti`、`unaccent` 与 `pg_trgm`。

### 核心模型

```sql
CREATE EXTENSION postgis;
CREATE EXTENSION pgsfti;
CREATE EXTENSION unaccent;
CREATE EXTENSION pg_trgm;
CREATE EXTENSION geohistorical_objects;

CREATE TABLE public.historic_places (
  place_id bigint GENERATED ALWAYS AS IDENTITY PRIMARY KEY
) INHERITS (geohistorical_object.geohistorical_object);

SELECT geohistorical_object.register_geohistorical_object_table(
  'public', 'historic_places'
);
```

应用应先填充 `historical_source` 与 `numerical_origin_process`，其中包括 `sfti` 类型的 `default_fuzzy_date`，以及包含非负 `default` 键的 `default_spatial_precision` JSON 对象。数据行应写入子表，而不是父表 `geohistorical_object`。注册辅助函数会为子表添加外键及 B-tree、GIN、GiST 索引。

扩展还提供 SFTI 转换函数与类型转换、关系表、名称规范化辅助函数以及模糊日期距离辅助函数。其设计使用 PostgreSQL 表继承，而不是声明式分区；唯一性与外键行为不会自动覆盖所有子表。

### 安装风险

1.0 版本安装脚本包含无条件的 `DROP TABLE ... CASCADE`，会创建名称范围较广的类型转换与辅助函数，并嵌入测试式语句。安装或重新安装可能删除目标模式中的已有对象。注册函数还会根据文本标识符拼接动态 DDL，但没有做标识符引用。

应完整审阅 SQL 文件，并在一次性数据库中演练安装、注册、转储恢复与移除。用于现有数据前必须准备经过验证的备份，只向注册函数传入可信的简单标识符，并在实际 PostgreSQL、PostGIS 与 pgsfti 版本上验证上游旧兼容性声明。
