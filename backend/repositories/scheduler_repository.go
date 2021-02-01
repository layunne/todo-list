package repositories

//type SchedulerRepository interface {
//	Add(schedulerRequest *models.ConfigurableDynamicRequest) error
//	GetById(id string) *models.ConfigurableDynamicRequest
//	GetAllWithPagination(limit int64, page int64) ([]*models.ConfigurableDynamicRequest, *mongopagination.PaginationData)
//	GetAll() []*models.ConfigurableDynamicRequest
//	Remove(id string)
//}
//
//func (r *schedulerRepository) GetAllWithPagination(limit int64, page int64) ([]*models.ConfigurableDynamicRequest, *mongopagination.PaginationData){
//
//	data := &([]*models.ConfigurableDynamicRequest{})
//
//	pag, err := r.mongo.GetAllWithPagination(r.collection, limit, page, data)
//	if err != nil {
//		log.Println("🔴 GetAllWithPagination error: ", err.Error())
//		return nil, nil
//	}
//
//	return *data, pag
//}

//func (r *schedulerRepository) GetAll() []*models.ConfigurableDynamicRequest {
//
//	data := &([]*models.ConfigurableDynamicRequest{})
//	err := r.mongo.GetAll(r.collection, data)
//	if err != nil {
//		log.Println("🔴 GetAll error: ", err.Error())
//		return nil
//	}
//
//	return *data
//}
