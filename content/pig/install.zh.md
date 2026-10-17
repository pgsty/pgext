---
title: 安装
description: 如何下载与安装 pig 包管理器
icon: fa-solid fa-download
weight: 300
breadcrumb: false
---

可通过 Linux 安装脚本、系统软件包或发布压缩包安装 PIG。当前版本为 **v{{< param pig_version >}}**，全部发布文件见[下载页](https://pig.pgsty.com/zh/download/)。

## 脚本安装

在 Linux 上执行：

```bash
curl -fsSL https://repo.pigsty.cc/pig | bash
```

中国大陆以外可使用：

```bash
curl -fsSL https://repo.pigsty.io/pig | bash
```

脚本检测 CPU 架构，下载 RPM 或 DEB 包，再调用 `rpm` 或 `dpkg` 安装。安装需要 root 或 `sudo` 权限。

检查安装后的版本：

```bash
pig version
```

指定已发布版本时，将版本号作为参数传入：

```bash
curl -fsSL https://repo.pigsty.cc/pig | bash -s {{< param pig_version >}}
```

## 发布产物下载

在[下载页](https://pig.pgsty.com/zh/download/#packages)或 [GitHub Release](https://github.com/pgsty/pig/releases/tag/v{{< param pig_version >}}) 中选择对应操作系统与架构的文件。

| 系统 | 格式 | 架构名称 |
|:---|:---|:---|
| EL | RPM | `x86_64`、`aarch64` |
| Debian、Ubuntu | DEB | `amd64`、`arm64` |
| Linux | tar.gz | `amd64`、`arm64` |
| macOS | tar.gz | `amd64`（Intel）、`arm64`（Apple Silicon） |

使用发布压缩包时，解压并将 `pig` 可执行文件放入 `PATH` 中的目录。例如，在 Apple Silicon 上执行：

```bash
curl -fLO https://github.com/pgsty/pig/releases/download/v{{< param pig_version >}}/pig-v{{< param pig_version >}}.darwin-arm64.tar.gz
tar -xzf pig-v{{< param pig_version >}}.darwin-arm64.tar.gz
sudo mkdir -p /usr/local/bin
sudo install -m 0755 pig /usr/local/bin/pig
pig version
```

每个版本提供 [SHA-256 校验和](https://github.com/pgsty/pig/releases/download/v{{< param pig_version >}}/checksums.txt)。Linux 用 `sha256sum`，macOS 用 `shasum -a 256` 计算下载文件的哈希，再与清单比对。

macOS 版支持目录查询、配置清单编辑等不依赖 Linux 包管理器的操作。安装 PostgreSQL 软件包请使用[受支持的 Linux 主机](/zh/pig/intro/#linux-兼容性)。

## 仓库安装

要用系统包管理器维护 PIG，先添加 [Pigsty Infra 仓库](https://pigsty.cc/docs/repo/infra/)。按主机选择对应标签页：

```bash {tab="APT (Debian / Ubuntu)" group="install-pig-repository" value="apt" copy="all"}
sudo tee /etc/apt/sources.list.d/pigsty-infra.list > /dev/null <<EOF
deb [trusted=yes] https://repo.pigsty.cc/apt/infra generic main
EOF

sudo apt update
sudo apt install -y pig
```

```bash {tab="DNF (EL)" value="dnf" copy="all"}
sudo tee /etc/yum.repos.d/pigsty-infra.repo > /dev/null <<-'EOF'
[pigsty-infra]
name=Pigsty Infra for $basearch
baseurl=https://repo.pigsty.cc/yum/infra/$basearch
enabled=1
gpgcheck=0
module_hotfixes=1
EOF

sudo dnf makecache
sudo dnf install -y pig
```

这些仓库定义关闭了 Pigsty 仓库的签名校验。密钥处理与校验设置见[仓库密钥](/zh/pig/cmd/repo/#repo-key)。

## 更新

在 Linux 上更新 PIG 程序：

```bash
sudo pig update
```

加 `-m` 使用中国镜像，加 `-v {{< param pig_version >}}` 选择该版本。macOS 压缩包安装请下载新版归档并替换已安装的可执行文件。

单独刷新扩展目录：

```bash
pig ext reload
```

此命令更新目录元数据，已安装的软件包保持不变。

## 卸载

按安装方式选择对应命令：

```bash
sudo apt remove pig              # Debian / Ubuntu 软件包
sudo dnf remove pig              # EL 软件包
sudo rm /usr/local/bin/pig        # 安装在此路径的归档二进制
```

如果归档二进制安装在其他位置，删除对应文件即可。PostgreSQL 及其扩展可以继续运行。

## 构建

安装 Git、Make，以及项目 [go.mod](https://github.com/pgsty/pig/blob/v{{< param pig_version >}}/go.mod) 指定的 Go 工具链。构建当前发布版本：

```bash
git clone --branch v{{< param pig_version >}} https://github.com/pgsty/pig.git
cd pig
go mod download
make build
./pig version
```

`make build` 在当前目录生成适用于本机平台的 `pig`。发布流程通过 GoReleaser 生成 RPM、DEB 和 tar.gz 文件。

接下来按[上手教程](/zh/pig/start/)安装 PostgreSQL 并启用扩展。
