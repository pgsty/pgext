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

  pig sty init    [-mpfvd]        # install pigsty (~/pigsty by default)
  pig sty boot    [-rmpk]         # install ansible and prepare offline pkg
  pig sty conf    [-cvmrsoxnpg --raw] # configure pigsty and generate config
  pig sty deploy                  # use pigsty to deploy everything (CAUTION!)
  pig sty get                     # download pigsty source tarball
  pig sty list                    # list available pigsty versions
  pig sty grafana <verb>          # manage grafana dashboards (native HTTP)

Examples:
  pig sty init                 # extract and init ~/pigsty
  pig sty boot                 # install ansible & other deps
  pig sty conf                 # generate pigsty.yml config file
  pig sty deploy               # run the deploy.yml playbook
```

| 命令            | 描述                     | 备注                |
|:--------------|:-----------------------|:------------------|
| `sty init`    | 安装 Pigsty              |                   |
| `sty boot`    | 安装 Ansible 依赖          | 需要 sudo 或 root 权限 |
| `sty conf`    | 生成配置                   |                   |
| `sty deploy`  | 运行部署 playbook          |                   |
| `sty list`    | 列出可用 Pigsty 版本         |                   |
| `sty get`     | 下载 Pigsty 源码压缩包        |                   |
| `sty grafana` | 管理 Grafana 仪表盘（别名 `gf`）| v1.6.0 新增         |
{.full-width}

> v1.6.0 起，原先的 `pig sty edit` / `validate` / `check` 已上移为根级
> [`pig inventory`](/pig/cmd/inventory/) 命令组；实验性的 `pig sty dashboard` 由
> `pig sty grafana` 取代。


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
pig sty init -m                # 优先从 pigsty.cc 镜像下载安装包
pig sty init -p /tmp/pigsty    # 安装到指定目录 /tmp/pigsty
pig sty init -v 3.4            # 获取并安装指定版本 v3.4.1
pig sty init 3                 # 获取并安装指定主版本 v3 最新
```

**选项：**
- `-p|--path`：目标安装目录（默认 "~/pigsty"）
- `-f|--force`：强制覆盖已存在的 pigsty 目录
- `-m|--mirror`：优先使用 `pigsty.cc` 镜像源
- `-v|--version`：pigsty 版本号
- `-d|--dir`：下载目录（默认 "/tmp"）


## sty boot

安装 Ansible 及其依赖。

```bash
pig sty boot                     # 安装 Ansible
pig sty boot -r china            # 使用中国区域镜像
pig sty boot -m                  # 等价于 --region china
pig sty boot -k                  # 保留已有仓库
pig sty boot -p /path/to/pkg     # 指定离线包路径
```

**选项：**
- `-r|--region`：区域（default, china, europe...）
- `-m|--mirror`：等价于 `--region china`
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
pig sty conf -v 18 -c rich         # 使用 conf/rich.yml 模板，PostgreSQL 18
pig sty conf -r china -s           # 使用中国区镜像源，跳过 IP 探测
pig sty conf -m -s                 # 使用镜像模式，跳过 IP 探测
pig sty conf -x                    # 从环境变量写入代理配置到配置文件
pig sty conf -c full -g -O ha.yml  # 完整 HA 模板，随机密码输出到 ha.yml
pig sty conf --raw                 # 使用旧版 shell configure 工作流
```

**选项：**
- `-c|--conf`：配置模板名称（meta/rich/slim/full/supabase/...）
- `--ip`：主节点 IP 地址
- `-v|--version`：PostgreSQL 主版本（18/17/16/15/14；19 beta 可显式指定）
- `-r|--region`：上游仓库区域（default/china/europe）
- `-m|--mirror`：等价于 `--region china`
- `-O|--output-file`：输出配置文件路径（默认：pigsty.yml）
- `-s|--skip`：使用占位 IP，并跳过管理员 SSH/sudo 预检
- `-p|--port`：SSH 端口
- `-x|--proxy`：从环境变量写入代理配置
- `-n|--non-interactive`：非交互模式
- `-g|--generate`：生成随机默认密码（推荐！）
- `--raw`：使用旧版 shell configure 工作流（生成的密码将保持可见）

详见：https://pigsty.io/docs/setup/install/#configure


## sty deploy

使用 deploy.yml 剧本部署 Pigsty。

```bash
pig sty deploy       # 执行 deploy.yml（如果找不到则使用 install.yml）
pig sty d            # 短别名
pig sty de           # 短别名
```

此命令从您的 Pigsty 安装目录执行 deploy.yml 剧本。为保持向后兼容性，如果 deploy.yml 不存在但 install.yml 存在，将使用 install.yml 代替。

> **警告**：此操作会修改您的系统，且**调用即执行**——deploy 不设 `--yes` 确认门，
> 误触发时请用 Ctrl+C 中断。（v1.6.0 起 `pig sty install` / `ins` 别名已移除。）


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
pig sty get -m                   # 优先使用 pigsty.cc 镜像
```


## sty grafana

自 v1.6.0 起，`pig sty grafana`（别名 `gf`）通过 Grafana 原生 HTTP API 管理仪表盘，
取代了实验性的 `pig sty dashboard`。`PATH` 参数可以指向 grafana 根目录、单个文件夹或单个仪表盘
JSON 文件；缺省时解析 `<PIGSTY_HOME>/files/grafana`，不会回退到当前目录。

```bash
pig sty grafana info             # 检查 Grafana 健康、认证与基本信息
pig sty grafana list             # 列出当前组织的全部仪表盘
pig sty grafana boot             # 围绕现有 pigsty 仪表盘引导 Grafana
pig sty grafana init             # 加载完整仪表盘集，然后引导 Grafana
pig sty grafana load [PATH]      # 按本地路径加载仪表盘
pig sty grafana dump [PATH]      # 导出远端仪表盘到本地路径
pig sty grafana clean [PATH]     # 删除本地路径对应的远端仪表盘（--dry-run/--yes）
pig sty grafana lang zh-Hans     # 设置组织与当前用户语言
pig sty grafana style            # 设置组织与当前用户界面风格
```

**连接与凭据**：

| 参数                | 说明                                                    |
|:------------------|:------------------------------------------------------|
| `--endpoint`      | Grafana 地址与路径前缀（默认 `http://i.pigsty/ui`）              |
| `--username`      | Grafana API 用户名                                       |
| `--password`      | Grafana API 密码（**不安全**：对进程列表与 shell 历史可见）             |
| `--password-file` | 仅属主可读的密码文件（推荐）                                        |
{.full-width}

密码解析顺序：`--password` → `--password-file` → `GRAFANA_PASSWORD` 环境变量 →
Inventory 中的 `all.vars.grafana_admin_password`。
HTTP 客户端带有超时与响应大小限制，并拒绝重定向；TLS 证书默认校验。
