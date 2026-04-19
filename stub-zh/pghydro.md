## 用法

来源：[README](https://github.com/pghydro/pghydro/blob/master/README.md)，[repo](https://github.com/pghydro/pghydro)，[releases](https://github.com/pghydro/pghydro/releases)

`pghydro` 是 PgHydro 套件中的核心扩展，用于在 PostgreSQL 和 PostGIS 之上做排水网络分析与水资源决策支持。

### 安装 PgHydro 套件

```sql
CREATE EXTENSION postgis;
CREATE EXTENSION postgis_raster;
CREATE EXTENSION pghydro;
CREATE EXTENSION pgh_raster;
CREATE EXTENSION pgh_hgm;
CREATE EXTENSION pgh_consistency;
CREATE EXTENSION pgh_output;
```

上游 README 将这些配套扩展一起介绍：

- `pghydro`：排水网络分析
- `pgh_raster`：基于 DEM 派生水文产品
- `pgh_hgm`：水文地貌分析
- `pgh_consistency`：Pfafstetter 一致性检查
- `pgh_output`：报表对象

### 上游覆盖的能力

- 河网中的流向修正
- Otto Pfafstetter 流域编码
- 上游和下游河段选择
- 距河口距离计算
- 上游汇水面积计算
- 河流等级和流域层级分析

### 环境要求

- PostgreSQL 9.1+
- PostGIS 3.x
- PostGIS Raster

### 说明

- 当前 README 的状态说明仍写明 master 分支跟踪 `6.6`，develop 分支跟踪 `6.7-dev`。
- 仓库里已经发布了更高版本的 tag，但面向用户的 README 仍以 `6.6` 的安装和教程流程为主。
