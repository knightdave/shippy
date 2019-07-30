func (s *service) Create(ctx context.Context, req *pb.Vessel, res *pb.Response) error {
	if err := s.repository.Create(req); err != nil {
		return err
	}
	res.Vessel = req
	res.Created = true
	return nil
}