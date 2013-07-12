package main

import (
	"github.com/emicklei/go-restful"
	"labix.org/v2/mgo"
	"log"
)

type MetaResource struct {
	session *mgo.Session
}

func (m MetaResource) Register() {
	ws := new(restful.WebService)
	ws.Path("/databases/{database}")
	ws.Consumes("*/*")
	ws.Route(ws.GET("/collections").To(m.getAllCollectionNames))
	restful.Add(ws)
}

func (m MetaResource) getAllCollectionNames(req *restful.Request, resp *restful.Response) {
	dbname := req.PathParameter("database")
	names, err := m.session.DB(dbname).CollectionNames()
	if err != nil {
		log.Printf("[mora] error:%v", err)
		resp.WriteError(500, err)
		return
	}
	resp.WriteEntity(names)
}