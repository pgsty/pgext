## 用法

来源：

- [README](https://github.com/thomasdarimont/pgauthz/blob/64ffad245fdfed11cae42ccc3849082340b9b420/extensions/pg-cel/README.md)
- [Control](https://github.com/thomasdarimont/pgauthz/blob/64ffad245fdfed11cae42ccc3849082340b9b420/extensions/pg-cel/pg_cel.control)
- [Implementation](https://github.com/thomasdarimont/pgauthz/blob/64ffad245fdfed11cae42ccc3849082340b9b420/extensions/pg-cel/src/lib.rs)
- [Cargo](https://github.com/thomasdarimont/pgauthz/blob/64ffad245fdfed11cae42ccc3849082340b9b420/extensions/pg-cel/Cargo.toml)

`pg_cel` 0.2.0 为 pgauthz 求值 Common Expression Language 条件。它是 Rust/pgrx 共享库，默认构建 PostgreSQL 18，Cargo 声明 13–18 的特性开关。核心用法不需要预加载或重启。

### 求值条件

```sql
CREATE SCHEMA IF NOT EXISTS authz;
CREATE EXTENSION pg_cel SCHEMA authz;
SELECT authz.cel_compile_check('request.age >= 18');
SELECT authz.cel_eval_bool('request.age >= 18',
                           '{"request":{"age":21}}');
```

### 集成与失败边界

`cel_compile_check` 拒绝语法错误表达式；`cel_eval_bool` 根据 JSON 输入求值，返回布尔值，结果不是布尔值时返回 NULL。编译或求值错误会产生 SQL 错误。策略调用方必须将错误、NULL 以及非 true 结果视为拒绝；pgauthz 引擎会如此处理，并要求函数位于 `authz`，而扩展本身允许迁移模式。

表达式求值没有外部副作用，输入变量名和类型必须匹配 JSON 上下文。这是 pgauthz 0.19.0 的可选 CEL 求值器，不替代引擎授权模型及其迁移。
