


## 用法

来源：[PGHydro README](https://github.com/pghydro/pghydro/blob/master/README.md)，[pgh_hgm SQL](https://github.com/pghydro/pghydro/blob/master/pgh_hgm--2.2.6.sql)

`pgh_hgm` 为 PgHydro 增加水文地貌分析函数。它结合排水线、排水区、DEM 以及派生栅格产品，计算水资源分析中常用的流域和河网指标。

### 启用

在核心 PgHydro 和栅格扩展之后安装：

```sql
CREATE EXTENSION postgis;
CREATE EXTENSION postgis_raster;
CREATE EXTENSION pghydro;
CREATE EXTENSION pgh_raster;
CREATE EXTENSION pgh_hgm;
```

### 初始化对象

扩展创建 `pgh_hgm` schema，其中包括 `pghft_drn_elevationprofile`、`pghft_upn_elevationprofile`、`pghft_hydro_intel` 等表。

```sql
SELECT pgh_hgm.pghfn_tables_initialize();
```

### 指标

SQL 对象包括高程剖面、排水区高程摘要、坡面流平均长度、轴长、圆形度、紧凑度、排水密度、高程统计、形状因子、水文密度、Kirpich 时间、周长、河段坡降、相对起伏、汇流时间、坡度方法、弯曲度、波速传播，以及 AMHG 深度/宽度估计等函数。

```sql
SELECT nspname, proname
FROM pg_proc p
JOIN pg_namespace n ON n.oid = p.pronamespace
WHERE nspname = 'pgh_hgm'
ORDER BY proname;
```

### 注意事项

运行指标计算前，需要先加载并校验 DEM、flow direction、flow accumulation 以及 PgHydro 排水图层。`pgh_hgm` 依赖 `pgh_raster` 提供栅格派生水文产品。
