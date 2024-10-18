post-compile: compile-proto golangci-lint-fix

compile-proto: proto
	@for idl in \
		"protocols/arbiter" \
		"protocols/audio" \
		"protocols/event" \
		"protocols/image" \
		"protocols/text" \
	; do \
		GOPATH=$(GOPATH) && \
		PATH=$(GOPATH)/bin:$$PATH && \
		cd $$idl && \
		find . -maxdepth 10 ! -name "prediction.thrift" -type f -exec rm -rf {} \; && \
		kitex -module github.com/ishumei/krpc -service $$idl prediction.thrift && \
		cd -; \
	done

include Makefile.mk
