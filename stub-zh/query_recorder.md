## 用法

来源：

- [官方 README](https://github.com/tvondra/query_recorder/blob/14e6509a0b92b962111647b48aa3fed9e909c060/README)
- [官方扩展控制文件](https://github.com/tvondra/query_recorder/blob/14e6509a0b92b962111647b48aa3fed9e909c060/query_recorder.control)
- [官方 PGXN 分发页](https://pgxn.org/dist/query_recorder/)

`query_recorder` 是一个 PostgreSQL 查询捕获模块，可将执行过的 SQL 及其计时元数据写入轮转文件。它适合受控的工作负载分析或重放准备，而不应作为通用审计日志：它可能捕获敏感 SQL 文本，文件中的顺序也没有保证；本次复核的 `1.0.1` 还是一个年代较早的预览版本。

### 核心流程

该模块需要分配共享内存，因此必须加入 `shared_preload_libraries` 并重启 PostgreSQL。启用捕获前，应配置可写的基础文件名和有界轮转策略：

```ini
shared_preload_libraries = 'query_recorder'
query_recorder.filename = '/var/log/postgresql/recorded-query'
query_recorder.max_files = 20
query_recorder.size_limit = '256MB'
query_recorder.buffer_size = '8MB'
query_recorder.enabled = false
query_recorder.normalize = true
```

重启后安装扩展，并且只在需要观察的会话或时间段内开启记录：

```sql
CREATE EXTENSION query_recorder;

SET query_recorder.enabled = true;
SELECT current_database();
SET query_recorder.enabled = false;
```

输出文件采用 `.000` 之类的数字后缀。每条记录包含微秒时间戳、后端标识、耗时、查询长度和查询文本。启用 `query_recorder.normalize` 后，查询内的换行符会替换为空格；否则，消费端必须按记录中的长度解析 SQL，不能依赖行边界。

### 配置索引

- `query_recorder.filename` 设置输出基础路径。
- `query_recorder.max_files` 限制轮转文件数量。
- `query_recorder.size_limit` 设置单个文件的轮转阈值。
- `query_recorder.buffer_size` 控制共享输出缓冲区。
- `query_recorder.enabled` 启停捕获，并可通过 `SET` 修改。
- `query_recorder.normalize` 将查询中的换行符转换为空格。

### 运维边界

启用记录前，应准备目录所有权、严格的文件权限、保留策略、磁盘监控和安全销毁流程。记录流可能包含凭据、个人数据和超大语句，并发后端的记录也可能乱序。请使用代表性流量测试目标构建及解析器，并尽量缩短捕获窗口；`query_recorder` 不提供审计级完整性、身份归属或防篡改保证。
