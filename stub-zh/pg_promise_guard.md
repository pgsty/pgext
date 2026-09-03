## 用法

来源：

- [PGXN 0.1.0 README](https://pgxn.org/dist/pg_promise_guard/0.1.0/README.html)
- [pg_promise_guard 控制文件](https://api.pgxn.org/src/pg_promise_guard/pg_promise_guard-0.1.0/pg_promise_guard.control)
- [pg_promise_guard 0.1.0 SQL 定义](https://api.pgxn.org/src/pg_promise_guard/pg_promise_guard-0.1.0/pg_promise_guard--0.1.0.sql)
- [PostgreSQL 许可证](https://api.pgxn.org/src/pg_promise_guard/pg_promise_guard-0.1.0/LICENSE)

`pg_promise_guard` 用于报告看似存在、实际上完全或部分未执行的模式保证。它只读取 PostgreSQL 系统目录，适合在迁移、批量装载或运维临时绕过之后周期性检查完整性。

### 核心流程

```sql
CREATE EXTENSION pg_promise_guard;

SELECT * FROM promise_breaks;
SELECT promises_kept();
SELECT promises_kept('app');
SELECT * FROM check_promises('app');
```

`promise_breaks` 是全库视图。`check_promises(text)` 可以把结果限制在一个模式内，`promises_kept(text)` 只有在存在当前违约时才返回 false。

### 检查结果

- 无效唯一索引属于 `breach`，因为唯一性没有得到执行。
- 被禁用的用户触发器属于 `breach`，因为其声明行为没有生效。
- 已启用但未强制执行的行级安全属于 `breach`，因为表所有者仍可绕过策略。
- `NOT VALID` 检查或外键约束属于 `gap`：新行会被检查，但已有行尚未验证。

结果类型 `promise_break` 包含结果类型、对象、关系、严重度与解释。扩展自有对象会被跳过，使检查聚焦于应用模式状态。

### 边界

检查只读取目录，不扫描用户表，也不获取应用锁。它报告当前状态，不记录由谁或何时造成。它不验证索引物理完整性、权限、默认权限、`search_path` 遮蔽或历史执行情况；这些问题应使用 `amcheck` 等工具和独立安全审计处理。

PGXN 元数据声明支持 PostgreSQL 13 及以上，但上游只测试过 18.6 与 19 beta。在依赖结果之前，应自行验证 PostgreSQL 13 至 17。该扩展是纯 SQL，无需预加载、共享库或外部依赖。
