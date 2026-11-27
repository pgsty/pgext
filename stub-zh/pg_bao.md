## 用法

来源：

- [Official source](https://github.com/LearnedSystems/BaoForPostgreSQL/blob/1c15600a6d0a5bf09b08405702889eb1783579ca/README.md)
- [Official source](https://github.com/LearnedSystems/BaoForPostgreSQL/blob/1c15600a6d0a5bf09b08405702889eb1783579ca/pg_extension/pg_bao.control)
- [Official source](https://github.com/LearnedSystems/BaoForPostgreSQL/blob/1c15600a6d0a5bf09b08405702889eb1783579ca/pg_extension/pg_bao--0.0.1.sql)
- [Official source](https://github.com/LearnedSystems/BaoForPostgreSQL/blob/1c15600a6d0a5bf09b08405702889eb1783579ca/pg_extension/main.c)
- [Official source](https://github.com/LearnedSystems/BaoForPostgreSQL/blob/1c15600a6d0a5bf09b08405702889eb1783579ca/docs/src/tutorial/1_pg_setup.md)
- [Official source](https://github.com/LearnedSystems/BaoForPostgreSQL/blob/1c15600a6d0a5bf09b08405702889eb1783579ca/docs/src/tutorial/2_bao_setup.md)
- [Official source](https://github.com/LearnedSystems/BaoForPostgreSQL/blob/1c15600a6d0a5bf09b08405702889eb1783579ca/LICENSE)

`pg_bao` 是 Bao 学习型查询优化器原型，控制文件版本为 0.0.1，配套 SQL 脚本为空。官方教程演示 PostgreSQL 12.3，该历史示例不代表已支持当前 PostgreSQL 大版本。

### 启用钩子

在研究集群安装匹配的库后，保留已有预加载条目、加入模块并重启 PostgreSQL。使用前启动独立的 Bao 模型服务器：

```ini
shared_preload_libraries = 'pg_bao'
```

```sql
SHOW enable_bao;
SET bao_host = 'localhost';
SET bao_port = 9381;
SET enable_bao = on;
EXPLAIN SELECT count(*) FROM existing_table;
SET enable_bao = off;
```

将示例表替换为已有关系。启用依靠预加载和会话设置，不依赖数据库 SQL 对象。`enable_bao` 默认关闭；启用后，`enable_bao_selection` 和 `enable_bao_rewards` 分别控制执行计划选择和延迟反馈采集。模型由独立 Bao 服务器维护和重新训练。

### 观察行为

启用时，`EXPLAIN` 包含 Bao 预测与推荐提示。尚未训练的服务器可能没有预测。训练需要有代表性的已完成工作负载和足够的观测量，应将规划加执行的总耗时与普通优化器比较。实现会顺序尝试多种提示，增加规划成本。

### 安全与维护

模型服务器不提供认证或加密，应限制在可信网络和预期客户端范围内。虽然执行计划表示省略谓词字面量，关系名、计划结构和执行测量仍会离开 PostgreSQL。上游明确将其视为尚未经过生产安全审计的原型。

训练和模型状态需在 PostgreSQL 备份之外另行保留。移除预加载需要重启服务器；关闭会话设置可以立即停止该会话使用优化器路径。项目没有 SQL 迁移链，也没有经过验证的现代大版本兼容矩阵。实现采用 AGPLv3 许可证。
