## 用法

来源：

- [extension/README.md](https://github.com/cpuchip/pg-ai-stewards/blob/c78212cf4dbca4a711cbc6416e06dad25d576b1d/extension/README.md)
- [extension/Cargo.toml](https://github.com/cpuchip/pg-ai-stewards/blob/c78212cf4dbca4a711cbc6416e06dad25d576b1d/extension/Cargo.toml)
- [extension/pg_ai_stewards.control](https://github.com/cpuchip/pg-ai-stewards/blob/c78212cf4dbca4a711cbc6416e06dad25d576b1d/extension/pg_ai_stewards.control)
- [extension/src/lib.rs](https://github.com/cpuchip/pg-ai-stewards/blob/c78212cf4dbca4a711cbc6416e06dad25d576b1d/extension/src/lib.rs)
- [extension/src/bgworker.rs](https://github.com/cpuchip/pg-ai-stewards/blob/c78212cf4dbca4a711cbc6416e06dad25d576b1d/extension/src/bgworker.rs)
- [extension/init/00-extensions.sql](https://github.com/cpuchip/pg-ai-stewards/blob/c78212cf4dbca4a711cbc6416e06dad25d576b1d/extension/init/00-extensions.sql)

`pg_ai_stewards` 0.3.0 在 PostgreSQL 18 中保存智能体状态、作业队列、模型配置及调度结果。Rust 原生核心依赖 `vector`，生成的安装 SQL 会嵌入所引用源码版本的迁移链。

### 启用与检查

添加预加载库时须保留其他已有项并重启 PostgreSQL。进程连接服务器环境变量 `POSTGRES_DB` 指定的数据库，默认为 `stewards`；应在该库创建扩展。

```conf
shared_preload_libraries = 'pg_ai_stewards'
```

```sql
CREATE EXTENSION vector;
CREATE EXTENSION pg_ai_stewards;
SELECT stewards.version();
SELECT * FROM stewards.providers_loaded();
SELECT stewards.enqueue('echo', 'echo', '{"hello":"world"}'::jsonb);
SELECT * FROM stewards.work_queue ORDER BY id DESC LIMIT 5;
```

### 运行边界

创建需要超级用户。预加载会注册调度进程；`STEWARDS_DISPATCHER_WORKERS` 默认为 4，范围限制在 1–16。服务提供方配置决定凭据、模型及请求目的地。echo 任务无需真实模型即可检查队列流程，实际推理仍需配置服务提供方。

`stewards` 模式保存队列、配置、流水线及历史。应限制凭据和调度／工具函数的访问，审核外部请求与工作负载成本。上游部署还会应用运行时迁移清单及可选附加配置，它们与核心扩展包分开。应按清单顺序执行，不能按文件名字母序重放 SQL。创建核心扩展不等于安装伴随扩展或用户界面。
