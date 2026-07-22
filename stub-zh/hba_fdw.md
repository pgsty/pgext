## 用法

来源：

- [Official README](https://github.com/fabriziomello/hba_fdw/blob/fcc1cd8eee8f997538b917f095caef5373a83e7d/README.md)
- [Extension SQL](https://github.com/fabriziomello/hba_fdw/blob/fcc1cd8eee8f997538b917f095caef5373a83e7d/hba_fdw--1.0.sql)
- [FDW implementation](https://github.com/fabriziomello/hba_fdw/blob/fcc1cd8eee8f997538b917f095caef5373a83e7d/hba_fdw.c)

hba_fdw 是一个进行中的外部数据包装器骨架，原计划通过 SQL 管理 pg_hba.conf。所核验的实现既不读取也不修改认证规则：每次扫描都会立即返回零行，验证器接受所有选项，而且没有插入、更新或删除回调。

### 核心流程

唯一完成的操作是注册这个空包装器：

```sql
CREATE EXTENSION hba_fdw;

SELECT fdwname
FROM pg_foreign_data_wrapper
WHERE fdwname = 'hba_fdw';
```

项目没有可据以配置的权威服务器选项或外部表结构，创建外部表也不会让 pg_hba.conf 变得可查询。

### 重要对象

- `hba_fdw_handler` 返回最小的扫描回调表。
- `hba_fdw_validator` 当前不执行任何校验。
- `hba_fdw` 由安装 SQL 注册为外部数据包装器。

### 运维说明

不要围绕 hba_fdw 构建认证管理。它没有解析器、重载机制、并发保护、权限模型或安全写入器，README 也只承诺未来功能。pg_hba.conf 应由受支持的配置管理工具维护，并用 PostgreSQL 自身工具校验变更。
