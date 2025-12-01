---
title: 安装
description: 如何下载与安装 pig 包管理器
icon: Download
weight: 300
breadcrumbs: false
---


## 脚本

安装 `pig` 最简单的方式是运行以下安装脚本：

{{< tabs items="默认,镜像" defaultIndex="1" >}}

{{< tab >}}
```bash
curl -fsSL https://repo.pigsty.io/pig | bash     # 从 Cloudflare 安装
```
{{< /tab >}}

{{< tab >}}
```bash
curl -fsSL https://repo.pigsty.cc/pig | bash     # 从中国 CDN 镜像站安装
```
{{< /tab >}}

{{< /tabs >}}

该脚本会从 Pigsty [软件仓库](/zh/repo/) 下载最新版 `pig` 的 RPM / DEB 包，并通过 `rpm` 或 `dpkg` 进行安装。




## 版本

您可以指定特定版本进行安装，将版本号作为参数传入即可：

{{< tabs items="默认,镜像" defaultIndex="1" >}}

{{< tab >}}
```bash
curl -fsSL https://repo.pigsty.io/pig | bash -s 0.7.4     # 安装 0.7.4 版本
```
{{< /tab >}}

{{< tab >}}
```bash
curl -fsSL https://repo.pigsty.cc/pig | bash -s 0.7.4     # 安装 0.7.4 版本（中国镜像）
```
{{< /tab >}}

{{< /tabs >}}





## 发布页下载

你也可以直接从 Pigsty 仓库下载 `pig` 安装包（`RPM`/`DEB`/ 压缩包）：[GitHub 最新版本发布页](https://github.com/pgsty/pig/releases/latest)

{{< filetree/container >}}
{{< filetree/file name="latest" >}}
{{< filetree/folder name="v0.7.3" state="open" >}}
{{< filetree/file name="https://repo.pigsty.io/pkg/pig/v0.7.4/pig_0.7.4-1_amd64.deb" >}}
{{< filetree/file name="https://repo.pigsty.io/pkg/pig/v0.7.4/pig_0.7.4-1_arm64.deb" >}}
{{< filetree/file name="https://repo.pigsty.io/pkg/pig/v0.7.4/pig-0.7.4-1.aarch64.rpm" >}}
{{< filetree/file name="https://repo.pigsty.io/pkg/pig/v0.7.4/pig-0.7.4-1.x86_64.rpm" >}}
{{< filetree/file name="https://repo.pigsty.io/pkg/pig/v0.7.4/pig-v0.7.4.linux-amd64.tar.gz" >}}
{{< filetree/file name="https://repo.pigsty.io/pkg/pig/v0.7.4/pig-v0.7.4.linux-arm64.tar.gz" >}}
{{< filetree/file name="https://repo.pigsty.io/pkg/pig/v0.7.4/pig-v0.7.4.darwin-amd64.tar.gz]" >}}
{{< filetree/file name="https://repo.pigsty.io/pkg/pig/v0.7.4/pig-v0.7.4.darwin-arm64.tar.gz" >}}

{{< /filetree/folder >}}
{{< filetree/folder name="v0.7.3" state="closed" >}}{{< /filetree/folder >}}
{{< filetree/folder name="......" state="closed" >}}{{< /filetree/folder >}}
{{< /filetree/container >}}

将其解压后，将二进制文件放入您的 PATH 系统路径中即可。


## 仓库

`pig` 软件位于 [`pigsty-infra`](/zh/repo/infra) 仓库中。你可以将该仓库添加到操作系统后，使用操作系统的包管理器进行安装：

### YUM

对于 RHEL，RockyLinux，CentOS，Alma Linux，OracleLinux 等 EL 系发行版：

```bash tab="yum"
sudo tee /etc/yum.repos.d/pigsty-infra.repo > /dev/null <<-'EOF'
[pigsty-infra]
name=Pigsty Infra for $basearch
baseurl=https://repo.pigsty.io/yum/infra/$basearch
enabled = 1
gpgcheck = 0
module_hotfixes=1
EOF

sudo yum makecache;
sudo yum install -y pig
```

### APT

对于 Debian，Ubuntu 等 DEB 系发行版：

```bash tab="apt"
sudo tee /etc/apt/sources.list.d/pigsty-infra.list > /dev/null <<EOF
deb [trusted=yes] https://repo.pigsty.io/apt/infra generic main
EOF

sudo apt update;
sudo apt install -y pig
```



## 更新

若要将现有 `pig` 版本升级至最新可用版本，可以使用以下命令：

```bash
pig update            # 将 pig 自身升级到最新版
```

若要将现有 `pig` 的扩展数据升级至最新可用版本，可以使用以下命令：

```bash
pig ext reload        # 将 pig 扩展数据更新至最新版本
```



## 卸载

```bash
apt remove -y pig     # Debian / Ubuntu 等 Debian 系统
yum remove -y pig     # RHEL / CentOS / RockyLinux 等 EL 系发行版
rm -rf /usr/bin/pig   # 若直接使用二进制安装，删除二进制文件即可
```



## 构建

你也可以自行构建 `pig`。`pig` 使用 Go 语言开发，构建非常容易，源码托管在 [github.com/pgsty/pig](https://github.com/pgsty/pig)

```bash
git clone https://github.com/pgsty/pig.git; cd pig
go get -u; go build
```

所有 RPM / DEB 包都通过 GitHub CI/CD 流程使用 goreleaser 自动化构建。

