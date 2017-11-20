list.txt: list1.txt list2.txt list3.txt list4.txt list5.txt list6.txt list7.txt list8.txt list9.txt list10.txt list11.txt list12.txt list13.txt list14.txt list15.txt list16.txt list17.txt list18.txt list19.txt list20.txt
	cp /tmp/go-make/outputfiles/list1.txt /tmp/go-make/outputfiles/list.txt ; cat /tmp/go-make/outputfiles/list2.txt >> /tmp/go-make/outputfiles/list.txt ; cat /tmp/go-make/outputfiles/list3.txt >> /tmp/go-make/outputfiles/list.txt ; cat /tmp/go-make/outputfiles/list4.txt >> /tmp/go-make/outputfiles/list.txt ; cat /tmp/go-make/outputfiles/list5.txt >> /tmp/go-make/outputfiles/list.txt ; cat /tmp/go-make/outputfiles/list6.txt >> /tmp/go-make/outputfiles/list.txt ; cat /tmp/go-make/outputfiles/list7.txt >> /tmp/go-make/outputfiles/list.txt ; cat /tmp/go-make/outputfiles/list8.txt >> /tmp/go-make/outputfiles/list.txt ; cat /tmp/go-make/outputfiles/list9.txt >> /tmp/go-make/outputfiles/list.txt ; cat /tmp/go-make/outputfiles/list10.txt >> /tmp/go-make/outputfiles/list.txt ; cat /tmp/go-make/outputfiles/list11.txt >> /tmp/go-make/outputfiles/list.txt ; cat /tmp/go-make/outputfiles/list12.txt >> /tmp/go-make/outputfiles/list.txt ; cat /tmp/go-make/outputfiles/list13.txt >> /tmp/go-make/outputfiles/list.txt ; cat /tmp/go-make/outputfiles/list14.txt >> /tmp/go-make/outputfiles/list.txt ; cat /tmp/go-make/outputfiles/list15.txt >> /tmp/go-make/outputfiles/list.txt ; cat /tmp/go-make/outputfiles/list16.txt >> /tmp/go-make/outputfiles/list.txt ; cat /tmp/go-make/outputfiles/list17.txt >> /tmp/go-make/outputfiles/list.txt ; cat /tmp/go-make/outputfiles/list18.txt >> /tmp/go-make/outputfiles/list.txt ; cat /tmp/go-make/outputfiles/list19.txt >> /tmp/go-make/outputfiles/list.txt ; cat /tmp/go-make/outputfiles/list20.txt >> /tmp/go-make/outputfiles/list.txt ;

list1.txt: premier
	premier 2 `echo 1*200000000/20-1 |bc` > /tmp/go-make/outputfiles/list1.txt

list2.txt: premier
	premier `echo 1*200000000/20 |bc` `echo 2*200000000/20-1 |bc` > /tmp/go-make/outputfiles/list2.txt

list3.txt: premier
	premier `echo 2*200000000/20 |bc` `echo 3*200000000/20-1 |bc` > /tmp/go-make/outputfiles/list3.txt

list4.txt: premier
	premier `echo 3*200000000/20 |bc` `echo 4*200000000/20-1 |bc` > /tmp/go-make/outputfiles/list4.txt

list5.txt: premier
	premier `echo 4*200000000/20 |bc` `echo 5*200000000/20-1 |bc` > /tmp/go-make/outputfiles/list5.txt

list6.txt: premier
	premier `echo 5*200000000/20 |bc` `echo 6*200000000/20-1 |bc` > /tmp/go-make/outputfiles/list6.txt

list7.txt: premier
	premier `echo 6*200000000/20 |bc` `echo 7*200000000/20-1 |bc` > /tmp/go-make/outputfiles/list7.txt

list8.txt: premier
	premier `echo 7*200000000/20 |bc` `echo 8*200000000/20-1 |bc` > /tmp/go-make/outputfiles/list8.txt

list9.txt: premier
	premier `echo 8*200000000/20 |bc` `echo 9*200000000/20-1 |bc` > /tmp/go-make/outputfiles/list9.txt

list10.txt: premier
	premier `echo 9*200000000/20 |bc` `echo 10*200000000/20-1 |bc` > /tmp/go-make/outputfiles/list10.txt

list11.txt: premier
	premier `echo 10*200000000/20 |bc` `echo 11*200000000/20-1 |bc` > /tmp/go-make/outputfiles/list11.txt

list12.txt: premier
	premier `echo 11*200000000/20 |bc` `echo 12*200000000/20-1 |bc` > /tmp/go-make/outputfiles/list12.txt

list13.txt: premier
	premier `echo 12*200000000/20 |bc` `echo 13*200000000/20-1 |bc` > /tmp/go-make/outputfiles/list13.txt

list14.txt: premier
	premier `echo 13*200000000/20 |bc` `echo 14*200000000/20-1 |bc` > /tmp/go-make/outputfiles/list14.txt

list15.txt: premier
	premier `echo 14*200000000/20 |bc` `echo 15*200000000/20-1 |bc` > /tmp/go-make/outputfiles/list15.txt

list16.txt: premier
	premier `echo 15*200000000/20 |bc` `echo 16*200000000/20-1 |bc` > /tmp/go-make/outputfiles/list16.txt

list17.txt: premier
	premier `echo 16*200000000/20 |bc` `echo 17*200000000/20-1 |bc` > /tmp/go-make/outputfiles/list17.txt

list18.txt: premier
	premier `echo 17*200000000/20 |bc` `echo 18*200000000/20-1 |bc` > /tmp/go-make/outputfiles/list18.txt

list19.txt: premier
	premier `echo 18*200000000/20 |bc` `echo 19*200000000/20-1 |bc` > /tmp/go-make/outputfiles/list19.txt

list20.txt: premier
	premier `echo 19*200000000/20 |bc` `echo 20*200000000/20-1 |bc` > /tmp/go-make/outputfiles/list20.txt

premier:
	gcc -o /tmp/go-make/outputfiles/premier premier.c -lm
