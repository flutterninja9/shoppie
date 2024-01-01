kill_services:
	-kill $$(lsof -ti :8000)
	-kill $$(lsof -ti :8080)
	-kill $$(lsof -ti :3000)
	-kill $$(lsof -ti :9000)

run_order_service:
	cd order_service && go mod tidy && go run main.go

run_product_service:
	cd product_service && go mod tidy  && go run main.go

run_user_service:
	cd user_service && go mod tidy  && go run main.go

run_cart_service:
	cd cart_service && go mod tidy  && go run main.go

run_bff:
	cd shoppie_bff && go mod tidy  && go run main.go

run_all: kill_services
	$(MAKE) run_user_service &
	$(MAKE) run_product_service &
	$(MAKE) run_order_service &
	$(MAKE) run_cart_service &
	$(MAKE) run_bff
	wait

