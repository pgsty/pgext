## 用法

来源：

- [README.md](https://github.com/mejedi/planscape-postgresql-module/blob/54e7e6407c35a68741430bb50cb9872a863bd62e/README.md)
- [planscape.cpp](https://github.com/mejedi/planscape-postgresql-module/blob/54e7e6407c35a68741430bb50cb9872a863bd62e/planscape.cpp)
- [pg_hooks.cpp](https://github.com/mejedi/planscape-postgresql-module/blob/54e7e6407c35a68741430bb50cb9872a863bd62e/pg_hooks.cpp)
- [hook_engine.c](https://github.com/mejedi/planscape-postgresql-module/blob/54e7e6407c35a68741430bb50cb9872a863bd62e/hook_engine.c)
- [Makefile](https://github.com/mejedi/planscape-postgresql-module/blob/54e7e6407c35a68741430bb50cb9872a863bd62e/Makefile)
- [planscape.control](https://github.com/mejedi/planscape-postgresql-module/blob/54e7e6407c35a68741430bb50cb9872a863bd62e/planscape.control)

`planscape` 在 PostgreSQL 会话中加载，用于在执行 EXPLAIN 时记录优化器备选路径。control 文件标识版本 1.0，但 Makefile 不安装扩展 SQL；启用方式是 LOAD，而非 CREATE EXTENSION。

### 核心用法

```sql
LOAD 'planscape';
CREATE TEMP TABLE planscape_example (a integer);
EXPLAIN (PLANSCAPE) SELECT avg(a) FROM planscape_example;
```

### 运行边界

该模块通过运行时函数钩子截取优化器与输出函数，将规划数据写为 JSON。官方入口为 `EXPLAIN (PLANSCAPE)`。加载共享库需要具有相应权限的会话。底层钩子实现对架构及服务器构建敏感，仓库未声明当前 PostgreSQL 支持范围或许可证。应在隔离的开发环境使用，并在测试敏感查询前检查文件输出行为。
