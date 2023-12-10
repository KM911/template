
# need some configuration right here

source_extension := c
object_extension := o
build_target_folder := ./

# compiler and flags
CC = gcc
CFLAGS = -Wall -g


pwd := $(shell pwd)
project := $(notdir $(pwd))
folders := $(wildcard */)

SRCS := $(wildcard *.$(source_extension))
SRCS += $(foreach dir,$(folders),$(wildcard $(dir)*.$(source_extension)))

OBJS := $(patsubst %.$(source_extension),%.$(object_extension),$(SRCS))




run: build
	./$(project)


build: $(OBJS)
	$(CC) $(CFLAGS) -o $(project) $^

# 隐含规则 自动推导出来的
# %.o:%.c
# 	$(CC) $(CFLAGS) -c $< -o $@

rm:
	-rm -rf $(OBJS)
	-rm -rf $(project)
	-rm -rf $(project).exe

clean: rm




## check argv 

source_extension:
	@echo $(source_extension)

object_extension:
	@echo $(object_extension)

pwd:
	@echo $(pwd)
dir:
	@echo $(dir)

project:
	@echo $(project)

folders:
	@echo $(folders)

src:
	@echo $(SRCS)

obj:
	@echo $(OBJS)
