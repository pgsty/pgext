## 用法

来源：

- [发行日志](https://api.pgxn.org/src/pg_improv/pg_improv-0.13.0/CHANGELOG.md)

- [官方资料](https://api.pgxn.org/src/pg_improv/pg_improv-0.13.0/crates/pg_improv/README.md)
- [官方资料](https://api.pgxn.org/src/pg_improv/pg_improv-0.13.0/crates/pg_improv/pg_improv.control)
- [官方资料](https://api.pgxn.org/src/pg_improv/pg_improv-0.13.0/crates/pg_improv/Cargo.toml)
- [SQL 升级脚本](https://api.pgxn.org/src/pg_improv/pg_improv-0.13.0/crates/pg_improv/sql/pg_improv--0.12.2--0.13.0.sql)

`pg_improv` 0.13.0 将 PostgreSQL 表绑定到多维模型，计算公式并物化结果。它支持 PostgreSQL 15–18，安装需要超级用户。以下示例使用手工同步模式，无需预加载库。

### 核心流程

```sql
CREATE EXTENSION pg_improv;
CREATE TABLE sales (quarter text, price float8, qty float8);
INSERT INTO sales VALUES ('Q1', 10, 100), ('Q2', 12, 200);
SELECT improv.create_model('plan');
SELECT improv.set_sync_mode('plan', 'manual');
SELECT improv.bind_category('plan', 'Time', 'sales', 'quarter');
SELECT improv.bind_measure('plan', 'Price', 'sales', 'price');
SELECT improv.bind_measure('plan', 'Quantity', 'sales', 'qty');
SELECT improv.define('plan', 'Revenue = Price * Quantity');
SELECT improv.recalc('plan');
SELECT * FROM improv_plan.revenue;
SELECT * FROM improv.status('plan');
```

先绑定类别，再绑定对应度量。重复坐标需要显式聚合，空值不会生成模型单元。结果表位于模型专属模式中。`improv.define` 可用 `replace => true` 替换公式，`improv.recalc` 返回发生变化的结果行数。

### 同步模式

`improv.set_sync_mode` 选择手工、触发器、提交或逻辑同步。触发器模式无需预加载；提交模式需要会话或服务器预加载。默认的逻辑模式要求配置 `shared_preload_libraries`、重启服务器、启用 `wal_level=logical`，并正确设置工作进程使用的 `improv.database`。逻辑同步结果异步可见；事务内需要立即看到当前结果时，应显式重算。

监控 `improv.status`，并通过 `max_slot_wal_keep_size` 限制保留的 WAL。工作进程停滞时，逻辑复制槽可能持续保留 WAL。绑定表可能创建发布并修改复制身份，启用逻辑模式前应评估这些影响。

### 公式与维护边界

`improv` 语言支持标量公式与模型表达式，但不能执行任意 SQL 调用或外部 I/O。公式错误通常转换为空值，`improv.raise_errors` 可以改变该策略。数值和大整数通过浮点运算求值，超过 2^53 的精度需要特别注意。`improv.max_cells` 限制载入的输入单元数。

使用数据库转储备份模型定义与结果状态，并演练恢复；逻辑模式恢复后需要新建复制槽并重算。升级库时需重启已预加载的进程，并更新扩展 SQL。模型计算结果与应用事务中的权威源表是不同层次的数据。

### 0.13.0 的模型访问策略

模型以所有者权限运行。向模型创建者授予 `improv_user`，并赋予必要的源表权限。0.13.0 通过 `improv.set_policies` 添加 Cedar 策略，以 `improv.policies` 查看策略，并用 `improv.explain_access` 解释权限判定。

```sql
SELECT * FROM improv.policies('plan');
SELECT * FROM improv.explain_access('plan', 'Read');
```

向受限用户授予 `improv_reader`（或 `improv_user`），并单独授予结果表访问权限。策略过滤结果行以及 `improv.categories`、`improv.items`、`improv.measures`、`improv.cells` 函数的结果。超级用户、所有者以及有权切换为所有者的角色绕过这些策略。未设置策略的模型保留普通授权行为；设置策略后默认拒绝，策略失效时也拒绝受限用户访问。

派生值先使用所有者可见的全部输入计算，再过滤结果。隐藏某项输入不会从用户可见的合计中扣除该项贡献。模型策略也不限制直接访问源表。获准执行的单元格编辑和策略修改记录在 `improv.policy_audit` 中，拒绝的操作记录到服务器日志。开放模型前应检查授权与合计值的信息边界。

### 升级到 0.13.0

安装匹配文件；如果库已预加载，先重启 PostgreSQL，再逐库更新：

```sql
ALTER EXTENSION pg_improv UPDATE TO '0.13.0';
```

0.12.2 到 0.13.0 的迁移添加策略对象与函数，不修改模型数据行。本版还修复了升级与逻辑同步工作进程之间的死锁。应备份角色和模型数据并演练升级。上游现在使用 pgrx 0.19.3，这不代表发行仓库已经提供匹配的软件包。
