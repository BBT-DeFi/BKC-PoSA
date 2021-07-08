mkdir build

solc --abi --bin --overwrite ./Contracts/Test.sol -o build
mkdir abi
abigen --bin=./build/Test.bin --abi=./build/Test.abi --pkg=test --out=abi/Test.go