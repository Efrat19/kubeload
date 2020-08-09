/*


Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
//https://github.com/kubernetes-sigs/kubebuilder/blob/master/docs/book/src/cronjob-tutorial/testdata/project/controllers/cronjob_controller.go

package controllers

import (
	"context"
	//"encoding/json"
	"fmt"
	//"k8s.io/apimachinery/pkg/types"
	//"k8s.io/client-go/kubernetes"
	"math"
	"time"

	kubeloadv1 "github.com/Efrat19/kubeload/api/v1"
	"github.com/go-logr/logr"
	batch "k8s.io/api/batch/v1"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// LoadManagerReconciler reconciles a LoadManager object
type LoadManagerReconciler struct {
	client.Client
	Log    logr.Logger
	Scheme *runtime.Scheme
}

// +kubebuilder:rbac:groups=kubeload.kubeload.efrat19.io,resources=loadmanagers,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=kubeload.kubeload.efrat19.io,resources=loadmanagers/status,verbs=get;update;patch

func (r *LoadManagerReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	ctx := context.Background()
	log := r.Log.WithValues("loadmanager", req.NamespacedName)
	var loadManager kubeloadv1.LoadManager
	if err := r.Get(ctx, req.NamespacedName, &loadManager); err != nil {
		log.Error(err, "unable to fetch loadManagers")
		// we'll ignore not-found errors, since they can't be fixed by an immediate
		// requeue (we'll need to wait for a new notification), and we can get them
		// on deleted requests.
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}
	log.Info(fmt.Sprintf("max load %v", loadManager.Spec.LoadSetup.MaxLoad))
	var childJobs batch.JobList
	if err := r.List(ctx, &childJobs, client.InNamespace(req.Namespace)); err != nil {
		log.Error(err, "unable to list child Jobs")
		return ctrl.Result{}, err
	}
	desiredLoad := getDesiredLoad(&loadManager.Spec.LoadSetup, loadManager.ObjectMeta.CreationTimestamp.Time)

	for _, job := range childJobs.Items {
		fmt.Printf("1: %+v\n", *job.Spec.Parallelism)
		fmt.Printf("2: %+v\n", desiredLoad)
		if *job.Spec.Parallelism != desiredLoad {
			var newJob batch.Job
			job.DeepCopyInto(&newJob)
			job.Spec.Parallelism = &desiredLoad
			job.SetManagedFields(nil)
			newJob.SetManagedFields(nil)
			applyOpts := []client.PatchOption{client.ForceOwnership, client.FieldOwner("kubeload")}
			err := r.Patch(ctx, &newJob, client.Apply, applyOpts...)
			if err != nil {
				log.Error(err, "unable to patch job")
				return ctrl.Result{}, err
			}
			job.SetManagedFields(nil)
			log.Info("patched")
			fmt.Printf("3: %v\n", job.Name)
			fmt.Printf("3: %v\n", *job.Spec.Parallelism)
			//r.patchJobParallelism(job.Name,desiredLoad)
		}
	}
	// your logic here

	return ctrl.Result{}, nil
}

func (r *LoadManagerReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&kubeloadv1.LoadManager{}).
		Complete(r)
}

func getDesiredLoad(ls *kubeloadv1.LoadSetup, createdAt time.Time) int32 {
	minSecondsInterval := 0.1
	intervalSeconds := math.Max(ls.Interval.Seconds(), minSecondsInterval)
	secondsPassed := time.Now().Sub(createdAt).Seconds()
	loadToBeAdded := uint64(secondsPassed/intervalSeconds) * ls.HatchRate
	return int32(math.Min(float64(ls.InitialLoad+loadToBeAdded), float64(ls.MaxLoad)))
}

type patchInt32Value struct {
	Op    string `json:"op"`
	Path  string `json:"path"`
	Value int32  `json:"value"`
}

//func (r *LoadManagerReconciler) patchJobParallelism(jobName string, parallelism int32) error {
//	payload := []patchInt32Value{{
//		Op:    "replace",
//		Path:  "/spec/parallelism",
//		Value: parallelism,
//	}}
//	payloadBytes, _ := json.Marshal(payload)
//	r.Patch()
//	_, err := clientSet.
//		BatchV1().
//		Jobs("default").
//		Patch(jobName, types.JSONPatchType, payloadBytes)
//	return err
//}