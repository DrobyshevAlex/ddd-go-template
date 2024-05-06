protoc:	
	protoc -I./proto \
   	--go_out=./src/internal/gen \
   	--go-grpc_out=./src/internal/gen \
   	./proto/*.proto

gen:
	mkdir -p ./src/internal/gen
	make protoc

run:
	cd src && go run cmd/app/main.go

lint:
	go-arch-lint check --project-path src

test:
	cd src && go test -v ./... -tags=unit -count=1

e2e:
	cd src && go test --tags=e2e ./test/tests/...

e2e_report:
#	mkdir -p src/test/tests/allure-results/history
#	cp allure-report/history/* src/test/tests/allure-results/history || true
	allure generate src/test/tests/allure-results --clean
	allure open allure-report

e2e_clean:
	rm src/test/tests/allure-results/*.json
