## 用法

来源：

- [官方上游 README](https://github.com/fenoman/pg_xclaim/blob/e26926e06a9442b8554e6032e22a8080360d5d1d/README.md)
- [官方扩展控制文件 (pg_xclaim.control)](https://github.com/fenoman/pg_xclaim/blob/e26926e06a9442b8554e6032e22a8080360d5d1d/pg_xclaim.control)
- [官方扩展 SQL (pg_xclaim--1.0.0-rc1.sql)](https://github.com/fenoman/pg_xclaim/blob/e26926e06a9442b8554e6032e22a8080360d5d1d/sql/pg_xclaim--1.0.0-rc1.sql)

`pg_xclaim` — **实验性** PostgreSQL 扩展：一种自定义的事务性 claim 存储机制，使用自己的分区共享内存哈希表。当应用程序需要此特定数据库功能时，请使用此扩展。上游明确表示此扩展尚未准备好用于生产环境。

### 核心工作流

```sql
CREATE EXTENSION pg_xclaim;
```

在目标数据库中安装扩展，当可用时运行上游提供的最小示例，并在将其集成到应用程序 SQL 之前验证已安装的版本和返回值。

### 重要对象

- `xclaim.count()` 是一个扩展函数，返回 `int8`。
- `xclaim.debug()` 是一个扩展函数，返回 `TABLE`。
- `xclaim.debug_inject_stale(scope int4, key int4)` 是一个扩展函数，返回 `void`。
- `xclaim.debug_snapshot()` 是一个扩展函数，返回 `TABLE`。
- `xclaim.session_reset()` 是一个扩展函数，返回 `void`。
- `xclaim.stats()` 是一个扩展函数，返回 `TABLE`。
- `xclaim.try(classid int4, objid int4)` 是一个扩展函数，返回 `boolean`。
- `xclaim.try(key int8)` 是一个扩展函数，返回 `boolean`。
- `xclaim.try_many(classid int4, objids int4[])` 是一个扩展函数，返回 `boolean[]`。
- `xclaim.try_many(keys int8[])` 是一个扩展函数，返回 `boolean[]`。

### 要求与注意事项

- 控制文件声明默认版本为 `1.0.0-rc1`。
- 控制文件标记此扩展为不可重定位。
- 控制文件要求超级用户进行安装。
- 控制文件标记此扩展为不受信任。
- 上游明确表示该项目尚未准备好用于生产环境。
- 上游将项目的一部分或全部标记为实验性。
- 在生产使用之前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况，以匹配固定源。
