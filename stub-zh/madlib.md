## 用法

来源：

- [Apache MADlib 项目](https://madlib.apache.org/)
- [MADlib 2.1.0 用户指南](https://madlib.apache.org/docs/latest/)
- [官方模块索引](https://madlib.apache.org/docs/latest/modules.html)

`madlib` 是 Apache MADlib 2.1.0，这是一个为 PostgreSQL 和 Greenplum 提供的数据库内分析库。它提供了 SQL 调用的统计、图、矩阵、机器学习、XGBoost 和深度学习工作流，同时保持训练数据在数据库中。

### 核心工作流

MADlib 是通过其 `madpack` 工具部署的，而不是 `CREATE EXTENSION`。安装支持目标数据库的版本后，在与目标数据库相关的连接中使用：

```sh
madpack -p postgres -c analyst@localhost:5432/analytics install
```

在开始算法工作流之前，请验证已安装的模式：

```sql
SELECT madlib.version();
```

从官方索引中选择一个模块，创建该模块的文档输入表，运行其训练或分析函数，并检查生成的模型或结果表。

### 主要模块族

- 回归、分类、聚类、抽样和假设检验
- 图算法和路径分析
- 数组、矩阵、分解和稀疏向量
- 模型准备、深度学习和 XGBoost
- 数据转换和实用函数

### 要求与注意事项

- MADlib 2.1.0 是经过审查的版本。请使用官方支持的数据库和操作系统矩阵来获取确切的构建。
- `madpack` 创建了一个包含大量函数和支持对象的重要模式，并拥有安装、升级和卸载行为；不要替换 `ALTER EXTENSION UPDATE`。
- 单个模块有不同的 Python 和本地库依赖项。在启用模块之前，请验证这些依赖项和资源需求。
- 训练函数通常会创建模型和摘要表。请使用专用模式，审查生成对象的名称，并测试回滚和清理行为。
