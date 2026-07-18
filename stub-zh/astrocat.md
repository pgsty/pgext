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

项目自称实验性质，并大量基于 pgSphere。信任结果前，应确认坐标单位、赤经回绕、赤纬边界、球面距离、空值行为和极点附近边界情况。

它没有当前兼容性或数值精度矩阵。应使用已知参考星表对比索引与顺序扫描结果，测试 GiST recheck 与并发更新，并基准选择率和 KNN 排序。如果索引键表示或距离语义改变，必须重建所有依赖索引。生产应优先使用持续维护的 pgSphere 或其他受支持天文空间栈。
