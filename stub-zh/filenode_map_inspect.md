## 用法

来源：

- [Official extension control file](https://github.com/ibarwick/filenode_map_inspect/blob/8cffe378ce38a6c3d7bff828636fff33e29e46f4/filenode_map_inspect.control)
- [Official upstream documentation](https://github.com/ibarwick/filenode_map_inspect/blob/8cffe378ce38a6c3d7bff828636fff33e29e46f4/README.md)

`filenode_map_inspect` 0.1 读取 PostgreSQL 的 `pg_filenode.map` 文件，用于集群级诊断。它能找出无法读取或不完整的映射文件，并列出映射的系统目录对象，但不会修复损坏，也不能替代常规恢复流程。

### 核心流程

```sql
CREATE EXTENSION filenode_map_inspect;

SELECT * FROM filenode_map_check();
SELECT * FROM filenode_map_list();
SELECT * FROM filenode_map_list_global();
```

- `filenode_map_check()` 定位数据库及全局映射文件，并报告每个文件的状态。
- `filenode_map_list()` 返回当前数据库映射中的关系名、OID 和 filenode。
- `filenode_map_list_global()` 返回全局映射中的对应条目。

### 运维边界

上游明确将其标为只面向教学与诊断的实验扩展。它会读取服务器端集群文件，因此只应由可信管理员安装并获得执行权限。错误报告只是进一步调查的证据；应保留数据目录并采用受支持的备份、还原或恢复流程，不要手工修改映射文件。
