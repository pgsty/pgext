## 用法

来源：

- [已复核 commit 的 pg_py_plan_forwarding README](https://github.com/eduard-gr/pg_py_plan_forwarding/blob/74205a8f864d9cfbefaa7085887ce26debabe059/README.md)
- [已复核 commit 的 pg_py_plan_forwarding.control](https://github.com/eduard-gr/pg_py_plan_forwarding/blob/74205a8f864d9cfbefaa7085887ce26debabe059/pg_py_plan_forwarding.control)
- [版本 0.1 的空安装 SQL](https://github.com/eduard-gr/pg_py_plan_forwarding/blob/74205a8f864d9cfbefaa7085887ce26debabe059/pg_py_plan_forwarding--0.1.sql)
- [规划器与执行器钩子实现](https://github.com/eduard-gr/pg_py_plan_forwarding/blob/74205a8f864d9cfbefaa7085887ce26debabe059/pg_py_plan_forwarding.c)
- [Python 回调示例](https://github.com/eduard-gr/pg_py_plan_forwarding/blob/74205a8f864d9cfbefaa7085887ce26debabe059/pg_py_plan_forwarding.py)
- [构建配置](https://github.com/eduard-gr/pg_py_plan_forwarding/blob/74205a8f864d9cfbefaa7085887ce26debabe059/Makefile)

`pg_py_plan_forwarding` 是用于学习优化器的代码，会安装由嵌入式 Python 支持的 `planner`、`executor_start` 和 `executor_end` 钩子。它用 `nodeToString` 序列化 PostgreSQL 内部节点，把字符串传给 Python，再用 `stringToNode` 重建返回的计划。

版本 `0.1` 的安装 SQL 为空，因此创建扩展只会完成登记。在针对源码所用确切 Python 和 PostgreSQL 版本构建的一次性后端中，可以显式激活库，进行仅规划器实验：

```sql
CREATE EXTENSION pg_py_plan_forwarding;
LOAD 'pg_py_plan_forwarding';
EXPLAIN SELECT 1;
```

Python 回调示例会记录解析树和计划，并原样返回计划。

### 注意事项

- 构建过程硬编码 Python 3.9 头文件，初始化又把 `/home/vagrant/pg_py_plan_forwarding` 硬编码为模块路径。它不是可移植的 Python 环境集成。
- Python 回调可以返回格式错误或与版本不兼容的 PostgreSQL 内部节点文本。把它重建为计划可能导致后端崩溃或破坏执行语义。
- 已复核执行器钩子的 fallback 在没有更早钩子时不会调用标准执行器，而且 Python 引用管理存在不安全路径。修复并审计 C 代码前，不要在激活后运行普通语句。
- 这不是稳定 SQL API，也不是生产计划管理功能。PostgreSQL 内部计划表示和钩子签名会随版本变化。
