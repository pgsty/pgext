## Usage

Sources:

- [mdbqs 0.1 install SQL at the reviewed commit](https://github.com/NikitOS94/MDBQT/blob/9b4c5b5b17973660ddf83a02d3b4ca2f8537af94/mdbqs--0.1.sql)
- [mdbqs regression SQL at the reviewed commit](https://github.com/NikitOS94/MDBQT/blob/9b4c5b5b17973660ddf83a02d3b4ca2f8537af94/sql/mdbqs_test.sql)
- [mdbqs.control at the reviewed commit](https://github.com/NikitOS94/MDBQT/blob/9b4c5b5b17973660ddf83a02d3b4ca2f8537af94/mdbqs.control)

`mdbqs` adds the `mdbquery` type and the `<=>` operator for matching a JSONB value against a MongoDB-style predicate. The regression suite exercises comparison and membership predicates such as `$lt`, `$gte`, and `$in`, logical forms such as `$and` and `$or`, and checks including `$exists` and `$type`.

### Match JSONB Against a Predicate

```sql
CREATE EXTENSION jsquery;
CREATE EXTENSION mdbqs;

SELECT '{"a": 2}'::jsonb
       <=> '{ a: { $lt: 3 } }'::mdbquery AS matches;

SELECT '{"quantity": 5, "price": 100}'::jsonb
       <=> '{ $and: [ { quantity: { $lt: 20 } }, { price: 100 } ] }'::mdbquery
       AS matches;
```

The install SQL defines the operator in both argument orders, so an `mdbquery` can also appear on the left of a JSONB value.

### Caveats

- The source invokes `jsquery` functionality and its regression setup creates that extension first, but `mdbqs.control` does not declare the dependency. Install `jsquery` explicitly.
- The repository and package are named MDBQT/`mdbqt`, while the installable extension is `mdbqs` version 0.1.
- Upstream publishes no README, license, release history, or modern compatibility matrix. The C source uses PostgreSQL internal APIs from its 2017-era codebase; validate or port it for the exact server version before use.
