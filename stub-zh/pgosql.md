## 用法

来源：

- [官方 README](https://github.com/pgoracle/pgora-osql/blob/abb37e20238dc80d29a5440b62d453625e7e766b/README.md)
- [版本 9.4 扩展 SQL](https://github.com/pgoracle/pgora-osql/blob/abb37e20238dc80d29a5440b62d453625e7e766b/sql/pgosql--9.4.sql)
- [官方 Makefile](https://github.com/pgoracle/pgora-osql/blob/abb37e20238dc80d29a5440b62d453625e7e766b/Makefile)

`pgosql` 是一个旧兼容性实验，README 将其描述为类 PL/SQL 过程语言。扩展 SQL 还会创建包含目录风格视图与兼容函数的 `sys` Schema。冻结的上游版本内部并不一致，因此更适合用于源码评估或迁移考古，而不是作为生产过程语言运行时。

### 预期安装与对象

control 文件要求超级用户安装。在已经修正上游库名不一致问题的构建中，预期入口如下：

```sql
CREATE EXTENSION pgosql;

SELECT sys.getdate(),
       sys.sysutcdatetime(),
       sys.isnull(NULL::integer, 42);

SELECT name, object_id, type_desc
FROM sys.objects
ORDER BY name;

SELECT lanname
FROM pg_language
WHERE lanname IN ('plosql', 'pltsql');
```

SQL 脚本定义 `sys.tables`、`sys.views` 与 `sys.objects`；时间辅助函数 `sys.sysdatetime()`、`sys.sysdatetimeoffset()`、`sys.sysutcdatetime()`、`sys.getdate()` 与 `sys.getutcdate()`；重载的 `sys.isnull` 函数；以及预期名为 `plosql` 的过程语言。

### 上游一致性警告

版本 `9.4` 无法提供可靠的未修改安装契约：

- Makefile 构建名为 `pgosql` 的共享模块，但扩展 SQL 从名为 `pgtsql` 的模块注册语言处理器。
- 扩展 SQL 创建 `plosql`，而仓库回归测试使用 `LANGUAGE pltsql`，并测试 `@` 变量与 `PRINT` 等 T-SQL 风格结构。
- README 将该语言称为 PL/SQL，但随附兼容对象与测试横跨不同 SQL 方言约定。

因此，除非下游软件包修补了这些名称或提供被引用模块，否则 `CREATE EXTENSION pgosql` 可能失败。不能假设带相同目录版本的软件包就实现了 Oracle PL/SQL 兼容性。

### 采用建议

先检查实际下游源码与软件包补丁，再测试扩展创建、处理器库解析、语言名、解析器行为、`sys` 对象语义、权限边界及转储恢复。只迁移回归测试覆盖的代码路径。该扩展没有声明预加载，但安装需要超级用户权限；未经专门安全审查，不应把已复核源码暴露给不可信过程代码。
