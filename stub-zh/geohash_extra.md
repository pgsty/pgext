## 用法

来源：

- [上游 README](https://github.com/geo-mathijs/geohash-extra/blob/f65bdc89f60c041bbf10a7b11badae237e30480f/README.md)
- [扩展 control 文件](https://github.com/geo-mathijs/geohash-extra/blob/f65bdc89f60c041bbf10a7b11badae237e30480f/geohash_extra.control)
- [SQL 封装](https://github.com/geo-mathijs/geohash-extra/blob/f65bdc89f60c041bbf10a7b11badae237e30480f/geohash_extra--0.0.1.sql)
- [C 实现](https://github.com/geo-mathijs/geohash-extra/blob/f65bdc89f60c041bbf10a7b11badae237e30480f/geohash_extra.c)

`geohash_extra` 为 PostGIS 补充两个集合返回操作：`geohash_neighbours()` 返回八个相邻 geohash，`geohashfromgeom()` 返回根据几何边界推导出的 geohash 网格。目录与 control 文件中的版本均为 `0.0.1`。

### 相邻单元与几何覆盖

```sql
CREATE EXTENSION postgis;
CREATE EXTENSION geohash_extra;

SELECT * FROM geohash_neighbours('u173zq37x014');

SELECT *
FROM geohashfromgeom(
  ST_MakeEnvelope(4.89, 52.36, 4.91, 52.38, 4326),
  7
);
```

`geohashfromgeom()` 会把非 4326 输入转换为 SRID 4326，再从包围盒角点的 geohash 构造矩形网格。它不是不规则几何的精确覆盖，可能返回位于几何之外的单元。control 文件没有声明 PostGIS 依赖，因此应先安装 PostGIS。

当前 C 实现为相邻计算的输入/输出缓冲区少分配了一个终止符字节，而且对 geohash 字符的校验不完整。因此空输入或畸形输入可能破坏内存或令后端崩溃。在修复并重新构建代码前，不应让不可信输入调用 `geohash_neighbours()`。该仓库没有提供现代 PostgreSQL 兼容矩阵。
