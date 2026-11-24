## 用法

来源：

- [PostGIS MEST README](https://github.com/MobilityDB/mest/blob/26503db3d1e6a019094da2316eb2862a8215f053/contrib/postgis/README.md)
- [Control](https://github.com/MobilityDB/mest/blob/26503db3d1e6a019094da2316eb2862a8215f053/contrib/postgis/postgis_mest.control)
- [SQL 1.0](https://github.com/MobilityDB/mest/blob/26503db3d1e6a019094da2316eb2862a8215f053/contrib/postgis/postgis_mest--1.0.sql)

`postgis_mest` 1.0 为 PostGIS `geometry` 提供多条目 R 树索引。几何集合同时产生整体包围盒和各元素的包围盒。上游要求 PostgreSQL 17 或 18、PostGIS 3 及 `mest`。

### 启用与建索引

上游要求预加载已安装的 PostGIS 库。将其加入现有配置并重启 PostgreSQL；文档中的库名为：

```ini
shared_preload_libraries = 'postgis-3'
```

安装扩展库后，使用具有相应权限的数据库会话：

```sql
CREATE EXTENSION postgis_mest CASCADE;
CREATE TABLE regions (id bigint, geom geometry);
CREATE INDEX regions_mgist_geom ON regions USING mgist (geom);
```

控制文件依赖 `postgis` 和 `mest`。`CASCADE` 可以启用已具备安装文件的依赖，但不会安装服务器库文件。

### 查询接口

`geometry_mgist_ops_2d` 是 `mgist` 上针对 `geometry` 的默认操作符类。实用查询操作符包括包围盒重叠 `&&` 和距离排序 `<->`。

```sql
SELECT id
FROM regions
WHERE geom && ST_MakeEnvelope(0, 0, 10, 10, 4326)
  AND ST_Intersects(geom, ST_MakeEnvelope(0, 0, 10, 10, 4326));
```

查询几何对象应与已存数据使用相同的空间参考。

### 限制

该访问方法尚未实现 PostGIS 的规划器支持函数。仅使用 `ST_Intersects` 或 `ST_Contains` 的谓词无法通过该支持机制选中 MGiST 索引；适用时需显式补充重叠条件。收益取决于集合形状和查询选择性，单个几何对象通常只产生一个包围盒。
