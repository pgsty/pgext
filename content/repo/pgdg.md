---
title: PGDG Repo
description: "The official PostgreSQL APT/YUM repository"
icon: fa-solid fa-database
weight: 400
---

The Pigsty PGSQL Repo is designed to work together with the official PostgreSQL Global Development Group ([PGDG](https://www.postgresql.org/download/linux/)) repo.
Together, they can provide up to [627 packaged PostgreSQL extensions](https://ext.pgsty.com/list) out-of-the-box.

Repository availability is derived from the current package indexes and reflected in the platform table below.

The catalog covers PostgreSQL 15–19, with PostgreSQL 18 remaining the configuration default. PG14 was retired from active build targets on 2026-11-21; EL8 remains supported for PostgreSQL 15–18 (`PG19` is `N/A` on EL8), while EL9/EL10, Debian, and Ubuntu cover PostgreSQL 15–19. See the [package matrix](/os/matrix/) for per-target availability.

## Quick Start

You can install [pig](https://pig.pgsty.com) - the cli tool, and add pgdg repo with it (recommended):

```bash
pig repo add pgdg                           # add pgdg repo file
pig repo add pgdg -u                        # add pgdg repo and update cache
pig repo add pgdg -u --region=default       # add pgdg repo, enforce using the default repo (postgresql.org)
pig repo add pgdg -u --region=china         # add PGDG repo, always use Tencent Cloud's China mirror
pig repo add pgsql -u                       # pgsql = pgdg + pigsty-pgsql (add pigsty + official PGDG)
pig repo add -u                             # all = node + pgsql (pgdg + pigsty) + infra
```


## Mirror

The China-region PGDG route uses Tencent Cloud's PostgreSQL mirror at
`https://mirrors.cloud.tencent.com/postgresql/repos/`. Pig and Pigsty select it
first for PGDG YUM/DNF and APT repositories when `--region=china` or `--mirror`
is used; YUM definitions retain the Pigsty mirror as a compatibility fallback.
The global `default` route remains the official PostgreSQL repository.

{{< pgext-platforms >}}

EL YUM/DNF Repo info:

```yaml
- { name: pgdg15         ,description: 'PostgreSQL 15'      ,module: pgsql   ,releases: [8,9,10]     ,arch: [x86_64, aarch64] ,baseurl: { default: 'https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-$releasever-$basearch' ,china: 'https://mirrors.cloud.tencent.com/postgresql/repos/yum/15/redhat/rhel-$releasever-$basearch https://repo.pigsty.cc/yum/pgdg/15/redhat/rhel-$releasever-$basearch' ,europe: 'https://mirrors.xtom.de/postgresql/repos/yum/15/redhat/rhel-$releasever-$basearch' }}
- { name: pgdg16         ,description: 'PostgreSQL 16'      ,module: pgsql   ,releases: [8,9,10]     ,arch: [x86_64, aarch64] ,baseurl: { default: 'https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-$releasever-$basearch' ,china: 'https://mirrors.cloud.tencent.com/postgresql/repos/yum/16/redhat/rhel-$releasever-$basearch https://repo.pigsty.cc/yum/pgdg/16/redhat/rhel-$releasever-$basearch' ,europe: 'https://mirrors.xtom.de/postgresql/repos/yum/16/redhat/rhel-$releasever-$basearch' }}
- { name: pgdg17         ,description: 'PostgreSQL 17'      ,module: pgsql   ,releases: [8,9,10]     ,arch: [x86_64, aarch64] ,baseurl: { default: 'https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-$releasever-$basearch' ,china: 'https://mirrors.cloud.tencent.com/postgresql/repos/yum/17/redhat/rhel-$releasever-$basearch https://repo.pigsty.cc/yum/pgdg/17/redhat/rhel-$releasever-$basearch' ,europe: 'https://mirrors.xtom.de/postgresql/repos/yum/17/redhat/rhel-$releasever-$basearch' }}
- { name: pgdg18         ,description: 'PostgreSQL 18'      ,module: pgsql   ,releases: [8,9,10]     ,arch: [x86_64, aarch64] ,baseurl: { default: 'https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-$releasever-$basearch' ,china: 'https://mirrors.cloud.tencent.com/postgresql/repos/yum/18/redhat/rhel-$releasever-$basearch https://repo.pigsty.cc/yum/pgdg/18/redhat/rhel-$releasever-$basearch' ,europe: 'https://mirrors.xtom.de/postgresql/repos/yum/18/redhat/rhel-$releasever-$basearch' }}
- { name: pgdg19         ,description: 'PostgreSQL 19'      ,module: pgsql   ,releases: [  9,10]     ,arch: [x86_64, aarch64] ,baseurl: { default: 'https://download.postgresql.org/pub/repos/yum/19/redhat/rhel-$releasever-$basearch' ,china: 'https://mirrors.cloud.tencent.com/postgresql/repos/yum/19/redhat/rhel-$releasever-$basearch https://repo.pigsty.cc/yum/pgdg/19/redhat/rhel-$releasever-$basearch' ,europe: 'https://mirrors.xtom.de/postgresql/repos/yum/19/redhat/rhel-$releasever-$basearch' }}
```

Debian / Ubuntu APT Repo info:

```yaml
- { name: pgdg           ,description: 'PGDG'               ,module: pgsql   ,releases: [12,13,22,24,26] ,arch: [x86_64, aarch64] ,baseurl: { default: 'http://apt.postgresql.org/pub/repos/apt/ ${distro_codename}-pgdg main'            ,china: 'https://mirrors.cloud.tencent.com/postgresql/repos/apt/ ${distro_codename}-pgdg main' }}
```
