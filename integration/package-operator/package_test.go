//go:build integration

package packageoperator

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"k8s.io/utils/ptr"

	"github.com/go-logr/logr"
	"github.com/go-logr/logr/testr"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	appsv1 "k8s.io/api/apps/v1"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"pkg.package-operator.run/cardboard/kubeutils/wait"
	"sigs.k8s.io/controller-runtime/pkg/client"

	corev1alpha1 "package-operator.run/apis/core/v1alpha1"
	manifestsv1alpha1 "package-operator.run/apis/manifests/v1alpha1"
)

func requireDeployPackage(ctx context.Context, t *testing.T, pkg, objectDeployment client.Object) {
	t.Helper()

	require.NoError(t, Client.Create(ctx, pkg))
	cleanupOnSuccess(ctx, t, pkg)

	timeoutOpt := wait.WithTimeout(40 * time.Second)

	require.NoError(t,
		Waiter.WaitForCondition(ctx, pkg, corev1alpha1.PackageUnpacked, metav1.ConditionTrue, timeoutOpt))
	// Condition Mapping from Deployment
	require.NoError(t,
		Waiter.WaitForCondition(ctx, pkg, "my-prefix/Progressing", metav1.ConditionTrue, timeoutOpt))
	require.NoError(t,
		Waiter.WaitForCondition(ctx, pkg, corev1alpha1.PackageAvailable, metav1.ConditionTrue, timeoutOpt))

	require.NoError(t, Client.Get(ctx, client.ObjectKey{
		Name: pkg.GetName(), Namespace: pkg.GetNamespace(),
	}, objectDeployment))
}

func newPackage(meta metav1.ObjectMeta, spec corev1alpha1.PackageSpec, namespaced bool) client.Object {
	if !namespaced {
		return &corev1alpha1.ClusterPackage{
			ObjectMeta: meta,
			Spec:       spec,
		}
	}

	pkg := &corev1alpha1.Package{
		ObjectMeta: meta,
		Spec:       spec,
	}
	pkg.SetNamespace("default")
	return pkg
}

// testNamespacedAndCluster constructs a (Cluster)Package from the 'meta' and 'spec' parameters
// adding a namespace, if needed. Then it ensures successful deployment of both versions of the package
// and optionally runs 'postCheck'.
func testNamespacedAndCluster(
	t *testing.T,
	meta metav1.ObjectMeta, spec corev1alpha1.PackageSpec,
	postCheck func(ctx context.Context, t *testing.T, namespace string),
) {
	t.Helper()

	for _, tc := range []struct {
		name             string
		namespace        string
		pkg              client.Object
		objectDeployment client.Object
	}{
		{"cluster", "", newPackage(meta, spec, false), &corev1alpha1.ClusterObjectDeployment{}},
		{"namespaced", "default", newPackage(meta, spec, true), &corev1alpha1.ObjectDeployment{}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			ctx := logr.NewContext(context.Background(), testr.New(t))

			requireDeployPackage(ctx, t, tc.pkg, tc.objectDeployment)

			if postCheck != nil {
				postCheck(ctx, t, tc.namespace)
			}
		})
	}
}

func TestPackage_simple(t *testing.T) {
	meta := metav1.ObjectMeta{
		Name: "success",
	}
	spec := corev1alpha1.PackageSpec{
		Image: SuccessTestPackageImage,
		Config: &runtime.RawExtension{
			Raw: []byte(fmt.Sprintf(`{"testStubImage": "%s"}`, TestStubImage)),
		},
	}
	postCheck := func(ctx context.Context, t *testing.T, namespace string) {
		t.Helper()

		if namespace == "" {
			return
		}

		// Test if environment information is injected successfully.
		deploy := &appsv1.Deployment{}
		err := Client.Get(ctx, client.ObjectKey{
			Name:      "test-stub-success",
			Namespace: namespace,
		}, deploy)
		require.NoError(t, err)

		var env manifestsv1alpha1.PackageEnvironment
		te := deploy.Annotations["test-environment"]
		err = json.Unmarshal([]byte(te), &env)
		require.NoError(t, err)
		assert.NotEmpty(t, env.Kubernetes.Version)
	}

	testNamespacedAndCluster(t, meta, spec, postCheck)
}

func TestPackage_simpleWithSlices(t *testing.T) {
	meta := metav1.ObjectMeta{
		Name: "success-slices",
		Annotations: map[string]string{
			// Manually force one slice per object.
			"packages.package-operator.run/chunking-strategy": "EachObject",
		},
	}
	spec := corev1alpha1.PackageSpec{
		Image: SuccessTestPackageImage,
		Config: &runtime.RawExtension{
			Raw: []byte(fmt.Sprintf(`{"testStubImage": "%s"}`, TestStubImage)),
		},
	}

	postCheck := func(ctx context.Context, t *testing.T, namespace string) {
		t.Helper()

		// Reminder: Every rendered object get's wrapped into its own ObjectSlice.
		// Or in go terms: `len(renderedObjects) == len(resultingObjectSlices)`.
		// The test package renders a deployment plus an additional namespace object
		// when the package is installed as a ClusterPackage.
		// When `namespace` is not empty, the test package has been installed as a Package.
		// When `namespace` is empty, the test package has been installed as a ClusterPackage.
		assertedAmount := 2 // Deployment and Namespace
		if namespace != "" {
			assertedAmount = 1 // only a Deployment
		}

		assertAmountOfSliceObjectsControlledPerObjectSet(ctx, t, types.NamespacedName{
			Namespace: namespace,
			Name:      "success-slices",
		}, assertedAmount)
	}

	testNamespacedAndCluster(t, meta, spec, postCheck)
}

func TestPackage_simpleWithoutSlices(t *testing.T) {
	meta := metav1.ObjectMeta{
		Name: "success-no-slices",
		Annotations: map[string]string{
			// Manually disable slicing.
			"packages.package-operator.run/chunking-strategy": "NoOp",
		},
	}
	spec := corev1alpha1.PackageSpec{
		Image: SuccessTestPackageImage,
		Config: &runtime.RawExtension{
			Raw: []byte(fmt.Sprintf(`{"testStubImage": "%s"}`, TestStubImage)),
		},
	}

	postCheck := func(ctx context.Context, t *testing.T, namespace string) {
		t.Helper()
		assertAmountOfSliceObjectsControlledPerObjectSet(ctx, t, types.NamespacedName{
			Namespace: namespace,
			Name:      "success-no-slices",
		}, 0)
	}

	testNamespacedAndCluster(t, meta, spec, postCheck)
}

func TestPackage_multi(t *testing.T) {
	meta := metav1.ObjectMeta{
		Name:      "success-multi",
		Namespace: "default",
	}
	spec := corev1alpha1.PackageSpec{
		Image: SuccessTestMultiPackageImage,
		Config: &runtime.RawExtension{
			Raw: []byte(fmt.Sprintf(`{"testStubMultiPackageImage": "%s","testStubImage": "%s"}`,
				SuccessTestMultiPackageImage, TestStubImage,
			)),
		},
	}

	postCheck := func(ctx context.Context, t *testing.T, namespace string) {
		t.Helper()

		var pkgBE, pkgFE client.Object
		var ns string
		if namespace != "" {
			pkgBE = &corev1alpha1.Package{}
			pkgFE = &corev1alpha1.Package{}
			ns = namespace
		} else {
			pkgBE = &corev1alpha1.ClusterPackage{}
			pkgFE = &corev1alpha1.ClusterPackage{}
			ns = "success-multi"
		}

		require.NoError(t, Client.Get(ctx, client.ObjectKey{Name: "success-multi-backend", Namespace: ns}, pkgBE))
		require.NoError(t, Client.Get(ctx, client.ObjectKey{Name: "success-multi-frontend", Namespace: ns}, pkgFE))
	}

	testNamespacedAndCluster(t, meta, spec, postCheck)
}

func TestPackage_cel(t *testing.T) {
	meta := metav1.ObjectMeta{
		Name:      "success-cel",
		Namespace: "default",
	}
	spec := corev1alpha1.PackageSpec{
		Image: SuccessTestCelPackageImage,
		Config: &runtime.RawExtension{
			Raw: []byte(fmt.Sprintf(`{"testStubCelPackageImage": "%s","testStubImage": "%s"}`,
				SuccessTestCelPackageImage, TestStubImage,
			)),
		},
	}

	postCheck := func(ctx context.Context, t *testing.T, namespace string) {
		t.Helper()

		var ns string
		if namespace != "" {
			ns = namespace
		} else {
			ns = "success-cel"
		}

		// deployment should be there
		deploy := &appsv1.Deployment{}
		err := Client.Get(ctx, client.ObjectKey{
			Name:      "test-deployment",
			Namespace: ns,
		}, deploy)
		require.NoError(t, err)

		// configMap should not be there
		cm := &v1.ConfigMap{}
		err = Client.Get(ctx, client.ObjectKey{
			Name:      "test-cm",
			Namespace: ns,
		}, cm)
		require.EqualError(t, err, "configmaps \"test-cm\" not found")

		// ignored configMap should not be there
		err = Client.Get(ctx, client.ObjectKey{
			Name:      "ignored-cm",
			Namespace: ns,
		}, cm)
		require.EqualError(t, err, "configmaps \"ignored-cm\" not found")

		// check that "cel-template-cm" was templated correctly
		celTemplateCm := &v1.ConfigMap{}
		err = Client.Get(ctx, client.ObjectKey{
			Name:      "cel-template-cm",
			Namespace: ns,
		}, celTemplateCm)
		require.NoError(t, err)

		v, ok := celTemplateCm.Data["banana"]
		require.True(t, ok)
		assert.Equal(t, "bread", v)

		_, ok = celTemplateCm.Data["should-not"]
		assert.False(t, ok)
	}

	testNamespacedAndCluster(t, meta, spec, postCheck)
}

func TestPackage_nonExistent(t *testing.T) {
	pkg := &corev1alpha1.Package{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "non-existent",
			Namespace: "default",
		},
		Spec: corev1alpha1.PackageSpec{
			Image: "quay.io/package-operator/non-existent:v123",
		},
	}

	ctx := logr.NewContext(context.Background(), testr.New(t))
	require.NoError(t, Client.Create(ctx, pkg))
	cleanupOnSuccess(ctx, t, pkg)

	require.NoError(t,
		Waiter.WaitForCondition(ctx, pkg, corev1alpha1.PackageUnpacked, metav1.ConditionFalse))

	existingPackage := &corev1alpha1.Package{}
	err := Client.Get(ctx, client.ObjectKey{Name: "non-existent", Namespace: "default"}, existingPackage)
	require.NoError(t, err)
	require.Equal(t, "ImagePullBackOff", existingPackage.Status.Conditions[0].Reason)
}

func TestPackage_externalObject_teardown(t *testing.T) {
	ctx := logr.NewContext(context.Background(), testr.New(t))
	meta := metav1.ObjectMeta{
		Name: "success",
	}
	spec := corev1alpha1.PackageSpec{
		Image: SuccessTestPackageImage,
		Config: &runtime.RawExtension{
			Raw: []byte(fmt.Sprintf(`{"testStubImage": "%s"}`, TestStubImage)),
		},
	}
	pkg := newPackage(meta, spec, true)
	requireDeployPackage(ctx, t, pkg, &corev1alpha1.ObjectDeployment{})
	// Test if environment information is injected successfully.
	deploy := &appsv1.Deployment{}
	err := Client.Get(ctx, client.ObjectKey{
		Name:      "test-stub-success",
		Namespace: "default",
	}, deploy)
	require.NoError(t, err)

	var env manifestsv1alpha1.PackageEnvironment
	te := deploy.Annotations["test-environment"]
	err = json.Unmarshal([]byte(te), &env)
	require.NoError(t, err)
	assert.NotEmpty(t, env.Kubernetes.Version)
	objectSet := &corev1alpha1.ObjectSet{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "external",
			Namespace: "default",
		},
		Spec: corev1alpha1.ObjectSetSpec{
			ObjectSetTemplateSpec: corev1alpha1.ObjectSetTemplateSpec{
				AvailabilityProbes: []corev1alpha1.ObjectSetProbe{
					{
						Selector: corev1alpha1.ProbeSelector{
							Kind: &corev1alpha1.PackageProbeKindSpec{
								Kind:  "Deployment",
								Group: "apps",
							},
						},
						Probes: []corev1alpha1.Probe{
							{
								FieldsEqual: &corev1alpha1.ProbeFieldsEqualSpec{
									FieldA: ".status.updatedReplicas",
									FieldB: ".status.replicas",
								},
							},
							{
								Condition: &corev1alpha1.ProbeConditionSpec{
									Status: string(metav1.ConditionTrue),
									Type:   string(appsv1.DeploymentAvailable),
								},
							},
						},
					},
				},
				Phases: []corev1alpha1.ObjectSetTemplatePhase{
					{
						Name: "deploy",
						ExternalObjects: []corev1alpha1.ObjectSetObject{
							{
								Object: unstructured.Unstructured{
									Object: map[string]any{
										"apiVersion": "apps/v1",
										"kind":       "Deployment",
										"metadata": map[string]any{
											"name": "test-stub-success",
										},
									},
								},
								ConditionMappings: []corev1alpha1.ConditionMapping{
									{
										SourceType:      string(appsv1.DeploymentAvailable),
										DestinationType: "a/DeploymentAvailable",
									},
									{
										SourceType:      string(appsv1.DeploymentProgressing),
										DestinationType: "a/DeploymentProgressing",
									},
								},
							},
						},
					},
				},
			},
		},
	}
	require.NoError(t, Client.Create(ctx, objectSet))
	require.NoError(t, Waiter.WaitForCondition(ctx, objectSet, corev1alpha1.ObjectSetAvailable, metav1.ConditionTrue))
	cleanupOnSuccess(ctx, t, objectSet)
	require.NoError(t, Client.Delete(ctx, objectSet))
	require.NoError(t,
		Waiter.WaitForObject(
			ctx, deploy, "external ownerReference to be removed",
			func(client.Object) (bool, error) {
				deploy = &appsv1.Deployment{}
				err = Client.Get(ctx, client.ObjectKey{
					Name:      "test-stub-success",
					Namespace: "default",
				}, deploy)
				externalOwner := false
				for _, own := range deploy.GetOwnerReferences() {
					if own.Name == "external" {
						externalOwner = true
					}
				}
				return len(deploy.GetOwnerReferences()) == 1 && !externalOwner, err
			},
			wait.WithTimeout(15*time.Second),
		))
}

func TestPackage_internalObject_teardown(t *testing.T) {
	ctx := logr.NewContext(context.Background(), testr.New(t))
	meta := metav1.ObjectMeta{
		Name: "success",
	}
	spec := corev1alpha1.PackageSpec{
		Image: SuccessTestPackageImage,
		Config: &runtime.RawExtension{
			Raw: []byte(fmt.Sprintf(`{"testStubImage": "%s"}`, TestStubImage)),
		},
	}
	pkg := newPackage(meta, spec, true)
	requireDeployPackage(ctx, t, pkg, &corev1alpha1.ObjectDeployment{})
	// Test if environment information is injected successfully.
	deploy := &appsv1.Deployment{}
	err := Client.Get(ctx, client.ObjectKey{
		Name:      "test-stub-success",
		Namespace: "default",
	}, deploy)
	require.NoError(t, err)
	cleanupOnSuccess(ctx, t, deploy)
	var env manifestsv1alpha1.PackageEnvironment
	te := deploy.Annotations["test-environment"]
	err = json.Unmarshal([]byte(te), &env)
	require.NoError(t, err)
	assert.NotEmpty(t, env.Kubernetes.Version)
	deploy.OwnerReferences[0].Controller = ptr.To(false)
	err = Client.Update(ctx, deploy)
	require.NoError(t, err)
	objDeploy := &corev1alpha1.ObjectDeployment{}
	err = Client.Get(ctx, client.ObjectKey{
		Name:      "success",
		Namespace: "default",
	}, objDeploy)
	require.NoError(t, err)
	templateHash := objDeploy.Status.TemplateHash
	objectSet := &corev1alpha1.ObjectSet{}
	err = Client.Get(ctx, client.ObjectKey{
		Name:      "success-" + templateHash,
		Namespace: "default",
	}, objectSet)
	require.NoError(t, err)
	require.NoError(t, Client.Delete(ctx, objectSet))
	require.NoError(t,
		Waiter.WaitForObject(
			ctx, deploy, "internal ownerReference to be removed",
			func(client.Object) (bool, error) {
				deploy = &appsv1.Deployment{}
				err = Client.Get(ctx, client.ObjectKey{
					Name:      "test-stub-success",
					Namespace: "default",
				}, deploy)
				internalOwner := false
				for _, own := range deploy.GetOwnerReferences() {
					if own.Name == "success" {
						internalOwner = true
					}
				}
				return len(deploy.GetOwnerReferences()) == 1 && !internalOwner, err
			},
			wait.WithTimeout(40*time.Second),
		))
}
