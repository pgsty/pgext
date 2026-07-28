## 用法

来源：

- [官方上游 README](https://github.com/seanwevans/pg_os/blob/b822d1f3a83657eac385d2004421236d5e8b1d4d/README.md)
- [官方扩展控制文件 (pg_os.control)](https://github.com/seanwevans/pg_os/blob/b822d1f3a83657eac385d2004421236d5e8b1d4d/pg_os.control)
- [官方扩展 SQL (pg_os--1.0.sql)](https://github.com/seanwevans/pg_os/blob/b822d1f3a83657eac385d2004421236d5e8b1d4d/pg_os--1.0.sql)

`pg_os` — pg_os 是一个完全在数据库内部建模操作系统概念的 PostgreSQL 扩展。当应用程序需要此特定数据库功能时，请使用它。在安装扩展及其依赖项并验证它们之前，请勿集成到应用程序 SQL 中。

### 核心工作流

```sql
CREATE EXTENSION pg_os;
```

在目标数据库中安装扩展，如果有可用的最小上游示例，请运行该示例，并在将其集成到应用程序 SQL 之前验证已安装的版本和返回值。

### 重要对象

- `acquire_semaphore(process_id INTEGER, sem_name TEXT)` 是一个扩展函数，返回 `VOID`。
- `allocate_memory(user_id INTEGER, process_id INTEGER, segment_size INTEGER)` 是一个扩展函数，返回 `VOID`。
- `allocate_page(thread_id INTEGER)` 是一个扩展函数，返回 `BIGINT`。
- `assign_role_to_user(user_id INTEGER, role_id INTEGER)` 是一个扩展函数，返回 `VOID`。
- `change_file_permissions(user_id INTEGER, file_id INTEGER, new_perms TEXT)` 是一个扩展函数，返回 `VOID`。
- `check_mail(user_id INTEGER)` 是一个扩展函数，返回 `SETOF`。
- `check_permission(p_user_id INTEGER, p_resource_type TEXT, p_action TEXT)` 是一个扩展函数，返回 `BOOLEAN`。
- `cleanup_terminated_processes(timeout_interval INTERVAL DEFAULT '1 hour')` 是一个扩展函数，返回 `VOID`。
- `create_file(user_id INTEGER, filename TEXT, parent_id INTEGER, is_dir BOOLEAN DEFAULT FALSE)` 是一个扩展函数，返回 `INTEGER`。
- `create_mutex(mutex_name TEXT)` 是一个扩展函数，返回 `VOID`。
- `create_role(role_name TEXT)` 是一个扩展函数，返回 `INTEGER`。
- `create_semaphore(sem_name TEXT, initial_count INTEGER, max_val INTEGER)` 是一个扩展函数，返回 `VOID`。
- `create_user(name TEXT)` 是一个扩展函数，返回 `INTEGER`。
- `enqueue_io_request(device_name TEXT, request_type TEXT, data TEXT)` 是一个扩展函数，返回 `VOID`。

### 要求与注意事项

- 控制文件声明默认版本为 `1.0`。
- 请先安装并验证确认的扩展依赖项：`plpgsql`。
- 控制文件标记该扩展为不可重定位。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况。
