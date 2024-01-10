

# 尝试让我们的makefile变成可配置的版本 


# need some configuration right here

source_extension := .go
CC = go
CFLAGS = 
#是否需要预处理不是吗


pwd := $(shell pwd)
project := $(notdir $(pwd))
folders := $(wildcard */)

SRCS := $(shell find . -name *$(source_extension))
# OBJS := $(patsubst %.c,%.o,$(SRCS))


build: $(SRCS)
	$(CC) $(CFLAGS) -o $(project) $^

run: build
	./$(project)

rm:
	rm -rf $(project)

clean: rm


all-argv: pwd project folders  src 

pwd:
	@echo pwd : $(pwd)
folders :
	@echo folders : $(folders)

project:
	@echo project : $(project)

src:
	@echo srcs : $(SRCS)
