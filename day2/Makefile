.PHONY: clean

day: a b
	@./a input2.txt
	@./b input2.txt
	
a: first
	@go build -o a ./first

b: second
	@go build -o b ./second

clean: 
	@$(RM) a b

