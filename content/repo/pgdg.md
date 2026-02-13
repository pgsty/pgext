---
title: PGDG Repo
description: "The official PostgreSQL APT/YUM repository"
icon: Database
weight: 400
---

The Pigsty PGSQL Repo is designed to work together with the official PostgreSQL Global Development Group ([PGDG](https://www.postgresql.org/download/linux/)) repo.
Together, they can provide up to [451 PostgreSQL Extensions](https://ext.pgsty.com/list) out-of-the-box.

Mirror synced at 2025-12-01 18:00:00

## Quick Start

You can install [pig](/pig) - the cli tool, and add pgdg repo with it (recommended):

```bash
pig repo add pgdg                           # add pgdg repo file
pig repo add pgdg -u                        # add pgdg repo and update cache
pig repo add pgdg -u --region=default       # add pgdg repo, enforce using the default repo (postgresql.org)
pig repo add pgdg -u --region=china         # add pgdg repo, always use the china mirror (repo.pigsty.cc)
pig repo add pgsql -u                       # pgsql = pgdg + pigsty-pgsql (add pigsty + official PGDG)
pig repo add -u                             # all = node + pgsql (pgdg + pigsty) + infra
```


## Mirror

Since 2025-05, PGDG has closed the rsync/ftp sync channel, which makes almost all mirror site out-of-sync.

Currently, Pigsty, Yandex, Xtom are providing regular synced mirror service.

The pigsty pgdg mirror is a subset of official pgdg repo, cover EL 7-10, Debian 11-13, Ubuntu 20.04 - 24.04, with x86_64 & arm6 and pg 13 - 19alpha. 

| OS Code                                                                               | Vendor     | Major            | PG Major                                                                                                                                                                                                                                                |                       Comment         |
|:----------------------------------------------------------------------------------------|:-------|:---:|:------------------|:------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|:----------------------------------------------:|
| {{< badge content="el7.x86_64"  color="green" >}}                                       | EL     |  7  |  {{< badge content="18" color="green" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} {{< badge content="13" color="green" >}} |   {{< badge content="EOL" color="green" >}}    |
| {{< badge content="el8.x86_64"  link="/os/el8.x86_64"  color="green" >}}              | EL     |  8  |  {{< badge content="18" color="green" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} {{< badge content="13" color="green" >}} | {{< badge content="Near EOL" color="green" >}} |
| {{< badge content="el8.aarch64" link="/os/el8.aarch64" color="green" border=false >}} | EL     |  8  |  {{< badge content="18" color="green" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} {{< badge content="13" color="green" >}} | {{< badge content="Near EOL" color="green" >}} |
| {{< badge content="el9.x86_64"  link="/os/el9.x86_64"  color="green" >}}              | EL     |  9  |  {{< badge content="18" color="green" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} {{< badge content="13" color="green" >}} |    {{< badge content="OK" color="green" >}}    |
| {{< badge content="el9.aarch64" link="/os/el9.aarch64" color="green" border=false >}} | EL     |  9  |  {{< badge content="18" color="green" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} {{< badge content="13" color="green" >}} |    {{< badge content="OK" color="green" >}}    |
| {{< badge content="el10.x86_64" link="/os/el10.x86_64" color="green" >}}              | EL     | 10  |  {{< badge content="18" color="green" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} {{< badge content="13" color="green" >}} |    {{< badge content="OK" color="green" >}}    |
| {{< badge content="el10.aarch64"link="/os/el10.aarch64"color="green" border=false >}} | EL     | 10  |  {{< badge content="18" color="green" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} {{< badge content="13" color="green" >}} |    {{< badge content="OK" color="green" >}}    |
| {{< badge content="d11.x86_64"  color="green" >}}                                       | Debian | 11  |  {{< badge content="18" color="green" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} {{< badge content="13" color="green" >}} |   {{< badge content="EOL" color="green" >}}    |
| {{< badge content="d11.aarch64" color="green" border=false >}}                          | Debian | 11  |  {{< badge content="18" color="green" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} {{< badge content="13" color="green" >}} |   {{< badge content="EOL" color="green" >}}    |
| {{< badge content="d12.x86_64"  link="/os/d12.x86_64"  color="green" >}}              | Debian | 12  |  {{< badge content="18" color="green" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} {{< badge content="13" color="green" >}} |    {{< badge content="OK" color="green" >}}    |
| {{< badge content="d12.aarch64" link="/os/d12.aarch64" color="green" border=false >}} | Debian | 12  |  {{< badge content="18" color="green" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} {{< badge content="13" color="green" >}} |    {{< badge content="OK" color="green" >}}    |
| {{< badge content="d13.x86_64"  link="/os/d13.x86_64"  color="green" >}}              | Debian | 13  |  {{< badge content="18" color="green" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} {{< badge content="13" color="green" >}} |    {{< badge content="OK" color="green" >}}    |
| {{< badge content="d13.aarch64" link="/os/d13.aarch64" color="green" border=false >}} | Debian | 13  |  {{< badge content="18" color="green" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} {{< badge content="13" color="green" >}} |    {{< badge content="OK" color="green" >}}    |
| {{< badge content="u20.x86_64" color="green" >}}                                        | Ubuntu | 20  |  {{< badge content="18" color="green" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} {{< badge content="13" color="green" >}} |   {{< badge content="EOL" color="green" >}}    |
| {{< badge content="u20.aarch64"color="green" border=false >}}                           | Ubuntu | 20  |  {{< badge content="18" color="green" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} {{< badge content="13" color="green" >}} |   {{< badge content="EOL" color="green" >}}    |
| {{< badge content="u22.x86_64"  link="/os/u22.x86_64"  color="green" >}}              | Ubuntu | 22  |  {{< badge content="18" color="green" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} {{< badge content="13" color="green" >}} |    {{< badge content="OK" color="green" >}}    |
| {{< badge content="u22.aarch64" link="/os/u22.aarch64" color="green" border=false >}} | Ubuntu | 22  |  {{< badge content="18" color="green" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} {{< badge content="13" color="green" >}} |    {{< badge content="OK" color="green" >}}    |
| {{< badge content="u24.x86_64"  link="/os/u24.x86_64"  color="green" >}}              | Ubuntu | 24  |  {{< badge content="18" color="green" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} {{< badge content="13" color="green" >}} |    {{< badge content="OK" color="green" >}}    |
| {{< badge content="u24.aarch64" link="/os/u24.aarch64" color="green" border=false >}} | Ubuntu | 24  |  {{< badge content="18" color="green" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} {{< badge content="13" color="green" >}} |    {{< badge content="OK" color="green" >}}    |

EL YUM/DNF Repo info:

```yaml
- { name: pgdg13         ,description: 'PostgreSQL 13'      ,module: pgsql   ,releases: [7,8,9,10]       ,arch: [x86_64, aarch64] ,baseurl: { default: 'https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-$releasever-$basearch' ,china: 'https://repo.pigsty.cc/yum/pgdg/13/redhat/rhel-$releasever-$basearch' ,europe: 'https://mirrors.xtom.de/postgresql/repos/yum/13/redhat/rhel-$releasever-$basearch' }}
- { name: pgdg14         ,description: 'PostgreSQL 14'      ,module: pgsql   ,releases: [7,8,9,10]       ,arch: [x86_64, aarch64] ,baseurl: { default: 'https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-$releasever-$basearch' ,china: 'https://repo.pigsty.cc/yum/pgdg/14/redhat/rhel-$releasever-$basearch' ,europe: 'https://mirrors.xtom.de/postgresql/repos/yum/14/redhat/rhel-$releasever-$basearch' }}
- { name: pgdg15         ,description: 'PostgreSQL 15'      ,module: pgsql   ,releases: [7,8,9,10]       ,arch: [x86_64, aarch64] ,baseurl: { default: 'https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-$releasever-$basearch' ,china: 'https://repo.pigsty.cc/yum/pgdg/15/redhat/rhel-$releasever-$basearch' ,europe: 'https://mirrors.xtom.de/postgresql/repos/yum/15/redhat/rhel-$releasever-$basearch' }}
- { name: pgdg16         ,description: 'PostgreSQL 16'      ,module: pgsql   ,releases: [  8,9,10]       ,arch: [x86_64, aarch64] ,baseurl: { default: 'https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-$releasever-$basearch' ,china: 'https://repo.pigsty.cc/yum/pgdg/16/redhat/rhel-$releasever-$basearch' ,europe: 'https://mirrors.xtom.de/postgresql/repos/yum/16/redhat/rhel-$releasever-$basearch' }}
- { name: pgdg17         ,description: 'PostgreSQL 17'      ,module: pgsql   ,releases: [  8,9,10]       ,arch: [x86_64, aarch64] ,baseurl: { default: 'https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-$releasever-$basearch' ,china: 'https://repo.pigsty.cc/yum/pgdg/17/redhat/rhel-$releasever-$basearch' ,europe: 'https://mirrors.xtom.de/postgresql/repos/yum/17/redhat/rhel-$releasever-$basearch' }}
- { name: pgdg18         ,description: 'PostgreSQL 18'      ,module: pgsql   ,releases: [  8,9,10]       ,arch: [x86_64, aarch64] ,baseurl: { default: 'https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-$releasever-$basearch' ,china: 'https://repo.pigsty.cc/yum/pgdg/18/redhat/rhel-$releasever-$basearch' ,europe: 'https://mirrors.xtom.de/postgresql/repos/yum/18/redhat/rhel-$releasever-$basearch' }}
- { name: pgdg19-beta    ,description: 'PostgreSQL 19 Beta' ,module: beta    ,releases: [  8,9,10]       ,arch: [x86_64, aarch64] ,baseurl: { default: 'https://download.postgresql.org/pub/repos/yum/testing/19/redhat/rhel-$releasever-$basearch' ,china: 'https://repo.pigsty.cc/yum/pgdg/testing/19/redhat/rhel-$releasever-$basearch' ,europe: 'https://mirrors.xtom.de/postgresql/repos/yum/testing/19/redhat/rhel-$releasever-$basearch' }}
```

Debian / Ubuntu APT Repo info:

```yaml
- { name: pgdg           ,description: 'PGDG'               ,module: pgsql   ,releases: [11,12,13,   22,24] ,arch: [x86_64, aarch64] ,baseurl: { default: 'http://apt.postgresql.org/pub/repos/apt/ ${distro_codename}-pgdg main'            ,china: 'https://repo.pigsty.cc/apt/pgdg/ ${distro_codename}-pgdg main' }}
- { name: pgdg-beta      ,description: 'PGDG Beta'          ,module: beta    ,releases: [11,12,13,   22,24] ,arch: [x86_64, aarch64] ,baseurl: { default: 'http://apt.postgresql.org/pub/repos/apt/ ${distro_codename}-pgdg-testing main 19' ,china: 'https://repo.pigsty.cc/apt/pgdg/ ${distro_codename}-pgdg-testing main 19' }}
```