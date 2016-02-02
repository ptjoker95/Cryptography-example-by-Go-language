package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
)

func askKey() int {
	var key int;
	fmt.Printf("Input Key: ");
	fmt.Scanf("%d", &key);

	fmt.Printf("You input %d", key)
	fmt.Printf(" as a key\n");

	return key
}

func askString() string {

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Input String for Encrypt: ")
	stringForEncrypt, _ := reader.ReadString('\n')
	fmt.Print("You input \"", strings.Replace(stringForEncrypt,"\n","",-1), "\" as for encrypt.\n")

	return stringForEncrypt
}

func decryptString ( key int, stringForDecrypt string ) string {
	var returnString string;

	for i:=0; i<len( stringForDecrypt ); i++ {
		if ( int(stringForDecrypt[i]) >= 65 && int(stringForDecrypt[i]) <= 90 ) {
			tempValue := int(stringForDecrypt[i]) - 65 - key
			if tempValue < 0 {
				tempValue += 26
			}
			returnString += string( tempValue + 65 )
		} else if ( int(stringForDecrypt[i]) >= 97 && int(stringForDecrypt[i]) <= 122 ) {
			tempValue := int(stringForDecrypt[i]) - 97 - key
			if tempValue < 0 {
				tempValue += 26
			}
			returnString += string( tempValue + 97 )
		} else {
			returnString += string( stringForDecrypt[i] )
		}
	}

	return returnString
}

func encryptString ( key int, stringForEncrypt string ) string {

	var returnString string;

	for i:=0; i<len( stringForEncrypt ); i++ {
		if (int(stringForEncrypt[i]) >= 65 && int(stringForEncrypt[i]) <= 90) {
	//		fmt.Print( string( (int(stringForEncrypt[i]) - 65 + key ) % 26 + 65 ))
			returnString += string( (int(stringForEncrypt[i]) - 65 + key ) % 26 + 65 )

		} else if (int(stringForEncrypt[i]) >= 97 && int(stringForEncrypt[i]) <= 122 ) {
			returnString += string( (int(stringForEncrypt[i]) - 97 + key ) % 26 + 97)
	//		fmt.Print( string( (int(stringForEncrypt[i]) - 97 + key ) % 26 + 97 ))
		} else {
			returnString += string(stringForEncrypt[i])
	//		fmt.Print( string(stringForEncrypt[i]) )
		}
	}

	//fmt.Print( "Return Value is ", returnString, "\n" );

	return returnString
}

func main() {

/*

	argsWithProg := os.Args
	argsWithoutProg := os.Args[1:]
	arg := os.Args[3]

	fmt.Println(argsWithProg)
	fmt.Println(argsWithoutProg)
	fmt.Println(arg)
*/

	args := os.Args

	if len(args) == 1 {

		fmt.Printf("Encrypt(1) or Decrypt(2)? ")

		var EnOrDe int;

		fmt.Scanf("%d", &EnOrDe)


		if EnOrDe == 1 {
			fmt.Printf("This is Encrypt practice about Go\n");

			key := askKey()

			stringForEncrypt := askString()

			fmt.Print( encryptString( key, stringForEncrypt ) )

		} else if EnOrDe == 2 {
			fmt.Printf("This is Decrypt practice about Go\n");

			key := askKey()

			stringForDecrypt := askString()

			fmt.Print( decryptString( key, stringForDecrypt ) )
		}
	} else {
		EnOrDe, err := strconv.Atoi(args[1])

		if err != nil {
			fmt.Print("Arguments: Encrypt(1) or Decrypt(2), key, string\n")
			os.Exit(-1)
		}

		if EnOrDe == 1 {

			key, err := strconv.Atoi(args[2])

			if err != nil {
				fmt.Print("Arguments: Encrypt(1) or Decrypt(2), key, string\n")
				os.Exit(-1)
			}

			fmt.Print( encryptString( key, args[3] ), "\n" )

		} else if EnOrDe == 2 {

			key, err := strconv.Atoi(args[2])

			if err != nil {
				fmt.Print("Arguments: Encrypt(1) or Decrypt(2), key, string\n")
				os.Exit(-1)
			}
			fmt.Print( decryptString( key, args[3] ), "\n" )
		}
	}

//	fmt.Printf("key is %d", key)
//	fmt.Printf("\n")
//	fmt.Printf("string is %s", string(int(stringForEncrypt[0])+key) )
//	fmt.Printf("\n")
/*
	fmt.Print("a is ", int('a'), "\n")
	fmt.Print("A is ", int('A'), "\n")
	fmt.Print("z is ", int('z'), "\n")
	fmt.Print("Z is ", int('Z'), "\n")
	fmt.Print("91 is ", string(91), "\n")
*/

/*
	for i:=0; i<len( stringForEncrypt ); i++ {
		if (int(stringForEncrypt[i]) >= 65 && int(stringForEncrypt[i]) <= 90) {
			fmt.Print( string( (int(stringForEncrypt[i]) - 65 + key ) % 26 + 65 ))

		} else if (int(stringForEncrypt[i]) >= 97 && int(stringForEncrypt[i]) <= 122 ) {
			fmt.Print( string( (int(stringForEncrypt[i]) - 97 + key ) % 26 + 97 ))
		} else {
			fmt.Print( string(stringForEncrypt[i]) )
		}
	}
*/
}
