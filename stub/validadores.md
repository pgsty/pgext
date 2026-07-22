## Usage

Sources:

- [Official upstream documentation](https://github.com/guedes/validadores/blob/63ea247a692f6995c54a24e80f1a39802ede59b6/README.md)
- [Official extension SQL](https://github.com/guedes/validadores/blob/63ea247a692f6995c54a24e80f1a39802ede59b6/sql/validadores.sql)
- [Official regression examples](https://github.com/guedes/validadores/blob/63ea247a692f6995c54a24e80f1a39802ede59b6/test/sql/base.sql)

`validadores` 0.0.2 is a pure-SQL Brazilian CPF check-digit helper. Despite the broader Portuguese project description mentioning CNPJ, the reviewed release installs only a CPF function, unary operator, and numeric domain; it contains no CNPJ validator.

### Validate and Constrain CPF Values

Create the extension, call the function or postfix operator, and use the domain when its exact acceptance rules fit the application:

```sql
CREATE EXTENSION validadores;

SELECT cpf_valido(59328253241);
SELECT 59328253241 #? AS cpf_is_valid;

CREATE TABLE pessoa (
  nro_cpf cpf
);
```

`cpf_valido(numeric)` left-pads the numeric text to 11 digits, calculates two check digits, and returns a boolean. The postfix `#?` operator calls the same function. The `cpf` domain is `numeric` with a `CHECK (cpf_valido(VALUE))` constraint.

### Data-Model Caveats

CPF is an identifier, not a quantity. A numeric column discards punctuation and can obscure leading-zero handling at application boundaries. If exact input representation matters, store canonical text and apply an explicitly reviewed validation function or domain instead of relying on implicit numeric conversion.

The source implements only its check-digit arithmetic. It does not explicitly reject all-equal digit sequences commonly treated as invalid CPFs, verify ownership or issuance, validate formatted text, or normalize CNPJ. Therefore, passing this function is not proof that an identifier is real or acceptable under current business or regulatory rules. Add application-specific repeated-digit and canonical-length tests, compare against an authoritative Brazilian validation specification, and test known valid and invalid fixtures before using the domain. The project is old; inspect behavior again before a PostgreSQL upgrade.
