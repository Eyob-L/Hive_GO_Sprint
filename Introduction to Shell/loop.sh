!#/bin/bash

number=$1
if [ $number -gt 100 ]; then
	number=100
fi

i=1
while [ $i -le $number ]; do
	echo "This is loop number $i"
	i=$((i + 1))
done
