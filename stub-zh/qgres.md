## 用法

来源：

- [官方扩展控制文件 (qgres.control)](https://api.pgxn.org/src/qgres/qgres-0.1.2/qgres.control)
- [官方扩展 SQL (qgres.sql)](https://api.pgxn.org/src/qgres/qgres-0.1.2/sql/qgres.sql)

`qgres` — 简单队列系统。当应用程序需要此特定数据库功能时使用它。在安装和验证其扩展依赖项之前，请勿使用它。

### 核心工作流

```sql
CREATE EXTENSION qgres;
```

在目标数据库中安装扩展，当可用时运行最小的上游示例，并在将其集成到应用程序 SQL 之前验证已安装的版本和返回值。

### 重要对象

- `consume(queue_id _queue.queue_id%TYPE , consumer_name _sp_consumer.consumer_name%TYPE , row_limit int DEFAULT 2^31-1)` 是一个扩展函数，返回 `TABLE`。
- `consume(queue_name _queue.queue_name%TYPE , consumer_name _sp_consumer.consumer_name%TYPE , row_limit int DEFAULT 2^31-1)` 是一个扩展函数，返回 `TABLE`。
- `consumer__drop(queue_name _queue.queue_name%TYPE , consumer_name _sp_consumer.consumer_name%TYPE)` 是一个扩展函数，返回 `void`。
- `consumer__register(queue_name _queue.queue_name%TYPE , consumer_name _sp_consumer.consumer_name%TYPE)` 是一个扩展函数，返回 `void`。
- `qgres_temp.build_add(first_arg text , call text , data_type regtype)` 是一个扩展函数，返回 `void`。
- `qgres_temp.build_publish(first_arg text , call text , data_type regtype)` 是一个扩展函数，返回 `void`。
- `qgres_temp.role__create(role_name name)` 是一个扩展函数，返回 `void`。
- `queue__drop(queue_name _queue.queue_name%TYPE , force boolean DEFAULT false)` 是一个扩展函数，返回 `void`。
- `queue__get(queue_id _queue.queue_id%TYPE)` 是一个扩展函数，返回 `queue`。
- `queue__get(queue_name _queue.queue_name%TYPE)` 是一个扩展函数，返回 `queue`。
- `queue__get_id(queue_name _queue.queue_name%TYPE)` 是一个扩展函数，返回 `int`。
- `queue_entry(bytea bytea DEFAULT NULL , jsonb jsonb DEFAULT NULL , text text DEFAULT NULL)` 是一个扩展函数，返回 `queue_entry`。
- `queue_entry` 是一个扩展定义的类型。
- `queue_type` 是一个扩展定义的类型。

### 要求与注意事项

- 审核后的控制文件声明默认版本为 `0.1.2`。
- 请先安装并验证确认的扩展依赖项：`plpgsql`, `citext`。
- 控制文件将扩展标记为可重定位。
- 控制文件不要求超级用户安装。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况与固定源的一致性。
