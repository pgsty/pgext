## 用法

来源：

- [已复核 commit 的 pg_healpix README](https://github.com/segasai/pg_healpix/blob/93220a4588974d9942d5bca7cf1e58048a441a04/README.md)
- [已复核 commit 的 pg_healpix.control](https://github.com/segasai/pg_healpix/blob/93220a4588974d9942d5bca7cf1e58048a441a04/pg_healpix.control)
- [版本 1.0 的 SQL 接口](https://github.com/segasai/pg_healpix/blob/93220a4588974d9942d5bca7cf1e58048a441a04/pg_healpix--1.0.sql)

`pg_healpix` 把 HEALPix 天球像素转换功能暴露为 PostgreSQL 函数。`healpix_ang2ipix_nest` 和 `healpix_ang2ipix_ring` 将 `nside`、赤经与赤纬转换为 NESTED 或 RING `ipix`。`healpix_ipix2ang_nest` 和 `healpix_ipix2ang_ring` 执行反向转换，并返回双元素 `double precision[]`。

### 基本转换

```sql
CREATE EXTENSION pg_healpix;

SELECT healpix_ang2ipix_nest(1024, 12.0, 25.0);
SELECT healpix_ang2ipix_ring(1024, 12.0, 25.0);

SELECT healpix_ipix2ang_nest(1024, 4323);
SELECT healpix_ipix2ang_ring(1024, 4323);
```

这些函数被声明为 `IMMUTABLE`、`PARALLEL SAFE` 和 `STRICT`；当输入约定适合业务时，可以在生成列、表达式索引和并行查询中使用。

### 注意事项

- NESTED 与 RING 标识符不能混用；反向函数必须与正向转换采用相同的排序约定。
- 导入外部天文数据前，应确认上游采用的赤经、赤纬约定以及有效 `nside` 范围。
- control 文件与目录都标识版本 `1.0`。仓库没有提供当前 PostgreSQL 兼容矩阵，因此应在计划运行的确切服务器大版本上测试。
