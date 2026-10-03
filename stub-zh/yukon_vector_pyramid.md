## 用法

来源：

- [pyramid/yukon_vector_pyramid.control](https://github.com/opengauss-mirror/Yukon/blob/7c1063e45f136dca9d456e5bb56fafd4ea973ef2/pyramid/yukon_vector_pyramid.control)
- [README_EN.md](https://github.com/opengauss-mirror/Yukon/blob/7c1063e45f136dca9d456e5bb56fafd4ea973ef2/README_EN.md)
- [pyramid/yukon_vector_pyramid--1.0.sql](https://github.com/opengauss-mirror/Yukon/blob/7c1063e45f136dca9d456e5bb56fafd4ea973ef2/pyramid/yukon_vector_pyramid--1.0.sql)
- [doc/yukon_doc/source/installation.rst](https://github.com/opengauss-mirror/Yukon/blob/7c1063e45f136dca9d456e5bb56fafd4ea973ef2/doc/yukon_doc/source/installation.rst)

`yukon_vector_pyramid` 是面向 openGauss 内核的 Yukon 模块。SQL 接口包含 ST_BuildPyramid、ST_ListPyramid、ST_HasPyramid、ST_DeletePyramid、ST_BuildTile、ST_UpdatePyramid 与 ST_AsTile。金字塔构建会产生派生数据，需要随源数据维护。

### 核心用法

```sql
CREATE EXTENSION postgis;
CREATE EXTENSION yukon_vector_pyramid;
SELECT yukon_pyramid_version();
```

### 运行边界

需要兼容的 Yukon／openGauss 环境，包括适配后的 PostGIS 类型和库。应由具有足够权限的数据库管理员启用，没有预加载要求。此条目不表示兼容原生 PostgreSQL。几何／SRID 约定、模型与网格编码遵循 Yukon API；应使用匹配的客户端，并在备份和维护时保留相关元数据。
