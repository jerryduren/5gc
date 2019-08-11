package main

import (
	"github.com/emicklei/go-restful"
	"github.com/jerryduren/5gc/platform/lib/repository"
)

var SmContext  = repository.NewEntityRepository()


	EntityRepository

func main() {
	smcServcie:=new(restful.WebService)
	smcServcie.
		Path("/nsmf-pdusession/v1").
		ApiVersion("v1").
		Consumes("multipart/related",restful.MIME_JSON).
		Produces("multipart/related",restful.MIME_JSON)
	
	
		
	
}
