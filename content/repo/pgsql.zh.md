---
title: PGSQL 仓库
description: PostgreSQL 扩展和内核分支仓库
icon: Database
weight: 100
---

`pigsty-pgsql` 仓库包含特定于 PostgreSQL 主版本的软件包。（通常也特定于特定的 Linux 发行版主版本）包括扩展和一些内核分支。

您可以参考 [发布 - RPM 变更日志](/zh/release/rpm) / [发布 - DEB 变更日志](/zh/release/deb) 查阅最近的版本更新情况。

|   系统 / 架构    |  OS  |                                                                                                                   x86_64                                                                                                                   |                                                                                                                  aarch64                                                                                                                   |
|:------------:|:----:|:------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------:|:------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------:|
|     EL8      | el8  | <Badge variant="blue-subtle">18</Badge><Badge variant="blue-subtle">17</Badge><Badge variant="blue-subtle">16</Badge><Badge variant="blue-subtle">15</Badge><Badge variant="blue-subtle">14</Badge><Badge variant="blue-subtle">13</Badge> | <Badge variant="blue-subtle">18</Badge><Badge variant="blue-subtle">17</Badge><Badge variant="blue-subtle">16</Badge><Badge variant="blue-subtle">15</Badge><Badge variant="blue-subtle">14</Badge><Badge variant="blue-subtle">13</Badge> |
|     EL9      | el9  | <Badge variant="blue-subtle">18</Badge><Badge variant="blue-subtle">17</Badge><Badge variant="blue-subtle">16</Badge><Badge variant="blue-subtle">15</Badge><Badge variant="blue-subtle">14</Badge><Badge variant="blue-subtle">13</Badge> | <Badge variant="blue-subtle">18</Badge><Badge variant="blue-subtle">17</Badge><Badge variant="blue-subtle">16</Badge><Badge variant="blue-subtle">15</Badge><Badge variant="blue-subtle">14</Badge><Badge variant="blue-subtle">13</Badge> |
|     EL10     | el10 | <Badge variant="blue-subtle">18</Badge><Badge variant="blue-subtle">17</Badge><Badge variant="blue-subtle">16</Badge><Badge variant="blue-subtle">15</Badge><Badge variant="blue-subtle">14</Badge><Badge variant="blue-subtle">13</Badge> | <Badge variant="blue-subtle">18</Badge><Badge variant="blue-subtle">17</Badge><Badge variant="blue-subtle">16</Badge><Badge variant="blue-subtle">15</Badge><Badge variant="blue-subtle">14</Badge><Badge variant="blue-subtle">13</Badge> |
|  Debian 12   | d12  | <Badge variant="blue-subtle">18</Badge><Badge variant="blue-subtle">17</Badge><Badge variant="blue-subtle">16</Badge><Badge variant="blue-subtle">15</Badge><Badge variant="blue-subtle">14</Badge><Badge variant="blue-subtle">13</Badge> | <Badge variant="blue-subtle">18</Badge><Badge variant="blue-subtle">17</Badge><Badge variant="blue-subtle">16</Badge><Badge variant="blue-subtle">15</Badge><Badge variant="blue-subtle">14</Badge><Badge variant="blue-subtle">13</Badge> |
|  Debian 13   | d13  | <Badge variant="blue-subtle">18</Badge><Badge variant="blue-subtle">17</Badge><Badge variant="blue-subtle">16</Badge><Badge variant="blue-subtle">15</Badge><Badge variant="blue-subtle">14</Badge><Badge variant="blue-subtle">13</Badge> | <Badge variant="blue-subtle">18</Badge><Badge variant="blue-subtle">17</Badge><Badge variant="blue-subtle">16</Badge><Badge variant="blue-subtle">15</Badge><Badge variant="blue-subtle">14</Badge><Badge variant="blue-subtle">13</Badge> |
| Ubuntu 22.04 | u22  | <Badge variant="blue-subtle">18</Badge><Badge variant="blue-subtle">17</Badge><Badge variant="blue-subtle">16</Badge><Badge variant="blue-subtle">15</Badge><Badge variant="blue-subtle">14</Badge><Badge variant="blue-subtle">13</Badge> | <Badge variant="blue-subtle">18</Badge><Badge variant="blue-subtle">17</Badge><Badge variant="blue-subtle">16</Badge><Badge variant="blue-subtle">15</Badge><Badge variant="blue-subtle">14</Badge><Badge variant="blue-subtle">13</Badge> |
| Ubuntu 24.04 | u24  | <Badge variant="blue-subtle">18</Badge><Badge variant="blue-subtle">17</Badge><Badge variant="blue-subtle">16</Badge><Badge variant="blue-subtle">15</Badge><Badge variant="blue-subtle">14</Badge><Badge variant="blue-subtle">13</Badge> | <Badge variant="blue-subtle">18</Badge><Badge variant="blue-subtle">17</Badge><Badge variant="blue-subtle">16</Badge><Badge variant="blue-subtle">15</Badge><Badge variant="blue-subtle">14</Badge><Badge variant="blue-subtle">13</Badge> |


## 快速上手

### PIG

您可以安装 [pig](/zh/pig) - CLI 工具，并使用它添加 pgdg / pigsty 仓库（推荐）：

```bash tab="pig"
pig repo add pigsty                         # 添加 pigsty-pgsql 仓库
pig repo add pigsty -u                      # 添加 pigsty-pgsql 仓库，并更新本地缓存
pig repo add pigsty -u --region=default     # 强制使用全球默认区域的仓库（pigsty.io）
pig repo add pigsty -u --region=china       # 使用中国镜像仓库 (pigsty.cc)
pig repo add pgsql -u                       # pgsql = pgdg + pigsty-pgsql (同时添加 Pigsty 与 PGDG 官方仓库)
pig repo add -u                             # all = node + pgsql (pgdg + pigsty) + infra，一次性添加所有仓库
```

> [!TIP] 如果您在中国大陆区域，可以考虑使用中国 CDN 镜像 (将 pigsty.io 替换为 pigsty.cc)

### APT

您也可以直接在 Debian / Ubuntu 上使用 `apt` 启用此仓库：

```bash tab="apt"
# 将 Pigsty 的 GPG 公钥添加到您的系统密钥链以验证包签名
curl -fsSL https://repo.pigsty.io/key | sudo gpg --dearmor -o /etc/apt/keyrings/pigsty.gpg

# 获取 Debian 发行版代号（distro_codename=jammy, focal, bullseye, bookworm），并将相应的上游仓库地址写入 APT List 文件
distro_codename=$(lsb_release -cs)
sudo tee /etc/apt/sources.list.d/pigsty-io.list > /dev/null <<EOF
deb [signed-by=/etc/apt/keyrings/pigsty.gpg] https://repo.pigsty.io/apt/pgsql/${distro_codename} ${distro_codename} main
EOF

# 刷新 APT 仓库缓存
sudo apt update
```


### DNF

您也可以直接在兼容 EL 的系统上使用 `dnf`/`yum` 启用此仓库：

```bash tab="dnf"
# 将 Pigsty 的 GPG 公钥添加到您的系统密钥链以验证包签名
curl -fsSL https://repo.pigsty.io/key | sudo tee /etc/pki/rpm-gpg/RPM-GPG-KEY-pigsty >/dev/null

# 将 Pigsty 仓库定义文件添加到 /etc/yum.repos.d/ 目录，包括两个仓库
sudo tee /etc/yum.repos.d/pigsty-pgsql.repo > /dev/null <<-'EOF'
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

# 刷新 YUM/DNF 仓库缓存
sudo yum makecache;
```


--------

## 源代码

此仓库的构建脚本与源代码在 GitHub 上开源：

- https://github.com/pgsty/rpm
- https://github.com/pgsty/deb

如果平台不受支持，您也可以自行从源代码构建软件包。

[`pig build`](/zh/pig/cmd/build) 提供了简便的扩展构建基础设施，您可以方便的在自己的系统上制作 RPM/DEB 包。