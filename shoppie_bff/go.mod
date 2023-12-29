module github.com/flutterninja9/shoppie/shoppie_bff

go 1.21.0

require (
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/flutterninja9/shoppie/user_sdk v0.0.1
	github.com/gofiber/fiber/v2 v2.51.0
	github.com/joho/godotenv v1.5.1
	github.com/sirupsen/logrus v1.9.3
	go.uber.org/dig v1.17.1

)

replace github.com/flutterninja9/shoppie/user_sdk => ../user_sdk

require (
	github.com/andybalholm/brotli v1.0.5 // indirect
	github.com/google/uuid v1.5.0 // indirect
	github.com/klauspost/compress v1.16.7 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/mattn/go-runewidth v0.0.15 // indirect
	github.com/rivo/uniseg v0.2.0 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	github.com/valyala/fasthttp v1.50.0 // indirect
	github.com/valyala/tcplisten v1.0.0 // indirect
	golang.org/x/sys v0.14.0 // indirect
)
