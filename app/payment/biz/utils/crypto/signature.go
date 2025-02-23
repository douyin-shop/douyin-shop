// @Author Adrian.Wang 2025/2/22 16:07:00
package crypto

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"
	"fmt"
)

// RSASign 使用RSA私钥对消息进行签名
func RSASign(privateKeyPEM []byte, message []byte) ([]byte, error) {
	// 解析PEM格式的私钥
	block, _ := pem.Decode(privateKeyPEM)
	if block == nil {
		return nil, fmt.Errorf("failed to decode PEM block containing private key")
	}

	// 解析私钥
	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	// 计算消息的哈希值
	hashed := sha256.Sum256(message)

	// 使用私钥进行签名
	signature, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, hashed[:])
	if err != nil {
		return nil, err
	}

	return signature, nil
}

// RSAVerify 使用RSA公钥验证签名
func RSAVerify(publicKeyPEM []byte, message, signature []byte) error {
	// 解析PEM格式的公钥
	block, _ := pem.Decode(publicKeyPEM)
	if block == nil {
		return fmt.Errorf("failed to decode PEM block containing public key")
	}

	// 解析公钥
	publicKeyInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return err
	}

	publicKey := publicKeyInterface.(*rsa.PublicKey)

	// 计算消息的哈希值
	hashed := sha256.Sum256(message)

	// 使用公钥验证签名
	return rsa.VerifyPKCS1v15(publicKey, crypto.SHA256, hashed[:], signature)
}
