## 用法

来源：

- [Official source](https://github.com/demitri/pgAST/blob/cc49e5a46f15275fc0d75917422d31a11c726d22/README.md)
- [Official source](https://github.com/demitri/pgAST/blob/cc49e5a46f15275fc0d75917422d31a11c726d22/src/pgast.control)
- [Official source](https://github.com/demitri/pgAST/blob/cc49e5a46f15275fc0d75917422d31a11c726d22/src/pgast--1.1.0.sql)
- [Official source](https://github.com/demitri/pgAST/blob/cc49e5a46f15275fc0d75917422d31a11c726d22/src/Makefile)

`pgast` 将 Starlink AST 的天文世界坐标和区域运算暴露给 PostgreSQL。已核对源码使用扩展版本 1.1.0，要求 AST 9.5.0 或更高版本。当前测试工具要求 PostgreSQL 18 及以上，Makefile 明确处理 PostgreSQL 19 的头文件要求。这不代表已有可用二进制包。

### 启用与查询

安装匹配的编译模块和 AST 依赖后，以超级用户创建扩展。扩展可重定位，无需共享预加载：

```sql
CREATE EXTENSION pgast;
SELECT pgast_point_in_polygon(150.0, 2.0,
  '((150.01,1.99),(149.99,1.99),(149.99,2.01),(150.01,2.01))'::polygon);
SELECT pgast_distance(point(150,2), point(151,2));
```

数据类型使用 PostgreSQL 原生点、圆和多边形，不依赖 PostGIS。坐标代表天球位置，通常为以度为单位的 ICRS 坐标。多边形边沿大圆连接，顶点顺序决定内部区域，反转顺序可能表示天空的其余部分。

### FITS 与区域运算

`pgast_bounding_polygon` 将 FITS 头转换为覆盖区域，`pgast_bounding_circle` 接受多边形或头信息。`pgast_point_in_polygon` 有多边形和头信息重载，`pgast_icrs_polygon_overlaps_hdu` 比较 ICRS 多边形与二维图像头，必要时转换参考系。`pgast_cutout` 接受头信息、天球位置及以角秒为单位的边长。

FITS 头应作为拼接了 80 字符记录的单个字符串传入。PostgreSQL 原生几何运算符使用平面坐标，不能替代跨赤经零点或极区的球面判断。包围圆是覆盖区域的近似，不一定是最小圆。

### 限制与维护

上游库可能将小多边形几乎精确的对跖点误判为内部。无效头信息或不适用的 WCS 会报错并中止语句，文档列出的 SQLSTATE 包括 `22023` 和 `XX000`。批量查询前需验证有代表性的覆盖区域。当前源码有安装脚本，但没有 SQL 升级链，应保留依赖对象并单独规划迁移，不能照搬 README 中会删除对象的重建示例。已核对修订中未找到项目许可证文件。
