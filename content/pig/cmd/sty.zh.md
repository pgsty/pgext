---
title: "CMD: sty"
description: 使用 pig sty 子命令管理 Pigsty 安装
icon: SquareTerminal
weight: 630
---

`pig sty` 命令管理 Pigsty 安装 - 一个功能齐全的 PostgreSQL 发行版，为您的 PostgreSQL 集群带来高可用、时间点恢复、全面监控、基础设施即代码以及 400+ 扩展支持。

------

## 子命令

| 命令 | 描述 |
|---------|--------------|
| `init`    | 安装 Pigsty（下载并解压） |
| `boot`    | 引导 Pigsty（安装 Ansible 和依赖） |
| `conf`    | 配置 Pigsty（生成配置文件） |
| `install` | 运行 pigsty install.yml playbook |
| `get`     | 下载 pigsty 源码压缩包 |
| `list`    | 列出可用的 pigsty 版本 |

------

## 概览

```bash
pig sty - 初始化（下载）、引导、配置和安装 Pigsty

  pig sty init    [-pfvd]      # 安装 pigsty（默认 ~/pigsty）
  pig sty boot    [-rpk]       # 安装 ansible 并准备离线包
  pig sty conf    [-civrsxn]   # 配置 pigsty 并生成配置
  pig sty install              # 使用 pigsty 安装和配置环境（危险！）
  pig sty get                  # 下载 pigsty 源码压缩包
  pig sty list                 # 列出可用的 pigsty 版本

用法:
  pig sty [command]

别名:
  sty, s, pigsty

示例:
  快速入门: https://pigsty.io/docs/setup/install/
  pig sty init                 # 解压并初始化 ~/pigsty
  pig sty boot                 # 安装 ansible 和其他依赖
  pig sty conf                 # 生成 pigsty.yml 配置文件
  pig sty install              # 运行 pigsty/install.yml playbook

可用命令:
  boot        引导 Pigsty
  conf        配置 Pigsty
  get         下载 pigsty 可用版本
  init        安装 Pigsty
  install     运行 pigsty install.yml playbook
  list        列出 pigsty 可用版本

参数:
  -h, --help   帮助信息

全局参数:
      --debug              启用调试模式
  -H, --home string        pigsty 主目录路径
  -i, --inventory string   配置清单路径
      --log-level string   日志级别: debug, info, warn, error, fatal, panic (默认 "info")
      --log-path string    日志文件路径，默认为终端
```

------

## 快速入门

完整的 Pigsty 安装工作流：

```bash
# 步骤 1：安装 Pigsty 到 ~/pigsty
pig sty init

# 步骤 2：使用 Ansible 引导
cd ~/pigsty
pig sty boot

# 步骤 3：生成配置
pig sty conf

# 步骤 4：运行安装 playbook
pig sty install
```

完成后，您将拥有一个完全配置的 PostgreSQL 环境，包括：
- 具备高可用的 PostgreSQL 集群
- 监控栈（Grafana、Prometheus、Loki）
- 备份和恢复工具
- Web 控制台：http://localhost:3000

------

## 什么是 Pigsty？

Pigsty 是一个全面的 PostgreSQL 发行版，提供：

### 核心特性
- **高可用**：Patroni 管理的自动故障转移
- **时间点恢复**：pgBackRest 用于备份和 PITR
- **监控栈**：Grafana 仪表板、Prometheus 指标、Loki 日志
- **基础设施即代码**：基于 Ansible 的部署和配置
- **扩展库**：400+ 预编译的 PostgreSQL 扩展

### 组件
- **PostgreSQL**：核心数据库，支持多个主版本
- **Patroni**：高可用编排器，用于自动故障转移
- **HAProxy**：负载均衡器，用于连接池化
- **pgBouncer**：连接池
- **pgBackRest**：备份和恢复解决方案
- **监控**：Grafana、Prometheus、Loki、AlertManager
- **管理**：Ansible playbooks 用于所有操作

------

## 命令参考

### `sty init`

下载并安装 Pigsty 到您的系统。

```bash
pig sty init                     # 安装到默认位置 ~/pigsty
pig sty init -p /opt/pigsty      # 自定义安装路径
pig sty init -v 3.0.4            # 指定版本
pig sty init -d                  # 仅下载，不解压
pig sty init -f                  # 强制覆盖现有安装
```

**选项：**
- `-p, --path`：安装路径（默认：`~/pigsty`）
- `-v, --version`：要安装的 Pigsty 版本
- `-d, --download`：仅下载不解压
- `-f, --force`：强制覆盖现有安装

**功能说明：**
1. 下载 Pigsty 源码压缩包
2. 解压到指定目录
3. 设置初始目录结构
4. 准备引导阶段

**示例：**
```bash
# 全新安装
pig sty init

# 检查安装
cd ~/pigsty
ls -la
```

------

### `sty boot`

通过安装 Ansible 和系统依赖来引导 Pigsty。

```bash
pig sty boot                     # 使用在线包引导
pig sty boot -r                  # 同时设置仓库
pig sty boot -p                  # 使用离线包（如果可用）
pig sty boot -k                  # 安装后保留包
```

**选项：**
- `-r, --repo`：设置系统仓库
- `-p, --pkg`：使用离线包（`/tmp/pkg.tgz`）
- `-k, --keep`：保留下载的包

**安装内容：**
- Ansible 和依赖
- Python 模块
- 系统工具
- PostgreSQL 客户端工具

**前提条件：**
- 必须先初始化 Pigsty（`pig sty init`）
- 需要 root 或 sudo 权限安装系统包

**示例：**
```bash
# 标准引导
cd ~/pigsty
pig sty boot

# 使用离线包
pig sty boot -p

# 验证 Ansible 安装
ansible --version
```

------

### `sty conf`

基于系统检测生成 Pigsty 配置文件。

```bash
pig sty conf                     # 生成默认配置
pig sty conf -c                  # 清洁模式（最小配置）
pig sty conf -i 192.168.1.10    # 指定 IP 地址
pig sty conf -v 17               # PostgreSQL 版本
pig sty conf -r                  # 包含区域设置
pig sty conf -s                  # 单节点设置
pig sty conf -x                  # 专家模式（所有参数）
pig sty conf -n mycluster       # 自定义集群名称
```

**选项：**
- `-c, --clean`：生成最小配置
- `-i, --ip`：主 IP 地址
- `-v, --version`：PostgreSQL 主版本（默认：17）
- `-r, --region`：包含区域特定设置
- `-s, --single`：单节点配置
- `-x, --expert`：包含所有配置参数
- `-n, --name`：集群名称

**生成文件：** `~/pigsty/pigsty.yml`

**配置内容：**
- 集群拓扑
- PostgreSQL 设置
- 监控配置
- 网络参数
- 扩展列表

**示例：**
```bash
# 为单节点生成配置
cd ~/pigsty
pig sty conf -s

# 自定义集群配置
pig sty conf -n prod-cluster -v 16 -i 10.10.10.10

# 查看生成的配置
cat pigsty.yml
```

------

### `sty install`

运行 Pigsty 安装 playbook 来配置环境。

```bash
pig sty install                  # 运行完整安装
```

**⚠️ 警告**：此命令将对您的系统进行重大更改：
- 安装 PostgreSQL 集群
- 配置监控栈
- 设置网络服务
- 修改系统设置

**前提条件：**
- Pigsty 已引导（`pig sty boot`）
- 已生成配置（`pig sty conf`）
- 足够的系统资源
- Root 或 sudo 访问权限

**功能说明：**
1. 验证配置
2. 安装基础设施组件
3. 部署 PostgreSQL 集群
4. 设置监控栈
5. 配置备份解决方案
6. 初始化 Web 控制台

**安装后访问：**
- Grafana：http://localhost:3000（admin/pigsty）
- Prometheus：http://localhost:9090
- AlertManager：http://localhost:9093
- PostgreSQL：localhost:5432

**示例：**
```bash
cd ~/pigsty
pig sty install

# 检查服务状态
systemctl status patroni
systemctl status grafana-server

# 访问 Web 控制台
open http://localhost:3000
```

------

### `sty get`

下载 Pigsty 源码压缩包但不安装。

```bash
pig sty get                      # 下载最新版本
pig sty get -v 3.0.4             # 指定版本
pig sty get -o /tmp/pigsty.tgz  # 自定义输出路径
```

**选项：**
- `-v, --version`：要下载的版本
- `-o, --output`：输出文件路径

**示例：**
```bash
# 下载最新版本
pig sty get

# 下载指定版本
pig sty get -v 3.0.3

# 检查下载的文件
ls -lh pigsty*.tgz
```

------

### `sty list`

列出所有可用的 Pigsty 版本。

```bash
pig sty list                     # 列出所有版本
```

**输出内容：**
- 版本号
- 发布日期
- 下载大小
- 主要特性

**示例：**
```bash
pig sty list

# 输出：
# 可用的 Pigsty 版本：
# - 3.0.4 (最新) - 2024-10-28
# - 3.0.3 - 2024-09-15
# - 3.0.2 - 2024-08-30
# ...
```

------

## 常用工作流

### 工作流 1：单节点安装

```bash
# 完整单节点设置
pig sty init
cd ~/pigsty
pig sty boot
pig sty conf -s
pig sty install

# 访问服务
open http://localhost:3000
```

### 工作流 2：生产集群设置

```bash
# 1. 在管理节点初始化
pig sty init
cd ~/pigsty

# 2. 引导
pig sty boot -r

# 3. 为生产配置
pig sty conf -x -n prod-db

# 4. 编辑配置
vim pigsty.yml
# 添加集群节点，调整参数

# 5. 部署
pig sty install
```

### 工作流 3：离线安装

```bash
# 在联网机器上：
pig sty get
pig repo cache
# 将 pigsty.tgz 和 pkg.tgz 传输到离线机器

# 在离线机器上：
pig sty init -p ~/pigsty
cd ~/pigsty
pig sty boot -p
pig sty conf
pig sty install
```

### 工作流 4：开发环境

```bash
# 快速开发设置
pig sty init
cd ~/pigsty
pig sty boot
pig sty conf -c -s  # 最小单节点配置
pig sty install

# 安装额外扩展
pig ext add pg_duckdb postgis pgvector
```

------

## 配置示例

### 单节点配置

```yaml
# 生成命令：pig sty conf -s
all:
  children:
    infra:
      hosts:
        10.10.10.10: { infra_seq: 1 }
    etcd:
      hosts:
        10.10.10.10: { etcd_seq: 1 }
    pg-meta:
      hosts:
        10.10.10.10: { pg_seq: 1, pg_role: primary }
      vars:
        pg_cluster: pg-meta
        pg_databases:
          - { name: meta }
        pg_users:
          - { name: dbuser_meta, password: DBUser.Meta }
```

### 多节点高可用集群

```yaml
# 三节点高可用集群
all:
  children:
    pg-prod:
      hosts:
        10.10.10.11: { pg_seq: 1, pg_role: primary }
        10.10.10.12: { pg_seq: 2, pg_role: replica }
        10.10.10.13: { pg_seq: 3, pg_role: replica }
      vars:
        pg_cluster: pg-prod
        pg_version: 17
        pg_vip_enabled: true
        pg_vip_address: 10.10.10.10/24
        pg_vip_interface: eth0
```

------

## 安装后任务

### 访问服务

```bash
# Grafana 仪表板
open http://localhost:3000
# 默认：admin/pigsty

# PostgreSQL
psql -h localhost -U postgres

# 检查集群状态
patronictl -c /etc/patroni/pg-meta.yml list
```

### 管理集群

```bash
# 添加新数据库
cd ~/pigsty
./pgsql-db.yml -l pg-meta

# 添加新用户
./pgsql-user.yml -l pg-meta

# 创建备份
pg_basebackup -h localhost -U replicator -D /tmp/backup
```

### 监控

```bash
# 查看指标
open http://localhost:3000/d/pgsql-overview

# 检查告警
open http://localhost:9093

# 查看日志
open http://localhost:3000/d/pgsql-logs
```

------

## 故障排除

### 安装失败

```bash
# 检查前提条件
cd ~/pigsty
./configure -n

# 验证 Ansible
ansible --version

# 检查日志
tail -f /tmp/pigsty-install.log
```

### 连接问题

```bash
# 检查服务
systemctl status postgresql
systemctl status patroni
systemctl status haproxy

# 测试连接
pg_isready -h localhost
```

### 配置问题

```bash
# 验证配置
cd ~/pigsty
ansible-playbook install.yml --check

# 重新生成配置
pig sty conf -x
```

### 服务恢复

```bash
# 重启 Patroni 集群
systemctl restart patroni

# 重载 HAProxy
systemctl reload haproxy

# 重启监控
systemctl restart prometheus grafana-server
```

------

## 最佳实践

1. **先测试**：在生产部署前始终在开发环境测试

2. **备份配置**：保留 `pigsty.yml` 和自定义配置的副本

3. **监控资源**：确保足够的 CPU、内存和磁盘空间

4. **使用版本控制**：在 Git 中跟踪配置更改

5. **定期更新**：保持 Pigsty 和扩展更新

6. **安全第一**：安装后立即更改默认密码

7. **规划拓扑**：部署前设计集群拓扑

8. **记录更改**：保留配置修改记录

9. **实践恢复**：定期测试备份和恢复过程

10. **加入社区**：参与 Pigsty 社区获取支持

------

## 资源

- **文档**：https://pigsty.io/docs/
- **GitHub**：https://github.com/Vonng/pigsty
- **演示**：https://demo.pigsty.cc
- **社区**：https://github.com/Vonng/pigsty/discussions