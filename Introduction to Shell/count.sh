!#/bin/bash
count=$(find . | wc -l)
total=$((count*5))
printf "\t\vTotal files * 5: %d\v\n" "$total"
