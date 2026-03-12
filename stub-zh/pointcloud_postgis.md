

## 用法

> [pointcloud_postgis: pgPointcloud 的 PostGIS 集成](https://github.com/pgpointcloud/pointcloud)

`pointcloud_postgis` 是将 pgPointcloud 扩展与 PostGIS 集成的桥接扩展。它实现了点云几何类型与 PostGIS 几何类型之间的转换。

```sql
CREATE EXTENSION pointcloud_postgis;
```

该扩展需要同时安装 `pointcloud` 和 `postgis`。它添加了在 `pcpoint`/`pcpatch` 类型与 PostGIS `geometry` 类型之间转换的函数，使得可以使用 PostGIS 运算符和函数对点云数据进行空间查询。
