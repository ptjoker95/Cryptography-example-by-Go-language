#!/bin/bash

word="YAMAMOTO"
encryptedWord="PRDRDFKF"
encrypt=""
i=1

while [ "$(Cryptography-example-by-Go-language 1 $i $word)" != "$encryptedWord" ]
do
	echo "key $i is" $(Cryptography-example-by-Go-language 1 $i $word)
	let "i=i+1"
done

echo "key $i is" $(Cryptography-example-by-Go-language 1 $i $word)

echo "key is $i"
