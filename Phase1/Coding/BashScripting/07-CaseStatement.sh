#!/bin/sh

# CASE STATEMENT
read -p "Are you Software Engineer: Y/N " ANSWER
case "$ANSWER" in
	[yY] | [yY][eE][sS])
	echo "So the problem will take 5 minutes to be solved"
	;;
   [nN] | [nN][oO])
     echo "Sorry, it will take from you 1 hour to be solved"
     ;;
   *)
     echo "Please enter y/yes or n/no"
     ;;
esac
