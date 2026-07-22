## 用法

来源：

- [Azure Storage 扩展概览](https://learn.microsoft.com/en-us/azure/postgresql/extensions/concepts-storage-extension)
- [Azure Storage 扩展配置](https://learn.microsoft.com/en-us/azure/postgresql/extensions/how-to-configure-azure-storage-extension)
- [Azure Storage 函数参考](https://learn.microsoft.com/en-us/azure/postgresql/extensions/reference-azure-storage-extension)

`azure_storage` 是 Microsoft 托管的扩展，可通过 SQL 在 Azure Database for PostgreSQL Flexible Server 与 Azure Blob Storage 之间移动数据。它支持列出、导入和导出 Blob；这不是供自托管 PostgreSQL 使用的可移植扩展。

### 核心流程

先按照 Microsoft 文档在 Flexible Server 配置中允许该扩展并加载其库，必要时重启，然后在每个需要使用它的数据库中创建扩展。Microsoft 建议使用托管身份而不是 Shared Key 身份验证。

```sql
CREATE EXTENSION azure_storage;

SELECT azure_storage.account_add(
    azure_storage.account_options_managed_identity('mystorage', 'blob')
);

SELECT azure_storage.account_user_add('mystorage', 'etl_role');

SELECT path, bytes, last_modified
FROM azure_storage.blob_list('mystorage', 'incoming', 'daily/');

SELECT *
FROM azure_storage.blob_get(
    'mystorage',
    'incoming',
    'daily/customers.csv',
    'csv',
    'none',
    azure_storage.options_csv_get(header => true)
) AS t(customer_id bigint, name text);
```

注册存储账户的角色必须属于 `azure_storage_admin`。其他角色在调用存储函数前，需要通过 `azure_storage.account_user_add` 显式授权。托管身份也必须具备相应的 Azure 数据平面权限，例如读取所需的 Blob Data Reader 或写入所需的 Blob Data Contributor。

### 主要对象

- `azure_storage.account_add`、`azure_storage.account_remove` 与 `azure_storage.account_list` 管理存储账户注册。
- `azure_storage.account_user_add` 与 `azure_storage.account_user_remove` 管理数据库角色访问权。
- `azure_storage.blob_list` 列出对象；`azure_storage.blob_get` 导入行；`azure_storage.blob_put` 导出行。
- `azure_storage.options_csv_get`、`azure_storage.options_tsv`、`azure_storage.options_binary` 与 `azure_storage.options_parquet` 构造解码选项。

### 注意事项

格式和压缩方式既可从 Blob 后缀推断，也可显式指定。支持的格式包括 CSV、TSV、文本类文件、PostgreSQL binary 与 Parquet；可用压缩算法取决于编码器。账户密钥属于敏感凭据，因此应优先使用托管身份，并轮换所有 Shared Keys。网络访问、容器存在、角色允许列表与 Azure IAM 是彼此独立的前置条件。加载库及修改服务器级设置可能需要重启。
