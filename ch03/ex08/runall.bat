cd bigfloat
go run main.go bigcomplex.go > out.png
cd ..
cd complex64
go run main.go > out.png
cd ..
cd complex128
go run main.go > out.png
cd ..
cd rat
go run main.go complex.go > out.png
cd ..