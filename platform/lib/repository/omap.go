// omap库体提供一个通用的有序的map数据类型，它的key和value可以是任意值
package repository

import "strings"

type Map struct {
	root *node																								// 树的根
	lessFunc func(key1 interface{},key2 interface{})bool					// 提供比较两个KEY大小的函数
	length int
}

type node struct {
	key,value interface{}
	red bool
	left,right *node								// define left and right leaf node of a tree
}

// 构造函数
// 通用的构造函数，即不管key是什么类型都行，但必须提供两个key比较的函数
func NewMap(lessFunc func(key1 interface{}, key2 interface{})bool)*Map{
	return &Map{lessFunc:lessFunc}
}

// 提供一个key是string的有序Map类型,大小写不敏感
func NewCaseFoldedKeyed()*Map{
	return &Map{lessFunc:func(a,b interface{})bool{
		return strings.ToLower(a.(string))<strings.ToLower(b.(string))}}		// a.(string)把 a转成string类型
}

// 提供一个key是string的有序Map类型,大小写敏感
func NewStringKeyed()*Map{
	return &Map{lessFunc:func(a,b interface{})bool{
		return a.(string)<b.(string)}}																		// a.(string)把 a转成string类型
}

// 提供一个key是Int 的有序Map类型
func NewIntKeyed()*Map{
	return &Map{lessFunc:func(a,b interface{})bool{
		return a.(int)<b.(int)}}																					// a.(int)把 a转成int类型
}

// 提供一个key是Float64的有序Map类型
func NewFloat64Keyed()*Map{
	return &Map{lessFunc:func(a,b interface{})bool{
		return a.(float64)<b.(float64)}}																					// a.(int)把 a转成int类型
}

// 提供一个key是struct的有序Map类型
func NewCustomizedStructureKeyed()*Map{
	return &Map{lessFunc:lessFuncForCustomizedStructure}
}

type CustomizedStructure struct {
	id1 string
	id2 string
	content string
}

func lessFuncForCustomizedStructure(a,b interface{})bool{
	c1,c2:=a.(CustomizedStructure),b.(CustomizedStructure)
	key1,key2:=c1.id1+c1.id2,c2.id1+c2.id2
	return key1<key2
}

// 这个函数可以看出典型的GO语言风格，主函数只关注步骤，具体步骤实现用辅助函数完成，如这里的insert
// 在Map中加入一个key，value对
func (m *Map)Insert(key,value interface{})(inserted bool){
	m.root,inserted=m.insert(m.root,key,value)
	m.root.red=false
	if inserted{
		m.length++
	}
	
	return inserted
}

// 这里实现的是迭代，而不是递归
// 查找一个node
func (m *Map)Find(key interface{})(value interface{},found bool){
	root:=m.root
	for root!=nil{
		if m.lessFunc(key,root.key){
			root = root.left
		}else if m.lessFunc(root.key,key){
			root = root.right
		}else {																											// 等于的比较逻辑是否对所有的类型都有效？
			return root.value,true
		}
	}
	
	return nil,false
}

// 删除一个node
func (m *Map)Delete(key interface{})(deleted bool){
	if m.root!=nil{
		if m.root,deleted=m.remove(m.root,key);m.root!=nil{
			m.root.red = false
		}
	}
	
	if deleted{
		m.length--
	}
	
	return deleted
}

func (m *Map)Len()int{
	return m.length
}

// 遍历这颗树找出每个key value对，然后调用function来处理
func (m *Map)Do(function func(interface{},interface{})){
	m.do(m.root,function)
}

// 左，中，右遍历
func (m *Map)do(root *node,function func(interface{},interface{})){
	if root!=nil{
		m.do(root.left,function)
		function(root.key,root.value)
		m.do(root.right,function)
	}
}

func (m *Map) remove(root *node, key interface{}) (*node, bool) {
	deleted := false
	if m.lessFunc(key, root.key) {
		if root.left != nil {
			if !isRed(root.left) && !isRed(root.left.left) {
				root = moveRedLeft(root)
			}
			root.left, deleted = m.remove(root.left, key)
		}
	} else {
		if isRed(root.left) {
			root = rotateRight(root)
		}
		if !m.lessFunc(key, root.key) && !m.lessFunc(root.key, key) &&
			root.right == nil {
			return nil, true
		}
		if root.right != nil {
			if !isRed(root.right) && !isRed(root.right.left) {
				root = moveRedRight(root)
			}
			if !m.lessFunc(key, root.key) && !m.lessFunc(root.key, key) {
				smallest := first(root.right)
				root.key = smallest.key
				root.value = smallest.value
				root.right = deleteMinimum(root.right)
				deleted = true
			} else {
				root.right, deleted = m.remove(root.right, key)
			}
		}
	}
	return fixUp(root), deleted
}

func fixUp(root *node) *node {
	if isRed(root.right) {
		root = rotateLeft(root)
	}
	if isRed(root.left) && isRed(root.left.left) {
		root = rotateRight(root)
	}
	if isRed(root.left) && isRed(root.right) {
		colorFlip(root)
	}
	return root
}

func deleteMinimum(root *node) *node {
	if root.left == nil {
		return nil
	}
	if !isRed(root.left) && !isRed(root.left.left) {
		root = moveRedLeft(root)
	}
	root.left = deleteMinimum(root.left)
	return fixUp(root)
}

func moveRedLeft(root *node) *node {
	colorFlip(root)
	if root.right != nil && isRed(root.right.left) {
		root.right = rotateRight(root.right)
		root = rotateLeft(root)
		colorFlip(root)
	}
	return root
}

func moveRedRight(root *node) *node {
	colorFlip(root)
	if root.left != nil && isRed(root.left.left) {
		root = rotateRight(root)
		colorFlip(root)
	}
	return root
}

// We do not provide an exported First() method because this is an
// implementation detail.
func first(root *node) *node {
	for root.left != nil {
		root = root.left
	}
	return root
}

// 在node里面插入一个节点
// 这是个牛逼的算法，要参见Robert Sedgewrick的论文了
func (m *Map)insert(root *node,key,value interface{})( *node, bool){
	inserted:=false
	if root==nil{				// 树为空，或者key已经在树里面
		return &node{key,value,true,nil,nil},true
	}
	if isRed(root.left)&&isRed(root.right){
		colorFlip(root)
	}
	if m.lessFunc(key,root.key){
		root.left,inserted=m.insert(root.left,key,value)
	}else if m.lessFunc(root.key,key){																// 体会一下为什么只要提供less函数，而不是< = >三个操作？
		root.right,inserted=m.insert(root.right,key,value)
	}else {																													// key has existed
		root.value = value																						// 相等，直接覆盖原来的值
	}
	if isRed(root.right)&&!isRed(root.left){
		root = rotateLeft(root)
	}
	if isRed(root.left)&&isRed(root.left.left){
		root = rotateRight(root)
	}
	
	return root, inserted
}

func rotateLeft(root *node)*node{
	x:=root.right
	root.right = x.left
	x.left=root
	x.red=root.red
	
	root.red=true
	return x
}
func rotateRight(root *node)*node{
	x:=root.left
	root.left=x.right
	x.right=root
	x.red=root.red
	
	root.red=true
	return x
}

func isRed( root *node)bool{
	return root!=nil && root.red
}

func colorFlip(root *node){
	root.red = !root.red
	if root.left!=nil{
		root.left.red=!root.left.red
	}
	if root.right!=nil{
		root.right.red=!root.right.red
	}
}