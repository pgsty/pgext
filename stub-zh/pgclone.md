
## 用法

> 语法：
>
> ```sql
> CREATE EXTENSION pgclone;
> SELECT pgclone_table('host=source-server dbname=mydb user=postgres password=secret',
>                      'public', 'customers', true);
> ```
>
> 来源：[README](https://github.com/valehdba/pgclone)，[使用指南](https://github.com/valehdba/pgclone/blob/main/docs/USAGE.md)

`pgclone` 可以直接从 SQL 克隆 PostgreSQL 数据库、模式、表、函数、角色和权限。上游 README 强调它使用 PostgreSQL 的 COPY 协议，避免了外部 `pg_dump` / `pg_restore` 工作流。

## 核心能力

README 列出的支持包括：

- 克隆表、模式、函数和完整数据库
- 索引、约束、触发器、视图、物化视图和序列
- 通过列选择和 `WHERE` 过滤进行选择性克隆
- 使用 error、skip、replace 或 rename 策略处理冲突
- 数据脱敏与敏感列发现
- 使用后台 worker 的异步与并行克隆

## 基本函数

```sql
SELECT pgclone_table(
    'host=source-server dbname=mydb user=postgres password=secret',
    'public', 'customers', true
);

SELECT pgclone_schema(
    'host=source-server dbname=mydb user=postgres password=secret',
    'sales', true
);

SELECT pgclone_database(
    'host=source-server dbname=mydb user=postgres password=secret',
    true
);
```

README 还记录了 `pgclone_version()`，用于安装后做版本确认。

## 异步模式

如果要使用后台 worker 功能，扩展必须预加载：

```ini
shared_preload_libraries = 'pgclone'
```

上游文档把异步操作、进度跟踪和多 worker 并行克隆拆分在独立文档中说明。

## 需求

当前上游要求如下：

- PostgreSQL 14 或更高版本
- PostgreSQL 开发头文件
- `libpq` 开发库
- GCC 或兼容的 C 编译器

项目主页是 [postgresql.az](https://postgresql.az/)，README 还链接了使用指南、异步操作、测试和架构等独立文档。
