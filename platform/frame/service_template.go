package frame

import (
	"fmt"
	"log"
	"net/http"
	"time"
	
	"github.com/gorilla/mux"
)

type ServiceRegister interface {
	RegisterRoutes()[]ServiceRoute
}

type ServiceRoute struct {
	OperatorId        string
	Method      string												// example: GET
	RelativePath     string										// example: /sm-contexts
	HandlerFunc http.HandlerFunc
}

type ServiceTemplate struct {
	ServiceRegister
	serviceName string
	rootPath string													// example: /nsmf-pdusession/v1  最后没有斜杠
	routes []ServiceRoute
	router *mux.Router
}

func (st *ServiceTemplate)configServiceRoute(){
	// Set default route: call Index function
	if st.routes == nil{
		st.routes=[]ServiceRoute{{"Index","GET","/",Index}}
	}
	
	// if ServiceRegister == nil do nothing
	// else add customized route
	if st.ServiceRegister == nil{
		return
	}
	
	for _,route:=range st.RegisterRoutes(){
		st.routes=append(st.routes,route)
	}
	
	st.router = mux.NewRouter().StrictSlash(true)
	for _, route := range st.routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = Logger(handler, route.OperatorId)
		
		st.router.
			Methods(route.Method).
			Path(st.rootPath+route.RelativePath).
			Name(route.OperatorId).
			Handler(handler)
	}
}

// example: StartService(":8080")
func (st *ServiceTemplate)StartService(addr string)error{
	if addr==""{
		addr = ":8080"
	}
	
	// according to ServiceRegister interface provided by concrete service, to config routes
	st.configServiceRoute()
	
	log.Printf("%s service is started.",st.serviceName)
	
	return http.ListenAndServe(addr, st.router)
}

// default function when accessing root path
func Index(w http.ResponseWriter, r *http.Request) {
	_,_=fmt.Fprintf(w, "This is a response to root path!")
}

// 缺省注册一个Index处理函数，对rootPath进行响应
func NewService(name string, rootPath string, sr ServiceRegister)*ServiceTemplate{
	return &ServiceTemplate{
		ServiceRegister: sr,
		serviceName:     name,
		rootPath:        rootPath,
		routes: []ServiceRoute{{"Index","GET","/",Index}},
		router:          nil,
	}
}

func Logger(inner http.Handler, name string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		
		inner.ServeHTTP(w, r)
		
		log.Printf(
			"%s %s %s %s",
			r.Method,
			r.RequestURI,
			name,
			time.Since(start),
		)
	})
}


// Example: 
type PduSessionService struct {
	
}

func (pss PduSessionService)RegisterRoutes()[]ServiceRoute{
	return []ServiceRoute{
		{"PostSmContexts","POST","/sm-contexts",pss.PostSmContexts},
		{"ReleasePduSession", "POST", "/pdu-sessions/{pduSessionRef}/release",pss.PostReleasePduSession},
		}
}
func (pss PduSessionService)PostSmContexts(w http.ResponseWriter, r *http.Request){
	fmt.Println("I receive a request")
	_,_=fmt.Fprintln(w,"This is Pdu Session Create Operator")
}

func (pss PduSessionService)PostReleasePduSession(w http.ResponseWriter, r *http.Request){
	_,_=fmt.Fprintln(w,"This is Pdu Session Release Operator")
}

func NewPduSessionService()*PduSessionService{
	return &PduSessionService{}
}

func DemoServiceTemplate(){
	pduSessionService:= NewService("Pdu Session Management Service", "/nsmf-pdusession/v1", NewPduSessionService())
	log.Fatal(pduSessionService.StartService(":9093"))
}
// Example: End
