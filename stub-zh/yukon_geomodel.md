## 用法

来源：

- [geomodel/src/yukon_geomodel.control](https://github.com/opengauss-mirror/Yukon/blob/7c1063e45f136dca9d456e5bb56fafd4ea973ef2/geomodel/src/yukon_geomodel.control)
- [README_EN.md](https://github.com/opengauss-mirror/Yukon/blob/7c1063e45f136dca9d456e5bb56fafd4ea973ef2/README_EN.md)
- [geomodel/src/yukon_geomodel--1.0.1.sql](https://github.com/opengauss-mirror/Yukon/blob/7c1063e45f136dca9d456e5bb56fafd4ea973ef2/geomodel/src/yukon_geomodel--1.0.1.sql)
- [doc/yukon_doc/source/installation.rst](https://github.com/opengauss-mirror/Yukon/blob/7c1063e45f136dca9d456e5bb56fafd4ea973ef2/doc/yukon_doc/source/installation.rst)
- [doc/yukon_doc/source/modular_api/geomodel_api.rst](https://github.com/opengauss-mirror/Yukon/blob/7c1063e45f136dca9d456e5bb56fafd4ea973ef2/doc/yukon_doc/source/modular_api/geomodel_api.rst)

`yukon_geomodel` 是面向 openGauss 内核的 Yukon 模块。提供 GEOMODEL 和 MODEL_ELEM 数据，应通过管理函数增删模型列及关联子表；上游明确要求不要直接用 DDL 管理这些结构。

### 核心用法

```sql
CREATE EXTENSION postgis;
CREATE EXTENSION yukon_geomodel;
CREATE TABLE model_sample (id integer);
SELECT AddGeoModelColumn('public', 'model_sample', 'model', 4326);
```

### 运行边界

需要兼容的 Yukon／openGauss 环境，包括适配后的 PostGIS 类型和库。应由具有足够权限的数据库管理员启用，没有预加载要求。此条目不表示兼容原生 PostgreSQL。几何／SRID 约定、模型与网格编码遵循 Yukon API；应使用匹配的客户端，并在备份和维护时保留相关元数据。
