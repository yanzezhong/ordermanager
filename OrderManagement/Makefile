gen:
	goctl api go -api ordermanagement.api -dir .
format:
	goctl api format --dir ordermanagement.api -declare
genWindows:
     goctl api go --api /home/yanzezhong/code/ordermanager/OrderManagement/ordermanagement.api --dir .	
formatWindows:	 
     goctl api format --dir /home/yanzezhong/code/ordermanager/OrderManagement/ordermanagement.api -declare
all:
	make gen  && make format
