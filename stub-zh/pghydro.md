## 用法

> 来源: [GitHub 仓库](https://github.com/pghydro/pghydro), [README](https://raw.githubusercontent.com/pghydro/pghydro/master/README.md), [releases](https://github.com/pghydro/pghydro/releases)
> PgHydro 套件的主扩展。

PgHydro 在 PostGIS 和 PostgreSQL 之上提供排水网络分析与水资源决策支持。该项目覆盖河网建模、流向分析、Otto Pfafstetter 流域编码、上游与下游河段选择、距河口距离计算、上游汇水面积分析、河流等级以及流域层级等能力。

```sql
CREATE EXTENSION postgis;
CREATE EXTENSION postgis_raster;
CREATE EXTENSION pghydro;
CREATE EXTENSION pgh_raster;
CREATE EXTENSION pgh_hgm;
CREATE EXTENSION pgh_consistency;
CREATE EXTENSION pgh_output;
```

### 组件

- `pghydro` 是核心的排水网络分析扩展。
- `pgh_raster` 使用从数字高程模型派生的水文产品。
- `pgh_hgm` 将 `pghydro` 与 `pgh_raster` 结合起来进行水文地貌分析。
- `pgh_output` 提供报表对象。
- `pgh_consistency` 增加 Pfafstetter 一致性检查。

### 环境要求

- PostgreSQL 9.1 或更高版本。
- PostGIS 3.x。
- PostGIS Raster。

### 说明

- README 表示 master 分支跟踪最新的小版本 6.6。
- CSV 中的主行对应核心 `pghydro` 包，但该仓库在同一发布树中还提供配套扩展。
