## 用法

来源：

- [Official extension control file](https://github.com/KLIEBHAN/pg_plsql_graphs/blob/22b2eaef35a90234aa364aee8bf21f600c639a0e/pg_plsql_graphs.control)
- [Official upstream documentation](https://github.com/KLIEBHAN/pg_plsql_graphs/blob/22b2eaef35a90234aa364aee8bf21f600c639a0e/README.md)

`pg_plsql_graphs` 1.0 在 PL/pgSQL 函数运行时捕获控制流图和数据依赖图，通过视图暴露 DOT 文本，并允许客户端导出给 Graphviz。它是较早期的实验性插桩模块，不是扫描数据库中全部函数的静态分析器。

### 核心流程

文档中的构建方式把 iGraph 集成进 PostgreSQL 源码树。预加载模块，重启 PostgreSQL，再创建扩展。

```conf
shared_preload_libraries = 'pg_plsql_graphs'
```

```sql
CREATE EXTENSION pg_plsql_graphs;

SELECT my_plpgsql_function();
SELECT * FROM pg_plsql_graphs;

\COPY (SELECT * FROM pg_plsql_last_flowgraph_dot) TO 'flow.dot'
\COPY (SELECT * FROM pg_plsql_last_pdgs_dot) TO 'pdg.dot'
```

`pg_plsql_graphs` 显示已调用函数的缩进控制流图和依赖图；`pg_plsql_graphs_trimmed` 面向导出。两个 `pg_plsql_last_*` 视图暴露最近一次捕获的图。

### 注意事项

README 面向 PostgreSQL 9.3 时代源码和 iGraph 0.7.1，要求侵入式修改服务器构建并使用预加载钩子。捕获状态驻留内存且依赖函数实际执行。应把它作为隔离研究工具：测试服务器兼容性和钩子交互，避免处理敏感函数体，也不要把它当作持久审计数据。
