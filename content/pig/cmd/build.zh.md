---
title: "命令：build"
description: 如何使用 pig build 子命令设置构建基础设施？
icon: SquareTerminal
weight: 640
---

`pig build` 命令是一个强大的工具，简化了从源码构建 PostgreSQL 扩展的整个工作流程。它提供了完整的构建基础设施设置、依赖管理，以及标准和自定义 PostgreSQL 扩展在不同操作系统上的编译环境。

```bash
pig build - Build Postgres Extension

Environment Setup:
  pig build spec                   # init build spec and directory (~ext)
  pig build repo                   # init build repo (=repo set -ru)
  pig build repo --beta            # init build repo with PostgreSQL beta repo
  pig build tool  [mini|full|...]  # init build toolset
  pig build rust  [-y] [-m]        # install Rust toolchain
  pig build pgrx  [-v <ver>] [-b]  # install & init pgrx (0.19.1)
  pig build proxy [id@host:port ]  # init build proxy (optional)

Package Building:
  pig build pkg   [ext|pkg...]     # complete pipeline: get + dep + ext
  pig build get   [ext|pkg...]     # download extension source tarball
  pig build dep   [ext|pkg...]     # install extension build dependencies
  pig build ext   [ext|pkg...]     # build extension package

Quick Start:
  pig build spec                   # setup build spec and directory
  pig build pkg citus              # build citus extension
```

| 命令            | 描述                 | 备注                |
|:--------------|:-------------------|:------------------|
| `build spec`  | 初始化构建规范目录          |                   |
| `build repo`  | 初始化所需仓库            | 需要 sudo 或 root 权限 |
| `build tool`  | 初始化构建工具            | 需要 sudo 或 root 权限 |
| `build rust`  | 安装 Rust 工具链        | 需要 sudo 或 root 权限 |
| `build pgrx`  | 安装并初始化 pgrx        | 需要 sudo 或 root 权限 |
| `build proxy` | 初始化构建代理            |                   |
| `build get`   | 下载源代码 tarball      |                   |
| `build dep`   | 安装扩展构建依赖           | 需要 sudo 或 root 权限 |
| `build ext`   | 构建扩展包              | 需要 sudo 或 root 权限 |
| `build pkg`   | 完整构建流程：get、dep、ext | 需要 sudo 或 root 权限 |
{.full-width}


## 快速入门

设置构建环境并构建扩展的最快方式：

```bash
# 步骤 1：初始化构建规范
pig build spec

# 步骤 2：构建扩展（完整流程）
pig build pkg citus

# 构建的包将位于：
# - EL: ~/ext/pkg/（同时可从 ~/rpmbuild/RPMS/ 访问）
# - Debian: ~/ext/pkg/（同时可从 ~/debbuild/DEBS/ 访问）
```

更精细的控制方式：

```bash
# 设置环境
pig build spec                   # 初始化构建规范
pig build repo                   # 设置仓库
pig build tool                   # 安装构建工具

# 构建过程
pig build get citus              # 下载源码
pig build dep citus              # 安装依赖
pig build ext citus              # 构建包

# 或一次完成所有三个步骤
pig build pkg citus              # get + dep + ext
```


## 构建基础设施

### 目录结构

```
~/ext/                           # 真实工作目录
├── pkg/                         # 构建产物输出目录
├── src/                         # 源码 tarball 下载目录
├── log/                         # 构建日志目录
└── tmp/                         # 临时目录

~/rpmbuild/                      # EL 构建目录
├── RPMS -> ~/ext/pkg            # RPM 产物软链接
├── SOURCES -> ~/ext/src         # 源码软链接
├── SPECS/
├── BUILD/
├── BUILDROOT/
└── SRPMS/

~/debbuild/                      # Debian / Ubuntu 构建目录
├── DEBS -> ~/ext/pkg            # DEB 产物软链接
├── SOURCES -> ~/ext/src         # 源码软链接
├── SPECS/
└── BUILD/
```

**构建输出位置：**
- **EL 系统**：`~/ext/pkg/`，并通过 `~/rpmbuild/RPMS/` 软链接访问
- **Debian 系统**：`~/ext/pkg/`，并通过 `~/debbuild/DEBS/` 软链接访问


## build spec

设置构建规范与目录结构。

```bash
pig build spec                   # 在默认位置初始化 ~/ext
pig build spec -f                # 强制重新下载构建规范 tarball
pig build spec -m                # 优先使用 pigsty.cc 中国镜像
```

**功能：**
1. 下载 RPM 或 DEB 构建规范 tarball
2. 创建 `~/ext/{pkg,src,log,tmp}` 与平台构建目录
3. 将 `RPMS/DEBS` 与 `SOURCES` 软链接到 `~/ext/pkg` 和 `~/ext/src`
4. 通过增量 `rsync` 同步 makefile、spec 与 Debian 打包文件

**工作目录：** 默认使用 `~/ext/` 保存源码、产物、日志与临时文件；平台打包目录为 `~/rpmbuild/` 或 `~/debbuild/`。


## build repo

初始化构建扩展所需的包仓库。

```bash
pig build repo                   # 等同于：pig repo set -ru
pig build repo -m                # 使用镜像/代理路径添加 PostgreSQL 仓库
pig build repo -b                # 包含 PostgreSQL beta 仓库
```

**功能：** 以 `pig repo set -ru` 初始化构建所需仓库：移除旧仓库、添加所需仓库并更新包缓存。

**选项：**

- `-m|--mirror`：对 PostgreSQL 仓库使用 Pigsty 镜像/代理路径
- `-b|--beta`：包含 PostgreSQL beta 仓库模块，当前用于 PG19 beta 构建


## build tool

安装必要的开发工具和编译器。

```bash
pig build tool                   # 安装默认工具集
pig build tool mini              # 最小工具集
pig build tool full              # 完整工具集
pig build tool rust              # 添加 Rust 开发工具
pig build tool -b                # 包含 PostgreSQL beta 构建包
```

**工具包：**

- **最小（`mini`）：** GCC/Clang 编译器、Make 和构建必需品、PostgreSQL 开发头文件、基本库
- **默认：** 所有最小工具、额外编译器（g++、clang++）、开发库、打包工具（rpmbuild、dpkg-dev）
- **完整（`full`）：** 所有默认工具、语言特定工具（Python、Perl、Ruby 开发）、高级调试工具、性能分析工具

在 EL 系统上，默认工具集会安装稳定活跃 PostgreSQL 主版本对应的显式 devel/server 包；需要当前 beta 主版本时使用 `--beta`。


## build rust

安装 Rust 编程语言工具链，基于 Rust 的扩展所需。

```bash
pig build rust                   # 带确认安装
pig build rust -y                # 强制重新安装 Rust 工具链
pig build rust -m                # 使用 rsproxy.cn 与 Cargo 镜像配置
```

**安装内容：** Rust 编译器（rustc）、Cargo 包管理器、Rust 标准库、开发工具。

**选项：**

- `-y|--yes`：即使 Cargo 已存在也强制重新安装 Rust
- `-m|--mirror`：从中国镜像下载 `rustup`，并写入面向 `rsproxy.cn` 的 `~/.cargo/config.toml`


## build pgrx

安装并初始化 PGRX（Rust 的 PostgreSQL 扩展框架）。

```bash
pig build pgrx                   # 安装最新稳定版 (0.19.1)
pig build pgrx -v 0.19.1         # 安装特定版本
pig build pgrx --pg 18,17,16     # 为指定 PG 版本初始化 pgrx
pig build pgrx --pg init         # 只执行 cargo pgrx init，不传 PG 参数
pig build pgrx -b                # 自动探测时包含 beta PostgreSQL pg_config
```

**前提条件：** 必须先安装 Rust 工具链、PostgreSQL 开发头文件。

**选项：**

- `-v|--pgrx`：要安装的 cargo-pgrx 版本
- `--pg`：逗号分隔的 PostgreSQL 主版本，`init` 表示不传 PostgreSQL 参数，留空则自动探测
- `-b|--beta`：自动探测时包含当前 beta 主版本


## build proxy

为受限互联网访问的构建环境设置代理配置。

```bash
pig build proxy                  # 交互式设置
pig build proxy user@host:8080   # 使用默认本地端口 127.0.0.1:12345
pig build proxy user@host:8080 127.0.0.1:1080
```


## build get

下载扩展源代码 tarball。

```bash
pig build get citus              # 单个扩展
pig build get citus pgvector     # 多个扩展
pig build get pdu pgdog          # 使用内置源码 alias
pig build get citus -f           # 已存在时仍重新下载
pig build get citus -m           # 优先使用 pigsty.cc 中国镜像
```

`pig build get` 的参数是扩展名、包名或源码文件名；未知名称会按源码文件名处理。它不会把 `all` 或 `std` 展开为内置集合。需要批量下载时，请显式列出目标包名。

一些源码包并不直接对应扩展名，`pig build get` 内置了特殊 alias 以便直接下载源码。

```bash
pig build get pdu                # 下载 pdu-3.0.25.12.tar.gz
pig build get pgdog              # 下载 pgdog-0.1.32.tar.gz
pig build get pgedge-17          # 下载 PostgreSQL 17 与 pgEdge 套件源码
pig build get babelfish-18       # 下载 Babelfish PG18 源码包
pig build get orioledb-16        # 下载 OrioleDB 与 PG16 patch bundle
pig build get onesparse          # 同时下载 onesparse、graphblas、lagraph
```

当前常见的特殊源码 alias 包括：`babelfish` / `babelfish-17` / `babelfish-18`、`ivorysql`、`orioledb` / `orioledb-16` / `orioledb-17` / `orioledb-18`、`cloudberry` / `cloudberry-backup` / `cloudberry-pxf`、`pgedge` / `pgedge-15` / `pgedge-16` / `pgedge-17` / `pgedge-18`、`polardb`、`polarstore`、`zlog`、`pdu`、`pgdog`、`rdkit`、`onesparse`、`libpgfeutils`、`libfq`。


## build dep

安装构建扩展所需的依赖。

```bash
pig build dep citus              # 单个扩展
pig build dep citus pgvector     # 多个扩展
pig build dep citus --pg 17,16   # 为特定 PG 版本
```

**选项：**

- `--pg`：指定一个或多个 PostgreSQL 大版本；未指定时按扩展元数据或本机安装自动推断


## build ext

编译扩展并创建安装包。

```bash
pig build ext citus              # 构建单个扩展
pig build ext citus pgvector     # 构建多个
pig build ext citus --pg 17      # 为特定 PG 版本
pig build ext citus -s           # 包含调试符号（仅 RPM）
```

**选项：**

- `--pg`：指定一个或多个 PostgreSQL 大版本
- `-s|--symbol`：构建调试符号包（仅 RPM）


## build pkg

执行完整的构建流程：下载、依赖和构建。

```bash
pig build pkg citus              # 构建单个扩展
pig build pkg citus pgvector     # 构建多个
pig build pkg citus --pg 17,16   # 为多个 PG 版本
pig build pkg citus -s           # 包含调试符号
pig build pkg citus -m           # 优先使用 pigsty.cc 中国镜像下载源码
pig build pkg cloudberry         # 按顺序构建 cloudberry、backup、pxf
```

**选项：**

- `--pg`：指定一个或多个 PostgreSQL 大版本
- `-s|--symbol`：构建调试符号包（仅 RPM）
- `-m|--mirror`：下载源码时优先使用 `pigsty.cc` 镜像

`pig build pkg cloudberry` 是套件构建流程：先下载完整 Cloudberry 源码包，构建并安装本地核心产物，然后继续构建 `cloudberry-backup` 与 `cloudberry-pxf`。


## 常见工作流

### 工作流 1：构建标准扩展

```bash
# 1. 设置构建环境（一次性）
pig build spec
pig build repo
pig build tool

# 2. 构建扩展
pig build pkg pg_partman

# 3. 安装构建的包
sudo rpm -ivh ~/ext/pkg/pg_partman*.rpm  # EL
sudo dpkg -i ~/ext/pkg/*partman*.deb     # Debian
```

### 工作流 2：构建 Rust 扩展

```bash
# 1. 设置 Rust 环境
pig build spec
pig build tool
pig build rust                   # 如需强制重装可追加 -y
pig build pgrx

# 2. 构建 Rust 扩展
pig build pkg pgmq

# 3. 安装
sudo pig ext add pgmq
```

### 工作流 3：构建多个版本

```bash
# 为多个 PostgreSQL 版本构建扩展
pig build pkg citus --pg 15,16,17

# 结果为每个版本生成包：
# citus_15-*.rpm
# citus_16-*.rpm
# citus_17-*.rpm
```


## 故障排除

### 找不到构建工具

```bash
# 安装构建工具
pig build tool

# 对于特定编译器
sudo dnf groupinstall "Development Tools"  # EL
sudo apt install build-essential          # Debian
```

### 缺少依赖

```bash
# 安装扩展依赖
pig build dep <extension>

# 检查错误消息以了解特定包
# 如需要，手动安装
sudo dnf install <package>  # EL
sudo apt install <package>  # Debian
```

### 找不到 PostgreSQL 头文件

```bash
# 安装 PostgreSQL 开发包
sudo pig ext install pg18-devel

# 或指定 pg_config 路径
export PG_CONFIG=/usr/pgsql-18/bin/pg_config
```

### Rust/PGRX 问题

```bash
# 重新安装 Rust
curl --proto '=https' --tlsv1.2 -sSf https://sh.rustup.rs | sh

# 更新 PGRX
cargo install --locked cargo-pgrx@0.19.1

# 重新初始化 PGRX
cargo pgrx init
```


## 扩展构建矩阵

### 常见构建的扩展

| 扩展          |     类型      | 构建时间 | 复杂度 | 特殊要求           |
|:------------|:-----------:|:-----|:----|:---------------|
| pg_repack   |      C      | 快速   | 简单  | 无              |
| pg_partman  | SQL/PLPGSQL | 快速   | 简单  | 无              |
| citus       |      C      | 中等   | 中等  | 无              |
| timescaledb |      C      | 慢    | 复杂  | CMake          |
| postgis     |      C      | 非常慢  | 复杂  | GDAL、GEOS、Proj |
| pg_duckdb   |     C++     | 中等   | 中等  | C++17 编译器      |
| pgroonga    |      C      | 中等   | 中等  | Groonga 库      |
| pgvector    |      C      | 快速   | 简单  | 无              |
| plpython3   |      C      | 中等   | 中等  | Python 开发      |
| pgrx 扩展     |    Rust     | 慢    | 复杂  | Rust、PGRX      |
{.full-width}
