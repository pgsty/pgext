---
title: 软件仓库
description: 用于交付 PostgreSQL 扩展的基础设施和软件包仓库
icon: fa-solid fa-warehouse
weight: 300
breadcrumb: false
---


Pigsty 为 16 个 [Linux 平台目标](/zh/os)（8 个发行版大版本，各含 `x86_64` 与 `aarch64`）提供 PostgreSQL 扩展仓库。

Pigsty 扩展仓库旨在与 [PGDG](https://www.postgresql.org/download/linux/) 官方仓库配合使用，共同使用时可以安装多达 [627 个已打包 PostgreSQL 扩展](/zh/list)。

{{< cards cols=2 >}}
{{< card link="/zh/repo/pgsql"  title="PGSQL Repo" subtitle="Pigsty PG 扩展仓库" icon="play" />}}
{{< card link="/zh/repo/infra"  title="INFRA Repo" subtitle="Pigsty 基础设施仓库" icon="sparkles" />}}
{{< /cards >}}

当前活跃目标为 **PG15–19**，共 **16 个 Linux 平台**（其中 EL8 覆盖 PG15–18）。


## 快速上手 

你可以使用 [**`pig`**](https://pig.pgsty.com/zh) 命令行工具来方便的启用 infra 和 pgsql 仓库，

{{< tabs group="mirror" default="mirror" >}}

{{< tab label="默认" value="global" >}}
```bash tab="pig"
curl https://repo.pigsty.io/pig | bash      # 下载安装最新版本的 pig 命令行工具
pig repo add all -u                         # 添加 Linux / PGDG / Pigsty 仓库并更新缓存
```
{{< /tab >}}

{{< tab label="镜像" value="mirror" >}}
```bash tab="pig"
curl https://repo.pigsty.cc/pig | bash      # 从中国镜像站下载安装最新版本的 pig 命令行工具
pig repo add -u                             # 添加 Linux / PGDG / Pigsty 仓库并更新缓存
```
{{< /tab >}}
{{< /tabs >}}


## 手工添加

您也可以使用经典的 `apt` / `dnf` / `yum` 命令 ，将它们 [手工添加](#手工添加) 到系统中。

{{< tabs group="pkgmgr" >}}
{{< tab label="APT" value="apt" >}}
```bash tab="apt"
# 将 Pigsty 的 GPG 公钥添加到系统密钥链中，以验证软件包签名
curl -fsSL https://repo.pigsty.io/key | sudo gpg --dearmor -o /etc/apt/keyrings/pigsty.gpg

# 获取 Debian / Ubuntu 发行版代号（bookworm、trixie、jammy、noble、resolute），并将对应仓库地址写入 /etc/apt/sources.list.d/ 中
distro_codename=$(lsb_release -cs)
sudo tee /etc/apt/sources.list.d/pigsty-io.list > /dev/null <<EOF
deb [signed-by=/etc/apt/keyrings/pigsty.gpg] https://repo.pigsty.io/apt/infra generic main
deb [signed-by=/etc/apt/keyrings/pigsty.gpg] https://repo.pigsty.io/apt/pgsql/${distro_codename} ${distro_codename} main
EOF

# 刷新 APT 仓库缓存
sudo apt update
```
{{< /tab >}}

{{< tab label="DNF/YUM" value="dnf" >}}
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
{{< /tab >}}
{{< /tabs >}}

所有的 RPM / DEB 软件包都使用指纹为 (`B9BD8B20`) 的 [GPG 密钥](/zh/repo/gpg) 进行签名，以确保软件包的完整性。




## 兼容性

Pigsty 仓库由两个主要部分组成：[`INFRA`](/zh/repo/infra) 和 [`PGSQL`](/zh/repo/pgsql)，提供 `x86_64` 和 `aarch64` 架构下的 DEB / RPM 包。

[`INFRA`](/zh/repo/infra) 仓库中的软件包与 PostgreSQL / Linux 大版本无关，包括 Prometheus、Grafana、以及一些 PostgreSQL 管理工具，
通常由 Go 等语言编写的，只有芯片架构（`x86_64` | `aarch64`）的区别。

| Linux  |  软件包  |                 x86_64                  |                 aarch64                 |
|:------:|:-----:|:---------------------------------------:|:---------------------------------------:|
|   EL   | `rpm` | {{< badge content="✓" color="green" >}} | {{< badge content="✓" color="green" >}} |
| Debian | `deb` | {{< badge content="✓" color="green" >}} | {{< badge content="✓" color="green" >}} |

[`PGSQL`](/zh/repo/pgsql) 仓库中的软件包通常特定于 Linux 大版本（例如 el9, d12），也通常与 PostgreSQL 大版本相关（例如 pg17，pg16 ）
这个仓库中包含了 PostgreSQL 内核分支、扩展插件与工具，通常由类 C 语言编写。

{{< pgext-platforms >}}
{.wide .expand}



## 源代码

用于构建仓库内软件的源代码文件位于以下仓库中：

- https://github.com/pgsty/rpm
- https://github.com/pgsty/deb
- https://github.com/pgsty/infra-pkg
