package client

import (
	"context"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

var (
	kubeconfig = `apiVersion: v1
clusters:
- cluster:
    certificate-authority-data: LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSUMvakNDQWVhZ0F3SUJBZ0lCQURBTkJna3Foa2lHOXcwQkFRc0ZBREFWTVJNd0VRWURWUVFERXdwcmRXSmwKY201bGRHVnpNQjRYRFRJeU1EVXhNekUxTlRFMU4xb1hEVE15TURVeE1ERTFOVEUxTjFvd0ZURVRNQkVHQTFVRQpBeE1LYTNWaVpYSnVaWFJsY3pDQ0FTSXdEUVlKS29aSWh2Y05BUUVCQlFBRGdnRVBBRENDQVFvQ2dnRUJBS21PCjRpM21rZFZjVnNwRnd1N3l5MFRWV3ovVVBCN0tUSlBnaHdXYzVRd3JWcjQyMm5PUithWUFEMGRYNHJIM0lzanYKZndVNXpIK1E1UDRraWE2WmFsc2twOElnS1RqU0l6U0F3OWVmbE9PWEFLM0dGN1Y3SU51MjVzTURRdWR1QytrVQo0L3lFcW9QaDZITjFXZFlnWlA3TGpabVB0d1B0UTJWbytzMjdvS2o0eW9uZWNxclBPR0NtYURoMnN2cGdaZDAxCmUyOFhEYnJsK0FQTVRiWUg5NFRQdG1DRGxBR252b0xwTHBEdUJ4bGhXRHgxZ0lIeEZ5RzBWZ1VuWHdXRlJ3Q3UKTkxSeDNMQjlhdFJuWkxRcXIvK2Z1UkVIOFFXQ2h5ZXZ3bk9BRnFvV1FEQ29MOGJFTS9vY3pKd0RkaUN4YzFKYwpmM283VlNLaWNiTG9ZWXpIZFc4Q0F3RUFBYU5aTUZjd0RnWURWUjBQQVFIL0JBUURBZ0trTUE4R0ExVWRFd0VCCi93UUZNQU1CQWY4d0hRWURWUjBPQkJZRUZKbkF2L2ozS0VmQzZTVjBvS1ZVdUhBU0RaZ0FNQlVHQTFVZEVRUU8KTUF5Q0NtdDFZbVZ5Ym1WMFpYTXdEUVlKS29aSWh2Y05BUUVMQlFBRGdnRUJBSHF5MXRsam1hanpPdWRBMmN2cApJU3UyRlhzOGcyTmpSNVpNcHcrdkI5YjhlQVdQZjdSalEvWGk2TlE4c3V0L0dnOTh1N2JlN0VoaW54Sm1VMFpRCkNkVStvNk9RMm9SMXFNUkthWDJ3S2VoRSt6ZEFxOGs5Vm96eWg3TWFOTW9yTVNsemIxbW5jT2l4NDZxQndtQk0Kb1pISGJpKzAzRXgvbEREaitwR21CbVI5d1pHWkVQcmZIWW1VYjA0bi83NnZBcVcyYWUzOXRLTlRjRExJbnFvago4VFdMTjUrWVQvRWpPQnB2NUJrc210Ym5ubklydldyTmhmSjhvSWlPVElVdjEvMHZpNnluRjB6UmN0enNGVkNICjdNcklyQ0dIV1B1bmNyWDRUbkxkMWVxTE4va0hraHdkWnNYcGwwNnNkQTdoeDJrZFVPcldpQ054U3NNemZLSWcKMUxVPQotLS0tLUVORCBDRVJUSUZJQ0FURS0tLS0tCg==
    server: https://127.0.0.1:6443
  name: kind-host
contexts:
- context:
    cluster: kind-host
    user: kind-host
  name: kind-host
current-context: kind-host
kind: Config
preferences: {}
users:
- name: kind-host
  user:
    client-certificate-data: LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSURJVENDQWdtZ0F3SUJBZ0lJSkFtQlNUYmdRK1l3RFFZSktvWklodmNOQVFFTEJRQXdGVEVUTUJFR0ExVUUKQXhNS2EzVmlaWEp1WlhSbGN6QWVGdzB5TWpBMU1UTXhOVFV4TlRkYUZ3MHlNekExTVRNeE5UVXhOVGxhTURReApGekFWQmdOVkJBb1REbk41YzNSbGJUcHRZWE4wWlhKek1Sa3dGd1lEVlFRREV4QnJkV0psY201bGRHVnpMV0ZrCmJXbHVNSUlCSWpBTkJna3Foa2lHOXcwQkFRRUZBQU9DQVE4QU1JSUJDZ0tDQVFFQXhKSEFVM0xTSTUrc0VCMnAKWlRJeVlDUThTeUsyUmZzN1gxR2xzQW5FNDlHU3RlclBGZlcwNU8raXNDQStBMURQbWc2REJHVFhpM2RFQnZJNgpZcXdKNDJ6SGFsY2Jla0ZQbzZHWHVHanJ2Uk5HRy9ySnZDT2tQN3JJZFdKQ1IxRWRSTjNVa1EzQjVFRDY5MlMvCkhOVm5ha2NXMTRQbnRQWDI2dGlRbUlTS1NPY3Y2bUVvQ1VPeWlUTHVzWDY4aVE5WkNSNUgxdTlIWXZmMU00QmsKM0t4czdYTHZRWFhXb3JLSzdMOFJ4dzFROGY1bkpZK0JtMThHNE1RbVRORHpPWm9BMVlKbkUwMWZmZkZwYkdocQpabytKbTMvUm1LTDhDVHpNS3gweDZOYVJmNFZpVWh5SjZwQkpDdnpxKzloSkdzWVVYRDAyTUFIbTd2NlRubDcyCnhjd0FhUUlEQVFBQm8xWXdWREFPQmdOVkhROEJBZjhFQkFNQ0JhQXdFd1lEVlIwbEJBd3dDZ1lJS3dZQkJRVUgKQXdJd0RBWURWUjBUQVFIL0JBSXdBREFmQmdOVkhTTUVHREFXZ0JTWndMLzQ5eWhId3VrbGRLQ2xWTGh3RWcyWQpBREFOQmdrcWhraUc5dzBCQVFzRkFBT0NBUUVBY0g3N2hGTlI3NzNpd0o1ZmI0NWdOc08ydUhPTUJZSWFBN0xtClVmTUZQY1J0c21maFBqcU9jTnhLMWtaUU5DQ0lsTmlOYjhFQ0JTNFdRZ3NnbDZYSENaWkxEcnJMS01DQlUyQSsKL0Zrd3REZFRsYXpzSG5nN0JESFdNU2pZUHRlL1kzVzQ0RjFHOWpnc0Q2blFvRHpsUTVjaFppSEJWVTZFakNMSApCcVo2RDhJRkJtUDM1QVZFazdkemtIcVVLZ0FMOVl0MU0xU1loQW9nSTZWQVVud0QxYmN6d2xLOXVIREpwcTdGClFiWXUyaUh0Uk00ZWkzWnZHR3poY1ZUMS9VY05UZ0NoL2RoM1NyUUtiOW53NFNvWXVsUVQ4a2xLZDNrd3JmWWkKVk5WNloxNkJvaEkvVGE5MFZrZ0ZQbVVPcHU4OWZYRTNqL0JYNytCNHdZYkE2NjJYRmc9PQotLS0tLUVORCBDRVJUSUZJQ0FURS0tLS0tCg==
    client-key-data: LS0tLS1CRUdJTiBSU0EgUFJJVkFURSBLRVktLS0tLQpNSUlFcFFJQkFBS0NBUUVBeEpIQVUzTFNJNStzRUIycFpUSXlZQ1E4U3lLMlJmczdYMUdsc0FuRTQ5R1N0ZXJQCkZmVzA1Tytpc0NBK0ExRFBtZzZEQkdUWGkzZEVCdkk2WXF3SjQyekhhbGNiZWtGUG82R1h1R2pydlJOR0cvckoKdkNPa1A3cklkV0pDUjFFZFJOM1VrUTNCNUVENjkyUy9ITlZuYWtjVzE0UG50UFgyNnRpUW1JU0tTT2N2Nm1FbwpDVU95aVRMdXNYNjhpUTlaQ1I1SDF1OUhZdmYxTTRCazNLeHM3WEx2UVhYV29yS0s3TDhSeHcxUThmNW5KWStCCm0xOEc0TVFtVE5Eek9ab0ExWUpuRTAxZmZmRnBiR2hxWm8rSm0zL1JtS0w4Q1R6TUt4MHg2TmFSZjRWaVVoeUoKNnBCSkN2enErOWhKR3NZVVhEMDJNQUhtN3Y2VG5sNzJ4Y3dBYVFJREFRQUJBb0lCQUZlRDlMYXlkakI2RkdjUQpiYXlxVHBkVFNxekJCWm5lb3E1cmNYTVF4bUlQbWx2MzhMNzhKOCtOaVVjVTg4Y1NJWHViWG1XRWFCcWx2Tm5DCjJvKzN2S2RPZFBJNVdmaHlQM3pBb3dYdFlKZExqM2xCakxPQXdzM0U0UjZ4NG9SUjd0QU1XMmxCVU1QSTBuTjIKblVlL0UwM1QzSzJUQW9Ra3hnd0U0MHVrSnRVUm9GTis0c3k5QmFNR3RwU1JOc0lwR2JaOTBaRjJkdUJQK0RlZgpveFV0amJadVp2VG5zUTljOWVHWExxWStKLy9ySVFwc21vSkxtTmh6dCt0R1NESVZEWGhydXZxbVpHT3VZRE1sCk5iWkNxUXp2UisxeG1Ua0NtVzVuZmRiRTNmQ1VCY3NMWFYzMkM3dE43cFlxNUhmaDdkcmgzWnBHNFR1NHZzZ1kKc1g2dmZBa0NnWUVBeldrQnRqYjcxa1dzenpXQlFnblZ1ZDdYQmdwR0pFajBhWjgzLzVkZDgxQ0lSSmxaejMrSQpSV1VsYk12NW84QTU4MW5iQnpQQ2hOUGF2eFhRaTVFOE0yc01qUW42WU5pSU96VmhhUmdUVjdmV0hWa3lTQys1Cnd2KzlRNWhaTC94RVFrV0RNVFAzejJaMUFmbTRFM3FOMElBSFRodWNVM2t6WTgzbUxCWm1ybThDZ1lFQTlQdFYKcXdlSjhOMHptY2RwTlZTUmhLTVB5ZmxmRVBmZ2FSb2VYUHN1TW91OGw2UXNWR3k1VXo4NW4yTmZOMXpOS2oyYgoxVlBzRTJWTEdxNS9nWUI4MUl0MDVGRGU2THFNZ1M3KzdQM0RkaENVM2lDVUVvZkxmRVdpbS9HK3F6ZUprMDhsCmVoeXNvOUNoY1ZWdG41SVhNTDV2V2t0OTJUUysxRFdNdzZ2dUtxY0NnWUVBdUEybUZnS1FoMytwQjRYMnl4aUsKNUdCVEpUdGhmRFBPcFRHZ2VLbkY2alkzMDlmZ3pIZUd3RCtRV0RzdzlkUlJXTWNqNWdFd0E5cmN2Nm1wVXRXUgpMclYxNm82TlJlZmZzY1gwQWJvcjRzWjcyWkpKNXJxMDVaQkhvMkRJVWFIbCs4ZlRkT3dPMlUwQi9RSW9PWFB2CnpHcGJvVlpHTGRtRS9hSEo0Nmt5Mm9zQ2dZRUFvSUpIOTkzaG9BR2VQR1F0NTNZNFBab0V6MFZtNXh3eFdVdDIKbDE4dlBvalZrTmxNL2llYUtSUGtzaXlPaHh4emcvaUhzSGJpMXpabnhkeU15QkdpT0RRQzYxQ2RMQWlGNUdJaApQcTlwTUdMZTFzYXJuWlNCV3pQWXZhbmZUaGorTjVrRXFnUTlqTHMxKzZhSVE2T1pOQ09obTV5WW9RWncvV0wwCmpvT0ljVU1DZ1lFQXlnaFR2a01VbzZkSThGQktBVnYvZ0dhNlIwQUoydXFpYW5VV1JMQys5WndRWXNoSnRRamIKZGVPNHJKOGdhQytTWnZ3cVhCUzZlZnB2SmJkdWJGNWRFc1FKSWtRQmZnMnk3UzNxQWRVelFHV0lXOGxnTVhPbgo5Yi91WlpvMVViNnpENVRIY0JUN1MzeHF1cGFRQjZnbUthSjMxMWR4QnNMdlo5Ui9pK0J5ckMwPQotLS0tLUVORCBSU0EgUFJJVkFURSBLRVktLS0tLQo=`
)

func TestClientSet(t *testing.T) {
	config, err := clientcmd.BuildConfigFromFlags("", filepath.Join(homedir.HomeDir(), ".kube", "config"))
	assert.NoError(t, err)

	client, err := kubernetes.NewForConfig(config)
	assert.NoError(t, err)
	assert.NotNil(t, client)
}

func TestClientSetWithContext(t *testing.T) {
	configLoadingRules := &clientcmd.ClientConfigLoadingRules{ExplicitPath: filepath.Join(homedir.HomeDir(), ".kube", "config")}
	configOverrides := &clientcmd.ConfigOverrides{CurrentContext: "kind-hub"}
	config, err := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(configLoadingRules, configOverrides).ClientConfig()
	assert.NoError(t, err)

	clientset, err := kubernetes.NewForConfig(config)
	assert.NoError(t, err)

	nodes, err := clientset.CoreV1().Nodes().List(context.TODO(), metav1.ListOptions{})
	assert.NoError(t, err)
	for _, node := range nodes.Items {
		t.Logf("node info: %s\n", node.Name)
	}
}

func TestListPodsWithResourceVersion(t *testing.T) {
	clientConfig, err := clientcmd.NewClientConfigFromBytes([]byte(kubeconfig))
	assert.NoError(t, err)
	config, err := clientConfig.ClientConfig()
	assert.NoError(t, err)

	client, err := kubernetes.NewForConfig(config)
	assert.NoError(t, err)

	tests := []struct {
		Name            string
		ResourceVersion string
	}{
		{
			Name:            "List pods with ResourceVersion: 0",
			ResourceVersion: "0",
		},
		{
			Name:            "List pods with empty ResourceVersion",
			ResourceVersion: "",
		},
	}
	for _, test := range tests {
		t.Logf("test for %s\n", test.Name)
		pods, err := client.CoreV1().Pods("").List(context.TODO(), metav1.ListOptions{
			ResourceVersion: test.ResourceVersion,
		})
		assert.NoError(t, err)
		for _, pod := range pods.Items {
			t.Logf("%v\t %v\t %v\n", pod.Namespace, pod.Status.Phase, pod.Name)
		}
	}
}

func TestCreateWithPrefix(t *testing.T) {
	config, err := clientcmd.BuildConfigFromFlags("", filepath.Join(homedir.HomeDir(), ".kube", "config"))
	assert.NoError(t, err)

	client, err := kubernetes.NewForConfig(config)
	assert.NoError(t, err)
	assert.NotNil(t, client)

	cm, err := client.CoreV1().ConfigMaps("kube-system").Create(context.TODO(), &corev1.ConfigMap{
		ObjectMeta: metav1.ObjectMeta{
			GenerateName: "test-",
		},
		Data: map[string]string{
			"test": "test",
		},
	}, metav1.CreateOptions{})
	assert.NoError(t, err)
	t.Log(cm)
}
