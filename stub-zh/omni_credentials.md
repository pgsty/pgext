


## 用法

> [omni_credentials: 应用凭证管理](https://docs.omnigres.org/omni_credentials/credentials/)

`omni_credentials` 扩展提供安全的加密凭证存储和访问控制。它是一个模板化扩展。

### 初始化

```sql
SELECT omni_credentials.instantiate(
    schema  => 'omni_credentials',
    env_var => 'OMNI_CREDENTIALS_MASTER_PASSWORD'
);
```

### 凭证视图

主要接口是 `credentials` 视图：

| 列 | 类型 | 用途 |
|:---|:-----|:-----|
| `name` | text | 凭证标识符 |
| `value` | bytea | 加密的凭证数据 |
| `kind` | credential_kind | 类型：`api_key`、`api_secret`、`password` 等 |
| `principal` | regrole | 拥有凭证的 PostgreSQL 角色 |
| `scope` | jsonb | 资源约束（`{"all": true}` 表示通用） |

直接查询和更新视图；更改会自动传播。行级安全确保只有对 `principal` 有授权的角色才能访问凭证。

### 文件存储（开发用）

```sql
SELECT omni_credentials.instantiate_file_store(filename, schema);
-- 从文件重新加载凭证：
SELECT credential_file_store_reload(filename);
```

将文件中的现有记录导入加密存储，并将表中缺失的凭证导出到文件。凭证变更时自动更新文件。
