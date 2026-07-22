## 用法

来源：

- [1.3 版本 SQL API](https://github.com/bmuratshin/zcurve/blob/76b88d80794a4ccde95c01111453ad8d7fbdbf36/zcurve--1.3.sql)
- [查询实现](https://github.com/bmuratshin/zcurve/blob/76b88d80794a4ccde95c01111453ad8d7fbdbf36/zcurve.c)
- [B-tree 遍历实现](https://github.com/bmuratshin/zcurve/blob/76b88d80794a4ccde95c01111453ad8d7fbdbf36/sp_tree.c)

`zcurve` 把两个或三个整数坐标交错编码成 Z 序键，并通过遍历既有 numeric B-tree 索引执行矩形范围查询。1.3 版本提供 `zcurve_val_from_xy`、`zcurve_num_from_xy`、`zcurve_num_from_xyz`，以及返回堆 TID 和解码坐标的二维/三维查询函数。

```sql
CREATE EXTENSION zcurve;
CREATE TABLE points (x integer, y integer, zkey numeric);
CREATE INDEX points_zkey_idx ON points (zkey);
SELECT * FROM zcurve_2d_lookup('points_zkey_idx', 0, 100, 0, 100);
```

查询参数是索引关系名，而不是表名。实现直接打开 PostgreSQL B-tree 内部结构，假定第一个键包含预期的 numeric Z 序编码，并返回物理 TID；应立即把结果连接回堆表并重新检查坐标谓词，因为更新或表重写会改变 TID。

这是已废弃的服务器内部 C 代码，复制了 numeric/B-tree 结构，含调试日志，没有用户 README，也未声明现代 PostgreSQL 兼容性。服务器大版本、索引定义、坐标范围或数值表示不匹配，都可能崩溃或返回错误行。它只适合在可丢弃数据上做历史研究；生产环境应使用受维护的 GiST/SP-GiST/PostGIS 等空间索引。
