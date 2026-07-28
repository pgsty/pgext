## 用法

来源：

- [官方上游 README](https://github.com/plumqqz/mbus4/blob/fb315dd9b5f585798e92a75ac8e09713bc8f39a3/readme.rus)
- [官方扩展控制文件 (mbus.control)](https://github.com/plumqqz/mbus4/blob/fb315dd9b5f585798e92a75ac8e09713bc8f39a3/mbus.control)
- [官方扩展 SQL (mbus--1.1.sql)](https://github.com/plumqqz/mbus4/blob/fb315dd9b5f585798e92a75ac8e09713bc8f39a3/mbus--1.1.sql)

`mbus` — 一个简单的消息总线用于 pg。当应用程序需要此特定数据库功能时，请使用它。在安装扩展及其依赖项并验证它们之前，请勿集成到应用程序 SQL 中。

### 核心工作流

```sql
CREATE EXTENSION mbus;
```

在目标数据库中安装扩展，当可用时运行最小的上游示例，并在将其集成到应用程序 SQL 之前验证已安装的版本和返回值。

### 重要对象

- `can_post(qname text)` 是一个扩展函数并返回 `boolean`。
- `clear_tempq()` 是一个扩展函数并返回 `void`。
- `consume` 是一个扩展函数。
- `consume_temp(tqname text)` 是一个扩展函数并返回 `SETOF`。
- `create_consumer` 是一个扩展函数。
- `create_run_function(qname text)` 是一个扩展函数并返回 `void`。
- `create_temporary_consumer` 是一个扩展函数。
- `create_temporary_queue()` 是一个扩展函数并返回 `text`。
- `create_trigger` 是一个扩展函数。
- `create_view(qname text, cname text default 'default', sname text default 'public', viewname text default null)` 是一个扩展函数并返回 `void`。
- `create_view_prop(qname text, cname text, sname text, viewname text, with_delay boolean default false, with_expire boolean default false)` 是一个扩展函数并返回 `void`。
- `drop_consumer(cname text, qname text)` 是一个扩展函数并返回 `void`。
- `drop_queue(qname text)` 是一个扩展函数并返回 `void`。
- `drop_trigger(src text, dst text)` 是一个扩展函数并返回 `void`。

### 要求与注意事项

- 控制文件声明默认版本为 `1.1`。
- 请首先安装并验证确认的扩展依赖项：`hstore`。
- 控制文件将该扩展标记为不可重定位。
- 控制文件要求超级用户进行安装。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况与固定源的一致性。
