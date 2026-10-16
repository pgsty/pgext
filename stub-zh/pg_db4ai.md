## 用法

来源：

- [README.md](https://github.com/wotchin/pg_db4ai/blob/edea3ae743545d558cf27574dfb03e19f55a19aa/README.md)
- [sql/pg_db4ai--0.1.0.sql](https://github.com/wotchin/pg_db4ai/blob/edea3ae743545d558cf27574dfb03e19f55a19aa/sql/pg_db4ai--0.1.0.sql)
- [docs/ARCHITECTURE.md](https://github.com/wotchin/pg_db4ai/blob/edea3ae743545d558cf27574dfb03e19f55a19aa/docs/ARCHITECTURE.md)
- [docs/PORTING.md](https://github.com/wotchin/pg_db4ai/blob/edea3ae743545d558cf27574dfb03e19f55a19aa/docs/PORTING.md)
- [.github/workflows/ci.yml](https://github.com/wotchin/pg_db4ai/blob/edea3ae743545d558cf27574dfb03e19f55a19aa/.github/workflows/ci.yml)
- [LICENSE](https://github.com/wotchin/pg_db4ai/blob/edea3ae743545d558cf27574dfb03e19f55a19aa/LICENSE)
- [pg_db4ai.control](https://github.com/wotchin/pg_db4ai/blob/edea3ae743545d558cf27574dfb03e19f55a19aa/pg_db4ai.control)

`pg_db4ai` 0.1.0 在 PostgreSQL 后端进程内执行训练和预测，提供原生 C 算法，并可选用捆绑的 XGBoost。它是受 openGauss DB4AI 启发的早期实现，算法和模型格式有所不同，不能直接载入旧 openGauss 模型。上游面向 Linux 和 PostgreSQL 16–18。安装需要超级用户，对象固定在 `db4ai` 中，无需预加载或重启。

### 核心用法

```sql
CREATE EXTENSION pg_db4ai;
CREATE TABLE db4ai_example (target float8, feature float8);
INSERT INTO db4ai_example VALUES (2, 1), (4, 2), (6, 3);
SELECT db4ai.train('linear_demo', 'linear_regression',
    'SELECT target, feature FROM db4ai_example ORDER BY feature',
    '{"max_iterations":300,"learning_rate":0.01}'::jsonb);
SELECT db4ai.predict_float8('linear_demo', ARRAY[4.0]::float8[]);
SELECT db4ai.explain_model('linear_demo');
SELECT db4ai.export_model('linear_demo');
```

### 运行边界

训练查询先返回标签，再返回数值或布尔类型的标量特征；KMeans 和 PCA 只接收特征。NULL 和非有限数值会被拒绝，预测时必须保持训练时的特征顺序。`db4ai.predict_float8`、`db4ai.predict_bool`、`db4ai.predict_cluster` 与 `db4ai.transform` 用于相应模型类型。`db4ai.explain_model` 报告训练集指标，不代表独立验证集精度。`db4ai.export_model` 返回字节数据，当前不提供任意外部模型导入。

模型保存在写 WAL 的 `db4ai.pg_db4ai_models` 表中，不可原地修改，并按当前角色隔离。可在同一事务中先调用 `db4ai.drop_model` 再重新训练。读取按所有者过滤，不支持跨角色共享。训练遵循调用者的表权限与 RLS，只接受一条 SELECT，但不会隔离获准函数的副作用。仅受保护的模型登记写入会切换为扩展所有者身份。

`db4ai.work_mem` 默认在输入达到 16 MB 后使用临时文件；管理员限制包括 `db4ai.max_rows`（10000000）、`db4ai.max_model_mb`（64 MB）和 `db4ai.xgboost_max_input_mb`（128 MB）。最后一项仅限制稠密输入副本，不限制 XGBoost 的全部进程内存。训练持有事务与快照，取消采用协作方式。模型数据和标识序列纳入逻辑备份，恢复前应先创建所需角色。全部 SQL 函数均不支持并行执行；部署这一预览版本前应评估资源与权限。
