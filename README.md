# TLV Decoder
This project presents a TLV decoder.

It recieves a list of bytes and returns a map of string indexed by string.
The map contains the fields indexed by the order in which they were found.

For example entering `A0511AB398765UJ1N230200` will result in a map containing:
```
"0": "A0511AB398765UJ1"
"1": "N230200"
```

## Initializing the project
To initalize go modules run:
```
go mod init falabella-tlv
```

## Runing the project
To start the project run:
```
go run main.go
```
The following list of commands will appear on the console:
```
decode		-> decode entered TLV fields into a map
help		-> shows this list of commands
exit 		-> to exit
```
Entering `decode` will allow the user to enter TLV fields and will show the resulting map.

The `help` command shows the list of commands.

To close the program enter `exit`.

## Run tests
To run tests with verbose output (-v) and code coverage (-cover) use:
```  
go test -v -cover ./tlv-decoder
```
