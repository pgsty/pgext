---
title: "命令：repo"
description: 如何使用 pig repo 子命令管理软件仓库？
icon: fa-solid fa-terminal
weight: 610
---

`pig repo` 管理 Linux 软件仓库：添加或移除 APT/YUM/DNF 软件源，更新索引，并制作或解压离线包。`repo create` 还支持在 macOS 上通过 SOW 创建仓库。

| 命令            | 描述              | 备注                |
|:--------------|:----------------|:------------------|
| `repo list`   | 打印可用仓库与模块列表     |                   |
| `repo info`   | 获取仓库详细信息        |                   |
| `repo status` | 显示当前仓库状态        |                   |
| `repo add`    | 添加新仓库           | 需要 sudo 或 root 权限 |
| `repo set`    | 清空、覆盖并更新仓库      | 需要 sudo 或 root 权限 |
| `repo rm`     | 移除仓库            | 需要 sudo 或 root 权限 |
| `repo update` | 更新仓库缓存          | 需要 sudo 或 root 权限 |
| `repo create` | 创建本地 YUM/APT 仓库 | Linux：sudo/root；macOS：当前用户 |
| `repo cache`  | 从本地仓库创建离线包      | 需要 sudo 或 root 权限 |
| `repo boot`   | 从离线包引导仓库        | 需要 sudo 或 root 权限 |
| `repo reload` | 刷新仓库目录          |                   |
{.full-width}

## 快速入门

为已配置系统软件源的主机添加 PostgreSQL 仓库，并更新索引：

```bash
pig repo add pgsql -u
```

在中国网络环境中，加上 `-m` 选择中国区域镜像：

```bash
pig repo add pgsql -m -u
```

需要由 PIG 替换整套软件源时，使用 `set`。该命令先备份现有定义，再添加 `node`、`infra`、`pgsql` 并更新索引：

```bash
pig repo set
```

使用 `pig repo info pgsql` 查看仓库定义，使用 `pig repo status` 查看主机当前配置。

## 模块

在 pig 中，APT/YUM 仓库被组织为 **模块** —— 服务于特定目的的一组仓库。

|    模块     | 说明                      | 仓库列表                                              |
|:---------:|:------------------------|:--------------------------------------------------|
|   `all`   | 安装 PG 所需的全部核心模块         | `node` + `infra` + `pgsql`                        |
|  `pgsql`  | PGDG + Pigsty PG 扩展     | `pigsty-pgsql` + `pgdg`                           |
| `pigsty`  | Pigsty Infra + PGSQL 仓库 | pigsty-infra, pigsty-pgsql                        |
|  `pgdg`   | PGDG 官方仓库               | EL：pgdg-common、pgdg14-18；APT：pgdg                            |
|  `node`   | Linux 系统仓库              | base, updates, extras, epel, baseos, appstream... |
|  `infra`  | 基础设施组件仓库 | pigsty-infra |
| `nginx` | Nginx 仓库 | nginx |
| `docker`  | Docker 仓库               | docker-ce                                         |
|  `beta`   | PostgreSQL 19 Beta 版本   | pgdg19-beta, pgdg-beta                            |
|  `extra`  | PGDG Non-Free 与三方扩展     | pgdg-extras, timescaledb, citus                   |
| `groonga` | PGroonga 仓库             | groonga                                           |
| `percona` | Percona PG + PG_TDE     | percona                                           |
|  `llvm`   | LLVM 工具链仓库              | llvm                                              |
|  `kube`   | Kubernetes 仓库           | kubernetes                                        |
| `grafana` | Grafana 仓库              | grafana                                           |
| `haproxy` | HAProxy 仓库              | haproxyd, haproxyu                                |
|  `redis`  | Redis 仓库                | redis                                             |
|  `mongo`  | MongoDB 仓库              | mongo                                             |
|  `mysql`  | MySQL 仓库                | mysql                                             |
|  `click`  | ClickHouse 仓库           | clickhouse                                        |
| `gitlab`  | GitLab 仓库               | gitlab-ce, gitlab-ee                              |
{.full-width}

模块中的仓库会按当前操作系统与架构筛选。用 `pig repo list` 查看本机可用项。

已有可用的系统软件源时，添加 `pgsql` 即可获取 PGDG 与 Pigsty 扩展包。`infra` 提供基础设施工具和部分 PostgreSQL 分支内核包。`all` 是 `node`、`infra`、`pgsql` 的组合；不带参数的 `repo add` 和 `repo set` 默认选择 `all`。

`repo add`、`repo set`、`repo info` 支持逗号分隔的名称，并忽略名称两侧的空白。
空参数、纯空白或 `node,` 中的空项会使整次请求在操作仓库前失败；显式空参数不等于省略参数。

## 仓库定义

Pigsty 中可用仓库的完整定义位于 [`cli/repo/assets/repo.yml`](https://github.com/pgsty/pig/blob/main/cli/repo/assets/repo.yml)。

创建 `~/.pig/repo.yml` 可覆盖内置目录。文件格式见上面的源码链接。`repo reload` 会更新此文件，因此应在更新前保存自己的修改。

`--region` 选择 `baseurl` 中的区域 URL；不存在对应区域时使用 `default`。`-m` 选择 `china`，包含 Pigsty 中国镜像及目录中配置的其他镜像。

### 信任策略与配置归属

`repo add` 写入所选模块的配置文件，并保留其他文件。`repo set` 或 `repo add --remove` 会先备份现有定义再替换。若仓库配置由 Ansible 等工具管理，请在对应配置中维护这些定义，避免下一次运行覆盖手工修改。

### Pigsty 公钥 {#repo-key}

添加包含 `pigsty-infra` 或 `pigsty-pgsql` 的模块时，PIG 会复用已有公钥文件，或从程序内嵌资源安装公钥，权限为 `0644`：

| 平台 | 公钥路径 |
|:---|:---|
| EL | `/etc/pki/rpm-gpg/RPM-GPG-KEY-pigsty` |
| Debian / Ubuntu | `/etc/apt/keyrings/pigsty.asc` |

此步骤可以离线完成。只添加 `pgdg`、`node` 等其他仓库时，不会准备 Pigsty 公钥。

**内置仓库默认不强制校验软件包签名**：EL 使用 `gpgcheck=0`，Debian/Ubuntu 使用 `trusted=yes`。普通 `repo add/set` 保留定义中的校验设置；[`pig sty boot`](/zh/pig/cmd/sty/#sty-boot) 在公钥准备成功时启用 Pigsty 签名校验。需要强制校验时，在仓库元数据中配置可信密钥并启用校验。

默认公钥安装失败时，命令会告警并继续。所选 Pigsty 仓库使用 `trusted=yes`，或 `gpgcheck=0` 与 `repo_gpgcheck=0`，并保留默认密钥引用。告警也会写入结构化结果的 `data.warnings`。APT 此时可能沿用旧索引，请检查缓存更新输出。

显式指定密钥时，准备失败会中止仓库配置：

- Debian/Ubuntu 的 `signed-by` 使用已有普通密钥文件的绝对路径，多个路径用逗号分隔。
- EL 的 `gpgkey` 使用路径或不含变量的 URL，PIG 会调用 `rpm --import` 导入。
- 显式引用默认 Pigsty 路径时，PIG 可从内嵌资源安装该文件。

默认自动安装仅准备公钥文件和仓库引用。显式 EL `gpgkey` 才会额外导入 RPM 数据库。实现依据见[设计记录](https://pig.pgsty.com/zh/design/automatic-pigsty-repo-key/)。

EL 仓库保留 DNF 的模块过滤规则。只有声明 `module_hotfixes=1` 的定义（如 Pigsty 与 PGDG）会绕过模块流过滤；生成 EL7 配置时会移除该键。

## repo list

`pig repo list` 将列出当前系统可用的所有仓库模块。

```bash
pig repo list                # 列出当前系统可用仓库
pig repo list all            # 列出所有仓库（不过滤）
```

## repo info

显示特定仓库或模块的详细信息，包括 URL、元数据和区域镜像，以及 `.repo` / `.list` 仓库文件内容。

```bash
pig repo info pgdg               # 显示 pgdg 模块的信息
pig repo info pigsty pgdg        # 显示多个模块的信息
pig repo info all                # 显示所有模块的信息
```

## repo status

显示系统上的当前仓库配置。

```bash
pig repo status
```

## repo add

添加仓库配置文件到系统。需要 root/sudo 权限。

```bash
pig repo add pgdg                # 添加 PGDG 仓库
pig repo add pgdg pigsty         # 添加多个仓库
pig repo add all                 # 添加所有必要仓库 (pgdg + pigsty + node)
pig repo add pigsty -u           # 添加并更新缓存
pig repo add all -r              # 添加前移除现有仓库
pig repo add all -ru             # 移除、添加并更新（完全重置）
pig repo add pgdg --region=china # 使用中国镜像
pig repo add pgsql -m            # 明确选择中国区域镜像
```

**选项：**

- `-r|--remove`：添加新仓库前移除现有仓库
- `-u|--update`：添加仓库后运行包缓存更新
- `-m|--mirror`：明确选择内置的 `china` 仓库定义
- `--region <region>`：使用区域镜像仓库（`default` / `china` / `europe`）

|   平台   | 模块位置                                    |
|:------:|:----------------------------------------|
|   EL   | `/etc/yum.repos.d/<module>.repo`        |
| Debian | `/etc/apt/sources.list.d/<module>.list` |
{.full-width}

## repo set

先备份并移除现有仓库定义，再写入所选模块，最后刷新软件包元数据。
这等同于 `repo add --remove --update`。省略模块时，PIG 使用 `all`。

```bash
pig repo set                     # 替换为默认仓库
pig repo set pgdg pigsty         # 替换为特定仓库并更新
pig repo set all --region=china  # 使用中国镜像
pig repo set -m                  # 明确选择中国区域镜像
```

`repo set` 支持与 `repo add` 相同的 `--region` 与 `-m|--mirror` 区域选项。

## repo rm

真正不带参数时，将全部软件源定义移到既有备份目录。显式模块参数只删除对应模块文件；
`all` 只展开为 `node`、`pgsql`、`infra`，不代表所有文件。空字符串、纯空白、逗号空项与
不安全名称会在任何仓库或缓存操作前失败，混合列表也不例外。`-u` 在成功全量备份或模块删除后
更新缓存。原有备份与恢复契约保持不变。
目标模块文件已不存在时仍返回成功；摘要报告已确保不存在的模块数，不宣称每次请求都删除了
既有文件。仍支持仓库目录内的自定义模块名称。
全量备份没有找到仓库文件时，摘要明确报告没有文件需要备份。
结构化结果的 `requested_modules` 保留原始参数，包括 `all` 和逗号分隔的参数；
`removed_repos` 列出展开后的模块及各项执行结果。

```bash
pig repo rm                      # 移除所有仓库
pig repo rm pgdg                 # 移除特定仓库
pig repo rm pgdg pigsty -u       # 移除并更新缓存
```

|   平台   | 备份位置                              |
|:------:|:----------------------------------|
|   EL   | `/etc/yum.repos.d/backup/`        |
| Debian | `/etc/apt/sources.list.d/backup/` |
{.full-width}

## repo update

更新包管理器缓存以反映仓库更改。

```bash
pig repo update                  # 更新包缓存
```

|   平台   | 等效命令            |
|:------:|:----------------|
|   EL   | `dnf makecache`（EL7：`yum makecache`） |
| Debian | `apt-get update`    |
{.full-width}

## repo create

为离线安装创建本地包仓库。

```bash
pig repo create                  # Linux：/www/pigsty；macOS：当前目录
pig repo create /srv/repo        # 在自定义位置创建
```

当前实现会优先使用 `PATH` 中的 [`sow`](https://sow.pgsty.com/zh/docs/reference/cli/create/)。在 Linux 上选择 SOW 后，会对每个目标目录以 sudo 执行等价命令：

```bash
sow create --pigsty --timeout 10m -- /absolute/repository/path
```

在 macOS 上，先安装 SOW。执行时不使用 `sudo`，默认目标为当前目录。

在 Linux 上，未安装 SOW 时，PIG 在 EL 上使用 `createrepo_c`，在 Debian/Ubuntu 上使用 `dpkg-dev` 提供的 `dpkg-scanpackages`。
SOW 执行失败时，PIG 直接返回错误。`10m` 限制等待 SOW 目录锁的时间。

SOW 的 `--pigsty` 模式会：

1. 只扫描顶层普通 `.rpm` 与 `.deb` 文件，不递归，也不跟随符号链接。
2. 删除 `i386` 架构的 DEB 包，以及包名为 `patroni`、上游版本为 `3.0.4` 的 RPM/DEB 包。
3. 生成 RPM/DEB 元数据，最后写入 `repo_complete`；标记文件按包文件名排序，记录剩余顶层软件包的 SHA-256。

`annobin-docs` 等文档包会保留，除非命中其他清理规则；不会仅因文件名包含 `docs` 而删除软件包。

SOW 保留非软件包文件和目录；软件包无法解析或同一软件包坐标冲突时停止。旧后端在 `repo_complete` 中记录 MD5，清理规则也与 SOW 不同。PIG 检查该文件是否存在，不校验其中的哈希；分发离线包前可用该清单检查文件完整性。

## repo cache

创建仓库内容的压缩 tarball 用于离线分发。

```bash
pig repo cache                   # 默认：/www 到 /tmp/pkg.tgz
pig repo cache -d /srv           # 自定义源目录
pig repo cache pigsty docker     # 包含多个顶层仓库
```

**选项：**

- `-d, --dir`：源目录（默认：`/www/`）
- `-p, --path`：输出路径（默认：`/tmp/pkg.tgz`）

离线包必须包含 `pigsty`，且 `pigsty/repo_complete` 必须是普通文件。请求中不存在的仓库会被
跳过，并与实际打包的仓库分别报告。原生 boot 解包器只接受目录和普通文件，因此 cache
拒绝符号链接。cache 失败时不会覆盖已有的输出包。

## repo boot

从离线包解压并设置本地仓库。

```bash
pig repo boot                    # 默认：/tmp/pkg.tgz 到 /www
pig repo boot -p /mnt/pkg.tgz   # 自定义包路径
pig repo boot -d /srv           # 自定义目标目录
```

**选项：**

- `-p, --path`：包路径（默认：`/tmp/pkg.tgz`）
- `-d, --dir`：目标目录（默认：`/www/`）

命令先解压到临时目录，检查目标路径，再移动仓库。附加仓库先移动，`pigsty` 最后移动。

- 目标已有完整的 `pigsty/repo_complete` 时，直接复用该仓库。命令不会读取离线包，也不会导入包中的附加仓库；结构化结果显示 `reused: true`。
- 目标路径冲突时停止。若上一次运行中断，先检查并移走已移动的冲突目录，再重试。
- 离线包必须包含普通文件 `pigsty/repo_complete`。PIG 检查该文件是否存在，不校验其内容。

## repo reload

从 Pigsty 仓库站点下载最新仓库目录。

```bash
pig repo reload                  # 刷新仓库目录
```

更新后的文件会放置于 `~/.pig/repo.yml` 中。
