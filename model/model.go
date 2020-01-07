package model

import "time"

type StorageInfo struct {
	Id        uint64 `json:"id"`
	Name      string `json:"name"`
	Path      string `json:"path"`
	Host      string `json:"host"`
	Port      uint   `json:"port"`
	Inventory string `json:"inventory"`
	Playbook  string `json:"playbook"`
	Role      string `json:"role"`
}

type CabinetTypeInfo struct {
	Id      int64     `json:"id"`      //ID
	Code    string    `json:"code"`    //柜子型号
	Name    string    `json:"name"`    //型号名称
	Vol     float64   `json:"vol"`     //额定充电电压
	Elastic float64   `json:"elastic"` //额定充电电流
	Size    int       `json:"size"`    //仓位数
	Remark  string    `json:"remark"`  //备注说明
	Ctime   time.Time `json:"ctime"`   //创建时间
	Utime   time.Time `json:"utime"`   //更新时间
}


type BatteryType struct {
	Id     int64     `json:"id"`
	Code   string    `json:"code"`   //电池型号
	Name   string    `json:"name"`   //型号名称
	Vol    float64   `json:"vol"`    //电压
	Ah     float64   `json:"ah"`     //安时
	Cap    float64   `json:"cap"`    //容量
	Remark string    `json:"remark"` //备注说明
	Ctime  time.Time `json:"ctime"`  //创建时间
	Utime  time.Time `json:"utime"`  //更新时间
}


type AnsCommand struct {
	Id   uint   `json:"id"`
	Pid  uint   `json:"pid"`
	Name string `json:"name"`
	Code string `json:"code"`
}

type ClusterInfo struct {
	Id        uint64 `json:"id"`
	Name      string `json:"name"`
	Code      string `json:"code"`
	Path      string `json:"path"`
	Host      string `json:"host"`
	Port      uint   `json:"port"`
	Inventory string `json:"inventory"`
	Playbook  string `json:"playbook"`
	Vars      string `json:"vars"`
	Role      string `json:"role"`
	Command   string `json:"command" gorm:"-"`
	Test      bool   `json:"test" gorm:"-"`
	ShowUrl   string `json:"show_url" gorm:"-"`
}

type User struct {
	ID             int64  `json:"id" form:"id" gorm:"AUTO_INCREMENT"`
	Loginname      string `json:"loginname" gorm:"loginname" form:"loginname"`
	Username       string `json:"username" gorm:"username"`
	Realname       string `json:"realname" gorm:"realname"`
	Email          string `json:"email"`
	Disabled       int    `json:"disabled"`
	Mobile         string `json:"mobile"`
	Password       string `json:"password"`
	Salt           string `json:"salt"`
	Token          string `json:"token"`
	LoginSessionId string `json:"loginSessionId" gorm:"column:loginSessionId"`
}

// 租赁类型
type LeaseType struct {
	ID    int64     `json:"id"`    //ID
	Rent  float64   `json:"rent"`  //租金
	Term  int       `json:"term"`  //租期
	Time  int       `json:"time"`  //次数
	Ctime time.Time `json:"ctime"` //创建时间
}
type Cluster struct {
	ID        int        `gorm:"primary_key"` //ID
	Name      string     `json:"name"` //Name
	TCP       int        `json:"tcp"` //TCP
	HTTP      int        `json:"http"` //HTTP
	Cluster   string     `json:"cluster"` //Cluster
	Version   string     `json:"version"` //Version
	Ctime     int        `json:"ctime"` //Ctime
	Instances []Instance `json:"instances"` //Instances
}

type Instance struct {
	Name     string `json:"name"`
	Path     string `json:"path"`
	DataPath string `json:"dataPath"`
}

type Source struct {
	ID      int64     `json:"id" gorm:"AUTO_INCREMENT;primary"` //ID
	Name    string    `json:"name"`                             //名称
	Version int       `json:"version"`                          //版本
	Path    string    `json:"path"`                             //路径
	CTime   time.Time `json:"ctime" gorm:"column:ctime"`        //创建时间
}


// 租赁换电
type LeaseCharging struct {
	ID     int64     `json:"id" gorm:"AUTO_INCREMENT;primary_key"` //ID
	IMEI   string    `json:"imei"`                                 //IMEI
	MID    int64     `json:"mid" gorm:"column:mid"`                //MID
	Status int8      `json:"status"`                               //租赁换电状态
	CTime  time.Time `json:"ctime" gorm:"column:ctime"`            //创建时间
}

// project project config
type Page struct {
	Name    string
	MName   string
	Columns []*Column
	Path    string
}

type Column struct {
	Label      string `json:"label"`
	Key        string `json:"key"`
	CanModify  bool   `json:"can_modify"`
	IgnoreEdit bool
	IgnoreShow bool
	IgFilter   bool
	IsID       bool
}

type Cabinet struct {
	ID     int64  `json:"id" gorm:"AUTO_INCREMENT;primary_key" s2a:"ignoreEdit"` //编号
	Name   string `json:"name"`                                                  // 名称
	Addr   string `json:"addr"`                                                  // 地址
	Desc   string `json:"desc"`                                                  // 介绍
	Lon    int64  `json:"lon"`                                                   // 经度
	Lat    int64  `json:"lat"`                                                   // 纬度
	Img    string `json:"img" s2a:"ignoreShow;igFilter"`                         // 图片
	Status string `json:"status"`                                                // 启用状态
	CTime  int64  `json:"ctime" gorm:"column:ctime" s2a:"ignoreEdit"`            // 创建时间
}

type Role struct {
	ID         int64  `json:"id" s2a:"ignoreEdit"`                   //编号
	ParentId   int64  `json:"parent_id"`                             //父角色id
	Status     string `json:"status"`                                //状态
	CreateTime int64  `json:"create_time" s2a:"ignoreEdit;igFilter"` //创建时间
	UpdateTime int64  `json:"update_time" s2a:"ignoreEdit;igFilter"` //更新时间
	ListOrder  string `json:"list_order"`                            //排序
	Name       string `json:"name"`                                  //角色名称
	Remark     string `json:"remark"`                                //备注
	Code       string `json:"code"`                                  //编码
	Devices    string `json:"devices"`                               //设备
}

func NewColumn(label, key string) Column {
	return NewColumnAll(label, key, true)
}
func NewColumnAll(label, key string, canModify bool) Column {
	return Column{Label: label, Key: key, CanModify: canModify, IgnoreEdit: false, IgnoreShow: false}

}

// RealnameApply is.
type RealnameApply struct {
	ID       int64     `json:"id" gorm:"AUTO_INCREMENT,primary_key"` //ID
	MID      int64     `json:"mid"`                                  //MID
	OMID     int64     `json:"omid"`                                 //OMID
	Realname string    `json:"realname"`                             //姓名
	Card     string    `json:"card"`                                 //身份证
	CTime    time.Time `json:"ctime" gorm:"column:ctime"`            //创建时间
	UTime    time.Time `json:"utime" gorm:"column:mtime"`            //更新时间
}

// 租赁申请
type LeaseApply struct {
	ID     int64     `json:"id" gorm:"AUTO_INCREMENT;primary_key"` //ID
	MID    int64     `json:"mid" gorm:"column:mid"`                //申请人
	OMID   int64     `json:"omid" gorm:"column:omid"`              //审核人
	IMEI   string    `json:"imei"`                                 //IMEI
	Status int       `json:"status"`                               //申请状态
	Type   int       `json:"type"`                                 //租赁类型
	Term   int       `json:"term"`                                 //租期
	CTime  time.Time `json:"ctime" gorm:"column:ctime"`            //创建时间
	UTime  time.Time `json:"utime" gorm:"column:utime"`            //更新时间
}

// 租赁
type Lease struct {
	ID     int64     `json:"id" gorm:"AUTO_INCREMENT;primary_key"` //ID
	IMEI   string    `json:"imei"`                                 //IMEI
	MID    int64     `json:"mid" gorm:"column:mid"`                //MID
	Type   int       `json:"type"`                                 //租赁类型
	Term   int       `json:"term"`                                 //租期
	Status int       `json:"status"`                               //租赁状态
	CTime  time.Time `json:"ctime" gorm:"column:ctime"`            //创建时间
	UTime  time.Time `json:"utime" gorm:"column:utime"`            //更新时间
}
