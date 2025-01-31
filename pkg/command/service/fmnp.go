// Copyright 2022 The Knative Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package service

import (
	"bytes"
	"context"
	"encoding/csv"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"log"
	"os/exec"
	"strings"
	"sync"
	"time"

	"github.com/spf13/cobra"
	vegeta "github.com/tsenart/vegeta/v12/lib"

	_ "k8s.io/client-go/plugin/pkg/client/auth"

	v1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"

	"k8s.io/apimachinery/pkg/watch"

	"knative.dev/kperf/pkg"
	"knative.dev/kperf/pkg/config"
	"knative.dev/serving/pkg/apis/serving"
	servingv1 "knative.dev/serving/pkg/apis/serving/v1"
	servingv1client "knative.dev/serving/pkg/client/clientset/versioned/typed/serving/v1"
)

const (
	FNMPOutputFilename = "ksvc_fnmp_time"
)

func NewServiceFNMPCommand(p *pkg.PerfParams) *cobra.Command {
	loadArgs := pkg.LoadArgs{}
	serviceLoadCommand := &cobra.Command{
		Use:   "fnmp",
		Short: "Load test and Find maximum number of pods ran concurrently",
		Long: `Scale Knative service from zero and measure the maximum number of pods deployed

For example:
# To measure a Knative Service scaling from zero to N
kperf service load --namespace ktest --svc-prefix ktest --range 0,3 --load-tool wrk --load-duration 60s --load-concurrency 40 --verbose --output /tmp

Try to seat a number of concurrency big enough to create more pods tha the node can handle,
also service needs to be applied onlt to one pod`,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			err := config.BindFlags(cmd, "service.fnmp.", nil)
			if err != nil {
				return err
			}
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return LoadServicesUpFromZero_fnmp(p, loadArgs)
		},
	}

	serviceLoadCommand.Flags().StringVarP(&loadArgs.Namespace, "namespace", "", "", "Service namespace")
	serviceLoadCommand.Flags().StringVarP(&loadArgs.NamespaceRange, "namespace-range", "", "", "Service namespace range")
	serviceLoadCommand.Flags().StringVarP(&loadArgs.NamespacePrefix, "namespace-prefix", "", "", "Service namespace prefix")
	serviceLoadCommand.Flags().StringVarP(&loadArgs.Svc, "svc", "", "", "Service name")
	serviceLoadCommand.Flags().StringVarP(&loadArgs.SvcPrefix, "svc-prefix", "", "", "Service name prefix")
	serviceLoadCommand.Flags().StringVarP(&loadArgs.SvcRange, "range", "r", "", "Desired service range")
	serviceLoadCommand.Flags().BoolVarP(&loadArgs.Verbose, "verbose", "v", false, "Service verbose result")
	serviceLoadCommand.Flags().BoolVarP(&loadArgs.ResolvableDomain, "resolvable", "", false, "If Service endpoint resolvable url")
	serviceLoadCommand.Flags().DurationVarP(&loadArgs.WaitPodsReadyDuration, "wait-time", "w", 10*time.Second, "Time to wait for all pods to be ready")
	serviceLoadCommand.Flags().StringVarP(&loadArgs.LoadTool, "load-tool", "t", "default", "Select the load test tool, use internal load test tool(vegeta) by default, also support external load tool(wrk and hey, require preinstallation)")
	serviceLoadCommand.Flags().StringVarP(&loadArgs.LoadConcurrency, "load-concurrency", "c", "30", "total number of workers to run concurrently for the load test tool")
	serviceLoadCommand.Flags().StringVarP(&loadArgs.LoadDuration, "load-duration", "d", "60s", "Duration of the test for the load test tool")
	serviceLoadCommand.Flags().StringVarP(&loadArgs.Output, "output", "o", ".", "Measure result location")
	serviceLoadCommand.Flags().BoolVarP(&loadArgs.Https, "https", "", false, "Use https with TLS")
	serviceLoadCommand.Flags().StringVarP(&loadArgs.Num_Reqs, "number-of-req", "", "", "Number of requests for the external load tool")

	return serviceLoadCommand
}

func LoadServicesUpFromZero_fnmp(params *pkg.PerfParams, inputs pkg.LoadArgs) error {
	ctx := context.Background()
	nsNameList, err := GetNamespaces(ctx, params, inputs.Namespace, inputs.NamespaceRange, inputs.NamespacePrefix)
	if err != nil {
		return err
	}
	loadFromZeroResult, _, err := loadAndMeasure_fnmp(ctx, params, inputs, nsNameList, getServices)
	if err != nil {
		return err
	}

	knativeVersion := GetKnativeVersion(params)
	ingressInfo := GetIngressController(params)
	loadFromZeroResult.KnativeInfo.ServingVersion = knativeVersion["serving"]
	loadFromZeroResult.KnativeInfo.EventingVersion = knativeVersion["eventing"]
	loadFromZeroResult.KnativeInfo.IngressController = ingressInfo["ingressController"]
	loadFromZeroResult.KnativeInfo.IngressVersion = ingressInfo["version"]

	rows := make([][]string, 0) // replicas ready duration of all services

	maxReplicasCount, replicasCountList := getReplicasCount_fnmp(loadFromZeroResult)

	//Add replicas ready duration results of each service to rows, only select results before the replicas reaches the maximum count
	for i := 0; i < len(loadFromZeroResult.Measurment); i++ {
		var row []string
		row = append(row, loadFromZeroResult.Measurment[i].ServiceName, loadFromZeroResult.Measurment[i].ServiceNamespace)
		for j := 0; j < len(loadFromZeroResult.Measurment[i].ReplicaResults); j++ {
			if loadFromZeroResult.Measurment[i].ReplicaResults[j].ReadyReplicasCount <= replicasCountList[i] {
				row = append(row, strconv.FormatFloat(loadFromZeroResult.Measurment[i].ReplicaResults[j].ReplicaReadyDuration, 'f', 3, 32))
			}
		}
		rows = append(rows, row)
	}

	sortSlice(rows)

	//Add the column name
	var row []string
	row = append(row, "svc_name", "svc_namespace")
	for i := 1; i <= maxReplicasCount; i++ {
		row = append(row, "replica_"+strconv.Itoa(i)+"_ready")
	}
	rows = append([][]string{row}, rows...)

	st_cont_runt := os.Getenv("CONT_RUNTIME")

	// generate CSV, HTML and JSON outputs from rows and loadFromZeroResult
	err = GenerateOutput(inputs.Output, "_"+st_cont_runt+FNMPOutputFilename, true, true, true, rows, loadFromZeroResult)
	if err != nil {
		fmt.Printf("failed to generate output: %s\n", err)
		return err
	}

	generateCSV_max_no_pods(inputs.Output, "_"+st_cont_runt+"_"+FNMPOutputFilename, loadFromZeroResult)

	return nil
}

func generateCSV_max_no_pods(output_dir string, outputFilename string, loadFromZeroResult pkg.FmnpResult) {
	outputPathPrefix, err := GenerateOutputPathPrefix(output_dir, outputFilename)

	// Create a new CSV file
	file, err := os.Create(outputPathPrefix + "max_pods.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Create a CSV writer
	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write the CSV header
	header := []string{"Service name", "max_nu_of_pods"}
	writer.Write(header)

	for _, val := range loadFromZeroResult.Measurment {
		record := []string{val.ServiceName + "_" + val.ServiceNamespace, fmt.Sprintf("%d", val.MaxNuPods)}
		writer.Write(record)
	}

	fmt.Printf("Measurement of maximum concurrent pods saved  in CSV file %s\n", outputPathPrefix+"max_pods.csv")
}

func loadAndMeasure_fnmp(ctx context.Context, params *pkg.PerfParams, inputs pkg.LoadArgs, nsNameList []string, servicesListFunc func(context.Context, servingv1client.ServingV1Interface, []string, string, string, string) ([]ServicesToScale, error)) (pkg.FmnpResult, int, error) {
	result := pkg.FmnpResult{}
	ksvcClient, err := params.NewServingClient()
	var max_nu_pod int
	if err != nil {
		return result, max_nu_pod, err
	}
	objs, err := servicesListFunc(ctx, ksvcClient, nsNameList, inputs.SvcPrefix, inputs.SvcRange, inputs.Svc)
	if err != nil {
		return result, max_nu_pod, err
	}
	count := len(objs)

	var wg sync.WaitGroup
	var m sync.Mutex
	// var st string = "10s"
	wg.Add(count)
	for i := 0; i < count; i++ {
		go func(ndx int, m *sync.Mutex) {
			defer wg.Done()

			loadToolOutput, loadResult, max_nu_pod, err := runLoadFromZero_fnmp(ctx, params, inputs, objs[ndx].Namespace, objs[ndx].Service)
			if err == nil {
				// print result(load test tool output, replicas result, pods result)
				if inputs.Verbose {
					fmt.Printf("\n[Verbose] Namespace %s, Service %s:\n", loadResult.ServiceNamespace, loadResult.ServiceName)
					fmt.Printf("\n[Verbose] Load tool(%s) output:\n%s\n", inputs.LoadTool, loadToolOutput)
					fmt.Printf("[Verbose] Deployment replicas changed from 0 to %d:\n", len(loadResult.ReplicaResults))
					fmt.Printf("replicas\tready_duration(seconds)\n")
					for i := 0; i < len(loadResult.ReplicaResults); i++ {
						fmt.Printf("%8d\t%23.3f\n", i, loadResult.ReplicaResults[i].ReplicaReadyDuration)
					}
					fmt.Printf("\n[Verbose] Pods changed from 0 to %d:\n", len(loadResult.PodResults))
					fmt.Printf("pods\tready_duration(seconds)\n")
					for i := 0; i < len(loadResult.PodResults); i++ {
						fmt.Printf("%4d\t%23.1f\n", i, loadResult.PodResults[i].PodReadyDuration)
					}
					fmt.Printf("\n[Verbose] Max number of pods running is (%d)\n", max_nu_pod)
					fmt.Printf("\n---------------------------------------------------------------------------------\n")
				}
				m.Lock()
				result.Measurment = append(result.Measurment, loadResult)
				m.Unlock()
			} else {
				fmt.Printf("failed in runLoadFromZero: %s\n", err)
			}
		}(i, &m)
	}
	wg.Wait()

	return result, max_nu_pod, nil
}

func runLoadFromZero_fnmp(ctx context.Context, params *pkg.PerfParams, inputs pkg.LoadArgs, namespace string, svc *servingv1.Service) (
	string, pkg.FmnpFromZeroResult, int, error) {
	selector := labels.SelectorFromSet(labels.Set{
		serving.ServiceLabelKey: svc.Name,
	})
	var loadOutput string
	var loadResult pkg.FmnpFromZeroResult
	var replicaResults []pkg.LoadReplicaResult
	var podResults []pkg.LoadPodResult

	watcher, err := params.ClientSet.AppsV1().Deployments(namespace).Watch(
		context.Background(), metav1.ListOptions{LabelSelector: selector.String()})
	if err != nil {
		m := fmt.Sprintf("unable to watch the deployment for the service: %v", err)
		log.Println(m)
		return "", loadResult, 0, errors.New(m)
	}
	defer watcher.Stop()

	//rdch := watcher.ResultChan() // replica duration channel
	pdch := make(chan struct{}) // pod duration channel
	errch := make(chan error, 1)

	host := svc.Status.RouteStatusFields.URL.URL().Host

	endpoint_env := os.Getenv("ENDPOINT_URL")

	if err != nil {
		fmt.Println("Error during conversion")
		return "", loadResult, 0, errors.New(err.Error())
	}

	fmt.Printf("Endpoint set to %s\n", endpoint_env)

	fmt.Printf(" The external tool will be like : hey -n %s -c %s -z %s -host %s %s \n", inputs.Num_Reqs, inputs.LoadConcurrency, inputs.LoadDuration, host, endpoint_env)

	loadStart := time.Now()
	log.Printf("Namespace %s, Service %s, load start\n", namespace, svc.Name)

	go func() {
		if inputs.LoadTool == "default" {
			loadOutput, err = runInternalVegeta_fnmp(inputs, endpoint_env, host)
			if err != nil {
				errch <- fmt.Errorf("failed to run internal load tool: %w", err)
				return
			}
		} else {
			loadOutput, err = runExternalLoadTool_fnmp(inputs, namespace, svc.Name, endpoint_env, host)
			if err != nil {
				errch <- fmt.Errorf("failed to run external load tool: %w", err)
				return
			}
		}
		loadEnd := time.Now()
		loadDuration := loadEnd.Sub(loadStart)
		log.Printf("Namespace %s, Service %s, load end, take off %.3f seconds\n", namespace, svc.Name, loadDuration.Seconds())

		time.Sleep(inputs.WaitPodsReadyDuration) //wait for all pods ready
		pdch <- struct{}{}
	}()

	var max_nu_pod int
	for {
		select {
		case <-pdch:
			podResults, max_nu_pod, err = getPodResults_fnmp(ctx, params, namespace, svc)
			if err != nil {
				return loadOutput, loadResult, max_nu_pod, err
			}
			// set loadResult
			loadResult = setLoadFromZeroResult_fnmp(namespace, svc, replicaResults, podResults)
			loadResult.MaxNuPods = max_nu_pod
			return loadOutput, loadResult, max_nu_pod, nil
		case err := <-errch:
			return "", loadResult, max_nu_pod, err
		}
	}
}

// getSvcPod gets pod list by namespace and service name.
func getSvcPods_fnmp(ctx context.Context, params *pkg.PerfParams, namespace string, svcName string) (PodList []corev1.Pod, err error) {
	selector := labels.SelectorFromSet(labels.Set{
		serving.ServiceLabelKey: svcName,
	})

	ksvcPodList, err := params.ClientSet.CoreV1().Pods(namespace).List(ctx, metav1.ListOptions{LabelSelector: selector.String()})
	if err != nil {
		return nil, err
	}

	return ksvcPodList.Items, nil
}

// getReplicasCount gets the maximum count of replicas in each service and the maximum count of replicas in all services
func getReplicasCount_fnmp(loadResult pkg.FmnpResult) (int, []int) {
	var maxReplicasCount int     // the maximum count in replicasCountList
	replicasCountList := []int{} // the maximum replicas count in each service
	for _, m := range loadResult.Measurment {
		count := 0

		for _, d := range m.ReplicaResults {
			if d.ReadyReplicasCount > count {
				count = d.ReadyReplicasCount
			}
		}
		replicasCountList = append(replicasCountList, count)
		if count > maxReplicasCount {
			maxReplicasCount = count
		}
	}
	return maxReplicasCount, replicasCountList
}

// runInternalVegeta runs internal load test tool(vegeta) using library, returns load output and error
func runInternalVegeta_fnmp(inputs pkg.LoadArgs, endpoint string, host string) (output string, err error) {
	concurrency, err := strconv.ParseUint(inputs.LoadConcurrency, 10, 64)
	if err != nil {
		return "", fmt.Errorf("failed to get load concurrency: %s", err)
	}

	duration, err := time.ParseDuration(inputs.LoadDuration)
	if err != nil {
		return "", fmt.Errorf("failed to get load duration: %s", err)
	}

	rate := vegeta.Rate{Freq: 8, Per: time.Second}
	targeter := vegeta.NewStaticTargeter(vegeta.Target{
		Method: "GET",
		URL:    endpoint,
		Header: http.Header{
			"Host": []string{
				host,
			},
		},
	})

	attacker := vegeta.NewAttacker(vegeta.Workers(concurrency))

	var metrics vegeta.Metrics
	for res := range attacker.Attack(targeter, rate, duration, "Big Bang!") {
		metrics.Add(res)
	}
	defer metrics.Close()

	repText := vegeta.NewTextReporter(&metrics)

	buf := new(bytes.Buffer)
	err = repText.Report(buf)
	if err != nil {
		return "", fmt.Errorf("failed to write result to buffer: %s", err)
	}

	return buf.String(), nil
}

// runExternalLoadTool runs external load test tool(wrk or hey) using command line, returns load output and error
func runExternalLoadTool_fnmp(inputs pkg.LoadArgs, namespace string, svcName string, endpoint string, host string) (output string, err error) {
	// Prepare command for load test tool
	cmd, wrkLua, err := loadCmdBuilder_fnmp(inputs, namespace, svcName, endpoint, host)
	if err != nil {
		return "", fmt.Errorf("failed to run loadCmdBuilder: %s", err)
	}
	defer func() {
		// Delete wrk lua script
		if strings.EqualFold(inputs.LoadTool, "wrk") {
			err := deleteFile(wrkLua)
			if err != nil {
				fmt.Printf("%s\n", err)
			}
		}
	}()

	runCmd := exec.Command("/bin/sh", "-c", cmd)
	var loadOutput []byte
	loadOutput, err = runCmd.Output()
	output = string(loadOutput)
	if err != nil {
		return "", fmt.Errorf("failed to run load command: %s", err)
	}
	return output, nil
}

// loadCmdBuilder builds the command to run load tool, returns command and wrk lua script.
func loadCmdBuilder_fnmp(inputs pkg.LoadArgs, namespace string, svcName string, endpoint string, host string) (string, string, error) {
	var cmd strings.Builder
	var wrkLuaFilename string
	if strings.EqualFold(inputs.LoadTool, "hey") {
		cmd.WriteString("hey")
		cmd.WriteString(" -n ")
		cmd.WriteString(inputs.Num_Reqs)
		cmd.WriteString(" -c ")
		cmd.WriteString(inputs.LoadConcurrency)
		cmd.WriteString(" -z ")
		cmd.WriteString(inputs.LoadDuration)
		cmd.WriteString(" -host ")
		cmd.WriteString(host)
		cmd.WriteString(" ")
		cmd.WriteString(endpoint)
		return cmd.String(), "", nil
	}

	if strings.EqualFold(inputs.LoadTool, "wrk") {
		// creat lua script to config host of URL
		wrkLuaFilename = "./wrk_" + namespace + "_" + svcName + ".lua"
		var content strings.Builder
		content.WriteString("wrk.host = \"")
		content.WriteString(host)
		content.WriteString("\"")
		data := []byte(content.String())
		err := os.WriteFile(wrkLuaFilename, data, 0644)
		if err != nil {
			return "", "", fmt.Errorf("write wrk lua script error: %w", err)
		}
		cmd.WriteString("wrk")
		cmd.WriteString(" -c ")
		cmd.WriteString(inputs.LoadConcurrency)
		cmd.WriteString(" -d ")
		cmd.WriteString(inputs.LoadDuration)
		cmd.WriteString(" -s ")
		cmd.WriteString(wrkLuaFilename)
		cmd.WriteString(" ")
		cmd.WriteString(endpoint)
		cmd.WriteString(" --latency")
		return cmd.String(), wrkLuaFilename, nil
	}

	return "", "", fmt.Errorf("kperf only support hey and wrk now")
}

// getReplicaResult get replicaResult by watching deployment, and append replicaResult to replicaResults
func getReplicaResult_fnmp(replicaResults []pkg.LoadReplicaResult, event watch.Event, loadStart time.Time) []pkg.LoadReplicaResult {
	var replicaResult pkg.LoadReplicaResult
	dm := event.Object.(*v1.Deployment)
	readyReplicas := int(dm.Status.ReadyReplicas)
	if event.Type == watch.Modified && readyReplicas > len(replicaResults) {
		replicaResult.ReplicaReadyTime = time.Now()
		replicaResult.ReadyReplicasCount = readyReplicas
		//apo tote pou ksekinise to load(hey) poso phre na ginei ready to replica
		replicaResult.ReplicaReadyDuration = replicaResult.ReplicaReadyTime.Sub(loadStart).Seconds()
		replicaResults = append(replicaResults, replicaResult)
	}
	return replicaResults
}

// getPodResults gets podReadyDuration of all pods and append result to podResults
func getPodResults_fnmp(ctx context.Context, params *pkg.PerfParams, namespace string, svc *servingv1.Service) ([]pkg.LoadPodResult, int, error) {
	var r pkg.LoadPodResult
	var podResults []pkg.LoadPodResult
	var count int
	podList, err := getSvcPods_fnmp(ctx, params, namespace, svc.Name)
	if err != nil {
		return nil, 0, err
	}
	count = getRunningPods(podList)
	for i := 0; i < len(podList); i++ {
		pod := podList[i]
		podCreatedTime := pod.GetCreationTimestamp().Rfc3339Copy()
		present, PodReadyCondition := getPodCondition(&pod.Status, corev1.PodReady)
		if present == -1 {
			log.Println("failed to find Pod Condition PodReady and skip measuring")
			continue
		}
		podReadyTime := PodReadyCondition.LastTransitionTime.Rfc3339Copy()
		//poso xrono phre na mexri na ftasei sto ready
		podReadyDuration := podReadyTime.Sub(podCreatedTime.Time).Seconds()
		r.PodCreateTime = podCreatedTime
		r.PodReadyTime = podReadyTime
		r.PodReadyDuration = podReadyDuration
		podResults = append(podResults, r)
	}
	return podResults, count, nil
}

// setLoadFromZeroResult sets every item of LoadFromZeroResult
func setLoadFromZeroResult_fnmp(namespace string, svc *servingv1.Service, replicaResults []pkg.LoadReplicaResult, podResults []pkg.LoadPodResult) pkg.FmnpFromZeroResult {
	var loadResult pkg.FmnpFromZeroResult
	var totalReadyReplicas int
	for _, value := range replicaResults {
		if value.ReadyReplicasCount > totalReadyReplicas {
			totalReadyReplicas = value.ReadyReplicasCount
		}
	}
	loadResult.TotalReadyReplicas = totalReadyReplicas
	loadResult.TotalReadyPods = len(podResults)
	loadResult.ReplicaResults = replicaResults
	loadResult.PodResults = podResults

	loadResult.ServiceName = svc.Name
	loadResult.ServiceNamespace = namespace

	return loadResult
}

// Count pod that have containers in Running State only
func getRunningPods(podList []corev1.Pod) int {
	count := 0
	for _, pd := range podList {
		if pd.Status.Phase == "Running" && checkContainersRunning(pd.Status.ContainerStatuses) {
			count++
		}
	}
	return count
}

// Check which containers are Running and not in Terminating state
func checkContainersRunning(contStList []corev1.ContainerStatus) bool {
	for _, val := range contStList {
		if val.State.Terminated != nil {
			return false
		}
	}
	return true
}
