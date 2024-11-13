package epargne

import "github.com/RazanakotoMandresy/go-gin-banks/pkg/common/models"

type getEpargnes struct {
	userUUID string
	h        Handler
}
type getSingleEparne struct {
	userUUID    string
	epargneUUID string
	h           Handler
}

func (s getSingleEparne) singleEconomie() (*models.Epargne, error) {
	var epargne models.Epargne
	res := s.h.DB.First(&epargne, "owner_uuid = ? AND id = ? AND is_economie = true", s.userUUID, s.epargneUUID)
	if res.Error != nil {
		return nil, res.Error
	}
	return &epargne, nil
}
func (s getSingleEparne) singleEpargneNonEconomie() (*models.Epargne, error) {
	var epargne models.Epargne
	res := s.h.DB.First(&epargne, "owner_uuid = ? AND id = ? AND is_economie = false", s.userUUID, s.epargneUUID)
	if res.Error != nil {
		return nil, res.Error
	}
	return &epargne, nil
}
func (g getEpargnes) economieCase() ([]models.Epargne, error) {
	var epargnes []models.Epargne
	res := g.h.DB.Where("owner_uuid = ? AND is_economie = true", g.userUUID).Find(&epargnes)
	if res.Error != nil {
		return nil, res.Error
	}
	return epargnes, nil
}
func (g getEpargnes) nonEconomieCase() ([]models.Epargne, error) {
	var epargnes []models.Epargne
	res := g.h.DB.Where("owner_uuid = ? AND is_economie = false", g.userUUID).Find(&epargnes)
	if res.Error != nil {
		return nil, res.Error
	}
	return epargnes, nil
}
