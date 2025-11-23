#==============================================================#
# File      :   Makefile
# Ctime     :   2024-10-30
# Mtime     :   2025-07-05
# Desc      :   Makefile shortcuts
# Path      :   Makefile
# Copyright (C) 2019-2025 Ruohang Feng
#==============================================================#
VERSION=v0.7.2
default: v d

PGURL="postgres:///vonng"

# run next dev server (3000)
d: dev
dev:
	hugo serve

# building (SSG) - new optimized version
b: build
build:
	hugo build

# building with legacy MDX files
build-mdx: gen-mdx
	TURBO_CONCURRENCY=25 npm run build

# open browser to view the app
v: view
view:
	open 'http://localhost:1313'

# remove the build output directory
c: clean
clean:
	rm -rf out

# run pnpm install
i: install
install:
	pnpm install

# update npm dependencies to the latest version
u: update
update:
	pnpm update


# serve wiht
s: serve
serve:
	cd out && python3 -m http.server 3001


# dump extension data to data dir
dump:
	psql $(PGURL) -c "COPY (SELECT * FROM pgext.pg         ORDER BY pg DESC)  TO STDOUT CSV HEADER;"  > db/pg.csv
	psql $(PGURL) -c "COPY (SELECT * FROM pgext.os         ORDER BY os_major) TO STDOUT CSV HEADER;"  > db/os.csv
	psql $(PGURL) -c "COPY (SELECT * FROM pgext.category   ORDER BY id)       TO STDOUT CSV HEADER;"  > db/category.csv
	psql $(PGURL) -c "COPY (SELECT * FROM pgext.repository ORDER BY id)       TO STDOUT CSV HEADER;"  > db/repository.csv
	psql $(PGURL) -c "COPY (SELECT * FROM pgext.extension  ORDER BY id)       TO STDOUT CSV HEADER;"  > db/extension.csv
load:
	cat db/pg.csv         | psql $(PGURL) -c "COPY pgext.pg         FROM STDIN CSV HEADER;"
	cat db/os.csv         | psql $(PGURL) -c "COPY pgext.os         FROM STDIN CSV HEADER;"
	cat db/category.csv   | psql $(PGURL) -c "COPY pgext.category   FROM STDIN CSV HEADER;"
	cat db/repository.csv | psql $(PGURL) -c "COPY pgext.repository FROM STDIN CSV HEADER;"
	cat db/extension.csv  | psql $(PGURL) -c "COPY pgext.extension  FROM STDIN CSV HEADER;"

# dump extension data to data dir
dump2:
	psql $(PGURL) -c "COPY (SELECT * FROM pgext.bin ORDER BY pg,os,name DESC)  TO STDOUT CSV HEADER;"  > db/bin.csv
	psql $(PGURL) -c "COPY (SELECT * FROM pgext.pkg ORDER BY pkg,os,pg) TO STDOUT CSV HEADER;"         > db/pkg.csv
load2:
	cat db/bin.csv         | psql $(PGURL) -c "COPY pgext.bin FROM STDIN CSV HEADER;"
	cat db/pkg.csv         | psql $(PGURL) -c "COPY pgext.pkg FROM STDIN CSV HEADER;"

gr:
	goreleaser release --clean --skip=publish

# load extension data from data dir
#load:
#	psql $(PGURL) -c "TRUNCATE ext.extension; COPY ext.extension FROM '/Users/vonng/pgsty/extension/data/extension.csv' CSV HEADER;"

# generate extension data in JSON format (optimized)
gen-ext:
	@echo "Generating extension JSON data..."
	@source ~/.venv/bin/activate && python bin/gen-ext-json.py
	@echo "Extension JSON data generated in data/extensions/"

# generate extension MDX files (legacy)
gen-mdx:
	@echo "Generating extension MDX files..."
	@source ~/.venv/bin/activate && python bin/gen-ext.py
	@echo "Extension MDX files generated in content/docs/ext/"

arm:
	CGO_ENABLED=0 GOOS=linux  GOARCH=arm64 go build -a -ldflags "$(LD_FLAGS) -extldflags '-static'" -o pgext
	upx pgext
amd:
	CGO_ENABLED=0 GOOS=linux  GOARCH=amd64 go build -a -ldflags "$(LD_FLAGS) -extldflags '-static'" -o pgext
	upx pgext

# inventory
.PHONY: default run gen dump save load gen-json gen-mdx build-mdx