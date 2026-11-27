VERSION ?= v1.8.0
COMMIT  := $(shell git rev-parse --short HEAD 2>/dev/null || echo unknown)
LD_FLAGS := -s -w -X pgext/cmd.Version=$(VERSION) -X pgext/cmd.Commit=$(COMMIT)
PGURL   ?= postgres:///data

.DEFAULT_GOAL := build
.PHONY: build b dev serve s gen check check-stubs d da dump load dumps loads dd d2 dump2 load2 gr arm amd 2j rs

b: build
build:
	go build -ldflags "$(LD_FLAGS)" -o pgext .

dev serve:
	go run . serve
s: serve

gen:
	go run . gen site

check: check-stubs
	go vet ./ ./cmd/... ./cli/... ./server/... ./db/...
	go test -count=1 ./cmd/... ./cli/... ./server/... ./db/...
	node --test server/*_test.mjs
	python3 -m unittest -v bin/test_changelog.py bin/test_rebuild_manifest.py

check-stubs:
	go run . gen lint

# Build, upload and restart the Go daemon on jp.
d: amd 2j rs

da: dump dd d2
# dump extension data to data dir
dump:
	psql $(PGURL) -c "COPY (SELECT * FROM pgext.pg         ORDER BY pg DESC)  TO STDOUT CSV HEADER;"  > db/pg.csv
	psql $(PGURL) -c "COPY (SELECT * FROM pgext.os         ORDER BY os_major, CASE WHEN os_code IN ('el10','d11','u22') THEN CASE os_arch WHEN 'aarch64' THEN 0 ELSE 1 END ELSE CASE os_arch WHEN 'x86_64' THEN 0 ELSE 1 END END) TO STDOUT CSV HEADER;"  > db/os.csv
	psql $(PGURL) -c "COPY (SELECT * FROM pgext.category   ORDER BY id)       TO STDOUT CSV HEADER;"  > db/category.csv
	psql $(PGURL) -c "COPY (SELECT * FROM pgext.repository ORDER BY id)       TO STDOUT CSV HEADER;"  > db/repository.csv
	psql $(PGURL) -c "COPY (SELECT * FROM pgext.extension  ORDER BY id)       TO STDOUT CSV HEADER;"  > db/extension.csv
	psql $(PGURL) -c "COPY (SELECT * FROM pgext.universe  ORDER BY id)       TO STDOUT CSV HEADER;"  > db/universe.csv
	psql $(PGURL) -c "COPY (SELECT * FROM pgext.changelog ORDER BY ds DESC, pkg) TO STDOUT CSV HEADER;" > db/changelog.csv
load:
	cat db/pg.csv         | psql $(PGURL) -c "COPY pgext.pg         FROM STDIN CSV HEADER;"
	cat db/os.csv         | psql $(PGURL) -c "COPY pgext.os         FROM STDIN CSV HEADER;"
	cat db/category.csv   | psql $(PGURL) -c "COPY pgext.category   FROM STDIN CSV HEADER;"
	cat db/repository.csv | psql $(PGURL) -c "COPY pgext.repository FROM STDIN CSV HEADER;"
	cat db/extension.csv  | psql $(PGURL) -c "COPY pgext.extension  FROM STDIN CSV HEADER;"
	cat db/universe.csv   | psql $(PGURL) -c "COPY pgext.universe   FROM STDIN CSV HEADER;"
	cat db/changelog.csv  | psql $(PGURL) -c "COPY pgext.changelog  FROM STDIN CSV HEADER;"

dumps:
	pg_dump $(PGURL) -c -x -n pgext -f tmp/pgext.sql
loads:
	psql $(PGURL) -f tmp/pgext.sql

# dump extension data to data dir
dd:
	psql $(PGURL) -c "COPY (SELECT * FROM pgext.category   ORDER BY id)       TO STDOUT CSV HEADER;"  > ~/pgsty/repo/ext/data/category.csv
	psql $(PGURL) -c "COPY (SELECT * FROM pgext.repository ORDER BY id)       TO STDOUT CSV HEADER;"  > ~/pgsty/repo/ext/data/repository.csv
	psql $(PGURL) -c "COPY (SELECT * FROM pgext.pig) TO STDOUT CSV HEADER;"  > ~/pgsty/repo/ext/data/extension.csv

# dump extension data to data dir
d2:
	psql $(PGURL) -c "COPY (SELECT * FROM pgext.pig) TO STDOUT CSV HEADER;"  > ~/pgsty/pig/cli/ext/assets/extension.csv


# dump extension data to data dir
dump2:
	psql $(PGURL) -c "COPY (SELECT * FROM pgext.bin ORDER BY pg,os,name DESC)  TO STDOUT CSV HEADER;"  > db/bin.csv
	psql $(PGURL) -c "COPY (SELECT * FROM pgext.pkg ORDER BY pkg,os,pg) TO STDOUT CSV HEADER;"         > db/pkg.csv
load2:
	cat db/bin.csv         | psql $(PGURL) -c "COPY pgext.bin FROM STDIN CSV HEADER;"
	cat db/pkg.csv         | psql $(PGURL) -c "COPY pgext.pkg FROM STDIN CSV HEADER;"

gr:
	goreleaser release --clean --skip=publish

arm:
	CGO_ENABLED=0 GOOS=linux  GOARCH=arm64 go build -ldflags "$(LD_FLAGS)" -o pgext .
amd:
	CGO_ENABLED=0 GOOS=linux  GOARCH=amd64 go build -ldflags "$(LD_FLAGS)" -o pgext .

# ship binary to jp (atomic swap avoids ETXTBSY on the running daemon)
2j:
	scp pgext jp:/usr/bin/pgext.new
	ssh jp 'chmod 755 /usr/bin/pgext.new && mv -f /usr/bin/pgext.new /usr/bin/pgext'

# restart pgext daemon on jp
rs:
	ssh jp 'systemctl restart pgext && systemctl is-active pgext'
