ceph对象存储没有提供go sdk，就连其他语言的支持，也是复用S3和swift client的SDK，由于S3的协议是xml的，有点耗费时间，所以决定先支持swift的协议：
https://docs.openstack.org/swift/latest/api/object_api_v1_overview.html

## ceph官网的推荐用法：
### PYTHON SWIFT EXAMPLES
#### CREATE A CONNECTION
This creates a connection so that you can interact with the server:
```
import swiftclient
user = 'account_name:username'
key = 'your_api_key'

conn = swiftclient.Connection(
        user=user,
        key=key,
        authurl='https://objects.dreamhost.com/auth',
)
```

#### CREATE A CONTAINER
This creates a new container called my-new-container:
```
container_name = 'my-new-container'
conn.put_container(container_name)
```

### CREATE AN OBJECT
This creates a file hello.txt from the file named my_hello.txt:
```
with open('hello.txt', 'r') as hello_file:
        conn.put_object(container_name, 'hello.txt',
                                        contents= hello_file.read(),
                                        content_type='text/plain')
```
#### LIST OWNED CONTAINERS
This gets a list of containers that you own, and prints out the container name:
```
for container in conn.get_account()[1]:
        print container['name']
The output will look something like this:

mahbuckat1
mahbuckat2
mahbuckat3
```

#### LIST A CONTAINER’S CONTENT
This gets a list of objects in the container, and prints out each object’s name, the file size, and last modified date:
```
for data in conn.get_container(container_name)[1]:
        print '{0}\t{1}\t{2}'.format(data['name'], data['bytes'], data['last_modified'])
The output will look something like this:

myphoto1.jpg 251262  2011-08-08T21:35:48.000Z
myphoto2.jpg 262518  2011-08-08T21:38:01.000Z
```

#### RETRIEVE AN OBJECT¶
This downloads the object hello.txt and saves it in ./my_hello.txt:
```
obj_tuple = conn.get_object(container_name, 'hello.txt')
with open('my_hello.txt', 'w') as my_hello:
        my_hello.write(obj_tuple[1])
```

#### DELETE AN OBJECT
This deletes the object hello.txt:
```
conn.delete_object(container_name, 'hello.txt')
DELETE A CONTAINER
Note The container must be empty! Otherwise the request won’t work!
conn.delete_container(container_name)
```
