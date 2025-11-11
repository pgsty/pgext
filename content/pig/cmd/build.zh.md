---
title: "命令：build"
description: 如何使用 pig build 子命令设置构建基础设施？
icon: SquareTerminal
weight: 640
---

`pig build` 命令是一个强大的工具，简化了从源码构建 PostgreSQL 扩展的整个工作流程。它提供了完整的构建基础设施设置、依赖管理，以及标准和自定义 PostgreSQL 扩展在不同操作系统上的编译环境。

------

## 子命令

| 命令      | 描述                   |
|---------|----------------------|
| `spec`  | 初始化构建规范仓库            |
| `repo`  | 初始化所需仓库              |
| `tool`  | 初始化构建工具              |
| `rust`  | 安装 Rust 工具链          |
| `pgrx`  | 安装并初始化 pgrx（需要 Rust） |
| `proxy` | 初始化构建代理              |
| `get`   | 下载源代码 tarball        |
| `dep`   | 安装扩展构建依赖             |
| `ext`   | 构建扩展包                |
| `pkg`   | 完整构建流程：get、dep、ext   |

------

## 概览

```bash
构建 Postgres 扩展

Usage:
  pig build [command]

别名:
  build, b

示例:
pig build - 构建 Postgres 扩展

环境设置:
  pig build spec                   # 初始化构建规范和目录 (~ext)
  pig build repo                   # 初始化构建仓库 (=repo set -ru)
  pig build tool  [mini|full|...]  # 初始化构建工具集
  pig build rust  [-y]             # 安装 Rust 工具链
  pig build pgrx  [-v <ver>]       # 安装并初始化 pgrx (0.16.1)
  pig build proxy [id@host:port ]  # 初始化构建代理（可选）

包构建:
  pig build pkg   [ext|pkg...]     # 完整流程：get + dep + ext
  pig build get   [ext|pkg...]     # 下载扩展源代码 tarball
  pig build dep   [ext|pkg...]     # 安装扩展构建依赖
  pig build ext   [ext|pkg...]     # 构建扩展包

快速开始:
  pig build spec                   # 设置构建规范和目录
  pig build pkg citus              # 构建 citus 扩展

可用命令:
  dep         安装扩展构建依赖
  ext         构建扩展包
  get         下载源代码 tarball
  pgrx        安装并初始化 pgrx（需要 Rust）
  pkg         完整构建流程：get、dep、ext
  proxy       初始化构建代理
  repo        初始化所需仓库
  rust        安装 Rust 工具链
  spec        初始化构建规范仓库
  tool        初始化构建工具

标志:
  -h, --help   获取帮助信息

全局标志:
      --debug              启用调试模式
  -H, --home string        pigsty 主目录路径
  -i, --inventory string   配置清单路径
      --log-level string   日志级别: debug, info, warn, error, fatal, panic (默认 "info")
      --log-path string    日志文件路径，默认为终端输出
```

------

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

------

## 构建基础设施

### 构建规范

构建系统使用规范文件来定义每个扩展应该如何构建。这些规范包括：
- 源代码位置和版本
- 构建依赖
- 编译标志
- PostgreSQL 版本兼容性
- 平台特定的构建指令

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

构建输出位置：
- **EL 系统**：`~/rpmbuild/RPMS/<arch>/`
- **Debian 系统**：`~/`（deb 文件）

------

## 命令参考

### `build spec`

设置构建规范仓库和目录结构。

```bash
pig build spec                   # 在默认位置初始化 ~/ext
```

**功能：**
1. 克隆或更新扩展构建规范仓库
2. 创建必要的目录结构
3. 设置 makefile 和构建脚本
4. 准备平台特定的打包文件

**仓库位置：**
- 默认：`~/ext/`
- 包含 100+ 扩展的构建规范

**示例：**
```bash
# 初始化规范
pig build spec

# 检查可用扩展
ls ~/ext/

# 查看特定扩展规范
cat ~/ext/citus/Makefile
```

------

### `build repo`

初始化构建扩展所需的包仓库。

```bash
pig build repo                   # 等同于：pig repo set -ru
```

**功能：**
- 移除现有仓库
- 添加所有必需仓库（pgdg、pigsty、node）
- 更新包缓存

这确保您可以访问构建所需的所有必要包和依赖。

------

### `build tool`

安装必要的开发工具和编译器。

```bash
pig build tool                   # 安装默认工具集
pig build tool mini              # 最小工具集
pig build tool full              # 完整工具集
pig build tool rust              # 添加 Rust 开发工具
```

**工具包：**

**最小（`mini`）：**
- GCC/Clang 编译器
- Make 和构建必需品
- PostgreSQL 开发头文件
- 基本库

**默认：**
- 所有最小工具
- 额外编译器（g++、clang++）
- 开发库
- 打包工具（rpmbuild、dpkg-dev）

**完整（`full`）：**
- 所有默认工具
- 语言特定工具（Python、Perl、Ruby 开发）
- 高级调试工具
- 性能分析工具

**常见安装的包：**
```bash
# EL 系统
gcc gcc-c++ make cmake
postgresql17-devel
rpm-build rpmdevtools
git wget curl

# Debian 系统
build-essential
postgresql-server-dev-17
dpkg-dev debhelper
git wget curl
```

------

### `build rust`

安装 Rust 编程语言工具链，基于 Rust 的扩展所需。

```bash
pig build rust                   # 带确认安装
pig build rust -y                # 自动确认安装
```

**安装内容：**
- Rust 编译器（rustc）
- Cargo 包管理器
- Rust 标准库
- 开发工具

**安装方法：**
- 使用 rustup 安装器
- 安装到用户主目录
- 自动配置 PATH

**示例：**
```bash
# 安装 Rust
pig build rust -y

# 验证安装
rustc --version
cargo --version
```

------

### `build pgrx`

安装并初始化 PGRX（Rust 的 PostgreSQL 扩展框架）。

```bash
pig build pgrx                   # 安装最新稳定版 (0.16.1)
pig build pgrx -v 0.15.0         # 安装特定版本
```

**前提条件：**
- 必须先安装 Rust 工具链
- PostgreSQL 开发头文件

**功能：**
1. 安装 cargo-pgrx 工具
2. 为已安装的 PostgreSQL 版本初始化 PGRX
3. 设置开发环境

**示例：**
```bash
# 首先安装 Rust
pig build rust -y

# 安装 PGRX
pig build pgrx

# 为特定 PG 版本初始化
cargo pgrx init --pg17=/usr/pgsql-17/bin/pg_config
```

------

### `build proxy`

为受限互联网访问的构建环境设置代理配置。

```bash
pig build proxy                  # 交互式设置
pig build proxy user@host:8080   # 直接配置
pig build proxy http://proxy.company.com:3128
```

**配置代理用于：**
- 系统环境（HTTP_PROXY、HTTPS_PROXY）
- 包管理器（yum/dnf、apt）
- 构建工具（cargo、npm、pip）
- Git 配置

**示例：**
```bash
# 设置企业代理
pig build proxy proxy.corp.com:8080

# 带认证
pig build proxy user:pass@proxy.corp.com:8080
```

------

### `build get`

下载扩展源代码 tarball。

```bash
pig build get citus              # 单个扩展
pig build get citus pgvector     # 多个扩展
pig build get all                # 所有可用扩展
pig build get std                # 标准扩展
```

**源位置：**
- 下载到：`~/ext/<extension>/`
- 源类型：
  - GitHub 发布 tarball
  - 官方项目下载
  - 源仓库

**示例：**
```bash
# 下载特定版本
pig build get citus

# 检查下载的源码
ls ~/ext/citus/*.tar.gz
```

------

### `build dep`

安装构建扩展所需的依赖。

```bash
pig build dep citus              # 单个扩展
pig build dep citus pgvector     # 多个扩展
pig build dep citus --pg 17,16   # 为特定 PG 版本
```

**选项：**
- `--pg`：逗号分隔的 PostgreSQL 版本（例如 '17,16'）

**依赖类型：**
- 构建工具（编译器、make）
- 开发库
- PostgreSQL 头文件
- 扩展特定要求

**示例：**
```bash
# 为 Citus 安装依赖
pig build dep citus

# 为多个 PostgreSQL 版本
pig build dep citus --pg 17,16,15
```

------

### `build ext`

编译扩展并创建安装包。

```bash
pig build ext citus              # 构建单个扩展
pig build ext citus pgvector     # 构建多个
pig build ext citus --pg 17      # 为特定 PG 版本
pig build ext citus -s           # 包含调试符号（仅 RPM）
```

**选项：**
- `--pg`：目标 PostgreSQL 版本
- `-s, --symbol`：包含调试符号

**构建过程：**
1. 解压源代码
2. 配置构建环境
3. 编译扩展
4. 创建平台包（RPM/DEB）
5. 签名包（如果配置）

**输出位置：**
- EL：`~/rpmbuild/RPMS/<arch>/*.rpm`
- Debian：`~/*.deb`

**示例：**
```bash
# 为 PostgreSQL 17 构建 Citus
pig build ext citus --pg 17

# 检查构建的包
ls ~/rpmbuild/RPMS/x86_64/citus*.rpm  # EL
ls ~/citus*.deb                        # Debian
```

------

### `build pkg`

执行完整的构建流程：下载、依赖和构建。

```bash
pig build pkg citus              # 构建单个扩展
pig build pkg citus pgvector     # 构建多个
pig build pkg citus --pg 17,16   # 为多个 PG 版本
pig build pkg citus -s           # 包含调试符号
```

**等同于运行：**
```bash
pig build get <extension>
pig build dep <extension>
pig build ext <extension>
```

**选项：**
- `--pg`：目标 PostgreSQL 版本
- `-s, --symbol`：包含调试符号

**示例：**
```bash
# 完整构建流程
pig build pkg citus --pg 17

# 构建多个扩展
pig build pkg citus pg_partman timescaledb
```

------

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
pig build rust -y
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

### 工作流 4：自定义扩展构建

```bash
# 1. 创建自定义规范
mkdir ~/ext/myextension
cd ~/ext/myextension

# 2. 创建 Makefile
cat > Makefile << 'EOF'
EXTENSION = myextension
VERSION = 1.0.0
REPO_URL = https://github.com/myorg/myextension
include ../Makefile
EOF

# 3. 构建
pig build pkg myextension
```

### 工作流 5：调试构建问题

```bash
# 带调试输出构建
pig build pkg citus --debug

# 带调试符号构建
pig build pkg citus -s

# 分步调试
pig build get citus              # 检查源码下载
pig build dep citus              # 验证依赖
pig build ext citus --debug     # 调试编译
```

### 工作流 6：为分发构建

```bash
# 1. 构建多个扩展
extensions="citus pg_partman timescaledb pgvector"
for ext in $extensions; do
    pig build pkg $ext --pg 16,17
done

# 2. 收集包
mkdir ~/packages
cp ~/rpmbuild/RPMS/x86_64/*.rpm ~/packages/  # EL
cp ~/*.deb ~/packages/                        # Debian

# 3. 创建仓库
pig repo create ~/packages

# 4. 分发
tar czf extensions.tar.gz ~/packages
```

------

## 扩展构建规范

### 标准扩展

具有官方构建支持的扩展：
- **分布式**：citus、pg_partman
- **时间序列**：timescaledb、pg_cron
- **分析**：pg_duckdb、parquet_fdw
- **搜索**：pgroonga、pg_bigm
- **语言**：plpython3、plperl、pltcl

### 构建要求

不同的扩展有不同的要求：

**简单 C 扩展：**
- PostgreSQL 开发头文件
- C 编译器（gcc/clang）
- Make

**复杂扩展：**
- 额外的库（例如 PostGIS 的 GDAL）
- 语言运行时（Python、Perl 等）
- 特殊编译器（Rust、Go）

**Rust 扩展（通过 PGRX）：**
- Rust 工具链
- PGRX 框架
- Cargo 构建系统

### 平台差异

**EL 系统（RHEL/Rocky/Alma）：**
- 使用 RPM 规范文件
- 使用 rpmbuild 构建
- 输出到 ~/rpmbuild/RPMS/

**Debian/Ubuntu 系统：**
- 使用 debian/ 目录结构
- 使用 dpkg-buildpackage 构建
- 输出到 ~/

------

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
sudo pig ext install pg17-devel

# 或指定 pg_config 路径
export PG_CONFIG=/usr/pgsql-17/bin/pg_config
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

### 源下载失败

```bash
# 检查网络/代理
pig build proxy

# 手动下载
cd ~/ext/<extension>
wget <source_url>

# 继续构建
pig build dep <extension>
pig build ext <extension>
```

### 构建失败

```bash
# 检查构建日志
tail -f ~/rpmbuild/BUILD/*/config.log  # EL

# 常见修复：
# 1. 更新构建工具
pig build tool full

# 2. 清除构建缓存
rm -rf ~/rpmbuild/BUILD/*

# 3. 尝试不同的 PG 版本
pig build pkg <extension> --pg 16
```

------

## 最佳实践

1. **一次设置，多次构建**：初始化构建环境一次，然后构建多个扩展

2. **使用 pkg 命令**：对于标准构建，使用 `pig build pkg` 完成整个流程

3. **检查兼容性**：构建前验证扩展支持您的 PostgreSQL 版本

4. **保持规范更新**：定期使用 `pig build spec` 更新构建规范

5. **测试构建**：在生产部署前在开发环境测试构建的包

6. **记录自定义构建**：保留自定义构建标志或修改的记录

7. **使用版本控制**：在版本控制中跟踪自定义构建规范

8. **并行构建**：并行构建独立扩展以提高效率

9. **缓存依赖**：保留常用依赖的安装以加快构建速度

10. **签名包**：为分发签名包，使用 GPG 确保安全

------

## 高级主题

### 自定义构建标志

```bash
# 设置自定义 CFLAGS
export CFLAGS="-O3 -march=native"
pig build ext <extension>

# PostgreSQL 特定标志
export PG_CPPFLAGS="-DUSE_SPECIAL_FEATURE"
pig build ext <extension>
```

### 交叉编译

```bash
# 为不同架构
export CC=aarch64-linux-gnu-gcc
export ARCH=arm64
pig build ext <extension>
```

### 调试符号

```bash
# 带调试符号构建
pig build ext <extension> -s

# 用于调试
export CFLAGS="-g -O0"
pig build ext <extension>
```

### 自定义 PostgreSQL 构建

```bash
# 对于自定义 PostgreSQL 安装
export PG_CONFIG=/opt/postgresql/bin/pg_config
export PKG_CONFIG_PATH=/opt/postgresql/lib/pkgconfig
pig build ext <extension>
```

### 持续集成

```bash
#!/bin/bash
# CI 构建脚本
set -e

# 设置
pig build spec
pig build tool

# 构建矩阵
for pg_version in 15 16 17; do
    for extension in citus pgvector timescaledb; do
        pig build pkg $extension --pg $pg_version
    done
done

# 测试安装
for rpm in ~/rpmbuild/RPMS/x86_64/*.rpm; do
    sudo rpm -Uvh --test $rpm
done
```

------

## 扩展构建矩阵

### 常见构建的扩展

| 扩展 | 类型 | 构建时间 | 复杂度 | 特殊要求 |
|-----------|------|------------|------------|---------------------|
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