---
title: 扩展待办与路线图
description: 当前扩展打包待办、候选项目、维护风险与退役项目。
weight: 700
---

本页根据 2026-08-12 的本地 `pgext` 目录与软件包矩阵整理。目录中出现新版本，并不等同于该版本已经交付；只有对应 RPM/DEB 完成构建、测试并进入仓库索引后，才算真正完成。

## 当前打包待办

### Pigsty 自维护升级

| 软件包 | 当前版本 | 目标版本 | 范围 | 备注 |
|:-------|:---------|:---------|:-----|:-----|
| `pg_profile` | 4.11 | 4.15 | DEB | PGDG RPM 已为 4.15；Pigsty DEB 与源码仍为 4.11。 |
| `pg_readonly` | 1.0.5 | 1.0.6 | DEB | PGDG RPM 已为 1.0.6。 |
| `pg_statement_rollback` | 1.5 | 1.6 | DEB | PGDG RPM 已为 1.6。 |
| `pgsodium` | 3.1.9 | 3.1.11 | Pigsty RPM/DEB | PGDG 在 EL10 上已有 3.1.11，Pigsty 与其他旧 PGDG 目标仍为 3.1.9。 |
| `topn` | 2.7.0 | 2.7.1 | DEB | Pigsty DEB/源码仍为 2.7.0；PGDG RPM 同时存在 2.7.0 与 2.7.1。 |

不要仅为消除表格差异而提前修改目录元数据；应先完成构建、测试、入库与重新扫描。

### 软件包矩阵缺口

- [`pg_statviz`](https://github.com/vyruss/pg_statviz) 是当前唯一仍含 `MISS` 单元格的软件包族：80 个活跃 PG/OS 目标中缺 20 个。缺口包括 EL8/9/10 的 RPM PG17、EL8/9 的 RPM PG18，以及 Ubuntu 22.04 上双架构的 DEB PG14-18。上游/控制文件版本为 1.1，PGDG DEB 为 1.1，PGDG RPM 仍为 0.9。需要决定由 Pigsty 补齐，还是将明确不支持的目标标记为 `N/A`，不应继续保留无法解释的 `MISS`。

### 外部仓库版本差异

以下差异存在于当前矩阵中，但由 PGDG 维护，不应自动视为 Pigsty 重构建任务：

| 软件包 | RPM | DEB |
|:-------|:----|:----|
| `credcheck` | 4.7 | 5.0 |
| `decoderbufs` | 3.5.0 | 3.6.0 |
| `pg_permissions` | 1.4.1 | 1.4 |
| `pg_stat_kcache` | 2.3.1 | 2.3.2 |
| `pgauditlogtofile` | 1.8.4 | 1.8.5 |
| `postgis` | 3.6.3 | 3.6.4 |
| `powa` | 5.1.0 | 5.2.0 |

## 打包候选

以下候选均已进入 `pgext.universe`，但尚未进入正式打包基线。

### 值得评估

- [`pgedge_vectorizer` 1.1](https://github.com/pgEdge/pgedge-vectorizer)：异步文本切分与嵌入生成；C 后台工作进程，需要预加载和 pgvector。当前 1.1 标签仍命名为 `v1.1-test1`。
- [`synchdb`](https://github.com/Hornetlabs/synchdb)：直接从 MySQL、SQL Server、Oracle 进行 CDC；上游发布版本为 1.4，控制文件版本仍为 1.0，且跨越较大的 C/Java 运行时边界。
- [`pg_onnx`](https://github.com/kibae/pg_onnx)：在 PostgreSQL 内执行 ONNX 推理；扩展版本为 1.2.1、项目发布版本为 1.28.0，C++ 运行时依赖较重。
- [`pg_deltax` 0.2.1](https://github.com/xataio/deltax)：活跃的 Rust 时序扩展；打包前需确认 PostgreSQL 大版本与 pgrx 支持范围。
- [`steampipe_postgres_fdw` 1.0](https://github.com/turbot/steampipe-postgres-fdw)：以零 ETL 方式访问云服务与 API；需评估 Go 运行时及插件分发边界。
- [`pg_mustach`](https://github.com/RekGRpth/pg_mustach)：体量较小的 C 语言 Mustache 实现；最新标签为 `v1.0.0`，但控制文件默认版本已经是 3.0，应先厘清发布边界。
- [`is_jsonb_valid` 0.1.4](https://github.com/furstenheim/is_jsonb_valid)：仍在维护的 C 语言 JSON Schema draft 4/7 校验实现；需评估与现有 JSON Schema 扩展的功能重叠。
- [`oai_fdw` 1.13](https://github.com/jimjonesbr/oai_fdw)：仍在活跃维护的 OAI-PMH FDW，但学术元数据场景较窄。
- [`pgjwt_rs` 0.1.2](https://github.com/vishvish/pgjwt)：支持 RS256 与 Ed25519 的 Rust JWT 校验扩展；需评估与现有 JWT/安全扩展的重叠。

### 暂存 / 待评审

- [`coldfront`](https://github.com/pgEdge/coldfront)：面向 PG16-18 的公开测试版；需要预加载、`pg_duckdb`、打补丁的 DuckDB/Iceberg 组件及辅助服务，尚不适合生产。
- [`ruvector`](https://github.com/ruvnet/RuVector)：覆盖面很广且变化迅速的 Rust/向量 monorepo；目录版本为 0.3.0，当前 PostgreSQL crate 为 2.0.6，而仓库标签已进入 2.2 系列。
- [`pg_deeplake`](https://github.com/activeloopai/deeplake) 与 [`vexdb_lite`](https://github.com/VexDB-THU/VexDB-Lite)：方向有吸引力，但 PostgreSQL 打包边界和运行时依赖需要单独评审。
- [`plrust`](https://github.com/pgcentralfoundation/plrust)：编译器与沙箱工具链复杂；项目发布版本为 1.2.8、控制文件默认版本仍为 1.1，上游也只声明 PG13-16 feature。
- [`pg_query_state`](https://github.com/postgrespro/pg_query_state)：依赖两处匹配的 PostgreSQL 内核补丁，不属于常规扩展打包候选。
- `pg_conda`、`pgfdb`、`postgres_ical`、`pgfaker`、`pgsloth`、`pg_kafka`、`pgspeck`、`dsef`、`pg_fsql`、`pg_liquid`、`pg_regresql`：保留在 Universe 中继续观察；优先级较低、仍属实验阶段、场景过窄，或缺少清晰的当前发布边界。

## 状态纠正

### 相比旧清单已完成

- `re2` 0.4.1、`spock` 5.0.10、`pg_lake` 3.4 与 `omnigres` 软件包族，在各自支持矩阵中均已入库且没有 `MISS`。
- `age` 1.8.0、`pg_jieba` 2.0.1、`onesparse` 1.0.0、`pgelog` 1.0.2、`rdf_fdw` 2.7.0、`pg_ttl_index` 3.0.0、`pgcalendar` 1.1.0 均已打包。
- `pg_statviz` 已从 PGDG 纳入目录，目前只剩上文明确列出的矩阵缺口。

### 仍在打包，并未退役

`pg_search`、`pg_net`、`pg_tle`、`pg_bigm`、`http`、`gzip`、`pg_dirtyread`、`pointcloud`、`pg_proctab`、`pgdd`、`pgx_ulid`、`hashtypes`、`pghydro` 仍在活跃打包目录中。旧页面将它们列入“退役”或“尚未规划”已经不再准确。

### 维护风险

- [`columnar` 1.1.2](https://github.com/hydradatabase/columnar) 仍为 PG14-16 提供软件包，但上游自 2025-02-10 后没有新提交，当前目录也没有 PG17/18 支持。
- Apache AGE 仍在活跃开发且已经打包，不再归类为“缺少维护”。

## 尚未规划或已经退役

- [`timescale/pgai`](https://github.com/timescale/pgai) 已归档；[`river`](https://github.com/riverqueue/river) 是 Go 作业队列库，不是 PostgreSQL 扩展。
- `pg_bm25` 已由 `pg_search` 取代；`pg_analytics` 已归档；`pg_lakehouse` 与 `embedding` 已弃用；`pg_sparse` 已并入 pgvector。
- PipelineDB 已弃用；`sql_firewall`、`zcurve`、`pg_comparator` 已停止维护；`weighted_mean` 与 `pg_paxos` 已归档。
- `pg_lz4` 与 `pg_query_state` 依赖打补丁的 PostgreSQL 内核。`vacuumlo`、`oid2name`、`pg_top` 是命令行程序，不是扩展软件包。
- 当前没有打包计划的老旧项目包括：`zson`、`pg_natural_sort_order`、`pgsampler`、`pg_amqp`、`tinyint`、`pg_blkchain`、`foreign_table_exposer`、`ldap_fdw`、`pg_backtrace`、`connection_limits`、`fixeddecimal`、`fuzzywuzzy`、`pg_scws`、`pg_themis`、`lsm3`、`monq`、`pg_recall`、`kmeans`。
- [`jsonb_apply` 0.1.0](https://github.com/Florents-Tselai/jsonb_apply) 仍被阻塞，因为上游仓库没有声明许可证。

## 单侧平台软件包

旧页面中的 EL 独有与 Debian 独有清单已经过时。当前目录只有以下单侧软件包族：

### 仅 RPM

- `db2_fdw` 18.2.0
- `informix_fdw` 0.6.3 — 位于 PGDG non-free，需要 IBM Informix Client SDK
- `pg_strom` 6.1 — GPU/NVMe 扩展，没有 DEB 软件包

### 仅 DEB

- `debversion` 1.2.0

## 相关资源

- [PGXN 最近发布](https://pgxn.org/recent/)
- [PGDG RPM 打包仓库](https://git.postgresql.org/gitweb/?p=pgrpms.git;a=summary)
- [PGDG Debian 打包仓库](https://salsa.debian.org/postgresql)
- [1000+ PostgreSQL 扩展清单](https://gist.github.com/joelonsql/e5aa27f8cc9bd22b8999b7de8aee9d47)
- [PostgreSQL Extension Network](https://www.pgextensions.org/)
