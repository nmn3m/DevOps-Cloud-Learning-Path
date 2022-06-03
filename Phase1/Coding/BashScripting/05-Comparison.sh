#!/bin/sh

# COMPARISON
NUM1=40
NUM2=30

if [ "$NUM1" -gt "$NUM2" ]
then 
	echo "$NUM1 is greater than $NUM2"
else 
	echo "$NUM1 is less than $NUM2"
fi

# NOTES
# val1 -eq val2 Returns true if the values are equal
# val1 -ne val2 Returns true if the values are not equal
# val1 -gt val2 Returns true if val1 is greater than val2
# val1 -ge val2 Returns true if val1 is greater than or equal to val2
# val1 -lt val2 Returns true if val1 is less than val2
# val1 -le val2 Returns true if val1 is less than or equal to val2

