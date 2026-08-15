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
| `sty boot`    | 原生引导 Pigsty 控制节点       | 需要时自动提权至 root  |
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
pig sty boot                     # 准备控制节点；需要时自动提权
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

使用 Go 原生工作流引导 Pigsty 控制节点。该命令能够准备可用的 Ansible 环境、处理在线与
离线仓库、修复常见控制节点前置条件，并返回结构化结果；整个过程不再委托给 Pigsty 旧版
`bootstrap` 脚本，下载与解压软件包也不依赖 `curl`、`wget`、`tar` 或 `gzip`。

```bash
sudo pig sty boot                         # 使用默认区域在线引导
sudo pig sty boot -r china                # 使用中国区域仓库
sudo pig sty boot -m                      # 等价于 --region china
sudo pig sty boot -k                      # 保留现有仓库定义
sudo pig sty boot -p /path/to/pkg.tgz     # 使用显式离线包
sudo pig sty boot -p https://host/pkg.tgz # 下载并使用离线包
sudo pig sty boot -o json                 # 返回机器可读的结果与告警
```

命令可以不带 `sudo` 直接调用：Pig 会先解析并下载显式来源，需要 root 权限时再通过 `sudo`
进行一次自重启。设置 `PIG_NO_SUDO=1` 可禁用自动提权；设置 `PIG_NON_INTERACTIVE=1` 可让
sudo 使用非交互模式。

### 引导阶段

原生工作流依次完成：

1. 在 Debian 12/13 上尽可能检查并修复 `en_US.UTF-8`，避免 Ansible 因继承到坏 locale 而无法启动。
2. 实际执行 `ansible-playbook`，发现它使用的 Python 解释器，并校验 `yaml`、`jmespath`，
   以及 `cryptography` 或 `OpenSSL` 两者之一；仅有二进制文件但无法运行，不会被判定为就绪。
3. 解析仓库来源，按需准备离线内容，并且只在 Ansible 缺失或不可用时安装精简的控制节点软件包集。
4. 安装后再次校验 Ansible；如果新软件包补齐了 locale 工具，也会重试 locale 准备。
5. 探测控制节点辅助工具，为发起调用的管理员用户修复到 `127.0.0.1` 的密钥 SSH，并尽可能
   初始化缺失的默认 `~/pigsty` 目录。

即使 Ansible 已经可用，显式指定、自动发现或已经提交的离线来源仍会被准备，因此可以在一个
已经就绪的控制节点上使用 `sty boot` 预置离线仓库。

### 来源选择与工作模式

结果中会记录以下四种引导模式之一：

| 模式 | 含义 |
|------|------|
| `ready` | Ansible 已经可用，也不需要准备离线来源。 |
| `offline` | 选择了显式、可信自动发现或已提交的离线仓库。 |
| `online` | 配置所选区域的在线仓库以修复控制节点。 |
| `existing` | 使用 `--keep` 在线刷新失败后，成功回退到现有仓库定义。 |

来源优先级与安全规则是确定的：

- `--path` 接受本地归档或 HTTP(S) URL。包含凭据的 URL 会被拒绝；显式来源无效时直接失败，
  不会悄悄回退到在线模式。
- 自动发现的 `/tmp/pkg.tgz` 必须是普通文件，不可被组或其他用户写入，且属主为 root 或发起
  sudo 的用户；不安全的候选会被忽略并产生告警。
- 已完整提交的 `/www/pigsty` 仓库优先于选中的离线包；两者同时存在时复用现有仓库，离线包
  保持不动并给出告警。
- Pig 使用 Go 原生能力下载并解压归档。如果 `/www` 不存在，会先创建 `/data/nginx` 与预期的
  `/www -> /data/nginx` 符号链接，再提交仓库内容。

离线模式只启用严格的 `pigsty-local` 仓库；在线模式配置所选区域，安装 Pigsty 内嵌签名密钥并
启用仓库签名校验，同时安装 `node` 与 `pigsty` 控制节点模块。

### 仓库事务与失败边界

默认策略会在替换仓库定义前创建备份。仓库配置或软件包安装失败时，Pig 会尝试恢复备份，并在
结果中明确标记回滚成功或失败。`--keep` 会切换为增量策略：保留现有定义，在线刷新失败时可以
回退到已有仓库，也不需要执行替换回滚。

显式来源无效、需要安装时软件包管理器不受支持、仓库或软件包操作失败，以及安装后 Ansible
仍不可用，都会让命令失败。locale 修复、可选辅助工具探测、本机 SSH 修复与 Pigsty 目录初始化
属于建议性收尾步骤；失败只会作为告警保留，不会否定已经可用的控制节点。

**选项：**

- `-r|--region`：区域（default, china, europe...）
- `-m|--mirror`：等价于 `--region china`；不能与 `--region` 同时使用
- `-p|--path`：离线包文件或 HTTP(S) URL；显式指定的来源无效时直接失败
- `-k|--keep`：保留现有仓库定义，不执行替换

### 结构化输出

自动化场景可使用全局 `-o json` 或 `-o yaml`。结果类型为 `pig.sty.boot/v2`，包含 Ansible
状态、工作模式与软件包管理器、仓库策略与回滚结果、来源与仓库路径、locale、本机 SSH 与
Pigsty 目录初始化状态、是否发生变更、告警，以及以下后续建议：

```bash
pig sty conf -g
pig inv edit
pig sty deploy
```

结构化模式会抑制动态进度信息，保证 stdout 可以直接被程序解析。

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

默认模式为 `meta`；`pig sty c` 与 `pig sty configure` 是命令别名。注意大写 `-O` 用于指定
Inventory 输出文件，全局小写 `-o` 用于选择 text、JSON 或 YAML 命令输出。

### 模板与输出安全

- 模式必须是 `<PIGSTY_HOME>/conf` 下用斜杠分隔的安全相对名称，`.yml` 后缀可省略；绝对路径、
  目录穿越、空路径段与路径逃逸都会被拒绝。
- 相对输出路径基于 `<PIGSTY_HOME>` 解析，绝对输出路径保持不变。
- 目标文件不能通过相同路径、已有符号链接、带符号链接的父目录或硬链接指回源模板；已有输出
  符号链接一律拒绝。
- Pig 会先解析源模板并拒绝冲突的 IP 映射，再执行外部预检。解析、变更、预检或校验失败均不
  会改动目标文件。
- 成功结果以 `0600` 权限原子写入。

### 结构化变更

命令操作解析后的 YAML 结构与有边界的标量，而不是进行宽泛的文本替换：

| 输入 | 原生行为 |
|------|----------|
| `--ip A,B,...` | 最多接收十个互不相同的地址，依次映射到 `10.10.10.10` 至 `10.10.10.19`；替换同时完成，因此地址互换安全，VIP 等无关地址保持不变。 |
| 未指定 `--ip` | 探测本机网卡；候选不唯一时交互选择，`--non-interactive` 或 stdin 已关闭时则失败并提示使用 `--ip`。 |
| `--domain NAME` | 只替换精确的 `i.pigsty`，不会误改 `cli.pigsty` 或 `i.pigsty.cc`；NAME 必须是合法 DNS 域名。 |
| 小规格控制节点 | 探测到 CPU 少于四核时，将 `node_tune: oltp` 与 `pg_conf: oltp.yml` 改为对应的 `tiny` 配置。 |
| `--region REGION` | 非默认区域会更新 `all.vars.region`；`china` 还会启用模板中已有的 Docker 与 pip 镜像值，但不会凭空补造模板中不存在的配置。 |
| `--proxy` | 将非空的 `HTTP_PROXY`/`http_proxy`、`HTTPS_PROXY`（缺失时回退到 `ALL_PROXY`）、`ALL_PROXY` 与 `NO_PROXY` 写入 `all.vars.proxy_env`；必要时补充安全的默认 no-proxy 列表。 |
| `--version MAJOR` | 通用模板支持 PostgreSQL 14-18，以及显式指定的 19 beta，并选择匹配的 locale；版本固定的 `mssql`、`polar` 与 `pgNN` 模式保留模板版本并给出告警。 |
| `--generate` | 每个已知凭据标识符生成一个 24 位随机值，并一致替换其生效值和文档化占位符。 |

如果某个 IP 映射会与未替换的 Inventory 键冲突，命令会按无效参数拒绝执行。某个已提供地址在
模板中没有对应占位槽时，不会被静默忽略，而是作为 discarded-IP 告警保留在结构化结果中。

指定 PostgreSQL 19 beta 时，如果模板包含预期的软件仓库列表，Pig 还会在 `pgsql` 后启用
`beta` 仓库。`conf/build/` 下的模式有意绕过 IP 映射与控制节点管理员预检，以保持构建模板可移植。

生效口令标识符包括 `grafana_admin_password`、`pg_admin_password`、
`pg_monitor_password`、`pg_replication_password`、`patroni_password`、
`haproxy_admin_password`、`minio_secret_key` 与 `etcd_root_password`。随机生成还覆盖文档中的
`DBUser.Meta`、`DBUser.Viewer`、`S3User.Backup`、`S3User.Meta`、`S3User.Data`、
`DBUser.Supa` 和 `Vibe.Coding` 占位符；同一标识符出现多次时会使用同一个生成值。

**选项：**

- `-c|--conf`：模板模式，等价于位置参数 `[mode]`，两种形式不能同时使用
- `--ip`：最多十个互不相同、逗号分隔的 IPv4 地址
- `--domain`：将精确的 `i.pigsty` 占位符替换为合法 DNS 域名
- `-v|--version`：PostgreSQL 主版本（18/17/16/15/14；19 beta 可显式指定）
- `-r|--region`：上游仓库区域（default/china/europe）
- `-m|--mirror`：等价于 `--region china`；不能与 `--region` 同时使用
- `-O|--output-file`：输出配置文件路径（默认：pigsty.yml）
- `-s|--skip`：保留占位 IP 并跳过管理员 SSH/sudo 预检；不能与 `--ip` 同时使用
- `-p|--port`：SSH 端口
- `-x|--proxy`：将非空代理环境变量写入 `all.vars.proxy_env`
- `-n|--non-interactive`：IP 候选不唯一时拒绝猜测，不进入交互选择
- `-g|--generate`：将已知演示口令替换为 24 位随机值

### 预检与校验

未使用 `--skip` 时，Pig 会检查内核、架构、软件包管理器、平台厂商、控制节点资源、sudo/管理员
权限、本机 SSH 与 Ansible 可用性；SSH 检查使用 `--port` 指定的端口。在 Inventory 仍可安全
生成时，这些诊断以可操作告警返回；无效参数与不安全的配置变换仍然是错误。

渲染候选必须通过 Pig 原生 Inventory 校验；如果存在 `ansible-inventory`，还会在提交文件前执行
一次有时间边界的外部解析。`--skip` 会保留占位 IP，并跳过管理员 SSH/sudo 预检，但不会禁用
模板解析、安全变更、Inventory 校验或原子写入。

### 结构化输出

使用全局 `-o json` 或 `-o yaml` 时，结果类型为 `pig.sty.configure/v1`，会报告模式、源模板与
输出路径、区域、所选主地址、已应用与被丢弃的 IP、域名、SSH 端口、请求与实际 PostgreSQL
版本、原生工作流标记、生成的机密标识符及告警。随机口令值绝不会输出。

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
