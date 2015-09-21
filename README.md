# プロビジョニングプロファイル

プロビジョニングプロファイルについてのメモ書き

主に [iOS Hacker's Handbook](http://www.amazon.co.jp/dp/1118204123) から

## DeveloperCertificates

base64 でエンコードされた証明書

```
-----BEGIN CERTIFICATE-----

-----END CERTIFICATE-----
```

で囲み `openssl x509 -in /path/to/file -text` で情報を出力可能

## 検証

libmis の MISProvisioningProfileCheckValidity 関数で検証が行なわれる 

1. Applle iPhone Certificate Authority によって発行された証明書である
2. 証明書の名前は Apple iPhone OS Provisioning Profile Signing である
3. 証明書チェーンは3つまで
4. ルート証明書は特定の SHA1 ハッシュ値を持つ
5. プロファイルのバージョンは1である
6. プロファイルにデバイスの UDID が存在するか、ProvisionsAllDevices キーが存在する
7. プロファイルが有効期限内である

[iOS Hacker's Handbook の P.77 から](http://www.amazon.co.jp/dp/1118204123)

## Authority

AppStore で公開されているアプリとそれ以外での証明書の違い

### AppStore

1. iPhone OS Application Signing
2. iPhone Certification Authority
3. Apple Root CA

### それ以外

1. iPhone Developer: 名前 (ID)
2. Apple Worldwide Developer Relations Certification Authority
3. Apple Root CA

`codesign -dvvv path/to/app` で確認できる

[iOS Hacker's Handbook の P.78 から](http://www.amazon.co.jp/dp/1118204123)