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
  pig build repo                   # init build repo (=repo set, with remove+update)
  pig build tool  [mini|full|...]  # init build toolset
  pig build rust  [-y]             # install Rust toolchain
  pig build pgrx  [-v <ver>]       # install & init pgrx (0.16.1)
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

| 命令 | 描述 | 备注 |
|:---|:---|:---|
| `build spec` | 初始化构建规范仓库 | |
| `build repo` | 初始化所需仓库 | 需要 sudo 或 root 权限 |
| `build tool` | 初始化构建工具 | 需要 sudo 或 root 权限 |
| `build rust` | 安装 Rust 工具链 | 需要 sudo 或 root 权限 |
| `build pgrx` | 安装并初始化 pgrx | 需要 sudo 或 root 权限 |
| `build proxy` | 初始化构建代理 | |
| `build get` | 下载源代码 tarball | |
| `build dep` | 安装扩展构建依赖 | 需要 sudo 或 root 权限 |
| `build ext` | 构建扩展包 | 需要 sudo 或 root 权限 |
| `build pkg` | 完整构建流程：get、dep、ext | 需要 sudo 或 root 权限 |
{.full-width}


## 快速入门

设置构建环境并构建扩展的最快方式：

```bash
# 步骤 1：初始化构建规范
pig build spec

# 步骤 2：构建扩展（完整流程）
pig build pkg citus

# 构建的包将位于：
# - EL: ~/rpmbuild/RPMS/
# - Debian: ~/
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
~/ext/                           # 默认构建规范目录
├── Makefile                     # 主构建 makefile
├── <extension>/                 # 每个扩展的目录
│   ├── Makefile                # 扩展特定的 makefile
│   ├── <extension>.spec        # RPM 规范文件（EL）
│   └── debian/                 # Debian 打包文件
│       ├── control
│       ├── rules
│       └── ...
```

**构建输出位置：**
- **EL 系统**：`~/rpmbuild/RPMS/<arch>/`
- **Debian 系统**：`~/`（deb 文件）


## build spec

设置构建规范仓库和目录结构。

```bash
pig build spec                   # 在默认位置初始化 ~/ext
```

**功能：**
1. 克隆或更新扩展构建规范仓库
2. 创建必要的目录结构
3. 设置 makefile 和构建脚本
4. 准备平台特定的打包文件

**仓库位置：** 默认 `~/ext/`，包含 100+ 扩展的构建规范。


## build repo

初始化构建扩展所需的包仓库。

```bash
pig build repo                   # 等同于：pig repo set（内部执行 remove+update）
```

**功能：** 移除现有仓库、添加所有必需仓库（pgdg、pigsty、node）、更新包缓存。


## build tool

安装必要的开发工具和编译器。

```bash
pig build tool                   # 安装默认工具集
pig build tool mini              # 最小工具集
pig build tool full              # 完整工具集
pig build tool rust              # 添加 Rust 开发工具
```

**工具包：**

- **最小（`mini`）：** GCC/Clang 编译器、Make 和构建必需品、PostgreSQL 开发头文件、基本库
- **默认：** 所有最小工具、额外编译器（g++、clang++）、开发库、打包工具（rpmbuild、dpkg-dev）
- **完整（`full`）：** 所有默认工具、语言特定工具（Python、Perl、Ruby 开发）、高级调试工具、性能分析工具


## build rust

安装 Rust 编程语言工具链，基于 Rust 的扩展所需。

```bash
pig build rust                   # 带确认安装
pig build rust -y                # 强制重新安装 Rust 工具链
```

**安装内容：** Rust 编译器（rustc）、Cargo 包管理器、Rust 标准库、开发工具。


## build pgrx

安装并初始化 PGRX（Rust 的 PostgreSQL 扩展框架）。

```bash
pig build pgrx                   # 安装最新稳定版 (0.16.1)
pig build pgrx -v 0.15.0         # 安装特定版本
```

**前提条件：** 必须先安装 Rust 工具链、PostgreSQL 开发头文件。


## build proxy

为受限互联网访问的构建环境设置代理配置。

```bash
pig build proxy                  # 交互式设置
pig build proxy user@host:8080   # 直接配置
pig build proxy http://proxy.company.com:3128
```


## build get

下载扩展源代码 tarball。

```bash
pig build get citus              # 单个扩展
pig build get citus pgvector     # 多个扩展
pig build get all                # 所有可用扩展
pig build get std                # 标准扩展
```


## build dep

安装构建扩展所需的依赖。

```bash
pig build dep citus              # 单个扩展
pig build dep citus pgvector     # 多个扩展
pig build dep citus --pg 17,16   # 为特定 PG 版本
```


## build ext

编译扩展并创建安装包。

```bash
pig build ext citus              # 构建单个扩展
pig build ext citus pgvector     # 构建多个
pig build ext citus --pg 17      # 为特定 PG 版本
pig build ext citus -s           # 包含调试符号（仅 RPM）
```


## build pkg

执行完整的构建流程：下载、依赖和构建。

```bash
pig build pkg citus              # 构建单个扩展
pig build pkg citus pgvector     # 构建多个
pig build pkg citus --pg 17,16   # 为多个 PG 版本
pig build pkg citus -s           # 包含调试符号
```


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
sudo rpm -ivh ~/rpmbuild/RPMS/x86_64/pg_partman*.rpm  # EL
sudo dpkg -i ~/pg_partman*.deb                         # Debian
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
cargo install cargo-pgrx --force

# 重新初始化 PGRX
cargo pgrx init
```


## 扩展构建矩阵

### 常见构建的扩展

| 扩展 | 类型 | 构建时间 | 复杂度 | 特殊要求 |
|:---|:---:|:---|:---|:---|
| pg_repack | C | 快速 | 简单 | 无 |
| pg_partman | SQL/PLPGSQL | 快速 | 简单 | 无 |
| citus | C | 中等 | 中等 | 无 |
| timescaledb | C | 慢 | 复杂 | CMake |
| postgis | C | 非常慢 | 复杂 | GDAL、GEOS、Proj |
| pg_duckdb | C++ | 中等 | 中等 | C++17 编译器 |
| pgroonga | C | 中等 | 中等 | Groonga 库 |
| pgvector | C | 快速 | 简单 | 无 |
| plpython3 | C | 中等 | 中等 | Python 开发 |
| pgrx 扩展 | Rust | 慢 | 复杂 | Rust、PGRX |
{.full-width}
