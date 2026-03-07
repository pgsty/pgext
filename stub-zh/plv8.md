

## 用法

```sql
CREATE EXTENSION plv8;

SELECT plv8_version();
SELECT plv8_info();

DO $$ plv8.elog(NOTICE, plv8.version); $$ LANGUAGE plv8;
```

示例：

```sql
CREATE FUNCTION plv8_test(keys TEXT[], vals TEXT[]) RETURNS JSON AS $$
    var o = {};
    for(var i=0; i<keys.length; i++){
        o[keys[i]] = vals[i];
    }
    return o;
$$ LANGUAGE plv8 IMMUTABLE STRICT;


SELECT plv8_test(ARRAY['name', 'age'], ARRAY['Tom', '29']);
```


## 构建

Plv8 在 EL10（x86/arm）上构建会遇到以下问题：

- 找不到 g++ 的问题
- g++ 14 需要包含 `<algorithm>` 头文件的问题
- LTO 问题，g++14 链接时优化（Link Time Optimization）相关错误



