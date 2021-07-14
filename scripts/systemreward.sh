echo "\n"
echo "#-------------------------------------------- System Reward Contract --------------------------------------------"

echo "system reward contract"
echo "\n"

echo "show help \n"
echo " > ../win systemreward -h"
echo "\n"
../win systemreward -h
echo "\n"

echo "show the alreadyInit status \n"
echo " > ../win systemreward -i"
../win systemreward -i
echo "\n"

echo "get balance of the contract \n"
echo " > ../win systemreward -b"
../win systemreward -b
echo "\n"

echo "get reward of an active validator \n"
echo " > ../win systemreward -r 0x67C7850bD403241fa8661D7B7d57E4CFD593EF92" # validator 4
../win systemreward -r 0x67C7850bD403241fa8661D7B7d57E4CFD593EF92
echo " > ../win systemreward -r 0xcc9963Cb2E935B92f7869bb8f34ae53297F6d2C7" # validator 3
../win systemreward -r 0xcc9963Cb2E935B92f7869bb8f34ae53297F6d2C7
echo "\n"

