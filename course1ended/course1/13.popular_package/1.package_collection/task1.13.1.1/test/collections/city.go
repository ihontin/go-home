package collections

import "sort"

type City struct {
	CityId int
}

type CityCollection struct {
	Items []*City
}

type SearchCallbackCity func(item *City) bool

//grizzly:replaceName New{{.Name}}Collection
func NewCityCollection(items []*City) *CityCollection {
	var collection CityCollection

	collection.Items = items

	return &collection
}

//grizzly:replaceName NewEmpty{{.Name}}Collection
func NewEmptyCityCollection() *CityCollection {
	return &CityCollection{}
}

func (c *CityCollection) Find(callback SearchCallbackCity) *City {
	for _, v := range c.Items {
		if callback(v) == true {
			return v
		}
	}

	return nil
}

func (c *CityCollection) Filter(callback SearchCallbackCity) *CityCollection {
	var newItems []*City

	for _, v := range c.Items {
		if callback(v) == true {
			newItems = append(newItems, v)
		}
	}

	return &CityCollection{Items: newItems}
}

func (c *CityCollection) MapToInt(callback func(item *City) int) []int {
	items := []int{}

	for _, v := range c.Items {
		items = append(items, callback(v))
	}

	return items
}

func (c *CityCollection) MapToString(callback func(item *City) string) []string {
	items := []string{}

	for _, v := range c.Items {
		items = append(items, callback(v))
	}

	return items
}

func (c *CityCollection) Push(item *City) *CityCollection {
	newItems := append(c.Items, item)

	return &CityCollection{Items: newItems}
}

func (c *CityCollection) Shift() *City {
	item := c.Items[0]
	c.Items = c.Items[1:]

	return item
}

func (c *CityCollection) Pop() *City {
	item := c.Items[len(c.Items)-1]
	c.Items = c.Items[:len(c.Items)-1]

	return item
}

func (c *CityCollection) Unshift(item *City) *CityCollection {
	newItems := append([]*City{item}, c.Items...)

	return &CityCollection{Items: newItems}
}

func (c *CityCollection) Len() int {
	return len(c.Items)
}

func (c *CityCollection) Get(index int) (model *City) {
	if index >= 0 && len(c.Items) > index {
		return c.Items[index]
	}

	return model
}

func (c *CityCollection) UniqByCityId() *CityCollection {
	collection := &CityCollection{}

	for _, item := range c.Items {
		searchItem := collection.Find(func(model *City) bool {
			return model.CityId == item.CityId
		})

		if searchItem == nil {
			collection = collection.Push(item)
		}
	}

	return collection
}

type byCityIdAsc []*City

func (a byCityIdAsc) Len() int           { return len(a) }
func (a byCityIdAsc) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byCityIdAsc) Less(i, j int) bool { return a[i].CityId < a[j].CityId }

type byCityIdDesc []*City

func (a byCityIdDesc) Len() int           { return len(a) }
func (a byCityIdDesc) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byCityIdDesc) Less(i, j int) bool { return a[i].CityId > a[j].CityId }

func (c *CityCollection) SortByCityId(mode string) *CityCollection {
	collection := &CityCollection{Items: c.Items}

	if mode == "desc" {
		sort.Sort(byCityIdDesc(collection.Items))
	} else {
		sort.Sort(byCityIdAsc(collection.Items))
	}

	return collection
}

func (c *CityCollection) ForEach(callback func(item *City)) {
	for _, i := range c.Items {
		callback(i)
	}
}
