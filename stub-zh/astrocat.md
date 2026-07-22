## 用法

来源：

- [已复核 commit 的上游 README](https://github.com/karpov-sv/astrocat/blob/5c714e05f7ee11f7cf61c482c189a919ec37edd9/README.md)
- [1.0 版本 SQL 对象](https://github.com/karpov-sv/astrocat/blob/5c714e05f7ee11f7cf61c482c189a919ec37edd9/astrocat--1.0.sql)
- [GiST 实现](https://github.com/karpov-sv/astrocat/blob/5c714e05f7ee11f7cf61c482c189a919ec37edd9/astrocat.c)

`astrocat` 是天文目录点索引试验台。它通过函数式 GiST 索引支持径向锥形搜索和 K 近邻排序，赤经与赤纬均以度表示。

```sql
CREATE EXTENSION astrocat;
CREATE TABLE star (ra double precision, dec double precision);
CREATE INDEX star_sky_idx ON star USING gist (skypoint(ra, dec));
SELECT *
FROM star
WHERE skypoint(ra, dec) @ skycircle(11.2, 67.0, 2.0)
ORDER BY skypoint(ra, dec) <-> skypoint(11.2, 67.0)
LIMIT 10;
```

项目自称实验性质，并大量基于 pgSphere。其构造器与文本输入函数既不验证解析是否成功，也不约束赤经、赤纬和半径范围。相等判断逐字节比较内部浮点表示。构造前必须验证输入，且不能假定数学上等价的坐标一定相等。

已复核的 GiST 代码存在正确性缺陷。直接 `@` 包含判断会接纳恰好落在圆周上的点，但叶节点一致性函数使用严格比较并设置 `recheck = false`，因此索引扫描可能漏掉该点。内部节点的 KNN 距离代码在一个分支中把查询的 y 坐标与键的 x 上界比较，可能破坏最近邻排序。必须对比索引与顺序扫描的锥形查询结果，并把 KNN 输出与独立计算、排序的球面距离对比，尤其要覆盖精确边界。

它没有当前兼容性或数值精度矩阵。在修复并独立验证前，不要让这个操作符类承担权威结果。如果索引键表示、包含判断或距离代码改变，必须重建所有依赖索引。生产应优先使用持续维护的 pgSphere 或其他受支持天文空间栈。
