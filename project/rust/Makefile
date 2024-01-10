Project = demo

.PHONY:

build:
	cargo build
	move target\debug\$(Project).exe $(Project).exe

run:
	cargo run

clean:
	cargo clean
	del .\$(Project).exe