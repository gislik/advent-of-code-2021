.PHONY: clean

day: a b
	@./a input2.txt
	@./b input2.txt
	
a: first.go
	@go build -o a first.go

b: second.go
	@go build -o b second.go

clean: 
	@$(RM) a b
