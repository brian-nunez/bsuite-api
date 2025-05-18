all: run

run:
	$(shell bash ./docker/run.dev.sh >&2)

stop:
	$(shell bash ./docker/stop.dev.sh >&2)

prod-run:
	$(shell bash ./docker/run.prod.sh >&2)

prod-stop:
	$(shell bash ./docker/stop.prod.sh >&2)

build-binary:
	$(shell bash ./docker/build.binary.sh >&2)
