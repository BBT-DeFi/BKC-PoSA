echo "\n"
echo "#-------------------------------------------- Stake Pool Contract --------------------------------------------"

echo "stake pool contract"
echo "\n"

echo "show help \n"
echo " > ../win stakepool -h"
echo "\n"
../win stakepool -h
echo "\n"

echo "show the alreadyInit status \n"
echo " > ../win stakepool -i"
../win stakepool -i
echo "\n"


echo "get all delegators of a validator \n"
echo " > ../win stakepool -d 0xB7720B04554b33a57b3beC8717aE2a2Cba128740" # validator 1
../win stakepool -d 0xB7720B04554b33a57b3beC8717aE2a2Cba128740
echo "\n"

echo "get total delegation exclude unbonding \n" 
echo " > ../win stakepool -t 0xB7720B04554b33a57b3beC8717aE2a2Cba128740 -e " # validator 1
../win stakepool -t 0xB7720B04554b33a57b3beC8717aE2a2Cba128740 -e 
echo "\n"

echo "get bonded amount of delegator delegated to a validator \n"
echo " > ../win stakepool -b 0xB7720B04554b33a57b3beC8717aE2a2Cba128740,0x71C507Fc560fE23b7c29F6bCd91113c222C39dfA" # validator 1 delegator 2
../win stakepool -b 0xB7720B04554b33a57b3beC8717aE2a2Cba128740,0x71C507Fc560fE23b7c29F6bCd91113c222C39dfA
echo " > ../win stakepool -b 0xcc9963Cb2E935B92f7869bb8f34ae53297F6d2C7,0x3C9A841aBfCE00c11b2989Afa6191f221b4055d4" # validator 3 delegator 1
../win stakepool -b 0xcc9963Cb2E935B92f7869bb8f34ae53297F6d2C7,0x3C9A841aBfCE00c11b2989Afa6191f221b4055d4
echo " > ../win stakepool -b 0xcc9963Cb2E935B92f7869bb8f34ae53297F6d2C7,0x71C507Fc560fE23b7c29F6bCd91113c222C39dfA" # validator 3 delegator 2
../win stakepool -b 0xcc9963Cb2E935B92f7869bb8f34ae53297F6d2C7,0x71C507Fc560fE23b7c29F6bCd91113c222C39dfA
echo "\n"

echo "get unbonding amount of delegator \n"
echo " > ../win stakepool -u 0xB7720B04554b33a57b3beC8717aE2a2Cba128740,0x71C507Fc560fE23b7c29F6bCd91113c222C39dfA" # validator 1 delegator 2
../win stakepool -u 0xB7720B04554b33a57b3beC8717aE2a2Cba128740,0x71C507Fc560fE23b7c29F6bCd91113c222C39dfA
echo " > ../win stakepool -b 0xcc9963Cb2E935B92f7869bb8f34ae53297F6d2C7,0x3C9A841aBfCE00c11b2989Afa6191f221b4055d4" # validator 3 delegator 1
../win stakepool -u 0xcc9963Cb2E935B92f7869bb8f34ae53297F6d2C7,0x3C9A841aBfCE00c11b2989Afa6191f221b4055d4
echo " > ../win stakepool -b 0xcc9963Cb2E935B92f7869bb8f34ae53297F6d2C7,0x71C507Fc560fE23b7c29F6bCd91113c222C39dfA" # validator 3 delegator 2
../win stakepool -u 0xcc9963Cb2E935B92f7869bb8f34ae53297F6d2C7,0x71C507Fc560fE23b7c29F6bCd91113c222C39dfA
echo "\n"

echo "get delegator total delegation of a validator \n"
echo " > ../win stakepool -v 0xcc9963Cb2E935B92f7869bb8f34ae53297F6d2C7,0x71C507Fc560fE23b7c29F6bCd91113c222C39dfA"
../win stakepool -v 0xcc9963Cb2E935B92f7869bb8f34ae53297F6d2C7,0x71C507Fc560fE23b7c29F6bCd91113c222C39dfA
echo "\n"

echo "get delegator unbond queue \n"
echo " > ../win stakepool -q 0x71C507Fc560fE23b7c29F6bCd91113c222C39dfA" # delegator 2
../win stakepool -q 0x71C507Fc560fE23b7c29F6bCd91113c222C39dfA
echo "\n"

echo "get total delegation of a validator\n"
echo " > ../win stakepool -t 0xcc9963Cb2E935B92f7869bb8f34ae53297F6d2C7" # validator 3
../win stakepool -t 0xcc9963Cb2E935B92f7869bb8f34ae53297F6d2C7
echo "\n"
