---
title: "CMD: sty"
description: 如何使用 pig sty 子命令管理 Pigsty
icon: SquareTerminal
weight: 630
---


**pig** 也可作为 Pigsty 的命令行工具使用 —— 这是一款开箱即用的免费 PostgreSQL RDS 解决方案。
它为你的 PostgreSQL 集群带来高可用（HA）、PITR、监控、基础设施即代码（IaC）以及丰富的扩展支持。

```bash
pig sty - 初始化（下载）、引导、配置与部署 Pigsty

  pig sty init    [-pfvd]         # 安装 pigsty（默认目录 ~/pigsty）
  pig sty boot    [-rpk]          # 安装 ansible 并准备离线包
  pig sty conf    [-cvrsoxnpg]   # 配置 pigsty 并生成配置文件
  pig sty deploy                  # 使用 pigsty 部署所有组件（注意！）
  pig sty get                     # 下载 pigsty 源码压缩包
  pig sty list                    # 列出可用 pigsty 版本

用法：
  pig sty [命令]

别名：
  sty, s, pigsty

示例：
  入门指南：https://pigsty.io/docs/setup/install/
  pig sty init                 # 解压并初始化 ~/pigsty
  pig sty boot                 # 安装 ansible 及其他依赖
  pig sty conf                 # 生成 pigsty.yml 配置文件
  pig sty deploy               # 执行 deploy.yml 剧本

可用命令：
  boot        引导 Pigsty
  conf        配置 Pigsty
  deploy      执行 pigsty deploy.yml 剧本
  get         下载 pigsty 可用版本
  init        安装 Pigsty
  list        列出 pigsty 可用版本

参数：
  -h, --help   获取 sty 帮助
```

你可以使用 `pig sty` 子命令在当前节点引导部署 Pigsty。



## `sty init`

```bash
pig sty init
  -p | --path    : 安装路径，默认为 ~/pigsty
  -f | --force   : 强制覆盖已存在的 pigsty 目录
  -v | --version : pigsty 版本，默认使用最新版本
  -d | --dir     : 下载目录，默认为 /tmp

用法：
  pig sty init [参数]

别名：
  init, i

示例：

  pig sty init                   # 使用最新版本安装到 ~/pigsty
  pig sty init -f                # 安装并覆盖已有 pigsty 目录
  pig sty init -p /tmp/pigsty    # 安装到指定目录 /tmp/pigsty
  pig sty init -v 3.4            # 获取并安装指定版本 v3.4.1
  pig sty init 3                 # 获取并安装指定主版本 v3 最新


参数：
  -d, --dir string       pigsty 下载目录（默认 "/tmp"）
  -f, --force            覆盖已存在 pigsty（默认不覆盖）
  -h, --help             获取 init 帮助
  -p, --path string      目标安装目录（默认 "~/pigsty"）
  -v, --version string   pigsty 版本号
```


## `sty boot`

```bash
pig sty boot
  [-r|--region <region>]   [default,china,europe]
  [-p|--path <path>]       指定离线包路径
  [-k|--keep]              引导时保留已有上游仓库

详见：https://pigsty.io/zh/docs/setup/offline/#bootstrap

用法：
  pig sty boot [参数]

别名：
  boot, b, bootstrap

参数：
  -h, --help            获取 boot 帮助
  -k, --keep            保留已有仓库
  -p, --path string     离线包路径
  -r, --region string   区域（default, china, europe...）
```


## `sty conf`

```bash
使用 ./configure 配置 pigsty

pig sty conf
  [-c|--conf <name>]      # 配置模板：[meta|rich|slim|full|supabase|...]
  [--ip <ip>]             # 主节点 IP 地址（-s 跳过）
  [-v|--version <pgver>]  # postgres 主版本：[18|17|16|15|14|13]
  [-r|--region <region>]  # 上游仓库区域：[default|china|europe]
  [-o|--output <file>]    # 输出配置文件路径（默认：pigsty.yml）
  [-s|--skip]             # 跳过 IP 探测
  [-p|--port <port>]      # 指定 SSH 端口（用于 SSH 可达性检查）
  [-x|--proxy]            # 从环境变量写入代理配置
  [-n|--non-interactive]  # 非交互模式

  [-g|--generate]         # 生成随机默认密码（推荐！）

详见：https://pigsty.io/docs/setup/install/#configure

用法：
  pig sty conf [参数]

别名：
  conf, c, configure

示例：

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


参数：
  -c, --conf string       配置模板名称
  -g, --generate          生成随机密码
  -h, --help              获取 conf 帮助
      --ip string         主节点 IP 地址
  -n, --non-interactive   非交互模式
  -o, --output string     输出配置文件路径
  -p, --port string       SSH 端口（仅在设置时使用）
  -x, --proxy             从环境变量写入代理配置
  -r, --region string     上游仓库区域
  -s, --skip              跳过 IP 探测
  -v, --version string    PostgreSQL 主版本
```


## `sty deploy`

```bash
使用 deploy.yml 剧本部署 Pigsty

此命令从您的 Pigsty 安装目录执行 deploy.yml 剧本。
为保持向后兼容性，如果 deploy.yml 不存在但 install.yml 存在，
将使用 install.yml 代替。

别名：deploy、d、de、install、ins（向后兼容）

警告：此操作会修改您的系统。请谨慎使用！

用法：
  pig sty deploy [参数]

别名：
  deploy, d, de, install, ins

示例：
  pig sty deploy       # 执行 deploy.yml（如果找不到则使用 install.yml）
  pig sty install      # 与 deploy 相同（向后兼容）
  pig sty d            # 短别名
  pig sty de           # 短别名
  pig sty ins          # 短别名

参数：
  -h, --help   获取 deploy 帮助
```
