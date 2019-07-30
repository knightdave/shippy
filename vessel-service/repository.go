// vessel-service/repository.go
func (repository *VesselRepository) Create(vessel *pb.Vessel) error {
	return repository.collection.Insert(vessel)
}