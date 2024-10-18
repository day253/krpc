package apollo

import "fmt"

const (
	KConstGlobal = "GLOBAL"
)

type OrganizationIndex struct {
	globalAppIdChannelIndex []int
	appIdChannelIndex       map[string]map[string][]int
}

func (i *OrganizationIndex) add(appId, channel string, count int) {
	if appId != "" && channel != "" {
		if _, exists := i.appIdChannelIndex[appId]; !exists {
			i.appIdChannelIndex[appId] = make(map[string][]int)
		}
		if _, exists := i.appIdChannelIndex[appId][channel]; !exists {
			i.appIdChannelIndex[appId][channel] = []int{}
		}
		i.appIdChannelIndex[appId][channel] = append(i.appIdChannelIndex[appId][channel], count)
	} else {
		i.globalAppIdChannelIndex = append(i.globalAppIdChannelIndex, count)
	}
}

func (i *OrganizationIndex) query(appId, channel string) []int {
	if appId != "" && channel != "" {
		if channelMap, exists := i.appIdChannelIndex[appId]; exists {
			if list, exists := channelMap[channel]; exists {
				list = append(list, i.globalAppIdChannelIndex...)
				return list
			}
		}
	}
	return i.globalAppIdChannelIndex
}

func (i *OrganizationIndex) stat() map[string]interface{} {
	res := make(map[string]interface{})
	res[KConstGlobal] = len(i.globalAppIdChannelIndex)
	for appId, channelMap := range i.appIdChannelIndex {
		for channel, list := range channelMap {
			res[fmt.Sprintf("%s_%s", appId, channel)] = len(list)
		}
	}
	return res
}

func NewOrganizationIndex() *OrganizationIndex {
	return &OrganizationIndex{
		globalAppIdChannelIndex: []int{},
		appIdChannelIndex:       make(map[string]map[string][]int),
	}
}

type GlobalOrganizationIndex struct {
	globalSceneIdIndex []int
	sceneIdIndex       map[string][]int
}

func (i *GlobalOrganizationIndex) add(sceneId string, count int) {
	if sceneId != "" {
		if _, exists := i.sceneIdIndex[sceneId]; !exists {
			i.sceneIdIndex[sceneId] = []int{}
		}
		i.sceneIdIndex[sceneId] = append(i.sceneIdIndex[sceneId], count)
	} else {
		i.globalSceneIdIndex = append(i.globalSceneIdIndex, count)
	}
}

func (i *GlobalOrganizationIndex) query(sceneId string) []int {
	if sceneId != "" {
		if list, exists := i.sceneIdIndex[sceneId]; exists {
			list = append(list, i.globalSceneIdIndex...)
			return list
		}
	}
	return i.globalSceneIdIndex
}

func (i *GlobalOrganizationIndex) stat() map[string]interface{} {
	res := make(map[string]interface{})
	res[KConstGlobal] = len(i.globalSceneIdIndex)
	for sceneId, list := range i.sceneIdIndex {
		res[sceneId] = len(list)
	}
	return res
}

func NewGlobalOrganizationIndex() *GlobalOrganizationIndex {
	return &GlobalOrganizationIndex{
		globalSceneIdIndex: []int{},
		sceneIdIndex:       make(map[string][]int),
	}
}

type ProductIndex struct {
	globalOrganizationIndex *GlobalOrganizationIndex
	organizationIndex       map[string]*OrganizationIndex
}

func (i *ProductIndex) add(organization, sceneId, appId, channel string, count int) {
	if organization == KConstGlobal {
		i.globalOrganizationIndex.add(sceneId, count)
	} else {
		if _, exists := i.organizationIndex[organization]; !exists {
			i.organizationIndex[organization] = NewOrganizationIndex()
		}
		i.organizationIndex[organization].add(appId, channel, count)
	}
}

func (i *ProductIndex) query(organization, sceneId, appId, channel string) []int {
	res := i.globalOrganizationIndex.query(sceneId)
	if organizationIndex, exists := i.organizationIndex[organization]; exists {
		res = append(res, organizationIndex.query(appId, channel)...)
	}
	return res
}

func (i *ProductIndex) stat() map[string]interface{} {
	res := make(map[string]interface{})
	res[KConstGlobal] = i.globalOrganizationIndex.stat()
	for organization, organizationIndex := range i.organizationIndex {
		res[organization] = organizationIndex.stat()
	}
	return res
}

func NewProductIndex() *ProductIndex {
	return &ProductIndex{
		globalOrganizationIndex: NewGlobalOrganizationIndex(),
		organizationIndex:       make(map[string]*OrganizationIndex),
	}
}
