all:
	@cat Makefile

run: format
	sudo `pwd`/caddy run -c `pwd`/Caddyfile

reload: format
	sudo `pwd`/caddy reload -c `pwd`/Caddyfile

format:
	./caddy fmt --overwrite ./Caddyfile

