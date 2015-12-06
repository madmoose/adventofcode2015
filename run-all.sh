for f in day*
do
	if [[ -x "$f" ]]
	then
		echo -en "$f\t"
		INPUT=${f:0:5}.txt
		if [[ -f $INPUT ]]
		then
			./$f < $INPUT
		else
			./$f
		fi
	fi
done
