## 用法

来源：

- [上游 README 与 API 示例](https://github.com/yuiseki/pg_s2/blob/0cb0f78877223dcb4c2429a29d4d9edb3e106c83/README.md)
- [扩展 control 文件](https://github.com/yuiseki/pg_s2/blob/0cb0f78877223dcb4c2429a29d4d9edb3e106c83/pg_s2.control)
- [Cargo 包元数据](https://github.com/yuiseki/pg_s2/blob/0cb0f78877223dcb4c2429a29d4d9edb3e106c83/Cargo.toml)
- [Rust SQL 实现](https://github.com/yuiseki/pg_s2/blob/0cb0f78877223dcb4c2429a29d4d9edb3e106c83/src/lib.rs)
- [API 规范](https://github.com/yuiseki/pg_s2/blob/0cb0f78877223dcb4c2429a29d4d9edb3e106c83/SPEC.md)

`pg_s2` 使用 PostgreSQL 内置空间类型提供紧凑的 S2 CellID API，无需 PostGIS。其保持顺序的 `s2cellid` 类型支持 B-tree 索引、token 与 bigint 转换、层级和相邻单元遍历、圆形/矩形覆盖、边界以及大圆距离。

### 创建并检查单元

PostgreSQL `point` 值以经度为 X、纬度为 Y：

```sql
CREATE EXTENSION pg_s2;

SELECT s2_lat_lng_to_cell(point(139.767, 35.681), 14);
SELECT s2_cell_to_token(
  s2_lat_lng_to_cell(point(139.767, 35.681), 14)
);
SELECT s2_cell_to_parent(
  s2_lat_lng_to_cell(point(139.767, 35.681), 14),
  12
);
SELECT *
FROM s2_cover_cap_ranges(point(139.767, 35.681), 2000.0, 12, 16);
```

进行索引空间搜索时，应存储单元格，使用 `s2_cover_cap_ranges()` 或 `s2_cover_rect_ranges()` 做粗粒度 B-tree 预过滤，再应用 `s2_great_circle_distance()` 等精确谓词。覆盖集有意采用过度近似，可能包含假阳性。用户可设置的默认值包括 `pg_s2.default_level`、`pg_s2.default_cover_level` 和 `pg_s2.earth_radius_m`。

版本 `0.0.6` 被上游描述为早期 MVP/测试版本。README 声明支持 PostgreSQL 14 至 17，而 Cargo 还暴露了 13 与 18 的构建 feature；额外 feature 并不等于支持承诺。无效 token、坐标或层级会报错，应用内也必须始终保持经纬度顺序一致。
