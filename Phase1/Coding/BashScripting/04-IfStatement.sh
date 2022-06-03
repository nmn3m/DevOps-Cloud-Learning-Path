#!/bin/sh

# IF STATEMENT
NAME0="Noureddine"
if [ "$NAME0" == "Noureddine" ]
then
	echo "Your name is: Noureddine"
fi


# IF-ELSE STATEMENT
NAME1="Ahmed"
if [ "$NAME1" == "Ahmed" ]
then 
	echo "Your name is: Ahmed"
else
	echo "Your name is NOT: Ahmed"
fi


# ELSE-IF (elif)
NAME2="AbdulMoniem"
if [ "$NAME2" == "AbdulMoniem" ]
then 
	echo "Your name is: AbdulMoniem"
elif [ "$NAME2" == "Mahmoud" ]
then 
	echo "Your name is: Mahmoud"
else
	echo "Your name is NOT: AbdulMoniem or Mahmoud"
fi
