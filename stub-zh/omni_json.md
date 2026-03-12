


## 用法

> [omni_json: JSON 工具箱](https://docs.omnigres.org/omni_json/table_mapping/)

`omni_json` 扩展提供表到 JSON 的映射功能，支持在行转换为 JSON 时进行自定义变换。

### 定义表映射

```sql
SELECT omni_json.define_table_mapping(example, '{}');
```

### 列重命名

```sql
SELECT omni_json.define_table_mapping(example, '{
    "columns": { "dob": { "path": "date_of_birth" } }
}');
```

### 嵌套键

```sql
SELECT omni_json.define_table_mapping(example, '{
    "columns": {
        "first_name": { "path": ["name", "first"] },
        "last_name":  { "path": ["name", "last"] }
    }
}');
```

### 列排除

```sql
SELECT omni_json.define_table_mapping(example, '{
    "columns": { "dob": { "exclude": true } }
}');
```

### 列变换

输入（JSON 到记录）和输出（记录到 JSON）变换：

```json
{
    "columns": {
        "password": {
            "transform": {
                "input":  { "type": "text", "function": "encrypt_password" },
                "output": { "type": "text", "function": "mask_password" }
            }
        }
    }
}
```

### 以 JSON 形式查询

```sql
SELECT to_jsonb(products.*) FROM products;
```

### 从 JSON 插入

```sql
INSERT INTO people (dob, first_name, last_name)
    (SELECT dob, first_name, last_name
     FROM jsonb_populate_record(null::people,
         '{"first_name": "Jane", "last_name": "Doe", "dob": "1981-12-12"}'));
```
