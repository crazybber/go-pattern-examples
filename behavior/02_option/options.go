package option

import (
	"fmt"
)

//Options is key struct,关键数据结构，聚合所有外部可传入的参数
type Options struct {
	UID     int
	GID     int
	Flags   int
	Company string
	Gender  bool //is male
}

//Option func is key func
type Option func(*Options)

//UID set User ID
func UID(userID int) Option {
	return func(args *Options) {
		args.UID = userID
	}
}

//GID set User Group ID
func GID(groupID int) Option {
	return func(args *Options) {
		args.GID = groupID
	}
}

//Company set Company Name
func Company(cname string) Option {
	return func(args *Options) {
		args.Company = cname
	}
}

//Gender for male or female
func Gender(gender bool) Option {
	return func(args *Options) {
		args.Gender = gender
	}
}

//Introduce someone
func Introduce(name string, setters ...Option /*传入闭包设置函数*/) {
	// Default Options
	args := &Options{
		UID:     0,
		GID:     0,
		Company: "",
		Gender:  true,
	}
	//模式的重点体现在这里，通过外部传入的闭包函数设置内在变量
	for _, setter := range setters {
		setter(args)
	}
	gender := "famale"
	if args.Gender {
		gender = "male"
	}
	fmt.Println("----------------------")
	fmt.Println("im am: ", name, "\nfrom: ", args.Company, "\ngender: ", gender, "\nUID: ", args.UID)
}
