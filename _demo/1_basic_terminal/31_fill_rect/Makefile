.PHONY: run debug

# for error (non-zero) exit codes, ensure terminal is reset

run:
	go run main.go || (code=$$?; reset; printf "ERROR: unexpected exit, code: %s\n" $$code)

debug:
	DEBUG=1 go run main.go || (code=$$?; reset; printf "ERROR: unexpected exit, code: %s\n" $$code)

