package main

import (
	"bufio"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha3"
	"crypto/sha512"
	"fmt"
	"my-go-project/hash_lab/utils/crack"
	"os"
)

// Lab 0
func lab0(input1 string, input2 string){
	input1_byte := []byte(input1)
	input2_byte := []byte(input2)

	input1_md5 := md5.Sum(input1_byte)
	input2_md5 := md5.Sum(input2_byte)

	if(input1_md5 == input2_md5){
		fmt.Printf("Hash MD5 : Match\nInput1 : %x\nInput2 : %x\n\n", input1_md5, input2_md5)
	}else{
		fmt.Printf("Hash MD5 : Not Match\nInput1 : %x\nInput2 : %x\n\n", input1_md5, input2_md5)
	}

	input1_sha1 := sha1.Sum(input1_byte)
	input2_sha1 := sha1.Sum(input2_byte)

	if(input1_md5 == input2_md5){
		fmt.Printf("Hash SHA1 : Match\nInput1 : %x\nInput2 : %x\n\n", input1_sha1, input2_sha1)
	}else{
		fmt.Printf("Hash SHA1 : Not Match\nInput1 : %x\nInput2 : %x\n\n", input1_sha1, input2_sha1)
	}

	input1_sha256 := sha256.Sum256(input1_byte)
	input2_sha256 := sha256.Sum256(input1_byte)

	if(input1_sha256 == input2_sha256){
		fmt.Printf("Hash SHA256 : Match\nInput1 : %x\nInput2 : %x\n\n", input1_sha256, input2_sha256)
	}else{
		fmt.Printf("Hash SHA256 : Not Match\nInput1 : %x\nInput2 : %x\n\n", input1_sha256, input2_sha256)
	}

	input1_sha3_256 := sha3.Sum256(input1_byte)
	input2_sha3_256 := sha3.Sum256(input2_byte)

	if(input1_sha256 == input2_sha256){
		fmt.Printf("Hash SHA3-256 : Match\nInput1 : %x\nInput2 : %x\n\n", input1_sha3_256, input2_sha3_256)
	}else{
		fmt.Printf("Hash SHA3-256 : Not Match\nInput1 : %x\nInput2 : %x\n\n", input1_sha3_256, input2_sha3_256)
	}

	input1_sha512 := sha512.Sum512(input1_byte)
	input2_sha512 := sha512.Sum512(input2_byte)

	if(input1_sha256 == input2_sha256){
		fmt.Printf("Hash SHA-512 : Match\nInput1 : %x\nInput2 : %x\n\n", input1_sha512, input2_sha512)
	}else{
		fmt.Printf("Hash SHA-512 : Not Match\nInput1 : %x\nInput2 : %x\n\n", input1_sha512, input2_sha512)
	}
}

// Lab 1
const MD5Hash = "6a85dfd77d9cb35770c9dc6728d73d3f"
func CrackMD5() {
    file, err := os.Open("hash_lab/nord_vpn.txt")
    if (err != nil){
        fmt.Println("Error:", err)
        return
    }
    defer file.Close()

	output, err := os.OpenFile("hash_lab/verbose.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if (err != nil){
        fmt.Println("Error:", err)
        return
    }
    defer output.Close()

    scanner := bufio.NewScanner(file)

    for scanner.Scan() {
        password := scanner.Text()
        hash := crack.MD5Hash(password)

        fmt.Println("Trying:", password, "MD5:", hash)

        if (hash == MD5Hash){
            fmt.Println("Password found:", password)

			fmt.Fprintln(output, "Target hash :", MD5Hash)
			fmt.Fprintln(output, "Password MD5 found : ", password)
			fmt.Fprintln(output)
            break
        }
    }

	err = scanner.Err()
	if (err != nil){
        fmt.Println("Reading error:", err)
        return
    }

}

// Lab 2
const SHA1Hash = "aa1c7d931cf140bb35a5a16adeb83a551649c3b9"
func CrackSha1(){
	file, err := os.Open("hash_lab/nord_vpn.txt")
	if (err != nil){
		fmt.Println("Error :", err)
		return
	}
	defer file.Close()

	output, err := os.OpenFile("hash_lab/verbose.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if(err != nil){
		fmt.Println("Error : ", err)
		return
	}
	defer output.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan(){
		password := scanner.Text()
		hash := crack.SHA1HASH(password)

		fmt.Println("Trying :", password, "SHA1 :", hash)

		if hash == SHA1Hash{
			fmt.Println("Password found:" ,password)

			fmt.Fprintln(output, "Target hash :", SHA1Hash)
			fmt.Fprintln(output, "Password SHA1 found :", password)
			fmt.Fprintln(output)
			break
		}
	}
	err = scanner.Err()
	if (err != nil){
    	fmt.Println("Reading error:", err)
        return
    }

}

// Lab3
const SHA512Hash = "485f5c36c6f8474f53a3b0e361369ee3e32ee0de2f368b87b847dd23cb284b316bb0f026ada27df76c31ae8fa8696708d14b4d8fa352dbd8a31991b90ca5dd38"
func CrackSha512(){
	file, err := os.Open("hash_lab/nord_vpn.txt")
	if(err != nil){
		fmt.Println("Error :", err)
		return
	}
	defer file.Close()

	output, err := os.OpenFile("hash_lab/verbose.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if(err != nil){
		fmt.Println("Error :", err)
		return
	}
	defer output.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan(){
		password := scanner.Text()
		hash := crack.SHA512HASH(password)

		fmt.Println("Trying :", password, "SHA512 :",hash)

		if(hash == SHA512Hash){
			fmt.Println("Password SHA512 found :", password)

			fmt.Fprintln(output, "Target hash :", SHA512Hash)
			fmt.Fprintln(output, "Password SHA512 found :", password)
			fmt.Fprintln(output)
			return
		}
	}
	err = scanner.Err()
	if(err != nil){
		fmt.Println("Reading Error :", err)
	}
}

func main(){
	// var input1 string
	// var input2 string

	// fmt.Print("Please input value 1 : ")
	// fmt.Scanln(&input1)

	// fmt.Print("Please input value 2 : ")
	// fmt.Scanln(&input2)

	// fmt.Println()

	// lab0(input1,input2)

	// CrackMD5()

	// CrackSha1()

	CrackSha512()
}
