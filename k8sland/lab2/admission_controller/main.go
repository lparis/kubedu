package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"k8s.io/api/admission/v1beta1"
	appv1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

const (
	denyLabel = "Grim-Reaper"
	appName   = "DepAdm"
	port      = ":443"
)

type (
	// Config contains the server (the webhook) cert and key.
	Config struct {
		CertFile string
		KeyFile  string
	}
	admitFunc func(v1beta1.AdmissionReview) *v1beta1.AdmissionResponse
)

func main() {
	r := mux.NewRouter()
	m := handlers.LoggingHandler(os.Stdout, r)
	r.Handle("/", http.HandlerFunc(handlePod))

	var config Config
	server := &http.Server{
		Addr:      port,
		Handler:   m,
		TLSConfig: configTLS(config, getClient()),
	}

	log.Printf("%s Listening on port %s...", appName, port)
	server.ListenAndServeTLS("", "")
}

func handlePod(w http.ResponseWriter, r *http.Request) {
	serve(w, r, admitDeployment)
}

// Reject fred's deployments
func admitDeployment(ar v1beta1.AdmissionReview) *v1beta1.AdmissionResponse {
	log.Println("Checking Deployment Admission...")

	// Check if we have the right resource. Hint! apps/v1/deployments
	depResource := metav1.GroupVersionResource{
		Group:    !!CHANGE_ME!!,
		Version:  !!CHANGE_ME!!,
		Resource: !!CHANGE_ME!!,
	}
	if ar.Request.Resource != depResource {
		err := fmt.Errorf("expect resource to be %s", depResource)
		log.Println("Boom", err)
		return toAdmissionResponse(err)
	}

	raw := ar.Request.Object.Raw
	var dep appv1.Deployment
	if _, _, err := codecs.UniversalDeserializer().Decode(raw, nil, &dep); err != nil {
		log.Println("Boom", err)
		return toAdmissionResponse(err)
	}

	reviewResponse := v1beta1.AdmissionResponse{Allowed: true}
	var msg string

	// Check deployment labels for the magic label. If found set Allowed to false.
	// Also set deny message!
	// Hint! msg = fmt.Sprintf("👻  Seriously `%s? No buzz kill allowed on this cluster!!", denyLabel)
	!!YOUR_CODE!!

	if !reviewResponse.Allowed {
		log.Printf("Rejecting Deployment %s", dep.ObjectMeta.Name)
		reviewResponse.Result = &metav1.Status{Message: strings.TrimSpace(msg)}
	} else {
		log.Printf("Admitting Deployment %s", dep.ObjectMeta.Name)
	}

	return &reviewResponse
}

func serve(w http.ResponseWriter, r *http.Request, admit admitFunc) {
	var body []byte
	if r.Body != nil {
		if data, err := ioutil.ReadAll(r.Body); err == nil {
			body = data
		}
		defer r.Body.Close()
	}

	// verify the content type is accurate
	contentType := r.Header.Get("Content-Type")
	if contentType != "application/json" {
		log.Println("Boom!", "contentType=%s, expect application/json", contentType)
		return
	}

	var (
		reviewResponse *v1beta1.AdmissionResponse
		ar             v1beta1.AdmissionReview
	)
	if _, _, err := codecs.UniversalDeserializer().Decode(body, nil, &ar); err != nil {
		log.Println("Boom!", err)
		reviewResponse = toAdmissionResponse(err)
	} else {
		reviewResponse = admit(ar)
	}

	var response v1beta1.AdmissionReview
	if reviewResponse != nil {
		response.Response = reviewResponse
		response.Response.UID = ar.Request.UID
	}
	// reset the Object and OldObject, they are not needed in a response.
	ar.Request.Object = runtime.RawExtension{}
	ar.Request.OldObject = runtime.RawExtension{}

	resp, err := json.Marshal(response)
	if err != nil {
		log.Println("Boom!", err)
		return
	}

	if _, err := w.Write(resp); err != nil {
		log.Println("Boom!", err)
	}
}

func toAdmissionResponse(err error) *v1beta1.AdmissionResponse {
	return &v1beta1.AdmissionResponse{
		Result: &metav1.Status{
			Message: err.Error(),
		},
	}
}
