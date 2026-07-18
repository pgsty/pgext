## 用法

来源：

- [已复核 commit 的 README](https://github.com/sangli00/pg_checkdestroy/blob/f78520270a5a411d2604056a8c0dd2473556e505/README.md)
- [扩展 control 文件](https://github.com/sangli00/pg_checkdestroy/blob/f78520270a5a411d2604056a8c0dd2473556e505/pg_checkdestroy.control)
- [空扩展安装 SQL](https://github.com/sangli00/pg_checkdestroy/blob/f78520270a5a411d2604056a8c0dd2473556e505/pg_checkdestroy--1.0.sql)
- [hook 实现](https://github.com/sangli00/pg_checkdestroy/blob/f78520270a5a411d2604056a8c0dd2473556e505/pg_checkdestroy.c)

`pg_checkdestroy` 是一个实验性 hook 模块，尝试压制表上的 `DROP`、`TRUNCATE` 和无条件 `DELETE` 语句。只有通过 `shared_preload_libraries` 加载时，它才会安装 parser、executor 与 utility hook；扩展 SQL 本身为空。

### 配置

```conf
shared_preload_libraries = 'pg_checkdestroy'
```

重启 PostgreSQL 后注册扩展，并确认其会话开关已启用：

```sql
CREATE EXTENSION pg_checkdestroy;
SHOW pg_checkdestroy.work;
```

`pg_checkdestroy.work` 默认为 `on`，但它是 `PGC_USERSET` 参数，任何有足够权限连接的会话都可以禁用它。

### 注意事项

- 上游明确表示代码未经完整测试，不建议在线上/生产环境使用。源码停留在 2018 年，并依赖旧版 PostgreSQL 内部 hook API。
- 该 hook 通过跳过正常执行来压制选定命令，而不是抛出明确的策略错误。因此，客户端可能收到误导性的完成行为。
- 这不是安全边界：`pg_checkdestroy.work` 可由用户设置，覆盖范围取决于特定 parse/plan 形态，特权用户也能绕过或卸载模块。
- 实现会串联多个全局 hook，并假定按查询维护的状态；这可能与错误、嵌套语句、预备计划或其他 hook 扩展产生不良交互。应使用原生权限、所有权控制和经过测试的审计机制，而不要依靠该模块保护数据。
