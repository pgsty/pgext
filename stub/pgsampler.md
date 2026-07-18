## Usage

Sources:

- [Official extension control file](https://github.com/no0p/pgsampler/blob/9feffce68a902dc9b35fd1ead5ef2c8095e25124/pgsampler.control)
- [Official upstream documentation](https://github.com/no0p/pgsampler/wiki)
- [Official project or provider page](https://no0p.github.io/pgsampler/)

`pgsampler` — Experimental PostgreSQL background worker that collects time-series status and performance metrics.

The reviewed catalog snapshot records version `1.0`, kind `preload`, and implementation language `C`.

This component has no standalone extension DDL in the curated record; integrate it only through the upstream-documented load or provider workflow.
The official material contains an experimental, deprecated, unsupported, or explicit warning boundary; read it in full and test failure cases before non-lab use.

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
