## 用法

来源：

- [geogridcoder/yukon_geogridcoder.control](https://github.com/opengauss-mirror/Yukon/blob/7c1063e45f136dca9d456e5bb56fafd4ea973ef2/geogridcoder/yukon_geogridcoder.control)
- [README_EN.md](https://github.com/opengauss-mirror/Yukon/blob/7c1063e45f136dca9d456e5bb56fafd4ea973ef2/README_EN.md)
- [geogridcoder/yukon_geogridcoder--1.0.1.sql](https://github.com/opengauss-mirror/Yukon/blob/7c1063e45f136dca9d456e5bb56fafd4ea973ef2/geogridcoder/yukon_geogridcoder--1.0.1.sql)
- [doc/yukon_doc/source/installation.rst](https://github.com/opengauss-mirror/Yukon/blob/7c1063e45f136dca9d456e5bb56fafd4ea973ef2/doc/yukon_doc/source/installation.rst)

`yukon_geogridcoder` 是面向 openGauss 内核的 Yukon 模块。提供 geosotgrid 类型、标量比较、网格数组运算，以及几何对象和 GeoSOT 网格之间的转换，包含二维与高度相关接口。

### 核心用法

```sql
CREATE EXTENSION postgis;
CREATE EXTENSION yukon_geogridcoder;
SELECT ST_GeoSOTGrid(ST_GeomFromText('POINT(116.4 39.9)', 4326), 15);
```

### 运行边界

需要兼容的 Yukon／openGauss 环境，包括适配后的 PostGIS 类型和库。应由具有足够权限的数据库管理员启用，没有预加载要求。此条目不表示兼容原生 PostgreSQL。几何／SRID 约定、模型与网格编码遵循 Yukon API；应使用匹配的客户端，并在备份和维护时保留相关元数据。
