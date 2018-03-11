# GoPrez
Application for sending images to other people's screen for presentations

# Required libraries
-- Pixel

`bash
go get github.com/faiface/pixel
`

-- PixelGL

`bash
go get github.com/faiface/pixel/pixelgl
`

# Cross compiling for windows
run

`bash
CGO_ENABLED=1 CC=x86_64-w64-mingw32-gcc CXX=x86_64-w64-mingw32-g++ GOOS=windows GOARCH=amd64 go build
`
