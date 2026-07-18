## 用法

来源：

- [上游 README](https://github.com/AbdulYadi/pgreadfile/blob/280468c96e641b47591c581e2729864df8aa0de1/README.md)
- [1.0 版安装 SQL](https://github.com/AbdulYadi/pgreadfile/blob/280468c96e641b47591c581e2729864df8aa0de1/pgreadfile--1.0.sql)
- [文件 I/O 实现](https://github.com/AbdulYadi/pgreadfile/blob/280468c96e641b47591c581e2729864df8aa0de1/pgreadfile.c)

pgreadfile 1.0 通过 pgreadfile(text) 与 pgwritefile(text, bytea) 暴露不受限制的服务器端文件读写。已发布扩展不适合多用户或生产数据库。

### 安装时撤销访问权

```sql
BEGIN;
CREATE EXTENSION pgreadfile;
REVOKE EXECUTE ON FUNCTION pgreadfile(text) FROM PUBLIC;
REVOKE EXECUTE ON FUNCTION pgwritefile(text, bytea) FROM PUBLIC;
COMMIT;
```

不要向应用角色授予写函数。如果经过审查的管理流程必须读取固定批准文件，只能向严格控制的角色授予读函数，并在该函数之外验证路径。

### 注意事项

- 安装 SQL 让 public 可以执行两个函数。否则任何数据库用户都能读取 PostgreSQL 操作系统账户可见的所有服务器文件，并覆盖或截断所有可写文件。
- 两个函数都被错误声明为不可变，尽管它们依赖并修改外部状态。规划器常量折叠或表达式复用可能在意外时机执行函数；写函数的声明尤其危险。
- 读函数会将整个文件载入后端内存，且没有大小上限。大文件可能耗尽内存。
- 文件打开失败与部分写入失败无法可靠报告，短写也会被忽略。符号链接与路径竞态没有得到控制。
- 上游只记录了针对 PostgreSQL 11 的编译，没有测试、权限模型或当前兼容证据。应优先使用带预定义角色与显式路径控制的受维护核心设施。
