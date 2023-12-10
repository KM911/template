
# need some configuration right here

source_extension := .c
CC = gcc
# CFLAGS = -Wall -g



pwd := $(shell pwd)
project := $(notdir $(pwd))
folders := $(wildcard */)
SRCS := $(foreach dir,$(folders),$(wildcard $(dir)*$(source_extension)))
# OBJS := $(patsubst %.c,%.o,$(SRCS))

# gcc  -o p c/main.c
build: $(SRCS)
	$(CC) $(CFLAGS) -o $(project) $^

rm:
	rm -rf $(project)

clean: rm



pwd:
	@echo $(pwd)

project:
	@echo $(project)

src:
	@echo $(SRCS)
