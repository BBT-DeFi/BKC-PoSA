mkdir build_ERC

solc --abi --bin --overwrite @openzeppelin/=$(pwd)/node_modules/@openzeppelin/ ./Contracts/KBD.sol -o build_ERC
mkdir abi_ERC
abigen --bin=./build_ERC/KBD.bin --abi=./build_ERC/KBD.abi --pkg=kbd --out=abi_ERC/KBD.go