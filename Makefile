
dao:
	gf gen dao
service:
	gf gen service
ctrl:
	gf gen ctrl -s app/http/api -d app/http/internal/controller
enums:
	$(eval _DIR  = $(shell pwd))
	gf gen enums -s ./internal/consts/ -p $(_DIR)/internal/boot/boot_enums.go