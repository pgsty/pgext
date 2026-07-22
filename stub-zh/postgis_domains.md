## 用法

来源：

- [官方 README](https://github.com/cahutton/postgis_domains/blob/cb9290194eebb0b63c85c7e59c406e5c0049f48e/README.md)
- [1.0 版扩展 SQL](https://github.com/cahutton/postgis_domains/blob/cb9290194eebb0b63c85c7e59c406e5c0049f48e/postgis_domains--1.0.sql)
- [扩展控制文件](https://github.com/cahutton/postgis_domains/blob/cb9290194eebb0b63c85c7e59c406e5c0049f48e/postgis_domains.control)

`postgis_domains` 为常见的 WGS 84 geometry 与 geography 形状定义可复用的 PostGIS 域。当模式需要在列边界拒绝几何族、维度、SRID、有效性或简单性错误的值，而不希望在每张表中重复检查时，可使用这些域。

### 核心流程

PostGIS 是已声明的依赖，因此安装本扩展还要求 `postgis` 可用。应选择类型与维度后缀符合目标数据的域：

```sql
CREATE EXTENSION postgis_domains;

CREATE TABLE places (
  place_id bigint GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  location point_geometry NOT NULL
);

INSERT INTO places (location)
VALUES (ST_SetSRID(ST_MakePoint(-73.9857, 40.7484), 4326));
```

输入只有满足域检查时才会被接受。因此，无论插入、更新、转换，还是任何使用该域的表，约束都会一致生效。

### 域家族

扩展为点、多点、线串、多线串、多边形、多多边形以及通用几何值定义了 geometry 与 geography 两类对应域。基础二维形式使用 `point_geometry`、`linestring_geometry`、`polygon_geography` 等名称；`m`、`z` 与 `zm` 变体编码所需坐标维度，例如 `pointz_geometry` 与 `polygonzm_geography`。

SQL 定义为每个域组合了相应检查：

- PostGIS 对象家族必须与域一致；
- 坐标维度必须匹配普通、M、Z 或 ZM 变体；
- SRID 必须为 4326；
- 值必须通过扩展定义的有效性与简单性检查。

在生成 DDL 中依赖具体域之前，应检查实际安装的定义：

```sql
SELECT domain_name, data_type
FROM information_schema.domains
WHERE domain_schema = 'public'
  AND domain_name LIKE '%\_geometry' ESCAPE '\\'
ORDER BY domain_name;
```

### 运维说明

1.0 版使用 PostgreSQL 10 与 PostGIS 2.5 开发；README 仅称其应适用于较新的版本。这不是经过验证的兼容矩阵，因此应在生产所用的确切 PostgreSQL/PostGIS 版本组合上测试所选域。

控制文件声明依赖 `postgis`、允许重定位，并且不要求超级用户安装或共享库预加载。域会成为表列类型的一部分，因此移除或替换扩展前必须先处理依赖列。内置域全部强制 WGS 84 SRID 4326；应显式转换源数据，不要仅通过设置新 SRID 来重新标记坐标。
