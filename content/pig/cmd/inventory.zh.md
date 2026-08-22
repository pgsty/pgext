---
title: "pig inventory"
description: "检视、编辑、校验、体检并与 CMDB 交换 Pigsty 配置清单"
icon: fa-solid fa-terminal
weight: 645
---

`pig inventory` 命令组（别名 `pig inv`）自 v1.6.0 起提供，用于以**无损**方式检视、编辑、
校验与体检 Pigsty 配置清单（`pigsty.yml`），并可通过实验性的 `cmdb` 子命令与 PostgreSQL CMDB 交换配置。

无损引擎会逐字节保留 YAML 的注释、格式、键序、锚点与换行风格；`edit` 在写盘前重新解析整个文档并原子写入，
**非法 YAML 不可能落盘**。

```bash
Inspect, edit, validate, check, and exchange Pigsty Inventory

Usage:
  pig inventory [command]

Aliases:
  inventory, inv

Available Commands:
  check       Check Inventory, controller, and opt-in target readiness
  cmdb        Exchange Inventory with Pigsty's existing PostgreSQL CMDB (experimental)
  diff        Compare declarations without emitting Inventory values
  edit        Edit Inventory or one selected YAML fragment
  list        List ordered Inventory topology and value kinds
  show        Show verbatim Inventory YAML (may contain secrets)
  status      Inspect the active Inventory source without executing it
  validate    Validate one complete static Pigsty Inventory
```

| 命令             | 别名   | 说明                            |
|:---------------|:-----|:------------------------------|
| `inv status`   |      | 检视当前生效的清单来源（不执行清单）            |
| `inv list`     | `ls` | 列出清单拓扑与取值类型（`--depth` 限制深度）   |
| `inv show`     |      | 原样显示清单 YAML（**可能包含敏感信息**）     |
| `inv edit`     | `e`  | 在 `$EDITOR` 中编辑清单或选中片段        |
| `inv validate` | `v`  | 校验一份完整的静态 Pigsty 清单           |
| `inv check`    | `ck` | 检查清单、控制器与目标节点就绪状态             |
| `inv diff`     |      | 对比两份清单的声明差异（不输出取值）            |
| `inv cmdb`     |      | 与 PostgreSQL CMDB 交换配置（**实验性**）|
{.full-width}

清单路径默认取自 pig 的配置解析（`-i/--inventory` 全局参数或 Pigsty 安装目录），
所有命令支持 `-o json|yaml` 结构化输出。


## 快速入门

```bash
pig inventory status                     # 当前生效的清单来源是什么？
pig inventory list                       # 浏览清单拓扑
pig inventory edit pg-meta               # 只编辑 pg-meta 集群片段
pig inventory validate                   # 全量校验
pig inventory check -p ssh --sudo        # 附加 SSH/sudo 就绪探测
pig inventory diff /path/other.yml       # 与另一份清单对比声明
```


## 选择器

`list` / `show` / `edit` 接受可选的**选择器**参数，定位清单中的一个片段：

```bash
pig inv e                                # 整个清单
pig inv e vars                           # all.vars 全局参数
pig inv e pg-meta                        # 一个集群分组
pig inv e pg-meta.vars.pg_databases      # 嵌套键路径
pig inv e 'pg-meta.hosts["10.10.10.10"]' # 含点的键（如 IP）用引号方括号
```


## inv edit

在 `$EDITOR` 中打开选中片段，保存退出后 pig 会：重新应用缩进与换行 → **重新解析整个文档**
（失败则中止并保留临时文件）→ 检查磁盘上的文件在编辑期间是否被并发修改 → 原子写入。

```bash
pig inventory edit                       # 编辑整个清单
pig inv e pg-meta                        # 编辑一个片段
pig inv e vars --from vars.yml           # 从文件（或 stdin：-）替换选中片段
```

| 参数       | 说明                            |
|:---------|:------------------------------|
| `--from` | 从常规文件或 stdin（`-`）替换选中片段，跳过编辑器 |
{.full-width}

> **注意**：编辑成功后，pig 会把清单文件权限收紧为 `0600`（清单可能包含数据库密码等敏感信息），
> 结果中的 `mode_tightened` 字段会予以提示。如有其他用户或工具直接读取该文件，请相应调整权限或属主。


## inv validate

校验一份完整的静态 Pigsty 清单：YAML 结构、Ansible 约定与 Pigsty 语义（如 `admin_ip`、`infra`
分组、主机 IPv4 键等）逐层检查，诊断信息不回显敏感取值。

```bash
pig inventory validate                   # 校验当前清单
pig inv v ./pigsty.yml                   # 校验指定文件
pig inv v --strict                       # 警告视为失败
pig inv v --ansible                      # 额外用 ansible-inventory 交叉解析
pig inv v --strict -o json               # 结构化输出（CI 友好）
```

| 参数          | 说明                                    |
|:------------|:--------------------------------------|
| `--strict`  | 把校验警告当作失败                             |
| `--ansible` | 额外使用带边界的 `ansible-inventory` 适配器交叉解析  |
| `--timeout` | `--ansible` 兼容性校验超时（默认 10s）           |
{.full-width}

> `validate` 是 **Pigsty 语义校验器**，不是通用 Ansible linter；其规则与 Pigsty 自带的
> `bin/validate` 保持对齐（个别地方更严格）。此外，清单解析对**重复键**与**多文档 YAML**
> 直接拒绝——此类文件连 `show` / `edit` 都无法使用。


## inv check

在静态校验之上做就绪体检：默认只检查清单与控制器本地条件，可用 `--profile` 追加探测。

```bash
pig inventory check                      # 清单 + 控制器基础检查
pig inv ck -p ansible                    # 追加 ansible 环境检查
pig inv ck -p network                    # 追加网络可达性探测
pig inv ck -p ssh --user admin           # 追加 SSH 登录探测
pig inv ck -p ssh --sudo                 # SSH 探测并验证 sudo -n
```

| 参数          | 简写   | 说明                              |
|:------------|:-----|:--------------------------------|
| `--profile` | `-p` | 追加探测：`ansible` / `network` / `ssh` |
| `--user`    |      | 显式 SSH 用户（默认控制器当前用户）            |
| `--port`    |      | TCP/SSH 目标端口（默认 22）             |
| `--sudo`    |      | 在 ssh 探测中额外验证 `sudo -n true`     |
{.full-width}


## inv cmdb（实验性）

> **实验性功能**：接口与行为可能变化。

`pig inventory cmdb` 通过原生 PostgreSQL 驱动，与 Pigsty 既有的 CMDB 模式
（`files/cmdb.sql`，`pigsty` / `pglog` schema）交换配置清单。连接解析顺序：
`-d/--database`（库名、URI 或 libpq conninfo）→ `METADB_URL` 环境变量 → `service=meta`。

| 子命令       | 说明                                                        |
|:----------|:----------------------------------------------------------|
| `check`   | 只读：验证 CMDB 投影，并可选校验与静态清单的一致性                              |
| `init`    | 应用 `cmdb.sql` 基线（存在既有 schema 时需 `--yes` 确认；支持 `--plan` 预览） |
| `load`    | 用静态清单**替换全部** CMDB 声明行（单事务；需 `--yes`；支持 `--plan`/`--strict`）|
| `dump`    | 将 CMDB 导出为静态清单文件（目标不同时需 `--force` 覆盖）                      |
| `enable`  | 守卫式切换 `ansible.cfg` 的 inventory 指向 CMDB（`inventory.sh`）    |
| `disable` | 切回静态 `pigsty.yml`（与 `enable` 均支持 `--plan`，原子写入可回滚）         |
{.full-width}

```bash
pig inventory cmdb check                 # 验证 CMDB 状态
pig inventory cmdb init --plan           # 预览初始化计划
pig inventory cmdb load -y               # 将静态清单载入 CMDB
pig inventory cmdb dump -f pigsty.yml    # 将 CMDB 导出为静态清单
pig inventory cmdb enable                # 切换 ansible 清单源到 CMDB
pig inventory cmdb disable               # 切回静态 pigsty.yml
```

> **注意**：`init` 会直接应用 `cmdb.sql` 基线，**不会**先备份既有 CMDB——对已有数据的 CMDB
> 执行前请自行备份；`load` 会替换全部声明行。破坏性操作均有基于目标指纹的确认门，
> 结构化输出模式下必须显式 `--yes`。
