---
title: PostgreSQL 扩展云
description: PostgreSQL 扩展大全，使用 pig 命令行工具，在 14+ Linux 发行版上一键安装 431 个 PostgreSQL 扩展。
cascade:
  type: docs
breadcrumbs: false
excludeSearch: true
width: full
toc: false
comment: false
---


Pigsty 提供了以下三样基础设施，帮助用户更好的利用 PostgreSQL 扩展生态系统的协同超能力：

- [**扩展目录**](/list) ： 查阅 [**431**](/e/) 个扩展插件的详细信息，找到满足您需求的插件
- [**扩展仓库**](/repo) ： 获取预先打包的 RPM/DEB 二进制包，在 [**14 个 Linux 大版本**](/zh/os) 上可用
- [**包管理器**](/pig) ： 屏蔽复杂度与操作系统与架构差异，使用 [**`pig`**](/zh/pig) 命令行工具一键完成所有步骤

{{< cards >}}
  {{< card link="/zh/list" title="扩展目录" icon="clipboard-list" subtitle="430+ 个可用 PostgreSQL 扩展的完整列表" >}}
  {{< card link="/zh/repo" title="软件仓库" icon="cube" subtitle="包含扩展二进制包的 APT/YUM 软件仓库" >}}
  {{< card link="/zh/pig" title="包管理器" icon="cash" subtitle="PostgreSQL 与扩展生态中缺失的包管理器" >}}
{{< /cards >}}

只需几行命令，即可在 [**主流 Linux 系统**](/zh/os) 上原生安装多达 431 个 PostgreSQL 扩展，组合使用，发挥 PostgreSQL 的超能力：

```bash
curl -fsSL https://repo.pigsty.io/pig | bash  # 安装 pig 命令行工具
pig repo set                  # 在您的 Linux 发行版上配置软件仓库
pig install pg18              # 从 PGDG 官方仓库安装 PostgreSQL 18 内核包
pig install pg_duckdb -v 18   # 例：针对 PG 18 安装 pg_duckdb
```

一切都可以用 PostgreSQL 解决！请参阅我们的博客文章：[**PostgreSQL 正在吞噬数据库世界！**](https://medium.com/@fengruohang/postgres-is-eating-the-database-world-157c204dcfc4)

[![ecosystem](/ecosystem.gif)](https://medium.com/@fengruohang/postgres-is-eating-the-database-world-157c204dcfc4)


## 扩展导航

| {{< category "time" >}} | {{< category "gis" >}}  | {{< category "rag" >}}   | {{< category "fts" >}}  | {{< category "olap" >}} | {{< category "feat" >}} | {{< category "lang" >}} | {{< category "type" >}} |
|-------------------------|-------------------------|--------------------------|-------------------------|-------------------------|-------------------------|-------------------------|-------------------------| 
| {{< category "util" >}} | {{< category "func" >}} | {{< category "admin" >}} | {{< category "stat" >}} | {{< category "sec" >}}  | {{< category "fdw" >}}  | {{< category "sim" >}}  | {{< category "etl" >}}  |

{{< cards cols=5 >}}
  {{< card link="/zh/list/cate#time" title="TIME" icon="clock" subtitle="时间时态扩展：时序数据库 TimescaleDB，时态数据库，版本控制表，定时任务，异步后台任务调度扩展。" >}}
  {{< card link="/zh/list/cate#gis" title="GIS" icon="globe" subtitle="地理空间扩展：PostGIS，地理空间数据类型、函数与索引，天空索引 Q3C，OGR FDW， 寻路算法，地理正/逆查询。" >}}
  {{< card link="/zh/list/cate#rag" title="RAG" icon="light-bulb" subtitle="AI与RAG扩展插件：向量数据库，DiskANN 向量索引，相似度度量函数集，库内机器学习与推理 pgml，等等。" >}}
  {{< card link="/zh/list/cate#fts" title="FTS" icon="search" subtitle="全文检索扩展：ES 替代 pg_search，BM25，中文分词，欧洲语言分词字典 hunspell，模糊检索，2-gram/3-gram 索引。" >}}
  {{< card link="/zh/list/cate#olap" title="OLAP" icon="chart-bar" subtitle="分析能力扩展：列式存储，DuckDB与外部数据源包装器，Parquet S3，数据冷热分级存储，分布式计算，透明分片，GPU加速。" >}}
  {{< card link="/zh/list/cate#feat" title="FEAT" icon="sparkles" subtitle="功能特性扩展：图数据库，Hyperloglog，Rum索引，GraphQL，JsonSchema，Hint，虚拟索引，增量物化视图，消息队列等等。" >}}
  {{< card link="/zh/list/cate#lang" title="LANG" icon="book-open" subtitle="存储过程语言扩展：使用各种编程语言开发，调试，打包，分发，测试 PostgreSQL 存储过程：Java，Js，Lua，R，SH，PRQL…" >}}
  {{< card link="/zh/list/cate#type" title="TYPE" icon="cube-transparent" subtitle="自定义类型扩展：前缀树，语义版本号，SI单位，位图，无符号整型，高精度数值，有理数，哈希值，IP地址段，球面，RRULE等。" >}}
  {{< card link="/zh/list/cate#util" title="UTIL" icon="cog" subtitle="实用功能扩展：HTTP请求，GZIP压缩，JWT处理，邮件客户端，正则，字符编码，编码解码，加密解密等实用功能。" >}}
  {{< card link="/zh/list/cate#func" title="FUNC" icon="variable" subtitle="标识聚合函数：ID生成器，各类聚合函数，摘要函数，数组处理函数，数学函数，统计量，伪随机，等等。" >}}
  {{< card link="/zh/list/cate#admin" title="ADMIN" icon="office-building" subtitle="管理工具扩展：膨胀治理，脏读，检视缓冲区，数据目录，校验和，系统目录腐败检查，优先级管理，权限管理，语句准备，限制批量更新等。" >}}
  {{< card link="/zh/list/cate#stat" title="STAT" icon="presentation-chart-line" subtitle="监控统计扩展：AWR报告，可观测性指标，显示执行计划，查询统计信息，内存使用，配置变更，等待事件采样，慢查询日志，等等。" >}}
  {{< card link="/zh/list/cate#sec" title="SEC" icon="shield-check" subtitle="安全功能扩展：强制密码强度，阉割超级用户，密钥管理，商密算法，PII匿名处理，扩展白名单，审计日志，变更追溯，反病毒等等。" >}}
  {{< card link="/zh/list/cate#fdw" title="FDW" icon="document-download" subtitle="外部数据源包装器：FDW开发框架 Wrappers，Multicorn，访问外部的 Mongo，MySQL，SQLite，HDFS，MSSQL，Oracle，DB2，……" >}}
  {{< card link="/zh/list/cate#sim" title="SIM" icon="terminal" subtitle="数据库兼容扩展：仿真其他 DBMS 的行为：MySQL，Memcache，Mongo，Oracle，Babelfish for Microsoft SQL Server……" >}}
  {{< card link="/zh/list/cate#etl" title="ETL" icon="truck" subtitle="数据复制扩展：逻辑复制，逻辑解码，DDL复制，JSON/BSON/Protobuf 变更抽取，数据迁移，数据导入，数据比对等。" >}}
{{< /cards >}}

| {{< license "MIT" >}}      | {{< license "ISC" >}}        | {{< license "PostgreSQL" >}} | {{< license "BSD 0-Clause" >}} | {{< license "BSD 2-Clause" >}} | {{< license "BSD 3-Clause" >}} |
|----------------------------|------------------------------|------------------------------|--------------------------------|--------------------------------|--------------------------------|
| {{< license "Artistic" >}} | {{< license "Apache-2.0" >}} | {{< license "MPL-2.0" >}}    |                                |                                |                                |
| {{< license "GPL-2.0" >}}  | {{< license "GPL-3.0" >}}    | {{< license "LGPL-2.1" >}}   | {{< license "LGPL-3.0" >}}     | {{< license "AGPL-3.0" >}}     | {{< license "Timescale" >}}    | 
| {{< language "c" >}}       | {{< language "c++" >}}       | {{< language "rust" >}}      | {{< language "java" >}}        | {{< language "python" >}}      | {{< language "sql" >}}         |{{< language "data" >}}|



## 用户

**PGSTY.CLOUD** 向用户提供无可比拟的扩展交付体验，一些 PostgreSQL 厂商使用本仓库进行交付：

{{< cards cols=2 >}}
{{< card link="https://github.com/github.com/pgsty/pigsty"     title="Pigsty"     icon="github" subtitle="The author and maintainer, Battery-Included PG RDS" >}}
{{< card link="https://docs.omnigres.org/quick_start/"         title="Omnigres"   icon="github" subtitle="The All-in-One PostgreSQL as a Platform" >}}
{{< card link="https://autobase.tech/docs/extensions/install"  title="AutoBase"   icon="github" subtitle="Self-hosted DBaaS / Postgres automation" >}}
{{< /cards >}}

服务仓库托管于 Cloudflare CDN 上（中国区域提供专用镜像仓库），这是一项向 PostgreSQL 用户社区提供的 **免费服务**。


## 系统兼容性

| OS                                                                                    | Vendor | Major |  Minor  | Fullname          | PG Major Version                                                                                                                                                                                                                                            |                     Comment                     |
|:--------------------------------------------------------------------------------------|:-------|:-----:|:-------:|:------------------|:------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|:-----------------------------------------------:|
| {{< badge content="el8.x86_64"  link="/os/el8.x86_64"  color="green" >}}              | EL     |   8   |  8.10   | RockyLinux 8 x86  | {{< badge content="18" color="green" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} {{< badge content="13" color="green" >}}       | {{< badge content="Near EOL" color="orange" >}} |
| {{< badge content="el8.aarch64" link="/os/el8.aarch64" color="green" border=false >}} | EL     |   8   |  8.10   | RockyLinux 8 ARM  | {{< badge content="18" color="green" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} {{< badge content="13" color="green" >}}       | {{< badge content="Near EOL" color="orange" >}} |
| {{< badge content="el9.x86_64"  link="/os/el9.x86_64"  color="green" >}}              | EL     |   9   |   9.6   | RockyLinux 9 x86  | {{< badge content="18" color="green" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} {{< badge content="13" color="green" >}}       |    {{< badge content="OK" color="green" >}}     |
| {{< badge content="el9.aarch64" link="/os/el9.aarch64" color="green" border=false >}} | EL     |   9   |   9.6   | RockyLinux 9 ARM  | {{< badge content="18" color="green" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} {{< badge content="13" color="green" >}}       |    {{< badge content="OK" color="green" >}}     |
| {{< badge content="el10.x86_64" link="/os/el10.x86_64" color="green" >}}              | EL     |  10   |  10.0   | RockyLinux 10 x86 | {{< badge content="18" color="green" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} {{< badge content="13" color="green" >}}       |    {{< badge content="OK" color="green" >}}     |
| {{< badge content="el10.aarch64"link="/os/el10.aarch64"color="green" border=false >}} | EL     |  10   |  10.0   | RockyLinux 10 ARM | {{< badge content="18" color="green" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} {{< badge content="13" color="green" >}}       |    {{< badge content="OK" color="green" >}}     |
| {{< badge content="d12.x86_64"  link="/os/d12.x86_64"  color="green" >}}              | Debian |  12   |  12.12  | Debian 12 x86     | {{< badge content="18" color="green" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} {{< badge content="13" color="green" >}}       |    {{< badge content="OK" color="green" >}}     |
| {{< badge content="d12.aarch64" link="/os/d12.aarch64" color="green" border=false >}} | Debian |  12   |  12.12  | Debian 12 ARM     | {{< badge content="18" color="green" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} {{< badge content="13" color="green" >}}       |    {{< badge content="OK" color="green" >}}     |
| {{< badge content="d13.x86_64"  link="/os/d13.x86_64"  color="green" >}}              | Debian |  13   |  13.1   | Debian 13 x86     | {{< badge content="18" color="green" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} {{< badge content="13" color="green" >}}       |    {{< badge content="OK" color="green" >}}     |
| {{< badge content="d13.aarch64" link="/os/d13.aarch64" color="green" border=false >}} | Debian |  13   |  13.1   | Debian 13 ARM     | {{< badge content="18" color="green" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} {{< badge content="13" color="green" >}}       |    {{< badge content="OK" color="green" >}}     |
| {{< badge content="u22.x86_64"  link="/os/u22.x86_64"  color="green" >}}              | Ubuntu |  22   | 22.04.5 | Ubuntu 22.04 x86  | {{< badge content="18" color="green" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} {{< badge content="13" color="green" >}}       |    {{< badge content="OK" color="green" >}}     |
| {{< badge content="u22.aarch64" link="/os/u22.aarch64" color="green" border=false >}} | Ubuntu |  22   | 22.04.5 | Ubuntu 22.04 ARM  | {{< badge content="18" color="green" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} {{< badge content="13" color="green" >}}       |    {{< badge content="OK" color="green" >}}     |
| {{< badge content="u24.x86_64"  link="/os/u24.x86_64"  color="green" >}}              | Ubuntu |  24   | 24.04.3 | Ubuntu 24.04 x86  | {{< badge content="18" color="green" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} {{< badge content="13" color="green" >}}       |    {{< badge content="OK" color="green" >}}     |
| {{< badge content="u24.aarch64" link="/os/u24.aarch64" color="green" border=false >}} | Ubuntu |  24   | 24.04.3 | Ubuntu 24.04 ARM  | {{< badge content="18" color="green" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} {{< badge content="13" color="green" >}}       |    {{< badge content="OK" color="green" >}}     |



--------

## 关于

PGEXT.CLOUD 项目由 [**PGSTY**](https://github.com/pgsty) / [**VONNG**](https://blog.vonng.com/en/) (rh@vonng.com) 开发维护，并使用  [**Apache 2.0**](https://github.com/pgsty/pig/?tab=Apache-2.0-1-ov-file#readme) 许可证开源。

| GitHub 仓库                                                        | 说明                |
|------------------------------------------------------------------|-------------------|
| [github.com/pgsty](https://github.com/pgsty)                     | **PGSTY** 组织主页    |
| [github.com/pgsty/pgext](https://github.com/pgsty/pgext)         | 本网站，扩展元数据，以及管理工具  |
| [github.com/pgsty/pigsty](https://github.com/pgsty/pigsty)       | PostgreSQL 数据库发行版 |
| [github.com/pgsty/pig](https://github.com/pgsty/pig)             | PostgreSQL 包管理器   |
| [github.com/pgsty/ext](https://github.com/pgsty/ext)             | 本文档，扩展目录          |
| [github.com/pgsty/rpm](https://github.com/pgsty/rpm)             | RPM 构建源代码         |
| [github.com/pgsty/deb](https://github.com/pgsty/deb)             | DEB 构建源代码         |
| [github.com/pgsty/infra-pkg](https://github.com/pgsty/infra-pkg) | 基础设施包仓库           |



