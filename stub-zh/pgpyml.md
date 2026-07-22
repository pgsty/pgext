## 用法

来源：

- [官方 README](https://github.com/Minoro/pgpyml/blob/70a119164c9af19d36a04d9592fd9e7f5a7ec521/README.md)
- [官方文档](https://minoro.is-a.dev/pgpyml/get-started/)
- [官方控制文件](https://github.com/Minoro/pgpyml/blob/70a119164c9af19d36a04d9592fd9e7f5a7ec521/pgpyml.control)

`pgpyml` 在 PostgreSQL 内加载 Python joblib/scikit-learn 模型，并提供预测函数与触发器辅助函数。上游明确将其标为尚在开发、未达到生产就绪。模型加载使用不可信 Python 与服务器本地文件，因此从数据库服务器安全角度看，模型制品就是可执行代码。

### 前置条件与核心流程

在 PostgreSQL 使用的 Python 环境中安装 `plpython3u` 及保存模型所需的软件包，再在专用模式中创建扩展：

```sql
CREATE SCHEMA IF NOT EXISTS pgpyml;
CREATE EXTENSION pgpyml SCHEMA pgpyml CASCADE;

SELECT *
FROM pgpyml.predict(
    '/srv/pgpyml/models/iris.joblib',
    ARRAY[[5.2, 3.5, 1.5, 0.2], [7.7, 2.8, 6.7, 2.0]]
);
```

模型路径在数据库服务器上解析，必须能由服务账号读取。嵌套数组每个特征向量对应一次预测，函数返回文本数组。

### 重要对象

- `predict` 把保存的模型应用于一个或多个特征数组。
- `predict_table_row` 先从表行读取指定列再预测。
- `classification_trigger` 在插入或更新时把预测写入目标列。
- `trigger_abort_if_prediction_is` 与 `trigger_abort_unless_prediction_is` 根据类别拒绝行。
- 组合触发器辅助函数既保存分类，也按条件拒绝操作。

触发器参数包含模型路径与列名。必须针对模式变更、空值、类型转换、模型特征顺序及预测错误测试生成的触发器。

### 安全与运维

`plpython3u` 不可信，而且 joblib/pickle 反序列化可以执行任意 Python 代码。只有严格受控的管理员才能安装扩展、写模型文件或调用带模型路径的 API。模型应存放在服务账号拥有的不可变目录，并在部署前验证摘要和来源。

推理同步运行在后端或行触发器中，会增加事务延迟与内存使用；异常可能中止写入。Python 环境、原生依赖、模型文件及精确 scikit-learn 版本不会被数据库转储捕获。应单独固定这些依赖、对并发做负载测试，并在任何非一次性用途前验证恢复与模型兼容性。
