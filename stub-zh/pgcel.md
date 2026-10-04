## 用法

来源：

- [pgcel.control](https://github.com/krashanoff/postgres-cel/blob/908363804df284f82a086526ab7e9fc98f98029f/pgcel.control)
- [README.md](https://github.com/krashanoff/postgres-cel/blob/908363804df284f82a086526ab7e9fc98f98029f/README.md)
- [Cargo.toml](https://github.com/krashanoff/postgres-cel/blob/908363804df284f82a086526ab7e9fc98f98029f/Cargo.toml)
- [src/lib.rs](https://github.com/krashanoff/postgres-cel/blob/908363804df284f82a086526ab7e9fc98f98029f/src/lib.rs)
- [tests/session_contracts.sh](https://github.com/krashanoff/postgres-cel/blob/908363804df284f82a086526ab7e9fc98f98029f/tests/session_contracts.sh)

`pgcel` 以 `celprogram` 类型保存 CEL 表达式源码，并对 JSONB 上下文求值，可用于 SQL 内可复用的判断或计算。实现会在初始化阶段登记共享内存，因此使用前需预加载库并重启。

### 核心用法

```ini
shared_preload_libraries = 'pgcel'
```

```sql
CREATE EXTENSION pgcel;
SELECT cel_eval(cel_compile('ctx.age >= 18'), '{"age": 21}'::jsonb);
SELECT cel_eval_json(cel_compile('ctx.a + ctx.b'), '{"a": 1, "b": 2}'::jsonb);
SET pgcel.program_cache_size = 2048;
```

### 运行边界

`cel_compile` 校验源码，`cel_eval` 要求结果为布尔值，`cel_eval_json` 将其他结果返回为 JSONB。输入对象通过 `ctx` 访问；无效表达式、缺失值或不兼容的结果类型可能报错。构建选项覆盖 PostgreSQL 14–17。control 标记为不可迁移且 superuser=false，仍需满足普通数据库和对象权限。`pgcel.program_cache_size` 默认每个后端缓存 1024 个编译程序，允许 1–65536。共享内存只保存校验哈希，不保存可执行程序，各后端仍独立编译，应据此规划连接池和缓存。这一早期版本不能替代 PostgreSQL 角色权限或 RLS。
