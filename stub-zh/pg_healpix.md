## 用法

来源：

- [已复核 commit 的 pg_healpix README](https://github.com/segasai/pg_healpix/blob/93220a4588974d9942d5bca7cf1e58048a441a04/README.md)
- [版本 1.0 的 SQL 接口](https://github.com/segasai/pg_healpix/blob/93220a4588974d9942d5bca7cf1e58048a441a04/pg_healpix--1.0.sql)
- [坐标转换实现](https://github.com/segasai/pg_healpix/blob/93220a4588974d9942d5bca7cf1e58048a441a04/pg_healpix.c)

`pg_healpix` 提供四个 C 函数，用于在天文学赤经、赤纬与 HEALPix 像素标识符之间转换。它同时支持 nested 和 ring 像素顺序。输入坐标与返回坐标的单位都是度；反向函数返回按赤经、赤纬排序的双元素数组。

### 核心工作流

```sql
CREATE EXTENSION pg_healpix;

SELECT healpix_ang2ipix_nest(1024, 12.0, 25.0);
SELECT healpix_ang2ipix_ring(1024, 12.0, 25.0);

SELECT healpix_ipix2ang_nest(1024, 4323);
SELECT healpix_ipix2ang_ring(1024, 4323);
```

反向转换必须使用与正向转换相同的排序方案。返回坐标代表所选 HEALPix 像素的中心，因此坐标到像素再到坐标的往返属于量化操作，不会精确恢复原始点。

### 函数索引

- `healpix_ang2ipix_nest(bigint, double precision, double precision)` 返回 nested `ipix`。
- `healpix_ang2ipix_ring(bigint, double precision, double precision)` 返回 ring `ipix`。
- `healpix_ipix2ang_nest(bigint, bigint)` 返回 nested `ipix` 对应的赤经/赤纬数组。
- `healpix_ipix2ang_ring(bigint, bigint)` 返回 ring `ipix` 对应的赤经/赤纬数组。

四个函数都声明为 `IMMUTABLE`、`PARALLEL SAFE` 和 `STRICT`；在 PostgreSQL 常规规则允许时，可用于表达式索引、生成表达式和并行计划。

### 输入约束

- `nside` 必须是大于零且不超过 `8192` 的二次幂。
- `ipix` 必须大于等于零且小于 `12 * nside * nside`。
- 赤纬必须映射为零到 pi 之间的极角，即传统的 `-90` 到 `90` 度范围。赤经从度转换，并允许绕过一周。
- 非有限角度输入返回 `NULL`；函数具有严格属性，因此 SQL 空输入也返回 `NULL`。

目录与 control 文件都标识版本 `1.0`。上游没有发布当前 PostgreSQL 兼容矩阵，目录中也只记录了较旧的服务器大版本；应在应用实际使用的服务器大版本与 HEALPix 工具链上测试编译、数值一致性、索引行为和有代表性的边界值。
