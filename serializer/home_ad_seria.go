package serializer

import "QUZHIYOU/models"

type Ad struct {
	Id uint `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
	Link string `json:"link"`
	IsShelves bool `json:"is_shelves"`
	Image string `json:"image"`
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