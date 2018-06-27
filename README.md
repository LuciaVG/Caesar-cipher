# Caesar-cipher
Programming Languages project. This in an implementation of the Innovative Enhancement Of The Caesar Cipher
Algorithm For Cryptography the full paper can be found [here](https://ieeexplore.ieee.org/stamp/stamp.jsp?arnumber=7749010 ).

## Set-Up
To run, you must previously have Golang [installed](https://golang.org/doc/install).
Then, clone the repo.
```bash
$ git clone https://github.com/LuciaVG/Caesar-cipher.git
```
Now, change to that directory and run the program.
```bash
$ go run caesar.go inputs.json
```
The argument passed should be the path to the JSON file containg the input for the program.

## JSON input
This is an example of how the input file should look like.
```
{
  "inputs": [
    {
      "type" : "en",
      "in_file" : "file.txt",
      "out_file" : "out.txt"
    },
    {
      "type" : "de",
      "in_file" : "file2.txt",
      "out_file" : "out2.txt"
    }
  ]
}
```
The field **type** should be "en" for encription and "de" for decription. The **in_file** field should be a string with the path to the file with the message that is going to be encrypted/decrypted and the **out_field** the one where the encrypted/decrypted should be. If the **out_field** file does not exist it will be created, if it does, it will be overwritten.
