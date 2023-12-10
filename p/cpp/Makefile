pwd := $(shell pwd)
project := $(notdir $(pwd))
SRCS := $(wildcard *.cpp)
folders := $(wildcard */)
SRCS += $(foreach dir,$(folders),$(wildcard $(dir)*.cpp))
OBJS := $(patsubst %.cpp,%.o,$(SRCS))


CC = g++
CFLAGS = -Wall -g

# 判段当前是否是windows系统


ifeq ($(OS),Windows_NT)
	project := $(project).exe
endif


# 通用的C语言模板

run:$(project)
	./$(project)
$(project):$(OBJS)
	@$(CC) $(CFLAGS)  $(OBJS) -L .  -o $(project)

build:
	@$(CC) main.c -o $(project)

clean:
	-rm *.exe
	-rm *.o
	-rm *.out


# 测试我们的参数

pwd:
	@echo $(pwd)

project:
	@echo $(project)

src:
	@echo $(SRCS)

obj:
	@echo $(OBJS)

argv:
	@echo $@

folder:
	@echo $(folders)