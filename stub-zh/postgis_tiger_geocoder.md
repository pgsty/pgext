

## 用法

> [PostGIS TIGER Geocoder：基于美国人口普查 TIGER/Line 数据的 PostGIS 地理编码](https://github.com/postgis/postgis)

PostGIS TIGER Geocoder 利用美国人口普查 TIGER/Line 数据提供地理编码和反向地理编码功能。它可以将地址字符串解析为标准化格式、查找地理坐标，以及将坐标反向解析为地址。

- [TIGER Geocoder 参考手册](https://postgis.net/docs/Extras.html)

### 安装

```sql
CREATE EXTENSION postgis_tiger_geocoder CASCADE;
```

这将创建包含地理编码器表和函数的 `tiger` 模式。

--------

## 加载 TIGER 数据

在进行地理编码之前，必须先加载所需州的 TIGER/Line 数据。该扩展提供了辅助函数来生成加载脚本：

```sql
-- 生成下载和加载某个州数据的脚本
-- （例如，马萨诸塞州 = 'MA'）
SELECT loader_generate_script(ARRAY['MA'], 'sh');
```

这会生成一个使用 `shp2pgsql` 加载 TIGER 形状文件的 shell 脚本。运行生成的脚本可将地址范围、边、面等数据填充到 `tiger_data` 模式中。

加载完成后：

```sql
-- 安装缺失的索引以提升性能
SELECT install_missing_indexes();

-- 更新统计信息
ANALYZE tiger.addr;
ANALYZE tiger.edges;
ANALYZE tiger.faces;
```

--------

## 地理编码

将美国地址字符串转换为地理坐标：

```sql
-- 基本地理编码
SELECT g.rating, ST_X(g.geomout) AS lon, ST_Y(g.geomout) AS lat,
       pprint_addy(g.addy) AS address
FROM geocode('1600 Pennsylvania Ave NW, Washington, DC 20500') AS g;
```

`rating` 表示匹配质量（越低越好，0 = 精确匹配）。

```sql
-- 限制返回结果数量的地理编码
SELECT g.rating, ST_AsText(g.geomout), pprint_addy(g.addy)
FROM geocode('100 Main St, Boston, MA', 3) AS g;

-- 从表中批量地理编码
SELECT a.id, g.rating, g.geomout, pprint_addy(g.addy)
FROM addresses a, LATERAL geocode(a.address_string, 1) AS g;
```

--------

## 反向地理编码

将坐标转换回街道地址：

```sql
SELECT pprint_addy(r.addy[1]) AS address
FROM reverse_geocode(ST_SetSRID(ST_MakePoint(-77.0365, 38.8977), 4326)) AS r;
```

--------

## 地址标准化

无需地理编码即可解析和标准化地址字符串：

```sql
SELECT *
FROM normalize_address('1600 Pennsylvania Avenue NW, Washington, DC 20500');
```

返回的组件包括：`address`（门牌号）、`predirAbbrev`、`streetName`、`streetTypeAbbrev`、`postdirAbbrev`、`internal`、`location`（城市）、`stateAbbrev`、`zip`、`parsed`。

```sql
-- 格式化输出标准化地址
SELECT pprint_addy(normalize_address('100 main street boston ma 02101'));
```

--------

## 配置

`tiger.geocode_settings` 表控制地理编码器的行为：

```sql
-- 查看当前设置
SELECT * FROM tiger.geocode_settings;

-- 调整设置（例如，增加调试级别）
UPDATE tiger.geocode_settings SET val = 'true' WHERE name = 'debug_geocode_address';
```
