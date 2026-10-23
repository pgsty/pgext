## 用法

来源：

- [官方资料](https://api.pgxn.org/src/pg_improv/pg_improv-0.10.0/crates/pg_improv/README.md)
- [官方资料](https://api.pgxn.org/src/pg_improv/pg_improv-0.10.0/crates/pg_improv/pg_improv.control)
- [官方资料](https://api.pgxn.org/src/pg_improv/pg_improv-0.10.0/crates/pg_improv/Cargo.toml)

`pg_improv` 0.10.0 将 PostgreSQL 表绑定到多维模型，计算公式并物化结果。它支持 PostgreSQL 15–18，安装需要超级用户。以下示例使用手工同步模式，无需预加载库。

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
