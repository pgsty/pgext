## 用法

来源：

- [官方文档](https://github.com/fabriziomello/unique_id/blob/95ac200ada23072cb0f4ad05741fd92c24c0470b/README.md)
- [扩展控制文件](https://github.com/fabriziomello/unique_id/blob/95ac200ada23072cb0f4ad05741fd92c24c0470b/unique_id.control)
- [官方仓库](https://github.com/fabriziomello/unique_id)

`unique_id` 受 Instagram 与 Sonyflake 启发的时间有序唯一标识符生成器。

### 启用

为目标服务器安装文件，然后在目标数据库中创建 `unique_id`：

```sql
CREATE EXTENSION unique_id;
```

### 核心流程

以下示例取自经审查的上游文档。使用前应调整对象名、路径、凭据与工作负载参数。

```sql
fabrizio=# CREATE EXTENSION unique_id;
CREATE EXTENSION
fabrizio=# \dx unique_id
                                     List of installed extensions
   Name    | Version | Schema |                              Description
-----------+---------+--------+------------------------------------------------------------------------
 unique_id | 1.0     | public | Non-Standard Time-based K-sorted, Lexicographically Unique Identifiers
(1 row)

fabrizio=# CREATE SEQUENCE instagram_seq;
CREATE SEQUENCE
fabrizio=# -- Using default shard_id = 0
fabrizio=# SELECT unique_id_instagram('instagram_seq');
 unique_id_instagram
---------------------
 2563729919292997633
(1 row)

fabrizio=# -- Using default shard_id = 1
fabrizio=# SELECT unique_id_instagram('instagram_seq', 1);
 unique_id_instagram
---------------------
 2563729919292998658
(1 row)
```

### 主要对象

经审查的安装接口包含以下主要对象：

| 对象 | 类型 | 运维作用 |
|---|---|---|
| `unique_id_instagram` | FUNCTION | 经审查安装接口中的可调用函数。 |
| `unique_id_sonyflake` | FUNCTION | 经审查安装接口中的可调用函数。 |

### 运维与边界

- 目录生命周期为 abandoned；生产使用前应测试升级、备份恢复与服务器兼容性。
