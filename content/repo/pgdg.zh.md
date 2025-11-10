---
title: PGDG Repo
description: "The official PostgreSQL APT/YUM repository"
icon: Database
weight: 400
---

Pigsty PGSQL 仓库旨在与 PostgreSQL [PGDG 官方仓库](https://www.postgresql.org/download/linux/) 配合使用。

Pigsty 依赖 PGDG 仓库中原生的 PostgreSQL 内核软件包，在此基础上提供了额外的 200+ 扩展插件。


## 快速上手

您可以安装 [pig](/zh/pig) - CLI 工具，并使用它添加 pgdg / pigsty 仓库（推荐）：

```bash tab="pig"
pig repo add pgdg                         # 添加 PGDG 仓库
pig repo add pgdg -u                      # 添加 PGDG 仓库，并更新本地缓存
pig repo add pgdg -u --region=default     # 强制使用全球默认区域的仓库（postgresql.org）
pig repo add pgdg -u --region=china       # 使用中国镜像仓库 (repo.pigsty.cc)
pig repo add pgsql -u                     # pgsql = pgdg + pigsty-pgsql (同时添加 Pigsty 与 PGDG 官方仓库)
pig repo add -u                           # all = node + pgsql (pgdg + pigsty) + infra，一次性添加所有仓库
```

## 镜像

2025年5月中旬，PGDG 关闭了 rsync/ftp 同步渠道，导致全球几乎所有 PGDG 镜像站失去同步，根据观察，目前只有 YANDEX，XTOM，PIGSTY 提供定期同步。

Pigsty 在中国区域提供了 PGDG 镜像的子集，更新频率约为一周一更新。对于 EL 7-10，Debian 11-13，Ubuntu 20.04 - 24.04 提供 x86_64 与 arm64 架构的镜像仓库。覆盖范围为所有生命周期内的 PG 大版本（PG12 - 19alpha）

| OS 系统代码                                                                               | 厂商     | 大版本            | PG 大版本                                                                                                                                                                                                                                                |                       备注                       |
|:----------------------------------------------------------------------------------------|:-------|:---:|:------------------|:------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|:----------------------------------------------:|
| {{< badge content="el7.x86_64"  color="green" >}}                                       | EL     |  7  |  {{< badge content="18" color="green" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} {{< badge content="13" color="green" >}} |   {{< badge content="EOL" color="green" >}}    |
| {{< badge content="el8.x86_64"  link="/list/el8.x86_64"  color="green" >}}              | EL     |  8  |  {{< badge content="18" color="green" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} {{< badge content="13" color="green" >}} | {{< badge content="Near EOL" color="green" >}} |
| {{< badge content="el8.aarch64" link="/list/el8.aarch64" color="green" border=false >}} | EL     |  8  |  {{< badge content="18" color="green" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} {{< badge content="13" color="green" >}} | {{< badge content="Near EOL" color="green" >}} |
| {{< badge content="el9.x86_64"  link="/list/el9.x86_64"  color="green" >}}              | EL     |  9  |  {{< badge content="18" color="green" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} {{< badge content="13" color="green" >}} |    {{< badge content="OK" color="green" >}}    |
| {{< badge content="el9.aarch64" link="/list/el9.aarch64" color="green" border=false >}} | EL     |  9  |  {{< badge content="18" color="green" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} {{< badge content="13" color="green" >}} |    {{< badge content="OK" color="green" >}}    |
| {{< badge content="el10.x86_64" link="/list/el10.x86_64" color="green" >}}              | EL     | 10  |  {{< badge content="18" color="green" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} {{< badge content="13" color="green" >}} |    {{< badge content="OK" color="green" >}}    |
| {{< badge content="el10.aarch64"link="/list/el10.aarch64"color="green" border=false >}} | EL     | 10  |  {{< badge content="18" color="green" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} {{< badge content="13" color="green" >}} |    {{< badge content="OK" color="green" >}}    |
| {{< badge content="d11.x86_64"  color="green" >}}                                       | Debian | 11  |  {{< badge content="18" color="green" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} {{< badge content="13" color="green" >}} |   {{< badge content="EOL" color="green" >}}    |
| {{< badge content="d11.aarch64" color="green" border=false >}}                          | Debian | 11  |  {{< badge content="18" color="green" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} {{< badge content="13" color="green" >}} |   {{< badge content="EOL" color="green" >}}    |
| {{< badge content="d12.x86_64"  link="/list/d12.x86_64"  color="green" >}}              | Debian | 12  |  {{< badge content="18" color="green" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} {{< badge content="13" color="green" >}} |    {{< badge content="OK" color="green" >}}    |
| {{< badge content="d12.aarch64" link="/list/d12.aarch64" color="green" border=false >}} | Debian | 12  |  {{< badge content="18" color="green" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} {{< badge content="13" color="green" >}} |    {{< badge content="OK" color="green" >}}    |
| {{< badge content="d13.x86_64"  link="/list/d13.x86_64"  color="green" >}}              | Debian | 13  |  {{< badge content="18" color="green" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} {{< badge content="13" color="green" >}} |    {{< badge content="OK" color="green" >}}    |
| {{< badge content="d13.aarch64" link="/list/d13.aarch64" color="green" border=false >}} | Debian | 13  |  {{< badge content="18" color="green" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} {{< badge content="13" color="green" >}} |    {{< badge content="OK" color="green" >}}    |
| {{< badge content="u20.x86_64" color="green" >}}                                        | Ubuntu | 20  |  {{< badge content="18" color="green" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} {{< badge content="13" color="green" >}} |   {{< badge content="EOL" color="green" >}}    |
| {{< badge content="u20.aarch64"color="green" border=false >}}                           | Ubuntu | 20  |  {{< badge content="18" color="green" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} {{< badge content="13" color="green" >}} |   {{< badge content="EOL" color="green" >}}    |
| {{< badge content="u22.x86_64"  link="/list/u22.x86_64"  color="green" >}}              | Ubuntu | 22  |  {{< badge content="18" color="green" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} {{< badge content="13" color="green" >}} |    {{< badge content="OK" color="green" >}}    |
| {{< badge content="u22.aarch64" link="/list/u22.aarch64" color="green" border=false >}} | Ubuntu | 22  |  {{< badge content="18" color="green" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} {{< badge content="13" color="green" >}} |    {{< badge content="OK" color="green" >}}    |
| {{< badge content="u24.x86_64"  link="/list/u24.x86_64"  color="green" >}}              | Ubuntu | 24  |  {{< badge content="18" color="green" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} {{< badge content="13" color="green" >}} |    {{< badge content="OK" color="green" >}}    |
| {{< badge content="u24.aarch64" link="/list/u24.aarch64" color="green" border=false >}} | Ubuntu | 24  |  {{< badge content="18" color="green" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} {{< badge content="13" color="green" >}} |    {{< badge content="OK" color="green" >}}    |

EL YUM/DNF 仓库信息：

```yaml
- { name: pgdg13         ,description: 'PostgreSQL 13'      ,module: pgsql   ,releases: [7,8,9,10]       ,arch: [x86_64, aarch64] ,baseurl: { default: 'https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-$releasever-$basearch' ,china: 'https://repo.pigsty.cc/yum/pgdg/13/redhat/rhel-$releasever-$basearch' ,europe: 'https://mirrors.xtom.de/postgresql/repos/yum/13/redhat/rhel-$releasever-$basearch' }}
- { name: pgdg14         ,description: 'PostgreSQL 14'      ,module: pgsql   ,releases: [7,8,9,10]       ,arch: [x86_64, aarch64] ,baseurl: { default: 'https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-$releasever-$basearch' ,china: 'https://repo.pigsty.cc/yum/pgdg/14/redhat/rhel-$releasever-$basearch' ,europe: 'https://mirrors.xtom.de/postgresql/repos/yum/14/redhat/rhel-$releasever-$basearch' }}
- { name: pgdg15         ,description: 'PostgreSQL 15'      ,module: pgsql   ,releases: [7,8,9,10]       ,arch: [x86_64, aarch64] ,baseurl: { default: 'https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-$releasever-$basearch' ,china: 'https://repo.pigsty.cc/yum/pgdg/15/redhat/rhel-$releasever-$basearch' ,europe: 'https://mirrors.xtom.de/postgresql/repos/yum/15/redhat/rhel-$releasever-$basearch' }}
- { name: pgdg16         ,description: 'PostgreSQL 16'      ,module: pgsql   ,releases: [  8,9,10]       ,arch: [x86_64, aarch64] ,baseurl: { default: 'https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-$releasever-$basearch' ,china: 'https://repo.pigsty.cc/yum/pgdg/16/redhat/rhel-$releasever-$basearch' ,europe: 'https://mirrors.xtom.de/postgresql/repos/yum/16/redhat/rhel-$releasever-$basearch' }}
- { name: pgdg17         ,description: 'PostgreSQL 17'      ,module: pgsql   ,releases: [  8,9,10]       ,arch: [x86_64, aarch64] ,baseurl: { default: 'https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-$releasever-$basearch' ,china: 'https://repo.pigsty.cc/yum/pgdg/17/redhat/rhel-$releasever-$basearch' ,europe: 'https://mirrors.xtom.de/postgresql/repos/yum/17/redhat/rhel-$releasever-$basearch' }}
- { name: pgdg18         ,description: 'PostgreSQL 18'      ,module: pgsql   ,releases: [  8,9,10]       ,arch: [x86_64, aarch64] ,baseurl: { default: 'https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-$releasever-$basearch' ,china: 'https://repo.pigsty.cc/yum/pgdg/18/redhat/rhel-$releasever-$basearch' ,europe: 'https://mirrors.xtom.de/postgresql/repos/yum/18/redhat/rhel-$releasever-$basearch' }}
- { name: pgdg19-beta    ,description: 'PostgreSQL 19 Beta' ,module: beta    ,releases: [  8,9,10]       ,arch: [x86_64, aarch64] ,baseurl: { default: 'https://download.postgresql.org/pub/repos/yum/testing/19/redhat/rhel-$releasever-$basearch' ,china: 'https://repo.pigsty.cc/yum/pgdg/testing/19/redhat/rhel-$releasever-$basearch' ,europe: 'https://mirrors.xtom.de/postgresql/repos/yum/testing/19/redhat/rhel-$releasever-$basearch' }}
```

Debian / Ubuntu APT 仓库信息：

```yaml
- { name: pgdg           ,description: 'PGDG'               ,module: pgsql   ,releases: [11,12,13,   22,24] ,arch: [x86_64, aarch64] ,baseurl: { default: 'http://apt.postgresql.org/pub/repos/apt/ ${distro_codename}-pgdg main'            ,china: 'https://repo.pigsty.cc/apt/pgdg/ ${distro_codename}-pgdg main' }}
- { name: pgdg-beta      ,description: 'PGDG Beta'          ,module: beta    ,releases: [11,12,13,   22,24] ,arch: [x86_64, aarch64] ,baseurl: { default: 'http://apt.postgresql.org/pub/repos/apt/ ${distro_codename}-pgdg-testing main 19' ,china: 'https://repo.pigsty.cc/apt/pgdg/ ${distro_codename}-pgdg-testing main 19' }}
```