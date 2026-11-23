## 用法

来源：

- [README.md](https://github.com/angelmorgaldev/pg-quick-sync/blob/45c160cec4f72673ec1a34fb1a93cb832d17d18d/README.md)
- [quick_sync.control](https://github.com/angelmorgaldev/pg-quick-sync/blob/45c160cec4f72673ec1a34fb1a93cb832d17d18d/quick_sync.control)
- [sql/quick_sync--0.1.0.sql](https://github.com/angelmorgaldev/pg-quick-sync/blob/45c160cec4f72673ec1a34fb1a93cb832d17d18d/sql/quick_sync--0.1.0.sql)
- [LICENSE](https://github.com/angelmorgaldev/pg-quick-sync/blob/45c160cec4f72673ec1a34fb1a93cb832d17d18d/LICENSE)

`quick_sync` 0.1.0 通过 PostgreSQL URI 管理 postgres_fdw 连接对象。项目名为 pg-quick-sync，SQL 扩展名则是 quick_sync。已实现的模块覆盖连接创建、列出、修改与删除；表映射和增量同步仍在开发中。上游要求 PostgreSQL 13 及以上版本。

### 管理连接

使用足够的管理权限，先安装依赖，再安装本扩展：

```sql
CREATE EXTENSION postgres_fdw;
CREATE EXTENSION quick_sync;
SELECT quick_sync.create_connection(
    'source', 'postgresql://reader:example@db.example.net:5432/app'
);
SELECT * FROM quick_sync.list_connections();
SELECT quick_sync.edit_connection('source', p_fetch_size => 10000);
```

上面的 URI 仅作示例；实际连接信息应通过客户端参数绑定，并保护 SQL 日志和历史记录。`quick_sync.create_connection()` 按规范化后的别名创建外部服务器，为当前数据库用户创建映射，并在 `quick_sync.connections` 中记录不含密码的元数据。默认抓取数量为 50000，且启用远端估算。创建操作只配置目录对象，不证明远端服务器能够连通。

### 接口与权限

`quick_sync.list_connections()` 返回登记的元数据、外部表数量以及外部服务器是否仍然存在。其活动标志只检查目录对象存在性，不检查网络健康。`quick_sync.edit_connection()` 接受新 URI，也可分别指定主机、端口、数据库、用户、密码与抓取数量。

函数使用调用者权限。调用者需要相应的外部服务器、用户映射权限，以及扩展元数据表的访问权限。密码保存在 PostgreSQL 用户映射中，不进入自定义清单表。这里的简单 URI 解析器并不是完整的 libpq URI 实现，不应假定支持 IPv6 字面量、百分号解码或任意连接选项。

### 谨慎删除

```sql
SELECT quick_sync.delete_connection('source', false);
```

第二个参数默认为 true，会用 CASCADE 删除服务器，可能连带删除依赖的外部表。传入 false 使用 RESTRICT，存在依赖时会阻止删除。扩展使用固定的 `quick_sync` 模式，不可迁移模式，无需共享预加载或重启。该 SQL 模块没有安装后台同步进程。
