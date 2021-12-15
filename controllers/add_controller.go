/*
Copyright 2021.

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

package controllers

import (
	"context"

	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	mathv1 "github.com/example/add-operator/api/v2"
)

// AddReconciler reconciles a Add object
type AddReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=math.example.com,resources=adds,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=math.example.com,resources=adds/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=math.example.com,resources=adds/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the Add object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.10.0/pkg/reconcile
func (r *AddReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	lg := log.FromContext(ctx)

	// Get
	cr := &mathv1.Add{}
	err := r.Get(ctx, req.NamespacedName, cr)
	if err != nil {
		return ctrl.Result{}, err
	}

	// update
	cr.Status.R = cr.Spec.A * cr.Spec.B
	err = r.Status().Update(ctx, cr)
	if err != nil {
		return ctrl.Result{}, err
	}

	lg.Info("Reconciled")

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *AddReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&mathv1.Add{}).
		Complete(r)
}
