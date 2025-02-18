gen:
	goctl api go -api ordermanagement.api -dir .
format:
	goctl api format --dir ordermanagement.api -declare
all:
	make gen  && make format