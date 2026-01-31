---
title: "pig pitr"
description: "使用 pig pitr 命令执行编排式时间点恢复（PITR）"
icon: Clock
weight: 670
---

`pig pitr` 命令用于执行**编排式时间点恢复**（Orchestrated Point-In-Time Recovery）。与 `pig pb restore` 不同，此命令会自动协调 Patroni、PostgreSQL 和 pgBackRest，完成完整的 PITR 工作流。

```bash
pig pitr - 执行编排式 PITR（自动管理 Patroni/PostgreSQL 生命周期）

此命令编排完整的 PITR 工作流：
  1. 停止 Patroni 服务（如果正在运行）
  2. 确保 PostgreSQL 已停止（带重试和降级策略）
  3. 执行 pgbackrest restore
  4. 启动 PostgreSQL
  5. 提供恢复后操作指引

恢复目标（至少需要指定一个）：
  --default, -d      恢复到 WAL 流末尾（最新数据）
  --immediate, -I    恢复到备份一致性点
  --time, -t         恢复到指定时间戳
  --name, -n         恢复到命名还原点
  --lsn, -l          恢复到指定 LSN
  --xid, -x          恢复到指定事务 ID

时间格式：
  - 完整格式: "2025-01-01 12:00:00+08"
  - 仅日期: "2025-01-01"（默认为 00:00:00）
  - 仅时间: "12:00:00"（默认为今天）

示例：
  pig pitr -d                      # 恢复到最新（最常用）
  pig pitr -t "2025-01-01 12:00"   # 恢复到指定时间
  pig pitr -I                      # 恢复到备份一致性点
  pig pitr -d --dry-run            # 仅显示执行计划
  pig pitr -d -y                   # 跳过确认（用于自动化）
  pig pitr -d --skip-patroni       # 跳过 Patroni 管理
  pig pitr -d --no-restart         # 恢复后不自动启动 PostgreSQL
```


## 命令概览

`pig pitr` 是一个高度自动化的恢复命令，它会：

1. 自动停止 Patroni 服务（如果正在运行）
2. 确保 PostgreSQL 已停止（带重试和降级策略）
3. 执行 pgBackRest 恢复
4. 启动 PostgreSQL
5. 提供恢复后的操作指引

**与 `pig pb restore` 的区别：**

| 特性 | `pig pitr` | `pig pb restore` |
|:----|:----------|:-----------------|
| 停止 Patroni | 自动 | 手动 |
| 停止 PostgreSQL | 自动（带重试） | 需要预先停止 |
| 启动 PostgreSQL | 自动 | 手动 |
| 恢复后指引 | 提供详细指引 | 无 |
| 适用场景 | 生产环境完整恢复 | 底层操作或脚本集成 |


## 快速入门

```bash
# 最常用：恢复到最新数据
pig pitr -d

# 恢复到指定时间点
pig pitr -t "2025-01-01 12:00:00+08"

# 恢复到备份一致性点（最快）
pig pitr -I

# 查看执行计划（不实际执行）
pig pitr -d --dry-run

# 跳过确认（用于自动化）
pig pitr -d -y

# 从特定备份集恢复
pig pitr -d -b 20251225-120000F

# 独立 PostgreSQL（非 Patroni 管理）
pig pitr -d --skip-patroni

# 恢复后不自动启动 PostgreSQL
pig pitr -d --no-restart
```


## 参数说明

### 恢复目标（必选其一）

| 参数 | 简写 | 说明 |
|:----|:----|:----|
| `--default` | `-d` | 恢复到 WAL 流末尾（最新数据） |
| `--immediate` | `-I` | 恢复到备份一致性点 |
| `--time` | `-t` | 恢复到指定时间戳 |
| `--name` | `-n` | 恢复到命名还原点 |
| `--lsn` | `-l` | 恢复到指定 LSN |
| `--xid` | `-x` | 恢复到指定事务 ID |

### 备份选择

| 参数 | 简写 | 说明 |
|:----|:----|:----|
| `--set` | `-b` | 从特定备份集恢复 |

### 流程控制

| 参数 | 简写 | 说明 |
|:----|:----|:----|
| `--skip-patroni` | `-S` | 跳过 Patroni 停止操作 |
| `--no-restart` | `-N` | 恢复后不自动启动 PostgreSQL |
| `--dry-run` | | 仅显示执行计划，不实际执行 |
| `--yes` | `-y` | 跳过确认倒计时 |

### 恢复选项

| 参数 | 简写 | 说明 |
|:----|:----|:----|
| `--exclusive` | `-X` | 排他模式：在目标前停止 |
| `--promote` | `-P` | 恢复后自动提升为主库 |

### 配置参数

| 参数 | 简写 | 说明 |
|:----|:----|:----|
| `--stanza` | `-s` | pgBackRest stanza 名称（自动检测） |
| `--config` | `-c` | pgBackRest 配置文件路径 |
| `--repo` | `-r` | 仓库编号（多仓库场景） |
| `--dbsu` | `-U` | 数据库超级用户（默认：`postgres`） |
| `--data` | `-D` | 目标数据目录 |


## 时间格式

`--time` 参数支持多种时间格式，自动补全时区：

| 格式 | 示例 | 说明 |
|:----|:----|:----|
| 完整格式 | `2025-01-01 12:00:00+08` | 包含时区的完整时间戳 |
| 仅日期 | `2025-01-01` | 自动补全为当天 00:00:00（当前时区） |
| 仅时间 | `12:00:00` | 自动补全为今天（当前时区） |


## 执行流程

### 第一阶段：预检查

- 验证恢复目标参数（必须且只能指定一个）
- 检查数据目录是否存在且已初始化
- 检测 Patroni 服务状态
- 检测 PostgreSQL 运行状态

### 第二阶段：停止 Patroni

如果 Patroni 服务正在运行且未指定 `--skip-patroni`：
- 执行 `systemctl stop patroni`
- 等待 PostgreSQL 随 Patroni 自动停止

### 第三阶段：确保 PostgreSQL 停止

采用渐进式策略确保 PostgreSQL 完全停止：

1. **等待自动停止**：Patroni 停止后等待 30 秒
2. **优雅停止**：使用 `pg_ctl stop -m fast`（最多重试 3 次，指数退避）
3. **立即停止**：使用 `pg_ctl stop -m immediate`
4. **强制终止**：使用 `kill -9`（最后手段）

### 第四阶段：执行恢复

调用 pgBackRest 执行实际的数据恢复：
```bash
pgbackrest restore --target-action=promote ...
```

### 第五阶段：启动 PostgreSQL

除非指定 `--no-restart`，否则自动启动 PostgreSQL：
- 等待启动完成（超时 120 秒）
- 验证进程确实运行

### 第六阶段：恢复后指引

显示详细的后续操作指引，包括：
- 如何验证恢复的数据
- 如何提升为主库
- 如何恢复 Patroni 集群管理
- 如何重新创建 pgBackRest stanza


## 使用示例

### 场景一：误删数据恢复

```bash
# 1. 查看可用的备份
pig pb info

# 2. 恢复到误删前的时间点
pig pitr -t "2025-01-15 09:30:00+08"

# 3. 验证数据
pig pg psql
SELECT * FROM important_table;

# 4. 确认无误后提升为主库
pig pg promote
```

### 场景二：恢复到最新状态

```bash
# 服务器故障后恢复到最新数据
pig pitr -d
```

### 场景三：快速恢复到备份点

```bash
# 恢复到备份一致性点（不需要回放 WAL）
pig pitr -I
```

### 场景四：自动化脚本

```bash
# 跳过所有确认，适合自动化
pig pitr -d -y
```

### 场景五：独立 PostgreSQL 实例

```bash
# 非 Patroni 管理的实例
pig pitr -d --skip-patroni
```

### 场景六：仅恢复不启动

```bash
# 恢复后手动检查，再决定是否启动
pig pitr -d --no-restart

# 检查恢复的数据目录
ls -la /pg/data/

# 手动启动
pig pg start
```


## 执行计划示例

执行 `pig pitr -d --dry-run` 会显示类似以下的执行计划：

```
══════════════════════════════════════════════════════════════════
 PITR Execution Plan
══════════════════════════════════════════════════════════════════

Current State:
  Data Directory:  /pg/data
  Database User:   postgres
  Patroni Service: active
  PostgreSQL:      running (PID: 12345)

Recovery Target:
  Latest (end of WAL stream)

Execution Steps:
  [1] Stop Patroni service
  [2] Ensure PostgreSQL is stopped
  [3] Execute pgBackRest restore
  [4] Start PostgreSQL
  [5] Print post-restore guidance

══════════════════════════════════════════════════════════════════

[Dry-run mode] No changes made.
```


## 恢复后操作

成功恢复后，命令会显示详细的后续操作指引：

```
══════════════════════════════════════════════════════════════════
 PITR Complete
══════════════════════════════════════════════════════════════════

[1] Verify recovered data:
   pig pg psql

[2] If satisfied, promote to primary:
   pig pg promote

[3] To resume Patroni cluster management:
   WARNING: Ensure data is correct before starting Patroni!
   systemctl start patroni

   Or if you want this node to be the leader:
   1. Promote PostgreSQL first: pig pg promote
   2. Then start Patroni: systemctl start patroni

[4] Re-create pgBackRest stanza if needed:
   pig pb create

══════════════════════════════════════════════════════════════════
```


## 安全机制

### 确认倒计时

除非使用 `--yes` 参数，命令执行前会显示 5 秒倒计时：

```
WARNING: This will overwrite the current database!
Press Ctrl+C to cancel, or wait for countdown...
Starting PITR in 5 seconds...
```

### 渐进式停止策略

为确保数据安全，停止 PostgreSQL 采用渐进式策略：
1. 先尝试优雅停止（保证数据一致性）
2. 失败后尝试立即停止
3. 最后才使用 kill -9（仅在极端情况）

### 恢复验证

恢复后自动验证 PostgreSQL 是否成功启动，如果失败会提示检查日志。


## 设计说明

**与其他命令的关系：**

- `pig pitr` 内部调用 `pig pt stop`、`pig pg stop`、`pig pg start` 和 `pig pb restore`
- 提供比单独命令更高级别的自动化协调
- 适合生产环境的完整 PITR 工作流

**错误处理：**

- 每个阶段都有详细的错误信息
- 失败时提示相关日志位置
- 支持中断后手动继续

**权限执行：**

- 如果当前用户已是 DBSU：直接执行命令
- 如果当前用户是 root：使用 `su - postgres -c "..."` 执行
- 其他用户：使用 `sudo -inu postgres -- ...` 执行

**平台支持：**

此命令专为 Linux 系统设计，依赖 Pigsty 的默认目录结构。
