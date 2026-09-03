## 用法

来源：

- [官方文档](https://github.com/fukanchik/probackup_ctl/blob/e03cb514d05c82830379d152924bfeaf81e42ef4/README.md)
- [扩展控制文件](https://github.com/fukanchik/probackup_ctl/blob/e03cb514d05c82830379d152924bfeaf81e42ef4/probackup_ctl.control)
- [官方仓库](https://github.com/fukanchik/probackup_ctl)

`probackup_ctl` 检查并操作 pg_probackup 目录，支持文件系统、SFTP 与 S3 存储。

### 启用

为目标服务器安装文件，然后在目标数据库中创建 `probackup_ctl`：

```sql
CREATE EXTENSION probackup_ctl;
```

### 核心流程

以下示例取自经审查的上游文档。使用前应调整对象名、路径、凭据与工作负载参数。

```sql
select * from probackup.log(catalog_id=>2,pg_instance=>'dba1',backup_id=>(select backup_id from probackup.show(catalog_id=>2)));
```

### 主要对象

经审查的安装接口包含以下主要对象：

| 对象 | 类型 | 运维作用 |
|---|---|---|
| `probackup.backup` | FUNCTION | 经审查安装接口中的可调用函数。 |
| `probackup.show` | FUNCTION | 经审查安装接口中的可调用函数。 |
| `probackup.catalogs` | TABLE | 扩展自有表；备份与升级时需计入其中数据。 |
| `probackup.register_catalog` | FUNCTION | 经审查安装接口中的可调用函数。 |
| `probackup.delete` | FUNCTION | 经审查安装接口中的可调用函数。 |
| `probackup.log` | FUNCTION | 经审查安装接口中的可调用函数。 |
| `probackup.s3_config` | TABLE | 扩展自有表；备份与升级时需计入其中数据。 |
| `probackup.sftp_config` | TABLE | 扩展自有表；备份与升级时需计入其中数据。 |

### 运维与边界

- 扩展会在 `probackup` 下固定或创建模式对象；权限与备份审查应包含这些对象。
