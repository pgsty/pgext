## 用法

来源：

- [README.md](https://github.com/bugraaktug/pgquarry/blob/2690fdb537bb325f84e4c2d3293738b05780e88b/README.md)
- [pgquarry.control](https://github.com/bugraaktug/pgquarry/blob/2690fdb537bb325f84e4c2d3293738b05780e88b/pgquarry.control)
- [sql/pgquarry--1.0.sql](https://github.com/bugraaktug/pgquarry/blob/2690fdb537bb325f84e4c2d3293738b05780e88b/sql/pgquarry--1.0.sql)

`pgquarry` 1.0 提供本地嵌入与文本生成所用的 SQL 队列和接口。它依赖 `vector`，并需要独立的 `pgquarry_worker` 操作系统进程加载本地 GGUF 模型。仅安装 SQL 扩展不会执行推理。

### 基本用法

```sql
CREATE EXTENSION pgquarry CASCADE;
SELECT pgquarry.embed_async('A document to embed');
SELECT id, status FROM pgquarry.jobs ORDER BY id DESC LIMIT 5;
```

在配置文件中设置进程连接的数据库和模型路径，再启动推理进程。无需 PostgreSQL 共享预加载或服务器重启。

### 队列、触发器与权限

`pgquarry.watch` 安装插入／更新触发器，将嵌入任务入队并把结果写回源表或另一张目标表。`pgquarry.watch_generate` 增加生成任务；`pgquarry.generate_async` 直接提交文本。同步过程需要传入文档要求的输出参数及超时值，仍依赖正在运行的进程。生成模型须单独配置。

创建扩展需要超级用户。脚本向 `PUBLIC` 授予模式、表和函数的访问权限，在共享数据库中使用前应审核这一权限边界。注册监听需拥有相关表。删除源行不会取消排队任务，跨表写回可能产生孤立结果。进程的保留策略会清理已完成任务，因此需规划审计保留时间与模型维度。
