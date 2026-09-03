## 用法

来源：

- [官方文档](https://github.com/FranckPachot/pg_ipinfo/blob/9eb3e7163d3d5a87ef0645b94f6f553a2657ba43/README.md)
- [扩展控制文件](https://github.com/FranckPachot/pg_ipinfo/blob/9eb3e7163d3d5a87ef0645b94f6f553a2657ba43/ipinfo.control)
- [官方仓库](https://github.com/FranckPachot/pg_ipinfo)

`ipinfo` 通过 ipinfo.io HTTP 服务查询 IP 地理位置的函数扩展。

### 启用

为目标服务器安装文件，然后在目标数据库中创建 `ipinfo`：

```sql
CREATE EXTENSION ipinfo;
```

### 核心流程

以下示例取自经审查的上游文档。使用前应调整对象名、路径、凭据与工作负载参数。

```sql
SELECT getipinfo();
```

### 主要对象

经审查的安装接口包含以下主要对象：

| 对象 | 类型 | 运维作用 |
|---|---|---|
| `getipinfo` | FUNCTION | 经审查安装接口中的可调用函数。 |

### 运维与边界

- 目录生命周期为 abandoned；生产使用前应测试升级、备份恢复与服务器兼容性。
