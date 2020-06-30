package serializer

import "QUZHIYOU/models"

type Ad struct {
	Id uint `json:"id"`
	Name string `json:"name,omitempty"`
	Type string `json:"type,omitempty"`
	Link string `json:"link,omitempty"`
	IsShelves bool `json:"is_shelves,omitempty"`
	Image string `json:"image,omitempty"`
}

func BuildAdSerializer(ad models.Ad)  *Ad{
	return &Ad{
		Id:        ad.ID,
		Name:      ad.Name,
		Type:      ad.Type,
		Link:      ad.Link,
		IsShelves: ad.IsShelves,
		Image:     ad.Image,
	}
}

func BuildAdSerializers(ad []*models.Ad)  (ads []*Ad) {

	for _,v:=range ad{
		ads = append(ads, BuildAdSerializer(*v))
	}
	return

}