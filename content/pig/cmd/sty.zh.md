---
title: "CMD: sty"
description: 如何使用 pig sty 子命令管理 Pigsty
icon: SquareTerminal
weight: 630
---


**pig** 也可作为 Pigsty 的命令行工具使用 —— 这是一款开箱即用的免费 PostgreSQL RDS 解决方案。
它为你的 PostgreSQL 集群带来高可用（HA）、PITR、监控、基础设施即代码（IaC）以及丰富的扩展支持。

```bash
pig sty - Init (Download), Bootstrap, Configure, and Deploy Pigsty

  pig sty init    [-pfvd]         # install pigsty (~/pigsty by default)
  pig sty boot    [-rpk]          # install ansible and prepare offline pkg
  pig sty conf    [-cvrsoxnpg]    # configure pigsty and generate config
  pig sty deploy                  # use pigsty to deploy everything (CAUTION!)
  pig sty get                     # download pigsty source tarball
  pig sty list                    # list available pigsty versions

Examples:
  pig sty init                 # extract and init ~/pigsty
  pig sty boot                 # install ansible & other deps
  pig sty conf                 # generate pigsty.yml config file
  pig sty deploy               # run the deploy.yml playbook
```

| 命令 | 描述 | 备注 |
|:---|:---|:---|
| `sty init` | 安装 Pigsty | |
| `sty boot` | 安装 Ansible 依赖 | 需要 sudo 或 root 权限 |
| `sty conf` | 生成配置 | |
| `sty deploy` | 运行部署 playbook | |
| `sty list` | 列出可用 Pigsty 版本 | |
| `sty get` | 下载 Pigsty 源码压缩包 | |
{.full-width}


## 快速入门

你可以使用 `pig sty` 子命令在当前节点引导部署 Pigsty。

```bash
pig sty init                     # 安装 Pigsty 到 ~/pigsty
pig sty boot                     # 安装 Ansible 依赖
pig sty conf                     # 生成配置
pig sty deploy                   # 运行部署 playbook
```

详细入门指南请参阅：https://pigsty.io/docs/setup/install/


## sty init

下载并安装 Pigsty 发行版到 `~/pigsty` 目录。

```bash
pig sty init                   # 使用最新版本安装到 ~/pigsty
pig sty init -f                # 安装并覆盖已有 pigsty 目录
pig sty init -p /tmp/pigsty    # 安装到指定目录 /tmp/pigsty
pig sty init -v 3.4            # 获取并安装指定版本 v3.4.1
pig sty init 3                 # 获取并安装指定主版本 v3 最新
```

**选项：**
- `-p|--path`：目标安装目录（默认 "~/pigsty"）
- `-f|--force`：强制覆盖已存在的 pigsty 目录
- `-v|--version`：pigsty 版本号
- `-d|--dir`：下载目录（默认 "/tmp"）


## sty boot

安装 Ansible 及其依赖。

```bash
pig sty boot                     # 安装 Ansible
pig sty boot -r china            # 使用中国区域镜像
pig sty boot -k                  # 保留已有仓库
pig sty boot -p /path/to/pkg     # 指定离线包路径
```

**选项：**
- `-r|--region`：区域（default, china, europe...）
- `-p|--path`：离线包路径
- `-k|--keep`：保留已有仓库

详见：https://pigsty.io/zh/docs/setup/offline/#bootstrap


## sty conf

使用 ./configure 配置 Pigsty，生成配置文件。

```bash
pig sty conf                       # 使用默认 meta.yml 配置
pig sty conf -g                    # 生成随机密码（推荐！）
pig sty conf -c rich               # 使用 conf/rich.yml 模板（包含更多扩展）
pig sty conf -c ha/full            # 使用 conf/ha/full.yml 4 节点高可用模板
pig sty conf -c slim               # 使用 conf/slim.yml 模板（最小化安装）
pig sty conf -c supabase           # 使用 conf/supabase.yml 模板（自托管）
pig sty conf -v 17 -c rich         # 使用 conf/rich.yml 模板，PostgreSQL 17
pig sty conf -r china -s           # 使用中国区镜像源，跳过 IP 探测
pig sty conf -x                    # 从环境变量写入代理配置到配置文件
pig sty conf -c full -g -o ha.yml  # 完整 HA 模板，随机密码输出到 ha.yml
```

**选项：**
- `-c|--conf`：配置模板名称（meta/rich/slim/full/supabase/...）
- `--ip`：主节点 IP 地址
- `-v|--version`：PostgreSQL 主版本（18/17/16/15/14/13）
- `-r|--region`：上游仓库区域（default/china/europe）
- `-o|--output`：输出配置文件路径（默认：pigsty.yml）
- `-s|--skip`：跳过 IP 探测
- `-p|--port`：SSH 端口
- `-x|--proxy`：从环境变量写入代理配置
- `-n|--non-interactive`：非交互模式
- `-g|--generate`：生成随机默认密码（推荐！）

详见：https://pigsty.io/docs/setup/install/#configure


## sty deploy

使用 deploy.yml 剧本部署 Pigsty。

```bash
pig sty deploy       # 执行 deploy.yml（如果找不到则使用 install.yml）
pig sty install      # 与 deploy 相同（向后兼容）
pig sty d            # 短别名
pig sty de           # 短别名
pig sty ins          # 短别名
```

此命令从您的 Pigsty 安装目录执行 deploy.yml 剧本。为保持向后兼容性，如果 deploy.yml 不存在但 install.yml 存在，将使用 install.yml 代替。

> **警告**：此操作会修改您的系统。请谨慎使用！


## sty list

列出可用的 Pigsty 版本。

```bash
pig sty list                     # 列出可用版本
```


## sty get

下载 Pigsty 源码压缩包。

```bash
pig sty get                      # 下载最新版本
pig sty get v3.4.0               # 下载指定版本
```
