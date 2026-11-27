## 用法

来源：

- [Official source](https://github.com/matroidbe/pg_extensions-releases/releases/tag/pg_augur-v0.4.2)
- [Official source](https://github.com/matroidbe/pg_extensions-releases/releases/download/pg_augur-v0.4.2/pg_augur-0.4.2-pg18-darwin-arm64.tar.gz)
- [Official source](https://github.com/matroidbe/pg_extensions-releases/releases/download/pg_augur-v0.4.2/SHA256SUMS)
- [Official source](https://github.com/matroidbe/pg_extensions-releases/releases/download/pg_augur-v0.4.2/LICENSE-pg_augur)

`pg_augur` 0.4.2 将 Augur 训练、模型存储与推理引入 PostgreSQL 18。本版本发布包含控制文件和 SQL 的二进制包，未在该发布处公开实现源码。它使用带有内部使用限制的 Matroid Source Available License v1.0，并非不受此限制的开源许可证。

### 启用与检查

安装与服务器及平台匹配的官方包。已发布平台为 Linux amd64 和 macOS arm64。Linux 需要 glibc 2.36 或更高版本及 libgomp1；macOS 构建要求 macOS 14 或更高版本，XGBoost 使用单线程训练。添加动态库时保留其他预加载项，然后重启：

```ini
shared_preload_libraries = 'pg_augur'
```

```sql
CREATE EXTENSION pg_augur;
SELECT pgaugur.extension_docs();
```

创建需要超级用户。扩展不受信任、不可重定位，使用固定的 `pgaugur` 模式；后台训练工作进程要求预加载。

### 训练与使用模型

对于已有且包含目标列的训练表，提交任务并保存返回的任务标识：

```sql
SELECT pgaugur.start_training(
    project_name => 'demo',
    source_table => 'public.training_rows',
    target_column => 'target',
    automl => true);
```

提交事务后，工作进程才能看到任务。将返回的 ID 传给 `pgaugur.training_status`，检查状态、进度、模型 ID 及错误；`pgaugur.cancel_training` 请求取消。使用 `pgaugur.deploy` 选择模型前，应检查指标与输入兼容性。`pgaugur.predict_row` 接收按名称组织的 JSONB 特征，`pgaugur.predict` 接收按位置排列的双精度数组，`pgaugur.inference_schema` 提供模型要求的输入结构。

### 模型状态与权限

项目、模型、训练任务、实验及运行产物存储在 `pgaugur` 表中，应整体备份并验证恢复。将模式和函数访问限制给预期角色，检查训练关系的权限；安装需要超级用户，并不等于所有 SQL 函数自动私有。训练可能消耗大量 CPU 与内存。本发布提供版本化升级脚本，但缺少可公开审计的实现源码；采用前应在可丢弃数据库中测试模型、升级和恢复。

### 从 0.4.0 升级

安装匹配的 0.4.2 二进制文件，重启预加载该库的服务器，再执行 `ALTER EXTENSION pg_augur UPDATE`。中间迁移增加项目漂移容差，为 `pgaugur.import_model` 和 `pgaugur.check_representative` 增加末尾可选 `tolerances` 参数，并新增 `pgaugur.set_drift_tolerances`。已有调用保持参数行为。0.4.1 到 0.4.2 的脚本不改变 SQL 对象，模型选择和 SHAP 行为变化位于二进制库中。
