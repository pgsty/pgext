## 用法

来源：

- [CHANGELOG.md](https://github.com/matroidbe/pg_extensions-releases/blob/ac623fe0885b79a2517eeefb9a1f4c92bddcc114/CHANGELOG.md)
- [extensions/pg_ml/pgbrew.toml](https://github.com/matroidbe/pg_extensions-releases/blob/ac623fe0885b79a2517eeefb9a1f4c92bddcc114/extensions/pg_ml/pgbrew.toml)
- [extensions/pg_ml/pg_ml.control](https://github.com/matroidbe/pg_extensions-releases/blob/ac623fe0885b79a2517eeefb9a1f4c92bddcc114/extensions/pg_ml/pg_ml.control)
- [README.md](https://github.com/matroidbe/pg_extensions-releases/blob/ac623fe0885b79a2517eeefb9a1f4c92bddcc114/README.md)
- [extensions/pg_ml/Cargo.toml](https://github.com/matroidbe/pg_extensions-releases/blob/ac623fe0885b79a2517eeefb9a1f4c92bddcc114/extensions/pg_ml/Cargo.toml)
- [extensions/pg_ml/src/lib.rs](https://github.com/matroidbe/pg_extensions-releases/blob/ac623fe0885b79a2517eeefb9a1f4c92bddcc114/extensions/pg_ml/src/lib.rs)
- [CHANGELOG.md](https://github.com/matroidbe/pg_extensions-releases/blob/ac623fe0885b79a2517eeefb9a1f4c92bddcc114/CHANGELOG.md)
- [LICENSE](https://github.com/matroidbe/pg_extensions-releases/blob/ac623fe0885b79a2517eeefb9a1f4c92bddcc114/LICENSE)
- [extensions/pg_ml/README.md](https://github.com/matroidbe/pg_extensions-releases/blob/ac623fe0885b79a2517eeefb9a1f4c92bddcc114/extensions/pg_ml/README.md)
- [extensions/pg_ml/pgbrew.toml](https://github.com/matroidbe/pg_extensions-releases/blob/ac623fe0885b79a2517eeefb9a1f4c92bddcc114/extensions/pg_ml/pgbrew.toml)

`pg_ml` 0.3.1 通过 `pgml` 管理 PyCaret 实验、模型训练与预测，并使用受管理的 Python 环境及后台任务。

### 核心用法

```ini
shared_preload_libraries = 'pg_ml'
```

```sql
CREATE EXTENSION pg_ml;
SELECT pgml.setup_venv();
SELECT * FROM pgml.load_dataset('iris');
SELECT * FROM pgml.setup('iris', 'species', project_name => 'iris_classifier');
SELECT * FROM pgml.create_model('iris_classifier', 'rf');
SELECT pgml.predict('iris_classifier', ARRAY[5.1,3.5,1.4,0.2]);
```

### 运行边界

控制文件要求超级用户安装。需要预加载并重启。Python 3.10+ 与 PyCaret 由上游环境初始化流程管理，该流程可能下载软件包。`setup`、`create_model` 与 `predict` 执行训练流程，`status` 和 `python_info` 查看运行时，0.3.0 增加按名称调用的概率预测。应限制模型／环境管理及外部数据访问。固定模式名称与其他机器学习项目存在重叠，不能假定可以同时安装。 这是采用 Matroid Source Available License 1.0 的无支持概念验证项目，API 可能变化。0.3.0 是新的升级起点：旧 0.2.0 安装需要预演数据迁移／重建，不能直接执行普通 ALTER EXTENSION UPDATE。遵循上游这一破坏性路径前，必须备份数据并检查依赖。

### 当前版本与升级

扩展版本 0.3.1 通过 `pg_ml.database` 选择后台工作进程使用的数据库。应在该库创建扩展；扩展尚不存在时，工作进程会等待，不再反复退出。调整需要重启的工作进程配置后须重启。 `pg_ml.training_database` 保留为已弃用的别名。 对扩展 0.3.0 及之后的版本，安装匹配文件后使用 ALTER EXTENSION UPDATE；更早版本仍需前述迁移。
