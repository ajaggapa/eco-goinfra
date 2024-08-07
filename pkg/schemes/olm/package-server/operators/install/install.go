package install

import (
	"k8s.io/apimachinery/pkg/runtime"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"

	"github.com/openshift-kni/eco-goinfra/pkg/schemes/olm/package-server/operators"
	operatorsv1 "github.com/openshift-kni/eco-goinfra/pkg/schemes/olm/package-server/operators/v1"
)

// Install registers API groups and adds types to a scheme.
func Install(scheme *runtime.Scheme) {
	utilruntime.Must(operators.AddToScheme(scheme))
	utilruntime.Must(operatorsv1.AddToScheme(scheme))
	utilruntime.Must(scheme.SetVersionPriority(operatorsv1.SchemeGroupVersion))
}
