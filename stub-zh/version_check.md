## 用法

来源：

- [version_check.control](https://github.com/VladlenPopolitov/version_control/blob/844aa0cd66391a9848f1242c1287a243d9d44baa/version_check.control)
- [README.md](https://github.com/VladlenPopolitov/version_control/blob/844aa0cd66391a9848f1242c1287a243d9d44baa/README.md)
- [version_check.c](https://github.com/VladlenPopolitov/version_control/blob/844aa0cd66391a9848f1242c1287a243d9d44baa/version_check.c)
- [version_check--5.0.sql](https://github.com/VladlenPopolitov/version_control/blob/844aa0cd66391a9848f1242c1287a243d9d44baa/version_check--5.0.sql)
- [Makefile](https://github.com/VladlenPopolitov/version_control/blob/844aa0cd66391a9848f1242c1287a243d9d44baa/Makefile)

`version_check` 提供 PostgreSQL 主版本布尔谓词。所核对的 5.0 SQL 定义了版本 13 至 17 的谓词，而当前 C 源码明确限定只能针对 PostgreSQL 17 编译。

### 核心用法

```sql
CREATE EXTENSION version_check;
SELECT getversion13(), getversion14(), getversion15(), getversion16(), getversion17();
```

### 运行边界

由超级用户安装；当前实现不要求预加载或重启。在 PG17 构建中，仅版本 17 的谓词返回真。这些函数反映编译目标，不能用作通用运行时兼容性探测。README 中仍有旧文件名与注释；此版本以控制文件、SQL 和编译保护条件为准。未找到明确许可证。
