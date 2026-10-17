---
title: "命令参考"
description: pig CLI 命令参考概览
icon: fa-solid fa-terminal
weight: 600
---

`pig` 管理 PostgreSQL 软件包、本地数据库服务和 Pigsty 部署。按任务选择下列命令；用 `pig COMMAND --help` 查看当前程序支持的选项。

- [**pig repo**](/zh/pig/cmd/repo/)：管理软件仓库
- [**pig ext**](/zh/pig/cmd/ext/)：管理 PostgreSQL 扩展
- [**pig build**](/zh/pig/cmd/build/)：从源码构建扩展
- **pig install**：使用原生包管理器安装包，并对 PostgreSQL 别名做翻译
- [**pig sty**](/zh/pig/cmd/sty/)：管理 Pigsty 安装与 Grafana 仪表盘
- [**pig inventory**](/zh/pig/cmd/inventory/)：检视、编辑、校验与交换 Pigsty 配置清单
- [**pig do**](/zh/pig/cmd/do/)：执行 Pigsty 管理 playbook 任务
- [**pig pe**](/zh/pig/cmd/pe/)：访问 pg_exporter 指标与配置
- [**pig pg**](/zh/pig/cmd/pg/)：管理本地 PostgreSQL 服务器
- [**pig pt**](/zh/pig/cmd/pt/)：透明运行 patronictl，附带服务与配置辅助命令
- [**pig pb**](/zh/pig/cmd/pb/)：管理 pgBackRest 备份与恢复
- [**pig pitr**](/zh/pig/cmd/pitr/)：进行完整 PITR 工作流
- **pig context**：输出面向人工和 Agent 的环境上下文快照
- **pig license**：签发和验证 Pigsty 许可证
- **pig status / update / version**：查看环境、升级 pig、打印版本信息

## 全局选项

| 选项 | 用途 |
|:---|:---|
| `-h, --help` | 查看帮助 |
| `-H, --home PATH` | 选择 Pigsty 安装目录 |
| `-i, --inventory PATH` | 选择 Inventory 配置清单 |
| `-o, --output FORMAT` | 输出格式：`text`（默认）、`json`、`yaml`、`json-pretty` |
| `--debug` | 启用调试日志 |
| `--log-level LEVEL` | 日志级别：`debug`、`info`（默认）、`warn`、`error`、`fatal`、`panic` |
| `--log-path PATH` | 日志文件；默认写入终端 |

PIG 自己管理的结果支持结构化输出。原始内容命令（如 `inventory show`）和透传命令（如 `pt list`）保留原生输出，具体格式见对应参考页。

结构化结果将诊断信息与被调用工具的普通输出放在 stderr。JSON/YAML 输出及
`PIG_NON_INTERACTIVE=1` 禁用 PIG 命令适配器的提示输入，并使用 `sudo -n`；
它们不代替破坏性操作的确认，需要时仍应传入 `--yes`。
`pig pg psql -f -` 等明确的数据输入通道继续可用。原生 Patronictl 转发保留自身的交互与输出契约，
不接受 PIG 的结构化输出模式。

```bash
pig repo list
pig ext list vector -o json
pig help inventory edit
```

## pig repo

管理 PostgreSQL 软件包的 APT/YUM 仓库，详情请参考 [`pig repo`](/zh/pig/cmd/repo/)

```bash
pig repo list                    # 列出可用仓库
pig repo info   pgdg             # 显示仓库详情
pig repo status                  # 检查当前仓库状态
pig repo add    pgdg pigsty -u   # 添加仓库
pig repo rm     old-repo         # 移除仓库
pig repo update                  # 更新软件包缓存
pig repo create /www/pigsty      # 创建本地仓库
pig repo cache                   # 创建离线包
pig repo boot                    # 从离线包引导
```

## pig ext

管理 PostgreSQL 扩展和内核包，详情请参考 [`pig ext`](/zh/pig/cmd/ext/)

```bash
pig ext list    duck             # 搜索扩展
pig ext info    pg_duckdb        # 扩展详情
pig ext status                   # 显示已安装的扩展
pig ext add     pg_duckdb -y     # 安装扩展
pig ext rm      old_extension    # 移除扩展
pig ext update  pg_duckdb        # 更新指定扩展软件包
pig ext scan                     # 扫描已安装的扩展
pig ext import  pg_duckdb        # 下载以供离线使用
pig ext link    17               # 链接 PG 版本到 PATH
pig ext reload                   # 刷新扩展目录
```

## pig build

从源码构建 PostgreSQL 扩展，详情请参考 [`pig build`](/zh/pig/cmd/build/)

```bash
# 环境设置
pig build spec                   # 初始化构建规格
pig build repo                   # 设置仓库
pig build repo --beta            # 设置仓库，并额外启用 PostgreSQL 19 beta 仓库
pig build tool                   # 安装构建工具
pig build tool --beta            # 安装构建工具，并额外安装 PG19 beta 构建包
pig build rust -y                # 强制重装 Rust（默认不重装）
pig build rust -m                # 使用中国镜像安装 Rust，并写入 Cargo 镜像配置
pig build pgrx                   # 安装 PGRX 框架
pig build pgrx -b                # 自动探测时包含 PostgreSQL 19 beta pg_config
pig build proxy                  # 安装或检查 Xray
pig build proxy client --from ./client.uri
pig build proxy server --host proxy.example.com --target www.sraoss.co.jp:443 --export ./client.uri

# 构建扩展
pig build pkg citus              # 完整构建流程 = get + dep + ext
pig build get citus              # 下载源码
pig build dep citus              # 安装依赖
pig build ext citus              # 构建包
```

## pig install

使用系统原生包管理器安装软件包，并对 PostgreSQL 内核、扩展及常用别名做包名翻译。需要直接传递系统包名时，可使用 `-n/--no-translation`。

```bash
pig install pg_duckdb            # 安装扩展并自动翻译包名
pig install pg18                 # 安装 PostgreSQL 18 内核包组
pig install nginx htop vim       # 安装多个系统包
pig install unknown-package -n   # 禁用翻译，按原始包名安装
pig install pg18 --plan          # 预览安装计划
pig install vector -y         # 自动确认安装
```

## pig sty

安装 Pigsty 发行版，详情请参考 [`pig sty`](/zh/pig/cmd/sty/)

```bash
pig sty init                     # 安装 Pigsty 到 ~/pigsty
pig sty boot                     # 安装 Ansible 依赖
pig sty conf -g                  # 生成配置并替换演示密码
pig sty deploy                   # 运行部署 playbook
pig sty list                     # 列出可用 Pigsty 版本
pig sty get {{< param pigsty_version >}}                # 下载内置的 Pigsty 版本
pig sty grafana list             # 管理 Grafana 仪表盘（info/list/boot/load/init/dump/clean/lang/style）
```

## pig inventory

检视、编辑、校验、体检并与 CMDB 交换 Pigsty 配置清单（`pigsty.yml`），根级命令组，别名 `inv`，
详情请参考 [`pig inventory`](/zh/pig/cmd/inventory/)。

```bash
pig inventory status             # 检视当前生效的清单来源
pig inventory list               # 列出清单拓扑与取值类型
pig inventory show               # 原样显示清单 YAML（可能包含敏感信息）
pig inventory edit               # 在 $EDITOR 中编辑清单或选中片段
pig inventory validate           # 校验一份完整的静态 Pigsty 清单
pig inventory check              # 检查清单与本地控制节点
pig inventory diff other.yml     # 对比声明差异（不输出取值）
pig inventory cmdb check         # 检查 PostgreSQL CMDB（实验性）
```

## pig do

执行 Pigsty 管理任务，底层调用对应的 Ansible Playbook。完整命令参考见
[`pig do`](/zh/pig/cmd/do/)。

```bash
pig do pgsql-add  <cls> [ip...]       # 添加集群或实例
pig do pgsql-rm   <cls> [ip...]       # 移除集群或实例
pig do pgsql-db   <cls> <dbname>      # 创建或更新数据库
pig do pgsql-user <cls> <username>    # 创建或更新用户
pig do pgsql-ext  <cls> [ext...]      # 安装扩展
pig do node-pkg   <sel> [pkg...]      # 安装节点软件包
```

## pig pe

访问 pg_exporter 暴露的 PostgreSQL 监控指标，默认连接 `127.0.0.1:9630`；
端点与操作详情见 [`pig pe`](/zh/pig/cmd/pe/)。

```bash
pig pe list                    # 列出可用指标类型
pig pe get                     # 获取所有 pg_ 前缀指标
pig pe stat                    # 查看 exporter 统计信息
pig pe reload                  # 重载 pg_exporter 配置
pig pe --host 127.0.0.1 -p 9630 get
```

## pig license {#license}

`pig license issue <name>` 将新签发的许可证交付给调用方。文本输出包含 Token；JSON/YAML 在
`data.captured_output` 中保留 Token。成功签发属于明确的凭据交付操作，应将结果作为凭据保存。
签名私钥不会出现在结果的参数或诊断信息中，私钥加载的原始错误替换为固定文案。
参数错误和不含私钥的签名密钥失败原因在文本、JSON、YAML 中均保留。
`license status` 与 `license verify` 不回显 Token。

## pig context

输出环境上下文快照，覆盖主机、PostgreSQL、Patroni、pgBackRest 与已安装扩展。该命令适合排障和自动化脚本快速了解当前节点状态。

```bash
pig context                      # 文本输出
pig context -o json              # JSON 输出
pig context -m postgres          # 只输出 PostgreSQL 模块（默认包含 host）
pig context -m postgres,!host    # 排除 host 模块
```

## pig pg

管理本地 PostgreSQL 服务器，详情请参考 [`pig pg`](/zh/pig/cmd/pg/)

```bash
pig pg init                      # 初始化数据目录
pig pg start                     # 启动 PostgreSQL
pig pg stop                      # 停止 PostgreSQL
pig pg status                    # 查看状态
pig pg psql mydb                 # 连接数据库
pig pg ps                        # 查看当前连接
pig pg vacuum mydb               # 清理数据库
pig pg tune -p olap              # 生成调优参数
pig pg fork dev                  # 创建本地一次性物理副本
pig pg fork list                 # 列出本地 fork
pig pg log tail                  # 实时查看日志
```

## pig pt

透明运行 `patronictl` 管理 Patroni HA 集群，详情请参考 [`pig pt`](/zh/pig/cmd/pt/)

```bash
pig pt list pg-meta              # 列出集群成员（原生透传）
pig pt show-config pg-meta       # 显示集群动态配置（原生透传）
pig pt set ttl=60                # 本地语法糖：修改 Patroni/PG 配置
pig pt restart pg-test --pending # 应用待重启成员（原生透传）
pig pt status                    # 本地命令：综合服务与集群状态
pig pt log -f                    # 本地命令：实时查看日志
```

## pig pb

管理 pgBackRest 备份与恢复，详情请参考 [`pig pb`](/zh/pig/cmd/pb/)

```bash
pig pb info                      # 显示备份信息
pig pb ls                        # 列出所有备份
pig pb backup                    # 创建备份
pig pb backup full               # 全量备份
pig pb restore -d                # 恢复到最新
pig pb restore -t "2025-01-01"   # 恢复到指定时间
pig pb log tail                  # 实时查看日志
```

## pig pitr

执行编排式时间点恢复（PITR），详情请参考 [`pig pitr`](/zh/pig/cmd/pitr/)

```bash
pig pitr -d                      # 恢复到最新数据
pig pitr -t "2025-01-01 12:00:00+08"  # 恢复到指定时间
pig pitr -I                      # 恢复到备份一致性点
pig pitr -d --plan               # 显示执行计划（不实际执行）
pig pitr -d -y                   # 跳过确认（自动化）
```

## 辅助命令

```bash
pig status                       # 显示当前环境状态
pig status -o json               # 结构化状态输出
pig update                       # 将 pig 自身升级到最新版
pig update -m                    # 使用 pigsty.cc 镜像升级
pig update -v X.Y.Z              # 升级到已发布在当前软件仓库中的指定版本
pig version                      # 显示 pig 版本信息
pig version -o json              # 结构化版本输出
```

`pig update` 在受支持的 Linux 发行版上调用系统包管理器。macOS 请按[安装文档](/zh/pig/install/)
下载新版本压缩包并替换可执行文件。

如果 `~/.pig/config.yml` 存在语法错误，请先修正，再运行 `pig status` 等读取配置的命令。`pig version` 和 `pig pe` HTTP 命令不读取该文件，仍可使用。

## 设计记录

- [从面向人类到 Agent-Native](https://pig.pgsty.com/zh/design/agent-native-cli/)
- [为什么 PIG 保持扁平的 Cobra 命令层](https://pig.pgsty.com/zh/design/flat-cobra-command-layer/)
- [PIG 运维 CLI 安全契约](https://pig.pgsty.com/zh/design/ops-cli-safety/)
