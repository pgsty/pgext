## 用法

来源：

- [官方上游 README](https://api.pgxn.org/src/narugn/narugn-0.3.0/README.md)
- [官方扩展控制文件 (narugn.control)](https://api.pgxn.org/src/narugn/narugn-0.3.0/narugn.control)
- [官方扩展 SQL (narugn--0.3.0.sql)](https://api.pgxn.org/src/narugn/narugn-0.3.0/sql/narugn--0.3.0.sql)

`narugn` — Narugn 是一个轻量级分布式计算机，由一个或多个本地连接的细胞组成。它需要带有 PL/Proxy 扩展的 PostgreSQL。使用它来进行相应的分析或存储工作流。在将其集成到应用程序 SQL 中之前，必须先安装并验证其扩展依赖项。

### 核心工作流

```sql
CREATE EXTENSION narugn;
```

在目标数据库中安装扩展，当可用时运行上游最小示例，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 重要对象

- `are_adjacent(s1 local_server , s2 local_server)` 是一个扩展函数，返回 `boolean`。
- `cell_logic(payload IN text[] , walked IN cdt[] , origin_tick IN bigint , rpfp IN boolean DEFAULT false , z OUT bigint , t OUT timestamp with time zone , output OUT text)` 是一个扩展函数，返回 `SETOF RECORD`。
- `cell_new_server(payload IN text[] , walked IN cdt[] , origin_tick IN bigint , rpfp IN boolean DEFAULT false , z OUT bigint , t OUT timestamp with time zone , output OUT text)` 是一个扩展函数，返回 `SETOF RECORD`。
- `cell_ping(payload IN text[] , walked IN cdt[] , origin_tick IN bigint , rpfp IN boolean DEFAULT false , z OUT bigint , t OUT timestamp with time zone , output OUT text)` 是一个扩展函数，返回 `SETOF RECORD`。
- `cell_rescan(payload IN text[] , walked IN cdt[] , origin_tick IN bigint , rpfp IN boolean DEFAULT false , z OUT bigint , t OUT timestamp with time zone , output OUT text)` 是一个扩展函数，返回 `SETOF RECORD`。
- `cell_version(payload IN text[] , walked IN cdt[] , origin_tick IN bigint , rpfp IN boolean DEFAULT false , z OUT bigint , t OUT timestamp with time zone , output OUT text)` 是一个扩展函数，返回 `SETOF RECORD`。
- `code_version()` 是一个扩展函数，返回 `text`。
- `configure_cell(cell IN cds , local_connstr IN text)` 是一个扩展函数，返回 `text`。
- `configure_cell(short_name IN text , full_name IN text , polygon IN polygon , connstr IN text , local_connstr IN text)` 是一个扩展函数，返回 `text`。
- `display_cct(this_cell cds , origin_cell cds , t bigint)` 是一个扩展函数，返回 `text`。
- `execute_sync(cell_function IN text , payload VARIADIC text[] DEFAULT '{}' , c OUT cds , z OUT bigint , dt OUT interval , output OUT text)` 是一个扩展函数，返回 `SETOF RECORD`。
- `execute_sync_abs(cell_function IN text , payload VARIADIC text[] DEFAULT '{}' , c OUT cds , z OUT bigint , t OUT timestamp with time zone , output OUT text)` 是一个扩展函数，返回 `SETOF RECORD`。
- `execute_sync_raw(cell_function IN text , payload IN text[] DEFAULT '{}' , ts IN timestamp with time zone DEFAULT clock_timestamp() , origin_cell IN cds DEFAULT NULL , origin_tick IN bigint DEFAULT NULL , walked IN cdt[] DEFAULT NULL , max_delay IN float DEFAULT 0.2 , c OUT cds…)` 是一个扩展函数，返回 `SETOF RECORD`。
- `global2dbname(i_c IN cds)` 是一个扩展函数，返回 `text`。

### 要求与注意事项

- 审核的控制文件声明默认版本为 `0.3.0`。
- 先安装并验证确认的扩展依赖项：`plproxy`。
- 在生产使用前，确认权限、支持的 PostgreSQL 版本、升级行为和失败情况。
