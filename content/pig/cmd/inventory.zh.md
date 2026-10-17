---
title: "pig inventory"
description: "检视、编辑、校验、体检并与 CMDB 交换 Pigsty 配置清单"
icon: fa-solid fa-terminal
weight: 645
---

`pig inventory`（别名 `pig inv`）查看、编辑和校验 Pigsty 配置清单 `pigsty.yml`，检查部署条件，并通过实验性的 `cmdb` 子命令与 PostgreSQL CMDB 交换配置。

编辑单个片段时，PIG 保留片段外的 YAML 内容与格式。保存前会重新解析完整文件；解析失败或原文件已被其他进程修改时，保留原文件并报告错误。

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

使用全局 `-i, --inventory` 选择配置清单，或通过 `-H, --home` 选择 Pigsty 安装目录。未指定时，PIG 从当前配置和 Pigsty 安装目录解析清单来源。除 `show` 原样输出 YAML 外，其他命令支持 `-o json` 或 `-o yaml` 结构化结果。

## 来源选择 {#source-selection}

PIG 按以下优先级选择第一个已配置的来源。选中的来源缺失或无效时报告该问题，
不会回退到另一份清单。

| 优先级 | 来源 | 相对路径的基准目录 |
|:---|:---|:---|
| 1 | `-i, --inventory PATH` | 当前工作目录 |
| 2 | `PIGSTY_INVENTORY` | 当前工作目录 |
| 3 | `~/.pig/config.yml` 中的 `inventory` | 已解析的 Pigsty 安装目录 |
| 4 | Pigsty 安装目录的 `ansible.cfg` 中的 `[defaults] inventory` | `ansible.cfg` 所在目录 |
| 5 | `pigsty.yml` | 已解析的 Pigsty 安装目录 |

PIG 在解析路径前展开 `~/`，包括 `ansible.cfg` 中带引号的路径。
`pig inv status` 只检查选中的来源，不执行动态清单。静态 Inventory 操作要求受支持的 YAML 来源。
[`pig do`](/zh/pig/cmd/do/) 可以将目录或可执行清单交给 Ansible；
[`pig sty deploy`](/zh/pig/cmd/sty/#sty-deploy) 要求普通文件。两者都显式向 Ansible 传入解析后的路径，
避免 Ansible 静默选择另一份默认清单。

## 快速入门

```bash
pig inventory status                     # 当前生效的清单来源是什么？
pig inventory list                       # 浏览清单拓扑
pig inventory edit pg-meta               # 只编辑 pg-meta 集群片段
pig inventory validate                   # 全量校验
pig inventory check -p ssh --sudo        # 附加 SSH/sudo 就绪探测
pig inventory diff /path/other.yml       # 与另一份清单对比声明
```

## inv status / list / show

```bash
pig inv status                          # 显示当前清单路径与来源类型
pig inv list --depth 2                  # 限制拓扑显示深度
pig inv list pg-meta                    # 查看一个集群的结构
pig inv show pg-meta                    # 原样输出集群 YAML
```

`status` 检查来源，不执行动态清单。`list` 显示分组、主机、键名与取值类型。`show` 输出实际取值，包括配置中的密码；加上 `-o json` 也不会改变其 YAML 输出。

## 选择器

`list` / `show` / `edit` 接受可选的 **选择器** 参数，定位清单中的一个片段：

```bash
pig inv e                                # 整个清单
pig inv e vars                           # all.vars 全局参数
pig inv e pg-meta                        # 一个集群分组
pig inv e pg-meta.vars.pg_databases      # 嵌套键路径
pig inv e 'pg-meta.hosts["10.10.10.10"]' # 含点的键（如 IP）用引号方括号
```

## inv edit

在 `$EDITOR` 中打开选中片段。保存退出后，PIG 恢复缩进和换行风格，重新解析完整文档，检查原文件是否被并发修改，再原子替换原文件。解析失败时会保留临时文件，供你修正。

```bash
pig inventory edit                       # 编辑整个清单
pig inv e pg-meta                        # 编辑一个片段
pig inv e vars --from vars.yml           # 从文件（或 stdin：-）替换选中片段
```

| 参数       | 说明                            |
|:---------|:------------------------------|
| `--from` | 从常规文件或 stdin（`-`）替换选中片段，跳过编辑器 |
{.full-width}

保存后的文件归当前用户所有，并移除组和其他用户的访问权限。通常权限变为 `0600`；原有更严格的仅属主权限会保留。结构化结果中的 `mode_tightened` 表示权限是否收紧。

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
| `--ansible` | 额外使用 `ansible-inventory` 解析 |
| `--timeout` | `--ansible` 兼容性校验超时（默认 10s）           |
{.full-width}

校验规则面向 Pigsty 配置，包括 `admin_ip`、`infra` 分组和主机地址要求。文件必须是单个 YAML 文档，且不能有重复键；这些解析规则同样适用于 `show` 和 `edit`。

`pg_users[].roles` 支持角色名字符串，也支持包含 `name` 的对象；可选字段为
`state`（`grant`、`revoke`、`absent`），以及允许 null 的布尔字段 `admin`、`set`、`inherit`。
编辑、序列化和 CMDB YAML 导出保留看起来像 YAML 1.1 标量的字符串类型，包括 `"1:30"`、`"010"`
以及作为字符串传入的日期和时间戳；必要的引号防止下游 YAML 解析器改变类型。已有的未加引号
或显式标记为时间戳的值保留原始内容，其标量类型为不透明的 `unknown`；需要字符串时应加引号。
`show` 和显式 CMDB 导出保留原始数据输出契约。

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

`--user` 和 `--port` 使用显式参数或上表中的默认值，不从 Inventory 推导有效 SSH 变量。可用逗号组合检查，例如 `--profile ansible,ssh`。

## inv diff

比较当前清单与另一份文件，输出声明的增加、删除和修改。结果包含路径与取值类型，不包含配置值：

```bash
pig inv diff other.yml
pig inv diff other.yml -o json
```

## inv cmdb（实验性）

`pig inventory cmdb` 通过原生 PostgreSQL 驱动，与 Pigsty 既有的 CMDB 模式
（`files/cmdb.sql`，`pigsty` / `pglog` schema）交换配置清单。连接解析顺序：
`-d/--database`（库名、URI 或 libpq conninfo）→ `METADB_URL` 环境变量 → `service=meta`。

| 子命令       | 说明                                                        |
|:----------|:----------------------------------------------------------|
| `check`   | 只读：验证 CMDB 投影，并可选校验与静态清单的一致性                              |
| `init`    | 应用 `cmdb.sql` 基线（覆盖既有 schema 需确认；`--plan` 预览） |
| `load`    | 用静态清单 **替换全部** CMDB 声明行（单事务；需确认；支持 `--plan`/`--strict`）|
| `dump`    | 将 CMDB 导出为静态清单文件（目标不同时需 `--force` 覆盖）                      |
| `enable`  | 切换 `ansible.cfg` 的 inventory 指向 CMDB（`inventory.sh`）    |
| `disable` | 切回静态 `pigsty.yml`（与 `enable` 均支持 `--plan`）         |
{.full-width}

```bash
pig inventory cmdb check                 # 验证 CMDB 状态
pig inventory cmdb init --plan           # 预览初始化计划
pig inventory cmdb load -y               # 将静态清单载入 CMDB
pig inventory cmdb dump -f pigsty.yml    # 将 CMDB 导出为静态清单
pig inventory cmdb enable                # 切换 ansible 清单源到 CMDB
pig inventory cmdb disable               # 切回静态 pigsty.yml
```

`init` 会执行 `cmdb.sql`，覆盖已有 `pigsty` / `pglog` schema；执行前先备份需要保留的数据。`load` 在单个事务中替换全部声明行。交互模式会要求确认目标；自动化和结构化输出模式使用 `--yes` 确认。用 `--plan` 查看操作内容。

## 设计记录

- [无损 Pigsty Inventory](https://pig.pgsty.com/zh/design/lossless-inventory/)
- [复用 Pigsty 已经拥有的 CMDB](https://pig.pgsty.com/zh/design/existing-cmdb-boundary/)
