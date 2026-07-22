## 用法

来源：

- [Monorepo README](https://github.com/ChristoferBerruz/hyper_db/blob/07c3b41e90ef434f29b3a97ac701bfce62e73363/README.md)
- [扩展实现](https://github.com/ChristoferBerruz/hyper_db/blob/07c3b41e90ef434f29b3a97ac701bfce62e73363/optim/src/lib.rs)
- [扩展控制文件](https://github.com/ChristoferBerruz/hyper_db/blob/07c3b41e90ef434f29b3a97ac701bfce62e73363/optim/optim.control)

`optim` 是一个来自数据库课程项目的早期 pgrx 概念验证。它公开了若干与优化有关的 PostgreSQL 类型和一个演示函数；目前并没有实现优化器，也没有形成有文档的超参数搜索流程。

### 核心流程

控制文件要求由超级用户安装这个不受信任的 C 兼容模块。安装后，唯一有文档依据的可调用操作是自动生成的 `hello_optim()` 函数：

```sql
CREATE EXTENSION optim;
SELECT hello_optim();
```

预期结果为 `Hello, optim`。这只是 pgrx 模块的冒烟测试，并不是优化结果。

### 源码定义的对象

- `OptimizationDirection`：pgrx `PostgresEnum`，包含 `Maximize` 和 `Minimize` 两个变体。
- `Hyperparameter`：pgrx `PostgresType`，包含一个 `f64` 值。
- `Trial`：pgrx `PostgresType`，包含一组 `Hyperparameter` 值和一个 `f64` 分数。
- `hello_optim()`：返回常量文本 `Hello, optim`。

使用前应在实际安装的构建中检查 pgrx 生成类型的准确 SQL 拼写与输入输出语法；上游没有给出这些类型的 SQL 示例。

### 运维说明

crate 版本为 `0.0.0`，控制文件版本由 Cargo 版本生成，README 将该仓库描述为学生数据库项目。控制文件声明 `relocatable = false`、`superuser = true` 和 `trusted = false`；没有定义预加载设置。上游没有记录持久化模型、规划器钩子流程、发布约定或兼容性保证。在上游定义稳定 SQL API 前，只应将 `optim` 用于源码评估或隔离实验。
