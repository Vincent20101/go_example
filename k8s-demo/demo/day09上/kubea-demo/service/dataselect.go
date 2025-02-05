package service
/**

此课程提供者：微信imax882

+微信imax882
办理会员 课程全部免费看

课程清单：https://leaaiv.cn

全网最全 最专业的 一手课程

成立十周年 会员特惠 速来抢购

**/
import (
	corev1 "k8s.io/api/core/v1"
	appsv1 "k8s.io/api/apps/v1"
	"sort"
	"strings"
	"time"
)

//dataSelector用于封装排序、过滤、分页的数据类型
type dataSelector struct {
	GenericDataList []DataCell
	dataSelectorQuery *DataSelectorQuery
}

//DataCell接口，用于各种资源list的类型转换，转换后可以使用dataSelector的排序、过滤、分页
type DataCell interface {
	GetCreation() time.Time
	GetName()     string
}

//DataSelectQuery 定义过滤和分页的属性，过滤：Name，分页：Limit和page
type DataSelectorQuery struct {
	FilterQuery *FilterQuery
	PaginateQuery *PaginateQuery
}

type FilterQuery struct {
	Name string
}

type PaginateQuery struct {
	Limit int
	Page  int
}

//排序，实现自定义结构的排序，需要重写Len、Swap、Less方法
//Len方法用于获取数组长度
func(d *dataSelector) Len() int {
	return len(d.GenericDataList)
}
//Swap方法用于数组中的元素在比较大小后怎么叫唤位置，可定义升降序
//i,j是切片的下标
func(d *dataSelector) Swap(i, j int) {
	d.GenericDataList[i], d.GenericDataList[j] = d.GenericDataList[j], d.GenericDataList[i]
}
//Less方法用于定义数组中元素排序的“大小”的比较方式
func(d *dataSelector) Less(i, j int) bool {
	a := d.GenericDataList[i].GetCreation()
	b := d.GenericDataList[j].GetCreation()
	return b.Before(a)
}

//重写以上3个方法用使用sort.Sort进行排序
func(d *dataSelector) Sort() *dataSelector {
	sort.Sort(d)
	return d
}

//过滤
//Filter方法用于过滤元素，比较元素的Name属性，若包含，则返回
func(d *dataSelector) Filter() *dataSelector {
	//若Name的传参为空，则返回所有元素
	if d.dataSelectorQuery.FilterQuery.Name == "" {
		return d
	}
	//若Name的传参不为空，则返回元素中包含Name的所有元素
	filteredList := []DataCell{}
	for _, value := range d.GenericDataList {
		matched := true
		objName := value.GetName()
		if !strings.Contains(objName, d.dataSelectorQuery.FilterQuery.Name) {
			matched = false
			continue
		}
		if matched {
			filteredList = append(filteredList, value)
		}
	}
	d.GenericDataList = filteredList
	return d
}

//Paginate方法用于数组分页，根据Limit和Page的传参，返回数据
func(d *dataSelector) Paginate() *dataSelector {
	limit := d.dataSelectorQuery.PaginateQuery.Limit
	page := d.dataSelectorQuery.PaginateQuery.Page
	//验证参数合法，若参数不合法，则返回所有数据
	if limit <= 0 || page <= 0 {
		return d
	}
	//定义offset
	//举例：25个元素的切片 limit10
	//page1 start0 end 10
	//page2 start10 end 20
	//page3 start20 end 30
	startIndex := limit * (page - 1)
	endIndex := limit * page
	if len(d.GenericDataList) < endIndex {
		endIndex = len(d.GenericDataList)
	}
	if startIndex > endIndex {
		return d
	}
	d.GenericDataList = d.GenericDataList[startIndex: endIndex]
	return d
}

//定义podCell类型，实现两个方法GetCreation GetName，可进行类型转换
type podCell corev1.Pod

func(p podCell) GetCreation() time.Time {
	return p.CreationTimestamp.Time
}

func(p podCell) GetName() string {
	return p.Name
}

type deploymentCell appsv1.Deployment

func(d deploymentCell) GetCreation() time.Time {
	return d.CreationTimestamp.Time
}

func(d deploymentCell) GetName() string {
	return d.Name
}
