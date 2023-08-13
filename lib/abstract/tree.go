package abstract

type ITree interface {
	GetID() string
	GetParentID() string
	GetName() string //主要是用在label显示
}

//树结构对象
type TreeObject struct {
	ID       string        `json:"id"`
	Name     string        `json:"name"`
	Node     ITree         `json:"node"`
	Children []*TreeObject `json:"children,omitempty"`
}

type ITreeArr[V ITree] []V

//只有一个根节点的
func (arr ITreeArr[V]) ToTree() []*TreeObject {
	dataMap := map[string]*TreeObject{}
	for i := 0; i < len(arr); i++ {
		id := arr[i].GetID()
		dataMap[id] = &TreeObject{
			ID:       id,
			Name:     arr[i].GetName(),
			Node:     arr[i],
			Children: make([]*TreeObject, 0, len(arr)/3),
		}
	}
	//增加一个空的
	root := &TreeObject{
		Node:     nil,
		Children: make([]*TreeObject, 0, len(arr)/3),
	}

	for i := 0; i < len(arr); i++ {
		id := arr[i].GetID()
		parentID := arr[i].GetParentID()
		parent, exist := dataMap[parentID]
		if exist {
			parent.Children = append(parent.Children, dataMap[id])
			dataMap[parentID] = parent
		} else { //不存在的  加入root中
			root.Children = append(root.Children, dataMap[id])
		}
	}
	return root.Children
}
