/*
 * this package is trying to realize ceph         object client with go
 */
package radosgw

type Connection struct {
	user string
	key string
	url string
}

func (con *Connection) create_container(container string) {

}

func (con *Connection) list_containers() (containers List){

}

func (con *Connection) list_container_contents(container string) (contents List){

}

func (con *Connection) delete_container(container string) {

}

func (con *Connection) create_object(object string){

}

func (con *Connection) put_object(object []byte){

}

func (con *Connection) delete_object(object string){

}