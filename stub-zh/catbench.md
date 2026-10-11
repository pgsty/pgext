## 用法

来源：

- [README.md](https://github.com/joelonsql/pg-catbench/blob/d81740c55a3aad7c599c502d044ba8db6f6e0837/README.md)
- [Makefile](https://github.com/joelonsql/pg-catbench/blob/d81740c55a3aad7c599c502d044ba8db6f6e0837/Makefile)
- [functions/new_benchmark.sql](https://github.com/joelonsql/pg-catbench/blob/d81740c55a3aad7c599c502d044ba8db6f6e0837/functions/new_benchmark.sql)
- [views/vreport.sql](https://github.com/joelonsql/pg-catbench/blob/d81740c55a3aad7c599c502d044ba8db6f6e0837/views/vreport.sql)
- [catbench.control](https://github.com/joelonsql/pg-catbench/blob/d81740c55a3aad7c599c502d044ba8db6f6e0837/catbench.control)

`catbench` 保存 PostgreSQL 系统目录函数的基准定义与结果。当前可以手工运行基准，自动跟踪提交仍是项目目标。

### 核心用法

```sql
CREATE EXTENSION timeit;
CREATE EXTENSION catbench;
SELECT * FROM catbench.vreport;
```

### 运行边界

先安装 `timeit`。`catbench.new_benchmark()` 注册基准定义，`catbench.vreport` 提供报告；上游配套可执行程序负责采集与执行。安装 SQL 由 Makefile 拼接。尽管控制文件声明模块路径，此扩展的 Makefile 只安装 SQL，没有定义原生模块。无需预加载，不可迁移模式；控制文件将 superuser 设为 false，但未声明扩展可信。上游没有说明 PostgreSQL 主版本支持范围。
