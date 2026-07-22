## 用法

来源：

- [官方 README](https://github.com/jezdez/pg_linegazer/blob/37456fde2b5c0100a4f25af56814f4fe92b71065/README.md)
- [扩展 SQL](https://github.com/jezdez/pg_linegazer/blob/37456fde2b5c0100a4f25af56814f4fe92b71065/pg_linegazer--1.0.sql)
- [PL/pgSQL 插件实现](https://github.com/jezdez/pg_linegazer/blob/37456fde2b5c0100a4f25af56814f4fe92b71065/src/pg_linegazer.c)
- [预加载配置](https://github.com/jezdez/pg_linegazer/blob/37456fde2b5c0100a4f25af56814f4fe92b71065/conf.add)

`pg_linegazer` 是一个 alpha 质量的 PL/pgSQL 行覆盖率收集器。它无需重写函数即可记录语句命中，并包含嵌套调用；但统计数据只存在于当前后端进程中，既不共享也不持久化。

### 配置与流程

先把库加入 `shared_preload_libraries` 并重启 PostgreSQL，再创建和使用扩展：

```conf
shared_preload_libraries = 'pg_linegazer'
```

```sql
CREATE EXTENSION pg_linegazer;

SELECT test_func();
SELECT * FROM linegazer_simple_report('test_func'::regprocedure);
SELECT linegazer_simple_coverage('test_func'::regprocedure);

SELECT linegazer_clear();
```

应在同一个数据库会话中运行被检测的 PL/pgSQL 函数并请求报告。使用连接池时，如果执行与报告落在不同后端，报告可能为空或不完整。

### 函数

- `linegazer_simple_report(regprocedure)` 返回源码行号、命中数与源码文本。
- `linegazer_simple_coverage(regprocedure)` 返回命中数大于零的可执行行占比。
- `linegazer_clear()` 清除当前后端中的所有覆盖率数据。

命中数达到 9,999 后不再增长。扩展只测量行执行；分支覆盖、共享存储、持久化存储和 LCOV 输出仍是路线图事项，并未实现。

### 运维边界

预加载会在每个后端安装进程全局的 `PLpgSQL_plugin` 回调。实现会替换唯一的 PL/pgSQL rendezvous 指针，而不是组合多个插件，因此可能根据加载顺序与其他 PL/pgSQL 检测工具冲突。应测试开销；未经源码级兼容性审查，不要与其他插件组合使用。

上游把项目标为 alpha，且实现最后更新于 2016 年。虽然 README 声称支持 PostgreSQL 9.5+，代码依赖会随主版本变化的内部 PL/pgSQL 结构与 API。应针对实际 PostgreSQL 主版本构建和测试，只在开发或受控测试环境使用；除非相关工作负载和报告在同一后端运行，否则不要把报告当作持久的发布证据。
