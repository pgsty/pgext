---
title: 下载
linkTitle: 下载
description: 获取 pig 包管理器，添加软件仓库，安装扩展
weight: 45
icon: fa-solid fa-download
sidebar_enabled: false
breadcrumb: false
page_width: wide
---

这里有三样东西可以下载，顺序就是你会用到它们的顺序：`pig` 包管理器、它负责配置的两个软件仓库，以及扩展软件包本身。

## 一 · 包管理器 {#pig}

`pig` 是一个静态编译的 Go 二进制，没有任何运行时依赖。按你管理这台机器的习惯，挑一个渠道即可。

{{< download pig >}}

## 二 · 软件仓库 {#repositories}

`pig repo add all -u` 会替你写好这些仓库文件，但它们本身就是普通的 APT / DNF 仓库，手工添加同样可以。

{{< cards cols=4 >}}
{{< card link="/repo/pgsql" title="PGSQL 仓库" icon="database" subtitle="576 个已打包扩展，按发行版代号各成一路" />}}
{{< card link="/repo/infra" title="INFRA 仓库" icon="cube" subtitle="与操作系统无关的基础设施软件包，pig 本身也在其中" />}}
{{< card link="/repo/pgdg" title="PGDG 仓库" icon="cloud-download" subtitle="这些扩展所针对构建的上游 PostgreSQL 内核仓库" />}}
{{< card link="/repo/gpg" title="GPG 密钥" icon="key" subtitle="签名公钥、指纹，以及如何校验一个软件包" />}}
{{< /cards >}}

完整的配置步骤（含中国大陆镜像）在 [软件仓库](/repo) 页面。

## 三 · 扩展软件包 {#packages}

每个扩展页面都自带一张下载表：一行一个平台与 PostgreSQL 大版本组合，给出确切的文件名、体积，以及指向仓库的直链。

{{< cards cols=4 >}}
{{< card link="/e" title="扩展清单" icon="clipboard-list" subtitle="全部 576 个已打包扩展，各自带有分平台下载表" />}}
{{< card link="/os" title="按平台" icon="server" subtitle="16 个 Linux 目标上分别有哪些扩展可用" />}}
{{< card link="/os/matrix" title="可用性矩阵" icon="view-grid" subtitle="所有软件包 × 所有操作系统 × 所有大版本" />}}
{{< card link="/matrix/pgext-global-matrix.csv" title="矩阵 CSV" icon="document-download" subtitle="同一张网格的原始数据，另有 JSON 导出" />}}
{{< /cards >}}

## 校验下载内容 {#verify}

仓库中的软件包都经过 GPG 签名，装好 [公钥](/repo/gpg) 之后， `apt` 与 `dnf` 会自动替你校验。GitHub 上的发布产物则由每个标签随附的 `checksums.txt` 覆盖：

```bash
curl -fsSLO https://github.com/pgsty/pig/releases/latest/download/checksums.txt
sha256sum -c checksums.txt --ignore-missing
```

## 源代码 {#source}

这里没有任何东西是从你看不到的代码里构建出来的：

| 仓库 | 构建产物 |
|:-----|:---------|
| [pgsty/pig](https://github.com/pgsty/pig) | 包管理器本身 |
| [pgsty/rpm](https://github.com/pgsty/rpm) | PGSQL 仓库中的每一个 RPM |
| [pgsty/deb](https://github.com/pgsty/deb) | PGSQL 仓库中的每一个 DEB |
| [pgsty/infra-pkg](https://github.com/pgsty/infra-pkg) | 与操作系统无关的基础设施软件包 |
