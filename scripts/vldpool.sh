echo "\n"
echo "#-------------------------------------------- Validator Pool Contract --------------------------------------------"

echo "validator pool contract"
echo "\n"

echo "show help \n"
echo " > ../win vldpool -h"
echo "\n"
../win vldpool -h
echo "\n"

echo "show the alreadyInit status \n"
echo " > ../win vldpool -i"
../win vldpool -i
echo "\n"

echo "show all validators in the pool \n"
echo " > ../win vldpool -v"
../win vldpool -v
echo "\n"

echo "get validator at index (starts at 0) \n"
echo " > ../win vldpool -a 0"
../win vldpool -a 0
echo " > ../win vldpool -a 1"
../win vldpool -a 1
echo "\n"

echo "get number of validators in the pool \n"
echo " > ../win vldpool -n"
../win vldpool -n
echo "\n"

echo "get total power exclude unbonding delegation of a validator \n"
echo " > ../win vldpool -p 0xcc9963Cb2E935B92f7869bb8f34ae53297F6d2C7" # validator 3
../win vldpool -p 0xcc9963Cb2E935B92f7869bb8f34ae53297F6d2C7
echo "\n"

echo "get the jail time of a validator \n"
echo " > ../win vldpool -j 0xcc9963Cb2E935B92f7869bb8f34ae53297F6d2C7" # validator 3 
../win vldpool -j 0xcc9963Cb2E935B92f7869bb8f34ae53297F6d2C7
echo " > ../win vldpool -j 0xB7720B04554b33a57b3beC8717aE2a2Cba128740" # validator 1
../win vldpool -j 0xB7720B04554b33a57b3beC8717aE2a2Cba128740
echo "\n"

echo "get remove time of a validator \n"
echo " > ../win vldpool -r 0xcc9963Cb2E935B92f7869bb8f34ae53297F6d2C7" # validator 3 
../win vldpool -r 0xcc9963Cb2E935B92f7869bb8f34ae53297F6d2C7
echo " > ../win vldpool -r 0xB7720B04554b33a57b3beC8717aE2a2Cba128740" # validator 1
../win vldpool -r 0xB7720B04554b33a57b3beC8717aE2a2Cba128740
echo "\n"

echo "get unbond time of a validator \n"
echo " > ../win vldpool -u 0x1a56A2579fFe9ab0bEA270576053de97F78212B1" # validator 0
../win vldpool -u 0x1a56A2579fFe9ab0bEA270576053de97F78212B1
echo " > ../win vldpool -u 0x784098965D74Eaf2aA5d644fFeeb1106443502ee" # validator 5
../win vldpool -u 0x784098965D74Eaf2aA5d644fFeeb1106443502ee
echo "\n"

echo "get index (starts at 1) of a validator's address \n"
echo " > ../win vldpool -m 0xcc9963Cb2E935B92f7869bb8f34ae53297F6d2C7" # validator 3 
../win vldpool -m 0xcc9963Cb2E935B92f7869bb8f34ae53297F6d2C7
echo " > ../win vldpool -m 0xE1e052d29352505e835E7ABF94a719678c62589E" # not a validator 
../win vldpool -m 0xE1e052d29352505e835E7ABF94a719678c62589E
echo "\n"