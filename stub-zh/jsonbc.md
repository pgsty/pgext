## 用法

来源：

- [jsonbc--1.0.sql](https://github.com/akorotkov/jsonbc/blob/c59545546211f17cabb419d9424e33774308f242/jsonbc--1.0.sql)
- [Makefile](https://github.com/akorotkov/jsonbc/blob/c59545546211f17cabb419d9424e33774308f242/Makefile)
- [jsonbc.control](https://github.com/akorotkov/jsonbc/blob/c59545546211f17cabb419d9424e33774308f242/jsonbc.control)

`jsonbc` 是历史上的压缩类 JSONB 类型。版本化 SQL 提供键字典、JSON 风格算子，以及 B-tree/hash/GIN 算子类。

### 核心用法

```sql
CREATE EXTENSION jsonbc;
SELECT '{"name":"Alice"}'::jsonbc;
SELECT '{"name":"Alice"}'::jsonbc->>'name';
```

### 运行边界

`get_id_by_name` 与 `get_name_by_id` 访问扩展的键字典；对象和数组访问、包含、存在判断及展开函数作用于自定义类型。字典与已存值共同组成持久化格式，不应单独删除字典数据。

安装要求超级用户及编译后的 C 库，未声明预加载要求。核对的仓库文件没有当前支持文档、主版本矩阵或许可证。这是历史源码收录，并非经过验证的原生 JSONB 替代品；写入数据前应测试安装、往返转换、备份和恢复。
