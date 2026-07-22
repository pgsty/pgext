## 用法

来源：

- [官方上游文档](https://github.com/ari-becker/pgcuid2/blob/0782b48bcb6eedbfc8785eca9c1e5b5f6215a22f/README.md)
- [官方 Rust 实现](https://github.com/ari-becker/pgcuid2/blob/0782b48bcb6eedbfc8785eca9c1e5b5f6215a22f/src/lib.rs)
- [官方 Rust 软件包清单](https://github.com/ari-becker/pgcuid2/blob/0782b48bcb6eedbfc8785eca9c1e5b5f6215a22f/Cargo.toml)

`pgcuid2` 通过 Rust `cuid2` 库添加一个生成 CUID2 文本标识符的 PostgreSQL 函数。只有应用明确需要 CUID2 兼容文本键时才应使用；结果不是 PostgreSQL `uuid`，而应视为不透明、区分大小写的字符串。

### 核心流程

安装由 pgrx 构建的扩展，并将生成器用作列默认值：

```sql
CREATE EXTENSION pgcuid2;

SELECT cuid2_create_id();

CREATE TABLE account (
  id text PRIMARY KEY DEFAULT cuid2_create_id(),
  display_name text NOT NULL
);
```

已复核版本返回 24 字符文本值。`cuid2_create_id()` 声明为 `VOLATILE`、parallel safe 和 security invoker，因此 PostgreSQL 会为每个插入行求值，而不会将它折叠成常量。

### 数据模型说明

应将生成值存入 `text`，或带有与固定实现一致的长度检查的域。不要将其转换为 `uuid`，不要从中推导创建时间，也不要把它当成时序标识符排序。文本主键比整数或 UUID 键更宽，会增加索引、外键和缓存占用；大表应对这一成本进行基准测试。

### 兼容性与维护

已归档的上游发布使用 pgrx 0.11，并声明了 PostgreSQL 11 至 16 构建特性。应同时固定扩展与 Rust 依赖版本，针对准确的 PostgreSQL 大版本构建，并在迁移前测试唯一性与输出格式。函数把生成工作委托给依赖库；不要作出超出固定 CUID2 实现文档的不可预测性或冲突保证。如果外部服务验证特定 CUID2 修订，应在应用测试套件中加入跨语言样例。
