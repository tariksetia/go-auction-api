package repo

import (
	"auction/pkg/entity"
	"auction/pkg/offers"
)

//IRepo in memory repo
type IRepo struct {
	m map[string]*offers.Offer
}

//NewInmemRepository create new repository
func NewInmemRepository() *IRepo {
	var m = map[string]*offers.Offer{}
	return &IRepo{
		m: m,
	}
}

//Store a Entity
func (r *IRepo) Insert(a *offers.Offer) (entity.ID, error) {
	r.m[string(a.ID)] = a
	return a.ID, nil
}

//Find a Entity
func (r *IRepo) Find(id entity.ID) (*offers.Offer, error) {
	if r.m[id.String()] == nil {
		return nil, entity.ErrNotFound
	}
	return r.m[id.String()], nil
}

//Search Entity
/*
func (r *IRepo) Search(query string) ([]*offer.Offer, error) {
	var d []*offer.Offer
	for _, j := range r.m {
		if strings.Contains(strings.ToLower(j.Name), query) {
			d = append(d, j)
		}
	}
	if len(d) == 0 {
		return nil, entity.ErrNotFound
	}

	return d, nil
}
*/

//FindAll Entity
func (r *IRepo) FindAll() ([]*offers.Offer, error) {
	var d []*offers.Offer
	for _, j := range r.m {
		d = append(d, j)
	}
	return d, nil
}

//Delete a Entity
func (r *IRepo) Delete(id entity.ID) error {
	if r.m[id.String()] == nil {
		return entity.ErrNotFound
	}
	r.m[id.String()] = nil
	return nil
}
