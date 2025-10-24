---
title: 仓库
description: 用于交付 PostgreSQL 扩展的基础设施和软件包仓库
icon: Warehouse
weight: 300
breadcrumbs: false
---


Pigsty 为主流 [Linux 发行版](https://pgsty.com/docs/prepare/linux) 提供了一个 PostgreSQL 扩展仓库，其中包含 200+ 额外的 PostgreSQL 扩展。

Pigsty 扩展仓库旨在与 [PGDG](https://www.postgresql.org/download/linux/) 官方仓库配合使用，共同使用时可以安装多达 [424 个 PostgreSQL 扩展](/zh/list)。


--------

## 快速上手

你可以使用 [pig](/pig/) CLI 工具启用 infra 和 pgsql 仓库，或者手动将它们添加到系统中：

```bash tab="pig"
curl https://repo.pigsty.io/pig | bash      # 下载安装最新版本的 pig 命令行工具
pig repo add all -u                         # 添加 Linux / PGDG / Pigsty 仓库并更新缓存
```
```bash tab="apt"
# 将 Pigsty 的 GPG 公钥添加到系统密钥链中，以验证软件包签名
curl -fsSL https://repo.pigsty.io/key | sudo gpg --dearmor -o /etc/apt/keyrings/pigsty.gpg

# 获取 Debian / Ubuntu 发行版的代号（jammy, focal, bullseye, bookworm），并将相应的上游仓库地址写入 /etc/apt/sources.list.d/ 中
distro_codename=$(lsb_release -cs)
sudo tee /etc/apt/sources.list.d/pigsty-io.list > /dev/null <<EOF
deb [signed-by=/etc/apt/keyrings/pigsty.gpg] https://repo.pigsty.io/apt/infra generic main
deb [signed-by=/etc/apt/keyrings/pigsty.gpg] https://repo.pigsty.io/apt/pgsql/${distro_codename} ${distro_codename} main
EOF

# 刷新 APT 仓库缓存
sudo apt update
```
```bash tab="yum"
# 将 Pigsty 的 GPG 公钥添加到系统密钥链中，以验证软件包签名
curl -fsSL https://repo.pigsty.io/key | sudo tee /etc/pki/rpm-gpg/RPM-GPG-KEY-pigsty >/dev/null

# 将 Pigsty 仓库的定义写入 /etc/yum.repos.d/ 目录中
sudo tee /etc/yum.repos.d/pigsty-io.repo > /dev/null <<-'EOF'
[pigsty-infra]
name=Pigsty Infra for $basearch
baseurl=https://repo.pigsty.io/yum/infra/$basearch
skip_if_unavailable = 1
enabled = 1
priority = 1
gpgcheck = 1
gpgkey=file:///etc/pki/rpm-gpg/RPM-GPG-KEY-pigsty
module_hotfixes=1

[pigsty-pgsql]
name=Pigsty PGSQL For el$releasever.$basearch
baseurl=https://repo.pigsty.io/yum/pgsql/el$releasever.$basearch
skip_if_unavailable = 1
enabled = 1
priority = 1
gpgcheck = 1
gpgkey=file:///etc/pki/rpm-gpg/RPM-GPG-KEY-pigsty
module_hotfixes=1
EOF

# 刷新 YUM 仓库缓存
sudo yum makecache;
```

所有的 RPM / DEB 软件包都使用指纹为 (`B9BD8B20`) 的 [GPG 密钥](/zh/repo/gpg) 进行签名，以确保软件包的完整性。


---------

## 兼容性

Pigsty 仓库由两个主要部分组成：[`INFRA`](/zh/repo/infra) 和 [`PGSQL`](/zh/repo/pgsql)，提供 `x86_64` 和 `aarch64` 架构下的 DEB / RPM 包。

[`INFRA`](/repo/infra) 仓库中的软件包与 PostgreSQL / Linux 大版本无关，包括 Prometheus、Grafana、以及一些 PostgreSQL 管理工具，
通常由 Go 等语言编写的，只有芯片架构（`x86_64` | `aarch64`）的区别。

| Linux  |  软件包  | x86_64 | aarch64 |
|:------:|:-----:|:------:|:-------:|
|   EL   | `rpm` |   ✓    |    ✓    |
| Debian | `deb` |   ✓    |    ✓    |

[`PGSQL`](/repo/pgsql) 仓库中的软件包通常特定于 Linux 大版本（例如 el9, d12），也通常与 PostgreSQL 大版本相关（例如 pg17，pg16 ）
这个仓库中包含 了 PostgreSQL 内核疯子、扩展插件与工具，通常由类 C 语言编写。

|   系统 / 架构    | OS  |                                                                                               x86_64                                                                                                |                                                                                               aarch64                                                                                               |
|:------------:|:---:|:---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------:|:---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------:|
|     EL8      | el8 | <Badge variant="blue-subtle">17</Badge><Badge variant="blue-subtle">16</Badge><Badge variant="blue-subtle">15</Badge><Badge variant="blue-subtle">14</Badge><Badge variant="blue-subtle">13</Badge> | <Badge variant="blue-subtle">17</Badge><Badge variant="blue-subtle">16</Badge><Badge variant="blue-subtle">15</Badge><Badge variant="blue-subtle">14</Badge><Badge variant="blue-subtle">13</Badge> |
|     EL9      | el9 | <Badge variant="blue-subtle">17</Badge><Badge variant="blue-subtle">16</Badge><Badge variant="blue-subtle">15</Badge><Badge variant="blue-subtle">14</Badge><Badge variant="blue-subtle">13</Badge> | <Badge variant="blue-subtle">17</Badge><Badge variant="blue-subtle">16</Badge><Badge variant="blue-subtle">15</Badge><Badge variant="blue-subtle">14</Badge><Badge variant="blue-subtle">13</Badge> |
|  Debian 12   | d12 | <Badge variant="blue-subtle">17</Badge><Badge variant="blue-subtle">16</Badge><Badge variant="blue-subtle">15</Badge><Badge variant="blue-subtle">14</Badge><Badge variant="blue-subtle">13</Badge> | <Badge variant="blue-subtle">17</Badge><Badge variant="blue-subtle">16</Badge><Badge variant="blue-subtle">15</Badge><Badge variant="blue-subtle">14</Badge><Badge variant="blue-subtle">13</Badge> |
| Ubuntu 22.04 | u22 | <Badge variant="blue-subtle">17</Badge><Badge variant="blue-subtle">16</Badge><Badge variant="blue-subtle">15</Badge><Badge variant="blue-subtle">14</Badge><Badge variant="blue-subtle">13</Badge> | <Badge variant="blue-subtle">17</Badge><Badge variant="blue-subtle">16</Badge><Badge variant="blue-subtle">15</Badge><Badge variant="blue-subtle">14</Badge><Badge variant="blue-subtle">13</Badge> |
| Ubuntu 24.04 | u24 | <Badge variant="blue-subtle">17</Badge><Badge variant="blue-subtle">16</Badge><Badge variant="blue-subtle">15</Badge><Badge variant="blue-subtle">14</Badge><Badge variant="blue-subtle">13</Badge> | <Badge variant="blue-subtle">17</Badge><Badge variant="blue-subtle">16</Badge><Badge variant="blue-subtle">15</Badge><Badge variant="blue-subtle">14</Badge><Badge variant="blue-subtle">13</Badge> |


|                        系统代码                         | 厂商     | 大版本 |   小版本   | 全名                |                                                                                                      PG 大版本                                                                                                       | 备注       |
|:---------------------------------------------------:|:-------|:---:|:-------:|:------------------|:-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------:|:---------|
|  {{< badge content="el7.x86_64"   color="red" >}}   | EL     |  7  |   7.9   | CentOS 7 x86      |      {{< badge content="18" color="red" >}} {{< badge content="17" color="red" >}} {{< badge content="16" color="red" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}}       | EOL      |
| {{< badge content="el8.x86_64"   color="green" >}}  | EL     |  8  |  8.10   | RockyLinux 8 x86  |   {{< badge content="18" color="green" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}}    | Near EOL |
| {{< badge content="el8.aarch64"  color="green" >}}  | EL     |  8  |  8.10   | RockyLinux 8 ARM  |   {{< badge content="18" color="green" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}}    | Near EOL |
| {{< badge content="el9.x86_64"   color="green" >}}  | EL     |  9  |   9.6   | RockyLinux 9 x86  |   {{< badge content="18" color="green" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}}    | OK       |
| {{< badge content="el9.aarch64"  color="green" >}}  | EL     |  9  |   9.6   | RockyLinux 9 ARM  |   {{< badge content="18" color="green" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}}    | OK       |
| {{< badge content="el10.x86_64"  color="yellow" >}} | EL     | 10  |  10.0   | RockyLinux 10 x86 |   {{< badge content="18" color="green" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}}    | TBD      |
| {{< badge content="el10.aarch64" color="yellow" >}} | EL     | 10  |  10.0   | RockyLinux 10 ARM |   {{< badge content="18" color="green" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}}    | TBD      |
|  {{< badge content="d11.x86_64"   color="red" >}}   | Debian | 11  |  11.11  | Debian 11 x86     | {{< badge content="18" color="yellow" >}} {{< badge content="17" color="yellow" >}} {{< badge content="16" color="yellow" >}} {{< badge content="15" color="yellow" >}} {{< badge content="14" color="yellow" >}} | EOL      |
|  {{< badge content="d11.aarch64"  color="red" >}}   | Debian | 11  |  11.11  | Debian 11 ARM     | {{< badge content="18" color="yellow" >}} {{< badge content="17" color="yellow" >}} {{< badge content="16" color="yellow" >}} {{< badge content="15" color="yellow" >}} {{< badge content="14" color="yellow" >}} | EOL      |
| {{< badge content="d12.x86_64"   color="green" >}}  | Debian | 12  |  12.11  | Debian 12 x86     |   {{< badge content="18" color="green" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}}    | OK       |
| {{< badge content="d12.aarch64"  color="green" >}}  | Debian | 12  |  12.11  | Debian 12 ARM     |   {{< badge content="18" color="green" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}}    | OK       |
| {{< badge content="d13.x86_64"   color="yellow" >}} | Debian | 13  |  13.1   | Debian 13 x86     |   {{< badge content="18" color="green" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}}    | TBD      |
| {{< badge content="d13.aarch64"  color="yellow" >}} | Debian | 13  |  13.1   | Debian 13 ARM     |   {{< badge content="18" color="green" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}}    | TBD      |
|  {{< badge content="u20.x86_64"   color="red" >}}   | Ubuntu | 20  | 20.04.6 | Ubuntu 20.04 x86  | {{< badge content="18" color="yellow" >}} {{< badge content="17" color="yellow" >}} {{< badge content="16" color="yellow" >}} {{< badge content="15" color="yellow" >}} {{< badge content="14" color="yellow" >}} | EOL      |
|  {{< badge content="u20.aarch64"  color="red" >}}   | Ubuntu | 20  | 20.04.6 | Ubuntu 20.04 ARM  | {{< badge content="18" color="yellow" >}} {{< badge content="17" color="yellow" >}} {{< badge content="16" color="yellow" >}} {{< badge content="15" color="yellow" >}} {{< badge content="14" color="yellow" >}} | EOL      |
| {{< badge content="u22.x86_64"   color="green" >}}  | Ubuntu | 22  | 22.04.5 | Ubuntu 22.04 x86  |   {{< badge content="18" color="green" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}}    | OK       |
| {{< badge content="u22.aarch64"  color="green" >}}  | Ubuntu | 22  | 22.04.5 | Ubuntu 22.04 ARM  |   {{< badge content="18" color="green" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}}    | OK       |
| {{< badge content="u24.x86_64"   color="green" >}}  | Ubuntu | 24  | 24.04.3 | Ubuntu 24.04 x86  |   {{< badge content="18" color="green" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}}    | OK       |
| {{< badge content="u24.aarch64"  color="green" >}}  | Ubuntu | 24  | 24.04.3 | Ubuntu 24.04 ARM  |   {{< badge content="18" color="green" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}}    | OK       |


------

## 源代码

用于构建仓库内软件的源代码文件位于以下仓库中：

- https://github.com/pgsty/rpm
- https://github.com/pgsty/deb
- https://github.com/pgsty/infra-pkg
