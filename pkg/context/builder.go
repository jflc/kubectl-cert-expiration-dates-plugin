package context

import (
	"fmt"

	"github.com/jflc/kubectl-cert-expiration-dates-plugin/pkg/util"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	"k8s.io/client-go/tools/clientcmd/api"
)

// CertExpirationDatesContextBuilder : provides information required to CertExpirationDatesContext
type CertExpirationDatesContextBuilder struct {
	configFlags *genericclioptions.ConfigFlags
}

// Builder : provides an instance of CertExpirationDatesContextBuilder
func Builder() CertExpirationDatesContextBuilder {
	return CertExpirationDatesContextBuilder{}
}

// ConfigFlags : set configFlags
func (builder *CertExpirationDatesContextBuilder) ConfigFlags(configFlags *genericclioptions.ConfigFlags) *CertExpirationDatesContextBuilder {
	builder.configFlags = configFlags
	return builder
}

// Build : build CertExpirationDatesContext
func (builder *CertExpirationDatesContextBuilder) Build() (CertExpirationDatesContext, error) {
	var result CertExpirationDatesContext
	var err error
	var rawConfig api.Config
	var contexts map[string]*api.Context

	rawConfig, err = builder.configFlags.ToRawKubeConfigLoader().RawConfig()
	if err != nil {
		return result, err
	}

	ctxFlag := *builder.configFlags.Context
	if len(ctxFlag) > 0 {
		if ctxValue, ok := rawConfig.Contexts[ctxFlag]; ok {
			contexts = map[string]*api.Context{
				ctxFlag: ctxValue,
			}
		} else {
			err = fmt.Errorf("Context was not found for specified context: %v", ctxFlag)
			return result, err
		}
	} else {
		contexts = rawConfig.Contexts
	}

	for k, v := range contexts {
		context := k
		cluster := v.Cluster
		user := v.AuthInfo
		authInfo := rawConfig.AuthInfos[user]

		if err = result.add(context, cluster, user, authInfo); err != nil {
			return result, err
		}
	}

	return result, err
}

func (ctx *CertExpirationDatesContext) add(context string, cluster string, user string, authInfo *api.AuthInfo) error {
	cert, err := buildCertificate(cluster, cluster, user, authInfo)
	if err != nil {
		return err
	}
	ctx.Certificates = append(ctx.Certificates, cert)
	return nil
}

func buildCertificate(context string, cluster string, user string, authInfo *api.AuthInfo) (Certificate, error) {
	var result Certificate
	var err error

	cert, err := util.ParseCertificate(authInfo.ClientCertificateData)
	if err != nil {
		return result, err
	}
	validFrom := cert.NotBefore
	validTo := cert.NotAfter

	result = Certificate{
		Context:   context,
		Cluster:   cluster,
		User:      user,
		ValidFrom: validFrom,
		ValidTo:   validTo,
	}

	return result, err
}
