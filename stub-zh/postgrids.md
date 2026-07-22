## 用法

来源：

- [官方 README](https://github.com/bis-brecon/postgrids/blob/ba0a85a344cfefc8e630c7428e930e6eb8df03aa/README.md)
- [OSGB SQL 实现](https://github.com/bis-brecon/postgrids/blob/ba0a85a344cfefc8e630c7428e930e6eb8df03aa/src/osgb.rs)
- [OSI SQL 实现](https://github.com/bis-brecon/postgrids/blob/ba0a85a344cfefc8e630c7428e930e6eb8df03aa/src/osi.rs)

`postgrids` 版本 `0.0.0` 为英国 OSGB 与爱尔兰 OSI 国家格网参考提供 PostgreSQL 类型和辅助函数。它可以解析、校验、格式化参考值并改变精度，但不会把参考值转换成 PostGIS 几何对象或其他坐标系。

### 核心流程

```sql
CREATE EXTENSION postgrids;

SELECT osgb_from_string('SO892437');
SELECT osgb_from_eastings_northings(389200, 243700, 100);
SELECT osgb_is_valid('SO892437');
SELECT osgb_precision(osgb_from_string('SO892437'));
SELECT osgb_recalculate(osgb_from_string('SO892437'), 1000);
```

对应的爱尔兰函数使用 `osi_` 前缀：

```sql
SELECT osi_from_string('O892437');
SELECT osi_is_valid('O892437');
SELECT osi_recalculate(osi_from_string('O892437'), 1000);
```

### 对象

自定义类型是 `OSGB` 与 `OSI`。每一族都提供 `*_from_string`、`*_from_eastings_northings`、`*_is_valid`、`*_precision` 和 `*_recalculate`。

精度以米表示。支持值为 `1`、`10`、`100`、`1000`、`2000`、`10000` 和 `100000`；不支持的值会报错。解析构造函数遇到无效参考值也会报错，而 `*_is_valid` 辅助函数返回布尔值。

### 成熟度边界

上游 README 把项目描述为非常早期的开发中状态，并明确说明尚不支持转换为地理空间基本对象。持久存储这些自定义类型前，应固定扩展构建并验证可接受的格网参考格式。
