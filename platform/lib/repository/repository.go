package repository

import (
	"errors"
	"fmt"
	
)

type Repository interface {
	// Add
	Add(value Entity)error
	AddAll([]Entity)error
	
	// Find
	FindById(key interface{})(value Entity,err error)
	// FindAll()([]Entity,error)
	// FindAllById([]string)([]Entity,error)
	
	// Delete
	DeleteById(interface{})
	DeleteAll()
	DeleteAllById([]interface{})

	 // other
	 ExistsById(key interface{})bool
	
	// Count
	Count() int
}

// 需要保存在EntityRepository里面的Entity需要实现这个接口
type Entity interface {
	GetKey()interface{}
	String()string
	Less(e1, e2 interface{})bool		// 这个接口是提供GetKey得到的值类型的比较方法，不是Entity的比较方法
}

type EntityRepository struct {
	*Map
}


func NewEntityRepository(entity Entity)*EntityRepository{
	return &EntityRepository{NewMap(entity.Less)}
}


func (er *EntityRepository)Add(e Entity)error{
	if e==nil{
		return nil
	}
	
	if er.Insert(e.GetKey(),e){
		return nil
	}
	return errors.New(fmt.Sprintf("Insert Entity %v failure",e.GetKey()))
}

// 事务操作，只有所有的元素添加成功才算成功
func (er *EntityRepository)AddAll(es[]Entity)(err error){
	if len(es)==0{
		return nil
	}
	for _,e:=range es{
		if failure:=er.Add(e);failure!=nil {
			err = failure																		// r任意一个元素添加失败就失败，那要不要删除已经添加成功的元素？
		}
	}
	
	if err!=nil{
		for _,e:=range es{
			er.DeleteById(e.GetKey())
		}
	}
	
	return
}

func (er *EntityRepository)FindById(key interface{})(Entity, error){
	if key ==nil{
		return nil,nil
	}
	value,found:=er.Find(key)
	if !found{
		return nil,errors.New(fmt.Sprintf("Not found entity of key = %v",key))
	}
	
	return value.(Entity),nil
}

// 只要找打一个元素就算成功，不保证全部能找到
func (er *EntityRepository)FindAllById(keys []interface{})(values []Entity,err error){
	if len(keys)==0{
		return nil,nil
	}
	
	values=make([]Entity,16)
	for _,key := range keys {
		value, hasErr := er.FindById(key)
		if hasErr != nil {
			values = append(values, value)
			err=nil																									// 只要找到一个err就返回nil
		}
	}
	
	return values,err
}


// implement String() string
func (er *EntityRepository)String()string{
	values:=make([]interface{},0)
	er.Do(func(key interface{},value interface{}){values=append(values,value)
	})
	
	result:=""
	for _,value :=range values{
		result = result+fmt.Sprintf("%v\n",value)
	}
	
	return result
}

// Delete
func (er *EntityRepository)DeleteById(key interface{}){
	if key != nil{
		er.Delete(key)
	}
}

func (er *EntityRepository)DeleteAll(){
	keys:=make([]interface{},0)
	er.Do(func(key interface{},value interface{}){keys=append(keys,key)})			// 找出所有的keys
	if len(keys)>0{
		er.DeleteAllById(keys)
	}
}

func (er *EntityRepository)DeleteAllById(keys []interface{}){
	if len(keys)==0{
		return
	}
	for _,key:=range keys{
		er.DeleteById(key)
	}
}

// other
func (er *EntityRepository)ExistsById(key interface{})bool{
	if key ==nil{
		return false
	}
	if _,found:=er.Find(key);found{
		return true
	}
	
	return false
}

// Count
func (er *EntityRepository)Count() int{
	return er.Len()
}
