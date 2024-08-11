---
{
  "title": "crypto",
  "tags": [ "crypto", "ras", "hmac" ],
  "layout": "blog/blog.base.gohtml",
  "cTime": "2024-08-11 14:33",
  "mTime": "2024-08-11 14:33"
}
---

# crypto

主要需要提供功能:

- 加簽: 需要確保加簽者唯一，也就是不能有多種方法其加簽出來的結果都指向某位使用者，這樣驗證就沒有意義了

  > 為了要能將加簽的物件，生成出一個唯一的識別碼，常用了話會使用SHA256, 512, ...的方式來生成

- 驗證:
  - 物件,key都相同，能驗證出加簽的對象
  - 物件相同但key不正確驗證需失敗

## HMAC (Keyed-Hash Message Authentication Code)

對稱式加密: 加簽與驗證都用同樣的密鑰

```go
key := []byte("...private_key...")
hasher := hmac.New(
		crypto.SHA256.New, // 雜湊值算法
		key,
)
// Sign
const signingString = "hello" // 假設這個是被加簽的內容
hasher.Write([]byte(signingString))
mac1 = hasher.Sum(nil)

// Verify
hasher2 := hmac.New(
  crypto.SHA256.New, // 生成出物件的指紋
  key, // 如果密鑰不同，則驗證要失敗
) // 要想辦法得知加簽時所用的內容和密鑰
hasher.Write([]byte(signingString))
mac2 = hasher.Sum(nil)

if hmac.Equal(mac1, mac2) {
    // Valid
}
```

## RSA

RSA是由羅納德·李維斯特（`R`on Rivest）、阿迪·薩莫爾（Adi `S`hamir）和倫納德·阿德曼（Leonard `A`dleman）在1977年一起提出的。

非對稱式加密: 加簽用私鑰; 驗證用公鑰

```go
// 生成私鑰
var rsaKey *rsa.PrivateKey
rsaKey, _ = rsa.GenerateKey(rand.Reader, 2048) // 2048, 3072, 4096

// Privacy-Enhanced Mail (PEM)
// 保存私鑰
_ = pem.Encode(writer, &pem.Block{
      Type:    "RSA PRIVATE KEY",
      Headers: map[string]string{},
      Bytes:   x509.MarshalPKCS1PrivateKey(rsaKey),
})

// 保存公鑰
_ = pem.Encode(writer2, &pem.Block{
			Type:    "RSA PUBLIC KEY",
			Headers: map[string]string{},
			Bytes:   x509.MarshalPKCS1PublicKey(&rsaKey.PublicKey),
})

// Sign
// 將我們要簽署的內容計算其SHA256, 生成出此物件的指紋
hash := crypto.SHA256
hasher := hash.New()
hasher.Write([]byte(signingString))
fingerprint := hasher.Sum(nil)
signedBytes, _ = rsa.SignPKCS1v15( // PKCS#1 是RSA實驗室提出的一系列標準, v1.5是其中的一個版本
  rand.Reader,
  rsaKey,
  hash,
  fingerprint
)

// Verify
var rsaPublicKey *rsa.PublicKey // 要想辦法得知公鑰，通常SERVER會告知您公鑰可以上拿取得
hash := crypto.SHA256 // 這邊一樣要想辦法得知加簽時所用的指紋演算法
hasher = hash.New()
hasher.Write([]byte(signingString))
singingBytes := hasher.Sum(nil)
if err := rsa.VerifyPKCS1v15(
  // 有以下四項即可進行驗證
  // {公鑰 + 指紋的演算法(sha) + 本次物件(如果沒辦串改就和以前加簽的物件相同) + 之前透過私鑰加簽出來的結果}
  rsaPublicKey, hash, singingBytes, signedBytes,
); err == nil {
  // Valid
}
```

rsa.SignPKCS1v15其實裡面還有做一些填充的方法，來將資料轉換成適合的長度和格式。例如:

- 格式: 固定前面多少byte是xxx
- 長度: 填充多少垃圾來達到滿長度

[playground](https://go.dev/play/p/W1vH5B7eL5z)
