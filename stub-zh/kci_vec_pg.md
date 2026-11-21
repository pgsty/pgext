## 用法

来源：

- [pg/README.md](https://github.com/KimYuki3/kci-vec-src/blob/4e03ffed84429defc1aa02629e93f1a7bfd08fca/pg/README.md)
- [pg/kci_vec_pg.control](https://github.com/KimYuki3/kci-vec-src/blob/4e03ffed84429defc1aa02629e93f1a7bfd08fca/pg/kci_vec_pg.control)
- [pg/kci_vec_pg--0.3.sql](https://github.com/KimYuki3/kci-vec-src/blob/4e03ffed84429defc1aa02629e93f1a7bfd08fca/pg/kci_vec_pg--0.3.sql)
- [LICENSE](https://github.com/KimYuki3/kci-vec-src/blob/4e03ffed84429defc1aa02629e93f1a7bfd08fca/LICENSE)

`kci_vec_pg` 通过 SQL 函数提供 KCI-Vec 近似最近邻检索。SQL 版本 0.3 是在 Linux x86-64 的 PostgreSQL 15–17 上测试的原型。源码采用 PolyForm Noncommercial 1.0.0 许可证。向量索引镜像保存在外部文件中，备份与备库运行需要单独处理。

### 构建与检索

以超级用户创建扩展，然后从 ID 可转为 bigint、向量可转为实数数组的表构建索引：

```sql
CREATE EXTENSION kci_vec_pg;
CREATE TABLE docs (id bigint PRIMARY KEY, embedding real[]);
INSERT INTO docs
SELECT i, ARRAY(SELECT random()::real FROM generate_series(1, 128))
FROM generate_series(1, 1000) AS i;
SELECT kci_vec_build('docs_index', 'docs', 'id', 'embedding', nlist => 16);
SELECT kci_vec_track('docs_index');
SELECT * FROM kci_vec_search(
  'docs_index', (SELECT embedding FROM docs WHERE id = 1), 10);
SELECT * FROM kci_vec_status();
```

构建支持余弦、平方欧氏距离和内积。检索返回 ID 与分数；余弦和内积分数越大越相似，平方距离越小越相似。提高探测与重排参数会增加工作量以换取召回率。另一个重载接受允许返回的 ID 数组用于过滤。规划器不会通过 PostgreSQL 索引访问方法自动选择这种检索。

### 变更与维护

`kci_vec_track()` 将插入、更新、删除和截断记录到普通表中；查询会合并待处理变更，包括调用者事务中的变更。`kci_vec_refresh()` 将变更写入镜像，积压过多会降低查询速度。`kci_vec_untrack()` 停止跟踪，`kci_vec_drop()` 删除索引。

可选的预加载可降低变更跟踪成本，并启用后台进程：

```ini
shared_preload_libraries = 'kci_vec_pg'
kci_vec_pg.databases = 'mydb'
kci_vec_pg.refresh_pending = 10000
```

修改预加载配置需要重启，且应保留原有库。普通函数接口不要求此配置。独立的 `vector` 扩展是可选项；需要向量类型重载时，应先安装它，再安装本扩展。

### 恢复与限制

`$PGDATA/kci_vec/` 下的镜像不进入 WAL 或 `pg_dump`，扩展也不对其加密。恢复后的数据库或备库必须通过 `kci_vec_materialize()` 重建本地镜像；镜像缺失或过期时，查询会拒绝执行。`kci_vec_image_here()` 检查镜像是否在本机存在。构建和刷新只能在主库执行。

删除扩展不会清理镜像，应先调用 `kci_vec_drop()`。文档要求升级时安装新文件、重启 PostgreSQL，再执行 `ALTER EXTENSION kci_vec_pg UPDATE`。NULL ID 或向量会被跳过。过滤条件选择性很高时，可能需要增加探测量或改用精确 SQL 检索。原型阶段的接口和目录表仍可能变化。
