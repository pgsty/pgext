## 用法

来源：

- [0.4.0 migration](https://github.com/univec-ai/postvec/blob/d972a6e7892d90363af677d916b91bef63e33a11/postvec/sql/postvec--0.3.0--0.4.0.sql)
- [postvec/README.md](https://github.com/univec-ai/postvec/blob/d972a6e7892d90363af677d916b91bef63e33a11/postvec/README.md)
- [postvec/postvec.control](https://github.com/univec-ai/postvec/blob/d972a6e7892d90363af677d916b91bef63e33a11/postvec/postvec.control)
- [postvec/Cargo.toml](https://github.com/univec-ai/postvec/blob/d972a6e7892d90363af677d916b91bef63e33a11/postvec/Cargo.toml)
- [README.md](https://github.com/univec-ai/postvec/blob/d972a6e7892d90363af677d916b91bef63e33a11/README.md)
- [postvec/src/api/search.rs](https://github.com/univec-ai/postvec/blob/d972a6e7892d90363af677d916b91bef63e33a11/postvec/src/api/search.rs)
- [postvec/src/api/registry.rs](https://github.com/univec-ai/postvec/blob/d972a6e7892d90363af677d916b91bef63e33a11/postvec/src/api/registry.rs)
- [postvec/src/api/status.rs](https://github.com/univec-ai/postvec/blob/d972a6e7892d90363af677d916b91bef63e33a11/postvec/src/api/status.rs)
- [postvec/sql/postvec--0.2.0--0.3.0.sql](https://github.com/univec-ai/postvec/blob/d972a6e7892d90363af677d916b91bef63e33a11/postvec/sql/postvec--0.2.0--0.3.0.sql)

`postvec` 0.4.0 异步维护嵌入向量列，并融合词法与语义搜索。所核验发布为 postvec-v0.4.0-1；软件包后缀与扩展版本分别计数。

### 启用与核心工作流

需要 PostgreSQL 16–18 和 `vector` 0.8 或更新版本。将 `postvec` 加入 `shared_preload_libraries`，配置目标数据库与推理后端后重启。发布的软件包包含内嵌与远程模式；源码构建若未启用内嵌特性，就必须使用远程推理。注册表前应确保所选模型可用。

```sql
CREATE EXTENSION postvec CASCADE;
SELECT postvec.refresh_models();
CREATE TABLE docs (id bigserial PRIMARY KEY, body text);
SELECT postvec.enable('public.docs', 'body',
  model => 'sentence-transformers-all-minilm-l6-v2', create_fts_index => true);
INSERT INTO docs(body) VALUES ('PostgreSQL extension development');
SELECT * FROM postvec.status();
```

提交插入事务，并等到 `postvec.status()` 报告 `pending_jobs = 0` 后再执行搜索：

```sql
SELECT d.body, s.rrf_score
FROM postvec.search('public.docs', 'body', 'database extensions') s
JOIN docs d ON d.id = s.pk_value::bigint ORDER BY s.rrf_score DESC;
```

### 对象与边界

`postvec.enable` 安装触发器、受管向量列与回填任务；处理完成后才会得到完整搜索结果。`postvec.adopt` 接管已有向量，`postvec.disable` 停止管理。`postvec.create_vector_index` 创建向量索引，自动建索引可能阻塞写入。`postvec.retry_dead` 重试符合条件的死信任务。

管理操作需要表所有者或超级用户权限。工作进程使用提升后的权限，应审查行安全限制。远程推理会将源文本发送到数据库主机之外。扩展使用 PostgreSQL 许可证，独立服务端和下载模型各有其许可条款。

### 0.4.0 升级

安装匹配的共享库与扩展文件并重启以载入预加载库，然后执行 `ALTER EXTENSION postvec UPDATE TO '0.4.0'`。从 0.3.0 到 0.4.0 的迁移获取扩展模式的咨询锁，没有声明额外模式变更。升级后检查工作进程健康状态、待处理任务及搜索结果。

本地推理使用 `postvec.mode = 'embedded'`，独立服务端使用 `postvec.mode = 'grpc'`。可选的托管提供方会接收输入，凭据需配置在推理主机上。跨越更旧版本时还需应用中间迁移：0.3.0 加固触发器 `search_path` 并规范化主键，可能使回填重新入队。有歧义的历史键会保留在 `jobs_dead`，直到问题解决。
