## 用法

来源：

- [1.0.2 release](https://github.com/bitsinfra/gnomon/releases/tag/v1.0.2)
- [官方资料](https://github.com/bitsinfra/gnomon/blob/v1.0.2/README.md)
- [官方资料](https://github.com/bitsinfra/gnomon/blob/v1.0.2/gnomon.control)
- [官方资料](https://github.com/bitsinfra/gnomon/blob/v1.0.2/LICENSE)
- [官方资料](https://github.com/bitsinfra/gnomon/blob/v1.0.2/Makefile)
- [官方资料](https://github.com/bitsinfra/gnomon/blob/v1.0.2/docs/installation.md)
- [官方资料](https://github.com/bitsinfra/gnomon/blob/v1.0.2/docs/security.md)
- [官方资料](https://github.com/bitsinfra/gnomon/blob/v1.0.2/docs/operations.md)
- [官方资料](https://github.com/bitsinfra/gnomon/releases/download/v1.0.2/gnomon-pgtle--1.0.2.sql)

gnomon 1.0.2 是提供表历史、时态读取、分支及经过审核的三方合并的 SQL 扩展，支持 PostgreSQL 15–18。可选 gnomon_c 辅助扩展可加速并增强路由，SQL 核心可独立工作。

### 开始跟踪

```sql
CREATE EXTENSION gnomon;
SELECT gnomon.track('public.orders');
SELECT gnomon.coordinate_text();
```

示例假定 public.orders 已存在。跟踪需要短暂锁并安装存储、视图与触发器，不回填过去历史。安装要求超级用户并创建集群级角色。分支选择与路由会改变会话对被跟踪表的视图，应遵循官方 branch/use 流程，不直接写入内部存储。

### 权限与持久性

绝不能把内部 gnomon_reader、gnomon_reader_s 或 gnomon_search 角色授予应用登录。gnomon_merger 与 gnomon_admin 具备 BYPASSRLS，属于高权限操作角色。表权限、分支权限与经过审核的合并权限彼此独立。合并会检测冲突，撤销本身也是历史变更，并可能因后续修改而拒绝。

精确持久重放要求开启 synchronous_commit。Unlogged 临时分支不参与复制，也不能像日志化历史一样跨崩溃保留。保留策略决定可读取的历史范围，备份必须同时覆盖应用表和 Gnomon 存储。

### 时间功能与升级

时间解析、调度和保留操作要求通过受管理的工作进程、pg_cron 或 CLI 大约每秒运行 gnomon.heartbeat()。没有心跳时分支与历史仍可用，但时间功能会偏离。无法使用事件触发器时，DDL 后需执行 gnomon.sync()。已有 1.0.0 或 1.0.1 安装应按文档升级到 1.0.2 并保留数据。服务器扩展、pg_tle 与纯 SQL 安装各有升级路径，应选择与当前安装版本匹配的脚本。

### 1.0.2 修正

SQL 升级会修复 1.0.1 在分支创建早于表跟踪时保存的读取集合。合并后会恢复会话原先选择的分支。C 伴随扩展会拒绝在分支上读取系统列，避免此前的崩溃。先将 SQL 核心更新到 1.0.2，再更新已安装的匹配 C 扩展；替换预加载库后需重启使用它的进程。
