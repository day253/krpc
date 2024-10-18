post-compile: compile-proto golangci-lint-fix

compile-proto: proto
	@for idl in \
		"arbiter" \
		"audio" \
		"event" \
		"image" \
		"text" \
	; do \
		GOPATH=$(GOPATH) && \
		PATH=$(GOPATH)/bin:$$PATH && \
		cd $$idl && \
		find . -maxdepth 10 ! -name "prediction.thrift" -type f -exec rm -rf {} \; && \
		kitex -module github.com/ishumei/krpc/protocols -service $$idl prediction.thrift && \
		cd -; \
	done

include Makefile.mk
