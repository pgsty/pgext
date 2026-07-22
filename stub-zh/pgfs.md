## 用法

来源：

- [0.0.1 版本 Rust 实现](https://github.com/tavallaie/pgfs/blob/d11f5da0f136a1b3f684c632be425ad38171c754/src/lib.rs)
- [扩展控制文件](https://github.com/tavallaie/pgfs/blob/d11f5da0f136a1b3f684c632be425ad38171c754/pgfs.control)
- [Rust 包清单](https://github.com/tavallaie/pgfs/blob/d11f5da0f136a1b3f684c632be425ad38171c754/Cargo.toml)

`pgfs` 暴露一个服务器端文件系统操作：`pgfs_copy_dir(source_dir, dest_dir)`。它以 PostgreSQL 操作系统账号递归复制目录。这是高权限管理原语，而不是客户端文件传输工具。

### 复制受控目录

`0.0.1` 版本声明只有超级用户能安装扩展。安装后应撤销函数的默认执行权限；确有需要时，也只能授予严格受控的维护角色：

```sql
CREATE EXTENSION pgfs;
REVOKE EXECUTE ON FUNCTION pgfs_copy_dir(text, text) FROM PUBLIC;

SELECT pgfs_copy_dir(
  '/srv/postgresql/export/source',
  '/srv/postgresql/export/copy'
);
```

函数会创建缺失的目标目录、遍历子目录、复制普通文件，并在成功时返回 `true`。对于多种路径不存在、目录创建、读取或复制错误，它会返回 `false` 并输出信息消息。

### 安全与文件系统语义

两个路径都在数据库服务器上解释，并使用数据库服务账号权限。绝不能把函数授予不可信 SQL 角色，也不能直接接受用户提供的路径。应使用操作系统权限限制访问，并只使用经过审阅、位于 PostgreSQL 数据目录与配置目录之外的绝对路径。

同名目标文件可能被覆盖。尽管源码注释如此声称，实现并没有显式保留目录或文件权限。目录判断可能跟随符号链接，递归链接可能造成不安全遍历，非 UTF-8 路径可能触发 panic，后续条目失败时也不会回滚已完成的部分复制。函数不提供 fsync、原子发布、所有者保留、校验和、取消清理或审计记录。生产流程应使用操作系统备份/复制工具；如果必须使用这个已废弃扩展，应先用一次性数据演练符号链接、失败、磁盘耗尽、覆盖行为与恢复。
