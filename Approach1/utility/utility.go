package main

import (
	"fmt"
	"os"
	"text/template"
)

type JobManifest struct {
	Name         string   // Name of the job  
	Namespace    string	  // Namespace of the pod whose information is to be fetched
	Action       string   // Action parameter
	PodName      string	  // Pod name whose information is to be fetched
	PID          string	  // Process ID parameter 
	UID          string   // UID parameter
	NameOverride string   // Name parameter
	Duration     string   // Duration parameter
	Egress       string   // Egress parameter
	Tags         string   // Tags parameter
}

func generateJobManifestFile(job JobManifest, filePath string) error {
	// Helper function for required field validation
	validateRequiredField := func(field, value string) error {
		if value == "" {
			return fmt.Errorf("%s parameter is required", field)
		}
		return nil
	}

	// Validate the required fields
	if err := validateRequiredField("Job name", job.Name); err != nil {
		return err
	}
	if err := validateRequiredField("Action", job.Action); err != nil {
		return err
	}
	if err := validateRequiredField("Namespace", job.Namespace); err != nil {
		return err
	}
	if err := validateRequiredField("Pod name", job.PodName); err != nil {
		return err
	}
	if err := validateRequiredField("Process ID", job.PID); err != nil {
		return err
	}
	if err := validateRequiredField("UID", job.UID); err != nil {
		return err
	}
	if err := validateRequiredField("Name", job.NameOverride); err != nil {
		return err
	}
	if err := validateRequiredField("Duration", job.Duration); err != nil {
		return err
	}
	if err := validateRequiredField("Egress", job.Egress); err != nil {
		return err
	}
	if err := validateRequiredField("Tags", job.Tags); err != nil {
		return err
	}

	templateStr := `
apiVersion: batch/v1
kind: Job
metadata:
  name: {{.Name}}
spec:
  template:
    metadata:
      name: client-monitor-pod
    spec:
      restartPolicy: Never
      containers:
      - name: client-monitor-container
        image: client:v1
        imagePullPolicy: Never
        resources:
          limits:
            memory: "1Gi"
            cpu: "500m"
        ports:
        - containerPort: 80
        volumeMounts:
        - name: response-volume
          mountPath: /app/response
        env:
        - name: ACTION
          value: "{{.Action}}"
        - name: PODNAME
          value: "{{.PodName}}"
        - name: NAMESPACE
          value: "{{.Namespace}}"
        - name: PID
          value: "{{.PID}}"
        - name: UID
          value: "{{.UID}}"
        - name: NAME
          value: "{{.NameOverride}}"
        - name: DURATION
          value: "{{.Duration}}"
        - name: EGRESS_PROVIDER
          value: "{{.Egress}}"
        - name: TAGS
          value: "{{.Tags}}"
      volumes:
      - name: response-volume
        emptyDir: {}
  completions: 1
  parallelism: 1
  backoffLimit: 0
  ttlSecondsAfterFinished: 90

# this section is created to give rbac(role based access control) on our cluster so that we can fetch pod ip address.
---
apiVersion: rbac.authorization.k8s.io/v1
kind: Role
metadata:
  name: pod-reader
  namespace: {{.Namespace}} # your pod's namespace 
rules:
  - apiGroups: [""]
    resources: ["pods", "services"]
    verbs: ["get", "list", "watch"]
  - apiGroups: ["extensions"]
    resources: ["deployments"]
    verbs: ["get", "list", "watch"]
---
apiVersion: rbac.authorization.k8s.io/v1
kind: RoleBinding
metadata:
  name: default-pod-reader
  namespace: {{.Namespace}} # your pod's namespace
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: Role
  name: pod-reader
subjects:
  - kind: ServiceAccount
    name: default
    namespace: default
`

	tmpl, err := template.New("jobTemplate").Parse(templateStr)
	if err != nil {
		return err
	}

	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	err = tmpl.Execute(file, job)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	job := JobManifest{
		Name:      "client-monitor-job",
		Namespace: "default",
		Action:    "processes",
		PodName:   "deploy-dotnet-monitor-77c45ff464-dsgr4",
		PID:       "NO_PID",
		UID:       "NO_UID",
		NameOverride: "NO_NAME",
		Duration:  "NO_DURATION",
		Egress:    "NO_EGRESS_PROVIDER",
		Tags:      "NO_TAG",
	}

	filePath := "job-by-utility.yml"
	err := generateJobManifestFile(job, filePath)
	if err != nil {
		fmt.Println("Error generating job manifest file:", err)
		return
	}

	fmt.Println("Job manifest file generated successfully!")
}
