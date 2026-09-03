## 用法

来源：

- [官方文档](https://github.com/fabriziomello/pg_toolkit_brazil/blob/25aa4d0472582d17af8b84e8c80cf8a1db8b8943/README.md)
- [扩展控制文件](https://github.com/fabriziomello/pg_toolkit_brazil/blob/25aa4d0472582d17af8b84e8c80cf8a1db8b8943/pg_toolkit_brazil.control)
- [官方仓库](https://github.com/fabriziomello/pg_toolkit_brazil)

`pg_toolkit_brazil` 提供巴西 CPF 与 CNPJ 类型、校验、格式化、转换和操作符。

### 启用

为目标服务器安装文件，然后在目标数据库中创建 `pg_toolkit_brazil`：

```sql
CREATE EXTENSION pg_toolkit_brazil;
```

### 核心流程

以下示例取自经审查的上游文档。使用前应调整对象名、路径、凭据与工作负载参数。

```sql
CREATE TABLE pessoa (id bigint GENERATED ALWAYS AS IDENTITY, cpf cpf);
INSERT INTO pessoa (cpf) VALUES (cpf '40100276300');
SELECT id, cpf FROM pessoa;
```

### 主要对象

经审查的安装接口包含以下主要对象：

| 对象 | 类型 | 运维作用 |
|---|---|---|
| `cnpj` | TYPE | 扩展创建的用户数据类型。 |
| `cnpj_cmp` | FUNCTION | 经审查安装接口中的可调用函数。 |
| `cnpj_eq` | FUNCTION | 经审查安装接口中的可调用函数。 |
| `cnpj_ge` | FUNCTION | 经审查安装接口中的可调用函数。 |
| `cnpj_gt` | FUNCTION | 经审查安装接口中的可调用函数。 |
| `cnpj_le` | FUNCTION | 经审查安装接口中的可调用函数。 |
| `cnpj_lt` | FUNCTION | 经审查安装接口中的可调用函数。 |
| `cnpj_ne` | FUNCTION | 经审查安装接口中的可调用函数。 |

### 运维与边界

- 目录生命周期为 abandoned；生产使用前应测试升级、备份恢复与服务器兼容性。
