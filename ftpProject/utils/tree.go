package utils

import "strconv"

type TreeMenuStruct struct {
	Menu_icon      string
	Menu_id        string
	Menu_name      string
	Menu_route     string
	Menu_up_id     string
	Menu_Dept      string
	Menu_leaf_flag string
	Menu_img       string
	Menu_color     string
}

func GetJSONMenuTree(node []TreeMenuStruct, id string, d int, result *[]TreeMenuStruct) {

	for _, val := range node {
		if val.Menu_up_id == id {

			var oneline TreeMenuStruct
			oneline.Menu_icon = val.Menu_icon
			oneline.Menu_id = val.Menu_id
			oneline.Menu_name = val.Menu_name
			oneline.Menu_up_id = val.Menu_up_id
			oneline.Menu_route = val.Menu_route
			oneline.Menu_leaf_flag = val.Menu_leaf_flag
			oneline.Menu_Dept = strconv.Itoa(d)
			oneline.Menu_img = val.Menu_img
			oneline.Menu_color = val.Menu_color
			*result = append(*result, oneline)
			GetJSONMenuTree(node, val.Menu_id, d+1, result)
		}
	}
}

//type TreeStruct struct {
//	Id    string
//	Name  string
//	Up_id string
//}

//type TreeResetSet struct {
//	Id    string
//	Name  string
//	Up_id string
//	Dept  string
//	Node  string
//}

//type DirTree struct {
//	FileId   string
//	FileName string
//	FileUpId string
//	Dept     int
//	IsDir    bool
//}
//func GetMultCheckBoxTree(node []TreeStruct, id string, d int, result *string) {

//	for _, val := range node {
//		if val.Up_id == id {
//			flag := 0
//			for _, v := range node {
//				if v.Up_id == val.Id {
//					flag = 1
//				}
//			}
//			if flag == 0 {
//				*result += "<li id='" + val.Id + "' data-dept='" + strconv.Itoa(d) + "' data-node='0' style='padding-left:" + strconv.Itoa(d*20) + "px'><input type='checkbox' onchange='mascheckbox(this)'/>" + val.Name + "</li>"
//			} else {
//				*result += "<li id='" + val.Id + "' data-dept='" + strconv.Itoa(d) + "' data-node='1' style='padding-left:" + strconv.Itoa(d*20) + "px'><input type='checkbox' onchange='mascheckbox(this)'/><span onclick='mascheckboxshow(this)' style='cursor:pointer'>" + val.Name + "<i class='caret'></i></span></li>"
//			}

//			GetMultCheckBoxTree(node, val.Id, d+1, result)
//		}
//	}
//}

//func GetMultCheckBoxUpTree(node []TreeStruct, id string, d int, result *string) {

//	for _, val := range node {
//		if val.Up_id == id {
//			flag := 0
//			for _, v := range node {
//				if v.Up_id == val.Id {
//					flag = 1
//				}
//			}
//			if flag == 0 {
//				*result += "<li id='" + val.Id + "' data-dept='" + strconv.Itoa(d) + "' data-node='0' style='padding-left:" + strconv.Itoa(d*20) + "px'><input type='checkbox' onchange='mascheckboxuptree(this)'/>" + val.Name + "</li>"
//			} else {
//				*result += "<li id='" + val.Id + "' data-dept='" + strconv.Itoa(d) + "' data-node='1' style='padding-left:" + strconv.Itoa(d*20) + "px'><input type='checkbox' onchange='mascheckboxuptree(this)'/><span onclick='mascheckboxshow(this)' style='cursor:pointer'>" + val.Name + "<i class='caret'></i></span></li>"
//			}

//			GetMultCheckBoxUpTree(node, val.Id, d+1, result)
//		}
//	}
//}

//func GetLeafTree(node []TreeStruct, id string, d int, result *string) {

//	for _, val := range node {
//		if val.Up_id == id {
//			flag := 0
//			for _, v := range node {
//				if v.Up_id == val.Id {
//					flag = 1
//				}
//			}
//			if flag == 0 {
//				*result += "<li id='" + val.Id + "' data-dept='" + strconv.Itoa(d) + "' data-node='0' onclick='masleafcheck(this)' style='padding-left:" + strconv.Itoa(d*20) + "px;cursor:pointer'>" + val.Name + "</li>"
//			} else {
//				*result += "<li id='" + val.Id + "' data-dept='" + strconv.Itoa(d) + "' data-node='1' onclick='masleafcheck(this)' style='padding-left:" + strconv.Itoa(d*20) + "px;cursor:pointer'>" + val.Name + "<i class='caret'></i></li>"
//			}

//			GetLeafTree(node, val.Id, d+1, result)
//		}
//	}
//}

//func GetNodeTree(node []TreeStruct, id string, d int, result *string) {

//	for _, val := range node {
//		if val.Up_id == id {
//			flag := 0
//			for _, v := range node {
//				if v.Up_id == val.Id {
//					flag = 1
//				}
//			}
//			if flag == 0 {
//				*result += "<li id='" + val.Id + "' data-dept='" + strconv.Itoa(d) + "' data-node='0' style='padding-left:" + strconv.Itoa(d*20) + "px'><i style='cursor:pointer' onclick='masnodecheck(this)'>" + val.Name + "</i></li>"
//			} else {
//				*result += "<li id='" + val.Id + "' data-dept='" + strconv.Itoa(d) + "' data-node='1' style='padding-left:" + strconv.Itoa(d*20) + "px'><span class='glyphicon glyphicon-plus-sign' onclick='masnodeshow(this)' ></span><i style='cursor:pointer' onclick='masnodecheck(this)'>" + val.Name + "</i></li>"
//			}

//			GetNodeTree(node, val.Id, d+1, result)
//		}
//	}
//}

//func GetJSONTree(node []TreeResetSet, id string, d int, result *[]TreeResetSet) {

//	for _, val := range node {
//		if val.Up_id == id {

//			var oneline TreeResetSet
//			oneline.Id = val.Id
//			oneline.Name = val.Name
//			oneline.Up_id = val.Up_id
//			oneline.Dept = strconv.Itoa(d)
//			oneline.Node = val.Node
//			*result = append(*result, oneline)
//			GetJSONTree(node, val.Id, d+1, result)
//		}
//	}
//}

//func GetDirTree(name string, upDir string, d int, result *[]DirTree) {
//	files, err := ioutil.ReadDir(name)
//	if err != nil {
//		panic(err)
//		return
//	}
//	v := runtime.GOOS
//	for _, val := range files {
//		if val.IsDir() == true {
//			var onelne DirTree
//			switch v {
//			case "windows":
//				onelne.FileId = name + "/" + strings.Replace(val.Name(), "\\", "/", -1)
//			case "linux":
//				onelne.FileId = name + "/" + val.Name()
//			default:
//				onelne.FileId = name + "/" + val.Name()
//			}

//			onelne.FileName = val.Name()
//			onelne.FileUpId = name
//			onelne.IsDir = val.IsDir()
//			onelne.Dept = d
//			*result = append(*result, onelne)
//			switch v {
//			case "windows":
//				GetDirTree(name+"/"+strings.Replace(val.Name(), "\\", "/", -1), name, d+1, result)
//			case "linux":
//				GetDirTree(name+"/"+val.Name(), name, d+1, result)
//			default:
//				GetDirTree(name+"/"+val.Name(), name, d+1, result)
//			}

//		} else {
//			var onelne DirTree
//			switch v {
//			case "windows":
//				onelne.FileId = name + "/" + strings.Replace(val.Name(), "\\", "/", -1)
//			case "linux":
//				onelne.FileId = name + "/" + val.Name()
//			default:
//				onelne.FileId = name + "/" + val.Name()
//			}
//			onelne.FileName = val.Name()
//			onelne.FileUpId = name
//			onelne.IsDir = val.IsDir()
//			onelne.Dept = d
//			*result = append(*result, onelne)
//		}
//	}
//}
