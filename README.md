### Email address verification utility ###

The package is example of batch email verification via SMTP lookup. 
It reads line by line email addresses from input file, verifies and write result to new csv file. 


```shell
./verimail -i=input.txt -w=2 -c=3
```

`-i` - input file with emails (every address at the new line) 
`-w` - max amount of workers 
`-c` - max amount of connections 