---
title: 软件包
description: 扩展包与别名
icon: Package
---


管理扩展与软件包并非易事，以下以两个常见扩展为例：

![](/img/pigsty/ext-mapping.png)

| 实体          | 示例 `pgvector`                        | 示例 `postgis`...                                         |
|---------------|----------------------------------------|----------------------------------------------------------|
| **扩展名**    | `vector`                               | `postgis`, `postgis_topology`, `postgis_raster`,...      |
| **软件包**    | `pgvector`                             | `postgis`                                                |
| **操作系统包**| pgvector_17                            | postgresql-16-postgis-3                                  |
| **RPM/DEB 包**| pgvector_17_0.8.0-1PGDG.rhel8.x86_64.rpm| postgresql-17-postgis-3_3.5.2+dfsg-1.pgdg22.04+1_amd64.deb |

如需高效安装正确的 RPM / DEB 包，可借助抽象层：**软件包别名**。
因此，你只需指定“规范化”名称（如 `pgvector` 或 `postgis`），即可完成扩展的下载与安装。
无需关心 PG 版本、操作系统版本、架构、扩展版本等细节。

软件包别名 `pkg` 用于扩展的下载与安装，但在数据库中执行 `CREATE EXTENSION` 时，需使用扩展名 `ext`（如在 `meta` 库中使用 `vector`）。
注意，部分扩展（如上例中的 `timescaledb`）需显式预加载。

此外，所有扩展被划分为 16 个主要类别，每个类别也有对应的别名，
便于批量下载与安装，例如：

```yaml title="将 17 替换为 16、15、14、13..."
repo_extra_packages: [ pg17-main ,pg17-core ,pg17-time ,pg17-gis ,pg17-rag ,pg17-fts ,pg17-olap ,pg17-feat ,pg17-lang ,pg17-type ,pg17-util ,pg17-func ,pg17-admin ,pg17-stat ,pg17-sec ,pg17-fdw ,pg17-sim ,pg17-etl]
pg_extensions: [pg17-time ,pg17-gis ,pg17-rag ,pg17-fts ,pg17-feat ,pg17-lang ,pg17-type ,pg17-util ,pg17-func ,pg17-admin ,pg17-stat ,pg17-sec ,pg17-fdw ,pg17-sim ,pg17-etl ] #,pg17-olap]
```

所有扩展**均可**同时安装，唯独 `olap` 类别下，`citus` 与 `hydra` 冲突，`pg_duckdb` 与 `pg_mooncake` 冲突。
因此可全部下载，但需分批安装。



