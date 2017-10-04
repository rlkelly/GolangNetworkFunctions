package main

import (
  "crypto/ecdsa"
 	"crypto/elliptic"
 	"crypto/md5"
 	"crypto/rand"
 	"fmt"
 	"hash"
)

type PrivateKey ecdsa.PrivateKey
type PublicKey ecdsa.PublicKey

func main() {
 	pubkeyCurve := elliptic.P256()
 	privatekey, _ := ecdsa.GenerateKey(pubkeyCurve, rand.Reader)

 	pubkey PublicKey = privatekey.PublicKey

 	fmt.Println("Private Key :")
 	fmt.Printf("%x \n", privatekey.D)

 	fmt.Println("Public Key :")
 	fmt.Printf("%x %x \n", pubkey.X, pubkey.Y)

 	// Sign ecdsa style

 	var h hash.Hash
 	h = md5.New()

 	signhash := h.Sum(nil)

 	r, s, _ := ecdsa.Sign(rand.Reader, privatekey, signhash)

 	signature := r.Bytes()
 	signature = append(signature, s.Bytes()...)

 	fmt.Printf("Signature : %x\n", signature)

 	// Verify
 	verifystatus := ecdsa.Verify(&pubkey, signhash, r, s)
 	fmt.Println(verifystatus) // should be true
}
