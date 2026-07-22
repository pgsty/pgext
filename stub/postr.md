## Usage

Sources:

- [Official README](https://github.com/mrcsparker/postr/blob/e1dbaa89db80e76a7d0cdff28cf3314047d5db2b/README.md)
- [Extension control file](https://github.com/mrcsparker/postr/blob/e1dbaa89db80e76a7d0cdff28cf3314047d5db2b/postr.control)
- [Official examples index](https://github.com/mrcsparker/postr/blob/e1dbaa89db80e76a7d0cdff28cf3314047d5db2b/examples/README.md)

`postr` embeds Ruby in PostgreSQL and registers an untrusted procedural language named `ruby`. It lets SQL functions, triggers, set-returning functions, and inline blocks execute Ruby in a backend process, but it is explicitly a superuser-only prototype rather than a sandboxed multi-tenant language.

### Core Workflow

After a trusted administrator installs the extension, define a Ruby function with a `def call(...)` entry point and invoke it through normal SQL:

```sql
CREATE EXTENSION postr;

CREATE FUNCTION public.ruby_greet(name text, excited boolean)
RETURNS text
LANGUAGE ruby
AS $ruby$
  def call(name:, excited:)
    message = "hello, #{name}"
    excited ? "#{message}!" : message
  end
$ruby$;

SELECT public.ruby_greet('postgres', true);
```

Arguments are available by name, through the `args` mapping, and in call order through `argv`. The runtime maps common PostgreSQL scalar, array, composite, network, and built-in range values to Ruby objects. Unsupported pseudotypes and output shapes are rejected during function validation.

### Database Access and Results

Ruby code can call back into the current PostgreSQL transaction through the `Postr` prelude:

- `Postr.exec` runs parameterized SQL and returns the processed row count.
- `Postr.select` returns row hashes or yields rows, while `Postr.first` and `Postr.value` select one row or value.
- `Postr.param` supplies an explicit PostgreSQL type for ambiguous Ruby values.
- `Postr.notice`, `Postr.warn`, and `Postr.quote_ident` expose server messaging and identifier quoting.
- `Postr.transaction` runs a Ruby block in an internal subtransaction; it behaves like a savepoint, not an independent commit.

The language also supports `DO LANGUAGE ruby`, row and statement triggers, iterable or yield-driven `SETOF` results, records, `OUT`/`INOUT`/`TABLE` signatures, and variadic arguments. Consult the official examples for the exact return contract of each shape.

### Curated Requires

Ruby `require` is restricted by PostgreSQL settings. Administrators may configure `postr.extra_load_paths` and `postr.allowed_requires` for approved features, or `postr.gem_home` and `postr.allowed_gems` for vendored gems:

```sql
SET LOCAL postr.extra_load_paths = '/opt/postr/ruby';
SET LOCAL postr.allowed_requires = 'postr_helpers';
```

These allowlists reduce accidental module exposure but do not sandbox Ruby. Policy state and compiled callables are cached per backend, so configuration and code changes should be exercised with fresh sessions during deployment tests.

### Security and Compatibility

The control file marks version 0.1.0 non-relocatable and superuser-only. Ruby executes inside the database backend and can consume its memory, CPU, and process privileges; only highly trusted roles should be able to define or alter `LANGUAGE ruby` code. External libraries and gems run in the same trust boundary.

The project's packaged smoke test targets PostgreSQL 17 on Linux, while source features cover PostgreSQL 13 through 18. Ruby support follows the embedded `magnus` library's documented range, and the README distinguishes local Ruby 4 testing from a compatibility guarantee. Pin PostgreSQL, Ruby, extension, and gem versions, apply statement/resource controls, and treat backend crashes or blocking Ruby calls as database availability risks.
