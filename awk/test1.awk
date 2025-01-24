BEGIN{printf "Col1\tCol2\tCol3\n"} 
/t/{

print $2"\t"$3"\t"$7

}
END{print "Done"}
