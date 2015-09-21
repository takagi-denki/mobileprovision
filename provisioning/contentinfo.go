package provisioning

import (
	"encoding/asn1"
)

type ContentInfo struct {
	ContentType asn1.ObjectIdentifier
	Content SignedData `asn1:"tag:0"`
}


type SignedData struct {
	Sequence Sequence
}

type Sequence struct {
	Version int
    DigestAlgorithms      interface{}//DigestAlgorithmIdentifiers,
    EncapContentInfo      EncapsulatedContentInfo
    Certificates          interface{}//[0] IMPLICIT CertificateSet OPTIONAL,
    Crls                  interface{}//[1] IMPLICIT RevocationInfoChoices OPTIONAL,
    SignerInfos           interface{}//SignerInfos 
}

type SignerInfo struct {
	Version int
}

type EncapsulatedContentInfo struct {
	ContentType asn1.ObjectIdentifier
	Content struct {
		Content []byte
	} `asn1:"tag:0"`
}

func (this *ContentInfo)GetContent() *MobileProvision {
	return NewMobileProvision(this.Content.Sequence.EncapContentInfo.Content.Content)
}

func NewContentInfo(buffer []byte) *ContentInfo {
	var info ContentInfo
	asn1.Unmarshal(buffer, &info)
	
	return &info
}
