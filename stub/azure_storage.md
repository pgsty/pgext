## Usage

Sources:

- [Azure Storage extension overview](https://learn.microsoft.com/en-us/azure/postgresql/extensions/concepts-storage-extension)
- [Azure Storage extension configuration](https://learn.microsoft.com/en-us/azure/postgresql/extensions/how-to-configure-azure-storage-extension)
- [Azure Storage function reference](https://learn.microsoft.com/en-us/azure/postgresql/extensions/reference-azure-storage-extension)

`azure_storage` is a Microsoft-managed extension for moving data between Azure Database for PostgreSQL Flexible Server and Azure Blob Storage with SQL. It supports listing, importing, and exporting blobs; it is not a portable extension for self-managed PostgreSQL.

### Core Workflow

First allow the extension and load its library through the Flexible Server configuration described by Microsoft, then restart when required and create it in every database that will use it. Microsoft recommends managed identity instead of Shared Key authentication.

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

The role that registers accounts must be a member of `azure_storage_admin`. Other roles need an explicit `azure_storage.account_user_add` grant before using storage functions. The managed identity also needs appropriate Azure data-plane permissions, such as Blob Data Reader for reads or Blob Data Contributor for writes.

### Main Objects

- `azure_storage.account_add`, `azure_storage.account_remove`, and `azure_storage.account_list` manage storage-account registrations.
- `azure_storage.account_user_add` and `azure_storage.account_user_remove` manage database-role access.
- `azure_storage.blob_list` lists objects; `azure_storage.blob_get` imports rows; `azure_storage.blob_put` exports rows.
- `azure_storage.options_csv_get`, `azure_storage.options_tsv`, `azure_storage.options_binary`, and `azure_storage.options_parquet` build decoder options.

### Caveats

Format and compression can be inferred from a blob suffix or supplied explicitly. Supported formats include CSV, TSV, text-like files, PostgreSQL binary, and Parquet; supported compression depends on the encoder. Account keys are sensitive credentials, so prefer managed identity and rotate any Shared Keys. Network access, container existence, role allowlists, and Azure IAM are all separate prerequisites. Loading the library and changing server-level settings can require a restart.
