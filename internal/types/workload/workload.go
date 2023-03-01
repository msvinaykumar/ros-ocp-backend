package workload

import "database/sql/driver"

type WorkloadType string

const (
	Deployment            WorkloadType = "deployment"
	Deploymentconfig      WorkloadType = "deploymentconfig"
	Replicaset            WorkloadType = "replicaset"
	Replicationcontroller WorkloadType = "replicationcontroller"
	Statefulsets          WorkloadType = "statefulsets"
)

func (p *WorkloadType) Scan(value interface{}) error {
	*p = WorkloadType(value.([]byte))
	return nil
}

func (p WorkloadType) Value() (driver.Value, error) {
	return string(p), nil
}
