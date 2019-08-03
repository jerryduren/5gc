package frame

import (
	"log"
	"net/http"
	
	"github.com/emicklei/go-restful"
)

type IConfigService interface {
	ConfigService()
}
type ServiceTemplate struct {
	configService IConfigService // 子类要提供注册PATH/METHOD/PARAM等对应的处理函数的方法
	ws            *restful.WebService
}

func (s *ServiceTemplate) createWebService()*restful.WebService {
	s.ws = &restful.WebService{}

	return s.ws
}

// Usage: s.StartWebService(":9090")
func (s *ServiceTemplate) StartWebService(address string) {
	s.createWebService()
	s.configService.ConfigService()			// 子类需要在这个接口函数里面提供必要的Webservice设置，请参考example
	restful.DefaultContainer.Add(s.ws	)
	
	log.Println("start listening on ",address)
	log.Fatal(http.ListenAndServe(address, nil))
}

func NewNfService()*ServiceTemplate{
	return &ServiceTemplate{configService:new(AddrMngService)}
}

// example
type AddrMngService struct {}
//  必须在这个方法里面设置service的相关属性
//  端口在启动方法里面提供
func (*AddrMngService)ConfigService(){
	ws:=&restful.WebService{}
	ws.
		Path("/addrmng").
		Consumes(restful.MIME_XML, restful.MIME_JSON).
		Produces(restful.MIME_JSON, restful.MIME_XML) // you can specify this per route as well
	
	ws.Route(ws.GET("/").To(GetAddr))
	
	// 以下是设置参考
	//tags := []string{"users"}
	//
	//ws.Route(ws.GET("/").To(u.findAllUsers).
	//	// docs
	//	Doc("get all users").
	//	Metadata(restfulspec.KeyOpenAPITags, tags).
	//	Writes([]User{}).
	//	Returns(200, "OK", []User{}))
	//
	//ws.Route(ws.GET("/{user-id}").To(u.findUser).
	//	// docs
	//	Doc("get a user").
	//	Param(ws.PathParameter("user-id", "identifier of the user").DataType("integer").DefaultValue("1")).
	//	Metadata(restfulspec.KeyOpenAPITags, tags).
	//	Writes(User{}). // on the response
	//	Returns(200, "OK", User{}).
	//	Returns(404, "Not Found", nil))
	//
	//ws.Route(ws.PUT("/{user-id}").To(u.updateUser).
	//	// docs
	//	Doc("update a user").
	//	Param(ws.PathParameter("user-id", "identifier of the user").DataType("string")).
	//	Metadata(restfulspec.KeyOpenAPITags, tags).
	//	Reads(User{})) // from the request
	//
	//ws.Route(ws.PUT("").To(u.createUser).
	//	// docs
	//	Doc("create a user").
	//	Metadata(restfulspec.KeyOpenAPITags, tags).
	//	Reads(User{})) // from the request
	//
	//ws.Route(ws.DELETE("/{user-id}").To(u.removeUser).
	//	// docs
	//	Doc("delete a user").
	//	Metadata(restfulspec.KeyOpenAPITags, tags).
	//	Param(ws.PathParameter("user-id", "identifier of the user").DataType("string")))
}

func GetAddr(request *restful.Request, response *restful.Response){
	response.StatusCode()=200
}


type IA interface {
	Config()
}

type Template{
	Ia IA
	a int
}

type Cl struct {

}

func (Cl)Config(){

}