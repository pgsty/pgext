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
  pig sty boot    [-rmpk]         # native controller bootstrap
  pig sty conf    [mode] [flags]  # native Inventory-aware configuration
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
| `sty boot`    | 原生引导 Pigsty 控制节点       | 需要 root 权限       |
| `sty conf`    | 原生生成并校验 Inventory      | Go 工作流           |
| `sty deploy`  | 运行部署 playbook          |                   |
| `sty list`    | 列出可用 Pigsty 版本         |                   |
| `sty get`     | 下载 Pigsty 源码压缩包        |                   |
| `sty grafana` | 管理 Grafana 仪表盘（别名 `gf`）| v1.6.0 新增         |
{.full-width}

> v1.8.0 起，`pig sty boot` 与 `pig sty conf` 均由 Go 原生实现，不再调用 Pigsty
> 旧版 `bootstrap` / `configure` Shell 脚本。v1.6.0 起，原先的
> `pig sty edit` / `validate` / `check` 已上移为根级
> [`pig inventory`](https://pig.pgsty.com/zh/inventory/) 命令组；实验性的 `pig sty dashboard` 由
> `pig sty grafana` 取代。

## 快速入门

你可以使用 `pig sty` 子命令在当前节点引导部署 Pigsty。

```bash
sudo pig sty boot                # 准备控制节点与 ~/pigsty
pig sty conf -g                  # 生成并校验 pigsty.yml
pig inventory edit              # 可选：检查并调整 Inventory
pig sty deploy                   # 运行部署 playbook
```

详细入门指南请参阅：https://pigsty.cc/docs/setup/install/

`sty boot` 会以尽力而为的方式初始化缺失的默认 `~/pigsty` 目录。若需要指定 Pigsty
版本或安装路径，请先显式执行 `pig sty init`。

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

使用 Go 原生工作流引导 Pigsty 控制节点。该命令需要 root 权限，能够准备可用的 Ansible
环境、处理在线与离线仓库、修复常见控制节点前置条件，并返回结构化结果；整个过程不再委托给
Pigsty 旧版 `bootstrap` 脚本。

```bash
sudo pig sty boot                         # 使用默认区域在线引导
sudo pig sty boot -r china                # 使用中国区域仓库
sudo pig sty boot -m                      # 等价于 --region china
sudo pig sty boot -k                      # 保留现有仓库定义
sudo pig sty boot -p /path/to/pkg.tgz     # 使用显式离线包
sudo pig sty boot -p https://host/pkg.tgz # 下载并使用离线包
sudo pig sty boot -o json                 # 返回机器可读的结果与告警
```

原生引导依次完成：

1. 在 Debian 12/13 上尽可能检查并修复 `en_US.UTF-8`，避免 Ansible 因继承到坏 locale 而无法启动。
2. 同时校验 `ansible-playbook` 与 Python 依赖，而不是只检查二进制文件是否存在。
3. 复用已提交的 `/www/pigsty` 仓库，使用显式软件包或 URL，接受可信的自动
   `/tmp/pkg.tgz`，或配置所选区域的在线仓库。
4. Ansible 缺失或不可用时安装控制节点软件包；即使 Ansible 已安装，显式离线输入仍会被准备。
5. 未指定 `--keep` 时，在替换仓库定义前创建备份；本地或在线安装失败时自动恢复。
6. 尽力修复本机免密 SSH，并初始化缺失的默认 `~/pigsty`。这些便利步骤失败时给出告警，
   不会把已经可用的控制节点误判为引导失败。

**选项：**

- `-r|--region`：区域（default, china, europe...）
- `-m|--mirror`：等价于 `--region china`
- `-p|--path`：离线包文件或 HTTP(S) URL；显式指定的来源无效时直接失败
- `-k|--keep`：保留现有仓库定义，不执行替换

自动发现的 `/tmp/pkg.tgz` 必须是普通文件，不可被组或其他用户写入，且属主为 root 或发起
sudo 的用户；不安全的自动候选会被忽略并产生告警。使用 `-o json` 或 `-o yaml` 可直接获取
引导模式、仓库策略、回滚状态、locale、SSH、Pigsty 初始化状态、告警与后续命令。

详见：https://pigsty.cc/docs/setup/offline/#bootstrap

## sty conf

使用 Go 原生工作流生成 Pigsty Inventory。`sty conf` 从 `<PIGSTY_HOME>/conf` 下读取一个模板，
执行有边界的结构化变更，校验完整候选配置，最后原子写入仅属主可读的 Inventory；它不会调用或
回退到 `./configure`。

```bash
pig sty conf                         # 使用 conf/meta.yml，写入 pigsty.yml
pig sty conf -g                      # 生成随机密码（推荐）
pig sty conf rich                    # 位置参数选择 conf/rich.yml
pig sty conf -c ha/full              # 等价的参数形式；不要与位置参数并用
pig sty conf ha/trio --ip 10.0.0.10,10.0.0.11,10.0.0.12
pig sty conf --domain infra.example.com
pig sty conf rich -v 18              # 为通用模板请求 PostgreSQL 18
pig sty conf -r china -s             # 中国区域、保留占位 IP、跳过管理员预检
pig sty conf -x                       # 写入代理环境变量
pig sty conf full -g -O ha.yml       # 写入自定义的仅属主可读文件
pig sty conf -n --ip 10.0.0.10 -o json
```

**选项：**

- `-c|--conf`：模板模式，等价于位置参数 `[mode]`，两种形式不能同时使用
- `--ip`：最多十个互不相同、逗号分隔的 IPv4 地址
- `--domain`：替换精确的 `i.pigsty` 占位域名
- `-v|--version`：PostgreSQL 主版本（18/17/16/15/14；19 beta 可显式指定）
- `-r|--region`：上游仓库区域（default/china/europe）
- `-m|--mirror`：等价于 `--region china`
- `-O|--output-file`：输出配置文件路径（默认：pigsty.yml）
- `-s|--skip`：保留占位 IP 并跳过管理员 SSH/sudo 预检；不能与 `--ip` 同时使用
- `-p|--port`：SSH 端口
- `-x|--proxy`：将非空代理环境变量写入 `all.vars.proxy_env`
- `-n|--non-interactive`：IP 候选不唯一时拒绝猜测，不进入交互选择
- `-g|--generate`：将已知演示口令替换为 24 位随机值

`--ip` 会按顺序映射模板中的 `10.10.10.10` 至 `10.10.10.19` 占位槽；替换是同时完成的，
VIP 等无关地址保持不变。未指定 `--ip` 时，交互模式会列出探测到的网卡供选择；非交互或输入已
关闭时会明确提示指定地址。`--domain` 只替换精确的 `i.pigsty`，不会误改 `cli.pigsty` 或
`i.pigsty.cc`。

模板模式只能是 `conf` 下的安全相对路径，绝对路径、目录穿越与路径逃逸都会被拒绝。输出文件
不能通过直接路径、符号链接、带符号链接的父目录或硬链接指回源模板。渲染后的 Inventory 会在
原子 `0600` 写入前完成校验，因此解析、变更、预检或校验失败都不会覆盖目标文件。结构化输出会
报告 IP 映射、实际 PostgreSQL 版本、生成的机密标识符与告警，但绝不会暴露随机口令值。

详见：https://pigsty.cc/docs/setup/install/#配置

## sty deploy

使用 deploy.yml 剧本部署 Pigsty。

```bash
pig sty deploy       # 执行 deploy.yml（如果找不到则使用 install.yml）
pig sty d            # 短别名
pig sty de           # 短别名
```

此命令从您的 Pigsty 安装目录执行 deploy.yml 剧本。为保持向后兼容性，如果 deploy.yml 不存在但 install.yml 存在，将使用 install.yml 代替。

> **警告**：此操作会修改您的系统，且 **调用即执行**——deploy 不设 `--yes` 确认门，
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

### 传统仪表盘与 schema v2 资源

`load` 与 `init` 同时接受传统 Grafana 仪表盘 JSON，以及具有以下精确身份的资源格式：

```json
{"apiVersion":"dashboard.grafana.app/v2","kind":"Dashboard","metadata":{"name":"pgsql-overview","namespace":"default"},"spec":{}}
```

加载时，PIG 不会把两种格式悄悄压平为同一种：

- 传统 JSON 从顶层 `uid` 取得身份，并调用旧版 dashboard API。
- Schema v2 从 `metadata.name` 取得 UID；缺少 namespace 时默认使用 `default`，`spec` 必须是对象，并通过 Grafana dashboard resource API 写入。
- JSON 文件名去掉 `.json` 后必须与解析出的 UID 一致。本地目录只允许一层文件夹，其目录名会成为 Grafana folder UID。
- 对 schema v2，PIG 保留 `spec` 与 `grafana.app/message` 注解，根据本地文件夹写入 `grafana.app/folder`，并在 upsert 前主动去掉由服务端管理的 metadata/status。
- `dump` 只有在目标文件已经以 v2 形式存在时才保持 schema v2；此时会使用该文件的 namespace 拉取原生 v2 资源。全新导出目标默认写成传统 JSON；仅存在于本地的文件不会被 `dump` 删除。

其他 `dashboard.grafana.app/*` 版本或结构不完整的资源封装会被直接拒绝，不会被静默当成传统仪表盘。因此，要往返保持 v2 格式，必须保留已有的本地 v2 文件作为格式契约。
