echo "\n"
echo "-------------------------------------------- Validator Set Contract ------------------------------------------"

echo "validator set contract"
echo "\n"

echo "show help \n"
echo " > ../win validatorset -h"
echo "\n"
../win validatorset -h
echo "\n"


echo "show the alreadyInit status \n"
echo " > ../win validatorset -i"
../win validatorset -i
echo "\n"


echo "show the current validator set by index \n"
echo " > ../win validatorset -s 0"
../win validatorset -s 0
echo " > ../win validatorset -s 1"
../win validatorset -s 1
echo "\n"


echo "show the validator set map address => index \n"
echo " > ../win validatorset -m 0x67C7850bD403241fa8661D7B7d57E4CFD593EF92" # validator 4
../win validatorset -m 0x67C7850bD403241fa8661D7B7d57E4CFD593EF92
echo " > ../win validatorset -m 0xcc9963Cb2E935B92f7869bb8f34ae53297F6d2C7" # validator 3
 ../win validatorset -m 0xcc9963Cb2E935B92f7869bb8f34ae53297F6d2C7
 echo " > ../win validatorset -m 0x784098965D74Eaf2aA5d644fFeeb1106443502ee" # validator 5
 ../win validatorset -m 0x784098965D74Eaf2aA5d644fFeeb1106443502ee
echo "\n"

echo "show the next end time \n"
echo " > ../win validatorset -e"
../win validatorset -e
echo "\n"

echo "show number of validators in the active set \n"
echo " > ../win validatorset -n"
../win validatorset -n
echo "\n"

echo "show all active validator \n"
echo " > ../win validatorset -v"
../win validatorset -v
echo "\n"
