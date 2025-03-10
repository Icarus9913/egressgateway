// Copyright 2022 Authors of spidernet-io
// SPDX-License-Identifier: Apache-2.0

package config

import (
	"fmt"
	"os"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	f, err := os.CreateTemp("", "example-")
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	// 输出临时文件的文件名
	fmt.Println("Created temp file:", f.Name())
	defer os.Remove(f.Name())

	os.Setenv("CONFIGMAP_PATH", f.Name())

	kubefile, err := os.CreateTemp("", "")
	if err != nil {
		t.Fatal(err)
	}
	defer kubefile.Close()
	_, err = kubefile.WriteString("apiVersion: v1\nclusters:\n- cluster:\n    certificate-authority-data: LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSUMvakNDQWVhZ0F3SUJBZ0lCQURBTkJna3Foa2lHOXcwQkFRc0ZBREFWTVJNd0VRWURWUVFERXdwcmRXSmwKY201bGRHVnpNQjRYRFRJek1EY3hNakE0TXpZMU5Wb1hEVE16TURjd09UQTRNelkxTlZvd0ZURVRNQkVHQTFVRQpBeE1LYTNWaVpYSnVaWFJsY3pDQ0FTSXdEUVlKS29aSWh2Y05BUUVCQlFBRGdnRVBBRENDQVFvQ2dnRUJBTWh1CjM3bkdBek5TSGtOb1c0cW1RSzVScXp1VUlldkNjVWF4eWlZbkQwTE9yYkZVZ0lFVnRZZUEvN2psbjJidHVpVXYKVEsrRWliRTFNMUs0OC9IYk1ZVlh2WEtERDhzbTBmZ3lJVDhnc04rbFBwVVpZdFc2cGFXbWFUVnRuUWFQNU8vaAp6MEorcFIxeTkyQTFJL2ZmSVBEa2xZbXdwSldMa1BFU1IvRmRZMm9Bb2UwejFKTjZ4VTlNSGVvcVZnckc5d3dLCnRvTTNTTnFoWXZNa3VVRnFGN0Zrc0U1aUJWRmxEUXJLblZrM0p6Q25PR0tSU0FidVdPS2huZ0g0eUNVVk5ydWIKdWpDbU1iTUFSQ09uazBCUnVvcXZSdnNOeU9SdXpxdXdXSk5lQXNXeGZrT09zT1hZWmFBV0k1Zm4vLzFqbTArVwp4VVZkL0dSVjJ6TjRodW5QNEcwQ0F3RUFBYU5aTUZjd0RnWURWUjBQQVFIL0JBUURBZ0trTUE4R0ExVWRFd0VCCi93UUZNQU1CQWY4d0hRWURWUjBPQkJZRUZFa21rTmY0WGFXT0pZRHhlNWFFVUxMUlVJTVFNQlVHQTFVZEVRUU8KTUF5Q0NtdDFZbVZ5Ym1WMFpYTXdEUVlKS29aSWh2Y05BUUVMQlFBRGdnRUJBS0tNMnJJMkozb01icGkySjdTdwpxU3hTN1FDcGljQUQyUkFKelBjMStvUmdOZndxbjhZNHVyY2dQazFkNWh3R2k5WTgxU0Nzb2YybitURldyeWNHCkFZOEhBc045M0RyUSt0ZUxWK1QzZ0xpd3BxNEp6QzFLTE1IZ3lDcU1uQXhRYjVkUUN6cFVLNjhaTG1NaVNvVnUKZnd1VTd5WjJyZmtJUUU2MVdsRW03NHQ2VjhkOFpQaVNFTXdTUDlzcE43Q0FHTHNJcElKREg5bEZtYXhjNnNDdwp4UjNUOXhqakE0SjFqSmY5RFdpQWZNWkRFMXhEREd5blZKZGdzeHRiMlFUMHVuRjRTYXZsaDB6alg2NmhrRkErCnhJd0hkWHMrbStTbi9ReWd4YkJKeXB5K0FlNkhZQU0rWE1BUnJtdXlzeHc3SFF4d2ZXM3FsSHVibThDK3JiUEwKcEJRPQotLS0tLUVORCBDRVJUSUZJQ0FURS0tLS0tCg==\n    server: https://10.6.1.21:6443\n  name: kubernetes\ncontexts:\n- context:\n    cluster: kubernetes\n    user: kubernetes-admin\n  name: kubernetes-admin@kubernetes\ncurrent-context: kubernetes-admin@kubernetes\nkind: Config\npreferences: {}\nusers:\n- name: kubernetes-admin\n  user:\n    client-certificate-data: LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSURJVENDQWdtZ0F3SUJBZ0lJRkxqUkdDa0FndlV3RFFZSktvWklodmNOQVFFTEJRQXdGVEVUTUJFR0ExVUUKQXhNS2EzVmlaWEp1WlhSbGN6QWVGdzB5TXpBM01USXdPRE0yTlRWYUZ3MHlOREEzTVRFd09ETTJOVGRhTURReApGekFWQmdOVkJBb1REbk41YzNSbGJUcHRZWE4wWlhKek1Sa3dGd1lEVlFRREV4QnJkV0psY201bGRHVnpMV0ZrCmJXbHVNSUlCSWpBTkJna3Foa2lHOXcwQkFRRUZBQU9DQVE4QU1JSUJDZ0tDQVFFQXdDWW1TQVZDNkxncnBXcFYKdUlWank4U0IrQWN0N0VKVmMyMWZXQzVUUlZhMUJ1eXk2aFNiQlhWWldTN0VZM1JySHZXWWlabkhOdmswRXdjKwptMEIzTzM1dFR5ZTYrcisxNGpkZVV6WEdhRVN6cFNma1U1VnN4TGhrWFV1dEVnNTFHZkJCSVFDNVl0cENSY3YwCld5V2daQ0Y4TEdNMkhFL1FjczQzeG9pSE04Yit5cWNaK0hKUlhSTU9kQlZsRlcwZG1KTmN0MlRQOFhGWW43d1cKbDhUWndKUVZ3T2JQVHFhVGUwTFNHY085cm9RMnN3TExtZjNMakltUXpsaVlLTFA3L3JvLzVpMnNQV2FxQnRuTwphcXBveXpNUllRTUtYREd0RmNOS1J1SkIzR0p4d0hJcnBidFNIOVdzamJROXh2c2p1emoyZDhmQkp2ODJNSGF2CmNNTUt4d0lEQVFBQm8xWXdWREFPQmdOVkhROEJBZjhFQkFNQ0JhQXdFd1lEVlIwbEJBd3dDZ1lJS3dZQkJRVUgKQXdJd0RBWURWUjBUQVFIL0JBSXdBREFmQmdOVkhTTUVHREFXZ0JSSkpwRFgrRjJsamlXQThYdVdoRkN5MFZDRApFREFOQmdrcWhraUc5dzBCQVFzRkFBT0NBUUVBVCtENjNvaVpYSUpMQ1UvWjVBS2laOWRUclgxbk9GalVMSEU3Cm5ZUHZHbXZPZmRqUGpjM2RNeXEycCtmTTlGLzVWOVNBbW1EK2Z3QWpOYk5OVUt6aHlFbVJBeDVsUGxYdk55L3kKdnc3REU0WUZYVE5zT1JLZjNIZ2JKam85dG1MMTFpZlBTZEs1V2dtVnJiMjZQVmMvYzBwWVRMb05oemg0MUNSaQpiMUkwTWpCdU9zQkFmdWRTdU9SQ0EwcFprek5LZkpJSExvTE03OVlBOFhMRWNoU2M2cXg1UWo0RVdvSzNHV0Y0CmxpU0xlRHE5ZXBJOTI1Y1BidU9MU1hGQ2NEN3lyblNLK2Vka2ZlUlhUQjhqSHBGcTlnWnR6VGpJSEhOT3NtdkYKVytaUWJEdXRIWGtrOGpsRXA2ckhsZ1BYNkZjWHB4ZDZOSGtXcXREYnAwTCtVSnFibUE9PQotLS0tLUVORCBDRVJUSUZJQ0FURS0tLS0tCg==\n    client-key-data: LS0tLS1CRUdJTiBSU0EgUFJJVkFURSBLRVktLS0tLQpNSUlFcFFJQkFBS0NBUUVBd0NZbVNBVkM2TGdycFdwVnVJVmp5OFNCK0FjdDdFSlZjMjFmV0M1VFJWYTFCdXl5CjZoU2JCWFZaV1M3RVkzUnJIdldZaVpuSE52azBFd2MrbTBCM08zNXRUeWU2K3IrMTRqZGVVelhHYUVTenBTZmsKVTVWc3hMaGtYVXV0RWc1MUdmQkJJUUM1WXRwQ1JjdjBXeVdnWkNGOExHTTJIRS9RY3M0M3hvaUhNOGIreXFjWgorSEpSWFJNT2RCVmxGVzBkbUpOY3QyVFA4WEZZbjd3V2w4VFp3SlFWd09iUFRxYVRlMExTR2NPOXJvUTJzd0xMCm1mM0xqSW1RemxpWUtMUDcvcm8vNWkyc1BXYXFCdG5PYXFwb3l6TVJZUU1LWERHdEZjTktSdUpCM0dKeHdISXIKcGJ0U0g5V3NqYlE5eHZzanV6ajJkOGZCSnY4Mk1IYXZjTU1LeHdJREFRQUJBb0lCQVFDdXh3U3pKZllDY09OaQpoeUtKd090UmdlRW1sb0V6RWZOZ0Z2Qk82WFJjOGMyZ0N0V0REbE1qMStYNXRReDEyb05GbWVleDRlclBHa1kvCnVLczkxSm1meUJQdG1Cbi8wem5DRnRMNXBVUmZ1MzRjai9pai9wcjlKU3hGb0h4QW5GM3Z4aFczeTB6Vm1lb0cKU3NwdHpmL2lsSUs2YlZQZTFNcXFZdUZnK1BiaUpHQksvalRUVWlpR2luSnRUaVJlNDBRelFLb2FueTI4N0EyaQp5R0Z3OExzTVEvS1NEbzRUNzBrcGFaZXltZnl4L21NR0VRSGJGWlhBK0tBNEJNTC9xS2svV0dBUFQ2UkhhY1BKCnFReHRhRVM5S0EydUdkWU1vcExITHAybmtDMEc3UTAvcTY1dlVNRXduSklkUmk3emhiM0FHWEdGdTdPMUN5M2wKTktqQkxYR2hBb0dCQU40eG5BdkNXa2ZJRjdNOERrU0xQRW8wYnY2RStHc1Rpd1lYZkJkRm1lUGcvQVJNTFBmYQo0TjUxWXFDQkYvK3FkZlhRaGFONmRVcDA0ZTJQNlpZKzd5dDhMeHZzNFZXd2EwN0RPTDVkbjhDUjkwM0M3RFZnCk5IM1g3WXNJQUpRTHlOdjdla3N6UXFveXZKS01aVTBEOS9LR1JxMmxIRVplaEF3MkpjSGJpRXh0QW9HQkFOMWkKVGR4TnFvZXA1RCtESUdOczhFYVRDSnBXM255eElIVkZTRTkzZC9OL3BMZE1iQTlvR2syR0pibDR0b2hyM1dEZQpKSTNpUHRORzVvMytBMG8vTWZtbWp1d1VLbHk5SDFBZVIvaE0vY3lIL1hMQXE0b0NrTDV4NVF5QWRWWWI3c0JtCngxY3ZIdTgxUEs2aTRZalAxMTNHY2dkR3MwV3FsRlY4eFNsV0lFdURBb0dCQUlVMUJMSmdFRFBjbDZqU3BsTWQKamtXR2JjeVU5MEZxYy94dzgrb1h4Z3pDQXhTb2ZvVVJhYUswaVM1a2RuakdQdlhoejF5VXUrQ3BkaEV3Si8vMQpOdm5BOTVVc1RHTk00dWhUVForRERaVXJiVEhuWENrYnhoeHo2V3RpbnNZaTBvWmZtNCtkNmFlVHgwMnNjY2JjClREZlBuR3ZhQXJ1RlNuRHZ2VzhkSi9kNUFvR0JBTWJDODlUUGhrTzNMTTQ1RkdNdjg2bnBhTmZwRm1ndFAwOEsKblJsNzBaNDFBOVh1THpiRjZKZWgwVXpzTERYZllpc09SeE44Qlp2N0ZCUjM4c3crWU1nYjJrWHE5UDIrYnRhbgoyVVg5R2dFQU4zVkh0cnQ2QWlwNlo0TUo4azhWVlE0NU9NLzE1bmd0L0FWdkI3NmxuRjc5UkhOejdwQ2x6ZmZTCnhkR1BHZit4QW9HQUoyYmxNeWhYd1ltMitiOVJzTWxJeGNHOFVKMm9DYVdDMVM0ZGQ1bzI5blJTS1Y4UmlJTG8KZUZaSlpjcDRtMlpNZkxkUVg1clNwRStaVWlFV2xuVWNSZktLSzNQbW9wU3VEK3BVTC9TaWZSbzlCMjNKNGZ6dwovaWVYVkpoajJEemJZSDZHRUtGaUttS1QzbW14WlNBY3B4OGJUYVhlT0IrK3hhdHlMaTJuTWZrPQotLS0tLUVORCBSU0EgUFJJVkFURSBLRVktLS0tLQo=")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(kubefile.Name())

	os.Setenv("KUBECONFIG", kubefile.Name())

	cfg, err := LoadConfig(false)
	if err != nil {
		t.Fatal(err)
	}
	cfg.PrintPrettyConfig()
}
