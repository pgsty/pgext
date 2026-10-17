---
title: "CMD: sty"
description: 如何使用 pig sty 子命令管理 Pigsty
icon: fa-solid fa-terminal
weight: 630
---

`pig sty` 安装 Pigsty，准备控制节点，生成配置，并运行部署剧本。它还通过 Grafana HTTP API 管理仪表盘。

| 命令 | 用途 |
|:---|:---|
| `sty init` | 下载并解压 Pigsty 到安装目录 |
| `sty boot` | 安装 Ansible 依赖并准备控制节点 |
| `sty conf` | 从模板生成并校验 Inventory |
| `sty deploy` | 运行部署剧本 |
| `sty list` / `sty get` | 查询或下载 Pigsty 发行版 |
| `sty grafana` | 管理 Grafana 仪表盘、语言和外观；别名 `gf` |

## 快速入门

在控制节点上依次运行：

```bash
pig sty init                     # 下载 Pigsty
pig sty boot                     # 准备 Ansible 与控制节点
pig sty conf -g                  # 生成配置并替换演示密码
pig inventory edit              # 检查目标主机与部署配置
pig sty deploy                   # 执行部署
```

`deploy` 会立即执行剧本，修改 Inventory 选中的主机。完整部署步骤见 [Pigsty 安装指南](https://pigsty.cc/docs/setup/install/)。

`boot` 也会尝试初始化缺失的 `~/pigsty`；需要指定版本或路径时，先执行 `init`。使用 [`pig inventory`](/zh/pig/cmd/inventory/) 编辑和校验已有配置。

## sty init

下载并安装 Pigsty 发行版到 `~/pigsty` 目录。

```bash
pig sty init                   # 使用最新版本安装到 ~/pigsty
pig sty init -f                # 安装并覆盖已有 pigsty 目录
pig sty init -m                # 优先从 pigsty.cc 镜像下载安装包
pig sty init -p /tmp/pigsty    # 安装到指定目录 /tmp/pigsty
pig sty init -v {{< param pigsty_version >}}       # 安装文档对应的 Pigsty 版本
pig sty init 4                 # 安装 4.x 系列的最新版本
```

**选项：**

- `-p|--path`：目标安装目录（默认 "~/pigsty"）
- `-f|--force`：强制覆盖已存在的 pigsty 目录
- `-m|--mirror`：优先使用 `pigsty.cc` 镜像源
- `-v|--version`：pigsty 版本号
- `-d|--dir`：下载目录（默认 `/tmp`）
- `--accept-eula`：非交互接受 Pigsty Pro EULA，仅 Pro 发行版需要

## sty boot

准备 Ansible 及其 Python 依赖，配置在线或离线软件源，并检查控制节点的本机 SSH 和 Pigsty 目录。

```bash
pig sty boot                         # 使用默认区域在线引导
pig sty boot -r china                # 使用中国区域仓库
pig sty boot -m                      # 等价于 --region china
pig sty boot -k                      # 保留现有仓库定义
pig sty boot -p /path/to/pkg.tgz     # 使用显式离线包
pig sty boot -p https://host/pkg.tgz # 下载并使用离线包
pig sty boot -o json                 # 返回机器可读的结果与告警
```

命令可以不带 `sudo` 直接调用：Pig 会先解析并下载显式来源，需要 root 权限时再通过 `sudo`
进行一次自重启。设置 `PIG_NO_SUDO=1` 可禁用自动提权。JSON/YAML 输出及
`PIG_NON_INTERACTIVE=1` 使用非交互 sudo；提权需要密码时直接失败，不等待输入。

### 引导阶段

1. 在 Debian 12/13 上检查并尝试修复 `en_US.UTF-8` locale。
2. 检查 Ansible 能否运行，以及其 Python 环境是否具备 `yaml`、`jmespath` 和 `cryptography` 或 `OpenSSL`。
3. 准备选中的离线仓库；Ansible 不可用时，安装控制节点依赖并再次检查。
4. 检查辅助工具，为当前管理员准备到 `127.0.0.1` 的密钥 SSH，并尝试初始化缺失的 `~/pigsty`。

即使 Ansible 已可用，指定或发现的离线仓库仍会准备，可用此命令为控制节点预置离线软件包。

### 来源选择与工作模式

结果中会记录以下四种引导模式之一：

| 模式 | 含义 |
|------|------|
| `ready` | Ansible 已经可用，也不需要准备离线来源。 |
| `offline` | 选择了显式指定、自动发现或已有的离线仓库。 |
| `online` | 配置所选区域的在线仓库以修复控制节点。 |
| `existing` | 使用 `--keep` 在线刷新失败后，成功回退到现有仓库定义。 |

来源选择规则：

- `--path` 接受本地归档或 HTTP(S) URL。包含凭据的 URL 会被拒绝；显式来源无效时直接失败，
  不会改用在线模式。
- 自动发现的 `/tmp/pkg.tgz` 必须是普通文件，不可被组或其他用户写入，且属主为 root 或发起
  sudo 的用户；不符合条件的文件会被忽略并产生告警。
- 已有完整的 `/www/pigsty` 仓库优先于选中的离线包；两者同时存在时复用现有仓库，离线包
  保持不动并给出告警。
- Pig 使用 Go 原生能力下载并解压归档。如果 `/www` 不存在，会先创建 `/data/nginx` 与预期的
  `/www -> /data/nginx` 符号链接，再提交仓库内容。

离线模式只启用严格的 `pigsty-local` 仓库；在线模式配置所选区域，安装 `node` 与 `pigsty`
控制节点模块，并复用或安装 Pigsty 内嵌签名密钥。公钥准备成功时启用签名校验；默认公钥准备
失败时记录告警，关闭所选 Pigsty 仓库的签名校验并继续。元数据显式指定 `gpgkey` 或
`signed-by` 时，准备该引用实际指定的密钥失败仍会终止，详见 [`pig repo`](/zh/pig/cmd/repo/#信任策略与配置归属)。

### 仓库替换与恢复

默认先备份现有仓库定义，再写入新配置。仓库配置或软件包安装失败时，PIG 尝试恢复备份，并报告恢复结果。`--keep` 保留现有定义；在线刷新失败时，可回退使用原有仓库。

Ansible 安装或验证失败会使命令失败。本机 SSH、locale 或 Pigsty 目录初始化未完成时，命令会报告告警。

**选项：**

- `-r|--region`：区域（default, china, europe...）
- `-m|--mirror`：等价于 `--region china`；不能与 `--region` 同时使用
- `-p|--path`：离线包文件或 HTTP(S) URL；显式指定的来源无效时直接失败
- `-k|--keep`：保留现有仓库定义，不执行替换

### 结构化输出

`-o json` 或 `-o yaml` 返回 `pig.sty.boot/v2` 结果，包括 Ansible 状态、工作模式、软件源、仓库恢复结果、本机 SSH、Pigsty 目录和告警。

离线部署步骤见 [Pigsty 离线安装](https://pigsty.cc/docs/setup/offline/#bootstrap)。

## sty conf

从 `<PIGSTY_HOME>/conf` 下的模板生成 Inventory。命令修改 YAML 配置，校验完整结果，再以 `0600` 权限原子写入输出文件。

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
pig sty conf ha/full -g -O ha.yml       # 写入自定义的仅属主可读文件
pig sty conf -n --ip 10.0.0.10 -o json
```

默认模式为 `meta`；`pig sty c` 与 `pig sty configure` 是命令别名。注意大写 `-O` 用于指定
Inventory 输出文件，全局小写 `-o` 用于选择 text、JSON 或 YAML 命令输出。

### 模板与输出路径

模板名相对于 `<PIGSTY_HOME>/conf`，可省略 `.yml` 后缀，例如 `ha/full`。使用相对名称，不带 `..`。

相对输出路径位于 `<PIGSTY_HOME>` 下，也可指定绝对路径。输出必须与模板是不同文件，且不能是符号链接。解析、变更或校验失败时，原输出文件保持不变。

### 结构化变更

命令对模板应用以下变更：

| 输入 | 配置变更 |
|------|----------|
| `--ip A,B,...` | 用最多十个互不相同的地址，依次替换 `10.10.10.10` 至 `10.10.10.19`。替换同时完成，其他 VIP 保持不变。 |
| 未指定 `--ip` | 探测本机网卡；候选不唯一时交互选择，`--non-interactive` 或 stdin 已关闭时则失败并提示使用 `--ip`。 |
| `--domain NAME` | 只替换精确的 `i.pigsty`，不会误改 `cli.pigsty` 或 `i.pigsty.cc`；NAME 必须是合法 DNS 域名。 |
| 小规格控制节点 | 探测到 CPU 少于四核时，将 `node_tune: oltp` 与 `pg_conf: oltp.yml` 改为对应的 `tiny` 配置。 |
| `--region REGION` | 非默认区域会更新 `all.vars.region`；`china` 还会启用模板中已有的 Docker 与 pip 镜像值。 |
| `--proxy` | 读取 `HTTP_PROXY`、`HTTPS_PROXY`、`ALL_PROXY` 与 `NO_PROXY`，每项都在大写值缺失时回退到对应的小写变量。HTTP/HTTPS 缺失时回退到 ALL_PROXY。将非空值写入 `all.vars.proxy_env`，必要时补充默认 no-proxy 列表。 |
| `--version MAJOR` | 通用模板支持 PostgreSQL 14-18，以及显式指定的 19 beta，并选择匹配的 locale；版本固定的 `mssql`、`polar` 与 `pgNN` 模式保留模板版本并给出告警。 |
| `--generate` | 每个已知凭据标识符生成一个 24 位随机值，并一致替换其生效值和文档化占位符。 |

IP 替换会产生重复键时，命令停止。提供的地址在模板中没有对应占位符时，PIG 会告警。

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

PIG 检查操作系统、软件包管理器、控制节点资源、SSH、sudo 和 Ansible。`--port` 指定 SSH 检查端口。`--skip` 跳过管理员 SSH/sudo 预检并保留占位 IP；其他检查仍会运行。

输出必须通过 Inventory 校验。存在 `ansible-inventory` 时，还会使用它解析结果。参数、模板或校验错误会阻止写入；可以继续生成配置的问题则以告警返回。

### 结构化输出

`-o json` 或 `-o yaml` 返回 `pig.sty.configure/v1`，包括模板、输出路径、区域、IP、域名、SSH 端口、PostgreSQL 版本和告警。生成的密码只写入 Inventory；结果列出对应标识符。

详细步骤见 [Pigsty 配置指南](https://pigsty.cc/docs/setup/install/#配置)。

## sty deploy

`sty deploy` 使用配置层已经选定的 Inventory 来源。`-i` 和 `PIGSTY_INVENTORY` 的相对路径
以调用者工作目录为基准；带引号的 `~/` 按用户主目录展开。选定的绝对路径通过 `-i` 传给 Ansible，
默认来源也一样；Ansible 的工作目录仍为 Pigsty home。来源不存在时失败，不静默回退。
deploy 仍要求普通文件；动态来源的执行仍由 Ansible 负责。

使用 deploy.yml 剧本部署 Pigsty。

```bash
pig sty deploy       # 执行 deploy.yml（如果找不到则使用 install.yml）
pig sty d            # 短别名
pig sty de           # 短别名
```

此命令从您的 Pigsty 安装目录执行 deploy.yml 剧本。为保持向后兼容性，如果 deploy.yml 不存在但 install.yml 存在，将使用 install.yml 代替。

剧本会立即开始执行，不再询问确认。运行前检查所选 Inventory。

## sty list

列出可用的 Pigsty 版本。

```bash
pig sty list                     # 列出近期版本
pig sty list all                 # 列出所有可用版本
pig sty list 4                   # 从最新的 4.x 稳定版起列出版本
```

版本前缀使用与 `sty get` 相同的补全规则，以最高匹配稳定版作为列表起点。
匹配按版本分段区分，`1.1` 不会匹配 `1.10`。精确版本（含预发行版）和 `latest` 的含义不变。

## sty get

下载 Pigsty 源码压缩包。

```bash
pig sty get                      # 下载最新版本
pig sty get {{< param pigsty_version >}}       # 下载文档对应的 Pigsty 版本
pig sty get -m                   # 优先使用 pigsty.cc 镜像
```

## sty grafana

`pig sty grafana`（别名 `gf`）通过 HTTP API 管理仪表盘。可选 `PATH` 指向 Grafana 根目录、其下一层目录或一个仪表盘 JSON 文件。默认路径为 `<PIGSTY_HOME>/files/grafana`。

```bash
pig sty grafana info             # 检查 Grafana 健康、认证与基本信息
pig sty grafana list             # 列出当前组织的全部仪表盘
pig sty grafana boot             # 围绕现有 pigsty 仪表盘引导 Grafana
pig sty grafana init             # 加载完整仪表盘集，然后引导 Grafana
pig sty grafana load [PATH]      # 按本地路径加载仪表盘
pig sty grafana dump [PATH]      # 导出远端仪表盘到本地路径
pig sty grafana clean [PATH]     # 删除本地路径对应的远端仪表盘（--dry-run/--yes）
pig sty grafana lang zh-Hans     # 设置组织与当前用户语言
pig sty grafana style dark       # 设置组织与当前用户界面风格
```

### 选择操作范围

- `list` 列出当前组织的全部仪表盘。
- `load` 按本地文件中的 UID 创建或更新仪表盘，保留其他远端仪表盘。
- `boot` 将当前组织命名为 Pigsty，并将已有的 `pigsty` 仪表盘收藏、设为组织与当前用户的首页。
- `init` 先加载完整仪表盘目录，再执行 `boot`；单个目录或文件使用 `load`。
- `dump -d DIR` 导出到指定目录；`--dir` 与位置参数 `PATH` 不能并用。
- `clean PATH` 删除该本地路径声明的远端仪表盘。`clean --all` 删除当前组织全部仪表盘和目录。
- `style` 接受 `dark`、`light` 或 `auto`。

先预览删除范围，再确认执行：

```bash
pig sty grafana clean --dry-run
pig sty grafana clean --yes
```

**连接与凭据**：

| 参数                | 说明                                                    |
|:------------------|:------------------------------------------------------|
| `--endpoint`      | Grafana 地址与路径前缀（默认 `http://i.pigsty/ui`）              |
| `--username`      | Grafana API 用户名                                       |
| `--password`      | Grafana API 密码（**不安全**：对进程列表与 shell 历史可见）             |
| `--password-file` | 仅属主可读的密码文件（推荐）                                        |
{.full-width}

`--password` 与 `--password-file` 不能同时使用。都未指定时，依次读取 `GRAFANA_PASSWORD` 环境变量和 Inventory 中的 `all.vars.grafana_admin_password`。
HTTP 客户端带有超时与响应大小限制，并拒绝重定向；TLS 证书默认校验。

### 传统仪表盘与 schema v2 资源

`load` 与 `init` 同时接受传统 Grafana 仪表盘 JSON，以及具有以下精确身份的资源格式：

```json
{"apiVersion":"dashboard.grafana.app/v2","kind":"Dashboard","metadata":{"name":"pgsql-overview","namespace":"default"},"spec":{}}
```

两种格式分别使用对应的 API：

- 传统 JSON 从顶层 `uid` 取得身份，并调用旧版 dashboard API。
- Schema v2 从 `metadata.name` 取得 UID；缺少 namespace 时默认使用 `default`，`spec` 必须是对象，并通过 Grafana dashboard resource API 写入。
- JSON 文件名去掉 `.json` 后必须与解析出的 UID 一致。本地目录只允许一层文件夹，其目录名会成为 Grafana folder UID。
- 对 schema v2，PIG 保留 `spec` 与 `grafana.app/message` 注解，根据本地文件夹写入 `grafana.app/folder`，并在 upsert 前主动去掉由服务端管理的 metadata/status。
- `dump` 只有在目标文件已经以 v2 形式存在时才保持 schema v2；此时会使用该文件的 namespace 拉取原生 v2 资源。全新导出目标默认写成传统 JSON；仅存在于本地的文件不会被 `dump` 删除。

其他 `dashboard.grafana.app/*` 版本暂不支持。需要导出时保留 v2 格式，请保留已有的本地 v2 文件。

## 设计记录

- [原生 sty boot](https://pig.pgsty.com/zh/design/native-sty-boot/)
- [原生 sty conf](https://pig.pgsty.com/zh/design/native-sty-conf/)
- [原生 Grafana 管理](https://pig.pgsty.com/zh/design/native-grafana-management/)
- [复用 Pigsty 已经拥有的 CMDB](https://pig.pgsty.com/zh/design/existing-cmdb-boundary/)
