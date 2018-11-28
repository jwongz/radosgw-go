/*
 * this package is trying to realize ceph's object client with go
 */
package radosgw

import (
	"github.com/astaxie/beego/httplib"
)

type Connection struct {
	url   string
	token string
}

func New(user string, key string, auth_url string) *Connection {
	b := httplib.Get(auth_url)
	b.Header("X-Auth-User", user)
	b.Header("X-Auth-Key", key)
	res, err := b.DoRequest()
	if err != nil {
		return nil
	}

	return &Connection{res.Header.Get("x-storage-url"), res.Header.Get("x-storage-token")}
}

func (conn *Connection) Create_container(container string) {
	url := conn.url + "/" + container
	b := httplib.Put(url)
	b.Header("X-Auth-Token", conn.token)
	b.DoRequest()
}

func (conn *Connection) List_containers(v interface{}) error {
	b := httplib.Get(conn.url)
	b.Header("X-Auth-Token", conn.token)

	return b.ToJSON(v)
}

func (conn *Connection) List_container_contents(container string, v interface{}) error {
	url := conn.url + container
	b := httplib.Get(url)
	b.Header("X-Auth-Token", conn.token)

	return b.ToJSON(v)
}

func (conn *Connection) Delete_container(container string) {
	url := conn.url + "/" + container
	b := httplib.Delete(url)
	b.Header("X-Auth-Token", conn.token)

	b.DoRequest()
}

// Create or update an object
func (conn *Connection) Put_object(container string, object string, data []byte) {
	url := conn.url + "/" + container + "/" + object
	b := httplib.Put(url)
	b.Header("X-Auth-Token", conn.token)
	b.Body(data)

	b.DoRequest()
}

func (conn *Connection) Delete_object(container string, object string) {
	url := conn.url + "/" + container + "/" + object
	b := httplib.Delete(url)
	b.Header("X-Auth-Token", conn.token)

	b.DoRequest()
}
