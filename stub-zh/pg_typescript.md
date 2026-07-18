## 用法

来源：

- [上游 README](https://github.com/isaacd9/pg_typescript/blob/e03ed5460decab37526e129f5577907b972b4898/README.md)
- [扩展 control 文件](https://github.com/isaacd9/pg_typescript/blob/e03ed5460decab37526e129f5577907b972b4898/pg_typescript.control)
- [Rust 包清单](https://github.com/isaacd9/pg_typescript/blob/e03ed5460decab37526e129f5577907b972b4898/Cargo.toml)
- [权限强制实现](https://github.com/isaacd9/pg_typescript/blob/e03ed5460decab37526e129f5577907b972b4898/src/permissions.rs)

`pg_typescript` `0.1.0` 版是 Alpha 阶段的过程语言，可在 PostgreSQL 16–18 中通过 Deno/V8 运行 TypeScript。它支持普通函数、远程或 npm 导入，以及通过 `_pg.execute` 访问数据库。每个后端会创建一个运行时；把库加入 `shared_preload_libraries` 只是可选预热，并非必需条件。

### 创建无依赖函数

```sql
LOAD 'pg_typescript';
CREATE EXTENSION pg_typescript;

CREATE FUNCTION slugify(title text)
RETURNS text
LANGUAGE typescript
IMMUTABLE STRICT
AS $$
  return title.trim().toLowerCase().replace(/\s+/g, '-');
$$;

SELECT slugify('Hello TypeScript');
```

control 文件把该语言标为可信，但安装扩展仅限超级用户，函数所有者还可通过用户级 `typescript.allow_*` 设置请求能力。超级用户配置的 `typescript.max_*` 上限才是真正的能力边界；上游默认允许导入和 `_pg.execute`，其他若干上限实际上没有封顶。授予语言使用权之前，应先设置严格的最大导入、网络、文件、环境、进程、FFI、系统与数据库执行策略。创建函数时会抓取远程模块并在本地缓存，因此每项依赖都要固定版本并审计。应把不可信 TypeScript 当作数据库代码，并在生产前测试内存、超时、后端崩溃、冷启动与供应链行为。
