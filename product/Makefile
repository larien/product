test:
	go test ./... -v -short -p 1 -cover

test-all:
	go test ./... -v -p 1 -cover

mock: # usage: make mock INTERFACE=value
	mockery --name=$(INTERFACE) --structname=Mock --inpackage