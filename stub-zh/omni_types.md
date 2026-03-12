


## 用法

> [omni_types: 高级类型](https://docs.omnigres.org/omni_types/function_signature_types/)

`omni_types` 扩展提供高级类型工具，包括捕获完整函数签名并允许直接调用的函数签名类型。

### 定义函数签名类型

**显式定义：**

```sql
SELECT omni_types.function_signature_type('sig', 'text', 'int');
-- 为接受 text 并返回 int 的函数创建 'sig' 类型
```

**从现有函数创建：**

```sql
SELECT omni_types.function_signature_type_of('sig', 'length(text)');
```

### 将函数转换为该类型

```sql
SELECT 'length'::sig;
```

非失败验证：

```sql
SELECT sig_conforming_function('length');  -- 无匹配时返回 NULL
```

### 调用类型化函数

```sql
SELECT call_sig('length', 'hello');  -- 返回 5
```

`call_<TYPE>` 函数为每个签名类型自动生成，使用提供的参数执行引用的函数。
