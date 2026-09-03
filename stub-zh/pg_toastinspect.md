## 用法

来源：

- [官方文档](https://github.com/yansheng836/pg_toastinspect/blob/65571a9267e3123eb0d9d477dff6c5a5246e948a/README.md)
- [扩展控制文件](https://github.com/yansheng836/pg_toastinspect/blob/65571a9267e3123eb0d9d477dff6c5a5246e948a/pg_toastinspect.control)
- [官方仓库](https://github.com/yansheng836/pg_toastinspect)

`pg_toastinspect` 低层检查 TOAST 分块标识、尺寸与存储值。

### 启用

为目标服务器安装文件，然后在目标数据库中创建 `pg_toastinspect`：

```sql
CREATE EXTENSION pg_toastinspect;
```

### 核心流程

以下示例取自经审查的上游文档。使用前应调整对象名、路径、凭据与工作负载参数。

```sql
-- Get the chunk ID of a TOASTed value
SELECT ctid, id, get_toast_chunk_id(response_content), response_content
FROM ai_call_log_copy1
WHERE id = 668310181431480320;

-- Get complete TOAST information
SELECT ctid, id, get_toast_info(response_content), response_content
FROM ai_call_log_copy1
WHERE id = 668310181431480320;

-- Access individual fields from the composite type
SELECT
    ctid,
    id,
    (get_toast_info(response_content)).raw_size,
    (get_toast_info(response_content)).ext_size,
    (get_toast_info(response_content)).chunk_id,
    (get_toast_info(response_content)).reltoastrelid
FROM ai_call_log_copy1
WHERE id = 668310181431480320;
```

### 主要对象

经审查的安装接口包含以下主要对象：

| 对象 | 类型 | 运维作用 |
|---|---|---|
| `get_toast_info` | FUNCTION | 经审查安装接口中的可调用函数。 |
| `get_toast_chunk_id` | FUNCTION | 经审查安装接口中的可调用函数。 |
| `toast_pointer_info` | TYPE | 扩展创建的用户数据类型。 |

### 运维与边界

- 应将扩展升级视为数据库变更：先审查上游升级路径、权限、锁以及备份恢复行为。
