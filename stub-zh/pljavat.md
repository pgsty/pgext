## 用法

来源：

- [gpdb/installation/pljavat.control](https://github.com/cloudberry-contrib/pljava/blob/d1a1c03507c1af6d8fce0c35a92169545e3a2a74/gpdb/installation/pljavat.control)
- [README.md](https://github.com/cloudberry-contrib/pljava/blob/d1a1c03507c1af6d8fce0c35a92169545e3a2a74/README.md)
- [gpdb/installation/pljavat--1.5.0.sql](https://github.com/cloudberry-contrib/pljava/blob/d1a1c03507c1af6d8fce0c35a92169545e3a2a74/gpdb/installation/pljavat--1.5.0.sql)
- [Makefile](https://github.com/cloudberry-contrib/pljava/blob/d1a1c03507c1af6d8fce0c35a92169545e3a2a74/Makefile)
- [COPYRIGHT](https://github.com/cloudberry-contrib/pljava/blob/d1a1c03507c1af6d8fce0c35a92169545e3a2a74/COPYRIGHT)

`pljavat` 是 Cloudberry 分支提供的仅可信 PL/Java 扩展。其最小 SQL 注册 Java 调用处理器和可信的 `java` 语言。控制版本仍为 1.5.0，与同仓库中版本更新的完整 PL/Java 控制文件不同。

### 核心用法

```sql
CREATE EXTENSION pljavat;
SELECT lanname, lanpltrusted FROM pg_language WHERE lanname = 'java';
```

### 运行边界

该条目面向 Cloudberry／Greenplum 分发，不声明原生 PostgreSQL 支持。应安装匹配的 PL/Java 原生库并配置 JVM，再由超级用户启用。语言被标记为可信不意味着普通用户可安装扩展。完整 PL/Java 扩展会创建重叠的处理器和语言对象，应选择合适的安装路径而非同时安装。此最小 SQL 不包含完整 SQL/J 仓库管理模式。
