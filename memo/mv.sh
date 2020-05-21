#https://qiita.com/NoName/items/fc59849ce7d497a7571a

ls *.jpg | awk '{ printf "mv %s source%04d.jpg\n", $0, NR-1 }' | sh

