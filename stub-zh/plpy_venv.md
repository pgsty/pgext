## 用法

来源：

- [官方文档](https://github.com/dpage/plpy_venv/blob/0fc90895033be8a10351a6383a275bcf09f2e404/README.md)
- [扩展控制文件](https://github.com/dpage/plpy_venv/blob/0fc90895033be8a10351a6383a275bcf09f2e404/plpy_venv.control)
- [官方仓库](https://github.com/dpage/plpy_venv)

`plpy_venv` 为 PL/Python 函数选择并管理 Python 虚拟环境。

### 启用

为目标服务器安装文件，然后在目标数据库中创建 `plpy_venv`：

```sql
CREATE EXTENSION plpy_venv CASCADE;
```

经审查的控制文件或官方流程要求 `plpython3u`。只有服务器已安装这些扩展文件时，`CASCADE` 才会成功。

### 核心流程

以下示例取自经审查的上游文档。使用前应调整对象名、路径、凭据与工作负载参数。

```sql
plpy=# SELECT plpy_venv.create_venv('myvenv');
ERROR:  plpy.Error: Virtual environment directory /path/to/postgresql/data/venvs/myvenv already exists.
CONTEXT:  Traceback (most recent call last):
  PL/Python function "create_venv", line 30, in <module>
    plpy.error('Virtual environment directory {} already exists.'.format(venv_dir))
PL/Python function "create_venv"
```

### 主要对象

官方来源通过动态方式或供应商工具定义扩展接口；授权前应检查实际安装版本。

### 运维与边界

- 扩展会在 `plpy_venv` 下固定或创建模式对象；权限与备份审查应包含这些对象。
- 目录生命周期为 abandoned；生产使用前应测试升级、备份恢复与服务器兼容性。
