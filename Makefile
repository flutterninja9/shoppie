kill_services:
	-kill $$(lsof -ti :8000)
	-kill $$(lsof -ti :8080)
	-kill $$(lsof -ti :3000)

run_order_service:
	cd order_service && go run main.go

run_product_service:
	cd product_service && go run main.go

run_user_service:
	cd user_service && go run main.go

run_all: kill_services
	$(MAKE) run_user_service &
	$(MAKE) run_product_service &
	$(MAKE) run_order_service &
	wait

