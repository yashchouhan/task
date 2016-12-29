$ vi myinfo
#
# Script to print current user information who currently logged IB, current date & time
#
clear
echo "Hello $USER"
echo "Today is \c "; date
echo "Number os user login : \c" ; who | wc -l
echo "Calendar"
cal
exit 0
