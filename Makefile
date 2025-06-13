tail-build:
	./tailwindcss --input assets/public/style/css/input.css --output assets/public/style/css/output.css --minify
	go run main.go
