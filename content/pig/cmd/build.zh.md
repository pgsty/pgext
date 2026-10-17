---
title: "命令：build"
description: 如何使用 pig build 子命令设置构建基础设施？
icon: fa-solid fa-terminal
weight: 640
---

`pig build` 使用 Pigsty 的构建配方下载源码、安装依赖并生成 RPM/DEB 软件包。它适用于 EL、Debian 和 Ubuntu 构建主机；Rust、pgrx 和 Xray 客户端也支持各自的 macOS 使用场景。

| 命令            | 描述                 | 备注                |
|:--------------|:-------------------|:------------------|
| `build spec`  | 初始化构建规范目录          |                   |
| `build repo`  | 初始化所需仓库            | 需要 sudo 或 root 权限 |
| `build tool`  | 初始化构建工具            | 需要 sudo 或 root 权限 |
| `build rust` | 安装 Rust 工具链 | 当前用户 |
| `build pgrx` | 安装并初始化 pgrx | 当前用户 |
| `build proxy` | 设置 Xray 客户端与服务端 | Linux 使用 root；macOS 客户端使用普通用户 |
| `build get`   | 下载源代码 tarball      |                   |
| `build dep`   | 安装扩展构建依赖           | 需要 sudo 或 root 权限 |
| `build ext` | 构建扩展包 | 当前用户 |
| `build pkg`   | 完整构建流程：get、dep、ext | 需要 sudo 或 root 权限 |
{.full-width}

## 快速入门

构建主机需要配置软件源时，运行 `pig build repo`。它会备份并替换现有仓库定义。
已有合适的软件源时，跳过此步骤。

然后安装构建工具、下载配方并构建扩展：

```bash
pig build tool                  # 安装构建工具与 PostgreSQL 开发包
pig build spec                  # 下载构建配方并建立目录
pig build pkg citus --pg 18     # 下载源码、安装依赖、构建 PG 18 软件包
```

软件包写入 `~/ext/pkg/`，RPM 产物可能位于其架构子目录中。

需要分别运行各阶段时，使用：

```bash
pig build get citus
pig build dep citus --pg 18
pig build ext citus --pg 18
```

`build pkg` 组合的是这三个阶段。它不会替你安装通用构建工具或下载构建配方。

## 构建基础设施

### 目录结构

```filetree {title="PIG 构建目录"}
- ~/ext/
  - pkg/   {open=false}
  - src/   {open=false}
  - log/   {open=false}
  - tmp/   {open=false}
- ~/rpmbuild/
  - RPMS -> ~/ext/pkg
  - SOURCES -> ~/ext/src
  - SPECS/   {open=false}
  - BUILD/   {open=false}
  - BUILDROOT/   {open=false}
  - SRPMS/   {open=false}
- ~/debbuild/
  - DEBS -> ~/ext/pkg
  - SOURCES -> ~/ext/src
  - SPECS/   {open=false}
  - BUILD/   {open=false}
```

`~/ext/` 是真实工作目录：`pkg` 保存构建产物，`src` 保存源码压缩包，`log` / `tmp` 分别保存构建日志和临时文件；平台专用目录则把输出与源码位置链接回这里。

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
pig build repo                   # 等同于：pig repo set
pig build repo -m                # 选择内置的中国区域镜像
pig build repo --beta            # 同时启用 PostgreSQL 19 beta 仓库模块
```

**功能：** 以 `pig repo set` 初始化构建所需仓库：移除旧仓库、添加所需仓库并更新包缓存。`--beta/-b` 会把 `beta` 模块追加到仓库模块列表，用于显式构建 PostgreSQL 19 beta 相关包；稳定默认路径仍只使用 PG14-18。

**选项：**

- `-b|--beta`：额外启用 PostgreSQL beta 仓库模块
- `-m|--mirror`：选择内置的 `china` 仓库定义

## build tool

安装必要的开发工具和编译器。

```bash
pig build tool                   # 安装默认工具集
pig build tool mini              # 最小工具集
pig build tool full              # 完整工具集
pig build rust                   # 单独安装 Rust 工具链
pig build tool --beta            # 额外安装 PG19 beta 构建依赖
```

**工具包：**

- **最小（`mini`）：** 构建工具和开发库，从安装列表中排除 PostgreSQL 与 PGDG 包
- **默认 / `full`：** Make、CMake、开发库、打包工具和 PostgreSQL 构建依赖。EL 安装 PG14-18 server/devel 包；Debian/Ubuntu 使用 `postgresql-all` 与 `postgresql-server-dev-all`
- **`--beta`：** 额外安装 PG19 beta server/devel 包；`mini` 仍排除 PostgreSQL 包

## build rust

安装 Rust 编程语言工具链，基于 Rust 的扩展所需。

```bash
pig build rust                   # Cargo 不存在时安装
pig build rust -y                # 强制重新安装 Rust 工具链
pig build rust -m                # 使用中国镜像安装 Rust，并写入 Cargo 镜像配置
```

PIG 通过 rustup 为当前用户安装 Rust，不询问交互确认。如果 `~/.cargo/bin/cargo` 已存在，默认跳过安装；`-y` 强制重新安装。`-m` 为 rustup 选择 `rsproxy.cn`，并写入 Cargo 镜像配置，即使 Rust 已安装也会写入该配置。

## build pgrx

安装并初始化 PGRX（Rust 的 PostgreSQL 扩展框架）。

```bash
pig build pgrx                   # 安装默认版本 (0.19.3)
pig build pgrx -v 0.19.3         # 安装特定版本
pig build pgrx --pg 18,17,16     # 为指定 PG 版本初始化 pgrx
pig build pgrx --pg init         # 只执行 cargo pgrx init，不传 PG 参数
pig build pgrx -b                # 自动探测时包含 PostgreSQL 19 beta pg_config
```

**前提条件：** 必须先安装 Rust 工具链、PostgreSQL 开发头文件。默认自动探测只覆盖稳定 PG14-18；需要 PG19 beta 时使用 `-b|--beta`，或通过 `--pg 19` 显式指定。

## build proxy

旧源码生成的无管理注释 macOS LaunchAgent，在完整内容匹配时可直接升级，无需 `--replace`。
识别依据是完整生成内容，不能只凭管理注释接管文件。外来或修改过的服务文件仍需
`--replace --yes`。一次性迁移注释后，相同配置重复执行不会重启服务。
此兼容结论针对已验证的旧源码格式，不代表覆盖所有已发布二进制。

为互联网访问受限的构建环境设置 Xray 客户端与服务端，保留 `x` 别名。
不带参数的 `pig build proxy` 只安装或确认 Xray。
设计依据见 [Xray 设计记录](https://pig.pgsty.com/zh/design/xray-client-server/)。

### 客户端

一行命令按需安装 Xray、写入客户端配置、启动服务，并分别检查 HTTP 与 SOCKS 的 HTTPS 请求。
首次监听默认使用 **`127.0.0.1:12345`**，同一端口提供 HTTP 和 SOCKS。
已有受支持客户端省略 `--listen` 时保留原监听地址。
协议采用 VLESS、RAW/TCP、REALITY 与 `xtls-rprx-vision`，使用 Chrome 指纹并禁用 mux。

```bash
pig build proxy client 'vless://UUID@proxy.example.com:443?encryption=none&security=reality&type=tcp&flow=xtls-rprx-vision&sni=www.sraoss.co.jp&fp=chrome&pbk=PUBLIC_KEY&sid=SHORT_ID&pqv=VERIFY_KEY'
pig build proxy client --server proxy.example.com:443 --id UUID --sni www.sraoss.co.jp --public-key PUBLIC_KEY --short-id SHORT_ID --pqv VERIFY_KEY
pig build proxy client --from ./client.uri
pig build proxy client --from ./client.uri --listen 127.0.0.1:8888
pig build proxy client --from ./client.uri --plan
```

连接输入三选一：位置参数 URI、直接连接参数，或 `--from FILE/-`。
`client.uri` 只是包含一条标准 `vless://` URI 的文本文件，可以有末尾换行，文件名任意。
`--from -` 从标准输入读取，不能混用不同输入方式。
服务端启用 ML-DSA-65 时，`--pqv` 提供验证密钥。
不支持的传输、重复 URI 参数及有歧义的输入会在设置前被拒绝。

**Linux** 需要 root 或 `sudo`、运行中的 systemd，以及已配置的 Pigsty `xray` 软件仓库，
例如先运行 `sudo pig repo add infra -u`。
配置使用 `/etc/xray.json`，权限 0640、所有者 `root:xray`；服务使用 `xray` 账户及 PIG 管理的 systemd drop-in。
**macOS** 使用普通用户运行，需要已有 Homebrew；配置使用 `~/.config/xray/pig-proxy.json`，
权限 0600，服务由独立的 `com.pigsty.xray-proxy` LaunchAgent 管理。

已有服务端可以将连接 URI 直接传给客户端：

```bash
# Linux 客户端：远端只读取并导出现有服务端连接。
ssh root@proxy.example.com 'pig build proxy server --host proxy.example.com --export-only --export -' | sudo pig build proxy client --from -
# macOS 客户端：不加 sudo。
ssh root@proxy.example.com 'pig build proxy server --host proxy.example.com --export-only --export -' | pig build proxy client --from -
```

重复运行会复用匹配的配置和服务。替换不同的已有客户端配置时，使用 `--replace --yes`；用 `--plan` 预览操作。

设置完成后，PIG 分别通过 HTTP 与 SOCKS 访问 `https://www.google.com/generate_204`，并要求返回 HTTP 204。检查失败时恢复原有配置与服务状态，已安装的软件包保留。

生成的 Shell 文件提供 `po`、`px`、`pck`。在当前 Shell 中启用代理变量：

```bash
# Linux
source /etc/profile.d/proxy.sh
po
# macOS
source ~/.config/xray/proxy.sh
po
```

### 服务端

服务端设置支持 Linux，需要相同的软件包、root 与 systemd 条件。
必须提供公网发布地址 `--host` 与 REALITY 伪装目标 `--target host:port`。
首次直连默认监听 `0.0.0.0:443`；`--port` 改变发布的公网端口及首次直连监听端口，
`--listen` 可指定独立的绑定地址。target 必须支持 TLS 1.3 与 HTTP/2。

```bash
# 直连公网入口，显式导出受保护的客户端 URI 文件。
sudo pig build proxy server --host proxy.example.com --target www.sraoss.co.jp:443 --export ./client.uri
# 已有可信 PROXY protocol 前端的后端：127.0.0.1:9443。
sudo pig build proxy server --host proxy.example.com --target www.sraoss.co.jp:443 --proxy-protocol --export ./client.uri
# 读取已有受支持服务端并输出连接，不更改部署。
sudo pig build proxy server --host proxy.example.com --export-only --export -
```

首次设置生成 UUID、X25519 密钥对、short ID，以及 ML-DSA-65 seed 和验证密钥。SNI 默认取 `--target` 的主机名。重复设置保留这些凭据。省略 `--listen` 时保留监听地址，指定时更新监听地址。匹配的配置不触发重启。需要更改目标或已有配置不受支持时，命令会报错。

连接导出选项：

| 选项 | 用途 |
|:---|:---|
| `--export FILE` | 将客户端 URI 写入权限为 `0600` 的文件；相同内容可重复导出，不覆盖不同内容 |
| `--export -` | 向标准输出写入完整客户端 URI，诊断写入标准错误 |
| `--export-only` | 读取现有配置并导出，不修改配置或服务；必须与 `--export` 一起使用 |
| `--plan` | 预览设置操作 |

导出必须使用文本输出。URI 包含客户端凭据，请通过文件或标准输入传递；导出内容不含服务端私钥或 seed。导出现有服务端时，`--host` 与 `--port` 指定客户端实际访问的公网地址，不是后端监听地址。

设置成功后，从真实客户端检查连通性。按网络环境开放防火墙或云安全组；`--proxy-protocol` 需要已有可信前端并使用回环监听。PIG 不会停止占用端口的其他服务或移除已有 V2Ray。

### 历史 VMess 形式

历史位置参数语法继续保留 V2Ray/VMess 行为，首次客户端端口仍为 12345：

```bash
pig build proxy user@host:8080
pig build proxy user@host:8080 127.0.0.1:1080
```

这个 Linux 兼容路径仍需要提供 `vray` 的仓库，写入 `/etc/v2ray.json` 与 `/etc/profile.d/proxy.sh`，
并重启 `v2ray`。需要 root 或 `sudo`；HTTPS 检查失败时返回非零。
远端用户 ID 属于凭据，普通结果会脱敏。

## build get

下载扩展源代码 tarball。

```bash
pig build get citus              # 单个扩展
pig build get citus pgvector     # 多个扩展
pig build get pdu pgdog          # 使用内置源码别名
pig build get citus -f           # 已存在时仍重新下载
pig build get citus -m           # 优先使用 pigsty.cc 中国镜像
```

`pig build get` 的参数是扩展名、包名或源码文件名；未知名称会按源码文件名处理。它不会把 `all` 或 `std` 展开为内置集合。需要批量下载时，请显式列出目标包名。

一些源码包并不直接对应扩展名，`pig build get` 内置了特殊别名 以便直接下载源码。

```bash
pig build get pdu                # 下载 pdu-3.0.25.12.tar.gz
pig build get pgdog              # 下载 pgdog-0.1.32.tar.gz
pig build get pgedge             # 同时下载 PostgreSQL 与 spock 源码
pig build get onesparse          # 同时下载 onesparse、graphblas、lagraph
```

当前常见的特殊源码别名 包括：`babelfish`、`agensgraph`、`orioledb`、`cloudberry`、`pgedge`、`pdu`、`pgdog`、`rdkit`、`onesparse`、`libpgfeutils`。

## build dep

安装构建扩展所需的依赖。

```bash
pig build dep citus              # 单个扩展
pig build dep citus pgvector     # 多个扩展
pig build dep citus --pg 17,16   # 为特定 PG 版本
```

**选项：**

- `--pg`：指定一个或多个 PostgreSQL 大版本；未指定时按扩展元数据或本机安装自动推断

指定 `--pg 17,16` 这样的版本列表时，带版本占位符的依赖包名会按每个选定版本展开，
去重后统一安装。`build pkg` 使用相同的依赖安装步骤。

## build ext

编译扩展并创建安装包。

`build ext` 和 `build pkg` 未指定 `--pg` 时，沿用扩展版本元数据和平台例外规则选择版本；
Debian/Ubuntu 还会参考本机安装的 PG 版本。包矩阵缺包或缺项不会移除构建目标，
显式 `--pg` 继续选择指定版本。
Debian/Ubuntu 构建保留原 `PATH` 的顺序；当前用户的 `~/.cargo/bin` 存在时，将其追加到末尾，
保留用户已选工具链的优先级。

调试包默认启用。包含编译产物的 RPM 构建通常会额外生成 `debuginfo` 与 `debugsource`，
基于 Debhelper 的构建通常会额外生成 `dbgsym`。纯 SQL、架构无关或采用特殊配方的包可能没有
可拆分的调试内容；包配方也可以做更窄的显式选择，PIG 不会改写 spec 或 `debian/rules`。

```bash
pig build ext citus              # 构建单个扩展
pig build ext citus pgvector     # 构建多个
pig build ext citus --pg 17      # 为特定 PG 版本
pig build ext citus --nodbg      # 显式省略自动调试包
```

**选项：**

- `--pg`：指定一个或多个 PostgreSQL 大版本
- `--nodbg`：在 RPM 构建中禁用自动 `debuginfo` / `debugsource` 包，在 DEB 构建中禁用
  `dbgsym` 包

## build pkg

对每个目标依次下载源码、安装依赖并构建。源码下载失败会停止该包的构建。依赖安装失败会告警后继续尝试构建，因为所需依赖可能已经安装。

```bash
pig build pkg citus              # 构建单个扩展
pig build pkg citus pgvector     # 构建多个
pig build pkg citus --pg 17,16   # 为多个 PG 版本
pig build pkg citus --nodbg      # 显式省略自动调试包
pig build pkg citus -m           # 优先使用 pigsty.cc 中国镜像下载源码
```

**选项：**

- `--pg`：指定一个或多个 PostgreSQL 大版本
- `--nodbg`：在 RPM 构建中禁用自动 `debuginfo` / `debugsource` 包，在 DEB 构建中禁用
  `dbgsym` 包
- `-m|--mirror`：下载源码时优先使用 `pigsty.cc` 镜像

## 常见工作流

### 构建并安装本地软件包

完成快速入门中的环境准备后，构建目标版本：

```bash
pig build pkg pg_partman --pg 18
find ~/ext/pkg -type f -name '*partman*'
```

从输出中选择与目标 PostgreSQL 版本和架构匹配的软件包，再通过系统包管理器安装。以下路径为占位符，请替换为实际文件路径：

```bash
sudo dnf install /absolute/path/to/package.rpm   # EL
sudo apt install /absolute/path/to/package.deb   # Debian / Ubuntu
```

`pig ext add NAME` 从已配置仓库安装软件包。安装本次构建产物时，使用上面的文件路径形式。

### 构建 Rust 扩展

在已准备好的构建环境中，为当前用户安装 Rust 和配方要求的 pgrx：

```bash
pig build rust
pig build pgrx --pg 18
pig build pkg pgmq --pg 18
```

需要中国镜像时，为 `build rust` 添加 `-m`。若配方要求特定 pgrx 版本，使用 `pig build pgrx -v VERSION --pg 18`。本地软件包按上一节安装。

### 构建多个 PostgreSQL 版本

```bash
pig build pkg citus --pg 15,16,17
```

结果按 PostgreSQL 主版本分别生成。所选版本需要匹配的开发包和构建配方支持。

### 构建 PostgreSQL 19 beta 扩展

```bash
pig build repo --beta
pig build tool --beta
pig build pgrx -b                # 仅 Rust/pgrx 扩展需要
pig build pkg EXTENSION --pg 19
```

将 `EXTENSION` 替换为支持该版本的扩展。这里的 beta 版本由 PIG 当前内置配置指定。

## 故障排除

| 问题 | 处理方式 |
|:---|:---|
| 缺少构建工具 | 运行 `pig build tool`；根据构建日志补齐配方要求的编译器 |
| 缺少扩展依赖 | 运行 `pig build dep NAME --pg MAJOR`，检查报错的软件包和仓库 |
| 找不到 PostgreSQL 头文件 | 安装 `pig ext add pg18-devel`，并用 `--pg 18` 选择对应构建目标 |
| pgrx 版本不匹配 | 按配方要求运行 `pig build pgrx -v VERSION --pg MAJOR` |
| 下载到旧源码 | 用 `pig build get NAME -f` 重新下载；先核对目录中的源码版本是否符合配方 |

构建日志位于 `~/ext/log/`。构建耗时取决于配方、CPU、内存、并行度和缓存；PostGIS、pg_duckdb 与 Rust 扩展通常需要更多资源。
