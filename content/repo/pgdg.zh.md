---
title: PGDG 仓库
description: "The official PostgreSQL APT/YUM repository"
icon: fa-solid fa-database
weight: 400
---

Pigsty PGSQL 仓库旨在与 PostgreSQL [PGDG 官方仓库](https://www.postgresql.org/download/linux/) 配合使用。

Pigsty 依赖 PGDG 仓库中原生的 PostgreSQL 内核软件包，并在此基础上提供额外的扩展插件。

仓库可用性以当前软件包索引为准，并反映在下方平台表中。


目录覆盖 PostgreSQL 15–19，默认配置仍使用 PostgreSQL 18。自 2026-11-21 起，PG14 已退出活跃构建目标；EL8 继续服务 PostgreSQL 15–18（EL8 上的 `PG19` 标记为 `N/A`），EL9/EL10、Debian 与 Ubuntu 覆盖 PostgreSQL 15–19。各目标实际可用性以[包矩阵](/zh/os/matrix/)为准。

## 快速上手

您可以安装 [pig](https://pig.pgsty.com/zh) - CLI 工具，并使用它添加 pgdg / pigsty 仓库（推荐）：

```bash tab="pig"
pig repo add pgdg                         # 添加 PGDG 仓库
pig repo add pgdg -u                      # 添加 PGDG 仓库，并更新本地缓存
pig repo add pgdg -u --region=default     # 强制使用全球默认区域的仓库（postgresql.org）
pig repo add pgdg -u --region=china       # 使用腾讯云 PGDG 中国镜像仓库
pig repo add pgsql -u                     # pgsql = pgdg + pigsty-pgsql (同时添加 Pigsty 与 PGDG 官方仓库)
pig repo add -u                           # all = node + pgsql (pgdg + pigsty) + infra，一次性添加所有仓库
```



## 镜像

中国区域的 PGDG 路由统一使用腾讯云 PostgreSQL 镜像：
`https://mirrors.cloud.tencent.com/postgresql/repos/`。使用
`--region=china` 或 `--mirror` 时，Pig 与 Pigsty 都会为 PGDG YUM/DNF 和
APT 仓库首选该地址；YUM 定义保留 Pigsty 镜像作为兼容回退。全球
`default` 路由仍使用 PostgreSQL 官方仓库。

{{< pgext-platforms >}}

EL YUM/DNF 仓库信息：

```yaml
- { name: pgdg15         ,description: 'PostgreSQL 15'      ,module: pgsql   ,releases: [8,9,10]     ,arch: [x86_64, aarch64] ,baseurl: { default: 'https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-$releasever-$basearch' ,china: 'https://mirrors.cloud.tencent.com/postgresql/repos/yum/15/redhat/rhel-$releasever-$basearch https://repo.pigsty.cc/yum/pgdg/15/redhat/rhel-$releasever-$basearch' ,europe: 'https://mirrors.xtom.de/postgresql/repos/yum/15/redhat/rhel-$releasever-$basearch' }}
- { name: pgdg16         ,description: 'PostgreSQL 16'      ,module: pgsql   ,releases: [8,9,10]     ,arch: [x86_64, aarch64] ,baseurl: { default: 'https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-$releasever-$basearch' ,china: 'https://mirrors.cloud.tencent.com/postgresql/repos/yum/16/redhat/rhel-$releasever-$basearch https://repo.pigsty.cc/yum/pgdg/16/redhat/rhel-$releasever-$basearch' ,europe: 'https://mirrors.xtom.de/postgresql/repos/yum/16/redhat/rhel-$releasever-$basearch' }}
- { name: pgdg17         ,description: 'PostgreSQL 17'      ,module: pgsql   ,releases: [8,9,10]     ,arch: [x86_64, aarch64] ,baseurl: { default: 'https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-$releasever-$basearch' ,china: 'https://mirrors.cloud.tencent.com/postgresql/repos/yum/17/redhat/rhel-$releasever-$basearch https://repo.pigsty.cc/yum/pgdg/17/redhat/rhel-$releasever-$basearch' ,europe: 'https://mirrors.xtom.de/postgresql/repos/yum/17/redhat/rhel-$releasever-$basearch' }}
- { name: pgdg18         ,description: 'PostgreSQL 18'      ,module: pgsql   ,releases: [8,9,10]     ,arch: [x86_64, aarch64] ,baseurl: { default: 'https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-$releasever-$basearch' ,china: 'https://mirrors.cloud.tencent.com/postgresql/repos/yum/18/redhat/rhel-$releasever-$basearch https://repo.pigsty.cc/yum/pgdg/18/redhat/rhel-$releasever-$basearch' ,europe: 'https://mirrors.xtom.de/postgresql/repos/yum/18/redhat/rhel-$releasever-$basearch' }}
- { name: pgdg19         ,description: 'PostgreSQL 19'      ,module: pgsql   ,releases: [  9,10]     ,arch: [x86_64, aarch64] ,baseurl: { default: 'https://download.postgresql.org/pub/repos/yum/19/redhat/rhel-$releasever-$basearch' ,china: 'https://mirrors.cloud.tencent.com/postgresql/repos/yum/19/redhat/rhel-$releasever-$basearch https://repo.pigsty.cc/yum/pgdg/19/redhat/rhel-$releasever-$basearch' ,europe: 'https://mirrors.xtom.de/postgresql/repos/yum/19/redhat/rhel-$releasever-$basearch' }}
```

Debian / Ubuntu APT 仓库信息：

```yaml
- { name: pgdg           ,description: 'PGDG'               ,module: pgsql   ,releases: [12,13,22,24,26] ,arch: [x86_64, aarch64] ,baseurl: { default: 'http://apt.postgresql.org/pub/repos/apt/ ${distro_codename}-pgdg main'            ,china: 'https://mirrors.cloud.tencent.com/postgresql/repos/apt/ ${distro_codename}-pgdg main' }}
```



## APT GPG

PGDG APT 仓库使用以下 GPG 密钥签名: `B97B0AFCAA1A47F044F244A07FCC7D46ACCC4CF8` (`ACCC4CF8`)

MD5 校验和为 `f54c5c1aa1329dc26e33b29762faaec4`，详情参考 https://www.postgresql.org/download/linux/debian/

```bash
sudo curl -fsSL https://www.postgresql.org/media/keys/ACCC4CF8.asc -o /usr/share/postgresql-common/pgdg/apt.postgresql.org.asc
. /etc/os-release
sudo sh -c "echo 'deb [signed-by=/usr/share/postgresql-common/pgdg/apt.postgresql.org.asc] https://apt.postgresql.org/pub/repos/apt $VERSION_CODENAME-pgdg main' > /etc/apt/sources.list.d/pgdg.list"
```

使用腾讯云镜像时，可以直接获取同一份 PGDG APT 公钥：

```bash
sudo curl -fsSL https://mirrors.cloud.tencent.com/postgresql/repos/apt/ACCC4CF8.asc -o /usr/share/postgresql-common/pgdg/apt.postgresql.org.asc
. /etc/os-release
sudo sh -c "echo 'deb [signed-by=/usr/share/postgresql-common/pgdg/apt.postgresql.org.asc] https://mirrors.cloud.tencent.com/postgresql/repos/apt/ $VERSION_CODENAME-pgdg main' > /etc/apt/sources.list.d/pgdg.list"
```


## YUM GPG

PGDG YUM 仓库使用 `https://ftp.postgresql.org/pub/repos/yum/keys/` 中的一系列密钥进行签名。请自行按需选取使用
