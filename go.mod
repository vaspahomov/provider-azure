module github.com/crossplane/provider-azure

go 1.13

require (
	github.com/Azure/azure-pipeline-go v0.2.2 // indirect
	github.com/Azure/azure-sdk-for-go v42.3.0+incompatible
	github.com/Azure/azure-storage-blob-go v0.7.0
	github.com/Azure/go-autorest/autorest v0.11.12
	github.com/Azure/go-autorest/autorest/adal v0.9.5
	github.com/Azure/go-autorest/autorest/azure/auth v0.4.0
	github.com/Azure/go-autorest/autorest/date v0.3.0
	github.com/Azure/go-autorest/autorest/to v0.3.0
	github.com/Azure/go-autorest/autorest/validation v0.2.0 // indirect
	github.com/crossplane/crossplane-runtime v0.14.1-0.20210722005935-0b469fcc77cd
	github.com/crossplane/crossplane-tools v0.0.0-20210320162312-1baca298c527
	github.com/google/go-cmp v0.5.5
	github.com/google/uuid v1.1.2
	github.com/mattn/go-ieproxy v0.0.0-20190805055040-f9202b1cfdeb // indirect
	github.com/onsi/gomega v1.13.0
	github.com/pkg/errors v0.9.1
	github.com/satori/go.uuid v1.2.0 // indirect
	gopkg.in/alecthomas/kingpin.v2 v2.2.6
	k8s.io/api v0.21.2
	k8s.io/apimachinery v0.21.2
	k8s.io/client-go v0.21.2
	sigs.k8s.io/controller-runtime v0.9.2
	sigs.k8s.io/controller-tools v0.4.0
)

replace github.com/crossplane/crossplane-runtime => github.com/vaspahomov/crossplane-runtime v0.14.1-0.20210723080850-07a4bc176f2b
