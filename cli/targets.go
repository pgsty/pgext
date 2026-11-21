package cli

// Retired repositories remain addressable for historical package provenance.
// Fetch/scan select active platforms and, for per-major RPM repositories,
// active PG majors. Shared APT and Pigsty repositories have no major suffix.
const activeRepositoryPredicate = `r.os IN (SELECT os FROM pgext.active_os)
	AND coalesce(substring(r.id FROM '[.](?:pgdg|pgnf)([0-9]+)$')::integer
	             IN (SELECT pg FROM pgext.active_pg), true)`
