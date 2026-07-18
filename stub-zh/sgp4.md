## 用法

来源：

- [项目 README 与精度免责声明](https://github.com/gabriel-russo/postgresql-sgp4/blob/bcf6ea2b45285856f1b7b68c9da78ba129a5249d/README.md)
- [扩展 control 文件](https://github.com/gabriel-russo/postgresql-sgp4/blob/bcf6ea2b45285856f1b7b68c9da78ba129a5249d/sgp4.control)
- [0.1 版 SQL API](https://github.com/gabriel-russo/postgresql-sgp4/blob/bcf6ea2b45285856f1b7b68c9da78ba129a5249d/sgp4--0.1.sql)
- [PostgreSQL 包装实现](https://github.com/gabriel-russo/postgresql-sgp4/blob/bcf6ea2b45285856f1b7b68c9da78ba129a5249d/sgp4.c)

`sgp4` 0.1 在 PostgreSQL 中嵌入 SGP4 轨道传播器。`satellite_geographic_position()` 接收 TLE 的两行和一个 `timestamp without time zone`，返回形如 `POINT Z(longitude latitude altitude)` 的 WKT 文本。经纬度是 WGS84 度数，高度单位为千米。

### 在 UTC 时刻传播 TLE

```sql
CREATE EXTENSION sgp4;

SELECT satellite_geographic_position(
  '1 44883U 19093E   26042.94689667  .00001181  00000+0  15674-3 0  9994',
  '2 44883  97.7645 114.1422 0001434 192.3199 167.7980 14.81620074332535',
  timestamp '2026-02-12 12:00:00'
);
```

该函数不执行时区转换。去掉时区前，应先把权威时间转换为 UTC，并在每个调用点记录此约定。传播前检查 TLE 校验和、epoch、目录标识和新鲜度；超出适用 epoch 窗口后，TLE 精度会衰减。

### 精度与空间边界

上游明确排除科学、天体测量、高精度大地测量、安全关键和监管用途。其 TEME 至地固坐标转换省略完整 IAU 岁差与章动、分点方程、极移和实时 UT1 修正。根据轨道和 epoch，与高精度库相比，经度差可能达到若干分之一度。

结果是文本而不是 PostGIS geometry。若转换到 SRID 4326，必须记住 Z 值仍以千米为单位，且水平 SRID 不定义垂直参考或单位。应在代表性轨道、错误 TLE、闰日与边界日期以及预期误差容限上，与受信 SGP4 实现比较。不能把没有数据库错误或得到看似合理的 WKT 字符串当成输入有效的证据。
